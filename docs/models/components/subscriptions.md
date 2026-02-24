# Subscriptions


## Supported Types

### 

```go
subscriptions := components.CreateSubscriptionsStr(string{/* values here */})
```

### TopicSubscriberIdentifierDto

```go
subscriptions := components.CreateSubscriptionsTopicSubscriberIdentifierDto(components.TopicSubscriberIdentifierDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch subscriptions.Type {
	case components.SubscriptionsTypeStr:
		// subscriptions.Str is populated
	case components.SubscriptionsTypeTopicSubscriberIdentifierDto:
		// subscriptions.TopicSubscriberIdentifierDto is populated
}
```
