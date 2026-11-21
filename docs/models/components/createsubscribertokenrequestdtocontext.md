# CreateSubscriberTokenRequestDtoContext


## Supported Types

### 

```go
createSubscriberTokenRequestDtoContext := components.CreateCreateSubscriberTokenRequestDtoContextStr(string{/* values here */})
```

### CreateSubscriberTokenRequestDtoContext2

```go
createSubscriberTokenRequestDtoContext := components.CreateCreateSubscriberTokenRequestDtoContextCreateSubscriberTokenRequestDtoContext2(components.CreateSubscriberTokenRequestDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createSubscriberTokenRequestDtoContext.Type {
	case components.CreateSubscriberTokenRequestDtoContextTypeStr:
		// createSubscriberTokenRequestDtoContext.Str is populated
	case components.CreateSubscriberTokenRequestDtoContextTypeCreateSubscriberTokenRequestDtoContext2:
		// createSubscriberTokenRequestDtoContext.CreateSubscriberTokenRequestDtoContext2 is populated
}
```
