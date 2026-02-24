# CreateTopicSubscriptionsRequestDtoContext


## Supported Types

### 

```go
createTopicSubscriptionsRequestDtoContext := components.CreateCreateTopicSubscriptionsRequestDtoContextStr(string{/* values here */})
```

### CreateTopicSubscriptionsRequestDtoContext2

```go
createTopicSubscriptionsRequestDtoContext := components.CreateCreateTopicSubscriptionsRequestDtoContextCreateTopicSubscriptionsRequestDtoContext2(components.CreateTopicSubscriptionsRequestDtoContext2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createTopicSubscriptionsRequestDtoContext.Type {
	case components.CreateTopicSubscriptionsRequestDtoContextTypeStr:
		// createTopicSubscriptionsRequestDtoContext.Str is populated
	case components.CreateTopicSubscriptionsRequestDtoContextTypeCreateTopicSubscriptionsRequestDtoContext2:
		// createTopicSubscriptionsRequestDtoContext.CreateTopicSubscriptionsRequestDtoContext2 is populated
}
```
