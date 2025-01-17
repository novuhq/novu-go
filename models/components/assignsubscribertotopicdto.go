// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AssignSubscriberToTopicDto struct {
	// List of successfully assigned subscriber IDs
	Succeeded []string `json:"succeeded"`
	// Details about failed assignments
	Failed *FailedAssignmentsDto `json:"failed,omitempty"`
}

func (o *AssignSubscriberToTopicDto) GetSucceeded() []string {
	if o == nil {
		return []string{}
	}
	return o.Succeeded
}

func (o *AssignSubscriberToTopicDto) GetFailed() *FailedAssignmentsDto {
	if o == nil {
		return nil
	}
	return o.Failed
}
