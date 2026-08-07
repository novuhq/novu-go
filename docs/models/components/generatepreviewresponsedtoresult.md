# GeneratePreviewResponseDtoResult

Preview result


## Supported Types

### 

```go
generatePreviewResponseDtoResult := components.CreateGeneratePreviewResponseDtoResultMapOfAny(map[string]any{/* values here */})
```

### Result2

```go
generatePreviewResponseDtoResult := components.CreateGeneratePreviewResponseDtoResultResult2(components.Result2{/* values here */})
```

### Three

```go
generatePreviewResponseDtoResult := components.CreateGeneratePreviewResponseDtoResultThree(components.Three{/* values here */})
```

### Result4

```go
generatePreviewResponseDtoResult := components.CreateGeneratePreviewResponseDtoResultResult4(components.Result4{/* values here */})
```

### Result5

```go
generatePreviewResponseDtoResult := components.CreateGeneratePreviewResponseDtoResultResult5(components.Result5{/* values here */})
```

### Six

```go
generatePreviewResponseDtoResult := components.CreateGeneratePreviewResponseDtoResultSix(components.Six{/* values here */})
```

### Seven

```go
generatePreviewResponseDtoResult := components.CreateGeneratePreviewResponseDtoResultSeven(components.Seven{/* values here */})
```

### Eight

```go
generatePreviewResponseDtoResult := components.CreateGeneratePreviewResponseDtoResultEight(components.Eight{/* values here */})
```

### Nine

```go
generatePreviewResponseDtoResult := components.CreateGeneratePreviewResponseDtoResultNine(components.Nine{/* values here */})
```

### Ten

```go
generatePreviewResponseDtoResult := components.CreateGeneratePreviewResponseDtoResultTen(components.Ten{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch generatePreviewResponseDtoResult.Type {
	case components.GeneratePreviewResponseDtoResultUnionTypeMapOfAny:
		// generatePreviewResponseDtoResult.MapOfAny is populated
	case components.GeneratePreviewResponseDtoResultUnionTypeResult2:
		// generatePreviewResponseDtoResult.Result2 is populated
	case components.GeneratePreviewResponseDtoResultUnionTypeThree:
		// generatePreviewResponseDtoResult.Three is populated
	case components.GeneratePreviewResponseDtoResultUnionTypeResult4:
		// generatePreviewResponseDtoResult.Result4 is populated
	case components.GeneratePreviewResponseDtoResultUnionTypeResult5:
		// generatePreviewResponseDtoResult.Result5 is populated
	case components.GeneratePreviewResponseDtoResultUnionTypeSix:
		// generatePreviewResponseDtoResult.Six is populated
	case components.GeneratePreviewResponseDtoResultUnionTypeSeven:
		// generatePreviewResponseDtoResult.Seven is populated
	case components.GeneratePreviewResponseDtoResultUnionTypeEight:
		// generatePreviewResponseDtoResult.Eight is populated
	case components.GeneratePreviewResponseDtoResultUnionTypeNine:
		// generatePreviewResponseDtoResult.Nine is populated
	case components.GeneratePreviewResponseDtoResultUnionTypeTen:
		// generatePreviewResponseDtoResult.Ten is populated
}
```
