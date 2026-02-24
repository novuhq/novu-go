# PushStepUpsertDtoControlValues

Control values for the Push step.


## Supported Types

### PushControlDto

```go
pushStepUpsertDtoControlValues := components.CreatePushStepUpsertDtoControlValuesPushControlDto(components.PushControlDto{/* values here */})
```

### 

```go
pushStepUpsertDtoControlValues := components.CreatePushStepUpsertDtoControlValuesMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch pushStepUpsertDtoControlValues.Type {
	case components.PushStepUpsertDtoControlValuesTypePushControlDto:
		// pushStepUpsertDtoControlValues.PushControlDto is populated
	case components.PushStepUpsertDtoControlValuesTypeMapOfAny:
		// pushStepUpsertDtoControlValues.MapOfAny is populated
}
```
