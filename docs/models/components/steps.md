# Steps


## Supported Types

### InAppStepUpsertDto

```go
steps := components.CreateStepsInApp(components.InAppStepUpsertDto{/* values here */})
```

### EmailStepUpsertDto

```go
steps := components.CreateStepsEmail(components.EmailStepUpsertDto{/* values here */})
```

### SmsStepUpsertDto

```go
steps := components.CreateStepsSms(components.SmsStepUpsertDto{/* values here */})
```

### PushStepUpsertDto

```go
steps := components.CreateStepsPush(components.PushStepUpsertDto{/* values here */})
```

### ChatStepUpsertDto

```go
steps := components.CreateStepsChat(components.ChatStepUpsertDto{/* values here */})
```

### DelayStepUpsertDto

```go
steps := components.CreateStepsDelay(components.DelayStepUpsertDto{/* values here */})
```

### DigestStepUpsertDto

```go
steps := components.CreateStepsDigest(components.DigestStepUpsertDto{/* values here */})
```

### ThrottleStepUpsertDto

```go
steps := components.CreateStepsThrottle(components.ThrottleStepUpsertDto{/* values here */})
```

### CustomStepUpsertDto

```go
steps := components.CreateStepsCustom(components.CustomStepUpsertDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch steps.Type {
	case components.StepsTypeInApp:
		// steps.InAppStepUpsertDto is populated
	case components.StepsTypeEmail:
		// steps.EmailStepUpsertDto is populated
	case components.StepsTypeSms:
		// steps.SmsStepUpsertDto is populated
	case components.StepsTypePush:
		// steps.PushStepUpsertDto is populated
	case components.StepsTypeChat:
		// steps.ChatStepUpsertDto is populated
	case components.StepsTypeDelay:
		// steps.DelayStepUpsertDto is populated
	case components.StepsTypeDigest:
		// steps.DigestStepUpsertDto is populated
	case components.StepsTypeThrottle:
		// steps.ThrottleStepUpsertDto is populated
	case components.StepsTypeCustom:
		// steps.CustomStepUpsertDto is populated
}
```
