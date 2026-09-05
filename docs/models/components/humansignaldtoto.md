# HumanSignalDtoTo

Novu subscriberId(s) allowed to settle this interaction (max 50). First valid answer wins. When omitted, the conversation subscriber is used. Subscriber ids only — not workflow topic recipients.


## Supported Types

### 

```go
humanSignalDtoTo := components.CreateHumanSignalDtoToStr(string{/* values here */})
```

### 

```go
humanSignalDtoTo := components.CreateHumanSignalDtoToArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch humanSignalDtoTo.Type {
	case components.HumanSignalDtoToTypeStr:
		// humanSignalDtoTo.Str is populated
	case components.HumanSignalDtoToTypeArrayOfStr:
		// humanSignalDtoTo.ArrayOfStr is populated
}
```
