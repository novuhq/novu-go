# Environments

## Overview

Environments allow you to manage different stages of your application development lifecycle. Each environment has its own set of API keys and configurations, enabling you to separate development, staging, and production workflows.
<https://docs.novu.co/platform/environments>

### Available Operations

* [GetTags](#gettags) - List environment tags
* [Diff](#diff) - Compare resources between environments
* [Publish](#publish) - Publish resources to target environment
* [Create](#create) - Create an environment
* [List](#list) - List all environments
* [Update](#update) - Update an environment
* [Delete](#delete) - Delete an environment

## GetTags

Retrieve all unique tags used in workflows within the specified environment. These tags can be used for filtering workflows.

### Example Usage

<!-- UsageSnippet language="go" operationID="EnvironmentsController_getEnvironmentTags" method="get" path="/v2/environments/{environmentId}/tags" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Environments.GetTags(ctx, "6615943e7ace93b0540ae377", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetEnvironmentTagsDtos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `environmentID`                                          | `string`                                                 | :heavy_check_mark:                                       | Environment internal ID (MongoDB ObjectId) or identifier | 6615943e7ace93b0540ae377                                 |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.EnvironmentsControllerGetEnvironmentTagsResponse](../../models/operations/environmentscontrollergetenvironmenttagsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Diff

Compares workflows and other resources between the source and target environments, returning detailed diff information including additions, modifications, and deletions.

### Example Usage

<!-- UsageSnippet language="go" operationID="EnvironmentsController_diffEnvironment" method="post" path="/v2/environments/{targetEnvironmentId}/diff" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"github.com/novuhq/novu-go/v3/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Environments.Diff(ctx, "6615943e7ace93b0540ae377", components.DiffEnvironmentRequestDto{
        SourceEnvironmentID: v3.Pointer("507f1f77bcf86cd799439011"),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DiffEnvironmentResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `targetEnvironmentID`                                                                        | `string`                                                                                     | :heavy_check_mark:                                                                           | Target environment ID (MongoDB ObjectId) to compare against                                  | 6615943e7ace93b0540ae377                                                                     |
| `diffEnvironmentRequestDto`                                                                  | [components.DiffEnvironmentRequestDto](../../models/components/diffenvironmentrequestdto.md) | :heavy_check_mark:                                                                           | Diff request configuration                                                                   |                                                                                              |
| `idempotencyKey`                                                                             | `*string`                                                                                    | :heavy_minus_sign:                                                                           | A header for idempotency purposes                                                            |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.EnvironmentsControllerDiffEnvironmentResponse](../../models/operations/environmentscontrollerdiffenvironmentresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Publish

Publishes all workflows and resources from the source environment to the target environment. Optionally specify specific resources to publish or use dryRun mode to preview changes.

### Example Usage

<!-- UsageSnippet language="go" operationID="EnvironmentsController_publishEnvironment" method="post" path="/v2/environments/{targetEnvironmentId}/publish" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"github.com/novuhq/novu-go/v3/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Environments.Publish(ctx, "6615943e7ace93b0540ae377", components.PublishEnvironmentRequestDto{
        SourceEnvironmentID: v3.Pointer("507f1f77bcf86cd799439011"),
        Resources: []components.ResourceToPublishDto{
            components.ResourceToPublishDto{
                ResourceType: components.ResourceTypeEnumRegular,
                ResourceID: "workflow-id-1",
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PublishEnvironmentResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `targetEnvironmentID`                                                                              | `string`                                                                                           | :heavy_check_mark:                                                                                 | Target environment ID (MongoDB ObjectId) to publish resources to                                   | 6615943e7ace93b0540ae377                                                                           |
| `publishEnvironmentRequestDto`                                                                     | [components.PublishEnvironmentRequestDto](../../models/components/publishenvironmentrequestdto.md) | :heavy_check_mark:                                                                                 | Publish request configuration                                                                      |                                                                                                    |
| `idempotencyKey`                                                                                   | `*string`                                                                                          | :heavy_minus_sign:                                                                                 | A header for idempotency purposes                                                                  |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.EnvironmentsControllerPublishEnvironmentResponse](../../models/operations/environmentscontrollerpublishenvironmentresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Create

Creates a new environment within the current organization. 
    Environments allow you to manage different stages of your application development lifecycle.
    Each environment has its own set of API keys and configurations.

### Example Usage

<!-- UsageSnippet language="go" operationID="EnvironmentsControllerV1_createEnvironment" method="post" path="/v1/environments" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"github.com/novuhq/novu-go/v3/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Environments.Create(ctx, components.CreateEnvironmentRequestDto{
        Name: "Production Environment",
        ParentID: v3.Pointer("60d5ecb8b3b3a30015f3e1a1"),
        Color: "#3498db",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.EnvironmentResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `createEnvironmentRequestDto`                                                                    | [components.CreateEnvironmentRequestDto](../../models/components/createenvironmentrequestdto.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `idempotencyKey`                                                                                 | `*string`                                                                                        | :heavy_minus_sign:                                                                               | A header for idempotency purposes                                                                |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.EnvironmentsControllerV1CreateEnvironmentResponse](../../models/operations/environmentscontrollerv1createenvironmentresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 402, 414                               | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## List

This API returns a list of environments for the current organization. 
    Each environment contains its configuration, API keys (if user has access), and metadata.

### Example Usage

<!-- UsageSnippet language="go" operationID="EnvironmentsControllerV1_listMyEnvironments" method="get" path="/v1/environments" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Environments.List(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.EnvironmentResponseDtos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.EnvironmentsControllerV1ListMyEnvironmentsResponse](../../models/operations/environmentscontrollerv1listmyenvironmentsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Update

Update an environment by its unique identifier **environmentId**. 
    You can modify the environment name, identifier, color, and other configuration settings.

### Example Usage

<!-- UsageSnippet language="go" operationID="EnvironmentsControllerV1_updateMyEnvironment" method="put" path="/v1/environments/{environmentId}" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"github.com/novuhq/novu-go/v3/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Environments.Update(ctx, "<id>", components.UpdateEnvironmentRequestDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.EnvironmentResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `environmentID`                                                                                  | `string`                                                                                         | :heavy_check_mark:                                                                               | The unique identifier of the environment                                                         |
| `updateEnvironmentRequestDto`                                                                    | [components.UpdateEnvironmentRequestDto](../../models/components/updateenvironmentrequestdto.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `idempotencyKey`                                                                                 | `*string`                                                                                        | :heavy_minus_sign:                                                                               | A header for idempotency purposes                                                                |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.EnvironmentsControllerV1UpdateMyEnvironmentResponse](../../models/operations/environmentscontrollerv1updatemyenvironmentresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Delete

Delete an environment by its unique identifier **environmentId**. 
    This action is irreversible and will remove the environment and all its associated data.

### Example Usage

<!-- UsageSnippet language="go" operationID="EnvironmentsControllerV1_deleteEnvironment" method="delete" path="/v1/environments/{environmentId}" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Environments.Delete(ctx, "<id>", nil)
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
| `environmentID`                                          | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the environment                 |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.EnvironmentsControllerV1DeleteEnvironmentResponse](../../models/operations/environmentscontrollerv1deleteenvironmentresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |