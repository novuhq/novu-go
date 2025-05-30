// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type TopicSubscriptionResponseDto struct {
	// The identifier of the subscription
	ID string `json:"_id"`
	// The date and time the subscription was created
	CreatedAt string `json:"createdAt"`
	// Topic information
	Topic TopicResponseDto `json:"topic"`
	// Subscriber information
	Subscriber SubscriberDto `json:"subscriber"`
}

func (o *TopicSubscriptionResponseDto) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TopicSubscriptionResponseDto) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *TopicSubscriptionResponseDto) GetTopic() TopicResponseDto {
	if o == nil {
		return TopicResponseDto{}
	}
	return o.Topic
}

func (o *TopicSubscriptionResponseDto) GetSubscriber() SubscriberDto {
	if o == nil {
		return SubscriberDto{}
	}
	return o.Subscriber
}
