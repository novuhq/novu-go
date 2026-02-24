# To

The recipients list of people who will receive the notification. Maximum number of recipients can be 100.


## Supported Types

### 

```go
to := components.CreateToArrayOfTo1([]components.To1{/* values here */})
```

### 

```go
to := components.CreateToStr(string{/* values here */})
```

### SubscriberPayloadDto

```go
to := components.CreateToSubscriberPayloadDto(components.SubscriberPayloadDto{/* values here */})
```

### TopicPayloadDto

```go
to := components.CreateToTopicPayloadDto(components.TopicPayloadDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch to.Type {
	case components.ToTypeArrayOfTo1:
		// to.ArrayOfTo1 is populated
	case components.ToTypeStr:
		// to.Str is populated
	case components.ToTypeSubscriberPayloadDto:
		// to.SubscriberPayloadDto is populated
	case components.ToTypeTopicPayloadDto:
		// to.TopicPayloadDto is populated
}
```
