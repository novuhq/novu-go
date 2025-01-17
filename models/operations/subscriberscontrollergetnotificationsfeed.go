// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/novuhq/novu-go/internal/utils"
	"github.com/novuhq/novu-go/models/components"
)

type SubscribersControllerGetNotificationsFeedRequest struct {
	SubscriberID string   `pathParam:"style=simple,explode=false,name=subscriberId"`
	Page         *float64 `queryParam:"style=form,explode=true,name=page"`
	Limit        *float64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	Read         *bool    `queryParam:"style=form,explode=true,name=read"`
	Seen         *bool    `queryParam:"style=form,explode=true,name=seen"`
	// Base64 encoded string of the partial payload JSON object
	Payload *string `queryParam:"style=form,explode=true,name=payload"`
}

func (s SubscribersControllerGetNotificationsFeedRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SubscribersControllerGetNotificationsFeedRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SubscribersControllerGetNotificationsFeedRequest) GetSubscriberID() string {
	if o == nil {
		return ""
	}
	return o.SubscriberID
}

func (o *SubscribersControllerGetNotificationsFeedRequest) GetPage() *float64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *SubscribersControllerGetNotificationsFeedRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *SubscribersControllerGetNotificationsFeedRequest) GetRead() *bool {
	if o == nil {
		return nil
	}
	return o.Read
}

func (o *SubscribersControllerGetNotificationsFeedRequest) GetSeen() *bool {
	if o == nil {
		return nil
	}
	return o.Seen
}

func (o *SubscribersControllerGetNotificationsFeedRequest) GetPayload() *string {
	if o == nil {
		return nil
	}
	return o.Payload
}

type SubscribersControllerGetNotificationsFeedResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	FeedResponseDto *components.FeedResponseDto
	Headers         map[string][]string
}

func (o *SubscribersControllerGetNotificationsFeedResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SubscribersControllerGetNotificationsFeedResponse) GetFeedResponseDto() *components.FeedResponseDto {
	if o == nil {
		return nil
	}
	return o.FeedResponseDto
}

func (o *SubscribersControllerGetNotificationsFeedResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
