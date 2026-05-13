# GenerateChatOauthURLRequestDtoConnectionMode

Connection mode that determines how the channel connection is scoped. Use "subscriber" (default) to associate the connection with a specific subscriber. Use "shared" to associate the connection with a context instead of a subscriber — subscriberId will not be stored on the connection.

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.GenerateChatOauthURLRequestDtoConnectionModeSubscriber
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `GenerateChatOauthURLRequestDtoConnectionModeSubscriber` | subscriber                                               |
| `GenerateChatOauthURLRequestDtoConnectionModeShared`     | shared                                                   |