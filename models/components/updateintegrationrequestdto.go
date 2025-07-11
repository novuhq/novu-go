// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UpdateIntegrationRequestDto struct {
	Name          *string `json:"name,omitempty"`
	Identifier    *string `json:"identifier,omitempty"`
	EnvironmentID *string `json:"_environmentId,omitempty"`
	// If the integration is active the validation on the credentials field will run
	Active      *bool           `json:"active,omitempty"`
	Credentials *CredentialsDto `json:"credentials,omitempty"`
	Check       *bool           `json:"check,omitempty"`
	Conditions  []StepFilterDto `json:"conditions,omitempty"`
}

func (o *UpdateIntegrationRequestDto) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateIntegrationRequestDto) GetIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *UpdateIntegrationRequestDto) GetEnvironmentID() *string {
	if o == nil {
		return nil
	}
	return o.EnvironmentID
}

func (o *UpdateIntegrationRequestDto) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *UpdateIntegrationRequestDto) GetCredentials() *CredentialsDto {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *UpdateIntegrationRequestDto) GetCheck() *bool {
	if o == nil {
		return nil
	}
	return o.Check
}

func (o *UpdateIntegrationRequestDto) GetConditions() []StepFilterDto {
	if o == nil {
		return nil
	}
	return o.Conditions
}
