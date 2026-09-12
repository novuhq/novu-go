# IntegrationResponseDtoKind

Distinguishes delivery integrations from agent-runtime integrations. Defaults to "delivery". Agent integrations do not have a channel.

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.IntegrationResponseDtoKindDelivery
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `IntegrationResponseDtoKindDelivery` | delivery                             |
| `IntegrationResponseDtoKindAgent`    | agent                                |