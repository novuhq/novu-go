// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ActivityNotificationResponseDtoPayload - Payload of the notification
type ActivityNotificationResponseDtoPayload struct {
}

// Controls associated with the notification
type Controls struct {
}

// ActivityNotificationResponseDtoTo - To field for subscriber definition
type ActivityNotificationResponseDtoTo struct {
}

type ActivityNotificationResponseDto struct {
	// Unique identifier of the notification
	ID *string `json:"_id,omitempty"`
	// Environment ID of the notification
	EnvironmentID string `json:"_environmentId"`
	// Organization ID of the notification
	OrganizationID string `json:"_organizationId"`
	// Subscriber ID of the notification
	SubscriberID string `json:"_subscriberId"`
	// Transaction ID of the notification
	TransactionID string `json:"transactionId"`
	// Template ID of the notification
	TemplateID *string `json:"_templateId,omitempty"`
	// Digested Notification ID
	DigestedNotificationID *string `json:"_digestedNotificationId,omitempty"`
	// Creation time of the notification
	CreatedAt *string `json:"createdAt,omitempty"`
	// Last updated time of the notification
	UpdatedAt *string        `json:"updatedAt,omitempty"`
	Channels  []StepTypeEnum `json:"channels,omitempty"`
	// Subscriber of the notification
	Subscriber *ActivityNotificationSubscriberResponseDto `json:"subscriber,omitempty"`
	// Template of the notification
	Template *ActivityNotificationTemplateResponseDto `json:"template,omitempty"`
	// Jobs of the notification
	Jobs []ActivityNotificationJobResponseDto `json:"jobs,omitempty"`
	// Payload of the notification
	Payload *ActivityNotificationResponseDtoPayload `json:"payload,omitempty"`
	// Tags associated with the notification
	Tags []string `json:"tags,omitempty"`
	// Controls associated with the notification
	Controls *Controls `json:"controls,omitempty"`
	// To field for subscriber definition
	To *ActivityNotificationResponseDtoTo `json:"to,omitempty"`
	// Topics of the notification
	Topics []ActivityTopicDto `json:"topics,omitempty"`
}

func (o *ActivityNotificationResponseDto) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ActivityNotificationResponseDto) GetEnvironmentID() string {
	if o == nil {
		return ""
	}
	return o.EnvironmentID
}

func (o *ActivityNotificationResponseDto) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *ActivityNotificationResponseDto) GetSubscriberID() string {
	if o == nil {
		return ""
	}
	return o.SubscriberID
}

func (o *ActivityNotificationResponseDto) GetTransactionID() string {
	if o == nil {
		return ""
	}
	return o.TransactionID
}

func (o *ActivityNotificationResponseDto) GetTemplateID() *string {
	if o == nil {
		return nil
	}
	return o.TemplateID
}

func (o *ActivityNotificationResponseDto) GetDigestedNotificationID() *string {
	if o == nil {
		return nil
	}
	return o.DigestedNotificationID
}

func (o *ActivityNotificationResponseDto) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ActivityNotificationResponseDto) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ActivityNotificationResponseDto) GetChannels() []StepTypeEnum {
	if o == nil {
		return nil
	}
	return o.Channels
}

func (o *ActivityNotificationResponseDto) GetSubscriber() *ActivityNotificationSubscriberResponseDto {
	if o == nil {
		return nil
	}
	return o.Subscriber
}

func (o *ActivityNotificationResponseDto) GetTemplate() *ActivityNotificationTemplateResponseDto {
	if o == nil {
		return nil
	}
	return o.Template
}

func (o *ActivityNotificationResponseDto) GetJobs() []ActivityNotificationJobResponseDto {
	if o == nil {
		return nil
	}
	return o.Jobs
}

func (o *ActivityNotificationResponseDto) GetPayload() *ActivityNotificationResponseDtoPayload {
	if o == nil {
		return nil
	}
	return o.Payload
}

func (o *ActivityNotificationResponseDto) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ActivityNotificationResponseDto) GetControls() *Controls {
	if o == nil {
		return nil
	}
	return o.Controls
}

func (o *ActivityNotificationResponseDto) GetTo() *ActivityNotificationResponseDtoTo {
	if o == nil {
		return nil
	}
	return o.To
}

func (o *ActivityNotificationResponseDto) GetTopics() []ActivityTopicDto {
	if o == nil {
		return nil
	}
	return o.Topics
}
