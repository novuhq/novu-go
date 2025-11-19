# Groups
(*Translations.Groups*)

## Overview

### Available Operations

* [Delete](#delete) - Delete a translation group
* [Retrieve](#retrieve) - Retrieve a translation group

## Delete

Delete an entire translation group and all its translations

### Example Usage

<!-- UsageSnippet language="go" operationID="TranslationController_deleteTranslationGroupEndpoint" method="delete" path="/v2/translations/{resourceType}/{resourceId}" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"github.com/novuhq/novu-go/v3/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Translations.Groups.Delete(ctx, operations.ResourceTypeWorkflow, "welcome-email", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |                                                                    |
| `resourceType`                                                     | [operations.ResourceType](../../models/operations/resourcetype.md) | :heavy_check_mark:                                                 | Resource type                                                      | workflow                                                           |
| `resourceID`                                                       | *string*                                                           | :heavy_check_mark:                                                 | Resource ID                                                        | welcome-email                                                      |
| `idempotencyKey`                                                   | **string*                                                          | :heavy_minus_sign:                                                 | A header for idempotency purposes                                  |                                                                    |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |                                                                    |

### Response

**[*operations.TranslationControllerDeleteTranslationGroupEndpointResponse](../../models/operations/translationcontrollerdeletetranslationgroupendpointresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Retrieve

Retrieves a single translation group by resource type (workflow, layout) and resource ID (workflowId, layoutId)

### Example Usage

<!-- UsageSnippet language="go" operationID="TranslationController_getTranslationGroupEndpoint" method="get" path="/v2/translations/group/{resourceType}/{resourceId}" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"github.com/novuhq/novu-go/v3/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Translations.Groups.Retrieve(ctx, operations.TranslationControllerGetTranslationGroupEndpointPathParamResourceTypeWorkflow, "welcome-email", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TranslationGroupDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                            | Type                                                                                                                                                                                 | Required                                                                                                                                                                             | Description                                                                                                                                                                          | Example                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                | :heavy_check_mark:                                                                                                                                                                   | The context to use for the request.                                                                                                                                                  |                                                                                                                                                                                      |
| `resourceType`                                                                                                                                                                       | [operations.TranslationControllerGetTranslationGroupEndpointPathParamResourceType](../../models/operations/translationcontrollergettranslationgroupendpointpathparamresourcetype.md) | :heavy_check_mark:                                                                                                                                                                   | Resource type                                                                                                                                                                        | workflow                                                                                                                                                                             |
| `resourceID`                                                                                                                                                                         | *string*                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                   | Resource ID                                                                                                                                                                          | welcome-email                                                                                                                                                                        |
| `idempotencyKey`                                                                                                                                                                     | **string*                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                   | A header for idempotency purposes                                                                                                                                                    |                                                                                                                                                                                      |
| `opts`                                                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                   | The options for this request.                                                                                                                                                        |                                                                                                                                                                                      |

### Response

**[*operations.TranslationControllerGetTranslationGroupEndpointResponse](../../models/operations/translationcontrollergettranslationgroupendpointresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |