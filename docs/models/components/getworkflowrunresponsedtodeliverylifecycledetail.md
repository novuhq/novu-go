# GetWorkflowRunResponseDtoDeliveryLifecycleDetail

Why the run reached its delivery lifecycle status, e.g. preference or missing_email. Absent when no reason was recorded.

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.GetWorkflowRunResponseDtoDeliveryLifecycleDetailStepCondition
```


## Values

| Name                                                                             | Value                                                                            |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `GetWorkflowRunResponseDtoDeliveryLifecycleDetailStepCondition`                  | step_condition                                                                   |
| `GetWorkflowRunResponseDtoDeliveryLifecycleDetailPreference`                     | preference                                                                       |
| `GetWorkflowRunResponseDtoDeliveryLifecycleDetailMissingPhone`                   | missing_phone                                                                    |
| `GetWorkflowRunResponseDtoDeliveryLifecycleDetailMissingEmail`                   | missing_email                                                                    |
| `GetWorkflowRunResponseDtoDeliveryLifecycleDetailMissingPushToken`               | missing_push_token                                                               |
| `GetWorkflowRunResponseDtoDeliveryLifecycleDetailMissingWebhookURL`              | missing_webhook_url                                                              |
| `GetWorkflowRunResponseDtoDeliveryLifecycleDetailSomeChannelsMissingCredentials` | some_channels_missing_credentials                                                |
| `GetWorkflowRunResponseDtoDeliveryLifecycleDetailWorkflowMissingChannelStep`     | workflow_missing_channel_step                                                    |
| `GetWorkflowRunResponseDtoDeliveryLifecycleDetailUnknownError`                   | unknown_error                                                                    |
| `GetWorkflowRunResponseDtoDeliveryLifecycleDetailExecutionStopped`               | execution_stopped                                                                |
| `GetWorkflowRunResponseDtoDeliveryLifecycleDetailExecutionCanceledByUser`        | execution_canceled_by_user                                                       |