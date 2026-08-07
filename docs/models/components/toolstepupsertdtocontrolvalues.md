# ToolStepUpsertDtoControlValues

Control values for the Tool step.


## Supported Types

### ToolControlDto

```go
toolStepUpsertDtoControlValues := components.CreateToolStepUpsertDtoControlValuesToolControlDto(components.ToolControlDto{/* values here */})
```

### 

```go
toolStepUpsertDtoControlValues := components.CreateToolStepUpsertDtoControlValuesMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch toolStepUpsertDtoControlValues.Type {
	case components.ToolStepUpsertDtoControlValuesTypeToolControlDto:
		// toolStepUpsertDtoControlValues.ToolControlDto is populated
	case components.ToolStepUpsertDtoControlValuesTypeMapOfAny:
		// toolStepUpsertDtoControlValues.MapOfAny is populated
}
```
