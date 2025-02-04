# Preferences
(*Subscribers.Preferences*)

## Overview

### Available Operations

* [List](#list) - Get subscriber preferences
* [UpdateGlobal](#updateglobal) - Update subscriber global preferences
* [RetrieveByLevel](#retrievebylevel) - Get subscriber preferences by level
* [UpdateLegacy](#updatelegacy) - Update subscriber preference
* [Retrieve](#retrieve) - Get subscriber preferences
* [Update](#update) - Update subscriber global or workflow specific preferences

## List

Get subscriber preferences

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

    res, err := s.Subscribers.Preferences.List(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateSubscriberPreferenceResponseDtos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                 | Type                                                                                                                      | Required                                                                                                                  | Description                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                     | :heavy_check_mark:                                                                                                        | The context to use for the request.                                                                                       |
| `subscriberID`                                                                                                            | *string*                                                                                                                  | :heavy_check_mark:                                                                                                        | N/A                                                                                                                       |
| `includeInactiveChannels`                                                                                                 | **bool*                                                                                                                   | :heavy_minus_sign:                                                                                                        | A flag which specifies if the inactive workflow channels should be included in the retrieved preferences. Default is true |
| `idempotencyKey`                                                                                                          | **string*                                                                                                                 | :heavy_minus_sign:                                                                                                        | A header for idempotency purposes                                                                                         |
| `opts`                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                  | :heavy_minus_sign:                                                                                                        | The options for this request.                                                                                             |

### Response

**[*operations.SubscribersV1ControllerListSubscriberPreferencesResponse](../../models/operations/subscribersv1controllerlistsubscriberpreferencesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

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

    res, err := s.Subscribers.Preferences.UpdateGlobal(ctx, "<id>", components.UpdateSubscriberGlobalPreferencesRequestDto{}, nil)
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

## RetrieveByLevel

Get subscriber preferences by level

### Example Usage

```go
package main

import(
	"context"
	"os"
	novugo "github.com/novuhq/novu-go"
	"github.com/novuhq/novu-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := novugo.New(
        novugo.WithSecurity(os.Getenv("NOVU_SECRET_KEY")),
    )

    res, err := s.Subscribers.Preferences.RetrieveByLevel(ctx, operations.ParameterTemplate, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSubscriberPreferencesResponseDtos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                 | Type                                                                                                                      | Required                                                                                                                  | Description                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                     | :heavy_check_mark:                                                                                                        | The context to use for the request.                                                                                       |
| `preferenceLevel`                                                                                                         | [operations.Parameter](../../models/operations/parameter.md)                                                              | :heavy_check_mark:                                                                                                        | the preferences level to be retrieved (template / global)                                                                 |
| `subscriberID`                                                                                                            | *string*                                                                                                                  | :heavy_check_mark:                                                                                                        | N/A                                                                                                                       |
| `includeInactiveChannels`                                                                                                 | **bool*                                                                                                                   | :heavy_minus_sign:                                                                                                        | A flag which specifies if the inactive workflow channels should be included in the retrieved preferences. Default is true |
| `idempotencyKey`                                                                                                          | **string*                                                                                                                 | :heavy_minus_sign:                                                                                                        | A header for idempotency purposes                                                                                         |
| `opts`                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                  | :heavy_minus_sign:                                                                                                        | The options for this request.                                                                                             |

### Response

**[*operations.SubscribersV1ControllerGetSubscriberPreferenceByLevelResponse](../../models/operations/subscribersv1controllergetsubscriberpreferencebylevelresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateLegacy

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

    res, err := s.Subscribers.Preferences.UpdateLegacy(ctx, "<id>", "<id>", components.UpdateSubscriberPreferenceRequestDto{}, nil)
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

## Retrieve

Get subscriber global and workflow specific preferences

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

    res, err := s.Subscribers.Preferences.Retrieve(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSubscriberPreferencesDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `subscriberID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | **string*                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SubscribersControllerGetSubscriberPreferencesResponse](../../models/operations/subscriberscontrollergetsubscriberpreferencesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Update

Update subscriber global or workflow specific preferences

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

    res, err := s.Subscribers.Preferences.Update(ctx, "<id>", components.PatchSubscriberPreferencesDto{
        Channels: components.PatchPreferenceChannelsDto{},
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSubscriberPreferencesDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `subscriberID`                                                                                       | *string*                                                                                             | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `patchSubscriberPreferencesDto`                                                                      | [components.PatchSubscriberPreferencesDto](../../models/components/patchsubscriberpreferencesdto.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `idempotencyKey`                                                                                     | **string*                                                                                            | :heavy_minus_sign:                                                                                   | A header for idempotency purposes                                                                    |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.SubscribersControllerUpdateSubscriberPreferencesResponse](../../models/operations/subscriberscontrollerupdatesubscriberpreferencesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |