# Domains.Routes

## Overview

### Available Operations

* [List](#list) - List routes for a domain
* [Create](#create) - Create a route
* [Retrieve](#retrieve) - Retrieve a route by address
* [Update](#update) - Update a route
* [Delete](#delete) - Delete a route
* [Test](#test) - Test an inbound route

## List

Returns a paginated list of routes attached to the domain. Optionally filter by an agent identifier to find routes pointing to a specific agent.

### Example Usage

<!-- UsageSnippet language="go" operationID="DomainsController_listDomainRoutes" method="get" path="/v1/domains/{domain}/routes" -->
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

    res, err := s.Domains.Routes.List(ctx, operations.DomainsControllerListDomainRoutesRequest{
        Domain: "fearless-fishery.com",
        Limit: v3.Pointer[float64](10.0),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListDomainRoutesResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.DomainsControllerListDomainRoutesRequest](../../models/operations/domainscontrollerlistdomainroutesrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.DomainsControllerListDomainRoutesResponse](../../models/operations/domainscontrollerlistdomainroutesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Create

Creates a route on the domain that forwards inbound mail addressed to `<address>@<domain>` to either a webhook or an agent. Each address on a domain may only have a single route.

### Example Usage

<!-- UsageSnippet language="go" operationID="DomainsController_createDomainRoute" method="post" path="/v1/domains/{domain}/routes" -->
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

    res, err := s.Domains.Routes.Create(ctx, "radiant-solvency.net", components.DomainRouteDto{
        Address: "6581 Birch Road",
        Type: components.DomainRouteDtoTypeWebhook,
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DomainRouteResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `domain`                                                               | `string`                                                               | :heavy_check_mark:                                                     | N/A                                                                    |
| `domainRouteDto`                                                       | [components.DomainRouteDto](../../models/components/domainroutedto.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `idempotencyKey`                                                       | `*string`                                                              | :heavy_minus_sign:                                                     | A header for idempotency purposes                                      |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.DomainsControllerCreateDomainRouteResponse](../../models/operations/domainscontrollercreatedomainrouteresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Retrieve

Returns the route bound to `<address>@<domain>`. Use `*` as the address to retrieve the wildcard route for the domain.

### Example Usage

<!-- UsageSnippet language="go" operationID="DomainsController_getDomainRoute" method="get" path="/v1/domains/{domain}/routes/{address}" -->
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

    res, err := s.Domains.Routes.Retrieve(ctx, "adolescent-petal.net", "42531 Green Lane", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DomainRouteResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `domain`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `address`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DomainsControllerGetDomainRouteResponse](../../models/operations/domainscontrollergetdomainrouteresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Update

Updates the destination of the route bound to `<address>@<domain>`. The address itself is the resource identity and cannot be changed; delete and recreate the route to rename it.

### Example Usage

<!-- UsageSnippet language="go" operationID="DomainsController_updateDomainRoute" method="patch" path="/v1/domains/{domain}/routes/{address}" -->
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

    res, err := s.Domains.Routes.Update(ctx, "cavernous-cycle.com", "70213 Gerlach Rue", components.UpdateDomainRouteDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DomainRouteResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `domain`                                                                           | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |
| `address`                                                                          | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |
| `updateDomainRouteDto`                                                             | [components.UpdateDomainRouteDto](../../models/components/updatedomainroutedto.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `idempotencyKey`                                                                   | `*string`                                                                          | :heavy_minus_sign:                                                                 | A header for idempotency purposes                                                  |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.DomainsControllerUpdateDomainRouteResponse](../../models/operations/domainscontrollerupdatedomainrouteresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Delete

Removes the route bound to `<address>@<domain>`. Inbound mail for that address will no longer be processed.

### Example Usage

<!-- UsageSnippet language="go" operationID="DomainsController_deleteDomainRoute" method="delete" path="/v1/domains/{domain}/routes/{address}" -->
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

    res, err := s.Domains.Routes.Delete(ctx, "corrupt-avalanche.biz", "753 W 4th Avenue", nil)
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
| `domain`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `address`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DomainsControllerDeleteDomainRouteResponse](../../models/operations/domainscontrollerdeletedomainrouteresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Test

Sends a synthetic inbound email through the same delivery path as production (outbound webhooks for webhook routes, signed HTTP to the agent for agent routes). Use `dryRun: true` to preview the payload without delivering.

### Example Usage

<!-- UsageSnippet language="go" operationID="DomainsController_testDomainRoute" method="post" path="/v1/domains/{domain}/routes/{address}/test" -->
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

    res, err := s.Domains.Routes.Test(ctx, "exalted-bonfire.com", "90499 Rowan Close", components.TestDomainRouteDto{
        From: components.TestDomainRouteFromDto{
            Address: "58851 Konopelski Overpass",
        },
        Subject: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TestDomainRouteResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `domain`                                                                       | `string`                                                                       | :heavy_check_mark:                                                             | N/A                                                                            |
| `address`                                                                      | `string`                                                                       | :heavy_check_mark:                                                             | N/A                                                                            |
| `testDomainRouteDto`                                                           | [components.TestDomainRouteDto](../../models/components/testdomainroutedto.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `idempotencyKey`                                                               | `*string`                                                                      | :heavy_minus_sign:                                                             | A header for idempotency purposes                                              |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.DomainsControllerTestDomainRouteResponse](../../models/operations/domainscontrollertestdomainrouteresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |