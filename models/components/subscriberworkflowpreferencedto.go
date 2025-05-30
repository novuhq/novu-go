// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type SubscriberWorkflowPreferenceDto struct {
	// Whether notifications are enabled for this workflow
	Enabled bool `json:"enabled"`
	// Channel-specific preference settings for this workflow
	Channels SubscriberPreferenceChannels `json:"channels"`
	// List of preference overrides
	Overrides []SubscriberPreferenceOverrideDto `json:"overrides"`
	// Workflow information
	Workflow SubscriberPreferencesWorkflowInfoDto `json:"workflow"`
}

func (o *SubscriberWorkflowPreferenceDto) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *SubscriberWorkflowPreferenceDto) GetChannels() SubscriberPreferenceChannels {
	if o == nil {
		return SubscriberPreferenceChannels{}
	}
	return o.Channels
}

func (o *SubscriberWorkflowPreferenceDto) GetOverrides() []SubscriberPreferenceOverrideDto {
	if o == nil {
		return []SubscriberPreferenceOverrideDto{}
	}
	return o.Overrides
}

func (o *SubscriberWorkflowPreferenceDto) GetWorkflow() SubscriberPreferencesWorkflowInfoDto {
	if o == nil {
		return SubscriberPreferencesWorkflowInfoDto{}
	}
	return o.Workflow
}
