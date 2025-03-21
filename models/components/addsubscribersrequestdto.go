// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AddSubscribersRequestDto struct {
	// List of subscriber identifiers that will be associated to the topic
	Subscribers []string `json:"subscribers"`
}

func (o *AddSubscribersRequestDto) GetSubscribers() []string {
	if o == nil {
		return []string{}
	}
	return o.Subscribers
}
