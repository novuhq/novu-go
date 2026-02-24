# Five


## Supported Types

### 

```go
five := components.CreateFiveStr(string{/* values here */})
```

### 

```go
five := components.CreateFiveNumber(float64{/* values here */})
```

### 

```go
five := components.CreateFiveBoolean(bool{/* values here */})
```

### 

```go
five := components.CreateFiveMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch five.Type {
	case components.FiveTypeStr:
		// five.Str is populated
	case components.FiveTypeNumber:
		// five.Number is populated
	case components.FiveTypeBoolean:
		// five.Boolean is populated
	case components.FiveTypeMapOfAny:
		// five.MapOfAny is populated
}
```
