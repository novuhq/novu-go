# NovuSubscribers
(*Topics.Subscribers*)

## Overview

### Available Operations

* [Retrieve](#retrieve) - Check topic subscriber

## Retrieve

Check if a subscriber belongs to a certain topic

### Example Usage

```go
package main

import(
	"context"
	novugo "github.com/novuhq/novu-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := novugo.New(
        novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Topics.Subscribers.Retrieve(ctx, "<id>", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TopicSubscriberDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `externalSubscriberID`                                   | *string*                                                 | :heavy_check_mark:                                       | The external subscriber id                               |
| `topicKey`                                               | *string*                                                 | :heavy_check_mark:                                       | The topic key                                            |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TopicsV1ControllerGetTopicSubscriberResponse](../../models/operations/topicsv1controllergettopicsubscriberresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |