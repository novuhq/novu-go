// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateTopicRequestDto struct {
	// User defined custom key and provided by the user that will be an unique identifier for the Topic created.
	Key string `json:"key"`
	// User defined custom name and provided by the user that will name the Topic created.
	Name string `json:"name"`
}

func (o *CreateTopicRequestDto) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *CreateTopicRequestDto) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
