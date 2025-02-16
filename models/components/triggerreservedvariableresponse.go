// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// TriggerReservedVariableResponseType - The type of the reserved variable
type TriggerReservedVariableResponseType string

const (
	TriggerReservedVariableResponseTypeTenant TriggerReservedVariableResponseType = "tenant"
	TriggerReservedVariableResponseTypeActor  TriggerReservedVariableResponseType = "actor"
)

func (e TriggerReservedVariableResponseType) ToPointer() *TriggerReservedVariableResponseType {
	return &e
}
func (e *TriggerReservedVariableResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tenant":
		fallthrough
	case "actor":
		*e = TriggerReservedVariableResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TriggerReservedVariableResponseType: %v", v)
	}
}

type TriggerReservedVariableResponse struct {
	// The type of the reserved variable
	Type TriggerReservedVariableResponseType `json:"type"`
	// The reserved variables of the trigger
	Variables []string `json:"variables"`
}

func (o *TriggerReservedVariableResponse) GetType() TriggerReservedVariableResponseType {
	if o == nil {
		return TriggerReservedVariableResponseType("")
	}
	return o.Type
}

func (o *TriggerReservedVariableResponse) GetVariables() []string {
	if o == nil {
		return []string{}
	}
	return o.Variables
}
