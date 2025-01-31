# github.com/novuhq/novu-go

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/novuhq/novu-go* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/novuhq/novu-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/novu/novu). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Novu API: Novu REST API. Please see https://docs.novu.co/api-reference for more details.

For more information about the API: [Novu Documentation](https://docs.novu.co)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/novuhq/novu-go](#githubcomnovuhqnovu-go)
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
  * [Maturity](#maturity)
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Integrations](docs/sdks/integrations/README.md)

* [List](docs/sdks/integrations/README.md#list) - Get integrations
* [Create](docs/sdks/integrations/README.md#create) - Create integration
* [ListActive](docs/sdks/integrations/README.md#listactive) - Get active integrations
* [Update](docs/sdks/integrations/README.md#update) - Update integration
* [Delete](docs/sdks/integrations/README.md#delete) - Delete integration
* [SetPrimary](docs/sdks/integrations/README.md#setprimary) - Set integration as primary

#### [Integrations.Webhooks](docs/sdks/webhooks/README.md)

* [GetProviderStatus](docs/sdks/webhooks/README.md#getproviderstatus) - Get webhook support status for provider

### [Messages](docs/sdks/messages/README.md)

* [List](docs/sdks/messages/README.md#list) - Get messages
* [Delete](docs/sdks/messages/README.md#delete) - Delete message
* [DeleteByTransactionID](docs/sdks/messages/README.md#deletebytransactionid) - Delete messages by transactionId

### [Notifications](docs/sdks/notifications/README.md)

* [List](docs/sdks/notifications/README.md#list) - Get notifications
* [Get](docs/sdks/notifications/README.md#get) - Get notification

#### [Notifications.Stats](docs/sdks/stats/README.md)

* [Get](docs/sdks/stats/README.md#get) - Get notification statistics
* [GetGraph](docs/sdks/stats/README.md#getgraph) - Get notification graph statistics

### [Novu SDK](docs/sdks/novu/README.md)

* [Trigger](docs/sdks/novu/README.md#trigger) - Trigger event
* [TriggerBulk](docs/sdks/novu/README.md#triggerbulk) - Bulk trigger event
* [Broadcast](docs/sdks/novu/README.md#broadcast) - Broadcast event to all
* [Cancel](docs/sdks/novu/README.md#cancel) - Cancel triggered event

### [Subscribers](docs/sdks/subscribers/README.md)

* [List](docs/sdks/subscribers/README.md#list) - Get subscribers
* [Create](docs/sdks/subscribers/README.md#create) - Create subscriber
* [Get](docs/sdks/subscribers/README.md#get) - Get subscriber
* [Update](docs/sdks/subscribers/README.md#update) - Update subscriber
* [~~DeleteLegacy~~](docs/sdks/subscribers/README.md#deletelegacy) - Delete subscriber :warning: **Deprecated**
* [CreateBulk](docs/sdks/subscribers/README.md#createbulk) - Bulk create subscribers
* [Search](docs/sdks/subscribers/README.md#search) - Search for subscribers
* [Retrieve](docs/sdks/subscribers/README.md#retrieve) - Get subscriber
* [Patch](docs/sdks/subscribers/README.md#patch) - Patch subscriber
* [Delete](docs/sdks/subscribers/README.md#delete) - Delete subscriber
* [UpdateOnlineStatus](docs/sdks/subscribers/README.md#updateonlinestatus) - Update subscriber online status

#### [Subscribers.Authentication](docs/sdks/authentication/README.md)

* [OauthCallback](docs/sdks/authentication/README.md#oauthcallback) - Handle providers oauth redirect
* [ChatAccessOauth](docs/sdks/authentication/README.md#chataccessoauth) - Handle chat oauth

#### [Subscribers.Credentials](docs/sdks/credentials/README.md)

* [Update](docs/sdks/credentials/README.md#update) - Update subscriber credentials
* [Append](docs/sdks/credentials/README.md#append) - Modify subscriber credentials
* [Delete](docs/sdks/credentials/README.md#delete) - Delete subscriber credentials by providerId

#### [Subscribers.Messages](docs/sdks/novumessages/README.md)

* [MarkAs](docs/sdks/novumessages/README.md#markas) - Mark a subscriber messages as seen, read, unseen or unread
* [MarkAll](docs/sdks/novumessages/README.md#markall) - Marks all the subscriber messages as read, unread, seen or unseen. Optionally you can pass feed id (or array) to mark messages of a particular feed.
* [UpdateAction](docs/sdks/novumessages/README.md#updateaction) - Mark message action as seen

#### [Subscribers.Notifications](docs/sdks/novusubscribersnotifications/README.md)

* [GetFeed](docs/sdks/novusubscribersnotifications/README.md#getfeed) - Get in-app notification feed for a particular subscriber
* [UnseenCount](docs/sdks/novusubscribersnotifications/README.md#unseencount) - Get the unseen in-app notifications count for subscribers feed

#### [Subscribers.Preferences](docs/sdks/preferences/README.md)

* [List](docs/sdks/preferences/README.md#list) - Get subscriber preferences
* [RetrieveByLevel](docs/sdks/preferences/README.md#retrievebylevel) - Get subscriber preferences by level

### [SubscribersPreferences](docs/sdks/subscriberspreferences/README.md)

* [UpdateGlobal](docs/sdks/subscriberspreferences/README.md#updateglobal) - Update subscriber global preferences
* [Update](docs/sdks/subscriberspreferences/README.md#update) - Update subscriber preference

### [Topics](docs/sdks/topics/README.md)

* [Create](docs/sdks/topics/README.md#create) - Topic creation
* [List](docs/sdks/topics/README.md#list) - Get topic list filtered 
* [Delete](docs/sdks/topics/README.md#delete) - Delete topic
* [Get](docs/sdks/topics/README.md#get) - Get topic
* [Rename](docs/sdks/topics/README.md#rename) - Rename a topic

#### [Topics.Subscribers](docs/sdks/novutopicssubscribers/README.md)

* [Check](docs/sdks/novutopicssubscribers/README.md#check) - Check topic subscriber
* [Remove](docs/sdks/novutopicssubscribers/README.md#remove) - Subscribers removal

### [TopicsSubscribers](docs/sdks/topicssubscribers/README.md)

* [Add](docs/sdks/topicssubscribers/README.md#add) - Subscribers addition

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
	"os"
)

func main() {
	ctx := context.Background()

	s := novugo.New(
		novugo.WithSecurity(os.Getenv("NOVU_SECRET_KEY")),
	)

	res, err := s.Subscribers.List(ctx, nil, novugo.Float64(10), nil)
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
	"os"
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

| #   | Server                   |
| --- | ------------------------ |
| 0   | `https://api.novu.co`    |
| 1   | `https://eu.api.novu.co` |

#### Example

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
		novugo.WithServerIndex(1),
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

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
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
		novugo.WithServerURL("https://api.novu.co"),
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

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/novuhq/novu-go&utm_campaign=go)
