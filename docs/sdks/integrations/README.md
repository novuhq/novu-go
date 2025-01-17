# Integrations
(*Integrations*)

## Overview

With the help of the Integration Store, you can easily integrate your favorite delivery provider. During the runtime of the API, the Integrations Store is responsible for storing the configurations of all the providers.
<https://docs.novu.co/channels-and-providers/integration-store>

### Available Operations

* [List](#list) - Get integrations
* [Create](#create) - Create integration
* [ListActive](#listactive) - Get active integrations
* [Update](#update) - Update integration
* [Delete](#delete) - Delete integration
* [SetPrimary](#setprimary) - Set integration as primary

## List

Return all the integrations the user has created for that organization. Review v.0.17.0 changelog for a breaking change

### Example Usage

```go
package main

import(
	"context"
	"os"
	novugo "github.com/novuhq/novu-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := novugo.New(
        novugo.WithSecurity(os.Getenv("NOVU_API_KEY")),
    )

    res, err := s.Integrations.List(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.IntegrationsControllerListIntegrationsResponse](../../models/operations/integrationscontrollerlistintegrationsresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 404, 409                | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## Create

Create an integration for the current environment the user is based on the API key provided

### Example Usage

```go
package main

import(
	"context"
	"os"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := novugo.New(
        novugo.WithSecurity(os.Getenv("NOVU_API_KEY")),
    )

    res, err := s.Integrations.Create(ctx, components.CreateIntegrationRequestDto{
        ProviderID: "<id>",
        Channel: components.CreateIntegrationRequestDtoChannelSms,
    })
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
| `request`                                                                                        | [components.CreateIntegrationRequestDto](../../models/components/createintegrationrequestdto.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.IntegrationsControllerCreateIntegrationResponse](../../models/operations/integrationscontrollercreateintegrationresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 404, 409                | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## ListActive

Return all the active integrations the user has created for that organization. Review v.0.17.0 changelog for a breaking change

### Example Usage

```go
package main

import(
	"context"
	"os"
	novugo "github.com/novuhq/novu-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := novugo.New(
        novugo.WithSecurity(os.Getenv("NOVU_API_KEY")),
    )

    res, err := s.Integrations.ListActive(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.IntegrationsControllerGetActiveIntegrationsResponse](../../models/operations/integrationscontrollergetactiveintegrationsresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 404, 409                | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## Update

Update integration

### Example Usage

```go
package main

import(
	"context"
	"os"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := novugo.New(
        novugo.WithSecurity(os.Getenv("NOVU_API_KEY")),
    )

    res, err := s.Integrations.Update(ctx, "<id>", components.UpdateIntegrationRequestDto{})
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
| `integrationID`                                                                                  | *string*                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `updateIntegrationRequestDto`                                                                    | [components.UpdateIntegrationRequestDto](../../models/components/updateintegrationrequestdto.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.IntegrationsControllerUpdateIntegrationByIDResponse](../../models/operations/integrationscontrollerupdateintegrationbyidresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 409                     | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## Delete

Delete integration

### Example Usage

```go
package main

import(
	"context"
	"os"
	novugo "github.com/novuhq/novu-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := novugo.New(
        novugo.WithSecurity(os.Getenv("NOVU_API_KEY")),
    )

    res, err := s.Integrations.Delete(ctx, "<id>")
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
| `integrationID`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.IntegrationsControllerRemoveIntegrationResponse](../../models/operations/integrationscontrollerremoveintegrationresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 404, 409                | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## SetPrimary

Set integration as primary

### Example Usage

```go
package main

import(
	"context"
	"os"
	novugo "github.com/novuhq/novu-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := novugo.New(
        novugo.WithSecurity(os.Getenv("NOVU_API_KEY")),
    )

    res, err := s.Integrations.SetPrimary(ctx, "<id>")
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
| `integrationID`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.IntegrationsControllerSetIntegrationAsPrimaryResponse](../../models/operations/integrationscontrollersetintegrationasprimaryresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 409                     | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |