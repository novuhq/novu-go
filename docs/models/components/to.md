# To

Recipient(s). Accepts a subscriberId string, subscriber object, topic object, or an array of those. When omitted, Novu falls back to the conversation subscriber.


## Supported Types

### 

```go
to := components.CreateToStr(string{/* values here */})
```

### 

```go
to := components.CreateToMapOfAny(map[string]any{/* values here */})
```

### 

```go
to := components.CreateToArrayOfTo3([]components.To3{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch to.Type {
	case components.ToTypeStr:
		// to.Str is populated
	case components.ToTypeMapOfAny:
		// to.MapOfAny is populated
	case components.ToTypeArrayOfTo3:
		// to.ArrayOfTo3 is populated
}
```
