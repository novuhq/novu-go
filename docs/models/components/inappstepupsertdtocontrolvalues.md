# InAppStepUpsertDtoControlValues

Control values for the In-App step.


## Supported Types

### InAppControlDto

```go
inAppStepUpsertDtoControlValues := components.CreateInAppStepUpsertDtoControlValuesInAppControlDto(components.InAppControlDto{/* values here */})
```

### 

```go
inAppStepUpsertDtoControlValues := components.CreateInAppStepUpsertDtoControlValuesMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inAppStepUpsertDtoControlValues.Type {
	case components.InAppStepUpsertDtoControlValuesTypeInAppControlDto:
		// inAppStepUpsertDtoControlValues.InAppControlDto is populated
	case components.InAppStepUpsertDtoControlValuesTypeMapOfAny:
		// inAppStepUpsertDtoControlValues.MapOfAny is populated
}
```
