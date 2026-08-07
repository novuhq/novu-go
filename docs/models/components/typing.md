# Typing

Per-turn typing/status control. Pass `{ status?: string }` to set/update the status (omit `status` for "Thinking…"), or `"stop"` to clear it. Best-effort per platform.


## Supported Types

### Typing1

```go
typing := components.CreateTypingTyping1(components.Typing1{/* values here */})
```

### TypingStatusDto

```go
typing := components.CreateTypingTypingStatusDto(components.TypingStatusDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch typing.Type {
	case components.TypingTypeTyping1:
		// typing.Typing1 is populated
	case components.TypingTypeTypingStatusDto:
		// typing.TypingStatusDto is populated
}
```
