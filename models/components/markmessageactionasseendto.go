// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// MarkMessageActionAsSeenDtoStatus - Message action status
type MarkMessageActionAsSeenDtoStatus string

const (
	MarkMessageActionAsSeenDtoStatusPending MarkMessageActionAsSeenDtoStatus = "pending"
	MarkMessageActionAsSeenDtoStatusDone    MarkMessageActionAsSeenDtoStatus = "done"
)

func (e MarkMessageActionAsSeenDtoStatus) ToPointer() *MarkMessageActionAsSeenDtoStatus {
	return &e
}
func (e *MarkMessageActionAsSeenDtoStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "done":
		*e = MarkMessageActionAsSeenDtoStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MarkMessageActionAsSeenDtoStatus: %v", v)
	}
}

// MarkMessageActionAsSeenDtoPayload - Message action payload
type MarkMessageActionAsSeenDtoPayload struct {
}

type MarkMessageActionAsSeenDto struct {
	// Message action status
	Status MarkMessageActionAsSeenDtoStatus `json:"status"`
	// Message action payload
	Payload *MarkMessageActionAsSeenDtoPayload `json:"payload,omitempty"`
}

func (o *MarkMessageActionAsSeenDto) GetStatus() MarkMessageActionAsSeenDtoStatus {
	if o == nil {
		return MarkMessageActionAsSeenDtoStatus("")
	}
	return o.Status
}

func (o *MarkMessageActionAsSeenDto) GetPayload() *MarkMessageActionAsSeenDtoPayload {
	if o == nil {
		return nil
	}
	return o.Payload
}
