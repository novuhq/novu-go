# Novu SDK

## Overview

Novu API: Novu REST API. Please see https://docs.novu.co/api-reference for more details.

Novu Documentation
<https://docs.novu.co>

### Available Operations

* [Trigger](#trigger) - Trigger event
* [TriggerBulk](#triggerbulk) - Bulk trigger event
* [Broadcast](#broadcast) - Broadcast event to all
* [Cancel](#cancel) - Cancel triggered event

## Trigger


    Trigger event is the main (and only) way to send notifications to subscribers. 
    The trigger identifier is used to match the particular workflow associated with it. 
    Additional information can be passed according the body interface below.
    

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
        novugo.WithSecurity(os.Getenv("NOVU_API_KEY")),
    )

    res, err := s.Trigger(ctx, components.TriggerEventRequestDto{
        Name: "workflow_identifier",
        Payload: map[string]any{
            "comment_id": "string",
            "post": map[string]any{
                "text": "string",
            },
        },
        BridgeURL: novugo.String("https://example.com/bridge"),
        Overrides: map[string]map[string]any{
            "fcm": map[string]any{
                "data": map[string]any{
                    "key": "value",
                },
            },
        },
        To: components.CreateToSubscriberPayloadDto(
            components.SubscriberPayloadDto{
                SubscriberID: "<id>",
            },
        ),
    })
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
| `request`                                                                              | [components.TriggerEventRequestDto](../../models/components/triggereventrequestdto.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.EventsControllerTriggerResponse](../../models/operations/eventscontrollertriggerresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 404, 409                | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## TriggerBulk


      Using this endpoint you can trigger multiple events at once, to avoid multiple calls to the API.
      The bulk API is limited to 100 events per request.
    

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
        novugo.WithSecurity(os.Getenv("NOVU_API_KEY")),
    )

    res, err := s.TriggerBulk(ctx, components.BulkTriggerEventDto{
        Events: []components.TriggerEventRequestDto{
            components.TriggerEventRequestDto{
                Name: "workflow_identifier",
                Payload: map[string]any{
                    "comment_id": "string",
                    "post": map[string]any{
                        "text": "string",
                    },
                },
                BridgeURL: novugo.String("https://example.com/bridge"),
                Overrides: map[string]map[string]any{
                    "fcm": map[string]any{
                        "data": map[string]any{
                            "key": "value",
                        },
                    },
                },
                To: components.CreateToSubscriberPayloadDto(
                    components.SubscriberPayloadDto{
                        SubscriberID: "<id>",
                    },
                ),
            },
            components.TriggerEventRequestDto{
                Name: "workflow_identifier",
                Payload: map[string]any{
                    "comment_id": "string",
                    "post": map[string]any{
                        "text": "string",
                    },
                },
                BridgeURL: novugo.String("https://example.com/bridge"),
                Overrides: map[string]map[string]any{
                    "fcm": map[string]any{
                        "data": map[string]any{
                            "key": "value",
                        },
                    },
                },
                To: components.CreateToArrayOf1(
                    []components.One{
                        components.CreateOneTopicPayloadDto(
                            components.TopicPayloadDto{
                                TopicKey: "<value>",
                                Type: components.TriggerRecipientsTypeEnumSubscriber,
                            },
                        ),
                    },
                ),
            },
            components.TriggerEventRequestDto{
                Name: "workflow_identifier",
                Payload: map[string]any{
                    "comment_id": "string",
                    "post": map[string]any{
                        "text": "string",
                    },
                },
                BridgeURL: novugo.String("https://example.com/bridge"),
                Overrides: map[string]map[string]any{
                    "fcm": map[string]any{
                        "data": map[string]any{
                            "key": "value",
                        },
                    },
                },
                To: components.CreateToArrayOf1(
                    []components.One{
                        components.CreateOneStr(
                            "SUBSCRIBER_ID",
                        ),
                        components.CreateOneStr(
                            "SUBSCRIBER_ID",
                        ),
                    },
                ),
            },
        },
    })
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
| `request`                                                                        | [components.BulkTriggerEventDto](../../models/components/bulktriggereventdto.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.EventsControllerTriggerBulkResponse](../../models/operations/eventscontrollertriggerbulkresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 404, 409                | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## Broadcast

Trigger a broadcast event to all existing subscribers, could be used to send announcements, etc.




      In the future could be used to trigger events to a subset of subscribers based on defined filters.

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
        novugo.WithSecurity(os.Getenv("NOVU_API_KEY")),
    )

    res, err := s.Broadcast(ctx, components.TriggerEventToAllRequestDto{
        Name: "<value>",
        Payload: map[string]any{
            "comment_id": "string",
            "post": map[string]any{
                "text": "string",
            },
        },
    })
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
| `request`                                                                                        | [components.TriggerEventToAllRequestDto](../../models/components/triggereventtoallrequestdto.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.EventsControllerBroadcastEventToAllResponse](../../models/operations/eventscontrollerbroadcasteventtoallresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 404, 409                | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## Cancel


    Using a previously generated transactionId during the event trigger,
     will cancel any active or pending workflows. This is useful to cancel active digests, delays etc...
    

### Example Usage

```go
package main

import(
	"context"
	"os"
	novugo "github.com/novuhq/novu-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := novugo.New(
        novugo.WithSecurity(os.Getenv("NOVU_API_KEY")),
    )

    res, err := s.Cancel(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DataBooleanDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `transactionID`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.EventsControllerCancelResponse](../../models/operations/eventscontrollercancelresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 404, 409                | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |