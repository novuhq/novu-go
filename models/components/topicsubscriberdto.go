// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type TopicSubscriberDto struct {
	// Unique identifier for the organization
	OrganizationID string `json:"_organizationId"`
	// Unique identifier for the environment
	EnvironmentID string `json:"_environmentId"`
	// Unique identifier for the subscriber
	SubscriberID string `json:"_subscriberId"`
	// Unique identifier for the topic
	TopicID string `json:"_topicId"`
	// Key associated with the topic
	TopicKey string `json:"topicKey"`
	// External identifier for the subscriber
	ExternalSubscriberID string `json:"externalSubscriberId"`
}

func (o *TopicSubscriberDto) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *TopicSubscriberDto) GetEnvironmentID() string {
	if o == nil {
		return ""
	}
	return o.EnvironmentID
}

func (o *TopicSubscriberDto) GetSubscriberID() string {
	if o == nil {
		return ""
	}
	return o.SubscriberID
}

func (o *TopicSubscriberDto) GetTopicID() string {
	if o == nil {
		return ""
	}
	return o.TopicID
}

func (o *TopicSubscriberDto) GetTopicKey() string {
	if o == nil {
		return ""
	}
	return o.TopicKey
}

func (o *TopicSubscriberDto) GetExternalSubscriberID() string {
	if o == nil {
		return ""
	}
	return o.ExternalSubscriberID
}
