// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type BulkTriggerEventDto struct {
	Events []TriggerEventRequestDto `json:"events"`
}

func (o *BulkTriggerEventDto) GetEvents() []TriggerEventRequestDto {
	if o == nil {
		return []TriggerEventRequestDto{}
	}
	return o.Events
}
