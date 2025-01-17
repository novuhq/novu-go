// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/novuhq/novu-go/models/components"
)

type IntegrationsControllerListIntegrationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The list of integrations belonging to the organization that are successfully returned.
	IntegrationResponseDtos []components.IntegrationResponseDto
	Headers                 map[string][]string
}

func (o *IntegrationsControllerListIntegrationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *IntegrationsControllerListIntegrationsResponse) GetIntegrationResponseDtos() []components.IntegrationResponseDto {
	if o == nil {
		return nil
	}
	return o.IntegrationResponseDtos
}

func (o *IntegrationsControllerListIntegrationsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
