# CreateChannelConnectionRequestDtoContext


## Supported Types

### 

```go
createChannelConnectionRequestDtoContext := components.CreateCreateChannelConnectionRequestDtoContextStr(string{/* values here */})
```

### CreateChannelConnectionRequestDtoContext2

```go
createChannelConnectionRequestDtoContext := components.CreateCreateChannelConnectionRequestDtoContextCreateChannelConnectionRequestDtoContext2(components.CreateChannelConnectionRequestDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createChannelConnectionRequestDtoContext.Type {
	case components.CreateChannelConnectionRequestDtoContextTypeStr:
		// createChannelConnectionRequestDtoContext.Str is populated
	case components.CreateChannelConnectionRequestDtoContextTypeCreateChannelConnectionRequestDtoContext2:
		// createChannelConnectionRequestDtoContext.CreateChannelConnectionRequestDtoContext2 is populated
}
```
