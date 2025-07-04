// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type SmsControlsMetadataResponseDto struct {
	// JSON Schema for data
	DataSchema map[string]any `json:"dataSchema,omitempty"`
	// UI Schema for rendering
	UISchema *UISchema `json:"uiSchema,omitempty"`
	// Control values specific to SMS
	Values SmsControlDto `json:"values"`
}

func (o *SmsControlsMetadataResponseDto) GetDataSchema() map[string]any {
	if o == nil {
		return nil
	}
	return o.DataSchema
}

func (o *SmsControlsMetadataResponseDto) GetUISchema() *UISchema {
	if o == nil {
		return nil
	}
	return o.UISchema
}

func (o *SmsControlsMetadataResponseDto) GetValues() SmsControlDto {
	if o == nil {
		return SmsControlDto{}
	}
	return o.Values
}
