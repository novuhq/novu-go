# Workflows
(*Workflows*)

## Overview

All notifications are sent via a workflow. Each workflow acts as a container for the logic and blueprint that are associated with a type of notification in your system.
<https://docs.novu.co/workflows>

### Available Operations

* [Create](#create) - Create a workflow
* [List](#list) - List all workflows
* [Update](#update) - Update a workflow
* [Get](#get) - Retrieve a workflow
* [Delete](#delete) - Delete a workflow
* [Patch](#patch) - Update a workflow
* [Sync](#sync) - Sync a workflow

## Create

Creates a new workflow in the Novu Cloud environment

### Example Usage

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

    res, err := s.Workflows.Create(ctx, components.CreateWorkflowDto{
        Name: "<value>",
        WorkflowID: "<id>",
        Steps: []components.Steps{},
        Preferences: &components.PreferencesRequestDto{
            User: novugo.Pointer(components.CreateUserUserWorkflowPreferencesDto(
                components.UserWorkflowPreferencesDto{
                    All: components.CreateUserAllWorkflowPreferenceDto(
                        components.WorkflowPreferenceDto{},
                    ),
                    Channels: map[string]components.ChannelPreferenceDto{
                        "email": components.ChannelPreferenceDto{},
                        "sms": components.ChannelPreferenceDto{
                            Enabled: novugo.Bool(false),
                        },
                    },
                },
            )),
            Workflow: &components.Workflow{
                All: components.CreatePreferencesRequestDtoAllWorkflowPreferenceDto(
                    components.WorkflowPreferenceDto{},
                ),
                Channels: map[string]components.ChannelPreferenceDto{
                    "email": components.ChannelPreferenceDto{},
                    "sms": components.ChannelPreferenceDto{
                        Enabled: novugo.Bool(false),
                    },
                },
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkflowResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `createWorkflowDto`                                                          | [components.CreateWorkflowDto](../../models/components/createworkflowdto.md) | :heavy_check_mark:                                                           | Workflow creation details                                                    |
| `idempotencyKey`                                                             | **string*                                                                    | :heavy_minus_sign:                                                           | A header for idempotency purposes                                            |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.WorkflowControllerCreateResponse](../../models/operations/workflowcontrollercreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## List

Retrieves a list of workflows with optional filtering and pagination

### Example Usage

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

    res, err := s.Workflows.List(ctx, operations.WorkflowControllerSearchWorkflowsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListWorkflowResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.WorkflowControllerSearchWorkflowsRequest](../../models/operations/workflowcontrollersearchworkflowsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.WorkflowControllerSearchWorkflowsResponse](../../models/operations/workflowcontrollersearchworkflowsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Update

Updates the details of an existing workflow, here **workflowId** is the identifier of the workflow

### Example Usage

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

    res, err := s.Workflows.Update(ctx, "<id>", components.UpdateWorkflowDto{
        Name: "<value>",
        Steps: []components.UpdateWorkflowDtoSteps{},
        Preferences: components.PreferencesRequestDto{
            User: novugo.Pointer(components.CreateUserUserWorkflowPreferencesDto(
                components.UserWorkflowPreferencesDto{
                    All: components.CreateUserAllWorkflowPreferenceDto(
                        components.WorkflowPreferenceDto{},
                    ),
                    Channels: map[string]components.ChannelPreferenceDto{
                        "email": components.ChannelPreferenceDto{},
                        "sms": components.ChannelPreferenceDto{
                            Enabled: novugo.Bool(false),
                        },
                    },
                },
            )),
            Workflow: &components.Workflow{
                All: components.CreatePreferencesRequestDtoAllWorkflowPreferenceDto(
                    components.WorkflowPreferenceDto{},
                ),
                Channels: map[string]components.ChannelPreferenceDto{
                    "email": components.ChannelPreferenceDto{},
                    "sms": components.ChannelPreferenceDto{
                        Enabled: novugo.Bool(false),
                    },
                },
            },
        },
        Origin: components.WorkflowOriginEnumExternal,
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkflowResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `workflowID`                                                                 | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `updateWorkflowDto`                                                          | [components.UpdateWorkflowDto](../../models/components/updateworkflowdto.md) | :heavy_check_mark:                                                           | Workflow update details                                                      |
| `idempotencyKey`                                                             | **string*                                                                    | :heavy_minus_sign:                                                           | A header for idempotency purposes                                            |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.WorkflowControllerUpdateResponse](../../models/operations/workflowcontrollerupdateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Get

Fetches details of a specific workflow by its unique identifier **workflowId**

### Example Usage

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

    res, err := s.Workflows.Get(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkflowResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `workflowID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `environmentID`                                          | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WorkflowControllerGetWorkflowResponse](../../models/operations/workflowcontrollergetworkflowresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Delete

Removes a specific workflow by its unique identifier **workflowId**

### Example Usage

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

    res, err := s.Workflows.Delete(ctx, "<id>", nil)
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
| `workflowID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WorkflowControllerRemoveWorkflowResponse](../../models/operations/workflowcontrollerremoveworkflowresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Patch

Partially updates a workflow by its unique identifier **workflowId**

### Example Usage

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

    res, err := s.Workflows.Patch(ctx, "<id>", components.PatchWorkflowDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkflowResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `workflowID`                                                               | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        |
| `patchWorkflowDto`                                                         | [components.PatchWorkflowDto](../../models/components/patchworkflowdto.md) | :heavy_check_mark:                                                         | Workflow patch details                                                     |
| `idempotencyKey`                                                           | **string*                                                                  | :heavy_minus_sign:                                                         | A header for idempotency purposes                                          |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.WorkflowControllerPatchWorkflowResponse](../../models/operations/workflowcontrollerpatchworkflowresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Sync

Synchronizes a workflow to the target environment

### Example Usage

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

    res, err := s.Workflows.Sync(ctx, "<id>", components.SyncWorkflowDto{
        TargetEnvironmentID: "<id>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkflowResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `workflowID`                                                             | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `syncWorkflowDto`                                                        | [components.SyncWorkflowDto](../../models/components/syncworkflowdto.md) | :heavy_check_mark:                                                       | Sync workflow details                                                    |
| `idempotencyKey`                                                         | **string*                                                                | :heavy_minus_sign:                                                       | A header for idempotency purposes                                        |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.WorkflowControllerSyncResponse](../../models/operations/workflowcontrollersyncresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |