# Domains.AutoConfigure

## Overview

### Available Operations

* [Retrieve](#retrieve) - Retrieve auto-configuration availability
* [Start](#start) - Start DNS auto-configuration

## Retrieve

Returns whether DNS auto-configuration (Domain Connect) is available for this domain. When `available` is `false`, `manualRecords` lists the DNS records the customer must add manually.

### Example Usage

<!-- UsageSnippet language="go" operationID="DomainsController_getDomainAutoConfigure" method="get" path="/v1/domains/{domain}/auto-configure" -->
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

    res, err := s.Domains.AutoConfigure.Retrieve(ctx, "hidden-subsidy.info", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DomainConnectStatusResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `domain`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DomainsControllerGetDomainAutoConfigureResponse](../../models/operations/domainscontrollergetdomainautoconfigureresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Start

Generates a signed redirect URL the customer can follow to apply Novu DNS records at their DNS provider. After the provider completes the flow, it redirects back to `redirectUri`.

### Example Usage

<!-- UsageSnippet language="go" operationID="DomainsController_startDomainAutoConfigure" method="post" path="/v1/domains/{domain}/auto-configure/start" -->
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

    res, err := s.Domains.AutoConfigure.Start(ctx, "criminal-other.name", components.CreateDomainConnectApplyURLDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DomainConnectApplyURLResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `domain`                                                                                               | `string`                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `createDomainConnectApplyURLDto`                                                                       | [components.CreateDomainConnectApplyURLDto](../../models/components/createdomainconnectapplyurldto.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `idempotencyKey`                                                                                       | `*string`                                                                                              | :heavy_minus_sign:                                                                                     | A header for idempotency purposes                                                                      |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.DomainsControllerStartDomainAutoConfigureResponse](../../models/operations/domainscontrollerstartdomainautoconfigureresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |