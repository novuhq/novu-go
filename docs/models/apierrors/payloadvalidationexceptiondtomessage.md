# PayloadValidationExceptionDtoMessage

Value that failed validation


## Supported Types

### 

```go
payloadValidationExceptionDtoMessage := apierrors.CreatePayloadValidationExceptionDtoMessageStr(string{/* values here */})
```

### 

```go
payloadValidationExceptionDtoMessage := apierrors.CreatePayloadValidationExceptionDtoMessageNumber(float64{/* values here */})
```

### 

```go
payloadValidationExceptionDtoMessage := apierrors.CreatePayloadValidationExceptionDtoMessageBoolean(bool{/* values here */})
```

### MessagePayloadValidationExceptionDto4

```go
payloadValidationExceptionDtoMessage := apierrors.CreatePayloadValidationExceptionDtoMessageMessagePayloadValidationExceptionDto4(apierrors.MessagePayloadValidationExceptionDto4{/* values here */})
```

### 

```go
payloadValidationExceptionDtoMessage := apierrors.CreatePayloadValidationExceptionDtoMessageArrayOfMessagePayloadValidationExceptionDto5([]*apierrors.MessagePayloadValidationExceptionDto5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch payloadValidationExceptionDtoMessage.Type {
	case apierrors.PayloadValidationExceptionDtoMessageTypeStr:
		// payloadValidationExceptionDtoMessage.Str is populated
	case apierrors.PayloadValidationExceptionDtoMessageTypeNumber:
		// payloadValidationExceptionDtoMessage.Number is populated
	case apierrors.PayloadValidationExceptionDtoMessageTypeBoolean:
		// payloadValidationExceptionDtoMessage.Boolean is populated
	case apierrors.PayloadValidationExceptionDtoMessageTypeMessagePayloadValidationExceptionDto4:
		// payloadValidationExceptionDtoMessage.MessagePayloadValidationExceptionDto4 is populated
	case apierrors.PayloadValidationExceptionDtoMessageTypeArrayOfMessagePayloadValidationExceptionDto5:
		// payloadValidationExceptionDtoMessage.ArrayOfMessagePayloadValidationExceptionDto5 is populated
}
```
