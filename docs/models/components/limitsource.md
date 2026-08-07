# LimitSource

Which constraint produced the limits. `plan` limits are lifted by upgrading; `system` limits (platform cap or per-organization override) require contacting the Novu team.

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.LimitSourcePlan
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `LimitSourcePlan`   | plan                |
| `LimitSourceSystem` | system              |