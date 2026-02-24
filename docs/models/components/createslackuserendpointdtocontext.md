# CreateSlackUserEndpointDtoContext


## Supported Types

### 

```go
createSlackUserEndpointDtoContext := components.CreateCreateSlackUserEndpointDtoContextStr(string{/* values here */})
```

### CreateSlackUserEndpointDtoContext2

```go
createSlackUserEndpointDtoContext := components.CreateCreateSlackUserEndpointDtoContextCreateSlackUserEndpointDtoContext2(components.CreateSlackUserEndpointDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createSlackUserEndpointDtoContext.Type {
	case components.CreateSlackUserEndpointDtoContextTypeStr:
		// createSlackUserEndpointDtoContext.Str is populated
	case components.CreateSlackUserEndpointDtoContextTypeCreateSlackUserEndpointDtoContext2:
		// createSlackUserEndpointDtoContext.CreateSlackUserEndpointDtoContext2 is populated
}
```
