# CreateGrafanaOnCallIntegrationEndpointDtoContext


## Supported Types

### 

```go
createGrafanaOnCallIntegrationEndpointDtoContext := components.CreateCreateGrafanaOnCallIntegrationEndpointDtoContextStr(string{/* values here */})
```

### CreateGrafanaOnCallIntegrationEndpointDtoContext2

```go
createGrafanaOnCallIntegrationEndpointDtoContext := components.CreateCreateGrafanaOnCallIntegrationEndpointDtoContextCreateGrafanaOnCallIntegrationEndpointDtoContext2(components.CreateGrafanaOnCallIntegrationEndpointDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createGrafanaOnCallIntegrationEndpointDtoContext.Type {
	case components.CreateGrafanaOnCallIntegrationEndpointDtoContextTypeStr:
		// createGrafanaOnCallIntegrationEndpointDtoContext.Str is populated
	case components.CreateGrafanaOnCallIntegrationEndpointDtoContextTypeCreateGrafanaOnCallIntegrationEndpointDtoContext2:
		// createGrafanaOnCallIntegrationEndpointDtoContext.CreateGrafanaOnCallIntegrationEndpointDtoContext2 is populated
}
```
