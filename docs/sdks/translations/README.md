# Translations
(*Translations*)

## Overview

Used to localize your notifications to different languages.
<https://docs.novu.co/platform/workflow/translations>

### Available Operations

* [Create](#create) - Create a translation
* [Retrieve](#retrieve) - Retrieve a translation
* [Delete](#delete) - Delete a translation
* [Upload](#upload) - Upload translation files

## Create

Create a translation for a specific workflow and locale, if the translation already exists, it will be updated

### Example Usage

<!-- UsageSnippet language="go" operationID="TranslationController_createTranslationEndpoint" method="post" path="/v2/translations" -->
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

    res, err := s.Translations.Create(ctx, components.CreateTranslationRequestDto{
        ResourceID: "welcome-email",
        ResourceType: components.ResourceTypeWorkflow,
        Locale: "en_US",
        Content: components.Content{},
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TranslationResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `createTranslationRequestDto`                                                                    | [components.CreateTranslationRequestDto](../../models/components/createtranslationrequestdto.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `idempotencyKey`                                                                                 | **string*                                                                                        | :heavy_minus_sign:                                                                               | A header for idempotency purposes                                                                |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.TranslationControllerCreateTranslationEndpointResponse](../../models/operations/translationcontrollercreatetranslationendpointresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Retrieve

Retrieve a specific translation by resource type, resource ID and locale

### Example Usage

<!-- UsageSnippet language="go" operationID="TranslationController_getSingleTranslation" method="get" path="/v2/translations/{resourceType}/{resourceId}/{locale}" -->
```go
package main

import(
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := novugo.New(
        novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Translations.Retrieve(ctx, operations.PathParamResourceTypeWorkflow, "welcome-email", "en_US", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TranslationResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `resourceType`                                                                       | [operations.PathParamResourceType](../../models/operations/pathparamresourcetype.md) | :heavy_check_mark:                                                                   | Resource type                                                                        |                                                                                      |
| `resourceID`                                                                         | *string*                                                                             | :heavy_check_mark:                                                                   | Resource ID                                                                          | welcome-email                                                                        |
| `locale`                                                                             | *string*                                                                             | :heavy_check_mark:                                                                   | Locale code                                                                          | en_US                                                                                |
| `idempotencyKey`                                                                     | **string*                                                                            | :heavy_minus_sign:                                                                   | A header for idempotency purposes                                                    |                                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.TranslationControllerGetSingleTranslationResponse](../../models/operations/translationcontrollergetsingletranslationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a specific translation by resource type, resource ID and locale

### Example Usage

<!-- UsageSnippet language="go" operationID="TranslationController_deleteTranslationEndpoint" method="delete" path="/v2/translations/{resourceType}/{resourceId}/{locale}" -->
```go
package main

import(
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := novugo.New(
        novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Translations.Delete(ctx, operations.TranslationControllerDeleteTranslationEndpointPathParamResourceTypeWorkflow, "<id>", "el", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                            | :heavy_check_mark:                                                                                                                                                               | The context to use for the request.                                                                                                                                              |
| `resourceType`                                                                                                                                                                   | [operations.TranslationControllerDeleteTranslationEndpointPathParamResourceType](../../models/operations/translationcontrollerdeletetranslationendpointpathparamresourcetype.md) | :heavy_check_mark:                                                                                                                                                               | Resource type                                                                                                                                                                    |
| `resourceID`                                                                                                                                                                     | *string*                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                               | Resource ID                                                                                                                                                                      |
| `locale`                                                                                                                                                                         | *string*                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                               | Locale code                                                                                                                                                                      |
| `idempotencyKey`                                                                                                                                                                 | **string*                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                               | A header for idempotency purposes                                                                                                                                                |
| `opts`                                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                                         | :heavy_minus_sign:                                                                                                                                                               | The options for this request.                                                                                                                                                    |

### Response

**[*operations.TranslationControllerDeleteTranslationEndpointResponse](../../models/operations/translationcontrollerdeletetranslationendpointresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Upload

Upload one or more JSON translation files for a specific workflow. Files name must match the locale, e.g. en_US.json

### Example Usage

<!-- UsageSnippet language="go" operationID="TranslationController_uploadTranslationFiles" method="post" path="/v2/translations/upload" -->
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

    res, err := s.Translations.Upload(ctx, components.UploadTranslationsRequestDto{
        ResourceID: "welcome-email",
        ResourceType: components.UploadTranslationsRequestDtoResourceTypeWorkflow,
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UploadTranslationsResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `uploadTranslationsRequestDto`                                                                     | [components.UploadTranslationsRequestDto](../../models/components/uploadtranslationsrequestdto.md) | :heavy_check_mark:                                                                                 | Translation files upload body details                                                              |
| `idempotencyKey`                                                                                   | **string*                                                                                          | :heavy_minus_sign:                                                                                 | A header for idempotency purposes                                                                  |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.TranslationControllerUploadTranslationFilesResponse](../../models/operations/translationcontrolleruploadtranslationfilesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |