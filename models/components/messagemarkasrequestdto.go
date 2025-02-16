// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/novuhq/novu-go/internal/utils"
)

type MessageIDType string

const (
	MessageIDTypeStr        MessageIDType = "str"
	MessageIDTypeArrayOfStr MessageIDType = "arrayOfStr"
)

type MessageID struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type MessageIDType
}

func CreateMessageIDStr(str string) MessageID {
	typ := MessageIDTypeStr

	return MessageID{
		Str:  &str,
		Type: typ,
	}
}

func CreateMessageIDArrayOfStr(arrayOfStr []string) MessageID {
	typ := MessageIDTypeArrayOfStr

	return MessageID{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *MessageID) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = MessageIDTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = MessageIDTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for MessageID", string(data))
}

func (u MessageID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type MessageID: all fields are null")
}

type MarkAs string

const (
	MarkAsRead   MarkAs = "read"
	MarkAsSeen   MarkAs = "seen"
	MarkAsUnread MarkAs = "unread"
	MarkAsUnseen MarkAs = "unseen"
)

func (e MarkAs) ToPointer() *MarkAs {
	return &e
}
func (e *MarkAs) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "read":
		fallthrough
	case "seen":
		fallthrough
	case "unread":
		fallthrough
	case "unseen":
		*e = MarkAs(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MarkAs: %v", v)
	}
}

type MessageMarkAsRequestDto struct {
	MessageID MessageID `json:"messageId"`
	MarkAs    MarkAs    `json:"markAs"`
}

func (o *MessageMarkAsRequestDto) GetMessageID() MessageID {
	if o == nil {
		return MessageID{}
	}
	return o.MessageID
}

func (o *MessageMarkAsRequestDto) GetMarkAs() MarkAs {
	if o == nil {
		return MarkAs("")
	}
	return o.MarkAs
}
