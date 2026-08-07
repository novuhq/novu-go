# CreateWebexRoomEndpointDtoContext


## Supported Types

### 

```go
createWebexRoomEndpointDtoContext := components.CreateCreateWebexRoomEndpointDtoContextStr(string{/* values here */})
```

### CreateWebexRoomEndpointDtoContext2

```go
createWebexRoomEndpointDtoContext := components.CreateCreateWebexRoomEndpointDtoContextCreateWebexRoomEndpointDtoContext2(components.CreateWebexRoomEndpointDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createWebexRoomEndpointDtoContext.Type {
	case components.CreateWebexRoomEndpointDtoContextTypeStr:
		// createWebexRoomEndpointDtoContext.Str is populated
	case components.CreateWebexRoomEndpointDtoContextTypeCreateWebexRoomEndpointDtoContext2:
		// createWebexRoomEndpointDtoContext.CreateWebexRoomEndpointDtoContext2 is populated
}
```
