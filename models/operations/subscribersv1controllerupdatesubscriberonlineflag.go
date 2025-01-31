// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/novuhq/novu-go/models/components"
)

type SubscribersV1ControllerUpdateSubscriberOnlineFlagRequest struct {
	SubscriberID string `pathParam:"style=simple,explode=false,name=subscriberId"`
	// A header for idempotency purposes
	IdempotencyKey                       *string                                         `header:"style=simple,explode=false,name=idempotency-key"`
	UpdateSubscriberOnlineFlagRequestDto components.UpdateSubscriberOnlineFlagRequestDto `request:"mediaType=application/json"`
}

func (o *SubscribersV1ControllerUpdateSubscriberOnlineFlagRequest) GetSubscriberID() string {
	if o == nil {
		return ""
	}
	return o.SubscriberID
}

func (o *SubscribersV1ControllerUpdateSubscriberOnlineFlagRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *SubscribersV1ControllerUpdateSubscriberOnlineFlagRequest) GetUpdateSubscriberOnlineFlagRequestDto() components.UpdateSubscriberOnlineFlagRequestDto {
	if o == nil {
		return components.UpdateSubscriberOnlineFlagRequestDto{}
	}
	return o.UpdateSubscriberOnlineFlagRequestDto
}

type SubscribersV1ControllerUpdateSubscriberOnlineFlagResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	SubscriberResponseDto *components.SubscriberResponseDto
	Headers               map[string][]string
}

func (o *SubscribersV1ControllerUpdateSubscriberOnlineFlagResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SubscribersV1ControllerUpdateSubscriberOnlineFlagResponse) GetSubscriberResponseDto() *components.SubscriberResponseDto {
	if o == nil {
		return nil
	}
	return o.SubscriberResponseDto
}

func (o *SubscribersV1ControllerUpdateSubscriberOnlineFlagResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
