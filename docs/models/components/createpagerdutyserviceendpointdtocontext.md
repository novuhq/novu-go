# CreatePagerDutyServiceEndpointDtoContext


## Supported Types

### 

```go
createPagerDutyServiceEndpointDtoContext := components.CreateCreatePagerDutyServiceEndpointDtoContextStr(string{/* values here */})
```

### CreatePagerDutyServiceEndpointDtoContext2

```go
createPagerDutyServiceEndpointDtoContext := components.CreateCreatePagerDutyServiceEndpointDtoContextCreatePagerDutyServiceEndpointDtoContext2(components.CreatePagerDutyServiceEndpointDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createPagerDutyServiceEndpointDtoContext.Type {
	case components.CreatePagerDutyServiceEndpointDtoContextTypeStr:
		// createPagerDutyServiceEndpointDtoContext.Str is populated
	case components.CreatePagerDutyServiceEndpointDtoContextTypeCreatePagerDutyServiceEndpointDtoContext2:
		// createPagerDutyServiceEndpointDtoContext.CreatePagerDutyServiceEndpointDtoContext2 is populated
}
```
