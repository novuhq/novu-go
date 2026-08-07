# CreateLineUserEndpointDtoContext


## Supported Types

### 

```go
createLineUserEndpointDtoContext := components.CreateCreateLineUserEndpointDtoContextStr(string{/* values here */})
```

### CreateLineUserEndpointDtoContext2

```go
createLineUserEndpointDtoContext := components.CreateCreateLineUserEndpointDtoContextCreateLineUserEndpointDtoContext2(components.CreateLineUserEndpointDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createLineUserEndpointDtoContext.Type {
	case components.CreateLineUserEndpointDtoContextTypeStr:
		// createLineUserEndpointDtoContext.Str is populated
	case components.CreateLineUserEndpointDtoContextTypeCreateLineUserEndpointDtoContext2:
		// createLineUserEndpointDtoContext.CreateLineUserEndpointDtoContext2 is populated
}
```
