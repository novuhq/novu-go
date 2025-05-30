// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type MetaDto struct {
	// The total count of subscriber IDs provided
	TotalCount float64 `json:"totalCount"`
	// The count of successfully created subscriptions
	Successful float64 `json:"successful"`
	// The count of failed subscription attempts
	Failed float64 `json:"failed"`
}

func (o *MetaDto) GetTotalCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.TotalCount
}

func (o *MetaDto) GetSuccessful() float64 {
	if o == nil {
		return 0.0
	}
	return o.Successful
}

func (o *MetaDto) GetFailed() float64 {
	if o == nil {
		return 0.0
	}
	return o.Failed
}
