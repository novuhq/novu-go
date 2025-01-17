// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/novuhq/novu-go/models/components"
)

type TopicsControllerRenameTopicRequest struct {
	// The topic key
	TopicKey              string                           `pathParam:"style=simple,explode=false,name=topicKey"`
	RenameTopicRequestDto components.RenameTopicRequestDto `request:"mediaType=application/json"`
}

func (o *TopicsControllerRenameTopicRequest) GetTopicKey() string {
	if o == nil {
		return ""
	}
	return o.TopicKey
}

func (o *TopicsControllerRenameTopicRequest) GetRenameTopicRequestDto() components.RenameTopicRequestDto {
	if o == nil {
		return components.RenameTopicRequestDto{}
	}
	return o.RenameTopicRequestDto
}

type TopicsControllerRenameTopicResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	RenameTopicResponseDto *components.RenameTopicResponseDto
	Headers                map[string][]string
}

func (o *TopicsControllerRenameTopicResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *TopicsControllerRenameTopicResponse) GetRenameTopicResponseDto() *components.RenameTopicResponseDto {
	if o == nil {
		return nil
	}
	return o.RenameTopicResponseDto
}

func (o *TopicsControllerRenameTopicResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
