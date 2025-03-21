// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ProviderID - The provider identifier for the credentials
type ProviderID string

const (
	ProviderIDSlack            ProviderID = "slack"
	ProviderIDDiscord          ProviderID = "discord"
	ProviderIDMsteams          ProviderID = "msteams"
	ProviderIDMattermost       ProviderID = "mattermost"
	ProviderIDRyver            ProviderID = "ryver"
	ProviderIDZulip            ProviderID = "zulip"
	ProviderIDGrafanaOnCall    ProviderID = "grafana-on-call"
	ProviderIDGetstream        ProviderID = "getstream"
	ProviderIDRocketChat       ProviderID = "rocket-chat"
	ProviderIDWhatsappBusiness ProviderID = "whatsapp-business"
	ProviderIDFcm              ProviderID = "fcm"
	ProviderIDApns             ProviderID = "apns"
	ProviderIDExpo             ProviderID = "expo"
	ProviderIDOneSignal        ProviderID = "one-signal"
	ProviderIDPushpad          ProviderID = "pushpad"
	ProviderIDPushWebhook      ProviderID = "push-webhook"
	ProviderIDPusherBeams      ProviderID = "pusher-beams"
)

func (e ProviderID) ToPointer() *ProviderID {
	return &e
}
func (e *ProviderID) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "slack":
		fallthrough
	case "discord":
		fallthrough
	case "msteams":
		fallthrough
	case "mattermost":
		fallthrough
	case "ryver":
		fallthrough
	case "zulip":
		fallthrough
	case "grafana-on-call":
		fallthrough
	case "getstream":
		fallthrough
	case "rocket-chat":
		fallthrough
	case "whatsapp-business":
		fallthrough
	case "fcm":
		fallthrough
	case "apns":
		fallthrough
	case "expo":
		fallthrough
	case "one-signal":
		fallthrough
	case "pushpad":
		fallthrough
	case "push-webhook":
		fallthrough
	case "pusher-beams":
		*e = ProviderID(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProviderID: %v", v)
	}
}

type ChannelSettingsDto struct {
	// The provider identifier for the credentials
	ProviderID ProviderID `json:"providerId"`
	// The integration identifier
	IntegrationIdentifier *string `json:"integrationIdentifier,omitempty"`
	// Credentials payload for the specified provider
	Credentials ChannelCredentials `json:"credentials"`
	// The unique identifier of the integration associated with this channel.
	IntegrationID string `json:"_integrationId"`
}

func (o *ChannelSettingsDto) GetProviderID() ProviderID {
	if o == nil {
		return ProviderID("")
	}
	return o.ProviderID
}

func (o *ChannelSettingsDto) GetIntegrationIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationIdentifier
}

func (o *ChannelSettingsDto) GetCredentials() ChannelCredentials {
	if o == nil {
		return ChannelCredentials{}
	}
	return o.Credentials
}

func (o *ChannelSettingsDto) GetIntegrationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationID
}
