# ControlValues

Control values for the In-App step.


## Supported Types

### InAppControlDto

```go
controlValues := components.CreateControlValuesInAppControlDto(components.InAppControlDto{/* values here */})
```

### 

```go
controlValues := components.CreateControlValuesMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch controlValues.Type {
	case components.ControlValuesTypeInAppControlDto:
		// controlValues.InAppControlDto is populated
	case components.ControlValuesTypeMapOfAny:
		// controlValues.MapOfAny is populated
}
```
