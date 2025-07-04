// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/novuhq/novu-go/internal/utils"
)

// ControlValues - Control values for the In-App step
type ControlValues struct {
	// JSONLogic filter conditions for conditionally skipping the step execution. Supports complex logical operations with AND, OR, and comparison operators. See https://jsonlogic.com/ for full typing reference.
	Skip map[string]any `json:"skip,omitempty"`
	// Content/body of the in-app message. Required if subject is empty.
	Body *string `json:"body,omitempty"`
	// Subject/title of the in-app message. Required if body is empty.
	Subject *string `json:"subject,omitempty"`
	// URL for an avatar image. Must be a valid URL or start with / or {{"{{"}} variable }}.
	Avatar *string `json:"avatar,omitempty"`
	// Primary action button details.
	PrimaryAction *ActionDto `json:"primaryAction,omitempty"`
	// Secondary action button details.
	SecondaryAction *ActionDto `json:"secondaryAction,omitempty"`
	// Redirection URL configuration for the main content click (if no actions defined/clicked)..
	Redirect *RedirectDto `json:"redirect,omitempty"`
	// Disable sanitization of the output.
	DisableOutputSanitization *bool `default:"false" json:"disableOutputSanitization"`
	// Additional data payload for the step.
	Data map[string]any `json:"data,omitempty"`
}

func (c ControlValues) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ControlValues) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ControlValues) GetSkip() map[string]any {
	if o == nil {
		return nil
	}
	return o.Skip
}

func (o *ControlValues) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *ControlValues) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *ControlValues) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *ControlValues) GetPrimaryAction() *ActionDto {
	if o == nil {
		return nil
	}
	return o.PrimaryAction
}

func (o *ControlValues) GetSecondaryAction() *ActionDto {
	if o == nil {
		return nil
	}
	return o.SecondaryAction
}

func (o *ControlValues) GetRedirect() *RedirectDto {
	if o == nil {
		return nil
	}
	return o.Redirect
}

func (o *ControlValues) GetDisableOutputSanitization() *bool {
	if o == nil {
		return nil
	}
	return o.DisableOutputSanitization
}

func (o *ControlValues) GetData() map[string]any {
	if o == nil {
		return nil
	}
	return o.Data
}

type InAppStepUpsertDto struct {
	// Unique identifier of the step
	ID *string `json:"_id,omitempty"`
	// Name of the step
	Name string `json:"name"`
	// Type of the step
	Type StepTypeEnum `json:"type"`
	// Control values for the In-App step
	ControlValues *ControlValues `json:"controlValues,omitempty"`
}

func (o *InAppStepUpsertDto) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InAppStepUpsertDto) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InAppStepUpsertDto) GetType() StepTypeEnum {
	if o == nil {
		return StepTypeEnum("")
	}
	return o.Type
}

func (o *InAppStepUpsertDto) GetControlValues() *ControlValues {
	if o == nil {
		return nil
	}
	return o.ControlValues
}
