# ChannelEndpoints

## Overview

### Available Operations

* [List](#list) - List all channel endpoints
* [Create](#create) - Create a channel endpoint
* [Retrieve](#retrieve) - Retrieve a channel endpoint
* [Update](#update) - Update a channel endpoint
* [Delete](#delete) - Delete a channel endpoint

## List

List all channel endpoints for a resource based on query filters.

### Example Usage

<!-- UsageSnippet language="go" operationID="ChannelEndpointsController_listChannelEndpoints" method="get" path="/v1/channel-endpoints" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"github.com/novuhq/novu-go/v3/models/components"
	"github.com/novuhq/novu-go/v3/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.ChannelEndpoints.List(ctx, operations.ChannelEndpointsControllerListChannelEndpointsRequest{
        Limit: v3.Pointer[float64](10.0),
        SubscriberID: v3.Pointer("subscriber-123"),
        ContextKeys: []string{
            "tenant:org-123",
            "region:us-east-1",
        },
        ProviderID: components.ProvidersIDEnumSlack.ToPointer(),
        IntegrationIdentifier: v3.Pointer("slack-prod"),
        ConnectionIdentifier: v3.Pointer("slack-connection-abc123"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListChannelEndpointsResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.ChannelEndpointsControllerListChannelEndpointsRequest](../../models/operations/channelendpointscontrollerlistchannelendpointsrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.ChannelEndpointsControllerListChannelEndpointsResponse](../../models/operations/channelendpointscontrollerlistchannelendpointsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Create

Create a new channel endpoint for a resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="ChannelEndpointsController_createChannelEndpoint" method="post" path="/v1/channel-endpoints" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"github.com/novuhq/novu-go/v3/models/components"
	"github.com/novuhq/novu-go/v3/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.ChannelEndpoints.Create(ctx, operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodySlackChannel(
        components.CreateSlackChannelEndpointDto{
            SubscriberID: "subscriber-123",
            IntegrationIdentifier: "slack-prod",
            Type: components.CreateSlackChannelEndpointDtoTypeSlackChannel,
            Endpoint: components.SlackChannelEndpointDto{
                ChannelID: "C123456789",
            },
        },
    ), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetChannelEndpointResponseDto != nil {
        switch res.GetChannelEndpointResponseDto.Endpoint.Type {
            case components.EndpointTypeSlackChannelEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.SlackChannelEndpointDto is populated
            case components.EndpointTypeSlackUserEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.SlackUserEndpointDto is populated
            case components.EndpointTypeWebhookEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.WebhookEndpointDto is populated
            case components.EndpointTypePhoneEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.PhoneEndpointDto is populated
            case components.EndpointTypeMsTeamsChannelEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.MsTeamsChannelEndpointDto is populated
            case components.EndpointTypeMsTeamsUserEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.MsTeamsUserEndpointDto is populated
            case components.EndpointTypeTelegramChatEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.TelegramChatEndpointDto is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `requestBody`                                                                                                                                                  | [operations.ChannelEndpointsControllerCreateChannelEndpointRequestBody](../../models/operations/channelendpointscontrollercreatechannelendpointrequestbody.md) | :heavy_check_mark:                                                                                                                                             | Channel endpoint creation request. The structure varies based on the type field.                                                                               |
| `idempotencyKey`                                                                                                                                               | `*string`                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                             | A header for idempotency purposes                                                                                                                              |
| `opts`                                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                             | The options for this request.                                                                                                                                  |

### Response

**[*operations.ChannelEndpointsControllerCreateChannelEndpointResponse](../../models/operations/channelendpointscontrollercreatechannelendpointresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Retrieve

Retrieve a specific channel endpoint by its unique identifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="ChannelEndpointsController_getChannelEndpoint" method="get" path="/v1/channel-endpoints/{identifier}" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"log"
	"github.com/novuhq/novu-go/v3/models/components"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.ChannelEndpoints.Retrieve(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetChannelEndpointResponseDto != nil {
        switch res.GetChannelEndpointResponseDto.Endpoint.Type {
            case components.EndpointTypeSlackChannelEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.SlackChannelEndpointDto is populated
            case components.EndpointTypeSlackUserEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.SlackUserEndpointDto is populated
            case components.EndpointTypeWebhookEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.WebhookEndpointDto is populated
            case components.EndpointTypePhoneEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.PhoneEndpointDto is populated
            case components.EndpointTypeMsTeamsChannelEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.MsTeamsChannelEndpointDto is populated
            case components.EndpointTypeMsTeamsUserEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.MsTeamsUserEndpointDto is populated
            case components.EndpointTypeTelegramChatEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.TelegramChatEndpointDto is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `identifier`                                             | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the channel endpoint            |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ChannelEndpointsControllerGetChannelEndpointResponse](../../models/operations/channelendpointscontrollergetchannelendpointresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Update

Update an existing channel endpoint by its unique identifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="ChannelEndpointsController_updateChannelEndpoint" method="patch" path="/v1/channel-endpoints/{identifier}" -->
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

    res, err := s.ChannelEndpoints.Update(ctx, "<value>", components.UpdateChannelEndpointRequestDto{
        Endpoint: components.CreateUpdateChannelEndpointRequestDtoEndpointSlackUserEndpointDto(
            components.SlackUserEndpointDto{
                UserID: "U123456789",
            },
        ),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetChannelEndpointResponseDto != nil {
        switch res.GetChannelEndpointResponseDto.Endpoint.Type {
            case components.EndpointTypeSlackChannelEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.SlackChannelEndpointDto is populated
            case components.EndpointTypeSlackUserEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.SlackUserEndpointDto is populated
            case components.EndpointTypeWebhookEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.WebhookEndpointDto is populated
            case components.EndpointTypePhoneEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.PhoneEndpointDto is populated
            case components.EndpointTypeMsTeamsChannelEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.MsTeamsChannelEndpointDto is populated
            case components.EndpointTypeMsTeamsUserEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.MsTeamsUserEndpointDto is populated
            case components.EndpointTypeTelegramChatEndpointDto:
                // res.GetChannelEndpointResponseDto.Endpoint.TelegramChatEndpointDto is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `identifier`                                                                                             | `string`                                                                                                 | :heavy_check_mark:                                                                                       | The unique identifier of the channel endpoint                                                            |
| `updateChannelEndpointRequestDto`                                                                        | [components.UpdateChannelEndpointRequestDto](../../models/components/updatechannelendpointrequestdto.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `idempotencyKey`                                                                                         | `*string`                                                                                                | :heavy_minus_sign:                                                                                       | A header for idempotency purposes                                                                        |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ChannelEndpointsControllerUpdateChannelEndpointResponse](../../models/operations/channelendpointscontrollerupdatechannelendpointresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Delete

Delete a specific channel endpoint by its unique identifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="ChannelEndpointsController_deleteChannelEndpoint" method="delete" path="/v1/channel-endpoints/{identifier}" -->
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

    res, err := s.ChannelEndpoints.Delete(ctx, "<value>", nil)
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
| `identifier`                                             | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the channel endpoint            |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ChannelEndpointsControllerDeleteChannelEndpointResponse](../../models/operations/channelendpointscontrollerdeletechannelendpointresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |