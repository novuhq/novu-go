# Actor

It is used to display the Avatar of the provided actor's subscriber id or actor object.
    If a new actor object is provided, we will create a new subscriber in our system


## Supported Types

### 

```go
actor := components.CreateActorStr(string{/* values here */})
```

### SubscriberPayloadDto

```go
actor := components.CreateActorSubscriberPayloadDto(components.SubscriberPayloadDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch actor.Type {
	case components.ActorTypeStr:
		// actor.Str is populated
	case components.ActorTypeSubscriberPayloadDto:
		// actor.SubscriberPayloadDto is populated
}
```
