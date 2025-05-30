// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type WorkflowResponseData struct {
}

type WorkflowIntegrationStatus struct {
}

type WorkflowResponse struct {
	ID                        *string                      `json:"_id,omitempty"`
	Name                      string                       `json:"name"`
	Description               string                       `json:"description"`
	Active                    bool                         `json:"active"`
	Draft                     bool                         `json:"draft"`
	PreferenceSettings        SubscriberPreferenceChannels `json:"preferenceSettings"`
	Critical                  bool                         `json:"critical"`
	Tags                      []string                     `json:"tags"`
	Steps                     []NotificationStepDto        `json:"steps"`
	OrganizationID            string                       `json:"_organizationId"`
	CreatorID                 string                       `json:"_creatorId"`
	EnvironmentID             string                       `json:"_environmentId"`
	Triggers                  []NotificationTrigger        `json:"triggers"`
	NotificationGroupID       string                       `json:"_notificationGroupId"`
	ParentID                  *string                      `json:"_parentId,omitempty"`
	Deleted                   bool                         `json:"deleted"`
	DeletedAt                 string                       `json:"deletedAt"`
	DeletedBy                 string                       `json:"deletedBy"`
	NotificationGroup         *NotificationGroup           `json:"notificationGroup,omitempty"`
	Data                      *WorkflowResponseData        `json:"data,omitempty"`
	WorkflowIntegrationStatus *WorkflowIntegrationStatus   `json:"workflowIntegrationStatus,omitempty"`
}

func (o *WorkflowResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WorkflowResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *WorkflowResponse) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *WorkflowResponse) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *WorkflowResponse) GetDraft() bool {
	if o == nil {
		return false
	}
	return o.Draft
}

func (o *WorkflowResponse) GetPreferenceSettings() SubscriberPreferenceChannels {
	if o == nil {
		return SubscriberPreferenceChannels{}
	}
	return o.PreferenceSettings
}

func (o *WorkflowResponse) GetCritical() bool {
	if o == nil {
		return false
	}
	return o.Critical
}

func (o *WorkflowResponse) GetTags() []string {
	if o == nil {
		return []string{}
	}
	return o.Tags
}

func (o *WorkflowResponse) GetSteps() []NotificationStepDto {
	if o == nil {
		return []NotificationStepDto{}
	}
	return o.Steps
}

func (o *WorkflowResponse) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *WorkflowResponse) GetCreatorID() string {
	if o == nil {
		return ""
	}
	return o.CreatorID
}

func (o *WorkflowResponse) GetEnvironmentID() string {
	if o == nil {
		return ""
	}
	return o.EnvironmentID
}

func (o *WorkflowResponse) GetTriggers() []NotificationTrigger {
	if o == nil {
		return []NotificationTrigger{}
	}
	return o.Triggers
}

func (o *WorkflowResponse) GetNotificationGroupID() string {
	if o == nil {
		return ""
	}
	return o.NotificationGroupID
}

func (o *WorkflowResponse) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *WorkflowResponse) GetDeleted() bool {
	if o == nil {
		return false
	}
	return o.Deleted
}

func (o *WorkflowResponse) GetDeletedAt() string {
	if o == nil {
		return ""
	}
	return o.DeletedAt
}

func (o *WorkflowResponse) GetDeletedBy() string {
	if o == nil {
		return ""
	}
	return o.DeletedBy
}

func (o *WorkflowResponse) GetNotificationGroup() *NotificationGroup {
	if o == nil {
		return nil
	}
	return o.NotificationGroup
}

func (o *WorkflowResponse) GetData() *WorkflowResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *WorkflowResponse) GetWorkflowIntegrationStatus() *WorkflowIntegrationStatus {
	if o == nil {
		return nil
	}
	return o.WorkflowIntegrationStatus
}
