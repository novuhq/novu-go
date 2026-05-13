# Kind

Distinguishes delivery integrations from agent-runtime integrations. Defaults to "delivery". Agent integrations do not have a channel.

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.KindDelivery
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `KindDelivery` | delivery       |
| `KindAgent`    | agent          |