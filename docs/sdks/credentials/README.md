# Credentials
(*Subscribers.Credentials*)

## Overview

### Available Operations

* [Update](#update) - Update provider credentials
* [Append](#append) - Upsert provider credentials
* [Delete](#delete) - Delete provider credentials

## Update

Update credentials for a provider such as **slack** and **FCM**. 
      **providerId** is required field. This API creates the **deviceTokens** or replaces the existing ones.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersV1Controller_updateSubscriberChannel" method="put" path="/v1/subscribers/{subscriberId}/credentials" -->
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

    res, err := s.Subscribers.Credentials.Update(ctx, "<id>", components.UpdateSubscriberChannelRequestDto{
        ProviderID: components.ChatOrPushProviderEnumSlack,
        Credentials: components.ChannelCredentials{
            WebhookURL: v3.Pointer("https://example.com/webhook"),
            Channel: v3.Pointer("general"),
            DeviceTokens: []string{
                "token1",
                "token2",
                "token3",
            },
            AlertUID: v3.Pointer("12345-abcde"),
            Title: v3.Pointer("Critical Alert"),
            ImageURL: v3.Pointer("https://example.com/image.png"),
            State: v3.Pointer("resolved"),
            ExternalURL: v3.Pointer("https://example.com/details"),
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriberResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `subscriberID`                                                                                               | *string*                                                                                                     | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `updateSubscriberChannelRequestDto`                                                                          | [components.UpdateSubscriberChannelRequestDto](../../models/components/updatesubscriberchannelrequestdto.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `idempotencyKey`                                                                                             | **string*                                                                                                    | :heavy_minus_sign:                                                                                           | A header for idempotency purposes                                                                            |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.SubscribersV1ControllerUpdateSubscriberChannelResponse](../../models/operations/subscribersv1controllerupdatesubscriberchannelresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Append

Upsert credentials for a provider such as **slack** and **FCM**. 
      **providerId** is required field. This API creates **deviceTokens** or appends to the existing ones.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersV1Controller_modifySubscriberChannel" method="patch" path="/v1/subscribers/{subscriberId}/credentials" -->
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

    res, err := s.Subscribers.Credentials.Append(ctx, "<id>", components.UpdateSubscriberChannelRequestDto{
        ProviderID: components.ChatOrPushProviderEnumOneSignal,
        Credentials: components.ChannelCredentials{
            WebhookURL: v3.Pointer("https://example.com/webhook"),
            Channel: v3.Pointer("general"),
            DeviceTokens: []string{
                "token1",
                "token2",
                "token3",
            },
            AlertUID: v3.Pointer("12345-abcde"),
            Title: v3.Pointer("Critical Alert"),
            ImageURL: v3.Pointer("https://example.com/image.png"),
            State: v3.Pointer("resolved"),
            ExternalURL: v3.Pointer("https://example.com/details"),
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriberResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `subscriberID`                                                                                               | *string*                                                                                                     | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `updateSubscriberChannelRequestDto`                                                                          | [components.UpdateSubscriberChannelRequestDto](../../models/components/updatesubscriberchannelrequestdto.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `idempotencyKey`                                                                                             | **string*                                                                                                    | :heavy_minus_sign:                                                                                           | A header for idempotency purposes                                                                            |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.SubscribersV1ControllerModifySubscriberChannelResponse](../../models/operations/subscribersv1controllermodifysubscriberchannelresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Delete

Delete subscriber credentials for a provider such as **slack** and **FCM** by **providerId**. 
    This action is irreversible and will remove the credentials for the provider for particular **subscriberId**.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersV1Controller_deleteSubscriberCredentials" method="delete" path="/v1/subscribers/{subscriberId}/credentials/{providerId}" -->
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

    res, err := s.Subscribers.Credentials.Delete(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `subscriberID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `providerID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SubscribersV1ControllerDeleteSubscriberCredentialsResponse](../../models/operations/subscribersv1controllerdeletesubscribercredentialsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |