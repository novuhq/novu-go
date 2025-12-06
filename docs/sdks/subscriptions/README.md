# Topics.Subscriptions

## Overview

### Available Operations

* [List](#list) - List topic subscriptions
* [Create](#create) - Create topic subscriptions
* [Delete](#delete) - Delete topic subscriptions
* [Update](#update) - Update a topic subscription

## List

List all subscriptions of subscribers for a topic.
    Checkout all available filters in the query section.

### Example Usage

<!-- UsageSnippet language="go" operationID="TopicsController_listTopicSubscriptions" method="get" path="/v2/topics/{topicKey}/subscriptions" -->
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

    res, err := s.Topics.Subscriptions.List(ctx, operations.TopicsControllerListTopicSubscriptionsRequest{
        TopicKey: "<value>",
        Limit: v3.Pointer[float64](10),
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
	"github.com/novuhq/novu-go/v3"
	"github.com/novuhq/novu-go/v3/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Topics.Subscriptions.Create(ctx, "<value>", components.CreateTopicSubscriptionsRequestDto{
        Subscriptions: []components.Subscriptions{
            components.CreateSubscriptionsTopicSubscriberIdentifierDto(
                components.TopicSubscriberIdentifierDto{
                    Identifier: "subscriber-123-subscription-a",
                    SubscriberID: "subscriber-123",
                },
            ),
            components.CreateSubscriptionsTopicSubscriberIdentifierDto(
                components.TopicSubscriberIdentifierDto{
                    Identifier: "subscriber-456-subscription-b",
                    SubscriberID: "subscriber-456",
                },
            ),
        },
        Name: v3.Pointer("My Topic"),
        Preferences: []components.Preferences{
            components.CreatePreferencesWorkflowPreferenceRequestDto(
                components.WorkflowPreferenceRequestDto{
                    Condition: map[string]any{
                        "===": []any{
                            map[string]any{
                                "var": "tier",
                            },
                            "premium",
                        },
                    },
                    WorkflowID: "workflow-123",
                },
            ),
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateSubscriptionsResponseDto != nil {
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
	"github.com/novuhq/novu-go/v3"
	"github.com/novuhq/novu-go/v3/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
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

## Update

Update a subscription by its unique identifier **subscriptionId** for a topic. You can update the preferences and name associated with the subscription.

### Example Usage

<!-- UsageSnippet language="go" operationID="TopicsController_updateTopicSubscription" method="patch" path="/v2/topics/{topicKey}/subscriptions/{subscriptionId}" -->
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

    res, err := s.Topics.Subscriptions.Update(ctx, "<value>", "<id>", components.UpdateTopicSubscriptionRequestDto{
        Name: v3.Pointer("My Subscription"),
        Preferences: []components.UpdateTopicSubscriptionRequestDtoPreferences{
            components.CreateUpdateTopicSubscriptionRequestDtoPreferencesWorkflowPreferenceRequestDto(
                components.WorkflowPreferenceRequestDto{
                    Condition: map[string]any{
                        "===": []any{
                            map[string]any{
                                "var": "tier",
                            },
                            "premium",
                        },
                    },
                    WorkflowID: "workflow-123",
                },
            ),
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `topicKey`                                                                                                   | *string*                                                                                                     | :heavy_check_mark:                                                                                           | The key identifier of the topic                                                                              |
| `subscriptionID`                                                                                             | *string*                                                                                                     | :heavy_check_mark:                                                                                           | The unique identifier of the subscription                                                                    |
| `updateTopicSubscriptionRequestDto`                                                                          | [components.UpdateTopicSubscriptionRequestDto](../../models/components/updatetopicsubscriptionrequestdto.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `idempotencyKey`                                                                                             | **string*                                                                                                    | :heavy_minus_sign:                                                                                           | A header for idempotency purposes                                                                            |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.TopicsControllerUpdateTopicSubscriptionResponse](../../models/operations/topicscontrollerupdatetopicsubscriptionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |