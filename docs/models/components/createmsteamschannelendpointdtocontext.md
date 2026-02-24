# CreateMsTeamsChannelEndpointDtoContext


## Supported Types

### 

```go
createMsTeamsChannelEndpointDtoContext := components.CreateCreateMsTeamsChannelEndpointDtoContextStr(string{/* values here */})
```

### CreateMsTeamsChannelEndpointDtoContext2

```go
createMsTeamsChannelEndpointDtoContext := components.CreateCreateMsTeamsChannelEndpointDtoContextCreateMsTeamsChannelEndpointDtoContext2(components.CreateMsTeamsChannelEndpointDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createMsTeamsChannelEndpointDtoContext.Type {
	case components.CreateMsTeamsChannelEndpointDtoContextTypeStr:
		// createMsTeamsChannelEndpointDtoContext.Str is populated
	case components.CreateMsTeamsChannelEndpointDtoContextTypeCreateMsTeamsChannelEndpointDtoContext2:
		// createMsTeamsChannelEndpointDtoContext.CreateMsTeamsChannelEndpointDtoContext2 is populated
}
```
