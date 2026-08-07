# Content

Replacement content. Exactly one of markdown, card, or toolApprovalCard.


## Supported Types

### MarkdownReplyContentDto

```go
content := components.CreateContentMarkdownReplyContentDto(components.MarkdownReplyContentDto{/* values here */})
```

### CardReplyContentDto

```go
content := components.CreateContentCardReplyContentDto(components.CardReplyContentDto{/* values here */})
```

### ToolApprovalCardReplyContentDto

```go
content := components.CreateContentToolApprovalCardReplyContentDto(components.ToolApprovalCardReplyContentDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch content.Type {
	case components.ContentTypeMarkdownReplyContentDto:
		// content.MarkdownReplyContentDto is populated
	case components.ContentTypeCardReplyContentDto:
		// content.CardReplyContentDto is populated
	case components.ContentTypeToolApprovalCardReplyContentDto:
		// content.ToolApprovalCardReplyContentDto is populated
}
```
