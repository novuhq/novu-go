// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/novuhq/novu-go/models/components"
)

type EventsControllerBroadcastEventToAllRequest struct {
	// A header for idempotency purposes
	IdempotencyKey              *string                                `header:"style=simple,explode=false,name=idempotency-key"`
	TriggerEventToAllRequestDto components.TriggerEventToAllRequestDto `request:"mediaType=application/json"`
}

func (o *EventsControllerBroadcastEventToAllRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *EventsControllerBroadcastEventToAllRequest) GetTriggerEventToAllRequestDto() components.TriggerEventToAllRequestDto {
	if o == nil {
		return components.TriggerEventToAllRequestDto{}
	}
	return o.TriggerEventToAllRequestDto
}

type EventsControllerBroadcastEventToAllResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	TriggerEventResponseDto *components.TriggerEventResponseDto
	Headers                 map[string][]string
}

func (o *EventsControllerBroadcastEventToAllResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *EventsControllerBroadcastEventToAllResponse) GetTriggerEventResponseDto() *components.TriggerEventResponseDto {
	if o == nil {
		return nil
	}
	return o.TriggerEventResponseDto
}

func (o *EventsControllerBroadcastEventToAllResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
