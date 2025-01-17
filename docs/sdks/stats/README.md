# Stats
(*Notifications.Stats*)

## Overview

### Available Operations

* [Get](#get) - Get notification statistics
* [GetGraph](#getgraph) - Get notification graph statistics

## Get

Get notification statistics

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
        novugo.WithSecurity(os.Getenv("NOVU_API_KEY")),
    )

    res, err := s.Notifications.Stats.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ActivityStatsResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.NotificationsControllerGetActivityStatsResponse](../../models/operations/notificationscontrollergetactivitystatsresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 404, 409                | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |

## GetGraph

Get notification graph statistics

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
        novugo.WithSecurity(os.Getenv("NOVU_API_KEY")),
    )

    res, err := s.Notifications.Stats.GetGraph(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ActivityGraphStatesResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `days`                                                   | **float64*                                               | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.NotificationsControllerGetActivityGraphStatsResponse](../../models/operations/notificationscontrollergetactivitygraphstatsresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.ErrorDto           | 400, 404, 409                | application/json             |
| apierrors.ValidationErrorDto | 422                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |