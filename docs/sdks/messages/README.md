# Messages

## Overview

A message in Novu represents a notification delivered to a recipient on a particular channel. Messages contain information about the request that triggered its delivery, a view of the data sent to the recipient, and a timeline of its lifecycle events. Learn more about messages.
<https://docs.novu.co/workflows/messages>

### Available Operations

* [Retrieve](#retrieve) - List all messages
* [Delete](#delete) - Delete a message
* [DeleteByTransactionID](#deletebytransactionid) - Delete messages by transactionId

## Retrieve

List all messages for the current environment. 
    This API supports filtering by **channel**, **subscriberId**, and **transactionId**. 
    This API returns a paginated list of messages.

### Example Usage

<!-- UsageSnippet language="go" operationID="MessagesController_getMessages" method="get" path="/v1/messages" -->
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

    res, err := s.Messages.Retrieve(ctx, operations.MessagesControllerGetMessagesRequest{
        ContextKeys: []string{
            "tenant:org-123",
            "region:us-east-1",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MessagesResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.MessagesControllerGetMessagesRequest](../../models/operations/messagescontrollergetmessagesrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.MessagesControllerGetMessagesResponse](../../models/operations/messagescontrollergetmessagesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Delete

Delete a message entity from the Novu platform by **messageId**. 
    This action is irreversible. **messageId** is required and of mongodbId type.

### Example Usage

<!-- UsageSnippet language="go" operationID="MessagesController_deleteMessage" method="delete" path="/v1/messages/{messageId}" -->
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

    res, err := s.Messages.Delete(ctx, "507f1f77bcf86cd799439011", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteMessageResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `messageID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 507f1f77bcf86cd799439011                                 |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.MessagesControllerDeleteMessageResponse](../../models/operations/messagescontrollerdeletemessageresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DeleteByTransactionID

Delete multiple messages from the Novu platform using **transactionId** of triggered event. 
    This API supports filtering by **channel** and delete all messages associated with the **transactionId**.

### Example Usage

<!-- UsageSnippet language="go" operationID="MessagesController_deleteMessagesByTransactionId" method="delete" path="/v1/messages/transaction/{transactionId}" -->
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

    res, err := s.Messages.DeleteByTransactionID(ctx, "507f1f77bcf86cd799439011", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |                                                           |
| `transactionID`                                           | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       | 507f1f77bcf86cd799439011                                  |
| `channel`                                                 | [*operations.Channel](../../models/operations/channel.md) | :heavy_minus_sign:                                        | The channel of the message to be deleted                  |                                                           |
| `idempotencyKey`                                          | **string*                                                 | :heavy_minus_sign:                                        | A header for idempotency purposes                         |                                                           |
| `opts`                                                    | [][operations.Option](../../models/operations/option.md)  | :heavy_minus_sign:                                        | The options for this request.                             |                                                           |

### Response

**[*operations.MessagesControllerDeleteMessagesByTransactionIDResponse](../../models/operations/messagescontrollerdeletemessagesbytransactionidresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |