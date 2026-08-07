# To3


## Supported Types

### 

```go
to3 := components.CreateTo3Str(string{/* values here */})
```

### 

```go
to3 := components.CreateTo3MapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch to3.Type {
	case components.To3TypeStr:
		// to3.Str is populated
	case components.To3TypeMapOfAny:
		// to3.MapOfAny is populated
}
```
