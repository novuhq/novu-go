// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type EmailBlock struct {
	// Type of the email block
	Type EmailBlockTypeEnum `json:"type"`
	// Content of the email block
	Content string `json:"content"`
	// URL associated with the email block, if any
	URL *string `json:"url,omitempty"`
	// Styles applied to the email block
	Styles *EmailBlockStyles `json:"styles,omitempty"`
}

func (o *EmailBlock) GetType() EmailBlockTypeEnum {
	if o == nil {
		return EmailBlockTypeEnum("")
	}
	return o.Type
}

func (o *EmailBlock) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *EmailBlock) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *EmailBlock) GetStyles() *EmailBlockStyles {
	if o == nil {
		return nil
	}
	return o.Styles
}
