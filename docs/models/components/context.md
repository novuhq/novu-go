# Context


## Supported Types

### 

```go
context := components.CreateContextStr(string{/* values here */})
```

### Two

```go
context := components.CreateContextTwo(components.Two{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch context.Type {
	case components.ContextTypeStr:
		// context.Str is populated
	case components.ContextTypeTwo:
		// context.Two is populated
}
```
