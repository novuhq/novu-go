# HmacSecretKeyEncoding

Email webhook: how `secretKey` is interpreted when signing webhook calls. `text` signs with the raw UTF-8 bytes; `base64`/`hex` decode it to binary first (e.g. for AWS KMS).

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.HmacSecretKeyEncodingText
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `HmacSecretKeyEncodingText`   | text                          |
| `HmacSecretKeyEncodingBase64` | base64                        |
| `HmacSecretKeyEncodingHex`    | hex                           |