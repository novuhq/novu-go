# TriggerEventRequestDtoTo

The recipients list of people who will receive the notification. Maximum number of recipients can be 100.


## Supported Types

### 

```go
triggerEventRequestDtoTo := components.CreateTriggerEventRequestDtoToArrayOfTo1([]components.To1{/* values here */})
```

### 

```go
triggerEventRequestDtoTo := components.CreateTriggerEventRequestDtoToStr(string{/* values here */})
```

### SubscriberPayloadDto

```go
triggerEventRequestDtoTo := components.CreateTriggerEventRequestDtoToSubscriberPayloadDto(components.SubscriberPayloadDto{/* values here */})
```

### TopicPayloadDto

```go
triggerEventRequestDtoTo := components.CreateTriggerEventRequestDtoToTopicPayloadDto(components.TopicPayloadDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch triggerEventRequestDtoTo.Type {
	case components.TriggerEventRequestDtoToTypeArrayOfTo1:
		// triggerEventRequestDtoTo.ArrayOfTo1 is populated
	case components.TriggerEventRequestDtoToTypeStr:
		// triggerEventRequestDtoTo.Str is populated
	case components.TriggerEventRequestDtoToTypeSubscriberPayloadDto:
		// triggerEventRequestDtoTo.SubscriberPayloadDto is populated
	case components.TriggerEventRequestDtoToTypeTopicPayloadDto:
		// triggerEventRequestDtoTo.TopicPayloadDto is populated
}
```
