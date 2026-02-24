# CreateSlackChannelEndpointDtoContext


## Supported Types

### 

```go
createSlackChannelEndpointDtoContext := components.CreateCreateSlackChannelEndpointDtoContextStr(string{/* values here */})
```

### CreateSlackChannelEndpointDtoContext2

```go
createSlackChannelEndpointDtoContext := components.CreateCreateSlackChannelEndpointDtoContextCreateSlackChannelEndpointDtoContext2(components.CreateSlackChannelEndpointDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createSlackChannelEndpointDtoContext.Type {
	case components.CreateSlackChannelEndpointDtoContextTypeStr:
		// createSlackChannelEndpointDtoContext.Str is populated
	case components.CreateSlackChannelEndpointDtoContextTypeCreateSlackChannelEndpointDtoContext2:
		// createSlackChannelEndpointDtoContext.CreateSlackChannelEndpointDtoContext2 is populated
}
```
