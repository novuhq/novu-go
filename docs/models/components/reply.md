# Reply

Outbound message content. Exactly one of `markdown`, `card`, or `toolApprovalCard`. Optional `files` attach to the message. Cannot be combined with `edit`.


## Supported Types

### MarkdownReplyContentDto

```go
reply := components.CreateReplyMarkdownReplyContentDto(components.MarkdownReplyContentDto{/* values here */})
```

### CardReplyContentDto

```go
reply := components.CreateReplyCardReplyContentDto(components.CardReplyContentDto{/* values here */})
```

### ToolApprovalCardReplyContentDto

```go
reply := components.CreateReplyToolApprovalCardReplyContentDto(components.ToolApprovalCardReplyContentDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch reply.Type {
	case components.ReplyTypeMarkdownReplyContentDto:
		// reply.MarkdownReplyContentDto is populated
	case components.ReplyTypeCardReplyContentDto:
		// reply.CardReplyContentDto is populated
	case components.ReplyTypeToolApprovalCardReplyContentDto:
		// reply.ToolApprovalCardReplyContentDto is populated
}
```
