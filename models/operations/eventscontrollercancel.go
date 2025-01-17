// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/novuhq/novu-go/models/components"
)

type EventsControllerCancelRequest struct {
	TransactionID string `pathParam:"style=simple,explode=false,name=transactionId"`
}

func (o *EventsControllerCancelRequest) GetTransactionID() string {
	if o == nil {
		return ""
	}
	return o.TransactionID
}

type EventsControllerCancelResponse struct {
	HTTPMeta       components.HTTPMetadata `json:"-"`
	DataBooleanDto *components.DataBooleanDto
	Headers        map[string][]string
}

func (o *EventsControllerCancelResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *EventsControllerCancelResponse) GetDataBooleanDto() *components.DataBooleanDto {
	if o == nil {
		return nil
	}
	return o.DataBooleanDto
}

func (o *EventsControllerCancelResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
