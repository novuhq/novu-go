// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateEnvironmentRequestDto struct {
	// Name of the environment to be created
	Name string `json:"name"`
	// MongoDB ObjectId of the parent environment (optional)
	ParentID *string `json:"parentId,omitempty"`
	// Hex color code for the environment
	Color string `json:"color"`
}

func (o *CreateEnvironmentRequestDto) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateEnvironmentRequestDto) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *CreateEnvironmentRequestDto) GetColor() string {
	if o == nil {
		return ""
	}
	return o.Color
}
