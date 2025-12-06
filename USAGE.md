<!-- Start SDK Example Usage [usage] -->
### Trigger Notification Event

```go
package main

import (
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

### Cancel Triggered Event

```go
package main

import (
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

### Broadcast Event to All

```go
package main

import (
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
				FirstName:    v3.Pointer("John"),
				LastName:     v3.Pointer("Doe"),
				Email:        v3.Pointer("john.doe@example.com"),
				Phone:        v3.Pointer("+1234567890"),
				Avatar:       v3.Pointer("https://example.com/avatar.jpg"),
				Locale:       v3.Pointer("en-US"),
				Timezone:     v3.Pointer("America/New_York"),
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

### Trigger Notification Events in Bulk

```go
package main

import (
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
<!-- End SDK Example Usage [usage] -->