// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UpdateSubscriberPreferenceResponseDto struct {
	// The workflow information and if it is critical or not
	Template TemplateResponse `json:"template"`
	// The preferences of the subscriber regarding the related workflow
	Preference Preference `json:"preference"`
}

func (o *UpdateSubscriberPreferenceResponseDto) GetTemplate() TemplateResponse {
	if o == nil {
		return TemplateResponse{}
	}
	return o.Template
}

func (o *UpdateSubscriberPreferenceResponseDto) GetPreference() Preference {
	if o == nil {
		return Preference{}
	}
	return o.Preference
}
