# Subscribers.Topics

## Overview

### Available Operations

* [List](#list) - Retrieve subscriber subscriptions

## List

Retrieve subscriber's topic subscriptions by its unique key identifier **subscriberId**. 
    Checkout all available filters in the query section.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_listSubscriberTopics" method="get" path="/v2/subscribers/{subscriberId}/subscriptions" -->
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

    res, err := s.Subscribers.Topics.List(ctx, operations.SubscribersControllerListSubscriberTopicsRequest{
        SubscriberID: "<id>",
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

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.SubscribersControllerListSubscriberTopicsRequest](../../models/operations/subscriberscontrollerlistsubscribertopicsrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |
| `opts`                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.SubscribersControllerListSubscriberTopicsResponse](../../models/operations/subscriberscontrollerlistsubscribertopicsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |