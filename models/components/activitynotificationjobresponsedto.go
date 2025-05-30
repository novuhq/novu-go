// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ActivityNotificationJobResponseDtoType - Type of the job
type ActivityNotificationJobResponseDtoType string

const (
	ActivityNotificationJobResponseDtoTypeInApp   ActivityNotificationJobResponseDtoType = "in_app"
	ActivityNotificationJobResponseDtoTypeEmail   ActivityNotificationJobResponseDtoType = "email"
	ActivityNotificationJobResponseDtoTypeSms     ActivityNotificationJobResponseDtoType = "sms"
	ActivityNotificationJobResponseDtoTypeChat    ActivityNotificationJobResponseDtoType = "chat"
	ActivityNotificationJobResponseDtoTypePush    ActivityNotificationJobResponseDtoType = "push"
	ActivityNotificationJobResponseDtoTypeDigest  ActivityNotificationJobResponseDtoType = "digest"
	ActivityNotificationJobResponseDtoTypeTrigger ActivityNotificationJobResponseDtoType = "trigger"
	ActivityNotificationJobResponseDtoTypeDelay   ActivityNotificationJobResponseDtoType = "delay"
	ActivityNotificationJobResponseDtoTypeCustom  ActivityNotificationJobResponseDtoType = "custom"
)

func (e ActivityNotificationJobResponseDtoType) ToPointer() *ActivityNotificationJobResponseDtoType {
	return &e
}
func (e *ActivityNotificationJobResponseDtoType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "in_app":
		fallthrough
	case "email":
		fallthrough
	case "sms":
		fallthrough
	case "chat":
		fallthrough
	case "push":
		fallthrough
	case "digest":
		fallthrough
	case "trigger":
		fallthrough
	case "delay":
		fallthrough
	case "custom":
		*e = ActivityNotificationJobResponseDtoType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ActivityNotificationJobResponseDtoType: %v", v)
	}
}

// ActivityNotificationJobResponseDtoPayload - Optional payload for the job
type ActivityNotificationJobResponseDtoPayload struct {
}

type ActivityNotificationJobResponseDto struct {
	// Unique identifier of the job
	ID string `json:"_id"`
	// Type of the job
	Type ActivityNotificationJobResponseDtoType `json:"type"`
	// Optional digest for the job, including metadata and events
	Digest *DigestMetadataDto `json:"digest,omitempty"`
	// Execution details of the job
	ExecutionDetails []ActivityNotificationExecutionDetailResponseDto `json:"executionDetails"`
	// Step details of the job
	Step ActivityNotificationStepResponseDto `json:"step"`
	// Optional context object for additional error details.
	Overrides map[string]any `json:"overrides,omitempty"`
	// Optional payload for the job
	Payload *ActivityNotificationJobResponseDtoPayload `json:"payload,omitempty"`
	// Provider ID of the job
	ProviderID ProvidersIDEnum `json:"providerId"`
	// Status of the job
	Status string `json:"status"`
	// Updated time of the notification
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

func (o *ActivityNotificationJobResponseDto) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ActivityNotificationJobResponseDto) GetType() ActivityNotificationJobResponseDtoType {
	if o == nil {
		return ActivityNotificationJobResponseDtoType("")
	}
	return o.Type
}

func (o *ActivityNotificationJobResponseDto) GetDigest() *DigestMetadataDto {
	if o == nil {
		return nil
	}
	return o.Digest
}

func (o *ActivityNotificationJobResponseDto) GetExecutionDetails() []ActivityNotificationExecutionDetailResponseDto {
	if o == nil {
		return []ActivityNotificationExecutionDetailResponseDto{}
	}
	return o.ExecutionDetails
}

func (o *ActivityNotificationJobResponseDto) GetStep() ActivityNotificationStepResponseDto {
	if o == nil {
		return ActivityNotificationStepResponseDto{}
	}
	return o.Step
}

func (o *ActivityNotificationJobResponseDto) GetOverrides() map[string]any {
	if o == nil {
		return nil
	}
	return o.Overrides
}

func (o *ActivityNotificationJobResponseDto) GetPayload() *ActivityNotificationJobResponseDtoPayload {
	if o == nil {
		return nil
	}
	return o.Payload
}

func (o *ActivityNotificationJobResponseDto) GetProviderID() ProvidersIDEnum {
	if o == nil {
		return ProvidersIDEnum("")
	}
	return o.ProviderID
}

func (o *ActivityNotificationJobResponseDto) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *ActivityNotificationJobResponseDto) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
