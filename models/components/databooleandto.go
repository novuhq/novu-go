// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DataBooleanDto struct {
	Data bool `json:"data"`
}

func (o *DataBooleanDto) GetData() bool {
	if o == nil {
		return false
	}
	return o.Data
}
