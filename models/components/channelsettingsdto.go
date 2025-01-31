// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ChannelSettingsDtoProviderID - The provider identifier for the credentials
type ChannelSettingsDtoProviderID string

const (
	ChannelSettingsDtoProviderIDSlack            ChannelSettingsDtoProviderID = "slack"
	ChannelSettingsDtoProviderIDDiscord          ChannelSettingsDtoProviderID = "discord"
	ChannelSettingsDtoProviderIDMsteams          ChannelSettingsDtoProviderID = "msteams"
	ChannelSettingsDtoProviderIDMattermost       ChannelSettingsDtoProviderID = "mattermost"
	ChannelSettingsDtoProviderIDRyver            ChannelSettingsDtoProviderID = "ryver"
	ChannelSettingsDtoProviderIDZulip            ChannelSettingsDtoProviderID = "zulip"
	ChannelSettingsDtoProviderIDGrafanaOnCall    ChannelSettingsDtoProviderID = "grafana-on-call"
	ChannelSettingsDtoProviderIDGetstream        ChannelSettingsDtoProviderID = "getstream"
	ChannelSettingsDtoProviderIDRocketChat       ChannelSettingsDtoProviderID = "rocket-chat"
	ChannelSettingsDtoProviderIDWhatsappBusiness ChannelSettingsDtoProviderID = "whatsapp-business"
	ChannelSettingsDtoProviderIDFcm              ChannelSettingsDtoProviderID = "fcm"
	ChannelSettingsDtoProviderIDApns             ChannelSettingsDtoProviderID = "apns"
	ChannelSettingsDtoProviderIDExpo             ChannelSettingsDtoProviderID = "expo"
	ChannelSettingsDtoProviderIDOneSignal        ChannelSettingsDtoProviderID = "one-signal"
	ChannelSettingsDtoProviderIDPushpad          ChannelSettingsDtoProviderID = "pushpad"
	ChannelSettingsDtoProviderIDPushWebhook      ChannelSettingsDtoProviderID = "push-webhook"
	ChannelSettingsDtoProviderIDPusherBeams      ChannelSettingsDtoProviderID = "pusher-beams"
)

func (e ChannelSettingsDtoProviderID) ToPointer() *ChannelSettingsDtoProviderID {
	return &e
}
func (e *ChannelSettingsDtoProviderID) UnmarshalJSON(data []byte) error {
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
		*e = ChannelSettingsDtoProviderID(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChannelSettingsDtoProviderID: %v", v)
	}
}

type ChannelSettingsDto struct {
	// The provider identifier for the credentials
	ProviderID ChannelSettingsDtoProviderID `json:"providerId"`
	// The integration identifier
	IntegrationIdentifier *string `json:"integrationIdentifier,omitempty"`
	// Credentials payload for the specified provider
	Credentials ChannelCredentials `json:"credentials"`
	// The unique identifier of the integration associated with this channel.
	IntegrationID string `json:"_integrationId"`
}

func (o *ChannelSettingsDto) GetProviderID() ChannelSettingsDtoProviderID {
	if o == nil {
		return ChannelSettingsDtoProviderID("")
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
