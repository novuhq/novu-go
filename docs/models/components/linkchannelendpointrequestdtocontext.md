# LinkChannelEndpointRequestDtoContext


## Supported Types

### 

```go
linkChannelEndpointRequestDtoContext := components.CreateLinkChannelEndpointRequestDtoContextStr(string{/* values here */})
```

### LinkChannelEndpointRequestDtoContext2

```go
linkChannelEndpointRequestDtoContext := components.CreateLinkChannelEndpointRequestDtoContextLinkChannelEndpointRequestDtoContext2(components.LinkChannelEndpointRequestDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch linkChannelEndpointRequestDtoContext.Type {
	case components.LinkChannelEndpointRequestDtoContextTypeStr:
		// linkChannelEndpointRequestDtoContext.Str is populated
	case components.LinkChannelEndpointRequestDtoContextTypeLinkChannelEndpointRequestDtoContext2:
		// linkChannelEndpointRequestDtoContext.LinkChannelEndpointRequestDtoContext2 is populated
}
```
