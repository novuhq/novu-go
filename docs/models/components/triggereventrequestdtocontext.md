# TriggerEventRequestDtoContext


## Supported Types

### 

```go
triggerEventRequestDtoContext := components.CreateTriggerEventRequestDtoContextStr(string{/* values here */})
```

### TriggerEventRequestDtoContext2

```go
triggerEventRequestDtoContext := components.CreateTriggerEventRequestDtoContextTriggerEventRequestDtoContext2(components.TriggerEventRequestDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch triggerEventRequestDtoContext.Type {
	case components.TriggerEventRequestDtoContextTypeStr:
		// triggerEventRequestDtoContext.Str is populated
	case components.TriggerEventRequestDtoContextTypeTriggerEventRequestDtoContext2:
		// triggerEventRequestDtoContext.TriggerEventRequestDtoContext2 is populated
}
```
