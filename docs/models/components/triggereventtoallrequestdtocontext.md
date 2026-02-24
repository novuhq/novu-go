# TriggerEventToAllRequestDtoContext


## Supported Types

### 

```go
triggerEventToAllRequestDtoContext := components.CreateTriggerEventToAllRequestDtoContextStr(string{/* values here */})
```

### TriggerEventToAllRequestDtoContext2

```go
triggerEventToAllRequestDtoContext := components.CreateTriggerEventToAllRequestDtoContextTriggerEventToAllRequestDtoContext2(components.TriggerEventToAllRequestDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch triggerEventToAllRequestDtoContext.Type {
	case components.TriggerEventToAllRequestDtoContextTypeStr:
		// triggerEventToAllRequestDtoContext.Str is populated
	case components.TriggerEventToAllRequestDtoContextTypeTriggerEventToAllRequestDtoContext2:
		// triggerEventToAllRequestDtoContext.TriggerEventToAllRequestDtoContext2 is populated
}
```
