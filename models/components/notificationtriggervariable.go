// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type NotificationTriggerVariable struct {
	// Name of the variable
	Name string `json:"name"`
}

func (o *NotificationTriggerVariable) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
