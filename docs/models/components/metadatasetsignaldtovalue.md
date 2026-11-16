# MetadataSetSignalDtoValue

JSON-serializable value to store.


## Supported Types

### 

```go
metadataSetSignalDtoValue := components.CreateMetadataSetSignalDtoValueStr(string{/* values here */})
```

### 

```go
metadataSetSignalDtoValue := components.CreateMetadataSetSignalDtoValueNumber(float64{/* values here */})
```

### 

```go
metadataSetSignalDtoValue := components.CreateMetadataSetSignalDtoValueBoolean(bool{/* values here */})
```

### 

```go
metadataSetSignalDtoValue := components.CreateMetadataSetSignalDtoValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
metadataSetSignalDtoValue := components.CreateMetadataSetSignalDtoValueArrayOfAny([]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metadataSetSignalDtoValue.Type {
	case components.MetadataSetSignalDtoValueTypeStr:
		// metadataSetSignalDtoValue.Str is populated
	case components.MetadataSetSignalDtoValueTypeNumber:
		// metadataSetSignalDtoValue.Number is populated
	case components.MetadataSetSignalDtoValueTypeBoolean:
		// metadataSetSignalDtoValue.Boolean is populated
	case components.MetadataSetSignalDtoValueTypeMapOfAny:
		// metadataSetSignalDtoValue.MapOfAny is populated
	case components.MetadataSetSignalDtoValueTypeArrayOfAny:
		// metadataSetSignalDtoValue.ArrayOfAny is populated
}
```
