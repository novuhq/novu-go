// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UpdateSubscriberChannelRequestDto struct {
	// The provider identifier for the credentials
	ProviderID ChatOrPushProviderEnum `json:"providerId"`
	// The integration identifier
	IntegrationIdentifier *string `json:"integrationIdentifier,omitempty"`
	// Credentials payload for the specified provider
	Credentials ChannelCredentials `json:"credentials"`
}

func (o *UpdateSubscriberChannelRequestDto) GetProviderID() ChatOrPushProviderEnum {
	if o == nil {
		return ChatOrPushProviderEnum("")
	}
	return o.ProviderID
}

func (o *UpdateSubscriberChannelRequestDto) GetIntegrationIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationIdentifier
}

func (o *UpdateSubscriberChannelRequestDto) GetCredentials() ChannelCredentials {
	if o == nil {
		return ChannelCredentials{}
	}
	return o.Credentials
}
