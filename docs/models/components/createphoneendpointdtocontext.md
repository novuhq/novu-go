# CreatePhoneEndpointDtoContext


## Supported Types

### 

```go
createPhoneEndpointDtoContext := components.CreateCreatePhoneEndpointDtoContextStr(string{/* values here */})
```

### CreatePhoneEndpointDtoContext2

```go
createPhoneEndpointDtoContext := components.CreateCreatePhoneEndpointDtoContextCreatePhoneEndpointDtoContext2(components.CreatePhoneEndpointDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createPhoneEndpointDtoContext.Type {
	case components.CreatePhoneEndpointDtoContextTypeStr:
		// createPhoneEndpointDtoContext.Str is populated
	case components.CreatePhoneEndpointDtoContextTypeCreatePhoneEndpointDtoContext2:
		// createPhoneEndpointDtoContext.CreatePhoneEndpointDtoContext2 is populated
}
```
