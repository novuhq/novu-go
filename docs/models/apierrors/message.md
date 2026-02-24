# Message

Value that failed validation


## Supported Types

### 

```go
message := apierrors.CreateMessageStr(string{/* values here */})
```

### 

```go
message := apierrors.CreateMessageNumber(float64{/* values here */})
```

### 

```go
message := apierrors.CreateMessageBoolean(bool{/* values here */})
```

### Four

```go
message := apierrors.CreateMessageFour(apierrors.Four{/* values here */})
```

### 

```go
message := apierrors.CreateMessageArrayOf5([]*apierrors.Five{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch message.Type {
	case apierrors.MessageTypeStr:
		// message.Str is populated
	case apierrors.MessageTypeNumber:
		// message.Number is populated
	case apierrors.MessageTypeBoolean:
		// message.Boolean is populated
	case apierrors.MessageTypeFour:
		// message.Four is populated
	case apierrors.MessageTypeArrayOf5:
		// message.ArrayOf5 is populated
}
```
