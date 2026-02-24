# Placeholder

Placeholder for the UI Schema Property


## Supported Types

### 

```go
placeholder := components.CreatePlaceholderStr(string{/* values here */})
```

### 

```go
placeholder := components.CreatePlaceholderNumber(float64{/* values here */})
```

### 

```go
placeholder := components.CreatePlaceholderBoolean(bool{/* values here */})
```

### 

```go
placeholder := components.CreatePlaceholderMapOfAny(map[string]any{/* values here */})
```

### 

```go
placeholder := components.CreatePlaceholderArrayOfPlaceholder5([]components.Placeholder5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch placeholder.Type {
	case components.PlaceholderTypeStr:
		// placeholder.Str is populated
	case components.PlaceholderTypeNumber:
		// placeholder.Number is populated
	case components.PlaceholderTypeBoolean:
		// placeholder.Boolean is populated
	case components.PlaceholderTypeMapOfAny:
		// placeholder.MapOfAny is populated
	case components.PlaceholderTypeArrayOfPlaceholder5:
		// placeholder.ArrayOfPlaceholder5 is populated
}
```
