// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/novuhq/novu-go/models/components"
)

type SubscribersV1ControllerChatAccessOauthRequest struct {
	SubscriberID string `pathParam:"style=simple,explode=false,name=subscriberId"`
	ProviderID   any    `pathParam:"style=simple,explode=false,name=providerId"`
	// HMAC hash for the request
	HmacHash string `queryParam:"style=form,explode=true,name=hmacHash"`
	// The ID of the environment, must be a valid MongoDB ID
	EnvironmentID string `queryParam:"style=form,explode=true,name=environmentId"`
	// Optional integration identifier
	IntegrationIdentifier *string `queryParam:"style=form,explode=true,name=integrationIdentifier"`
	// A header for idempotency purposes
	IdempotencyKey *string `header:"style=simple,explode=false,name=idempotency-key"`
}

func (o *SubscribersV1ControllerChatAccessOauthRequest) GetSubscriberID() string {
	if o == nil {
		return ""
	}
	return o.SubscriberID
}

func (o *SubscribersV1ControllerChatAccessOauthRequest) GetProviderID() any {
	if o == nil {
		return nil
	}
	return o.ProviderID
}

func (o *SubscribersV1ControllerChatAccessOauthRequest) GetHmacHash() string {
	if o == nil {
		return ""
	}
	return o.HmacHash
}

func (o *SubscribersV1ControllerChatAccessOauthRequest) GetEnvironmentID() string {
	if o == nil {
		return ""
	}
	return o.EnvironmentID
}

func (o *SubscribersV1ControllerChatAccessOauthRequest) GetIntegrationIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationIdentifier
}

func (o *SubscribersV1ControllerChatAccessOauthRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

type SubscribersV1ControllerChatAccessOauthResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Headers  map[string][]string
}

func (o *SubscribersV1ControllerChatAccessOauthResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SubscribersV1ControllerChatAccessOauthResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
