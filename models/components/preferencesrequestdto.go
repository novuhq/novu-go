// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/novuhq/novu-go/internal/utils"
)

type UserAllType string

const (
	UserAllTypeWorkflowPreferenceDto UserAllType = "WorkflowPreferenceDto"
)

// UserAll - A preference for the workflow. The values specified here will be used if no preference is specified for a channel.
type UserAll struct {
	WorkflowPreferenceDto *WorkflowPreferenceDto `queryParam:"inline"`

	Type UserAllType
}

func CreateUserAllWorkflowPreferenceDto(workflowPreferenceDto WorkflowPreferenceDto) UserAll {
	typ := UserAllTypeWorkflowPreferenceDto

	return UserAll{
		WorkflowPreferenceDto: &workflowPreferenceDto,
		Type:                  typ,
	}
}

func (u *UserAll) UnmarshalJSON(data []byte) error {

	var workflowPreferenceDto WorkflowPreferenceDto = WorkflowPreferenceDto{}
	if err := utils.UnmarshalJSON(data, &workflowPreferenceDto, "", true, true); err == nil {
		u.WorkflowPreferenceDto = &workflowPreferenceDto
		u.Type = UserAllTypeWorkflowPreferenceDto
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UserAll", string(data))
}

func (u UserAll) MarshalJSON() ([]byte, error) {
	if u.WorkflowPreferenceDto != nil {
		return utils.MarshalJSON(u.WorkflowPreferenceDto, "", true)
	}

	return nil, errors.New("could not marshal union type UserAll: all fields are null")
}

type UserWorkflowPreferencesDto struct {
	// A preference for the workflow. The values specified here will be used if no preference is specified for a channel.
	All UserAll `json:"all"`
	// Preferences for different communication channels
	Channels map[string]ChannelPreferenceDto `json:"channels"`
}

func (o *UserWorkflowPreferencesDto) GetAll() UserAll {
	if o == nil {
		return UserAll{}
	}
	return o.All
}

func (o *UserWorkflowPreferencesDto) GetChannels() map[string]ChannelPreferenceDto {
	if o == nil {
		return map[string]ChannelPreferenceDto{}
	}
	return o.Channels
}

type UserType string

const (
	UserTypeUserWorkflowPreferencesDto UserType = "user_WorkflowPreferencesDto"
)

// User workflow preferences
type User struct {
	UserWorkflowPreferencesDto *UserWorkflowPreferencesDto `queryParam:"inline"`

	Type UserType
}

func CreateUserUserWorkflowPreferencesDto(userWorkflowPreferencesDto UserWorkflowPreferencesDto) User {
	typ := UserTypeUserWorkflowPreferencesDto

	return User{
		UserWorkflowPreferencesDto: &userWorkflowPreferencesDto,
		Type:                       typ,
	}
}

func (u *User) UnmarshalJSON(data []byte) error {

	var userWorkflowPreferencesDto UserWorkflowPreferencesDto = UserWorkflowPreferencesDto{}
	if err := utils.UnmarshalJSON(data, &userWorkflowPreferencesDto, "", true, true); err == nil {
		u.UserWorkflowPreferencesDto = &userWorkflowPreferencesDto
		u.Type = UserTypeUserWorkflowPreferencesDto
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for User", string(data))
}

func (u User) MarshalJSON() ([]byte, error) {
	if u.UserWorkflowPreferencesDto != nil {
		return utils.MarshalJSON(u.UserWorkflowPreferencesDto, "", true)
	}

	return nil, errors.New("could not marshal union type User: all fields are null")
}

type PreferencesRequestDtoAllType string

const (
	PreferencesRequestDtoAllTypeWorkflowPreferenceDto PreferencesRequestDtoAllType = "WorkflowPreferenceDto"
)

// PreferencesRequestDtoAll - A preference for the workflow. The values specified here will be used if no preference is specified for a channel.
type PreferencesRequestDtoAll struct {
	WorkflowPreferenceDto *WorkflowPreferenceDto `queryParam:"inline"`

	Type PreferencesRequestDtoAllType
}

func CreatePreferencesRequestDtoAllWorkflowPreferenceDto(workflowPreferenceDto WorkflowPreferenceDto) PreferencesRequestDtoAll {
	typ := PreferencesRequestDtoAllTypeWorkflowPreferenceDto

	return PreferencesRequestDtoAll{
		WorkflowPreferenceDto: &workflowPreferenceDto,
		Type:                  typ,
	}
}

func (u *PreferencesRequestDtoAll) UnmarshalJSON(data []byte) error {

	var workflowPreferenceDto WorkflowPreferenceDto = WorkflowPreferenceDto{}
	if err := utils.UnmarshalJSON(data, &workflowPreferenceDto, "", true, true); err == nil {
		u.WorkflowPreferenceDto = &workflowPreferenceDto
		u.Type = PreferencesRequestDtoAllTypeWorkflowPreferenceDto
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for PreferencesRequestDtoAll", string(data))
}

func (u PreferencesRequestDtoAll) MarshalJSON() ([]byte, error) {
	if u.WorkflowPreferenceDto != nil {
		return utils.MarshalJSON(u.WorkflowPreferenceDto, "", true)
	}

	return nil, errors.New("could not marshal union type PreferencesRequestDtoAll: all fields are null")
}

// Workflow - Workflow-specific preferences
type Workflow struct {
	// A preference for the workflow. The values specified here will be used if no preference is specified for a channel.
	All PreferencesRequestDtoAll `json:"all"`
	// Preferences for different communication channels
	Channels map[string]ChannelPreferenceDto `json:"channels"`
}

func (o *Workflow) GetAll() PreferencesRequestDtoAll {
	if o == nil {
		return PreferencesRequestDtoAll{}
	}
	return o.All
}

func (o *Workflow) GetChannels() map[string]ChannelPreferenceDto {
	if o == nil {
		return map[string]ChannelPreferenceDto{}
	}
	return o.Channels
}

type PreferencesRequestDto struct {
	// User workflow preferences
	User *User `json:"user,omitempty"`
	// Workflow-specific preferences
	Workflow *Workflow `json:"workflow,omitempty"`
}

func (o *PreferencesRequestDto) GetUser() *User {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *PreferencesRequestDto) GetWorkflow() *Workflow {
	if o == nil {
		return nil
	}
	return o.Workflow
}
