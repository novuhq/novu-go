// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// EmailBlockTypeEnum - Type of the email block
type EmailBlockTypeEnum string

const (
	EmailBlockTypeEnumButton EmailBlockTypeEnum = "button"
	EmailBlockTypeEnumText   EmailBlockTypeEnum = "text"
)

func (e EmailBlockTypeEnum) ToPointer() *EmailBlockTypeEnum {
	return &e
}
func (e *EmailBlockTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "button":
		fallthrough
	case "text":
		*e = EmailBlockTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EmailBlockTypeEnum: %v", v)
	}
}
