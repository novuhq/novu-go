# CreateWebexPersonEndpointDtoContext


## Supported Types

### 

```go
createWebexPersonEndpointDtoContext := components.CreateCreateWebexPersonEndpointDtoContextStr(string{/* values here */})
```

### CreateWebexPersonEndpointDtoContext2

```go
createWebexPersonEndpointDtoContext := components.CreateCreateWebexPersonEndpointDtoContextCreateWebexPersonEndpointDtoContext2(components.CreateWebexPersonEndpointDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createWebexPersonEndpointDtoContext.Type {
	case components.CreateWebexPersonEndpointDtoContextTypeStr:
		// createWebexPersonEndpointDtoContext.Str is populated
	case components.CreateWebexPersonEndpointDtoContextTypeCreateWebexPersonEndpointDtoContext2:
		// createWebexPersonEndpointDtoContext.CreateWebexPersonEndpointDtoContext2 is populated
}
```
