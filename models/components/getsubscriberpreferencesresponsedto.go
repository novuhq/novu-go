// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type GetSubscriberPreferencesResponseDto struct {
	// The workflow information and if it is critical or not
	Template *TemplateResponse `json:"template,omitempty"`
	// The preferences of the subscriber regarding the related workflow
	Preference Preference `json:"preference"`
}

func (o *GetSubscriberPreferencesResponseDto) GetTemplate() *TemplateResponse {
	if o == nil {
		return nil
	}
	return o.Template
}

func (o *GetSubscriberPreferencesResponseDto) GetPreference() Preference {
	if o == nil {
		return Preference{}
	}
	return o.Preference
}
