// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UpdatedSubscriberDto struct {
	// The ID of the subscriber that was updated.
	SubscriberID string `json:"subscriberId"`
}

func (o *UpdatedSubscriberDto) GetSubscriberID() string {
	if o == nil {
		return ""
	}
	return o.SubscriberID
}
