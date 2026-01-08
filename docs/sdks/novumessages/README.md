# Subscribers.Messages

## Overview

### Available Operations

* [UpdateAsSeen](#updateasseen) - Update notification action status
* [MarkAll](#markall) - Update all notifications state
* [MarkAllAs](#markallas) - Update notifications state

## UpdateAsSeen

Update in-app (inbox) notification's action status by its unique key identifier **messageId** and type field **type**. 
      **type** field can be **primary** or **secondary**

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersV1Controller_markActionAsSeen" method="post" path="/v1/subscribers/{subscriberId}/messages/{messageId}/actions/{type}" -->
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

    res, err := s.Subscribers.Messages.UpdateAsSeen(ctx, operations.SubscribersV1ControllerMarkActionAsSeenRequest{
        MessageID: "<id>",
        Type: "<value>",
        SubscriberID: "<id>",
        MarkMessageActionAsSeenDto: components.MarkMessageActionAsSeenDto{
            Status: components.MarkMessageActionAsSeenDtoStatusPending,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MessageResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.SubscribersV1ControllerMarkActionAsSeenRequest](../../models/operations/subscribersv1controllermarkactionasseenrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.SubscribersV1ControllerMarkActionAsSeenResponse](../../models/operations/subscribersv1controllermarkactionasseenresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## MarkAll

Update all subscriber in-app (inbox) notifications state such as read, unread, seen or unseen by **subscriberId**.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersV1Controller_markAllUnreadAsRead" method="post" path="/v1/subscribers/{subscriberId}/messages/mark-all" -->
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

    res, err := s.Subscribers.Messages.MarkAll(ctx, "<id>", components.MarkAllMessageAsRequestDto{
        MarkAs: components.MarkAsRead,
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Number != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `subscriberID`                                                                                 | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `markAllMessageAsRequestDto`                                                                   | [components.MarkAllMessageAsRequestDto](../../models/components/markallmessageasrequestdto.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `idempotencyKey`                                                                               | **string*                                                                                      | :heavy_minus_sign:                                                                             | A header for idempotency purposes                                                              |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.SubscribersV1ControllerMarkAllUnreadAsReadResponse](../../models/operations/subscribersv1controllermarkallunreadasreadresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## MarkAllAs

Update subscriber's multiple in-app (inbox) notifications state such as seen, read, unseen or unread by **subscriberId**. 
      **messageId** is of type mongodbId of notifications

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersV1Controller_markMessagesAs" method="post" path="/v1/subscribers/{subscriberId}/messages/mark-as" -->
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

    res, err := s.Subscribers.Messages.MarkAllAs(ctx, "<id>", components.MessageMarkAsRequestDto{
        MessageID: components.CreateMessageIDArrayOfStr(
            []string{},
        ),
        MarkAs: components.MessageMarkAsRequestDtoMarkAsSeen,
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.MessageResponseDtos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `subscriberID`                                                                           | *string*                                                                                 | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `messageMarkAsRequestDto`                                                                | [components.MessageMarkAsRequestDto](../../models/components/messagemarkasrequestdto.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `idempotencyKey`                                                                         | **string*                                                                                | :heavy_minus_sign:                                                                       | A header for idempotency purposes                                                        |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.SubscribersV1ControllerMarkMessagesAsResponse](../../models/operations/subscribersv1controllermarkmessagesasresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |