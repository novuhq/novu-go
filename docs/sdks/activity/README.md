# Activity

## Overview

### Available Operations

* [Track](#track) - Track activity and engagement events

## Track

Track activity and engagement events for a specific delivery provider

### Example Usage

<!-- UsageSnippet language="go" operationID="InboundWebhooksController_handleWebhook" method="post" path="/v2/inbound-webhooks/delivery-providers/{environmentId}/{integrationId}" -->
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

    res, err := s.Activity.Track(ctx, "<id>", "<id>", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookResultDtos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `environmentID`                                          | *string*                                                 | :heavy_check_mark:                                       | The environment identifier                               |
| `integrationID`                                          | *string*                                                 | :heavy_check_mark:                                       | The integration identifier for the delivery provider     |
| `requestBody`                                            | *any*                                                    | :heavy_check_mark:                                       | Webhook event payload from the delivery provider         |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.InboundWebhooksControllerHandleWebhookResponse](../../models/operations/inboundwebhookscontrollerhandlewebhookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |