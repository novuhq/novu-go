# GenerateLinkUserOauthURLRequestDtoContext


## Supported Types

### 

```go
generateLinkUserOauthURLRequestDtoContext := components.CreateGenerateLinkUserOauthURLRequestDtoContextStr(string{/* values here */})
```

### GenerateLinkUserOauthURLRequestDtoContext2

```go
generateLinkUserOauthURLRequestDtoContext := components.CreateGenerateLinkUserOauthURLRequestDtoContextGenerateLinkUserOauthURLRequestDtoContext2(components.GenerateLinkUserOauthURLRequestDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch generateLinkUserOauthURLRequestDtoContext.Type {
	case components.GenerateLinkUserOauthURLRequestDtoContextTypeStr:
		// generateLinkUserOauthURLRequestDtoContext.Str is populated
	case components.GenerateLinkUserOauthURLRequestDtoContextTypeGenerateLinkUserOauthURLRequestDtoContext2:
		// generateLinkUserOauthURLRequestDtoContext.GenerateLinkUserOauthURLRequestDtoContext2 is populated
}
```
