<!-- speakeasy-ignore-start -->

<div align="center">
  <a href="https://novu.co?utm_source=github" target="_blank">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://user-images.githubusercontent.com/2233092/213641039-220ac15f-f367-4d13-9eaf-56e79433b8c1.png">
    <img alt="Novu Logo" src="https://user-images.githubusercontent.com/2233092/213641043-3bbb3f21-3c53-4e67-afe5-755aeb222159.png" width="280"/>
  </picture>
  </a>
</div>

<br/>

<p align="center">
   <a href="https://www.producthunt.com/products/novu">
    <img src="https://img.shields.io/badge/Product%20Hunt-Golden%20Kitty%20Award%202023-yellow" alt="Product Hunt">
  </a>
  <a href="https://news.ycombinator.com/item?id=38419513"><img src="https://img.shields.io/badge/Hacker%20News-%231-%23FF6600" alt="Hacker News"></a>
  <a href="https://www.npmjs.com/package/@novu/node">
    <img src="https://img.shields.io/npm/dm/@novu/node" alt="npm downloads">
  </a>
</p>

<h1 align="center">The &lt;Inbox /&gt; infrastructure for modern products</h1>

<div align="center">
The notification platform that turns complex multi-channel delivery into a single <Inbox /> component. Built for developers, designed for growth, powered by open source.
</div>

<!-- speakeasy-ignore-end -->


# Novu's API v2 Go SDK

Novu's API exposes the entire Novu features via a standardized programmatic interface. Please refer to the [full documentation](https://docs.novu.co/) to learn more.

Developer-friendly & type-safe Go SDK specifically catered to leverage Novu API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/novuhq/novu-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />

<!-- Start Summary [summary] -->
## Summary

Novu API: Novu REST API. Please see https://docs.novu.co/api-reference for more details.

