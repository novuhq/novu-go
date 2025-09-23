# Layouts
(*Layouts*)

## Overview

Layouts are reusable wrappers for your email notifications.
<https://docs.novu.co/platform/workflow/layouts>

### Available Operations

* [Create](#create) - Create a layout
* [List](#list) - List all layouts
* [Update](#update) - Update a layout
* [Retrieve](#retrieve) - Retrieve a layout
* [Delete](#delete) - Delete a layout
* [Duplicate](#duplicate) - Duplicate a layout
* [GeneratePreview](#generatepreview) - Generate layout preview
* [Usage](#usage) - Get layout usage

## Create

Creates a new layout in the Novu Cloud environment

### Example Usage

<!-- UsageSnippet language="go" operationID="LayoutsController_create" method="post" path="/v2/layouts" -->
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

    res, err := s.Layouts.Create(ctx, components.CreateLayoutDto{
        LayoutID: "<id>",
        Name: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.LayoutResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `createLayoutDto`                                                        | [components.CreateLayoutDto](../../models/components/createlayoutdto.md) | :heavy_check_mark:                                                       | Layout creation details                                                  |
| `idempotencyKey`                                                         | **string*                                                                | :heavy_minus_sign:                                                       | A header for idempotency purposes                                        |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.LayoutsControllerCreateResponse](../../models/operations/layoutscontrollercreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## List

Retrieves a list of layouts with optional filtering and pagination

### Example Usage

<!-- UsageSnippet language="go" operationID="LayoutsController_list" method="get" path="/v2/layouts" -->
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

    res, err := s.Layouts.List(ctx, operations.LayoutsControllerListRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListLayoutResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.LayoutsControllerListRequest](../../models/operations/layoutscontrollerlistrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.LayoutsControllerListResponse](../../models/operations/layoutscontrollerlistresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Update

Updates the details of an existing layout, here **layoutId** is the identifier of the layout

### Example Usage

<!-- UsageSnippet language="go" operationID="LayoutsController_update" method="put" path="/v2/layouts/{layoutId}" -->
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

    res, err := s.Layouts.Update(ctx, "<id>", components.UpdateLayoutDto{
        Name: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.LayoutResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `layoutID`                                                               | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `updateLayoutDto`                                                        | [components.UpdateLayoutDto](../../models/components/updatelayoutdto.md) | :heavy_check_mark:                                                       | Layout update details                                                    |
| `idempotencyKey`                                                         | **string*                                                                | :heavy_minus_sign:                                                       | A header for idempotency purposes                                        |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.LayoutsControllerUpdateResponse](../../models/operations/layoutscontrollerupdateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Retrieve

Fetches details of a specific layout by its unique identifier **layoutId**

### Example Usage

<!-- UsageSnippet language="go" operationID="LayoutsController_get" method="get" path="/v2/layouts/{layoutId}" -->
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

    res, err := s.Layouts.Retrieve(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.LayoutResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `layoutID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.LayoutsControllerGetResponse](../../models/operations/layoutscontrollergetresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Delete

Removes a specific layout by its unique identifier **layoutId**

### Example Usage

<!-- UsageSnippet language="go" operationID="LayoutsController__delete" method="delete" path="/v2/layouts/{layoutId}" -->
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

    res, err := s.Layouts.Delete(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `layoutID`                                               | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the layout                      |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.LayoutsControllerDeleteResponse](../../models/operations/layoutscontrollerdeleteresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Duplicate

Duplicates a layout by its unique identifier **layoutId**. This will create a new layout with the content of the original layout.

### Example Usage

<!-- UsageSnippet language="go" operationID="LayoutsController_duplicate" method="post" path="/v2/layouts/{layoutId}/duplicate" -->
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

    res, err := s.Layouts.Duplicate(ctx, "<id>", components.DuplicateLayoutDto{
        Name: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.LayoutResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `layoutID`                                                                     | *string*                                                                       | :heavy_check_mark:                                                             | N/A                                                                            |
| `duplicateLayoutDto`                                                           | [components.DuplicateLayoutDto](../../models/components/duplicatelayoutdto.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `idempotencyKey`                                                               | **string*                                                                      | :heavy_minus_sign:                                                             | A header for idempotency purposes                                              |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.LayoutsControllerDuplicateResponse](../../models/operations/layoutscontrollerduplicateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## GeneratePreview

Generates a preview for a layout by its unique identifier **layoutId**

### Example Usage

<!-- UsageSnippet language="go" operationID="LayoutsController_generatePreview" method="post" path="/v2/layouts/{layoutId}/preview" -->
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

    res, err := s.Layouts.GeneratePreview(ctx, "<id>", components.LayoutPreviewRequestDto{
        PreviewPayload: &components.LayoutPreviewPayloadDto{
            Subscriber: &components.SubscriberResponseDtoOptional{
                Channels: []components.ChannelSettingsDto{
                    components.ChannelSettingsDto{
                        ProviderID: components.ChatOrPushProviderEnumMattermost,
                        Credentials: components.ChannelCredentials{
                            WebhookURL: novugo.Pointer("https://example.com/webhook"),
                            Channel: novugo.Pointer("general"),
                            DeviceTokens: []string{
                                "token1",
                                "token2",
                                "token3",
                            },
                            AlertUID: novugo.Pointer("12345-abcde"),
                            Title: novugo.Pointer("Critical Alert"),
                            ImageURL: novugo.Pointer("https://example.com/image.png"),
                            State: novugo.Pointer("resolved"),
                            ExternalURL: novugo.Pointer("https://example.com/details"),
                        },
                        IntegrationID: "<id>",
                    },
                },
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GenerateLayoutPreviewResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `layoutID`                                                                               | *string*                                                                                 | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `layoutPreviewRequestDto`                                                                | [components.LayoutPreviewRequestDto](../../models/components/layoutpreviewrequestdto.md) | :heavy_check_mark:                                                                       | Layout preview generation details                                                        |
| `idempotencyKey`                                                                         | **string*                                                                                | :heavy_minus_sign:                                                                       | A header for idempotency purposes                                                        |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.LayoutsControllerGeneratePreviewResponse](../../models/operations/layoutscontrollergeneratepreviewresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Usage

Retrieves information about workflows that use the specified layout by its unique identifier **layoutId**

### Example Usage

<!-- UsageSnippet language="go" operationID="LayoutsController_getUsage" method="get" path="/v2/layouts/{layoutId}/usage" -->
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

    res, err := s.Layouts.Usage(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetLayoutUsageResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `layoutID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.LayoutsControllerGetUsageResponse](../../models/operations/layoutscontrollergetusageresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |