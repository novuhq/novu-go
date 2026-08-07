# Method

Optional HTTP method override for this webhook. Defaults to the integration-level method.

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.MethodPost
```


## Values

| Name          | Value         |
| ------------- | ------------- |
| `MethodPost`  | POST          |
| `MethodPut`   | PUT           |
| `MethodPatch` | PATCH         |