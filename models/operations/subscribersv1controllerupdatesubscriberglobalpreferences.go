// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/novuhq/novu-go/models/components"
)

type SubscribersV1ControllerUpdateSubscriberGlobalPreferencesRequest struct {
	SubscriberID string `pathParam:"style=simple,explode=false,name=subscriberId"`
	// A header for idempotency purposes
	IdempotencyKey                              *string                                                `header:"style=simple,explode=false,name=idempotency-key"`
	UpdateSubscriberGlobalPreferencesRequestDto components.UpdateSubscriberGlobalPreferencesRequestDto `request:"mediaType=application/json"`
}

func (o *SubscribersV1ControllerUpdateSubscriberGlobalPreferencesRequest) GetSubscriberID() string {
	if o == nil {
		return ""
	}
	return o.SubscriberID
}

func (o *SubscribersV1ControllerUpdateSubscriberGlobalPreferencesRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *SubscribersV1ControllerUpdateSubscriberGlobalPreferencesRequest) GetUpdateSubscriberGlobalPreferencesRequestDto() components.UpdateSubscriberGlobalPreferencesRequestDto {
	if o == nil {
		return components.UpdateSubscriberGlobalPreferencesRequestDto{}
	}
	return o.UpdateSubscriberGlobalPreferencesRequestDto
}

type SubscribersV1ControllerUpdateSubscriberGlobalPreferencesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	UpdateSubscriberPreferenceGlobalResponseDto *components.UpdateSubscriberPreferenceGlobalResponseDto
	Headers                                     map[string][]string
}

func (o *SubscribersV1ControllerUpdateSubscriberGlobalPreferencesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SubscribersV1ControllerUpdateSubscriberGlobalPreferencesResponse) GetUpdateSubscriberPreferenceGlobalResponseDto() *components.UpdateSubscriberPreferenceGlobalResponseDto {
	if o == nil {
		return nil
	}
	return o.UpdateSubscriberPreferenceGlobalResponseDto
}

func (o *SubscribersV1ControllerUpdateSubscriberGlobalPreferencesResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
