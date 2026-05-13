# Body

Request body as a raw JSON string. Key-value arrays are supported for legacy workflows.


## Supported Types

### 

```go
body := components.CreateBodyStr(string{/* values here */})
```

### 

```go
body := components.CreateBodyArrayOfHTTPRequestKeyValuePairDto([]components.HTTPRequestKeyValuePairDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch body.Type {
	case components.BodyTypeStr:
		// body.Str is populated
	case components.BodyTypeArrayOfHTTPRequestKeyValuePairDto:
		// body.ArrayOfHTTPRequestKeyValuePairDto is populated
}
```
