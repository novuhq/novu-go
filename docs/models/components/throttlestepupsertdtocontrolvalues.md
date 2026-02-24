# ThrottleStepUpsertDtoControlValues

Control values for the Throttle step.


## Supported Types

### ThrottleControlDto

```go
throttleStepUpsertDtoControlValues := components.CreateThrottleStepUpsertDtoControlValuesThrottleControlDto(components.ThrottleControlDto{/* values here */})
```

### 

```go
throttleStepUpsertDtoControlValues := components.CreateThrottleStepUpsertDtoControlValuesMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch throttleStepUpsertDtoControlValues.Type {
	case components.ThrottleStepUpsertDtoControlValuesTypeThrottleControlDto:
		// throttleStepUpsertDtoControlValues.ThrottleControlDto is populated
	case components.ThrottleStepUpsertDtoControlValuesTypeMapOfAny:
		// throttleStepUpsertDtoControlValues.MapOfAny is populated
}
```
