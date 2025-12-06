# Novu SDK

## Overview

Novu API: Novu REST API. Please see https://docs.novu.co/api-reference for more details.

Novu Documentation
<https://docs.novu.co>

### Available Operations

* [Trigger](#trigger) - Trigger event
* [Cancel](#cancel) - Cancel triggered event
* [TriggerBroadcast](#triggerbroadcast) - Broadcast event to all
* [TriggerBulk](#triggerbulk) - Bulk trigger event

## Trigger


    Trigger event is the main (and only) way to send notifications to subscribers. The trigger identifier is used to match the particular workflow associated with it. Additional information can be passed according the body interface below.
    To prevent duplicate triggers, you can optionally pass a **transactionId** in the request body. If the same **transactionId** is used again, the trigger will be ignored. The retention period depends on your billing tier.

### Example Usage

<!-- UsageSnippet language="go" operationID="EventsController_trigger" method="post" path="/v1/events/trigger" -->
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

    res, err := s.Trigger(ctx, components.TriggerEventRequestDto{
        WorkflowID: "workflow_identifier",
        Payload: map[string]any{
            "comment_id": "string",
            "post": map[string]any{
                "text": "string",
            },
        },
        Overrides: &components.Overrides{},
        To: components.CreateToStr(
            "SUBSCRIBER_ID",
        ),
        Actor: v3.Pointer(components.CreateActorStr(
            "<value>",
        )),
        Context: map[string]components.Context{
            "key": components.CreateContextStr(
                "org-acme",
            ),
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TriggerEventResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `triggerEventRequestDto`                                                               | [components.TriggerEventRequestDto](../../models/components/triggereventrequestdto.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `idempotencyKey`                                                                       | **string*                                                                              | :heavy_minus_sign:                                                                     | A header for idempotency purposes                                                      |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.EventsControllerTriggerResponse](../../models/operations/eventscontrollertriggerresponse.md), error**

### Errors

| Error Type                              | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| apierrors.PayloadValidationExceptionDto | 400                                     | application/json                        |
| apierrors.ErrorDto                      | 414                                     | application/json                        |
| apierrors.ErrorDto                      | 401, 403, 404, 405, 409, 413, 415       | application/json                        |
| apierrors.ValidationErrorDto            | 422                                     | application/json                        |
| apierrors.ErrorDto                      | 500                                     | application/json                        |
| apierrors.APIError                      | 4XX, 5XX                                | \*/\*                                   |

## Cancel


    Using a previously generated transactionId during the event trigger,
     will cancel any active or pending workflows. This is useful to cancel active digests, delays etc...
    

### Example Usage

<!-- UsageSnippet language="go" operationID="EventsController_cancel" method="delete" path="/v1/events/trigger/{transactionId}" -->
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

    res, err := s.Cancel(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Boolean != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `transactionID`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.EventsControllerCancelResponse](../../models/operations/eventscontrollercancelresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## TriggerBroadcast

Trigger a broadcast event to all existing subscribers, could be used to send announcements, etc.
      In the future could be used to trigger events to a subset of subscribers based on defined filters.

### Example Usage

<!-- UsageSnippet language="go" operationID="EventsController_broadcastEventToAll" method="post" path="/v1/events/trigger/broadcast" -->
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

    res, err := s.TriggerBroadcast(ctx, components.TriggerEventToAllRequestDto{
        Name: "<value>",
        Payload: map[string]any{
            "comment_id": "string",
            "post": map[string]any{
                "text": "string",
            },
        },
        Overrides: &components.TriggerEventToAllRequestDtoOverrides{
            AdditionalProperties: map[string]map[string]any{
                "fcm": map[string]any{
                    "data": map[string]any{
                        "key": "value",
                    },
                },
            },
        },
        Actor: v3.Pointer(components.CreateTriggerEventToAllRequestDtoActorSubscriberPayloadDto(
            components.SubscriberPayloadDto{
                FirstName: v3.Pointer("John"),
                LastName: v3.Pointer("Doe"),
                Email: v3.Pointer("john.doe@example.com"),
                Phone: v3.Pointer("+1234567890"),
                Avatar: v3.Pointer("https://example.com/avatar.jpg"),
                Locale: v3.Pointer("en-US"),
                Timezone: v3.Pointer("America/New_York"),
                SubscriberID: "<id>",
            },
        )),
        Context: map[string]components.TriggerEventToAllRequestDtoContext{
            "key": components.CreateTriggerEventToAllRequestDtoContextStr(
                "org-acme",
            ),
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TriggerEventResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `triggerEventToAllRequestDto`                                                                    | [components.TriggerEventToAllRequestDto](../../models/components/triggereventtoallrequestdto.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `idempotencyKey`                                                                                 | **string*                                                                                        | :heavy_minus_sign:                                                                               | A header for idempotency purposes                                                                |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.EventsControllerBroadcastEventToAllResponse](../../models/operations/eventscontrollerbroadcasteventtoallresponse.md), error**

### Errors

| Error Type                              | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| apierrors.PayloadValidationExceptionDto | 400                                     | application/json                        |
| apierrors.ErrorDto                      | 414                                     | application/json                        |
| apierrors.ErrorDto                      | 401, 403, 404, 405, 409, 413, 415       | application/json                        |
| apierrors.ValidationErrorDto            | 422                                     | application/json                        |
| apierrors.ErrorDto                      | 500                                     | application/json                        |
| apierrors.APIError                      | 4XX, 5XX                                | \*/\*                                   |

## TriggerBulk


      Using this endpoint you can trigger multiple events at once, to avoid multiple calls to the API.
      The bulk API is limited to 100 events per request.
    

### Example Usage

<!-- UsageSnippet language="go" operationID="EventsController_triggerBulk" method="post" path="/v1/events/trigger/bulk" -->
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

    res, err := s.TriggerBulk(ctx, components.BulkTriggerEventDto{
        Events: []components.TriggerEventRequestDto{
            components.TriggerEventRequestDto{
                WorkflowID: "workflow_identifier",
                Payload: map[string]any{
                    "comment_id": "string",
                    "post": map[string]any{
                        "text": "string",
                    },
                },
                Overrides: &components.Overrides{},
                To: components.CreateToStr(
                    "SUBSCRIBER_ID",
                ),
            },
            components.TriggerEventRequestDto{
                WorkflowID: "workflow_identifier",
                Payload: map[string]any{
                    "comment_id": "string",
                    "post": map[string]any{
                        "text": "string",
                    },
                },
                Overrides: &components.Overrides{},
                To: components.CreateToStr(
                    "SUBSCRIBER_ID",
                ),
            },
            components.TriggerEventRequestDto{
                WorkflowID: "workflow_identifier",
                Payload: map[string]any{
                    "comment_id": "string",
                    "post": map[string]any{
                        "text": "string",
                    },
                },
                Overrides: &components.Overrides{},
                To: components.CreateToStr(
                    "SUBSCRIBER_ID",
                ),
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TriggerEventResponseDtos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `bulkTriggerEventDto`                                                            | [components.BulkTriggerEventDto](../../models/components/bulktriggereventdto.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `idempotencyKey`                                                                 | **string*                                                                        | :heavy_minus_sign:                                                               | A header for idempotency purposes                                                |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.EventsControllerTriggerBulkResponse](../../models/operations/eventscontrollertriggerbulkresponse.md), error**

### Errors

| Error Type                              | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| apierrors.PayloadValidationExceptionDto | 400                                     | application/json                        |
| apierrors.ErrorDto                      | 414                                     | application/json                        |
| apierrors.ErrorDto                      | 401, 403, 404, 405, 409, 413, 415       | application/json                        |
| apierrors.ValidationErrorDto            | 422                                     | application/json                        |
| apierrors.ErrorDto                      | 500                                     | application/json                        |
| apierrors.APIError                      | 4XX, 5XX                                | \*/\*                                   |