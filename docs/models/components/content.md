# Content

Content of the message, can be an email block or a string


## Supported Types

### 

```go
content := components.CreateContentArrayOfEmailBlock([]components.EmailBlock{/* values here */})
```

### 

```go
content := components.CreateContentStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch content.Type {
	case components.ContentTypeArrayOfEmailBlock:
		// content.ArrayOfEmailBlock is populated
	case components.ContentTypeStr:
		// content.Str is populated
}
```
