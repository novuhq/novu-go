# To1


## Supported Types

### SubscriberPayloadDto

```go
to1 := components.CreateTo1SubscriberPayloadDto(components.SubscriberPayloadDto{/* values here */})
```

### TopicPayloadDto

```go
to1 := components.CreateTo1TopicPayloadDto(components.TopicPayloadDto{/* values here */})
```

### 

```go
to1 := components.CreateTo1Str(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch to1.Type {
	case components.To1TypeSubscriberPayloadDto:
		// to1.SubscriberPayloadDto is populated
	case components.To1TypeTopicPayloadDto:
		// to1.TopicPayloadDto is populated
	case components.To1TypeStr:
		// to1.Str is populated
}
```
