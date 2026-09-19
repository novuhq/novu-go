# TriggerSignalDtoTo

Recipient(s). Accepts a subscriberId string, subscriber object, topic object, or an array of those. When omitted, Novu falls back to the conversation subscriber.


## Supported Types

### 

```go
triggerSignalDtoTo := components.CreateTriggerSignalDtoToStr(string{/* values here */})
```

### 

```go
triggerSignalDtoTo := components.CreateTriggerSignalDtoToMapOfAny(map[string]any{/* values here */})
```

### 

```go
triggerSignalDtoTo := components.CreateTriggerSignalDtoToArrayOfTo3([]components.To3{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch triggerSignalDtoTo.Type {
	case components.TriggerSignalDtoToTypeStr:
		// triggerSignalDtoTo.Str is populated
	case components.TriggerSignalDtoToTypeMapOfAny:
		// triggerSignalDtoTo.MapOfAny is populated
	case components.TriggerSignalDtoToTypeArrayOfTo3:
		// triggerSignalDtoTo.ArrayOfTo3 is populated
}
```
