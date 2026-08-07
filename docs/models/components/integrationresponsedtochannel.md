# IntegrationResponseDtoChannel

The channel type for the integration, which defines how it communicates (e.g., email, SMS). Not set for agent-kind integrations.

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.IntegrationResponseDtoChannelInApp
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `IntegrationResponseDtoChannelInApp` | in_app                               |
| `IntegrationResponseDtoChannelEmail` | email                                |
| `IntegrationResponseDtoChannelSms`   | sms                                  |
| `IntegrationResponseDtoChannelChat`  | chat                                 |
| `IntegrationResponseDtoChannelPush`  | push                                 |
| `IntegrationResponseDtoChannelTool`  | tool                                 |