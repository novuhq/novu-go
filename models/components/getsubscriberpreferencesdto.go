// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type GetSubscriberPreferencesDto struct {
	// Global preference settings
	Global GlobalPreferenceDto `json:"global"`
	// Workflow-specific preference settings
	Workflows []WorkflowPreferenceDto `json:"workflows"`
}

func (o *GetSubscriberPreferencesDto) GetGlobal() GlobalPreferenceDto {
	if o == nil {
		return GlobalPreferenceDto{}
	}
	return o.Global
}

func (o *GetSubscriberPreferencesDto) GetWorkflows() []WorkflowPreferenceDto {
	if o == nil {
		return []WorkflowPreferenceDto{}
	}
	return o.Workflows
}
