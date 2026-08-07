# Agents.Integrations

## Overview

### Available Operations

* [Create](#create) - Create an agent integration
* [List](#list) - List agent integrations
* [Update](#update) - Update an agent integration
* [Delete](#delete) - Delete an agent integration

## Create

Create a link between an agent (by identifier) and an integration (by integration **identifier**, not the internal _id).

### Example Usage

<!-- UsageSnippet language="go" operationID="AgentIntegrationsController_addAgentIntegration" method="post" path="/v1/agents/{identifier}/integrations" -->
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

    res, err := s.Agents.Integrations.Create(ctx, "<value>", components.AddAgentIntegrationRequestDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentIntegrationResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `identifier`                                                                                         | `string`                                                                                             | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `addAgentIntegrationRequestDto`                                                                      | [components.AddAgentIntegrationRequestDto](../../models/components/addagentintegrationrequestdto.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `idempotencyKey`                                                                                     | `*string`                                                                                            | :heavy_minus_sign:                                                                                   | A header for idempotency purposes                                                                    |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.AgentIntegrationsControllerAddAgentIntegrationResponse](../../models/operations/agentintegrationscontrolleraddagentintegrationresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## List

Retrieve integration links for an agent identified by its external identifier. Supports cursor pagination via **after**, **before**, **limit**, **orderBy**, and **orderDirection**.

### Example Usage

<!-- UsageSnippet language="go" operationID="AgentIntegrationsController_listAgentIntegrations" method="get" path="/v1/agents/{identifier}/integrations" -->
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

    res, err := s.Agents.Integrations.List(ctx, operations.AgentIntegrationsControllerListAgentIntegrationsRequest{
        Identifier: "<value>",
        Limit: v3.Pointer[float64](10.0),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAgentIntegrationsResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.AgentIntegrationsControllerListAgentIntegrationsRequest](../../models/operations/agentintegrationscontrollerlistagentintegrationsrequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.AgentIntegrationsControllerListAgentIntegrationsResponse](../../models/operations/agentintegrationscontrollerlistagentintegrationsresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Update

Update which integration a link points to (by integration **identifier**, not the internal _id).

### Example Usage

<!-- UsageSnippet language="go" operationID="AgentIntegrationsController_updateAgentIntegration" method="patch" path="/v1/agents/{identifier}/integrations/{agentIntegrationId}" -->
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

    res, err := s.Agents.Integrations.Update(ctx, "<value>", "<id>", components.UpdateAgentIntegrationRequestDto{
        IntegrationIdentifier: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentIntegrationResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `identifier`                                                                                               | `string`                                                                                                   | :heavy_check_mark:                                                                                         | N/A                                                                                                        |
| `agentIntegrationID`                                                                                       | `string`                                                                                                   | :heavy_check_mark:                                                                                         | N/A                                                                                                        |
| `updateAgentIntegrationRequestDto`                                                                         | [components.UpdateAgentIntegrationRequestDto](../../models/components/updateagentintegrationrequestdto.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |
| `idempotencyKey`                                                                                           | `*string`                                                                                                  | :heavy_minus_sign:                                                                                         | A header for idempotency purposes                                                                          |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.AgentIntegrationsControllerUpdateAgentIntegrationResponse](../../models/operations/agentintegrationscontrollerupdateagentintegrationresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Delete

Delete a specific agent-integration link by its document id.

### Example Usage

<!-- UsageSnippet language="go" operationID="AgentIntegrationsController_removeAgentIntegration" method="delete" path="/v1/agents/{identifier}/integrations/{agentIntegrationId}" -->
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

    res, err := s.Agents.Integrations.Delete(ctx, "<value>", "<id>", nil)
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
| `identifier`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `agentIntegrationID`                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AgentIntegrationsControllerRemoveAgentIntegrationResponse](../../models/operations/agentintegrationscontrollerremoveagentintegrationresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |