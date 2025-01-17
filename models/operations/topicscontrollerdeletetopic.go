// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/novuhq/novu-go/models/components"
)

type TopicsControllerDeleteTopicRequest struct {
	// The topic key
	TopicKey string `pathParam:"style=simple,explode=false,name=topicKey"`
}

func (o *TopicsControllerDeleteTopicRequest) GetTopicKey() string {
	if o == nil {
		return ""
	}
	return o.TopicKey
}

type TopicsControllerDeleteTopicResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Headers  map[string][]string
}

func (o *TopicsControllerDeleteTopicResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *TopicsControllerDeleteTopicResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
