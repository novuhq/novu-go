// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type SyncWorkflowDto struct {
	// Target environment identifier to sync the workflow to
	TargetEnvironmentID string `json:"targetEnvironmentId"`
}

func (o *SyncWorkflowDto) GetTargetEnvironmentID() string {
	if o == nil {
		return ""
	}
	return o.TargetEnvironmentID
}
