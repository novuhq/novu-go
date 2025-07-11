// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/novuhq/novu-go/internal/utils"
)

type StepsType string

const (
	StepsTypeInApp  StepsType = "in_app"
	StepsTypeEmail  StepsType = "email"
	StepsTypeSms    StepsType = "sms"
	StepsTypePush   StepsType = "push"
	StepsTypeChat   StepsType = "chat"
	StepsTypeDelay  StepsType = "delay"
	StepsTypeDigest StepsType = "digest"
	StepsTypeCustom StepsType = "custom"
)

type Steps struct {
	InAppStepUpsertDto  *InAppStepUpsertDto  `queryParam:"inline"`
	EmailStepUpsertDto  *EmailStepUpsertDto  `queryParam:"inline"`
	SmsStepUpsertDto    *SmsStepUpsertDto    `queryParam:"inline"`
	PushStepUpsertDto   *PushStepUpsertDto   `queryParam:"inline"`
	ChatStepUpsertDto   *ChatStepUpsertDto   `queryParam:"inline"`
	DelayStepUpsertDto  *DelayStepUpsertDto  `queryParam:"inline"`
	DigestStepUpsertDto *DigestStepUpsertDto `queryParam:"inline"`
	CustomStepUpsertDto *CustomStepUpsertDto `queryParam:"inline"`

	Type StepsType
}

func CreateStepsInApp(inApp InAppStepUpsertDto) Steps {
	typ := StepsTypeInApp

	typStr := StepTypeEnum(typ)
	inApp.Type = typStr

	return Steps{
		InAppStepUpsertDto: &inApp,
		Type:               typ,
	}
}

func CreateStepsEmail(email EmailStepUpsertDto) Steps {
	typ := StepsTypeEmail

	typStr := StepTypeEnum(typ)
	email.Type = typStr

	return Steps{
		EmailStepUpsertDto: &email,
		Type:               typ,
	}
}

func CreateStepsSms(sms SmsStepUpsertDto) Steps {
	typ := StepsTypeSms

	typStr := StepTypeEnum(typ)
	sms.Type = typStr

	return Steps{
		SmsStepUpsertDto: &sms,
		Type:             typ,
	}
}

func CreateStepsPush(push PushStepUpsertDto) Steps {
	typ := StepsTypePush

	typStr := StepTypeEnum(typ)
	push.Type = typStr

	return Steps{
		PushStepUpsertDto: &push,
		Type:              typ,
	}
}

func CreateStepsChat(chat ChatStepUpsertDto) Steps {
	typ := StepsTypeChat

	typStr := StepTypeEnum(typ)
	chat.Type = typStr

	return Steps{
		ChatStepUpsertDto: &chat,
		Type:              typ,
	}
}

func CreateStepsDelay(delay DelayStepUpsertDto) Steps {
	typ := StepsTypeDelay

	typStr := StepTypeEnum(typ)
	delay.Type = typStr

	return Steps{
		DelayStepUpsertDto: &delay,
		Type:               typ,
	}
}

func CreateStepsDigest(digest DigestStepUpsertDto) Steps {
	typ := StepsTypeDigest

	typStr := StepTypeEnum(typ)
	digest.Type = typStr

	return Steps{
		DigestStepUpsertDto: &digest,
		Type:                typ,
	}
}

func CreateStepsCustom(custom CustomStepUpsertDto) Steps {
	typ := StepsTypeCustom

	typStr := StepTypeEnum(typ)
	custom.Type = typStr

	return Steps{
		CustomStepUpsertDto: &custom,
		Type:                typ,
	}
}

