// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/novuhq/novu-go/models/components"
)

type SubscribersControllerChatAccessOauthRequest struct {
	SubscriberID string `pathParam:"style=simple,explode=false,name=subscriberId"`
	ProviderID   any    `pathParam:"style=simple,explode=false,name=providerId"`
	// HMAC hash for the request
	HmacHash string `queryParam:"style=form,explode=true,name=hmacHash"`
	// The ID of the environment, must be a valid MongoDB ID
	EnvironmentID string `queryParam:"style=form,explode=true,name=environmentId"`
	// Optional integration identifier
	IntegrationIdentifier *string `queryParam:"style=form,explode=true,name=integrationIdentifier"`
}

func (o *SubscribersControllerChatAccessOauthRequest) GetSubscriberID() string {
	if o == nil {
		return ""
	}
	return o.SubscriberID
}

func (o *SubscribersControllerChatAccessOauthRequest) GetProviderID() any {
	if o == nil {
		return nil
	}
	return o.ProviderID
}

func (o *SubscribersControllerChatAccessOauthRequest) GetHmacHash() string {
	if o == nil {
		return ""
	}
	return o.HmacHash
}

func (o *SubscribersControllerChatAccessOauthRequest) GetEnvironmentID() string {
	if o == nil {
		return ""
	}
	return o.EnvironmentID
}

func (o *SubscribersControllerChatAccessOauthRequest) GetIntegrationIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationIdentifier
}

type SubscribersControllerChatAccessOauthResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Headers  map[string][]string
}

func (o *SubscribersControllerChatAccessOauthResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SubscribersControllerChatAccessOauthResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
