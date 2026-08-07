# CreateOpsgenieIntegrationEndpointDtoContext


## Supported Types

### 

```go
createOpsgenieIntegrationEndpointDtoContext := components.CreateCreateOpsgenieIntegrationEndpointDtoContextStr(string{/* values here */})
```

### CreateOpsgenieIntegrationEndpointDtoContext2

```go
createOpsgenieIntegrationEndpointDtoContext := components.CreateCreateOpsgenieIntegrationEndpointDtoContextCreateOpsgenieIntegrationEndpointDtoContext2(components.CreateOpsgenieIntegrationEndpointDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createOpsgenieIntegrationEndpointDtoContext.Type {
	case components.CreateOpsgenieIntegrationEndpointDtoContextTypeStr:
		// createOpsgenieIntegrationEndpointDtoContext.Str is populated
	case components.CreateOpsgenieIntegrationEndpointDtoContextTypeCreateOpsgenieIntegrationEndpointDtoContext2:
		// createOpsgenieIntegrationEndpointDtoContext.CreateOpsgenieIntegrationEndpointDtoContext2 is populated
}
```
