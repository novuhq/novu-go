// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type BulkSubscriberCreateDto struct {
	// An array of subscribers to be created in bulk.
	Subscribers []CreateSubscriberRequestDto `json:"subscribers"`
}

func (o *BulkSubscriberCreateDto) GetSubscribers() []CreateSubscriberRequestDto {
	if o == nil {
		return []CreateSubscriberRequestDto{}
	}
	return o.Subscribers
}
