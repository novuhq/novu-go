# ConnectionMode

Scope results relative to the subscriber. `subscriber` returns only the subscriber-owned connections, `shared` returns only shared (workspace-level) connections. Omit to return both.

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/operations"
)

value := operations.ConnectionModeSubscriber
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ConnectionModeSubscriber` | subscriber                 |
| `ConnectionModeShared`     | shared                     |