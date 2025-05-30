// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateTopicSubscriptionsRequestDto struct {
	// List of subscriber identifiers to subscribe to the topic (max: 100)
	SubscriberIds []string `json:"subscriberIds"`
}

func (o *CreateTopicSubscriptionsRequestDto) GetSubscriberIds() []string {
	if o == nil {
		return []string{}
	}
	return o.SubscriberIds
}
