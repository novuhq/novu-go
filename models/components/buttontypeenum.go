// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ButtonTypeEnum - Type of button for the action result
type ButtonTypeEnum string

const (
	ButtonTypeEnumPrimary   ButtonTypeEnum = "primary"
	ButtonTypeEnumSecondary ButtonTypeEnum = "secondary"
)

func (e ButtonTypeEnum) ToPointer() *ButtonTypeEnum {
	return &e
}
func (e *ButtonTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "primary":
		fallthrough
	case "secondary":
		*e = ButtonTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ButtonTypeEnum: %v", v)
	}
}
