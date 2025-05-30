// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// NotificationTriggerDtoType - Type of the trigger
type NotificationTriggerDtoType string

const (
	NotificationTriggerDtoTypeEvent NotificationTriggerDtoType = "event"
)

func (e NotificationTriggerDtoType) ToPointer() *NotificationTriggerDtoType {
	return &e
}
func (e *NotificationTriggerDtoType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "event":
		*e = NotificationTriggerDtoType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NotificationTriggerDtoType: %v", v)
	}
}

type NotificationTriggerDto struct {
	// Type of the trigger
	Type NotificationTriggerDtoType `json:"type"`
	// Identifier of the trigger
	Identifier string `json:"identifier"`
	// Variables of the trigger
	Variables []NotificationTriggerVariable `json:"variables"`
	// Subscriber variables of the trigger
	SubscriberVariables []NotificationTriggerVariable `json:"subscriberVariables,omitempty"`
}

func (o *NotificationTriggerDto) GetType() NotificationTriggerDtoType {
	if o == nil {
		return NotificationTriggerDtoType("")
	}
	return o.Type
}

func (o *NotificationTriggerDto) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *NotificationTriggerDto) GetVariables() []NotificationTriggerVariable {
	if o == nil {
		return []NotificationTriggerVariable{}
	}
	return o.Variables
}

func (o *NotificationTriggerDto) GetSubscriberVariables() []NotificationTriggerVariable {
	if o == nil {
		return nil
	}
	return o.SubscriberVariables
}
