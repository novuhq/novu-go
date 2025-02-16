// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// OrdinalEnum - Ordinal position for the digest
type OrdinalEnum string

const (
	OrdinalEnumOne   OrdinalEnum = "1"
	OrdinalEnumTwo   OrdinalEnum = "2"
	OrdinalEnumThree OrdinalEnum = "3"
	OrdinalEnumFour  OrdinalEnum = "4"
	OrdinalEnumFive  OrdinalEnum = "5"
	OrdinalEnumLast  OrdinalEnum = "last"
)

func (e OrdinalEnum) ToPointer() *OrdinalEnum {
	return &e
}
func (e *OrdinalEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "1":
		fallthrough
	case "2":
		fallthrough
	case "3":
		fallthrough
	case "4":
		fallthrough
	case "5":
		fallthrough
	case "last":
		*e = OrdinalEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OrdinalEnum: %v", v)
	}
}
