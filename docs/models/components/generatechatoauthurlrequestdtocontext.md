# GenerateChatOauthURLRequestDtoContext


## Supported Types

### 

```go
generateChatOauthURLRequestDtoContext := components.CreateGenerateChatOauthURLRequestDtoContextStr(string{/* values here */})
```

### GenerateChatOauthURLRequestDtoContext2

```go
generateChatOauthURLRequestDtoContext := components.CreateGenerateChatOauthURLRequestDtoContextGenerateChatOauthURLRequestDtoContext2(components.GenerateChatOauthURLRequestDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch generateChatOauthURLRequestDtoContext.Type {
	case components.GenerateChatOauthURLRequestDtoContextTypeStr:
		// generateChatOauthURLRequestDtoContext.Str is populated
	case components.GenerateChatOauthURLRequestDtoContextTypeGenerateChatOauthURLRequestDtoContext2:
		// generateChatOauthURLRequestDtoContext.GenerateChatOauthURLRequestDtoContext2 is populated
}
```
