# TriggerEventToAllRequestDtoActor

It is used to display the Avatar of the provided actor's subscriber id or actor object.
    If a new actor object is provided, we will create a new subscriber in our system
    


## Supported Types

### 

```go
triggerEventToAllRequestDtoActor := components.CreateTriggerEventToAllRequestDtoActorStr(string{/* values here */})
```

### SubscriberPayloadDto

```go
triggerEventToAllRequestDtoActor := components.CreateTriggerEventToAllRequestDtoActorSubscriberPayloadDto(components.SubscriberPayloadDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch triggerEventToAllRequestDtoActor.Type {
	case components.TriggerEventToAllRequestDtoActorTypeStr:
		// triggerEventToAllRequestDtoActor.Str is populated
	case components.TriggerEventToAllRequestDtoActorTypeSubscriberPayloadDto:
		// triggerEventToAllRequestDtoActor.SubscriberPayloadDto is populated
}
```
