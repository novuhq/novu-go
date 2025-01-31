# SubscribersPreferences
(*SubscribersPreferences*)

## Overview

### Available Operations

* [UpdateGlobal](#updateglobal) - Update subscriber global preferences
* [Update](#update) - Update subscriber preference

## UpdateGlobal

Update subscriber global preferences

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

    res, err := s.SubscribersPreferences.UpdateGlobal(ctx, "<id>", components.UpdateSubscriberGlobalPreferencesRequestDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateSubscriberPreferenceGlobalResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `subscriberID`                                                                                                                   | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `updateSubscriberGlobalPreferencesRequestDto`                                                                                    | [components.UpdateSubscriberGlobalPreferencesRequestDto](../../models/components/updatesubscriberglobalpreferencesrequestdto.md) | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `idempotencyKey`                                                                                                                 | **string*                                                                                                                        | :heavy_minus_sign:                                                                                                               | A header for idempotency purposes                                                                                                |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.SubscribersV1ControllerUpdateSubscriberGlobalPreferencesResponse](../../models/operations/subscribersv1controllerupdatesubscriberglobalpreferencesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Update

Update subscriber preference

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

    res, err := s.SubscribersPreferences.Update(ctx, "<id>", "<id>", components.UpdateSubscriberPreferenceRequestDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateSubscriberPreferenceResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `subscriberID`                                                                                                     | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `workflowID`                                                                                                       | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `updateSubscriberPreferenceRequestDto`                                                                             | [components.UpdateSubscriberPreferenceRequestDto](../../models/components/updatesubscriberpreferencerequestdto.md) | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `idempotencyKey`                                                                                                   | **string*                                                                                                          | :heavy_minus_sign:                                                                                                 | A header for idempotency purposes                                                                                  |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.SubscribersV1ControllerUpdateSubscriberPreferenceResponse](../../models/operations/subscribersv1controllerupdatesubscriberpreferenceresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |