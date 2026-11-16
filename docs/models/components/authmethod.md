# AuthMethod

Twilio: `apiKey` authenticates with `apiKey` (API Key SID) and `secretKey` (API Key Secret) on behalf of `accountSid`. When omitted, `token` (Auth Token) is used.

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.AuthMethodAPIKey
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `AuthMethodAPIKey` | apiKey             |