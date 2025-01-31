# TopicsSubscribers
(*TopicsSubscribers*)

## Overview

### Available Operations

* [Add](#add) - Subscribers addition

## Add

Add subscribers to a topic by key

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
        novugo.WithSecurity(os.Getenv("NOVU_SECRET_KEY")),
    )

    res, err := s.TopicsSubscribers.Add(ctx, "<value>", components.AddSubscribersRequestDto{
        Subscribers: []string{
            "<value>",
            "<value>",
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AssignSubscriberToTopicDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `topicKey`                                                                                 | *string*                                                                                   | :heavy_check_mark:                                                                         | The topic key                                                                              |
| `addSubscribersRequestDto`                                                                 | [components.AddSubscribersRequestDto](../../models/components/addsubscribersrequestdto.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `idempotencyKey`                                                                           | **string*                                                                                  | :heavy_minus_sign:                                                                         | A header for idempotency purposes                                                          |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.TopicsControllerAssignResponse](../../models/operations/topicscontrollerassignresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |