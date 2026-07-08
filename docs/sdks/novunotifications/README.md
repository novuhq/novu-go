# Subscribers.Notifications

## Overview

### Available Operations

* [List](#list) - Retrieve subscriber notifications
* [Delete](#delete) - Delete a notification
* [CompleteAction](#completeaction) - Complete a notification action
* [RevertAction](#revertaction) - Revert a notification action
* [Archive](#archive) - Archive a notification
* [MarkAsRead](#markasread) - Mark a notification as read
* [Snooze](#snooze) - Snooze a notification
* [Unarchive](#unarchive) - Unarchive a notification
* [MarkAsUnread](#markasunread) - Mark a notification as unread
* [Unsnooze](#unsnooze) - Unsnooze a notification
* [ArchiveAll](#archiveall) - Archive all notifications
* [Count](#count) - Retrieve subscriber notifications count
* [DeleteAll](#deleteall) - Delete all notifications
* [MarkAllAsRead](#markallasread) - Mark all notifications as read
* [ArchiveAllRead](#archiveallread) - Archive all read notifications
* [MarkAsSeen](#markasseen) - Mark notifications as seen
* [~~Feed~~](#feed) - Retrieve subscriber notifications :warning: **Deprecated**
* [~~UnseenCount~~](#unseencount) - Retrieve unseen notifications count :warning: **Deprecated**

## List

Retrieve in-app (inbox) notifications for a subscriber by its unique key identifier **subscriberId**. 
    Supports filtering by tags, read/archived/snoozed/seen state, data attributes, severity, date range, and context keys.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_getSubscriberNotifications" method="get" path="/v2/subscribers/{subscriberId}/notifications" -->
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

    res, err := s.Subscribers.Notifications.List(ctx, operations.SubscribersControllerGetSubscriberNotificationsRequest{
        SubscriberID: "<id>",
        Offset: v3.Pointer[float64](0.0),
        CreatedGte: v3.Pointer[float64](1704067200000.0),
        CreatedLte: v3.Pointer[float64](1735689599999.0),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSubscriberNotificationsResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.SubscribersControllerGetSubscriberNotificationsRequest](../../models/operations/subscriberscontrollergetsubscribernotificationsrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.SubscribersControllerGetSubscriberNotificationsResponse](../../models/operations/subscriberscontrollergetsubscribernotificationsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Delete

Delete a specific in-app (inbox) notification permanently by its unique identifier **notificationId**.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_deleteNotification" method="delete" path="/v2/subscribers/{subscriberId}/notifications/{notificationId}" -->
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

    res, err := s.Subscribers.Notifications.Delete(ctx, "<id>", "<id>", nil, nil)
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
| `subscriberID`                                           | `string`                                                 | :heavy_check_mark:                                       | The identifier of the subscriber                         |
| `notificationID`                                         | `string`                                                 | :heavy_check_mark:                                       | The identifier of the notification                       |
| `contextKeys`                                            | []`string`                                               | :heavy_minus_sign:                                       | Context keys for filtering                               |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SubscribersControllerDeleteNotificationResponse](../../models/operations/subscriberscontrollerdeletenotificationresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CompleteAction

Mark a single in-app (inbox) notification's action (primary or secondary) as completed by its unique identifier **notificationId** and action type **actionType**.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_completeNotificationAction" method="patch" path="/v2/subscribers/{subscriberId}/notifications/{notificationId}/actions/{actionType}/complete" -->
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

    res, err := s.Subscribers.Notifications.CompleteAction(ctx, operations.SubscribersControllerCompleteNotificationActionRequest{
        SubscriberID: "<id>",
        NotificationID: "<id>",
        ActionType: operations.ActionTypeSecondary,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.InboxNotificationDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.SubscribersControllerCompleteNotificationActionRequest](../../models/operations/subscriberscontrollercompletenotificationactionrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.SubscribersControllerCompleteNotificationActionResponse](../../models/operations/subscriberscontrollercompletenotificationactionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RevertAction

Revert a single in-app (inbox) notification's action (primary or secondary) to pending state by its unique identifier **notificationId** and action type **actionType**.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_revertNotificationAction" method="patch" path="/v2/subscribers/{subscriberId}/notifications/{notificationId}/actions/{actionType}/revert" -->
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

    res, err := s.Subscribers.Notifications.RevertAction(ctx, operations.SubscribersControllerRevertNotificationActionRequest{
        SubscriberID: "<id>",
        NotificationID: "<id>",
        ActionType: operations.PathParamActionTypePrimary,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.InboxNotificationDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.SubscribersControllerRevertNotificationActionRequest](../../models/operations/subscriberscontrollerrevertnotificationactionrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.SubscribersControllerRevertNotificationActionResponse](../../models/operations/subscriberscontrollerrevertnotificationactionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Archive

Archive a specific in-app (inbox) notification by its unique identifier **notificationId**.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_archiveNotification" method="patch" path="/v2/subscribers/{subscriberId}/notifications/{notificationId}/archive" -->
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

    res, err := s.Subscribers.Notifications.Archive(ctx, "<id>", "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.InboxNotificationDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `subscriberID`                                           | `string`                                                 | :heavy_check_mark:                                       | The identifier of the subscriber                         |
| `notificationID`                                         | `string`                                                 | :heavy_check_mark:                                       | The identifier of the notification                       |
| `contextKeys`                                            | []`string`                                               | :heavy_minus_sign:                                       | Context keys for filtering                               |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SubscribersControllerArchiveNotificationResponse](../../models/operations/subscriberscontrollerarchivenotificationresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## MarkAsRead

Mark a specific in-app (inbox) notification as read by its unique identifier **notificationId**.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_markNotificationAsRead" method="patch" path="/v2/subscribers/{subscriberId}/notifications/{notificationId}/read" -->
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

    res, err := s.Subscribers.Notifications.MarkAsRead(ctx, "<id>", "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.InboxNotificationDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `subscriberID`                                           | `string`                                                 | :heavy_check_mark:                                       | The identifier of the subscriber                         |
| `notificationID`                                         | `string`                                                 | :heavy_check_mark:                                       | The identifier of the notification                       |
| `contextKeys`                                            | []`string`                                               | :heavy_minus_sign:                                       | Context keys for filtering                               |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SubscribersControllerMarkNotificationAsReadResponse](../../models/operations/subscriberscontrollermarknotificationasreadresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Snooze

Snooze a specific in-app (inbox) notification by its unique identifier **notificationId** until a specified time.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_snoozeNotification" method="patch" path="/v2/subscribers/{subscriberId}/notifications/{notificationId}/snooze" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"github.com/novuhq/novu-go/v3/types"
	"github.com/novuhq/novu-go/v3/models/components"
	"github.com/novuhq/novu-go/v3/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Subscribers.Notifications.Snooze(ctx, operations.SubscribersControllerSnoozeNotificationRequest{
        SubscriberID: "<id>",
        NotificationID: "<id>",
        SnoozeSubscriberNotificationDto: components.SnoozeSubscriberNotificationDto{
            SnoozeUntil: types.MustTimeFromString("2026-03-01T10:00:00Z"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.InboxNotificationDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.SubscribersControllerSnoozeNotificationRequest](../../models/operations/subscriberscontrollersnoozenotificationrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.SubscribersControllerSnoozeNotificationResponse](../../models/operations/subscriberscontrollersnoozenotificationresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Unarchive

Unarchive a specific in-app (inbox) notification by its unique identifier **notificationId**.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_unarchiveNotification" method="patch" path="/v2/subscribers/{subscriberId}/notifications/{notificationId}/unarchive" -->
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

    res, err := s.Subscribers.Notifications.Unarchive(ctx, "<id>", "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.InboxNotificationDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `subscriberID`                                           | `string`                                                 | :heavy_check_mark:                                       | The identifier of the subscriber                         |
| `notificationID`                                         | `string`                                                 | :heavy_check_mark:                                       | The identifier of the notification                       |
| `contextKeys`                                            | []`string`                                               | :heavy_minus_sign:                                       | Context keys for filtering                               |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SubscribersControllerUnarchiveNotificationResponse](../../models/operations/subscriberscontrollerunarchivenotificationresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## MarkAsUnread

Mark a specific in-app (inbox) notification as unread by its unique identifier **notificationId**.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_markNotificationAsUnread" method="patch" path="/v2/subscribers/{subscriberId}/notifications/{notificationId}/unread" -->
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

    res, err := s.Subscribers.Notifications.MarkAsUnread(ctx, "<id>", "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.InboxNotificationDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `subscriberID`                                           | `string`                                                 | :heavy_check_mark:                                       | The identifier of the subscriber                         |
| `notificationID`                                         | `string`                                                 | :heavy_check_mark:                                       | The identifier of the notification                       |
| `contextKeys`                                            | []`string`                                               | :heavy_minus_sign:                                       | Context keys for filtering                               |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SubscribersControllerMarkNotificationAsUnreadResponse](../../models/operations/subscriberscontrollermarknotificationasunreadresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Unsnooze

Unsnooze a specific in-app (inbox) notification by its unique identifier **notificationId**.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_unsnoozeNotification" method="patch" path="/v2/subscribers/{subscriberId}/notifications/{notificationId}/unsnooze" -->
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

    res, err := s.Subscribers.Notifications.Unsnooze(ctx, "<id>", "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.InboxNotificationDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `subscriberID`                                           | `string`                                                 | :heavy_check_mark:                                       | The identifier of the subscriber                         |
| `notificationID`                                         | `string`                                                 | :heavy_check_mark:                                       | The identifier of the notification                       |
| `contextKeys`                                            | []`string`                                               | :heavy_minus_sign:                                       | Context keys for filtering                               |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SubscribersControllerUnsnoozeNotificationResponse](../../models/operations/subscriberscontrollerunsnoozenotificationresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ArchiveAll

Archive all in-app (inbox) notifications matching the specified filters. Supports context-based filtering.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_archiveAllNotifications" method="post" path="/v2/subscribers/{subscriberId}/notifications/archive" -->
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

    res, err := s.Subscribers.Notifications.ArchiveAll(ctx, "<id>", components.UpdateAllSubscriberNotificationsDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `subscriberID`                                                                                                   | `string`                                                                                                         | :heavy_check_mark:                                                                                               | The identifier of the subscriber                                                                                 |
| `updateAllSubscriberNotificationsDto`                                                                            | [components.UpdateAllSubscriberNotificationsDto](../../models/components/updateallsubscribernotificationsdto.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `idempotencyKey`                                                                                                 | `*string`                                                                                                        | :heavy_minus_sign:                                                                                               | A header for idempotency purposes                                                                                |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.SubscribersControllerArchiveAllNotificationsResponse](../../models/operations/subscriberscontrollerarchiveallnotificationsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Count

Retrieve count of in-app (inbox) notifications for a subscriber by its unique key identifier **subscriberId**. 
    Supports multiple filters to count in-app (inbox) notifications by different criteria, including context keys.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_getSubscriberNotificationsCount" method="get" path="/v2/subscribers/{subscriberId}/notifications/count" -->
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

    res, err := s.Subscribers.Notifications.Count(ctx, "<id>", "[{\"read\":false,\"archived\":false},{\"tags\":[\"important\"]},{\"tags\":{\"and\":[{\"or\":[\"a\",\"b\"]},{\"or\":[\"c\"]}]}}]", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSubscriberNotificationsCountResponseDtos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               | Example                                                                                                   |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |                                                                                                           |
| `subscriberID`                                                                                            | `string`                                                                                                  | :heavy_check_mark:                                                                                        | The identifier of the subscriber                                                                          |                                                                                                           |
| `filters`                                                                                                 | `string`                                                                                                  | :heavy_check_mark:                                                                                        | Array of filter objects (max 30) to count notifications by different criteria                             | [{"read":false,"archived":false},{"tags":["important"]},{"tags":{"and":[{"or":["a","b"]},{"or":["c"]}]}}] |
| `idempotencyKey`                                                                                          | `*string`                                                                                                 | :heavy_minus_sign:                                                                                        | A header for idempotency purposes                                                                         |                                                                                                           |
| `opts`                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                  | :heavy_minus_sign:                                                                                        | The options for this request.                                                                             |                                                                                                           |

### Response

**[*operations.SubscribersControllerGetSubscriberNotificationsCountResponse](../../models/operations/subscriberscontrollergetsubscribernotificationscountresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DeleteAll

Permanently delete all in-app (inbox) notifications matching the specified filters. Supports context-based filtering.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_deleteAllNotifications" method="post" path="/v2/subscribers/{subscriberId}/notifications/delete" -->
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

    res, err := s.Subscribers.Notifications.DeleteAll(ctx, "<id>", components.UpdateAllSubscriberNotificationsDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `subscriberID`                                                                                                   | `string`                                                                                                         | :heavy_check_mark:                                                                                               | The identifier of the subscriber                                                                                 |
| `updateAllSubscriberNotificationsDto`                                                                            | [components.UpdateAllSubscriberNotificationsDto](../../models/components/updateallsubscribernotificationsdto.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `idempotencyKey`                                                                                                 | `*string`                                                                                                        | :heavy_minus_sign:                                                                                               | A header for idempotency purposes                                                                                |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.SubscribersControllerDeleteAllNotificationsResponse](../../models/operations/subscriberscontrollerdeleteallnotificationsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## MarkAllAsRead

Mark all in-app (inbox) notifications matching the specified filters as read. Supports context-based filtering.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_markAllNotificationsAsRead" method="post" path="/v2/subscribers/{subscriberId}/notifications/read" -->
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

    res, err := s.Subscribers.Notifications.MarkAllAsRead(ctx, "<id>", components.UpdateAllSubscriberNotificationsDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `subscriberID`                                                                                                   | `string`                                                                                                         | :heavy_check_mark:                                                                                               | The identifier of the subscriber                                                                                 |
| `updateAllSubscriberNotificationsDto`                                                                            | [components.UpdateAllSubscriberNotificationsDto](../../models/components/updateallsubscribernotificationsdto.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `idempotencyKey`                                                                                                 | `*string`                                                                                                        | :heavy_minus_sign:                                                                                               | A header for idempotency purposes                                                                                |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.SubscribersControllerMarkAllNotificationsAsReadResponse](../../models/operations/subscriberscontrollermarkallnotificationsasreadresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ArchiveAllRead

Archive all read in-app (inbox) notifications matching the specified filters. Supports context-based filtering.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_archiveAllReadNotifications" method="post" path="/v2/subscribers/{subscriberId}/notifications/read-archive" -->
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

    res, err := s.Subscribers.Notifications.ArchiveAllRead(ctx, "<id>", components.UpdateAllSubscriberNotificationsDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `subscriberID`                                                                                                   | `string`                                                                                                         | :heavy_check_mark:                                                                                               | The identifier of the subscriber                                                                                 |
| `updateAllSubscriberNotificationsDto`                                                                            | [components.UpdateAllSubscriberNotificationsDto](../../models/components/updateallsubscribernotificationsdto.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `idempotencyKey`                                                                                                 | `*string`                                                                                                        | :heavy_minus_sign:                                                                                               | A header for idempotency purposes                                                                                |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.SubscribersControllerArchiveAllReadNotificationsResponse](../../models/operations/subscriberscontrollerarchiveallreadnotificationsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## MarkAsSeen

Mark specific and multiple in-app (inbox) notifications as seen. Supports context-based filtering.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_markNotificationsAsSeen" method="post" path="/v2/subscribers/{subscriberId}/notifications/seen" -->
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

    res, err := s.Subscribers.Notifications.MarkAsSeen(ctx, "<id>", components.MarkSubscriberNotificationsAsSeenDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `subscriberID`                                                                                                     | `string`                                                                                                           | :heavy_check_mark:                                                                                                 | The identifier of the subscriber                                                                                   |
| `markSubscriberNotificationsAsSeenDto`                                                                             | [components.MarkSubscriberNotificationsAsSeenDto](../../models/components/marksubscribernotificationsasseendto.md) | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `idempotencyKey`                                                                                                   | `*string`                                                                                                          | :heavy_minus_sign:                                                                                                 | A header for idempotency purposes                                                                                  |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.SubscribersControllerMarkNotificationsAsSeenResponse](../../models/operations/subscriberscontrollermarknotificationsasseenresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ~~Feed~~

This API is deprecated, use v2 API instead. Retrieve subscriber in-app notifications by its unique key identifier **subscriberId**.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersV1Controller_getNotificationsFeed" method="get" path="/v1/subscribers/{subscriberId}/notifications/feed" -->
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

    res, err := s.Subscribers.Notifications.Feed(ctx, operations.SubscribersV1ControllerGetNotificationsFeedRequest{
        SubscriberID: "<id>",
        Page: v3.Pointer[float64](0.0),
        Payload: v3.Pointer("btoa(JSON.stringify({ foo: 123 })) results in base64 encoded string like eyJmb28iOjEyM30="),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FeedResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.SubscribersV1ControllerGetNotificationsFeedRequest](../../models/operations/subscribersv1controllergetnotificationsfeedrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.SubscribersV1ControllerGetNotificationsFeedResponse](../../models/operations/subscribersv1controllergetnotificationsfeedresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ~~UnseenCount~~

This API is deprecated, use v2 API instead. Retrieve unseen in-app notifications count for a subscriber by its unique key identifier **subscriberId**.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersV1Controller_getUnseenCount" method="get" path="/v1/subscribers/{subscriberId}/notifications/unseen" -->
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

    res, err := s.Subscribers.Notifications.UnseenCount(ctx, "<id>", v3.Pointer(false), v3.Pointer[float64](100.0), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnseenCountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `subscriberID`                                           | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `seen`                                                   | `*bool`                                                  | :heavy_minus_sign:                                       | Indicates whether to count seen notifications.           |
| `limit`                                                  | `*float64`                                               | :heavy_minus_sign:                                       | The maximum number of notifications to return.           |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SubscribersV1ControllerGetUnseenCountResponse](../../models/operations/subscribersv1controllergetunseencountresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |