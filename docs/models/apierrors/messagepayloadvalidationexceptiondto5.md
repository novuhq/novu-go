# MessagePayloadValidationExceptionDto5


## Supported Types

### 

```go
messagePayloadValidationExceptionDto5 := apierrors.CreateMessagePayloadValidationExceptionDto5Str(string{/* values here */})
```

### 

```go
messagePayloadValidationExceptionDto5 := apierrors.CreateMessagePayloadValidationExceptionDto5Number(float64{/* values here */})
```

### 

```go
messagePayloadValidationExceptionDto5 := apierrors.CreateMessagePayloadValidationExceptionDto5Boolean(bool{/* values here */})
```

### 

```go
messagePayloadValidationExceptionDto5 := apierrors.CreateMessagePayloadValidationExceptionDto5MapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch messagePayloadValidationExceptionDto5.Type {
	case apierrors.MessagePayloadValidationExceptionDto5TypeStr:
		// messagePayloadValidationExceptionDto5.Str is populated
	case apierrors.MessagePayloadValidationExceptionDto5TypeNumber:
		// messagePayloadValidationExceptionDto5.Number is populated
	case apierrors.MessagePayloadValidationExceptionDto5TypeBoolean:
		// messagePayloadValidationExceptionDto5.Boolean is populated
	case apierrors.MessagePayloadValidationExceptionDto5TypeMapOfAny:
		// messagePayloadValidationExceptionDto5.MapOfAny is populated
}
```
