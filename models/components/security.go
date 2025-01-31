// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Security struct {
	SecretKey *string `security:"scheme,type=apiKey,subtype=header,name=Authorization,env=novu_secret_key"`
}

func (o *Security) GetSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.SecretKey
}
