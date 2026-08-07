# Agents

## Overview

Agents are conversational assistants that receive inbound messages from connected channels and respond through a custom code bridge or a managed runtime provider.
<https://docs.novu.co/agents>

### Available Operations

* [Create](#create) - Create an agent
* [List](#list) - List all agents
* [SendReply](#sendreply) - Send an agent reply
* [Retrieve](#retrieve) - Retrieve an agent
* [Update](#update) - Update an agent
* [Delete](#delete) - Delete an agent
* [UpdateBridge](#updatebridge) - Update an agent bridge

## Create

Create an agent scoped to the current environment. The identifier must be unique per environment. Set `runtime` to `managed` and supply `managedRuntime` to provision a provider-hosted agent brain.

### Example Usage

<!-- UsageSnippet language="go" operationID="AgentsController_createAgent" method="post" path="/v1/agents" -->
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

    res, err := s.Agents.Create(ctx, "<value>", components.CreateAgentRequestDto{
        Name: "<value>",
        Identifier: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `novuAnalyticsSource`                                                                | `string`                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `createAgentRequestDto`                                                              | [components.CreateAgentRequestDto](../../models/components/createagentrequestdto.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `idempotencyKey`                                                                     | `*string`                                                                            | :heavy_minus_sign:                                                                   | A header for idempotency purposes                                                    |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.AgentsControllerCreateAgentResponse](../../models/operations/agentscontrollercreateagentresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## List

Retrieve a cursor-paginated list of agents for the current environment. Use **after**, **before**, **limit**, **orderBy**, and **orderDirection** query parameters.

### Example Usage

<!-- UsageSnippet language="go" operationID="AgentsController_listAgents" method="get" path="/v1/agents" -->
```go
package main

import(
	"context"
	"github.com/novuhq/novu-go/v3"
	"github.com/novuhq/novu-go/v3/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity("YOUR_SECRET_KEY_HERE"),
    )

    res, err := s.Agents.List(ctx, operations.AgentsControllerListAgentsRequest{
        Limit: v3.Pointer[float64](10.0),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAgentsResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.AgentsControllerListAgentsRequest](../../models/operations/agentscontrollerlistagentsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.AgentsControllerListAgentsResponse](../../models/operations/agentscontrollerlistagentsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ErrorDto                     | 414                                    | application/json                       |
| apierrors.ErrorDto                     | 400, 401, 403, 404, 405, 409, 413, 415 | application/json                       |
| apierrors.ValidationErrorDto           | 422                                    | application/json                       |
| apierrors.ErrorDto                     | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## SendReply

Send a message or side-effect into an existing agent conversation from your backend.

Use this endpoint when you are not using `@novu/framework` (for example Python, Go, PHP, .NET, or Java SDKs),
or when a server process outside the bridge needs to post into a live conversation.

**Message actions**
- `reply` — markdown, interactive card, or tool-approval card (optional `files`)
- `edit` — update a previously delivered message in place
- `deleteMessages` — remove rendered platform messages (history is kept)
- `addReactions` — add emoji reactions to existing messages

**Turn control**
- `typing` — `{ status?: string }` to set status, or `"stop"` to clear
- `resolve` — mark the conversation resolved (optionally with a final reply)
- `error: true` — report a customer-runtime failure (cannot combine with other actions)

**Signals & tools**
- `signals` — metadata set/delete/clear, or trigger a Novu workflow
- `toolResults` — persist tool outputs into conversation history
- `toolApprovalRequest` — ledger a gated tool call (pair with an approval card reply)

Returns `{ data: { messageId, platformThreadId } }` when a reply or edit is delivered;
otherwise `{ data: null }`.

### Example Usage: addReaction

<!-- UsageSnippet language="go" operationID="AgentReplyController_handleAgentReplyHandler" method="post" path="/v1/agents/{agentId}/reply" example="addReaction" -->
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

    res, err := s.Agents.SendReply(ctx, "support-agent", components.AgentReplyPayloadDto{
        ConversationID: "64f5a1c2e8b7a3d9f0c1b2a3",
        IntegrationIdentifier: "slack-support",
        AddReactions: []components.AddReactionPayloadDto{
            components.AddReactionPayloadDto{
                MessageID: "1712345678.123456",
                EmojiName: "white_check_mark",
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: cardReply

<!-- UsageSnippet language="go" operationID="AgentReplyController_handleAgentReplyHandler" method="post" path="/v1/agents/{agentId}/reply" example="cardReply" -->
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

    res, err := s.Agents.SendReply(ctx, "support-agent", components.AgentReplyPayloadDto{
        ConversationID: "64f5a1c2e8b7a3d9f0c1b2a3",
        IntegrationIdentifier: "slack-support",
        Reply: v3.Pointer(components.CreateReplyCardReplyContentDto(
            components.CardReplyContentDto{
                Card: map[string]any{
                    "type": "card",
                    "title": "Order #123",
                    "children": []any{
                        map[string]any{
                            "type": "text",
                            "content": "Your order is ready for pickup.",
                        },
                        map[string]any{
                            "type": "button",
                            "id": "confirm",
                            "label": "Confirm",
                            "style": "primary",
                        },
                    },
                },
            },
        )),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: deleteMessage

<!-- UsageSnippet language="go" operationID="AgentReplyController_handleAgentReplyHandler" method="post" path="/v1/agents/{agentId}/reply" example="deleteMessage" -->
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

    res, err := s.Agents.SendReply(ctx, "support-agent", components.AgentReplyPayloadDto{
        ConversationID: "64f5a1c2e8b7a3d9f0c1b2a3",
        IntegrationIdentifier: "slack-support",
        DeleteMessages: []components.DeleteMessagePayloadDto{
            components.DeleteMessagePayloadDto{
                MessageID: "1712345678.123456",
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: editMessage

<!-- UsageSnippet language="go" operationID="AgentReplyController_handleAgentReplyHandler" method="post" path="/v1/agents/{agentId}/reply" example="editMessage" -->
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

    res, err := s.Agents.SendReply(ctx, "support-agent", components.AgentReplyPayloadDto{
        ConversationID: "64f5a1c2e8b7a3d9f0c1b2a3",
        IntegrationIdentifier: "slack-support",
        Edit: &components.EditPayloadDto{
            MessageID: "1712345678.123456",
            Content: components.CreateContentMarkdownReplyContentDto(
                components.MarkdownReplyContentDto{
                    Markdown: "Updated: the report is now final.",
                },
            ),
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: markdownReply

<!-- UsageSnippet language="go" operationID="AgentReplyController_handleAgentReplyHandler" method="post" path="/v1/agents/{agentId}/reply" example="markdownReply" -->
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

    res, err := s.Agents.SendReply(ctx, "support-agent", components.AgentReplyPayloadDto{
        ConversationID: "64f5a1c2e8b7a3d9f0c1b2a3",
        IntegrationIdentifier: "slack-support",
        Reply: v3.Pointer(components.CreateReplyMarkdownReplyContentDto(
            components.MarkdownReplyContentDto{
                Markdown: "**Report ready.** Your weekly summary is attached.",
            },
        )),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: metadataSignal

<!-- UsageSnippet language="go" operationID="AgentReplyController_handleAgentReplyHandler" method="post" path="/v1/agents/{agentId}/reply" example="metadataSignal" -->
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

    res, err := s.Agents.SendReply(ctx, "support-agent", components.AgentReplyPayloadDto{
        ConversationID: "64f5a1c2e8b7a3d9f0c1b2a3",
        IntegrationIdentifier: "slack-support",
        Signals: []components.Signals{
            components.CreateSignalsTriggerSignalDto(
                components.TriggerSignalDto{
                    Type: components.TriggerSignalDtoTypeTrigger,
                    WorkflowID: "order-shipped",
                    To: v3.Pointer(components.CreateToStr(
                        "subscriber-123",
                    )),
                    Payload: map[string]any{
                        "orderId": "ORD-42",
                    },
                },
            ),
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: replyWithFile

<!-- UsageSnippet language="go" operationID="AgentReplyController_handleAgentReplyHandler" method="post" path="/v1/agents/{agentId}/reply" example="replyWithFile" -->
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

    res, err := s.Agents.SendReply(ctx, "support-agent", components.AgentReplyPayloadDto{
        ConversationID: "64f5a1c2e8b7a3d9f0c1b2a3",
        IntegrationIdentifier: "slack-support",
        Reply: v3.Pointer(components.CreateReplyMarkdownReplyContentDto(
            components.MarkdownReplyContentDto{
                Markdown: "Here is your report.",
                Files: []components.FileRefDto{
                    components.FileRefDto{
                        Filename: "report.pdf",
                        MimeType: v3.Pointer("application/pdf"),
                        URL: v3.Pointer("https://example.com/files/report.pdf"),
                    },
                },
            },
        )),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: resolveConversation

<!-- UsageSnippet language="go" operationID="AgentReplyController_handleAgentReplyHandler" method="post" path="/v1/agents/{agentId}/reply" example="resolveConversation" -->
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

    res, err := s.Agents.SendReply(ctx, "support-agent", components.AgentReplyPayloadDto{
        ConversationID: "64f5a1c2e8b7a3d9f0c1b2a3",
        IntegrationIdentifier: "slack-support",
        Reply: v3.Pointer(components.CreateReplyMarkdownReplyContentDto(
            components.MarkdownReplyContentDto{
                Markdown: "Glad that helped — marking this as resolved.",
            },
        )),
        Resolve: &components.ResolveDto{
            Summary: v3.Pointer("Answered billing question about invoice INV-42."),
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: toolApprovalRequest

<!-- UsageSnippet language="go" operationID="AgentReplyController_handleAgentReplyHandler" method="post" path="/v1/agents/{agentId}/reply" example="toolApprovalRequest" -->
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

    res, err := s.Agents.SendReply(ctx, "support-agent", components.AgentReplyPayloadDto{
        ConversationID: "64f5a1c2e8b7a3d9f0c1b2a3",
        IntegrationIdentifier: "slack-support",
        Reply: v3.Pointer(components.CreateReplyToolApprovalCardReplyContentDto(
            components.ToolApprovalCardReplyContentDto{
                ToolApprovalCard: map[string]any{
                    "type": "tool-approval-card",
                    "title": "Approve refund?",
                    "subtitle": "issue_refund · ORD-42 · $25.00",
                    "approveLabel": "Approve",
                    "denyLabel": "Deny",
                },
            },
        )),
        ToolApprovalRequest: &components.ToolApprovalRequestPayloadDto{
            ApprovalID: "apr_01HZX",
            ToolCallID: "call_refund_1",
            Name: "issue_refund",
            Input: map[string]any{
                "orderId": "ORD-42",
                "amountCents": 2500,
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: toolResult

<!-- UsageSnippet language="go" operationID="AgentReplyController_handleAgentReplyHandler" method="post" path="/v1/agents/{agentId}/reply" example="toolResult" -->
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

    res, err := s.Agents.SendReply(ctx, "support-agent", components.AgentReplyPayloadDto{
        ConversationID: "64f5a1c2e8b7a3d9f0c1b2a3",
        IntegrationIdentifier: "slack-support",
        Reply: v3.Pointer(components.CreateReplyMarkdownReplyContentDto(
            components.MarkdownReplyContentDto{
                Markdown: "Your order **ORD-42** has shipped and should arrive by July 16.",
            },
        )),
        ToolResults: []components.ToolResultDto{
            components.ToolResultDto{
                ToolCallID: "call_abc123",
                ToolName: v3.Pointer("lookup_order"),
                Output: &components.Output{},
                Preview: v3.Pointer("Order ORD-42 is shipped"),
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: triggerWorkflow

<!-- UsageSnippet language="go" operationID="AgentReplyController_handleAgentReplyHandler" method="post" path="/v1/agents/{agentId}/reply" example="triggerWorkflow" -->
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

    res, err := s.Agents.SendReply(ctx, "support-agent", components.AgentReplyPayloadDto{
        ConversationID: "64f5a1c2e8b7a3d9f0c1b2a3",
        IntegrationIdentifier: "slack-support",
        Signals: []components.Signals{
            components.CreateSignalsTriggerSignalDto(
                components.TriggerSignalDto{
                    Type: components.TriggerSignalDtoTypeTrigger,
                    WorkflowID: "order-shipped",
                    To: v3.Pointer(components.CreateToStr(
                        "subscriber-123",
                    )),
                    Payload: map[string]any{
                        "orderId": "ORD-42",
                    },
                },
            ),
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: turnError

<!-- UsageSnippet language="go" operationID="AgentReplyController_handleAgentReplyHandler" method="post" path="/v1/agents/{agentId}/reply" example="turnError" -->
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

    res, err := s.Agents.SendReply(ctx, "support-agent", components.AgentReplyPayloadDto{
        ConversationID: "64f5a1c2e8b7a3d9f0c1b2a3",
        IntegrationIdentifier: "slack-support",
        Error: v3.Pointer(true),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: typingStart

<!-- UsageSnippet language="go" operationID="AgentReplyController_handleAgentReplyHandler" method="post" path="/v1/agents/{agentId}/reply" example="typingStart" -->
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

    res, err := s.Agents.SendReply(ctx, "support-agent", components.AgentReplyPayloadDto{
        ConversationID: "64f5a1c2e8b7a3d9f0c1b2a3",
        IntegrationIdentifier: "slack-support",
        Typing: v3.Pointer(components.CreateTypingTypingStatusDto(
            components.TypingStatusDto{
                Status: v3.Pointer("Looking up your order…"),
            },
        )),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: typingStop

<!-- UsageSnippet language="go" operationID="AgentReplyController_handleAgentReplyHandler" method="post" path="/v1/agents/{agentId}/reply" example="typingStop" -->
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

    res, err := s.Agents.SendReply(ctx, "support-agent", components.AgentReplyPayloadDto{
        ConversationID: "64f5a1c2e8b7a3d9f0c1b2a3",
        IntegrationIdentifier: "slack-support",
        Typing: v3.Pointer(components.CreateTypingTyping1(
            components.Typing1Stop,
        )),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                       | Type                                                                                                                                                                                                                                            | Required                                                                                                                                                                                                                                        | Description                                                                                                                                                                                                                                     | Example                                                                                                                                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                              | The context to use for the request.                                                                                                                                                                                                             |                                                                                                                                                                                                                                                 |
| `agentID`                                                                                                                                                                                                                                       | `string`                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                              | Agent identifier (slug) for the agent that owns the conversation.                                                                                                                                                                               | support-agent                                                                                                                                                                                                                                   |
| `agentReplyPayloadDto`                                                                                                                                                                                                                          | [components.AgentReplyPayloadDto](../../models/components/agentreplypayloaddto.md)                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                              | Reply payload. Provide at least one action: `reply`, `edit`, `resolve`, `signals`, `toolResults`, `toolApprovalRequest`, `addReactions`, `deleteMessages`, `typing`, or `error`. See named examples for common shapes used by server-side SDKs. |                                                                                                                                                                                                                                                 |
| `idempotencyKey`                                                                                                                                                                                                                                | `*string`                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                              | A header for idempotency purposes                                                                                                                                                                                                               |                                                                                                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                              | The options for this request.                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                 |

### Response

**[*operations.AgentReplyControllerHandleAgentReplyHandlerResponse](../../models/operations/agentreplycontrollerhandleagentreplyhandlerresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Retrieve

Retrieve an agent by its external identifier (not the internal MongoDB id).

### Example Usage

<!-- UsageSnippet language="go" operationID="AgentsController_getAgent" method="get" path="/v1/agents/{identifier}" -->
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

    res, err := s.Agents.Retrieve(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `identifier`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AgentsControllerGetAgentResponse](../../models/operations/agentscontrollergetagentresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Update

Update an agent by its external identifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="AgentsController_updateAgent" method="patch" path="/v1/agents/{identifier}" -->
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

    res, err := s.Agents.Update(ctx, "<value>", components.UpdateAgentRequestDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `identifier`                                                                         | `string`                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `updateAgentRequestDto`                                                              | [components.UpdateAgentRequestDto](../../models/components/updateagentrequestdto.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `idempotencyKey`                                                                     | `*string`                                                                            | :heavy_minus_sign:                                                                   | A header for idempotency purposes                                                    |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.AgentsControllerUpdateAgentResponse](../../models/operations/agentscontrollerupdateagentresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Delete

Delete an agent by identifier, remove all agent-integration links, and clear the agent assignment from any workflows that reference it. For managed-runtime agents, pass `deleteFromProvider=true` to also archive the agent on the provider side (e.g. Anthropic). By default only the Novu record is deleted and the provider agent is left intact.

### Example Usage

<!-- UsageSnippet language="go" operationID="AgentsController_deleteAgent" method="delete" path="/v1/agents/{identifier}" -->
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

    res, err := s.Agents.Delete(ctx, "<value>", "<value>", nil)
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
| `identifier`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `deleteFromProvider`                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `idempotencyKey`                                         | `*string`                                                | :heavy_minus_sign:                                       | A header for idempotency purposes                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AgentsControllerDeleteAgentResponse](../../models/operations/agentscontrollerdeleteagentresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## UpdateBridge

Update the bridge URL configuration for an agent. Used by the CLI to register dev tunnel URLs. Refuses to activate dev bridges on production environments.

### Example Usage

<!-- UsageSnippet language="go" operationID="AgentsController_updateAgentBridge" method="put" path="/v1/agents/{identifier}/bridge" -->
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

    res, err := s.Agents.UpdateBridge(ctx, "<value>", components.UpdateAgentBridgeRequestDto{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentResponseDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `identifier`                                                                                     | `string`                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `updateAgentBridgeRequestDto`                                                                    | [components.UpdateAgentBridgeRequestDto](../../models/components/updateagentbridgerequestdto.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `idempotencyKey`                                                                                 | `*string`                                                                                        | :heavy_minus_sign:                                                                               | A header for idempotency purposes                                                                |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.AgentsControllerUpdateAgentBridgeResponse](../../models/operations/agentscontrollerupdateagentbridgeresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.ErrorDto                | 414                               | application/json                  |
| apierrors.ErrorDto                | 400, 401, 403, 405, 409, 413, 415 | application/json                  |
| apierrors.ValidationErrorDto      | 422                               | application/json                  |
| apierrors.ErrorDto                | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |