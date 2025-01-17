# NovuMessages
(*Subscribers.Messages*)

## Overview

### Available Operations

* [MarkAs](#markas) - Mark a subscriber messages as seen, read, unseen or unread
* [MarkAll](#markall) - Marks all the subscriber messages as read, unread, seen or unseen. Optionally you can pass feed id (or array) to mark messages of a particular feed.
* [UpdateAction](#updateaction) - Mark message action as seen

## MarkAs

Mark a subscriber messages as seen, read, unseen or unread

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

    res, err := s.Subscribers.Messages.MarkAs(ctx, "<id>", components.MessageMarkAsRequestDto{
        MessageID: components.CreateMessageIDStr(
            "<id>",
        ),
        MarkAs: components.MarkAsUnread,
    })
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
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.SubscribersControllerMarkMessagesAsResponse](../../models/operations/subscriberscontrollermarkmessagesasresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 404, 409                | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## MarkAll

Marks all the subscriber messages as read, unread, seen or unseen. Optionally you can pass feed id (or array) to mark messages of a particular feed.

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

    res, err := s.Subscribers.Messages.MarkAll(ctx, "<id>", components.MarkAllMessageAsRequestDto{
        MarkAs: components.MarkAllMessageAsRequestDtoMarkAsSeen,
    })
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
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.SubscribersControllerMarkAllUnreadAsReadResponse](../../models/operations/subscriberscontrollermarkallunreadasreadresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 404, 409                | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## UpdateAction

Mark message action as seen

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

    res, err := s.Subscribers.Messages.UpdateAction(ctx, "<id>", "<value>", "<id>", components.MarkMessageActionAsSeenDto{
        Status: components.MarkMessageActionAsSeenDtoStatusDone,
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `messageID`                                                                                    | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `type_`                                                                                        | *any*                                                                                          | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `subscriberID`                                                                                 | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `markMessageActionAsSeenDto`                                                                   | [components.MarkMessageActionAsSeenDto](../../models/components/markmessageactionasseendto.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.SubscribersControllerMarkActionAsSeenResponse](../../models/operations/subscriberscontrollermarkactionasseenresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 404, 409                | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |