package hooks

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Ensure interface compliance at compile-time
var (
	_ sdkInitHook       = (*NovuHooks)(nil)
	_ beforeRequestHook = (*NovuHooks)(nil)
	_ afterSuccessHook  = (*NovuHooks)(nil)
	_ afterErrorHook    = (*NovuHooks)(nil)
)

type NovuHooks struct {
	mu sync.Mutex
}

// SDKInit implements the sdkInitHook interface
func (h *NovuHooks) SDKInit(baseURL string, client HTTPClient) (string, HTTPClient) {
	// Similar to previous implementation
	return baseURL, client
}

// AfterError implements the afterErrorHook interface
func (h *NovuHooks) AfterError(hookCtx AfterErrorContext, res *http.Response, err error) (*http.Response, error) {
	// If both response and error are nil, return an error
	if res == nil && err == nil {
		return nil, fmt.Errorf("both response and error cannot be nil")
	}
	return res, err
}

// Rest of the methods remain the same as in the previous implementation
// (BeforeRequest, AfterSuccess, generateIdempotencyKey, etc.)

// generateSecureRandomString generates a cryptographically secure random string
func generateSecureRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}

// generateIdempotencyKey creates a unique key
func (h *NovuHooks) generateIdempotencyKey() (string, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	randomStr, err := generateSecureRandomString(9)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d%s", timestamp, randomStr), nil
}

// BeforeRequest modifies request headers before sending
func (h *NovuHooks) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	const (
		authKey        = "Authorization"
		idempotencyKey = "Idempotency-Key"
		apiKeyPrefix   = "ApiKey"
	)

	// Ensure Authorization header is prefixed with ApiKey if needed
	authHeader := req.Header.Get(authKey)
	if authHeader != "" && !strings.HasPrefix(authHeader, apiKeyPrefix) {
		req.Header.Set(authKey, fmt.Sprintf("%s %s", apiKeyPrefix, authHeader))
	}

	// Add idempotency key if not present
	if req.Header.Get(idempotencyKey) == "" {
		key, err := h.generateIdempotencyKey()
		if err != nil {
			return nil, fmt.Errorf("failed to generate idempotency key: %w", err)
		}
		req.Header.Set(idempotencyKey, key)
	}

	return req, nil
}

// AfterSuccess modifies the response after a successful request
func (h *NovuHooks) AfterSuccess(hookCtx AfterSuccessContext, res *http.Response) (*http.Response, error) {
	// Check content type
	contentType := res.Header.Get("Content-Type")
	if res.Body == nil || !strings.Contains(contentType, "application/json") {
		return res, nil
	}

	// Read the body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return res, fmt.Errorf("failed to read response body: %w", err)
	}
	defer res.Body.Close()

	// If body is empty, return original response
	if len(bodyBytes) == 0 {
		return res, nil
	}

	// Try to parse JSON with more robust handling
	var jsonResponse map[string]json.RawMessage
	if err := json.Unmarshal(bodyBytes, &jsonResponse); err != nil {
		// If not JSON, restore original body and return
		res.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		return res, nil
	}

	// Check if response has a 'data' key
	dataBytes, ok := jsonResponse["data"]
	if ok {
		newBody := io.NopCloser(bytes.NewBuffer(dataBytes))
		res.Body = newBody
		res.ContentLength = int64(len(dataBytes))
	}

	return res, nil
}
