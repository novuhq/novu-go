// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// StepResponseDtoSlug - Slug of the step
type StepResponseDtoSlug struct {
}

type StepResponseDto struct {
	// Controls metadata for the step
	Controls ControlsMetadataDto `json:"controls"`
	// Control values for the step (alias for controls.values)
	ControlValues map[string]any `json:"controlValues,omitempty"`
	// JSON Schema for variables, follows the JSON Schema standard
	Variables map[string]any `json:"variables"`
	// Unique identifier of the step
	StepID string `json:"stepId"`
	// Database identifier of the step
	ID string `json:"_id"`
	// Name of the step
	Name string `json:"name"`
	// Slug of the step
	Slug StepResponseDtoSlug `json:"slug"`
	// Type of the step
	Type StepTypeEnum `json:"type"`
	// Origin of the workflow
	Origin WorkflowOriginEnum `json:"origin"`
	// Workflow identifier
	WorkflowID string `json:"workflowId"`
	// Workflow database identifier
	WorkflowDatabaseID string `json:"workflowDatabaseId"`
	// Issues associated with the step
	Issues *StepIssuesDto `json:"issues,omitempty"`
}

func (o *StepResponseDto) GetControls() ControlsMetadataDto {
	if o == nil {
		return ControlsMetadataDto{}
	}
	return o.Controls
}

func (o *StepResponseDto) GetControlValues() map[string]any {
	if o == nil {
		return nil
	}
	return o.ControlValues
}

func (o *StepResponseDto) GetVariables() map[string]any {
	if o == nil {
		return map[string]any{}
	}
	return o.Variables
}

func (o *StepResponseDto) GetStepID() string {
	if o == nil {
		return ""
	}
	return o.StepID
}

func (o *StepResponseDto) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *StepResponseDto) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *StepResponseDto) GetSlug() StepResponseDtoSlug {
	if o == nil {
		return StepResponseDtoSlug{}
	}
	return o.Slug
}

func (o *StepResponseDto) GetType() StepTypeEnum {
	if o == nil {
		return StepTypeEnum("")
	}
	return o.Type
}

func (o *StepResponseDto) GetOrigin() WorkflowOriginEnum {
	if o == nil {
		return WorkflowOriginEnum("")
	}
	return o.Origin
}

func (o *StepResponseDto) GetWorkflowID() string {
	if o == nil {
		return ""
	}
	return o.WorkflowID
}

func (o *StepResponseDto) GetWorkflowDatabaseID() string {
	if o == nil {
		return ""
	}
	return o.WorkflowDatabaseID
}

func (o *StepResponseDto) GetIssues() *StepIssuesDto {
	if o == nil {
		return nil
	}
	return o.Issues
}
