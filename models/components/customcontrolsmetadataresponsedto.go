// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CustomControlsMetadataResponseDto struct {
	// JSON Schema for data
	DataSchema map[string]any `json:"dataSchema,omitempty"`
	// UI Schema for rendering
	UISchema *UISchema `json:"uiSchema,omitempty"`
	// Control values specific to Custom step
	Values CustomControlDto `json:"values"`
}

func (o *CustomControlsMetadataResponseDto) GetDataSchema() map[string]any {
	if o == nil {
		return nil
	}
	return o.DataSchema
}

func (o *CustomControlsMetadataResponseDto) GetUISchema() *UISchema {
	if o == nil {
		return nil
	}
	return o.UISchema
}

func (o *CustomControlsMetadataResponseDto) GetValues() CustomControlDto {
	if o == nil {
		return CustomControlDto{}
	}
	return o.Values
}
