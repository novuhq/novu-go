// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DigestControlsMetadataResponseDto struct {
	// JSON Schema for data
	DataSchema map[string]any `json:"dataSchema,omitempty"`
	// UI Schema for rendering
	UISchema *UISchema `json:"uiSchema,omitempty"`
	// Control values specific to Digest
	Values DigestControlDto `json:"values"`
}

func (o *DigestControlsMetadataResponseDto) GetDataSchema() map[string]any {
	if o == nil {
		return nil
	}
	return o.DataSchema
}

func (o *DigestControlsMetadataResponseDto) GetUISchema() *UISchema {
	if o == nil {
		return nil
	}
	return o.UISchema
}

func (o *DigestControlsMetadataResponseDto) GetValues() DigestControlDto {
	if o == nil {
		return DigestControlDto{}
	}
	return o.Values
}
