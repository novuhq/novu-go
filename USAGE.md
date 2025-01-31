<!-- Start SDK Example Usage [usage] -->
### Trigger Notification Event

```go
package main

import (
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithSecurity(os.Getenv("NOVU_SECRET_KEY")),
	)

	res, err := s.Trigger(ctx, components.TriggerEventRequestDto{
		WorkflowID: "workflow_identifier",
		Payload: map[string]any{
			"comment_id": "string",
			"post": map[string]any{
				"text": "string",
			},
		},
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
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.TriggerEventResponseDto != nil {
		// handle response
	}
}

```

### Trigger Notification Events in Bulk

```go
package main

import (
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithSecurity(os.Getenv("NOVU_SECRET_KEY")),
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
				WorkflowID: "workflow_identifier",
				Payload: map[string]any{
					"comment_id": "string",
					"post": map[string]any{
						"text": "string",
					},
				},
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
								Type:     components.TriggerRecipientsTypeEnumSubscriber,
							},
						),
					},
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
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.TriggerEventResponseDtos != nil {
		// handle response
	}
}

```

### Broadcast Event to All

```go
package main

import (
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithSecurity(os.Getenv("NOVU_SECRET_KEY")),
	)

	res, err := s.Broadcast(ctx, components.TriggerEventToAllRequestDto{
		Name: "<value>",
		Payload: map[string]any{
			"comment_id": "string",
			"post": map[string]any{
				"text": "string",
			},
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

### Cancel Triggered Event

```go
package main

import (
	"context"
	novugo "github.com/novuhq/novu-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithSecurity(os.Getenv("NOVU_SECRET_KEY")),
	)

	res, err := s.Cancel(ctx, "<id>", nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.DataBooleanDto != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->