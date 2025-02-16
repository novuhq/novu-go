// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ActivityNotificationExecutionDetailResponseDto struct {
	// Unique identifier of the execution detail
	ID string `json:"_id"`
	// Creation time of the execution detail
	CreatedAt *string `json:"createdAt,omitempty"`
	// Status of the execution detail
	Status ExecutionDetailsStatusEnum `json:"status"`
	// Detailed information about the execution
	Detail string `json:"detail"`
	// Whether the execution is a retry or not
	IsRetry bool `json:"isRetry"`
	// Whether the execution is a test or not
	IsTest bool `json:"isTest"`
	// Provider ID of the job
	ProviderID ProvidersIDEnum `json:"providerId"`
	// Raw data of the execution
	Raw *string `json:"raw,omitempty"`
	// Source of the execution detail
	Source ExecutionDetailsSourceEnum `json:"source"`
}

func (o *ActivityNotificationExecutionDetailResponseDto) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ActivityNotificationExecutionDetailResponseDto) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ActivityNotificationExecutionDetailResponseDto) GetStatus() ExecutionDetailsStatusEnum {
	if o == nil {
		return ExecutionDetailsStatusEnum("")
	}
	return o.Status
}

func (o *ActivityNotificationExecutionDetailResponseDto) GetDetail() string {
	if o == nil {
		return ""
	}
	return o.Detail
}

func (o *ActivityNotificationExecutionDetailResponseDto) GetIsRetry() bool {
	if o == nil {
		return false
	}
	return o.IsRetry
}

func (o *ActivityNotificationExecutionDetailResponseDto) GetIsTest() bool {
	if o == nil {
		return false
	}
	return o.IsTest
}

func (o *ActivityNotificationExecutionDetailResponseDto) GetProviderID() ProvidersIDEnum {
	if o == nil {
		return ProvidersIDEnum("")
	}
	return o.ProviderID
}

func (o *ActivityNotificationExecutionDetailResponseDto) GetRaw() *string {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *ActivityNotificationExecutionDetailResponseDto) GetSource() ExecutionDetailsSourceEnum {
	if o == nil {
		return ExecutionDetailsSourceEnum("")
	}
	return o.Source
}
