# Workflows.Steps

## Overview

### Available Operations

* [GeneratePreview](#generatepreview) - Generate a step preview
* [Retrieve](#retrieve) - Retrieve workflow step

## GeneratePreview

Generates a preview for a specific workflow step by its unique identifier **stepId**

### Example Usage

<!-- UsageSnippet language="go" operationID="WorkflowController_generatePreview" method="post" path="/v2/workflows/{workflowId}/step/{stepId}/preview" -->
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

    res, err := s.Workflows.Steps.GeneratePreview(ctx, "<id>", "<id>", components.GeneratePreviewRequestDto{
        PreviewPayload: &components.PreviewPayloadDto{
            Subscriber: &components.SubscriberResponseDtoOptional{
                Channels: []components.ChannelSettingsDto{
                    components.ChannelSettingsDto{
                        ProviderID: components.ChatOrPushProviderEnumNovuSlack,
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
                        IntegrationID: "<id>",
                    },
                },
            },
            Actor: &components.SubscriberResponseDtoOptional{
                Channels: []components.ChannelSettingsDto{
                    components.ChannelSettingsDto{
                        ProviderID: components.ChatOrPushProviderEnumTelegram,
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
                        IntegrationID: "<id>",
                    },
                },
            },
            Context: map[string]components.PreviewPayloadDtoContext{
                "key": components.CreatePreviewPayloadDtoContextStr(
                    "org-acme",
                ),
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GeneratePreviewResponseDto != nil {
        switch res.GeneratePreviewResponseDto.Result.Type {
            case components.GeneratePreviewResponseDtoResultUnionTypeMapOfAny:
                // res.GeneratePreviewResponseDto.Result.MapOfAny is populated
            case components.GeneratePreviewResponseDtoResultUnionTypeResult2:
                // res.GeneratePreviewResponseDto.Result.Result2 is populated
            case components.GeneratePreviewResponseDtoResultUnionTypeThree:
                // res.GeneratePreviewResponseDto.Result.Three is populated
            case components.GeneratePreviewResponseDtoResultUnionTypeResult4:
                // res.GeneratePreviewResponseDto.Result.Result4 is populated
            case components.GeneratePreviewResponseDtoResultUnionTypeResult5:
                // res.GeneratePreviewResponseDto.Result.Result5 is populated
            case components.GeneratePreviewResponseDtoResultUnionTypeSix:
                // res.GeneratePreviewResponseDto.Result.Six is populated
            case components.GeneratePreviewResponseDtoResultUnionTypeSeven:
                // res.GeneratePreviewResponseDto.Result.Seven is populated
            case components.GeneratePreviewResponseDtoResultUnionTypeEight:
                // res.GeneratePreviewResponseDto.Result.Eight is populated
            case components.GeneratePreviewResponseDtoResultUnionTypeNine:
                // res.GeneratePreviewResponseDto.Result.Nine is populated
            case components.GeneratePreviewResponseDtoResultUnionTypeTen:
                // res.GeneratePreviewResponseDto.Result.Ten is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `workflowID`                                                                                 | `string`                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `stepID`                                                                                     | `string`                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `generatePreviewRequestDto`                                                                  | [components.GeneratePreviewRequestDto](../../models/components/generatepreviewrequestdto.md) | :heavy_check_mark:                                                                           | Preview generation details                                                                   |
| `idempotencyKey`                                                                             | `*string`                                                                                    | :heavy_minus_sign:                                                                           | A header for idempotency purposes                                                            |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.WorkflowControllerGeneratePreviewResponse](../../models/operations/workflowcontrollergeneratepreviewresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Retrieve

Retrieves data for a specific step in a workflow

### Example Usage

<!-- UsageSnippet language="go" operationID="WorkflowController_getWorkflowStepData" method="get" path="/v2/workflows/{workflowId}/steps/{stepId}" -->
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

    res, err := s.Workflows.Steps.Retrieve(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.StepResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `workflowID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `stepID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WorkflowControllerGetWorkflowStepDataResponse](../../models/operations/workflowcontrollergetworkflowstepdataresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |