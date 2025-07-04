// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UISchema struct {
	// Group of the UI Schema
	Group *UISchemaGroupEnum `json:"group,omitempty"`
	// Properties of the UI Schema
	Properties map[string]UISchemaProperty `json:"properties,omitempty"`
}

func (o *UISchema) GetGroup() *UISchemaGroupEnum {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *UISchema) GetProperties() map[string]UISchemaProperty {
	if o == nil {
		return nil
	}
	return o.Properties
}
