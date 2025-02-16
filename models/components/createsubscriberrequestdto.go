// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/novuhq/novu-go/internal/utils"
)

type CreateSubscriberRequestDtoDataType string

const (
	CreateSubscriberRequestDtoDataTypeStr        CreateSubscriberRequestDtoDataType = "str"
	CreateSubscriberRequestDtoDataTypeArrayOfStr CreateSubscriberRequestDtoDataType = "arrayOfStr"
	CreateSubscriberRequestDtoDataTypeBoolean    CreateSubscriberRequestDtoDataType = "boolean"
	CreateSubscriberRequestDtoDataTypeNumber     CreateSubscriberRequestDtoDataType = "number"
)

type CreateSubscriberRequestDtoData struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`
	Boolean    *bool    `queryParam:"inline"`
	Number     *float64 `queryParam:"inline"`

	Type CreateSubscriberRequestDtoDataType
}

func CreateCreateSubscriberRequestDtoDataStr(str string) CreateSubscriberRequestDtoData {
	typ := CreateSubscriberRequestDtoDataTypeStr

	return CreateSubscriberRequestDtoData{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateSubscriberRequestDtoDataArrayOfStr(arrayOfStr []string) CreateSubscriberRequestDtoData {
	typ := CreateSubscriberRequestDtoDataTypeArrayOfStr

	return CreateSubscriberRequestDtoData{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func CreateCreateSubscriberRequestDtoDataBoolean(boolean bool) CreateSubscriberRequestDtoData {
	typ := CreateSubscriberRequestDtoDataTypeBoolean

	return CreateSubscriberRequestDtoData{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateCreateSubscriberRequestDtoDataNumber(number float64) CreateSubscriberRequestDtoData {
	typ := CreateSubscriberRequestDtoDataTypeNumber

	return CreateSubscriberRequestDtoData{
		Number: &number,
		Type:   typ,
	}
}

func (u *CreateSubscriberRequestDtoData) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateSubscriberRequestDtoDataTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = CreateSubscriberRequestDtoDataTypeArrayOfStr
		return nil
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = CreateSubscriberRequestDtoDataTypeBoolean
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = CreateSubscriberRequestDtoDataTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateSubscriberRequestDtoData", string(data))
}

func (u CreateSubscriberRequestDtoData) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type CreateSubscriberRequestDtoData: all fields are null")
}

type CreateSubscriberRequestDto struct {
	// The internal identifier you used to create this subscriber, usually correlates to the id the user in your systems
	SubscriberID string `json:"subscriberId"`
	// The email address of the subscriber.
	Email *string `json:"email,omitempty"`
	// The first name of the subscriber.
	FirstName *string `json:"firstName,omitempty"`
	// The last name of the subscriber.
	LastName *string `json:"lastName,omitempty"`
	// The phone number of the subscriber.
	Phone *string `json:"phone,omitempty"`
	// An HTTP URL to the profile image of your subscriber.
	Avatar *string `json:"avatar,omitempty"`
	// The locale of the subscriber.
	Locale *string `json:"locale,omitempty"`
	// An optional payload object that can contain any properties.
	Data map[string]CreateSubscriberRequestDtoData `json:"data,omitempty"`
	// An optional array of subscriber channels.
	Channels []SubscriberChannelDto `json:"channels,omitempty"`
}

func (o *CreateSubscriberRequestDto) GetSubscriberID() string {
	if o == nil {
		return ""
	}
	return o.SubscriberID
}

func (o *CreateSubscriberRequestDto) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *CreateSubscriberRequestDto) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *CreateSubscriberRequestDto) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *CreateSubscriberRequestDto) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *CreateSubscriberRequestDto) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *CreateSubscriberRequestDto) GetLocale() *string {
	if o == nil {
		return nil
	}
	return o.Locale
}

func (o *CreateSubscriberRequestDto) GetData() map[string]CreateSubscriberRequestDtoData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *CreateSubscriberRequestDto) GetChannels() []SubscriberChannelDto {
	if o == nil {
		return nil
	}
	return o.Channels
}
