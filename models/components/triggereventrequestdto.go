// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/novuhq/novu-go/internal/utils"
)

type OneType string

const (
	OneTypeSubscriberPayloadDto OneType = "SubscriberPayloadDto"
	OneTypeTopicPayloadDto      OneType = "TopicPayloadDto"
	OneTypeStr                  OneType = "str"
)

type One struct {
	SubscriberPayloadDto *SubscriberPayloadDto `queryParam:"inline"`
	TopicPayloadDto      *TopicPayloadDto      `queryParam:"inline"`
	Str                  *string               `queryParam:"inline"`

	Type OneType
}

func CreateOneSubscriberPayloadDto(subscriberPayloadDto SubscriberPayloadDto) One {
	typ := OneTypeSubscriberPayloadDto

	return One{
		SubscriberPayloadDto: &subscriberPayloadDto,
		Type:                 typ,
	}
}

func CreateOneTopicPayloadDto(topicPayloadDto TopicPayloadDto) One {
	typ := OneTypeTopicPayloadDto

	return One{
		TopicPayloadDto: &topicPayloadDto,
		Type:            typ,
	}
}

func CreateOneStr(str string) One {
	typ := OneTypeStr

	return One{
		Str:  &str,
		Type: typ,
	}
}

func (u *One) UnmarshalJSON(data []byte) error {

	var topicPayloadDto TopicPayloadDto = TopicPayloadDto{}
	if err := utils.UnmarshalJSON(data, &topicPayloadDto, "", true, true); err == nil {
		u.TopicPayloadDto = &topicPayloadDto
		u.Type = OneTypeTopicPayloadDto
		return nil
	}

	var subscriberPayloadDto SubscriberPayloadDto = SubscriberPayloadDto{}
	if err := utils.UnmarshalJSON(data, &subscriberPayloadDto, "", true, true); err == nil {
		u.SubscriberPayloadDto = &subscriberPayloadDto
		u.Type = OneTypeSubscriberPayloadDto
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = OneTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for One", string(data))
}

func (u One) MarshalJSON() ([]byte, error) {
	if u.SubscriberPayloadDto != nil {
		return utils.MarshalJSON(u.SubscriberPayloadDto, "", true)
	}

	if u.TopicPayloadDto != nil {
		return utils.MarshalJSON(u.TopicPayloadDto, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type One: all fields are null")
}

type ToType string

const (
	ToTypeArrayOf1             ToType = "arrayOf1"
	ToTypeStr                  ToType = "str"
	ToTypeSubscriberPayloadDto ToType = "SubscriberPayloadDto"
	ToTypeTopicPayloadDto      ToType = "TopicPayloadDto"
)

// To - The recipients list of people who will receive the notification.
type To struct {
	ArrayOf1             []One                 `queryParam:"inline"`
	Str                  *string               `queryParam:"inline"`
	SubscriberPayloadDto *SubscriberPayloadDto `queryParam:"inline"`
	TopicPayloadDto      *TopicPayloadDto      `queryParam:"inline"`

	Type ToType
}

func CreateToArrayOf1(arrayOf1 []One) To {
	typ := ToTypeArrayOf1

	return To{
		ArrayOf1: arrayOf1,
		Type:     typ,
	}
}

func CreateToStr(str string) To {
	typ := ToTypeStr

	return To{
		Str:  &str,
		Type: typ,
	}
}

func CreateToSubscriberPayloadDto(subscriberPayloadDto SubscriberPayloadDto) To {
	typ := ToTypeSubscriberPayloadDto

	return To{
		SubscriberPayloadDto: &subscriberPayloadDto,
		Type:                 typ,
	}
}

func CreateToTopicPayloadDto(topicPayloadDto TopicPayloadDto) To {
	typ := ToTypeTopicPayloadDto

	return To{
		TopicPayloadDto: &topicPayloadDto,
		Type:            typ,
	}
}

func (u *To) UnmarshalJSON(data []byte) error {

	var topicPayloadDto TopicPayloadDto = TopicPayloadDto{}
	if err := utils.UnmarshalJSON(data, &topicPayloadDto, "", true, true); err == nil {
		u.TopicPayloadDto = &topicPayloadDto
		u.Type = ToTypeTopicPayloadDto
		return nil
	}

	var subscriberPayloadDto SubscriberPayloadDto = SubscriberPayloadDto{}
	if err := utils.UnmarshalJSON(data, &subscriberPayloadDto, "", true, true); err == nil {
		u.SubscriberPayloadDto = &subscriberPayloadDto
		u.Type = ToTypeSubscriberPayloadDto
		return nil
	}

	var arrayOf1 []One = []One{}
	if err := utils.UnmarshalJSON(data, &arrayOf1, "", true, true); err == nil {
		u.ArrayOf1 = arrayOf1
		u.Type = ToTypeArrayOf1
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ToTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for To", string(data))
}

func (u To) MarshalJSON() ([]byte, error) {
	if u.ArrayOf1 != nil {
		return utils.MarshalJSON(u.ArrayOf1, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.SubscriberPayloadDto != nil {
		return utils.MarshalJSON(u.SubscriberPayloadDto, "", true)
	}

	if u.TopicPayloadDto != nil {
		return utils.MarshalJSON(u.TopicPayloadDto, "", true)
	}

	return nil, errors.New("could not marshal union type To: all fields are null")
}

type ActorType string

const (
	ActorTypeStr                  ActorType = "str"
	ActorTypeSubscriberPayloadDto ActorType = "SubscriberPayloadDto"
)

// Actor - It is used to display the Avatar of the provided actor's subscriber id or actor object.
//
//	If a new actor object is provided, we will create a new subscriber in our system
type Actor struct {
	Str                  *string               `queryParam:"inline"`
	SubscriberPayloadDto *SubscriberPayloadDto `queryParam:"inline"`

	Type ActorType
}

func CreateActorStr(str string) Actor {
	typ := ActorTypeStr

	return Actor{
		Str:  &str,
		Type: typ,
	}
}

func CreateActorSubscriberPayloadDto(subscriberPayloadDto SubscriberPayloadDto) Actor {
	typ := ActorTypeSubscriberPayloadDto

	return Actor{
		SubscriberPayloadDto: &subscriberPayloadDto,
		Type:                 typ,
	}
}

func (u *Actor) UnmarshalJSON(data []byte) error {

	var subscriberPayloadDto SubscriberPayloadDto = SubscriberPayloadDto{}
	if err := utils.UnmarshalJSON(data, &subscriberPayloadDto, "", true, true); err == nil {
		u.SubscriberPayloadDto = &subscriberPayloadDto
		u.Type = ActorTypeSubscriberPayloadDto
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ActorTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Actor", string(data))
}

func (u Actor) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.SubscriberPayloadDto != nil {
		return utils.MarshalJSON(u.SubscriberPayloadDto, "", true)
	}

	return nil, errors.New("could not marshal union type Actor: all fields are null")
}

type TenantType string

const (
	TenantTypeStr              TenantType = "str"
	TenantTypeTenantPayloadDto TenantType = "TenantPayloadDto"
)

// Tenant - It is used to specify a tenant context during trigger event.
//
//	Existing tenants will be updated with the provided details.
type Tenant struct {
	Str              *string           `queryParam:"inline"`
	TenantPayloadDto *TenantPayloadDto `queryParam:"inline"`

	Type TenantType
}

func CreateTenantStr(str string) Tenant {
	typ := TenantTypeStr

	return Tenant{
		Str:  &str,
		Type: typ,
	}
}

func CreateTenantTenantPayloadDto(tenantPayloadDto TenantPayloadDto) Tenant {
	typ := TenantTypeTenantPayloadDto

	return Tenant{
		TenantPayloadDto: &tenantPayloadDto,
		Type:             typ,
	}
}

func (u *Tenant) UnmarshalJSON(data []byte) error {

	var tenantPayloadDto TenantPayloadDto = TenantPayloadDto{}
	if err := utils.UnmarshalJSON(data, &tenantPayloadDto, "", true, true); err == nil {
		u.TenantPayloadDto = &tenantPayloadDto
		u.Type = TenantTypeTenantPayloadDto
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = TenantTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Tenant", string(data))
}

func (u Tenant) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.TenantPayloadDto != nil {
		return utils.MarshalJSON(u.TenantPayloadDto, "", true)
	}

	return nil, errors.New("could not marshal union type Tenant: all fields are null")
}

type TriggerEventRequestDto struct {
	// The trigger identifier of the workflow you wish to send. This identifier can be found on the workflow page.
	Name string `json:"name"`
	// The payload object is used to pass additional custom information that could be
	//     used to render the workflow, or perform routing rules based on it.
	//       This data will also be available when fetching the notifications feed from the API to display certain parts of the UI.
	Payload map[string]any `json:"payload,omitempty"`
	// A URL to bridge for additional processing.
	BridgeURL *string `json:"bridgeUrl,omitempty"`
	// This could be used to override provider specific configurations
	Overrides map[string]map[string]any `json:"overrides,omitempty"`
	// The recipients list of people who will receive the notification.
	To To `json:"to"`
	// A unique identifier for this transaction, we will generate a UUID if not provided.
	TransactionID *string `json:"transactionId,omitempty"`
	// It is used to display the Avatar of the provided actor's subscriber id or actor object.
	//
	//
	//     If a new actor object is provided, we will create a new subscriber in our system
	Actor *Actor `json:"actor,omitempty"`
	// It is used to specify a tenant context during trigger event.
	//     Existing tenants will be updated with the provided details.
	Tenant *Tenant `json:"tenant,omitempty"`
	// Additional control configurations.
	Controls *WorkflowToStepControlValuesDto `json:"controls,omitempty"`
}

func (o *TriggerEventRequestDto) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *TriggerEventRequestDto) GetPayload() map[string]any {
	if o == nil {
		return nil
	}
	return o.Payload
}

func (o *TriggerEventRequestDto) GetBridgeURL() *string {
	if o == nil {
		return nil
	}
	return o.BridgeURL
}

func (o *TriggerEventRequestDto) GetOverrides() map[string]map[string]any {
	if o == nil {
		return nil
	}
	return o.Overrides
}

func (o *TriggerEventRequestDto) GetTo() To {
	if o == nil {
		return To{}
	}
	return o.To
}

func (o *TriggerEventRequestDto) GetTransactionID() *string {
	if o == nil {
		return nil
	}
	return o.TransactionID
}

func (o *TriggerEventRequestDto) GetActor() *Actor {
	if o == nil {
		return nil
	}
	return o.Actor
}

func (o *TriggerEventRequestDto) GetTenant() *Tenant {
	if o == nil {
		return nil
	}
	return o.Tenant
}

func (o *TriggerEventRequestDto) GetControls() *WorkflowToStepControlValuesDto {
	if o == nil {
		return nil
	}
	return o.Controls
}
