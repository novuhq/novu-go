# DeliveryLifecycleDetail

Why the run reached its delivery lifecycle status, e.g. preference or missing_email. Absent when no reason was recorded.

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.DeliveryLifecycleDetailStepCondition
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `DeliveryLifecycleDetailStepCondition`                  | step_condition                                          |
| `DeliveryLifecycleDetailPreference`                     | preference                                              |
| `DeliveryLifecycleDetailMissingPhone`                   | missing_phone                                           |
| `DeliveryLifecycleDetailMissingEmail`                   | missing_email                                           |
| `DeliveryLifecycleDetailMissingPushToken`               | missing_push_token                                      |
| `DeliveryLifecycleDetailMissingWebhookURL`              | missing_webhook_url                                     |
| `DeliveryLifecycleDetailSomeChannelsMissingCredentials` | some_channels_missing_credentials                       |
| `DeliveryLifecycleDetailWorkflowMissingChannelStep`     | workflow_missing_channel_step                           |
| `DeliveryLifecycleDetailUnknownError`                   | unknown_error                                           |
| `DeliveryLifecycleDetailExecutionStopped`               | execution_stopped                                       |
| `DeliveryLifecycleDetailExecutionCanceledByUser`        | execution_canceled_by_user                              |