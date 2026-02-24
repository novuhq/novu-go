# DelayStepUpsertDtoControlValues

Control values for the Delay step.


## Supported Types

### DelayControlDto

```go
delayStepUpsertDtoControlValues := components.CreateDelayStepUpsertDtoControlValuesDelayControlDto(components.DelayControlDto{/* values here */})
```

### 

```go
delayStepUpsertDtoControlValues := components.CreateDelayStepUpsertDtoControlValuesMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch delayStepUpsertDtoControlValues.Type {
	case components.DelayStepUpsertDtoControlValuesTypeDelayControlDto:
		// delayStepUpsertDtoControlValues.DelayControlDto is populated
	case components.DelayStepUpsertDtoControlValuesTypeMapOfAny:
		// delayStepUpsertDtoControlValues.MapOfAny is populated
}
```
