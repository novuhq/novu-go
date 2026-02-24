# UpdateWorkflowDtoSteps


## Supported Types

### InAppStepUpsertDto

```go
updateWorkflowDtoSteps := components.CreateUpdateWorkflowDtoStepsInApp(components.InAppStepUpsertDto{/* values here */})
```

### EmailStepUpsertDto

```go
updateWorkflowDtoSteps := components.CreateUpdateWorkflowDtoStepsEmail(components.EmailStepUpsertDto{/* values here */})
```

### SmsStepUpsertDto

```go
updateWorkflowDtoSteps := components.CreateUpdateWorkflowDtoStepsSms(components.SmsStepUpsertDto{/* values here */})
```

### PushStepUpsertDto

```go
updateWorkflowDtoSteps := components.CreateUpdateWorkflowDtoStepsPush(components.PushStepUpsertDto{/* values here */})
```

### ChatStepUpsertDto

```go
updateWorkflowDtoSteps := components.CreateUpdateWorkflowDtoStepsChat(components.ChatStepUpsertDto{/* values here */})
```

### DelayStepUpsertDto

```go
updateWorkflowDtoSteps := components.CreateUpdateWorkflowDtoStepsDelay(components.DelayStepUpsertDto{/* values here */})
```

### DigestStepUpsertDto

```go
updateWorkflowDtoSteps := components.CreateUpdateWorkflowDtoStepsDigest(components.DigestStepUpsertDto{/* values here */})
```

### CustomStepUpsertDto

```go
updateWorkflowDtoSteps := components.CreateUpdateWorkflowDtoStepsCustom(components.CustomStepUpsertDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateWorkflowDtoSteps.Type {
	case components.UpdateWorkflowDtoStepsTypeInApp:
		// updateWorkflowDtoSteps.InAppStepUpsertDto is populated
	case components.UpdateWorkflowDtoStepsTypeEmail:
		// updateWorkflowDtoSteps.EmailStepUpsertDto is populated
	case components.UpdateWorkflowDtoStepsTypeSms:
		// updateWorkflowDtoSteps.SmsStepUpsertDto is populated
	case components.UpdateWorkflowDtoStepsTypePush:
		// updateWorkflowDtoSteps.PushStepUpsertDto is populated
	case components.UpdateWorkflowDtoStepsTypeChat:
		// updateWorkflowDtoSteps.ChatStepUpsertDto is populated
	case components.UpdateWorkflowDtoStepsTypeDelay:
		// updateWorkflowDtoSteps.DelayStepUpsertDto is populated
	case components.UpdateWorkflowDtoStepsTypeDigest:
		// updateWorkflowDtoSteps.DigestStepUpsertDto is populated
	case components.UpdateWorkflowDtoStepsTypeCustom:
		// updateWorkflowDtoSteps.CustomStepUpsertDto is populated
}
```
