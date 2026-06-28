# Integrations

## Overview

With the help of the Integration Store, you can easily integrate your favorite delivery provider. During the runtime of the API, the Integrations Store is responsible for storing the configurations of all the providers.
<https://docs.novu.co/platform/integrations/overview>

### Available Operations

* [List](#list) - List all integrations
* [Create](#create) - Create an integration
* [Update](#update) - Update an integration
* [Delete](#delete) - Delete an integration
* [IntegrationsControllerAutoConfigureIntegration](#integrationscontrollerautoconfigureintegration) - Auto-configure an integration for inbound webhooks
* [SetAsPrimary](#setasprimary) - Update integration as primary
* [ListActive](#listactive) - List active integrations
* [GenerateConnectOAuthURL](#generateconnectoauthurl) - Generate OAuth URL for a workspace/tenant connection
* [GenerateLinkUserOAuthURL](#generatelinkuseroauthurl) - Generate OAuth URL to link a subscriber user identity
* [~~GenerateChatOAuthURL~~](#generatechatoauthurl) - Generate chat OAuth URL :warning: **Deprecated**

## List

List all the channels integrations created in the organization. Only integration metadata is returned, credentials field is returned as an empty object.

### Example Usage

<!-- UsageSnippet language="go" operationID="IntegrationsController_listIntegrations" method="get" path="/v1/integrations" -->
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

    res, err := s.Integrations.List(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IntegrationResponseDtos != nil {
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

**[*operations.IntegrationsControllerListIntegrationsResponse](../../models/operations/integrationscontrollerlistintegrationsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Create

Create an integration for the current environment the user is based on the API key provided. 
    Each provider supports different credentials, check the provider documentation for more details. Only integration metadata is returned, credentials field is returned as an empty object.

### Example Usage

<!-- UsageSnippet language="go" operationID="IntegrationsController_createIntegration" method="post" path="/v1/integrations" -->
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

    res, err := s.Integrations.Create(ctx, components.CreateIntegrationRequestDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IntegrationResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `createIntegrationRequestDto`                                                                    | [components.CreateIntegrationRequestDto](../../models/components/createintegrationrequestdto.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `idempotencyKey`                                                                                 | `*string`                                                                                        | :heavy_minus_sign:                                                                               | A header for idempotency purposes                                                                |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.IntegrationsControllerCreateIntegrationResponse](../../models/operations/integrationscontrollercreateintegrationresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Update

Update an integration by its unique key identifier **integrationId**. 
    Each provider supports different credentials, check the provider documentation for more details. Only integration metadata is returned, credentials field is returned as an empty object.

### Example Usage

<!-- UsageSnippet language="go" operationID="IntegrationsController_updateIntegrationById" method="put" path="/v1/integrations/{integrationId}" -->
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

    res, err := s.Integrations.Update(ctx, "<id>", components.UpdateIntegrationRequestDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IntegrationResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `integrationID`                                                                                  | `string`                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `updateIntegrationRequestDto`                                                                    | [components.UpdateIntegrationRequestDto](../../models/components/updateintegrationrequestdto.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `idempotencyKey`                                                                                 | `*string`                                                                                        | :heavy_minus_sign:                                                                               | A header for idempotency purposes                                                                |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.IntegrationsControllerUpdateIntegrationByIDResponse](../../models/operations/integrationscontrollerupdateintegrationbyidresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Delete

Delete an integration by its unique key identifier **integrationId**. 
    This action is irreversible. Only integration metadata is returned, credentials field is returned as empty object.

### Example Usage

<!-- UsageSnippet language="go" operationID="IntegrationsController_removeIntegration" method="delete" path="/v1/integrations/{integrationId}" -->
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

    res, err := s.Integrations.Delete(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IntegrationResponseDtos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `integrationID`                                          | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.IntegrationsControllerRemoveIntegrationResponse](../../models/operations/integrationscontrollerremoveintegrationresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## IntegrationsControllerAutoConfigureIntegration

Auto-configure an integration by its unique key identifier **integrationId** for inbound webhook support. 
    This will automatically generate required webhook signing keys and configure webhook endpoints. Only integration metadata is returned, credentials field is returned as an empty object.

### Example Usage

<!-- UsageSnippet language="go" operationID="IntegrationsController_autoConfigureIntegration" method="post" path="/v1/integrations/{integrationId}/auto-configure" -->
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

    res, err := s.Integrations.IntegrationsControllerAutoConfigureIntegration(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AutoConfigureIntegrationResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `integrationID`                                          | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.IntegrationsControllerAutoConfigureIntegrationResponse](../../models/operations/integrationscontrollerautoconfigureintegrationresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## SetAsPrimary

Update an integration as **primary** by its unique key identifier **integrationId**. 
    This API will set the integration as primary for that channel in the current environment. 
    Primary integration is used to deliver notification for sms and email channels in the workflow. 
    Only integration metadata is returned, credentials field is returned as an empty object.

### Example Usage

<!-- UsageSnippet language="go" operationID="IntegrationsController_setIntegrationAsPrimary" method="post" path="/v1/integrations/{integrationId}/set-primary" -->
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

    res, err := s.Integrations.SetAsPrimary(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IntegrationResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `integrationID`                                          | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.IntegrationsControllerSetIntegrationAsPrimaryResponse](../../models/operations/integrationscontrollersetintegrationasprimaryresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## ListActive

List all the active integrations created in the organization. Only integration metadata is returned, credentials field is returned as an empty object.

### Example Usage

<!-- UsageSnippet language="go" operationID="IntegrationsController_getActiveIntegrations" method="get" path="/v1/integrations/active" -->
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

    res, err := s.Integrations.ListActive(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IntegrationResponseDtos != nil {
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

**[*operations.IntegrationsControllerGetActiveIntegrationsResponse](../../models/operations/integrationscontrollergetactiveintegrationsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## GenerateConnectOAuthURL

Generate an OAuth URL that creates a workspace or tenant-level channel connection (Slack workspace install or MS Teams admin consent). 
    The generated URL expires after 5 minutes.

### Example Usage

<!-- UsageSnippet language="go" operationID="IntegrationsController_generateConnectOAuthUrl" method="post" path="/v1/integrations/channel-connections/oauth" -->
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

    res, err := s.Integrations.GenerateConnectOAuthURL(ctx, components.GenerateConnectOauthURLRequestDto{
        SubscriberID: v3.Pointer("subscriber-123"),
        IntegrationIdentifier: "<value>",
        ConnectionIdentifier: v3.Pointer("slack-connection-abc123"),
        Context: map[string]components.GenerateConnectOauthURLRequestDtoContext{
            "key": components.CreateGenerateConnectOauthURLRequestDtoContextStr(
                "org-acme",
            ),
        },
        Scope: []string{
            "chat:write",
            "chat:write.public",
            "channels:read",
        },
        ConnectionMode: components.GenerateConnectOauthURLRequestDtoConnectionModeShared.ToPointer(),
        AutoLinkUser: v3.Pointer(true),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GenerateChatOAuthURLResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `generateConnectOauthURLRequestDto`                                                                          | [components.GenerateConnectOauthURLRequestDto](../../models/components/generateconnectoauthurlrequestdto.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `idempotencyKey`                                                                                             | `*string`                                                                                                    | :heavy_minus_sign:                                                                                           | A header for idempotency purposes                                                                            |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.IntegrationsControllerGenerateConnectOAuthURLResponse](../../models/operations/integrationscontrollergenerateconnectoauthurlresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## GenerateLinkUserOAuthURL

Generate an OAuth URL that links a specific subscriber to their chat identity (Slack user ID or MS Teams user OID). 
    The generated URL expires after 5 minutes.

### Example Usage

<!-- UsageSnippet language="go" operationID="IntegrationsController_generateLinkUserOAuthUrl" method="post" path="/v1/integrations/channel-endpoints/oauth" -->
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

    res, err := s.Integrations.GenerateLinkUserOAuthURL(ctx, components.GenerateLinkUserOauthURLRequestDto{
        SubscriberID: "subscriber-123",
        IntegrationIdentifier: "<value>",
        ConnectionIdentifier: v3.Pointer("slack-connection-abc123"),
        Context: map[string]components.GenerateLinkUserOauthURLRequestDtoContext{
            "key": components.CreateGenerateLinkUserOauthURLRequestDtoContextStr(
                "org-acme",
            ),
        },
        UserScope: []string{
            "identity.basic",
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GenerateChatOAuthURLResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `generateLinkUserOauthURLRequestDto`                                                                           | [components.GenerateLinkUserOauthURLRequestDto](../../models/components/generatelinkuseroauthurlrequestdto.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `idempotencyKey`                                                                                               | `*string`                                                                                                      | :heavy_minus_sign:                                                                                             | A header for idempotency purposes                                                                              |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.IntegrationsControllerGenerateLinkUserOAuthURLResponse](../../models/operations/integrationscontrollergeneratelinkuseroauthurlresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ~~GenerateChatOAuthURL~~

**Deprecated** — use `POST /integrations/channel-connections/oauth` (connect) or `POST /integrations/channel-endpoints/oauth` (link_user) instead.
    Generate an OAuth URL for chat integrations like Slack and MS Teams. 
    This URL allows subscribers to authorize the integration, enabling the system to send messages 
    through their chat workspace. The generated URL expires after 5 minutes.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="IntegrationsController_getChatOAuthUrl" method="post" path="/v1/integrations/chat/oauth" -->
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

    res, err := s.Integrations.GenerateChatOAuthURL(ctx, components.GenerateChatOauthURLRequestDto{
        SubscriberID: v3.Pointer("subscriber-123"),
        IntegrationIdentifier: "<value>",
        ConnectionIdentifier: v3.Pointer("slack-connection-abc123"),
        Context: map[string]components.GenerateChatOauthURLRequestDtoContext{
            "key": components.CreateGenerateChatOauthURLRequestDtoContextStr(
                "org-acme",
            ),
        },
        Scope: []string{
            "chat:write",
            "chat:write.public",
            "channels:read",
            "groups:read",
            "users:read",
            "users:read.email",
            "incoming-webhook",
        },
        UserScope: []string{
            "identity.basic",
        },
        Mode: components.ModeLinkUser.ToPointer(),
        ConnectionMode: components.GenerateChatOauthURLRequestDtoConnectionModeShared.ToPointer(),
        AutoLinkUser: v3.Pointer(true),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GenerateChatOAuthURLResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `generateChatOauthURLRequestDto`                                                                       | [components.GenerateChatOauthURLRequestDto](../../models/components/generatechatoauthurlrequestdto.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `idempotencyKey`                                                                                       | `*string`                                                                                              | :heavy_minus_sign:                                                                                     | A header for idempotency purposes                                                                      |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.IntegrationsControllerGetChatOAuthURLResponse](../../models/operations/integrationscontrollergetchatoauthurlresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |