# Master
(*Translations.Master*)

## Overview

### Available Operations

* [Retrieve](#retrieve) - Retrieve master translations JSON
* [Import](#import) - Import master translations JSON
* [Upload](#upload) - Upload master translations JSON file

## Retrieve

Retrieve all translations for a locale in master JSON format organized by resourceId (workflowId)

### Example Usage

<!-- UsageSnippet language="go" operationID="TranslationController_getMasterJsonEndpoint" method="get" path="/v2/translations/master-json" -->
```go
package main

import(
	"context"
	novugo "github.com/novuhq/novu-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := novugo.New(
        novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Translations.Master.Retrieve(ctx, novugo.Pointer("en_US"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetMasterJSONResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |                                                                        |
| `locale`                                                               | **string*                                                              | :heavy_minus_sign:                                                     | Locale to export. If not provided, exports organization default locale | en_US                                                                  |
| `idempotencyKey`                                                       | **string*                                                              | :heavy_minus_sign:                                                     | A header for idempotency purposes                                      |                                                                        |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |                                                                        |

### Response

**[*operations.TranslationControllerGetMasterJSONEndpointResponse](../../models/operations/translationcontrollergetmasterjsonendpointresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Import

Import translations for multiple workflows from master JSON format for a specific locale

### Example Usage

<!-- UsageSnippet language="go" operationID="TranslationController_importMasterJsonEndpoint" method="post" path="/v2/translations/master-json" -->
```go
package main

import(
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := novugo.New(
        novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Translations.Master.Import(ctx, components.ImportMasterJSONRequestDto{
        Locale: "en_US",
        MasterJSON: components.MasterJSON{},
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ImportMasterJSONResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `importMasterJSONRequestDto`                                                                   | [components.ImportMasterJSONRequestDto](../../models/components/importmasterjsonrequestdto.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `idempotencyKey`                                                                               | **string*                                                                                      | :heavy_minus_sign:                                                                             | A header for idempotency purposes                                                              |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.TranslationControllerImportMasterJSONEndpointResponse](../../models/operations/translationcontrollerimportmasterjsonendpointresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Upload

Upload a master JSON file containing translations for multiple workflows. Locale is automatically detected from filename (e.g., en_US.json)

### Example Usage

<!-- UsageSnippet language="go" operationID="TranslationController_uploadMasterJsonEndpoint" method="post" path="/v2/translations/master-json/upload" -->
```go
package main

import(
	"context"
	novugo "github.com/novuhq/novu-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := novugo.New(
        novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Translations.Master.Upload(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ImportMasterJSONResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TranslationControllerUploadMasterJSONEndpointResponse](../../models/operations/translationcontrolleruploadmasterjsonendpointresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |