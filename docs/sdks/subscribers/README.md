# Subscribers
(*Subscribers*)

## Overview

A subscriber in Novu represents someone who should receive a message. A subscriber's profile information contains important attributes about the subscriber that will be used in messages (name, email). The subscriber object can contain other key-value pairs that can be used to further personalize your messages.
<https://docs.novu.co/subscribers/subscribers>

### Available Operations

* [Search](#search) - Search subscribers
* [Create](#create) - Create a subscriber
* [Retrieve](#retrieve) - Retrieve a subscriber
* [Patch](#patch) - Update a subscriber
* [Delete](#delete) - Delete a subscriber
* [CreateBulk](#createbulk) - Bulk create subscribers

## Search

Search subscribers by their **email**, **phone**, **subscriberId** and **name**. 
    The search is case sensitive and supports pagination.Checkout all available filters in the query section.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_searchSubscribers" method="get" path="/v2/subscribers" -->
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

    res, err := s.Subscribers.Search(ctx, operations.SubscribersControllerSearchSubscribersRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListSubscribersResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.SubscribersControllerSearchSubscribersRequest](../../models/operations/subscriberscontrollersearchsubscribersrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.SubscribersControllerSearchSubscribersResponse](../../models/operations/subscriberscontrollersearchsubscribersresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Create

Create a subscriber with the subscriber attributes. 
      **subscriberId** is a required field, rest other fields are optional, if the subscriber already exists, it will be updated

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_createSubscriber" method="post" path="/v2/subscribers" -->
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

    res, err := s.Subscribers.Create(ctx, components.CreateSubscriberRequestDto{
        FirstName: v3.Pointer("John"),
        LastName: v3.Pointer("Doe"),
        Email: v3.Pointer("john.doe@example.com"),
        Phone: v3.Pointer("+1234567890"),
        Avatar: v3.Pointer("https://example.com/avatar.jpg"),
        Locale: v3.Pointer("en-US"),
        Timezone: v3.Pointer("America/New_York"),
        SubscriberID: "<id>",
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriberResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `createSubscriberRequestDto`                                                                   | [components.CreateSubscriberRequestDto](../../models/components/createsubscriberrequestdto.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `failIfExists`                                                                                 | **bool*                                                                                        | :heavy_minus_sign:                                                                             | If true, the request will fail if a subscriber with the same subscriberId already exists       |
| `idempotencyKey`                                                                               | **string*                                                                                      | :heavy_minus_sign:                                                                             | A header for idempotency purposes                                                              |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.SubscribersControllerCreateSubscriberResponse](../../models/operations/subscriberscontrollercreatesubscriberresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.SubscriberResponseDto   | 409                               | application/json                  |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 404, 405, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Retrieve

Retrieve a subscriber by its unique key identifier **subscriberId**. 
    **subscriberId** field is required.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_getSubscriber" method="get" path="/v2/subscribers/{subscriberId}" -->
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

    res, err := s.Subscribers.Retrieve(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriberResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `subscriberID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SubscribersControllerGetSubscriberResponse](../../models/operations/subscriberscontrollergetsubscriberresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Patch

Update a subscriber by its unique key identifier **subscriberId**. 
    **subscriberId** is a required field, rest other fields are optional

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_patchSubscriber" method="patch" path="/v2/subscribers/{subscriberId}" -->
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

    res, err := s.Subscribers.Patch(ctx, "<id>", components.PatchSubscriberRequestDto{
        FirstName: v3.Pointer("John"),
        LastName: v3.Pointer("Doe"),
        Email: v3.Pointer("john.doe@example.com"),
        Phone: v3.Pointer("+1234567890"),
        Avatar: v3.Pointer("https://example.com/avatar.jpg"),
        Locale: v3.Pointer("en-US"),
        Timezone: v3.Pointer("America/New_York"),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriberResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `subscriberID`                                                                               | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `patchSubscriberRequestDto`                                                                  | [components.PatchSubscriberRequestDto](../../models/components/patchsubscriberrequestdto.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `idempotencyKey`                                                                             | **string*                                                                                    | :heavy_minus_sign:                                                                           | A header for idempotency purposes                                                            |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.SubscribersControllerPatchSubscriberResponse](../../models/operations/subscriberscontrollerpatchsubscriberresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Delete

Deletes a subscriber entity from the Novu platform along with associated messages, preferences, and topic subscriptions. 
      **subscriberId** is a required field.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_removeSubscriber" method="delete" path="/v2/subscribers/{subscriberId}" -->
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

    res, err := s.Subscribers.Delete(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RemoveSubscriberResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `subscriberID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SubscribersControllerRemoveSubscriberResponse](../../models/operations/subscriberscontrollerremovesubscriberresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateBulk


      Using this endpoint multiple subscribers can be created at once. The bulk API is limited to 500 subscribers per request.
    

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersV1Controller_bulkCreateSubscribers" method="post" path="/v1/subscribers/bulk" -->
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

    res, err := s.Subscribers.CreateBulk(ctx, components.BulkSubscriberCreateDto{
        Subscribers: []components.CreateSubscriberRequestDto{
            components.CreateSubscriberRequestDto{
                SubscriberID: "<id>",
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.BulkCreateSubscriberResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `bulkSubscriberCreateDto`                                                                | [components.BulkSubscriberCreateDto](../../models/components/bulksubscribercreatedto.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `idempotencyKey`                                                                         | **string*                                                                                | :heavy_minus_sign:                                                                       | A header for idempotency purposes                                                        |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.SubscribersV1ControllerBulkCreateSubscribersResponse](../../models/operations/subscribersv1controllerbulkcreatesubscribersresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |