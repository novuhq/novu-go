# Topics

## Overview

Topics are a way to group subscribers together so that they can be notified of events at once. A topic is identified by a custom key. This can be helpful for things like sending out marketing emails or notifying users of new features. Topics can also be used to send notifications to the subscribers who have been grouped together based on their interests, location, activities and much more.
<https://docs.novu.co/subscribers/topics>

### Available Operations

* [List](#list) - List all topics
* [Create](#create) - Create a topic
* [Get](#get) - Retrieve a topic
* [Update](#update) - Update a topic
* [Delete](#delete) - Delete a topic

## List

This api returns a paginated list of topics.
    Topics can be filtered by **key**, **name**, or **includeCursor** to paginate through the list. 
    Checkout all available filters in the query section.

### Example Usage

<!-- UsageSnippet language="go" operationID="TopicsController_listTopics" method="get" path="/v2/topics" -->
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

    res, err := s.Topics.List(ctx, operations.TopicsControllerListTopicsRequest{
        Limit: v3.Pointer[float64](10),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListTopicsResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.TopicsControllerListTopicsRequest](../../models/operations/topicscontrollerlisttopicsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.TopicsControllerListTopicsResponse](../../models/operations/topicscontrollerlisttopicsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Create

Creates a new topic if it does not exist, or updates an existing topic if it already exists. Use ?failIfExists=true to prevent updates.

### Example Usage

<!-- UsageSnippet language="go" operationID="TopicsController_upsertTopic" method="post" path="/v2/topics" -->
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

    res, err := s.Topics.Create(ctx, components.CreateUpdateTopicRequestDto{
        Key: "task:12345",
        Name: v3.Pointer("Task Title"),
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TopicResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `createUpdateTopicRequestDto`                                                                    | [components.CreateUpdateTopicRequestDto](../../models/components/createupdatetopicrequestdto.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `failIfExists`                                                                                   | **bool*                                                                                          | :heavy_minus_sign:                                                                               | If true, the request will fail if a topic with the same key already exists                       |
| `idempotencyKey`                                                                                 | **string*                                                                                        | :heavy_minus_sign:                                                                               | A header for idempotency purposes                                                                |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.TopicsControllerUpsertTopicResponse](../../models/operations/topicscontrollerupserttopicresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.TopicResponseDto        | 409                               | application/json                  |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 404, 405, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Get

Retrieve a topic by its unique key identifier **topicKey**

### Example Usage

<!-- UsageSnippet language="go" operationID="TopicsController_getTopic" method="get" path="/v2/topics/{topicKey}" -->
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

    res, err := s.Topics.Get(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TopicResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `topicKey`                                               | *string*                                                 | :heavy_check_mark:                                       | The key identifier of the topic                          |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TopicsControllerGetTopicResponse](../../models/operations/topicscontrollergettopicresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Update

Update a topic name by its unique key identifier **topicKey**

### Example Usage

<!-- UsageSnippet language="go" operationID="TopicsController_updateTopic" method="patch" path="/v2/topics/{topicKey}" -->
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

    res, err := s.Topics.Update(ctx, "<value>", components.UpdateTopicRequestDto{
        Name: "Updated Topic Name",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TopicResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `topicKey`                                                                           | *string*                                                                             | :heavy_check_mark:                                                                   | The key identifier of the topic                                                      |
| `updateTopicRequestDto`                                                              | [components.UpdateTopicRequestDto](../../models/components/updatetopicrequestdto.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `idempotencyKey`                                                                     | **string*                                                                            | :heavy_minus_sign:                                                                   | A header for idempotency purposes                                                    |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.TopicsControllerUpdateTopicResponse](../../models/operations/topicscontrollerupdatetopicresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Delete

Delete a topic by its unique key identifier **topicKey**. 
    This action is irreversible and will remove all subscriptions to the topic.

### Example Usage

<!-- UsageSnippet language="go" operationID="TopicsController_deleteTopic" method="delete" path="/v2/topics/{topicKey}" -->
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

    res, err := s.Topics.Delete(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteTopicResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `topicKey`                                               | *string*                                                 | :heavy_check_mark:                                       | The key identifier of the topic                          |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TopicsControllerDeleteTopicResponse](../../models/operations/topicscontrollerdeletetopicresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |