# Credentials
(*Subscribers.Credentials*)

## Overview

### Available Operations

* [Update](#update) - Update subscriber credentials
* [Append](#append) - Modify subscriber credentials
* [Delete](#delete) - Delete subscriber credentials by providerId

## Update

Subscriber credentials associated to the delivery methods such as slack and push tokens.

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
        novugo.WithSecurity(os.Getenv("NOVU_SECRET_KEY")),
    )

    res, err := s.Subscribers.Credentials.Update(ctx, "<id>", components.UpdateSubscriberChannelRequestDto{
        ProviderID: components.UpdateSubscriberChannelRequestDtoProviderIDPushpad,
        Credentials: components.ChannelCredentials{
            WebhookURL: novugo.String("https://example.com/webhook"),
            Channel: novugo.String("general"),
            DeviceTokens: []string{
                "token1",
                "token2",
                "token3",
            },
            AlertUID: novugo.String("12345-abcde"),
            Title: novugo.String("Critical Alert"),
            ImageURL: novugo.String("https://example.com/image.png"),
            State: novugo.String("resolved"),
            ExternalURL: novugo.String("https://example.com/details"),
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

Subscriber credentials associated to the delivery methods such as slack and push tokens.



    This endpoint appends provided credentials and deviceTokens to the existing ones.

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
        novugo.WithSecurity(os.Getenv("NOVU_SECRET_KEY")),
    )

    res, err := s.Subscribers.Credentials.Append(ctx, "<id>", components.UpdateSubscriberChannelRequestDto{
        ProviderID: components.UpdateSubscriberChannelRequestDtoProviderIDZulip,
        Credentials: components.ChannelCredentials{
            WebhookURL: novugo.String("https://example.com/webhook"),
            Channel: novugo.String("general"),
            DeviceTokens: []string{
                "token1",
                "token2",
                "token3",
            },
            AlertUID: novugo.String("12345-abcde"),
            Title: novugo.String("Critical Alert"),
            ImageURL: novugo.String("https://example.com/image.png"),
            State: novugo.String("resolved"),
            ExternalURL: novugo.String("https://example.com/details"),
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

Delete subscriber credentials such as slack and expo tokens.

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
        novugo.WithSecurity(os.Getenv("NOVU_SECRET_KEY")),
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