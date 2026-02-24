# PayloadValidationErrorDtoValue

The actual value that failed validation


## Supported Types

### 

```go
payloadValidationErrorDtoValue := components.CreatePayloadValidationErrorDtoValueStr(string{/* values here */})
```

### 

```go
payloadValidationErrorDtoValue := components.CreatePayloadValidationErrorDtoValueNumber(float64{/* values here */})
```

### 

```go
payloadValidationErrorDtoValue := components.CreatePayloadValidationErrorDtoValueBoolean(bool{/* values here */})
```

### Value4

```go
payloadValidationErrorDtoValue := components.CreatePayloadValidationErrorDtoValueValue4(components.Value4{/* values here */})
```

### 

```go
payloadValidationErrorDtoValue := components.CreatePayloadValidationErrorDtoValueArrayOfValue5([]*components.Value5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch payloadValidationErrorDtoValue.Type {
	case components.PayloadValidationErrorDtoValueTypeStr:
		// payloadValidationErrorDtoValue.Str is populated
	case components.PayloadValidationErrorDtoValueTypeNumber:
		// payloadValidationErrorDtoValue.Number is populated
	case components.PayloadValidationErrorDtoValueTypeBoolean:
		// payloadValidationErrorDtoValue.Boolean is populated
	case components.PayloadValidationErrorDtoValueTypeValue4:
		// payloadValidationErrorDtoValue.Value4 is populated
	case components.PayloadValidationErrorDtoValueTypeArrayOfValue5:
		// payloadValidationErrorDtoValue.ArrayOfValue5 is populated
}
```
