# SmsStepUpsertDtoControlValues

Control values for the SMS step.


## Supported Types

### SmsControlDto

```go
smsStepUpsertDtoControlValues := components.CreateSmsStepUpsertDtoControlValuesSmsControlDto(components.SmsControlDto{/* values here */})
```

### 

```go
smsStepUpsertDtoControlValues := components.CreateSmsStepUpsertDtoControlValuesMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch smsStepUpsertDtoControlValues.Type {
	case components.SmsStepUpsertDtoControlValuesTypeSmsControlDto:
		// smsStepUpsertDtoControlValues.SmsControlDto is populated
	case components.SmsStepUpsertDtoControlValuesTypeMapOfAny:
		// smsStepUpsertDtoControlValues.MapOfAny is populated
}
```
