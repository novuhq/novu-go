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

