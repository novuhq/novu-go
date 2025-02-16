// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/novuhq/novu-go/models/components"
)

type NotificationsControllerGetActivityStatsRequest struct {
	// A header for idempotency purposes
	IdempotencyKey *string `header:"style=simple,explode=false,name=idempotency-key"`
}

func (o *NotificationsControllerGetActivityStatsRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

type NotificationsControllerGetActivityStatsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ActivityStatsResponseDto *components.ActivityStatsResponseDto
	Headers                  map[string][]string
}

func (o *NotificationsControllerGetActivityStatsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *NotificationsControllerGetActivityStatsResponse) GetActivityStatsResponseDto() *components.ActivityStatsResponseDto {
	if o == nil {
		return nil
	}
	return o.ActivityStatsResponseDto
}

func (o *NotificationsControllerGetActivityStatsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
