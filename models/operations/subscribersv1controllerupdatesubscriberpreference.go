// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/novuhq/novu-go/models/components"
)

type SubscribersV1ControllerUpdateSubscriberPreferenceRequest struct {
	SubscriberID string `pathParam:"style=simple,explode=false,name=subscriberId"`
	WorkflowID   string `pathParam:"style=simple,explode=false,name=parameter"`
	// A header for idempotency purposes
	IdempotencyKey                       *string                                         `header:"style=simple,explode=false,name=idempotency-key"`
	UpdateSubscriberPreferenceRequestDto components.UpdateSubscriberPreferenceRequestDto `request:"mediaType=application/json"`
}

func (o *SubscribersV1ControllerUpdateSubscriberPreferenceRequest) GetSubscriberID() string {
	if o == nil {
		return ""
	}
	return o.SubscriberID
}

func (o *SubscribersV1ControllerUpdateSubscriberPreferenceRequest) GetWorkflowID() string {
	if o == nil {
		return ""
	}
	return o.WorkflowID
}

func (o *SubscribersV1ControllerUpdateSubscriberPreferenceRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *SubscribersV1ControllerUpdateSubscriberPreferenceRequest) GetUpdateSubscriberPreferenceRequestDto() components.UpdateSubscriberPreferenceRequestDto {
	if o == nil {
		return components.UpdateSubscriberPreferenceRequestDto{}
	}
	return o.UpdateSubscriberPreferenceRequestDto
}

type SubscribersV1ControllerUpdateSubscriberPreferenceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	UpdateSubscriberPreferenceResponseDto *components.UpdateSubscriberPreferenceResponseDto
	Headers                               map[string][]string
}

func (o *SubscribersV1ControllerUpdateSubscriberPreferenceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SubscribersV1ControllerUpdateSubscriberPreferenceResponse) GetUpdateSubscriberPreferenceResponseDto() *components.UpdateSubscriberPreferenceResponseDto {
	if o == nil {
		return nil
	}
	return o.UpdateSubscriberPreferenceResponseDto
}

func (o *SubscribersV1ControllerUpdateSubscriberPreferenceResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
