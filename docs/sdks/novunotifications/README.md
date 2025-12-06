# Subscribers.Notifications

## Overview

### Available Operations

* [Feed](#feed) - Retrieve subscriber notifications
* [UnseenCount](#unseencount) - Retrieve unseen notifications count

## Feed

Retrieve subscriber in-app (inbox) notifications by its unique key identifier **subscriberId**.

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
        Page: v3.Pointer[float64](0),
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

## UnseenCount

Retrieve unseen in-app (inbox) notifications count for a subscriber by its unique key identifier **subscriberId**.

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

    res, err := s.Subscribers.Notifications.UnseenCount(ctx, "<id>", v3.Pointer(false), v3.Pointer[float64](100), nil)
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
| `subscriberID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `seen`                                                   | **bool*                                                  | :heavy_minus_sign:                                       | Indicates whether to count seen notifications.           |
| `limit`                                                  | **float64*                                               | :heavy_minus_sign:                                       | The maximum number of notifications to return.           |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
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