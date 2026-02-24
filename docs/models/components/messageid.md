# MessageID


## Supported Types

### 

```go
messageID := components.CreateMessageIDStr(string{/* values here */})
```

### 

```go
messageID := components.CreateMessageIDArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch messageID.Type {
	case components.MessageIDTypeStr:
		// messageID.Str is populated
	case components.MessageIDTypeArrayOfStr:
		// messageID.ArrayOfStr is populated
}
```
