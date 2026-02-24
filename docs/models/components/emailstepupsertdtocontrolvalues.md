# EmailStepUpsertDtoControlValues

Control values for the Email step.


## Supported Types

### EmailControlDto

```go
emailStepUpsertDtoControlValues := components.CreateEmailStepUpsertDtoControlValuesEmailControlDto(components.EmailControlDto{/* values here */})
```

### 

```go
emailStepUpsertDtoControlValues := components.CreateEmailStepUpsertDtoControlValuesMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch emailStepUpsertDtoControlValues.Type {
	case components.EmailStepUpsertDtoControlValuesTypeEmailControlDto:
		// emailStepUpsertDtoControlValues.EmailControlDto is populated
	case components.EmailStepUpsertDtoControlValuesTypeMapOfAny:
		// emailStepUpsertDtoControlValues.MapOfAny is populated
}
```
