# Mode

OAuth flow mode. Use "connect" (default) to create a workspace channel connection, or "link_user" to identify the subscriber's Slack user ID without creating a connection.

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.ModeConnect
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `ModeConnect`  | connect        |
| `ModeLinkUser` | link_user      |