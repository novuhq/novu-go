// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type Channels string

const (
	ChannelsInApp Channels = "in_app"
	ChannelsEmail Channels = "email"
	ChannelsSms   Channels = "sms"
	ChannelsChat  Channels = "chat"
	ChannelsPush  Channels = "push"
)

func (e Channels) ToPointer() *Channels {
	return &e
}
func (e *Channels) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "in_app":
		fallthrough
	case "email":
		fallthrough
	case "sms":
		fallthrough
	case "chat":
		fallthrough
	case "push":
		*e = Channels(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Channels: %v", v)
	}
}

type ActivityGraphStatesResponse struct {
	ID        string     `json:"_id"`
	Count     float64    `json:"count"`
	Templates []string   `json:"templates"`
	Channels  []Channels `json:"channels"`
}

func (o *ActivityGraphStatesResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ActivityGraphStatesResponse) GetCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Count
}

func (o *ActivityGraphStatesResponse) GetTemplates() []string {
	if o == nil {
		return []string{}
	}
	return o.Templates
}

func (o *ActivityGraphStatesResponse) GetChannels() []Channels {
	if o == nil {
		return []Channels{}
	}
	return o.Channels
}
