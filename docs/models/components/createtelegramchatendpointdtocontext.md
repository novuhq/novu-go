# CreateTelegramChatEndpointDtoContext


## Supported Types

### 

```go
createTelegramChatEndpointDtoContext := components.CreateCreateTelegramChatEndpointDtoContextStr(string{/* values here */})
```

### CreateTelegramChatEndpointDtoContext2

```go
createTelegramChatEndpointDtoContext := components.CreateCreateTelegramChatEndpointDtoContextCreateTelegramChatEndpointDtoContext2(components.CreateTelegramChatEndpointDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createTelegramChatEndpointDtoContext.Type {
	case components.CreateTelegramChatEndpointDtoContextTypeStr:
		// createTelegramChatEndpointDtoContext.Str is populated
	case components.CreateTelegramChatEndpointDtoContextTypeCreateTelegramChatEndpointDtoContext2:
		// createTelegramChatEndpointDtoContext.CreateTelegramChatEndpointDtoContext2 is populated
}
```
