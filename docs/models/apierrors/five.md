# Five


## Supported Types

### 

```go
five := apierrors.CreateFiveStr(string{/* values here */})
```

### 

```go
five := apierrors.CreateFiveNumber(float64{/* values here */})
```

### 

```go
five := apierrors.CreateFiveBoolean(bool{/* values here */})
```

### 

```go
five := apierrors.CreateFiveMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch five.Type {
	case apierrors.FiveTypeStr:
		// five.Str is populated
	case apierrors.FiveTypeNumber:
		// five.Number is populated
	case apierrors.FiveTypeBoolean:
		// five.Boolean is populated
	case apierrors.FiveTypeMapOfAny:
		// five.MapOfAny is populated
}
```