func (u *Steps) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		Type string `json:"type"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Type {
	case "in_app":
		inAppStepUpsertDto := new(InAppStepUpsertDto)
		if err := utils.UnmarshalJSON(data, &inAppStepUpsertDto, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == in_app) type InAppStepUpsertDto within Steps: %w", string(data), err)
		}

		u.InAppStepUpsertDto = inAppStepUpsertDto
		u.Type = StepsTypeInApp
		return nil
	case "email":
		emailStepUpsertDto := new(EmailStepUpsertDto)
		if err := utils.UnmarshalJSON(data, &emailStepUpsertDto, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == email) type EmailStepUpsertDto within Steps: %w", string(data), err)
		}

		u.EmailStepUpsertDto = emailStepUpsertDto
		u.Type = StepsTypeEmail
		return nil
	case "sms":
		smsStepUpsertDto := new(SmsStepUpsertDto)
		if err := utils.UnmarshalJSON(data, &smsStepUpsertDto, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == sms) type SmsStepUpsertDto within Steps: %w", string(data), err)
		}

		u.SmsStepUpsertDto = smsStepUpsertDto
		u.Type = StepsTypeSms
		return nil
	case "push":
		pushStepUpsertDto := new(PushStepUpsertDto)
		if err := utils.UnmarshalJSON(data, &pushStepUpsertDto, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == push) type PushStepUpsertDto within Steps: %w", string(data), err)
		}

		u.PushStepUpsertDto = pushStepUpsertDto
		u.Type = StepsTypePush
		return nil
	case "chat":
		chatStepUpsertDto := new(ChatStepUpsertDto)
		if err := utils.UnmarshalJSON(data, &chatStepUpsertDto, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == chat) type ChatStepUpsertDto within Steps: %w", string(data), err)
		}

		u.ChatStepUpsertDto = chatStepUpsertDto
		u.Type = StepsTypeChat
		return nil
	case "delay":
		delayStepUpsertDto := new(DelayStepUpsertDto)
		if err := utils.UnmarshalJSON(data, &delayStepUpsertDto, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == delay) type DelayStepUpsertDto within Steps: %w", string(data), err)
		}

		u.DelayStepUpsertDto = delayStepUpsertDto
		u.Type = StepsTypeDelay
		return nil
	case "digest":
		digestStepUpsertDto := new(DigestStepUpsertDto)
		if err := utils.UnmarshalJSON(data, &digestStepUpsertDto, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == digest) type DigestStepUpsertDto within Steps: %w", string(data), err)
		}

		u.DigestStepUpsertDto = digestStepUpsertDto
		u.Type = StepsTypeDigest
		return nil
	case "custom":
		customStepUpsertDto := new(CustomStepUpsertDto)
		if err := utils.UnmarshalJSON(data, &customStepUpsertDto, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == custom) type CustomStepUpsertDto within Steps: %w", string(data), err)
		}

		u.CustomStepUpsertDto = customStepUpsertDto
		u.Type = StepsTypeCustom
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Steps", string(data))
}

func (u Steps) MarshalJSON() ([]byte, error) {
	if u.InAppStepUpsertDto != nil {
		return utils.MarshalJSON(u.InAppStepUpsertDto, "", true)
	}

	if u.EmailStepUpsertDto != nil {
		return utils.MarshalJSON(u.EmailStepUpsertDto, "", true)
	}

	if u.SmsStepUpsertDto != nil {
		return utils.MarshalJSON(u.SmsStepUpsertDto, "", true)
	}

	if u.PushStepUpsertDto != nil {
		return utils.MarshalJSON(u.PushStepUpsertDto, "", true)
	}

	if u.ChatStepUpsertDto != nil {
		return utils.MarshalJSON(u.ChatStepUpsertDto, "", true)
	}

	if u.DelayStepUpsertDto != nil {
		return utils.MarshalJSON(u.DelayStepUpsertDto, "", true)
	}

	if u.DigestStepUpsertDto != nil {
		return utils.MarshalJSON(u.DigestStepUpsertDto, "", true)
	}

	if u.CustomStepUpsertDto != nil {
		return utils.MarshalJSON(u.CustomStepUpsertDto, "", true)
	}

	return nil, errors.New("could not marshal union type Steps: all fields are null")
}

type CreateWorkflowDto struct {
	// Name of the workflow
	Name string `json:"name"`
	// Description of the workflow
	Description *string `json:"description,omitempty"`
	// Tags associated with the workflow
	Tags []string `json:"tags,omitempty"`
	// Whether the workflow is active
	Active *bool `default:"false" json:"active"`
	// Unique identifier for the workflow
	WorkflowID string `json:"workflowId"`
	// Steps of the workflow
	Steps []Steps `json:"steps"`
	// Source of workflow creation
	Source *WorkflowCreationSourceEnum `default:"editor" json:"__source"`
	// Workflow preferences
	Preferences *PreferencesRequestDto `json:"preferences,omitempty"`
	// The payload JSON Schema for the workflow
	PayloadSchema map[string]any `json:"payloadSchema,omitempty"`
	// Enable or disable payload schema validation
	ValidatePayload *bool `json:"validatePayload,omitempty"`
}

func (c CreateWorkflowDto) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateWorkflowDto) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateWorkflowDto) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateWorkflowDto) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateWorkflowDto) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateWorkflowDto) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *CreateWorkflowDto) GetWorkflowID() string {
	if o == nil {
		return ""
	}
	return o.WorkflowID
}

func (o *CreateWorkflowDto) GetSteps() []Steps {
	if o == nil {
		return []Steps{}
	}
	return o.Steps
}

func (o *CreateWorkflowDto) GetSource() *WorkflowCreationSourceEnum {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *CreateWorkflowDto) GetPreferences() *PreferencesRequestDto {
	if o == nil {
		return nil
	}
	return o.Preferences
}

func (o *CreateWorkflowDto) GetPayloadSchema() map[string]any {
	if o == nil {
		return nil
	}
	return o.PayloadSchema
}

func (o *CreateWorkflowDto) GetValidatePayload() *bool {
	if o == nil {
		return nil
	}
	return o.ValidatePayload
}
