# CreateMsTeamsUserEndpointDtoContext


## Supported Types

### 

```go
createMsTeamsUserEndpointDtoContext := components.CreateCreateMsTeamsUserEndpointDtoContextStr(string{/* values here */})
```

### CreateMsTeamsUserEndpointDtoContext2

```go
createMsTeamsUserEndpointDtoContext := components.CreateCreateMsTeamsUserEndpointDtoContextCreateMsTeamsUserEndpointDtoContext2(components.CreateMsTeamsUserEndpointDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createMsTeamsUserEndpointDtoContext.Type {
	case components.CreateMsTeamsUserEndpointDtoContextTypeStr:
		// createMsTeamsUserEndpointDtoContext.Str is populated
	case components.CreateMsTeamsUserEndpointDtoContextTypeCreateMsTeamsUserEndpointDtoContext2:
		// createMsTeamsUserEndpointDtoContext.CreateMsTeamsUserEndpointDtoContext2 is populated
}
```
