// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ActorTypeEnum - The type of the actor, indicating the role in the notification process.
type ActorTypeEnum string

const (
	ActorTypeEnumNone         ActorTypeEnum = "none"
	ActorTypeEnumUser         ActorTypeEnum = "user"
	ActorTypeEnumSystemIcon   ActorTypeEnum = "system_icon"
	ActorTypeEnumSystemCustom ActorTypeEnum = "system_custom"
)

func (e ActorTypeEnum) ToPointer() *ActorTypeEnum {
	return &e
}
func (e *ActorTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "user":
		fallthrough
	case "system_icon":
		fallthrough
	case "system_custom":
		*e = ActorTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ActorTypeEnum: %v", v)
	}
}
