# StepIssueSeverityEnum

Blocking severity of the issue. `error` (default when omitted) blocks save; `warning` is a non-blocking notice.

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.StepIssueSeverityEnumError
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `StepIssueSeverityEnumError`   | error                          |
| `StepIssueSeverityEnumWarning` | warning                        |