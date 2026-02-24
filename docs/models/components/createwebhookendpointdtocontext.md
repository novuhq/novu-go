# CreateWebhookEndpointDtoContext


## Supported Types

### 

```go
createWebhookEndpointDtoContext := components.CreateCreateWebhookEndpointDtoContextStr(string{/* values here */})
```

### CreateWebhookEndpointDtoContext2

```go
createWebhookEndpointDtoContext := components.CreateCreateWebhookEndpointDtoContextCreateWebhookEndpointDtoContext2(components.CreateWebhookEndpointDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createWebhookEndpointDtoContext.Type {
	case components.CreateWebhookEndpointDtoContextTypeStr:
		// createWebhookEndpointDtoContext.Str is populated
	case components.CreateWebhookEndpointDtoContextTypeCreateWebhookEndpointDtoContext2:
		// createWebhookEndpointDtoContext.CreateWebhookEndpointDtoContext2 is populated
}
```
