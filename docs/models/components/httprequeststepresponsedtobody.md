# HTTPRequestStepResponseDtoBody

Request body as a raw JSON string. Key-value arrays are supported for legacy workflows.


## Supported Types

### 

```go
httpRequestStepResponseDtoBody := components.CreateHTTPRequestStepResponseDtoBodyStr(string{/* values here */})
```

### 

```go
httpRequestStepResponseDtoBody := components.CreateHTTPRequestStepResponseDtoBodyArrayOfHTTPRequestKeyValuePairDto([]components.HTTPRequestKeyValuePairDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch httpRequestStepResponseDtoBody.Type {
	case components.HTTPRequestStepResponseDtoBodyTypeStr:
		// httpRequestStepResponseDtoBody.Str is populated
	case components.HTTPRequestStepResponseDtoBodyTypeArrayOfHTTPRequestKeyValuePairDto:
		// httpRequestStepResponseDtoBody.ArrayOfHTTPRequestKeyValuePairDto is populated
}
```
