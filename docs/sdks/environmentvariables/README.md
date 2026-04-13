# EnvironmentVariables

## Overview

### Available Operations

* [List](#list) - List all variables
* [Create](#create) - Create a variable
* [Retrieve](#retrieve) - Get environment variable
* [Update](#update) - Update a variable
* [Delete](#delete) - Delete environment variable
* [Usage](#usage) - Retrieve a variable usage

## List

Returns all environment variables for the current organization. Secret values are masked.

### Example Usage

<!-- UsageSnippet language="go" operationID="EnvironmentVariablesController_listEnvironmentVariables" method="get" path="/v1/environment-variables" -->
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

    res, err := s.EnvironmentVariables.List(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.EnvironmentVariableResponseDtos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `search`                                                 | `*string`                                                | :heavy_minus_sign:                                       | Filter variables by key (case-insensitive partial match) |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.EnvironmentVariablesControllerListEnvironmentVariablesResponse](../../models/operations/environmentvariablescontrollerlistenvironmentvariablesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Create

Creates a new environment variable. Keys must be uppercase with underscores only (e.g. BASE_URL). Secret variables are encrypted at rest and masked in API responses.

### Example Usage

<!-- UsageSnippet language="go" operationID="EnvironmentVariablesController_createEnvironmentVariable" method="post" path="/v1/environment-variables" -->
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

    res, err := s.EnvironmentVariables.Create(ctx, components.CreateEnvironmentVariableRequestDto{
        Key: "<key>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.EnvironmentVariableResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `createEnvironmentVariableRequestDto`                                                                            | [components.CreateEnvironmentVariableRequestDto](../../models/components/createenvironmentvariablerequestdto.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `idempotencyKey`                                                                                                 | `*string`                                                                                                        | :heavy_minus_sign:                                                                                               | A header for idempotency purposes                                                                                |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.EnvironmentVariablesControllerCreateEnvironmentVariableResponse](../../models/operations/environmentvariablescontrollercreateenvironmentvariableresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 404, 405, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Retrieve

Returns a single environment variable by key. Secret values are masked.

### Example Usage

<!-- UsageSnippet language="go" operationID="EnvironmentVariablesController_getEnvironmentVariable" method="get" path="/v1/environment-variables/{variableKey}" -->
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

    res, err := s.EnvironmentVariables.Retrieve(ctx, "BASE_URL", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.EnvironmentVariableResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |                                                            |
| `variableKey`                                              | `string`                                                   | :heavy_check_mark:                                         | The unique key of the environment variable (e.g. BASE_URL) | BASE_URL                                                   |
| `idempotencyKey`                                           | `*string`                                                  | :heavy_minus_sign:                                         | A header for idempotency purposes                          |                                                            |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |                                                            |

### Response

**[*operations.EnvironmentVariablesControllerGetEnvironmentVariableResponse](../../models/operations/environmentvariablescontrollergetenvironmentvariableresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Update

Updates an existing environment variable. Providing values replaces all existing per-environment values.

### Example Usage

<!-- UsageSnippet language="go" operationID="EnvironmentVariablesController_updateEnvironmentVariable" method="patch" path="/v1/environment-variables/{variableKey}" -->
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

    res, err := s.EnvironmentVariables.Update(ctx, "BASE_URL", components.UpdateEnvironmentVariableRequestDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.EnvironmentVariableResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |                                                                                                                  |
| `variableKey`                                                                                                    | `string`                                                                                                         | :heavy_check_mark:                                                                                               | The unique key of the environment variable (e.g. BASE_URL)                                                       | BASE_URL                                                                                                         |
| `updateEnvironmentVariableRequestDto`                                                                            | [components.UpdateEnvironmentVariableRequestDto](../../models/components/updateenvironmentvariablerequestdto.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `idempotencyKey`                                                                                                 | `*string`                                                                                                        | :heavy_minus_sign:                                                                                               | A header for idempotency purposes                                                                                |                                                                                                                  |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |                                                                                                                  |

### Response

**[*operations.EnvironmentVariablesControllerUpdateEnvironmentVariableResponse](../../models/operations/environmentvariablescontrollerupdateenvironmentvariableresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Delete

Deletes an environment variable by key.

### Example Usage

<!-- UsageSnippet language="go" operationID="EnvironmentVariablesController_deleteEnvironmentVariable" method="delete" path="/v1/environment-variables/{variableKey}" -->
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

    res, err := s.EnvironmentVariables.Delete(ctx, "BASE_URL", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |                                                            |
| `variableKey`                                              | `string`                                                   | :heavy_check_mark:                                         | The unique key of the environment variable (e.g. BASE_URL) | BASE_URL                                                   |
| `idempotencyKey`                                           | `*string`                                                  | :heavy_minus_sign:                                         | A header for idempotency purposes                          |                                                            |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |                                                            |

### Response

**[*operations.EnvironmentVariablesControllerDeleteEnvironmentVariableResponse](../../models/operations/environmentvariablescontrollerdeleteenvironmentvariableresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Usage

Returns the workflows that reference this environment variable via `{{env.KEY}}` in their step controls. **variableId** is required.

### Example Usage

<!-- UsageSnippet language="go" operationID="EnvironmentVariablesController_getEnvironmentVariableUsage" method="get" path="/v1/environment-variables/{variableKey}/usage" -->
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

    res, err := s.EnvironmentVariables.Usage(ctx, "BASE_URL", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetEnvironmentVariableUsageResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |                                                            |
| `variableKey`                                              | `string`                                                   | :heavy_check_mark:                                         | The unique key of the environment variable (e.g. BASE_URL) | BASE_URL                                                   |
| `idempotencyKey`                                           | `*string`                                                  | :heavy_minus_sign:                                         | A header for idempotency purposes                          |                                                            |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |                                                            |

### Response

**[*operations.EnvironmentVariablesControllerGetEnvironmentVariableUsageResponse](../../models/operations/environmentvariablescontrollergetenvironmentvariableusageresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |