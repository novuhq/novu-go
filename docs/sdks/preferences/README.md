# Preferences
(*Subscribers.Preferences*)

## Overview

### Available Operations

* [List](#list) - Retrieve subscriber preferences
* [Update](#update) - Update subscriber preferences

## List

Retrieve subscriber channel preferences by its unique key identifier **subscriberId**. 
    This API returns all five channels preferences for all workflows and global preferences.

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_getSubscriberPreferences" method="get" path="/v2/subscribers/{subscriberId}/preferences" -->
```go
package main

import(
	"context"
	novugo "github.com/novuhq/novu-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := novugo.New(
        novugo.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Subscribers.Preferences.List(ctx, "<id>", nil)
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

Update subscriber preferences by its unique key identifier **subscriberId**. 
    **workflowId** is optional field, if provided, this API will update that workflow preference, 
    otherwise it will update global preferences

### Example Usage

<!-- UsageSnippet language="go" operationID="SubscribersController_updateSubscriberPreferences" method="patch" path="/v2/subscribers/{subscriberId}/preferences" -->
```go
package main

import(
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