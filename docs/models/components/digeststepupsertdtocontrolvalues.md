# DigestStepUpsertDtoControlValues

Control values for the Digest step.


## Supported Types

### DigestControlDto

```go
digestStepUpsertDtoControlValues := components.CreateDigestStepUpsertDtoControlValuesDigestControlDto(components.DigestControlDto{/* values here */})
```

### 

```go
digestStepUpsertDtoControlValues := components.CreateDigestStepUpsertDtoControlValuesMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch digestStepUpsertDtoControlValues.Type {
	case components.DigestStepUpsertDtoControlValuesTypeDigestControlDto:
		// digestStepUpsertDtoControlValues.DigestControlDto is populated
	case components.DigestStepUpsertDtoControlValuesTypeMapOfAny:
		// digestStepUpsertDtoControlValues.MapOfAny is populated
}
```
