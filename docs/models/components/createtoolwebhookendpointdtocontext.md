# CreateToolWebhookEndpointDtoContext


## Supported Types

### 

```go
createToolWebhookEndpointDtoContext := components.CreateCreateToolWebhookEndpointDtoContextStr(string{/* values here */})
```

### CreateToolWebhookEndpointDtoContext2

```go
createToolWebhookEndpointDtoContext := components.CreateCreateToolWebhookEndpointDtoContextCreateToolWebhookEndpointDtoContext2(components.CreateToolWebhookEndpointDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createToolWebhookEndpointDtoContext.Type {
	case components.CreateToolWebhookEndpointDtoContextTypeStr:
		// createToolWebhookEndpointDtoContext.Str is populated
	case components.CreateToolWebhookEndpointDtoContextTypeCreateToolWebhookEndpointDtoContext2:
		// createToolWebhookEndpointDtoContext.CreateToolWebhookEndpointDtoContext2 is populated
}
```
