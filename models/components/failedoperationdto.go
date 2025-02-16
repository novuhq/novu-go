// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type FailedOperationDto struct {
	// The error message associated with the failed operation.
	Message *string `json:"message,omitempty"`
	// The subscriber ID associated with the failed operation. This field is optional.
	SubscriberID *string `json:"subscriberId,omitempty"`
}

func (o *FailedOperationDto) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *FailedOperationDto) GetSubscriberID() *string {
	if o == nil {
		return nil
	}
	return o.SubscriberID
}
