# HTTPRequestStepUpsertDtoControlValues

Control values for the HTTP Request step.


## Supported Types

### HTTPRequestControlDto

```go
httpRequestStepUpsertDtoControlValues := components.CreateHTTPRequestStepUpsertDtoControlValuesHTTPRequestControlDto(components.HTTPRequestControlDto{/* values here */})
```

### 

```go
httpRequestStepUpsertDtoControlValues := components.CreateHTTPRequestStepUpsertDtoControlValuesMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch httpRequestStepUpsertDtoControlValues.Type {
	case components.HTTPRequestStepUpsertDtoControlValuesTypeHTTPRequestControlDto:
		// httpRequestStepUpsertDtoControlValues.HTTPRequestControlDto is populated
	case components.HTTPRequestStepUpsertDtoControlValuesTypeMapOfAny:
		// httpRequestStepUpsertDtoControlValues.MapOfAny is populated
}
```
