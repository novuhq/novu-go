# MessageResponseDtoContent

Content of the message, can be an email block or a string


## Supported Types

### 

```go
messageResponseDtoContent := components.CreateMessageResponseDtoContentArrayOfEmailBlock([]components.EmailBlock{/* values here */})
```

### 

```go
messageResponseDtoContent := components.CreateMessageResponseDtoContentStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch messageResponseDtoContent.Type {
	case components.MessageResponseDtoContentTypeArrayOfEmailBlock:
		// messageResponseDtoContent.ArrayOfEmailBlock is populated
	case components.MessageResponseDtoContentTypeStr:
		// messageResponseDtoContent.Str is populated
}
```
