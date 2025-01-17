// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
)

type ErrorDto struct {
	// HTTP status code of the error response.
	StatusCode float64 `json:"statusCode"`
	// Timestamp of when the error occurred.
	Timestamp string `json:"timestamp"`
	// The path where the error occurred.
	Path string `json:"path"`
	// A detailed error message.
	Message string `json:"message"`
	// Optional context object for additional error details.
	Ctx map[string]any `json:"ctx,omitempty"`
	// Optional unique identifier for the error, useful for tracking using Sentry and
	//       New Relic, only available for 500.
	ErrorID *string `json:"errorId,omitempty"`
}

var _ error = &ErrorDto{}

func (e *ErrorDto) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
