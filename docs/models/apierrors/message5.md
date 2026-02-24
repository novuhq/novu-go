# Message5


## Supported Types

### 

```go
message5 := apierrors.CreateMessage5Str(string{/* values here */})
```

### 

```go
message5 := apierrors.CreateMessage5Number(float64{/* values here */})
```

### 

```go
message5 := apierrors.CreateMessage5Boolean(bool{/* values here */})
```

### 

```go
message5 := apierrors.CreateMessage5MapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch message5.Type {
	case apierrors.Message5TypeStr:
		// message5.Str is populated
	case apierrors.Message5TypeNumber:
		// message5.Number is populated
	case apierrors.Message5TypeBoolean:
		// message5.Boolean is populated
	case apierrors.Message5TypeMapOfAny:
		// message5.MapOfAny is populated
}
```
