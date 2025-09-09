# Subscriptions
(*Topics.Subscriptions*)

## Overview

### Available Operations

* [List](#list) - List topic subscriptions
* [Create](#create) - Create topic subscriptions
* [Delete](#delete) - Delete topic subscriptions

## List

List all subscriptions of subscribers for a topic.
    Checkout all available filters in the query section.

### Example Usage

<!-- UsageSnippet language="go" operationID="TopicsController_listTopicSubscriptions" method="get" path="/v2/topics/{topicKey}/subscriptions" -->
```go
package main

import(
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := novugo.New(
        novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Topics.Subscriptions.List(ctx, operations.TopicsControllerListTopicSubscriptionsRequest{
        TopicKey: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListTopicSubscriptionsResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.TopicsControllerListTopicSubscriptionsRequest](../../models/operations/topicscontrollerlisttopicsubscriptionsrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.TopicsControllerListTopicSubscriptionsResponse](../../models/operations/topicscontrollerlisttopicsubscriptionsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Create

This api will create subscription for subscriberIds for a topic. 
      Its like subscribing to a common interest group. if topic does not exist, it will be created.

### Example Usage

<!-- UsageSnippet language="go" operationID="TopicsController_createTopicSubscriptions" method="post" path="/v2/topics/{topicKey}/subscriptions" -->
```go
package main

import(
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := novugo.New(
        novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Topics.Subscriptions.Create(ctx, "<value>", components.CreateTopicSubscriptionsRequestDto{
        SubscriberIds: []string{
            "subscriberId1",
            "subscriberId2",
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateTopicSubscriptionsResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `topicKey`                                                                                                     | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The key identifier of the topic                                                                                |
| `createTopicSubscriptionsRequestDto`                                                                           | [components.CreateTopicSubscriptionsRequestDto](../../models/components/createtopicsubscriptionsrequestdto.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `idempotencyKey`                                                                                               | **string*                                                                                                      | :heavy_minus_sign:                                                                                             | A header for idempotency purposes                                                                              |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.TopicsControllerCreateTopicSubscriptionsResponse](../../models/operations/topicscontrollercreatetopicsubscriptionsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Delete

Delete subscriptions for subscriberIds for a topic.

### Example Usage

<!-- UsageSnippet language="go" operationID="TopicsController_deleteTopicSubscriptions" method="delete" path="/v2/topics/{topicKey}/subscriptions" -->
```go
package main

import(
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := novugo.New(
        novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Topics.Subscriptions.Delete(ctx, "<value>", components.DeleteTopicSubscriptionsRequestDto{
        SubscriberIds: []string{
            "subscriberId1",
            "subscriberId2",
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteTopicSubscriptionsResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `topicKey`                                                                                                     | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The key identifier of the topic                                                                                |
| `deleteTopicSubscriptionsRequestDto`                                                                           | [components.DeleteTopicSubscriptionsRequestDto](../../models/components/deletetopicsubscriptionsrequestdto.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `idempotencyKey`                                                                                               | **string*                                                                                                      | :heavy_minus_sign:                                                                                             | A header for idempotency purposes                                                                              |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.TopicsControllerDeleteTopicSubscriptionsResponse](../../models/operations/topicscontrollerdeletetopicsubscriptionsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |