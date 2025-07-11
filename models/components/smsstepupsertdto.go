// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// SmsStepUpsertDtoControlValues - Control values for the SMS step
type SmsStepUpsertDtoControlValues struct {
	// JSONLogic filter conditions for conditionally skipping the step execution. Supports complex logical operations with AND, OR, and comparison operators. See https://jsonlogic.com/ for full typing reference.
	Skip map[string]any `json:"skip,omitempty"`
	// Content of the SMS message.
	Body *string `json:"body,omitempty"`
}

func (o *SmsStepUpsertDtoControlValues) GetSkip() map[string]any {
	if o == nil {
		return nil
	}
	return o.Skip
}

func (o *SmsStepUpsertDtoControlValues) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

type SmsStepUpsertDto struct {
	// Unique identifier of the step
	ID *string `json:"_id,omitempty"`
	// Name of the step
	Name string `json:"name"`
	// Type of the step
	Type StepTypeEnum `json:"type"`
	// Control values for the SMS step
	ControlValues *SmsStepUpsertDtoControlValues `json:"controlValues,omitempty"`
}

func (o *SmsStepUpsertDto) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SmsStepUpsertDto) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SmsStepUpsertDto) GetType() StepTypeEnum {
	if o == nil {
		return StepTypeEnum("")
	}
	return o.Type
}

func (o *SmsStepUpsertDto) GetControlValues() *SmsStepUpsertDtoControlValues {
	if o == nil {
		return nil
	}
	return o.ControlValues
}
