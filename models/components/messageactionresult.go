// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// MessageActionResultPayload - Payload of the action result
type MessageActionResultPayload struct {
}

type MessageActionResult struct {
	// Payload of the action result
	Payload *MessageActionResultPayload `json:"payload,omitempty"`
	// Type of button for the action result
	Type *ButtonTypeEnum `json:"type,omitempty"`
}

func (o *MessageActionResult) GetPayload() *MessageActionResultPayload {
	if o == nil {
		return nil
	}
	return o.Payload
}

func (o *MessageActionResult) GetType() *ButtonTypeEnum {
	if o == nil {
		return nil
	}
	return o.Type
}
