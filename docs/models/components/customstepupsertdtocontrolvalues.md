# CustomStepUpsertDtoControlValues

Control values for the Custom step.


## Supported Types

### CustomControlDto

```go
customStepUpsertDtoControlValues := components.CreateCustomStepUpsertDtoControlValuesCustomControlDto(components.CustomControlDto{/* values here */})
```

### 

```go
customStepUpsertDtoControlValues := components.CreateCustomStepUpsertDtoControlValuesMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customStepUpsertDtoControlValues.Type {
	case components.CustomStepUpsertDtoControlValuesTypeCustomControlDto:
		// customStepUpsertDtoControlValues.CustomControlDto is populated
	case components.CustomStepUpsertDtoControlValuesTypeMapOfAny:
		// customStepUpsertDtoControlValues.MapOfAny is populated
}
```
