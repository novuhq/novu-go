# Placeholder5


## Supported Types

### 

```go
placeholder5 := components.CreatePlaceholder5Str(string{/* values here */})
```

### 

```go
placeholder5 := components.CreatePlaceholder5Number(float64{/* values here */})
```

### 

```go
placeholder5 := components.CreatePlaceholder5Boolean(bool{/* values here */})
```

### 

```go
placeholder5 := components.CreatePlaceholder5MapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch placeholder5.Type {
	case components.Placeholder5TypeStr:
		// placeholder5.Str is populated
	case components.Placeholder5TypeNumber:
		// placeholder5.Number is populated
	case components.Placeholder5TypeBoolean:
		// placeholder5.Boolean is populated
	case components.Placeholder5TypeMapOfAny:
		// placeholder5.MapOfAny is populated
}
```
