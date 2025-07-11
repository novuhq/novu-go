// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ControlsMetadataDto struct {
	// JSON Schema for data
	DataSchema map[string]any `json:"dataSchema,omitempty"`
	// UI Schema for rendering
	UISchema *UISchema `json:"uiSchema,omitempty"`
}

func (o *ControlsMetadataDto) GetDataSchema() map[string]any {
	if o == nil {
		return nil
	}
	return o.DataSchema
}

func (o *ControlsMetadataDto) GetUISchema() *UISchema {
	if o == nil {
		return nil
	}
	return o.UISchema
}
