# ValidationErrorDtoMessage

Value that failed validation


## Supported Types

### 

```go
validationErrorDtoMessage := apierrors.CreateValidationErrorDtoMessageStr(string{/* values here */})
```

### 

```go
validationErrorDtoMessage := apierrors.CreateValidationErrorDtoMessageNumber(float64{/* values here */})
```

### 

```go
validationErrorDtoMessage := apierrors.CreateValidationErrorDtoMessageBoolean(bool{/* values here */})
```

### Message4

```go
validationErrorDtoMessage := apierrors.CreateValidationErrorDtoMessageMessage4(apierrors.Message4{/* values here */})
```

### 

```go
validationErrorDtoMessage := apierrors.CreateValidationErrorDtoMessageArrayOfMessage5([]*apierrors.Message5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch validationErrorDtoMessage.Type {
	case apierrors.ValidationErrorDtoMessageTypeStr:
		// validationErrorDtoMessage.Str is populated
	case apierrors.ValidationErrorDtoMessageTypeNumber:
		// validationErrorDtoMessage.Number is populated
	case apierrors.ValidationErrorDtoMessageTypeBoolean:
		// validationErrorDtoMessage.Boolean is populated
	case apierrors.ValidationErrorDtoMessageTypeMessage4:
		// validationErrorDtoMessage.Message4 is populated
	case apierrors.ValidationErrorDtoMessageTypeArrayOfMessage5:
		// validationErrorDtoMessage.ArrayOfMessage5 is populated
}
```
