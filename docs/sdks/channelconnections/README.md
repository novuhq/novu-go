# ChannelConnections

## Overview

### Available Operations

* [List](#list) - List all channel connections
* [Create](#create) - Create a channel connection
* [Retrieve](#retrieve) - Retrieve a channel connection
* [Update](#update) - Update a channel connection
* [Delete](#delete) - Delete a channel connection

## List

List all channel connections for a resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="ChannelConnectionsController_listChannelConnections" method="get" path="/v1/channel-connections" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"github.com/novuhq/novu-go/v3/models/operations"
	"github.com/novuhq/novu-go/v3/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.ChannelConnections.List(ctx, operations.ChannelConnectionsControllerListChannelConnectionsRequest{
        Limit: v3.Pointer[float64](10.0),
        SubscriberID: v3.Pointer("subscriber-123"),
        Channel: operations.ChannelChat.ToPointer(),
        ProviderID: components.ProvidersIDEnumSlack.ToPointer(),
        IntegrationIdentifier: v3.Pointer("slack-prod"),
        ContextKeys: []string{
            "tenant:org-123",
            "region:us-east-1",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListChannelConnectionsResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.ChannelConnectionsControllerListChannelConnectionsRequest](../../models/operations/channelconnectionscontrollerlistchannelconnectionsrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |
| `opts`                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                           | The options for this request.                                                                                                                                |

### Response

**[*operations.ChannelConnectionsControllerListChannelConnectionsResponse](../../models/operations/channelconnectionscontrollerlistchannelconnectionsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Create

Create a new channel connection for a resource for given integration. Only one channel connection is allowed per resource and integration.

### Example Usage

<!-- UsageSnippet language="go" operationID="ChannelConnectionsController_createChannelConnection" method="post" path="/v1/channel-connections" -->
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

    res, err := s.ChannelConnections.Create(ctx, components.CreateChannelConnectionRequestDto{
        Identifier: v3.Pointer("slack-prod-user123-abc4"),
        SubscriberID: v3.Pointer("subscriber-123"),
        Context: map[string]components.CreateChannelConnectionRequestDtoContext{
            "key": components.CreateCreateChannelConnectionRequestDtoContextStr(
                "org-acme",
            ),
        },
        IntegrationIdentifier: "slack-prod",
        Workspace: components.WorkspaceDto{
            ID: "T123456",
            Name: v3.Pointer("Acme HQ"),
        },
        Auth: components.AuthDto{
            AccessToken: "Workspace access token",
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetChannelConnectionResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `createChannelConnectionRequestDto`                                                                          | [components.CreateChannelConnectionRequestDto](../../models/components/createchannelconnectionrequestdto.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `idempotencyKey`                                                                                             | `*string`                                                                                                    | :heavy_minus_sign:                                                                                           | A header for idempotency purposes                                                                            |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ChannelConnectionsControllerCreateChannelConnectionResponse](../../models/operations/channelconnectionscontrollercreatechannelconnectionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Retrieve

Retrieve a specific channel connection by its unique identifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="ChannelConnectionsController_getChannelConnectionByIdentifier" method="get" path="/v1/channel-connections/{identifier}" -->
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

    res, err := s.ChannelConnections.Retrieve(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetChannelConnectionResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `identifier`                                             | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the channel connection          |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ChannelConnectionsControllerGetChannelConnectionByIdentifierResponse](../../models/operations/channelconnectionscontrollergetchannelconnectionbyidentifierresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Update

Update an existing channel connection by its unique identifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="ChannelConnectionsController_updateChannelConnection" method="patch" path="/v1/channel-connections/{identifier}" -->
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

    res, err := s.ChannelConnections.Update(ctx, "<value>", components.UpdateChannelConnectionRequestDto{
        Workspace: components.WorkspaceDto{
            ID: "T123456",
            Name: v3.Pointer("Acme HQ"),
        },
        Auth: components.AuthDto{
            AccessToken: "Workspace access token",
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetChannelConnectionResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `identifier`                                                                                                 | `string`                                                                                                     | :heavy_check_mark:                                                                                           | The unique identifier of the channel connection                                                              |
| `updateChannelConnectionRequestDto`                                                                          | [components.UpdateChannelConnectionRequestDto](../../models/components/updatechannelconnectionrequestdto.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `idempotencyKey`                                                                                             | `*string`                                                                                                    | :heavy_minus_sign:                                                                                           | A header for idempotency purposes                                                                            |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ChannelConnectionsControllerUpdateChannelConnectionResponse](../../models/operations/channelconnectionscontrollerupdatechannelconnectionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Delete

Delete a specific channel connection by its unique identifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="ChannelConnectionsController_deleteChannelConnection" method="delete" path="/v1/channel-connections/{identifier}" -->
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

    res, err := s.ChannelConnections.Delete(ctx, "<value>", nil)
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
| `identifier`                                             | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the channel connection          |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ChannelConnectionsControllerDeleteChannelConnectionResponse](../../models/operations/channelconnectionscontrollerdeletechannelconnectionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |