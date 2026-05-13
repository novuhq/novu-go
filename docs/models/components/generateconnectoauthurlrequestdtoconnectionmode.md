# GenerateConnectOauthURLRequestDtoConnectionMode

Connection mode that determines how the channel connection is scoped. "subscriber" (default) associates the connection with a specific subscriber. "shared" associates the connection with a context instead of a subscriber.

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.GenerateConnectOauthURLRequestDtoConnectionModeSubscriber
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `GenerateConnectOauthURLRequestDtoConnectionModeSubscriber` | subscriber                                                  |
| `GenerateConnectOauthURLRequestDtoConnectionModeShared`     | shared                                                      |