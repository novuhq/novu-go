# Result

Preview result


## Supported Types

### One

```go
result := components.CreateResultOne(components.One{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch result.Type {
	case components.ResultUnionTypeOne:
		// result.One is populated
}
```