For more information about the API: [Novu Documentation](https://docs.novu.co)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [Novu's API v2 Go SDK](#novus-api-v2-go-sdk)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Pagination](#pagination)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/novuhq/novu-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Trigger Notification Event

```go
package main

import (
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
	)

	res, err := s.Trigger(ctx, components.TriggerEventRequestDto{
		WorkflowID: "workflow_identifier",
		Payload: map[string]any{
			"comment_id": "string",
			"post": map[string]any{
				"text": "string",
			},
		},
		Overrides: &components.Overrides{
			Steps: map[string]components.StepsOverrides{
				"email-step": components.StepsOverrides{
					Providers: map[string]map[string]any{
						"sendgrid": map[string]any{
							"templateId": "1234567890",
						},
					},
				},
			},
			Providers: map[string]map[string]any{
				"sendgrid": map[string]any{
					"templateId": "1234567890",
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

### Cancel Triggered Event

```go
package main

import (
	"context"
	novugo "github.com/novuhq/novu-go"
	"log"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
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
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
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
			Steps: map[string]components.StepsOverrides{
				"email-step": components.StepsOverrides{
					Providers: map[string]map[string]any{
						"sendgrid": map[string]any{
							"templateId": "1234567890",
						},
					},
				},
			},
			Providers: map[string]map[string]any{
				"sendgrid": map[string]any{
					"templateId": "1234567890",
				},
			},
			AdditionalProperties: map[string]map[string]any{
				"fcm": map[string]any{
					"data": map[string]any{
						"key": "value",
					},
				},
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

### Trigger Notification Events in Bulk

```go
package main

import (
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
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
				Overrides: &components.Overrides{
					Steps: map[string]components.StepsOverrides{
						"email-step": components.StepsOverrides{
							Providers: map[string]map[string]any{
								"sendgrid": map[string]any{
									"templateId": "1234567890",
								},
							},
						},
					},
					Providers: map[string]map[string]any{
						"sendgrid": map[string]any{
							"templateId": "1234567890",
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
				Overrides: &components.Overrides{
					Steps: map[string]components.StepsOverrides{
						"email-step": components.StepsOverrides{
							Providers: map[string]map[string]any{
								"sendgrid": map[string]any{
									"templateId": "1234567890",
								},
							},
						},
					},
					Providers: map[string]map[string]any{
						"sendgrid": map[string]any{
							"templateId": "1234567890",
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
				Overrides: &components.Overrides{
					Steps: map[string]components.StepsOverrides{
						"email-step": components.StepsOverrides{
							Providers: map[string]map[string]any{
								"sendgrid": map[string]any{
									"templateId": "1234567890",
								},
							},
						},
					},
					Providers: map[string]map[string]any{
						"sendgrid": map[string]any{
							"templateId": "1234567890",
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name        | Type   | Scheme  | Environment Variable |
| ----------- | ------ | ------- | -------------------- |
| `SecretKey` | apiKey | API key | `NOVU_SECRET_KEY`    |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
	)

	res, err := s.Trigger(ctx, components.TriggerEventRequestDto{
		WorkflowID: "workflow_identifier",
		Payload: map[string]any{
			"comment_id": "string",
			"post": map[string]any{
				"text": "string",
			},
		},
		Overrides: &components.Overrides{
			Steps: map[string]components.StepsOverrides{
				"email-step": components.StepsOverrides{
					Providers: map[string]map[string]any{
						"sendgrid": map[string]any{
							"templateId": "1234567890",
						},
					},
				},
			},
			Providers: map[string]map[string]any{
				"sendgrid": map[string]any{
					"templateId": "1234567890",
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Integrations](docs/sdks/integrations/README.md)

* [List](docs/sdks/integrations/README.md#list) - Get integrations
* [Create](docs/sdks/integrations/README.md#create) - Create integration
* [Update](docs/sdks/integrations/README.md#update) - Update integration
* [Delete](docs/sdks/integrations/README.md#delete) - Delete integration
* [SetAsPrimary](docs/sdks/integrations/README.md#setasprimary) - Set integration as primary
* [ListActive](docs/sdks/integrations/README.md#listactive) - Get active integrations

#### [Integrations.Webhooks](docs/sdks/webhooks/README.md)

* [Retrieve](docs/sdks/webhooks/README.md#retrieve) - Get webhook support status for provider

### [Messages](docs/sdks/messages/README.md)

* [Retrieve](docs/sdks/messages/README.md#retrieve) - Get messages
* [Delete](docs/sdks/messages/README.md#delete) - Delete message
* [DeleteByTransactionID](docs/sdks/messages/README.md#deletebytransactionid) - Delete messages by transactionId

### [Notifications](docs/sdks/notifications/README.md)

* [List](docs/sdks/notifications/README.md#list) - Get notifications
* [Retrieve](docs/sdks/notifications/README.md#retrieve) - Get notification

#### [Notifications.Stats](docs/sdks/stats/README.md)

* [Graph](docs/sdks/stats/README.md#graph) - Get notification graph statistics
* [Retrieve](docs/sdks/stats/README.md#retrieve) - Get notification statistics

### [Novu SDK](docs/sdks/novu/README.md)

* [Trigger](docs/sdks/novu/README.md#trigger) - Trigger event
* [Cancel](docs/sdks/novu/README.md#cancel) - Cancel triggered event
* [TriggerBroadcast](docs/sdks/novu/README.md#triggerbroadcast) - Broadcast event to all
* [TriggerBulk](docs/sdks/novu/README.md#triggerbulk) - Bulk trigger event

### [Subscribers](docs/sdks/subscribers/README.md)

* [Search](docs/sdks/subscribers/README.md#search) - Search for subscribers
* [Create](docs/sdks/subscribers/README.md#create) - Create subscriber
* [Retrieve](docs/sdks/subscribers/README.md#retrieve) - Get subscriber
* [Patch](docs/sdks/subscribers/README.md#patch) - Patch subscriber
* [Delete](docs/sdks/subscribers/README.md#delete) - Delete subscriber
* [List](docs/sdks/subscribers/README.md#list) - Get subscribers
* [Upsert](docs/sdks/subscribers/README.md#upsert) - Upsert subscriber
* [CreateBulk](docs/sdks/subscribers/README.md#createbulk) - Bulk create subscribers

#### [Subscribers.Authentication](docs/sdks/authentication/README.md)

* [ChatAccessOauth](docs/sdks/authentication/README.md#chataccessoauth) - Handle chat oauth
* [ChatAccessOauthCallBack](docs/sdks/authentication/README.md#chataccessoauthcallback) - Handle providers oauth redirect

#### [Subscribers.Credentials](docs/sdks/credentials/README.md)

* [Update](docs/sdks/credentials/README.md#update) - Update subscriber credentials
* [Append](docs/sdks/credentials/README.md#append) - Modify subscriber credentials
* [Delete](docs/sdks/credentials/README.md#delete) - Delete subscriber credentials by providerId

#### [Subscribers.Messages](docs/sdks/novumessages/README.md)

* [UpdateAsSeen](docs/sdks/novumessages/README.md#updateasseen) - Mark message action as seen
* [MarkAll](docs/sdks/novumessages/README.md#markall) - Marks all the subscriber messages as read, unread, seen or unseen.
* [MarkAllAs](docs/sdks/novumessages/README.md#markallas) - Mark a subscriber messages as seen, read, unseen or unread

#### [Subscribers.Notifications](docs/sdks/novunotifications/README.md)

* [Feed](docs/sdks/novunotifications/README.md#feed) - Get in-app notification feed for a particular subscriber
* [UnseenCount](docs/sdks/novunotifications/README.md#unseencount) - Get the unseen in-app notifications count for subscribers feed

#### [Subscribers.Preferences](docs/sdks/preferences/README.md)

* [List](docs/sdks/preferences/README.md#list) - Get subscriber preferences
* [Update](docs/sdks/preferences/README.md#update) - Update subscriber global or workflow specific preferences

#### [Subscribers.Properties](docs/sdks/properties/README.md)

* [UpdateOnlineFlag](docs/sdks/properties/README.md#updateonlineflag) - Update subscriber online status

#### [Subscribers.Topics](docs/sdks/novutopics/README.md)

* [List](docs/sdks/novutopics/README.md#list) - List topics a subscriber is subscribed to

### [Topics](docs/sdks/topics/README.md)

* [List](docs/sdks/topics/README.md#list) - Get topics list
* [Create](docs/sdks/topics/README.md#create) - Create or update a topic
* [Get](docs/sdks/topics/README.md#get) - Get topic by key
* [Update](docs/sdks/topics/README.md#update) - Update topic by key
* [Delete](docs/sdks/topics/README.md#delete) - Delete topic by key

#### [Topics.Subscribers](docs/sdks/novusubscribers/README.md)

* [Retrieve](docs/sdks/novusubscribers/README.md#retrieve) - Check topic subscriber

#### [Topics.Subscriptions](docs/sdks/subscriptions/README.md)

* [List](docs/sdks/subscriptions/README.md#list) - List topic subscriptions
* [Create](docs/sdks/subscriptions/README.md#create) - Create topic subscriptions, if the topic does not exist, it will be created.
* [Delete](docs/sdks/subscriptions/README.md#delete) - Delete topic subscriptions

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	novugo "github.com/novuhq/novu-go"
	"log"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
	)

	res, err := s.Subscribers.List(ctx, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"github.com/novuhq/novu-go/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
	)

	res, err := s.Trigger(ctx, components.TriggerEventRequestDto{
		WorkflowID: "workflow_identifier",
		Payload: map[string]any{
			"comment_id": "string",
			"post": map[string]any{
				"text": "string",
			},
		},
		Overrides: &components.Overrides{
			Steps: map[string]components.StepsOverrides{
				"email-step": components.StepsOverrides{
					Providers: map[string]map[string]any{
						"sendgrid": map[string]any{
							"templateId": "1234567890",
						},
					},
				},
			},
			Providers: map[string]map[string]any{
				"sendgrid": map[string]any{
					"templateId": "1234567890",
				},
			},
		},
		To: components.CreateToSubscriberPayloadDto(
			components.SubscriberPayloadDto{
				SubscriberID: "<id>",
			},
		),
	}, nil, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.TriggerEventResponseDto != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"github.com/novuhq/novu-go/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
	)

	res, err := s.Trigger(ctx, components.TriggerEventRequestDto{
		WorkflowID: "workflow_identifier",
		Payload: map[string]any{
			"comment_id": "string",
			"post": map[string]any{
				"text": "string",
			},
		},
		Overrides: &components.Overrides{
			Steps: map[string]components.StepsOverrides{
				"email-step": components.StepsOverrides{
					Providers: map[string]map[string]any{
						"sendgrid": map[string]any{
							"templateId": "1234567890",
						},
					},
				},
			},
			Providers: map[string]map[string]any{
				"sendgrid": map[string]any{
					"templateId": "1234567890",
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Trigger` function may return the following errors:

| Error Type                   | Status Code                            | Content Type     |
| ---------------------------- | -------------------------------------- | ---------------- |
| apierrors.ErrorDto           | 414                                    | application/json |
| apierrors.ErrorDto           | 400, 401, 403, 404, 405, 409, 413, 415 | application/json |
| apierrors.ValidationErrorDto | 422                                    | application/json |
| apierrors.ErrorDto           | 500                                    | application/json |
| apierrors.APIError           | 4XX, 5XX                               | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/apierrors"
	"github.com/novuhq/novu-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
	)

	res, err := s.Trigger(ctx, components.TriggerEventRequestDto{
		WorkflowID: "workflow_identifier",
		Payload: map[string]any{
			"comment_id": "string",
			"post": map[string]any{
				"text": "string",
			},
		},
		Overrides: &components.Overrides{
			Steps: map[string]components.StepsOverrides{
				"email-step": components.StepsOverrides{
					Providers: map[string]map[string]any{
						"sendgrid": map[string]any{
							"templateId": "1234567890",
						},
					},
				},
			},
			Providers: map[string]map[string]any{
				"sendgrid": map[string]any{
					"templateId": "1234567890",
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

		var e *apierrors.ErrorDto
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.ErrorDto
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.ValidationErrorDto
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.ErrorDto
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex(serverIndex int)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                   | Description |
| --- | ------------------------ | ----------- |
| 0   | `https://api.novu.co`    |             |
| 1   | `https://eu.api.novu.co` |             |

#### Example

```go
package main

import (
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithServerIndex(1),
		novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
	)

	res, err := s.Trigger(ctx, components.TriggerEventRequestDto{
		WorkflowID: "workflow_identifier",
		Payload: map[string]any{
			"comment_id": "string",
			"post": map[string]any{
				"text": "string",
			},
		},
		Overrides: &components.Overrides{
			Steps: map[string]components.StepsOverrides{
				"email-step": components.StepsOverrides{
					Providers: map[string]map[string]any{
						"sendgrid": map[string]any{
							"templateId": "1234567890",
						},
					},
				},
			},
			Providers: map[string]map[string]any{
				"sendgrid": map[string]any{
					"templateId": "1234567890",
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

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithServerURL("https://eu.api.novu.co"),
		novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
	)

	res, err := s.Trigger(ctx, components.TriggerEventRequestDto{
		WorkflowID: "workflow_identifier",
		Payload: map[string]any{
			"comment_id": "string",
			"post": map[string]any{
				"text": "string",
			},
		},
		Overrides: &components.Overrides{
			Steps: map[string]components.StepsOverrides{
				"email-step": components.StepsOverrides{
					Providers: map[string]map[string]any{
						"sendgrid": map[string]any{
							"templateId": "1234567890",
						},
					},
				},
			},
			Providers: map[string]map[string]any{
				"sendgrid": map[string]any{
					"templateId": "1234567890",
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/novuhq/novu-go&utm_campaign=go)
