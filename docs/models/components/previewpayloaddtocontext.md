# PreviewPayloadDtoContext


## Supported Types

### 

```go
previewPayloadDtoContext := components.CreatePreviewPayloadDtoContextStr(string{/* values here */})
```

### PreviewPayloadDtoContext2

```go
previewPayloadDtoContext := components.CreatePreviewPayloadDtoContextPreviewPayloadDtoContext2(components.PreviewPayloadDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch previewPayloadDtoContext.Type {
	case components.PreviewPayloadDtoContextTypeStr:
		// previewPayloadDtoContext.Str is populated
	case components.PreviewPayloadDtoContextTypePreviewPayloadDtoContext2:
		// previewPayloadDtoContext.PreviewPayloadDtoContext2 is populated
}
```
