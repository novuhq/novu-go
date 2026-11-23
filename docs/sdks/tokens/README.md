# Subscribers.Tokens

## Overview

### Available Operations

* [Create](#create) - Create a subscriber login token

## Create

Create a short-lived, Novu-signed login token for the subscriber identified by **subscriberId**.
    Call it from your backend with your secret key, then pass the token to the browser, which exchanges it for an Inbox session.
    Creating a token does not read or change the subscriber, so the subscriber does not need to exist yet.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_createSubscriberToken" method="post" path="/v2/subscribers/{subscriberId}/tokens" -->
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

    res, err := s.Subscribers.Tokens.Create(ctx, "<id>", components.CreateSubscriberTokenRequestDto{
        Context: map[string]components.CreateSubscriberTokenRequestDtoContext{
            "key": components.CreateCreateSubscriberTokenRequestDtoContextStr(
                "org-acme",
            ),
        },
        Subscriber: &components.BaseSubscriberFieldsDto{
            FirstName: v3.Pointer("John"),
            LastName: v3.Pointer("Doe"),
            Email: v3.Pointer("john.doe@example.com"),
            Phone: v3.Pointer("+1234567890"),
            Avatar: v3.Pointer("https://example.com/avatar.jpg"),
            Locale: v3.Pointer("en-US"),
            Timezone: v3.Pointer("America/New_York"),
        },
        Scope: []components.SubscriberTokenScopeEnum{
            components.SubscriberTokenScopeEnumChannelsShared,
        },
        ExpiresIn: v3.Pointer[int64](600),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateSubscriberTokenResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `subscriberID`                                                                                           | `string`                                                                                                 | :heavy_check_mark:                                                                                       | The identifier of the subscriber                                                                         |
| `createSubscriberTokenRequestDto`                                                                        | [components.CreateSubscriberTokenRequestDto](../../models/components/createsubscribertokenrequestdto.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `idempotencyKey`                                                                                         | `*string`                                                                                                | :heavy_minus_sign:                                                                                       | A header for idempotency purposes                                                                        |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.SubscribersControllerCreateSubscriberTokenResponse](../../models/operations/subscriberscontrollercreatesubscribertokenresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |