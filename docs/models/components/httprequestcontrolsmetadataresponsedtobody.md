# HTTPRequestControlsMetadataResponseDtoBody

Request body as a raw JSON string. Key-value arrays are supported for legacy workflows.


## Supported Types

### 

```go
httpRequestControlsMetadataResponseDtoBody := components.CreateHTTPRequestControlsMetadataResponseDtoBodyStr(string{/* values here */})
```

### 

```go
httpRequestControlsMetadataResponseDtoBody := components.CreateHTTPRequestControlsMetadataResponseDtoBodyArrayOfHTTPRequestKeyValuePairDto([]components.HTTPRequestKeyValuePairDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch httpRequestControlsMetadataResponseDtoBody.Type {
	case components.HTTPRequestControlsMetadataResponseDtoBodyTypeStr:
		// httpRequestControlsMetadataResponseDtoBody.Str is populated
	case components.HTTPRequestControlsMetadataResponseDtoBodyTypeArrayOfHTTPRequestKeyValuePairDto:
		// httpRequestControlsMetadataResponseDtoBody.ArrayOfHTTPRequestKeyValuePairDto is populated
}
```
