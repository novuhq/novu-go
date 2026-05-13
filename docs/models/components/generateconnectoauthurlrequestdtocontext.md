# GenerateConnectOauthURLRequestDtoContext


## Supported Types

### 

```go
generateConnectOauthURLRequestDtoContext := components.CreateGenerateConnectOauthURLRequestDtoContextStr(string{/* values here */})
```

### GenerateConnectOauthURLRequestDtoContext2

```go
generateConnectOauthURLRequestDtoContext := components.CreateGenerateConnectOauthURLRequestDtoContextGenerateConnectOauthURLRequestDtoContext2(components.GenerateConnectOauthURLRequestDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch generateConnectOauthURLRequestDtoContext.Type {
	case components.GenerateConnectOauthURLRequestDtoContextTypeStr:
		// generateConnectOauthURLRequestDtoContext.Str is populated
	case components.GenerateConnectOauthURLRequestDtoContextTypeGenerateConnectOauthURLRequestDtoContext2:
		// generateConnectOauthURLRequestDtoContext.GenerateConnectOauthURLRequestDtoContext2 is populated
}
```
