# CreateIntegrationRequestDtoKind

Distinguishes delivery integrations from agent-runtime integrations. Defaults to "delivery". Agent integrations do not require a channel.

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.CreateIntegrationRequestDtoKindDelivery
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `CreateIntegrationRequestDtoKindDelivery` | delivery                                  |
| `CreateIntegrationRequestDtoKindAgent`    | agent                                     |