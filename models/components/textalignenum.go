// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// TextAlignEnum - Text alignment for the email block
type TextAlignEnum string

const (
	TextAlignEnumCenter TextAlignEnum = "center"
	TextAlignEnumLeft   TextAlignEnum = "left"
	TextAlignEnumRight  TextAlignEnum = "right"
)

func (e TextAlignEnum) ToPointer() *TextAlignEnum {
	return &e
}
func (e *TextAlignEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "center":
		fallthrough
	case "left":
		fallthrough
	case "right":
		*e = TextAlignEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TextAlignEnum: %v", v)
	}
}
