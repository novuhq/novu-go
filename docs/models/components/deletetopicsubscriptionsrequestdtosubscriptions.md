# DeleteTopicSubscriptionsRequestDtoSubscriptions


## Supported Types

### 

```go
deleteTopicSubscriptionsRequestDtoSubscriptions := components.CreateDeleteTopicSubscriptionsRequestDtoSubscriptionsStr(string{/* values here */})
```

### DeleteTopicSubscriberIdentifierDto

```go
deleteTopicSubscriptionsRequestDtoSubscriptions := components.CreateDeleteTopicSubscriptionsRequestDtoSubscriptionsDeleteTopicSubscriberIdentifierDto(components.DeleteTopicSubscriberIdentifierDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch deleteTopicSubscriptionsRequestDtoSubscriptions.Type {
	case components.DeleteTopicSubscriptionsRequestDtoSubscriptionsTypeStr:
		// deleteTopicSubscriptionsRequestDtoSubscriptions.Str is populated
	case components.DeleteTopicSubscriptionsRequestDtoSubscriptionsTypeDeleteTopicSubscriberIdentifierDto:
		// deleteTopicSubscriptionsRequestDtoSubscriptions.DeleteTopicSubscriberIdentifierDto is populated
}
```
