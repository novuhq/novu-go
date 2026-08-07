# WorkflowResponseDtoSteps


## Supported Types

### InAppStepResponseDto

```go
workflowResponseDtoSteps := components.CreateWorkflowResponseDtoStepsInApp(components.InAppStepResponseDto{/* values here */})
```

### EmailStepResponseDto

```go
workflowResponseDtoSteps := components.CreateWorkflowResponseDtoStepsEmail(components.EmailStepResponseDto{/* values here */})
```

### SmsStepResponseDto

```go
workflowResponseDtoSteps := components.CreateWorkflowResponseDtoStepsSms(components.SmsStepResponseDto{/* values here */})
```

### PushStepResponseDto

```go
workflowResponseDtoSteps := components.CreateWorkflowResponseDtoStepsPush(components.PushStepResponseDto{/* values here */})
```

### ChatStepResponseDto

```go
workflowResponseDtoSteps := components.CreateWorkflowResponseDtoStepsChat(components.ChatStepResponseDto{/* values here */})
```

### DelayStepResponseDto

```go
workflowResponseDtoSteps := components.CreateWorkflowResponseDtoStepsDelay(components.DelayStepResponseDto{/* values here */})
```

### DigestStepResponseDto

```go
workflowResponseDtoSteps := components.CreateWorkflowResponseDtoStepsDigest(components.DigestStepResponseDto{/* values here */})
```

### CustomStepResponseDto

```go
workflowResponseDtoSteps := components.CreateWorkflowResponseDtoStepsCustom(components.CustomStepResponseDto{/* values here */})
```

### ThrottleStepResponseDto

```go
workflowResponseDtoSteps := components.CreateWorkflowResponseDtoStepsThrottle(components.ThrottleStepResponseDto{/* values here */})
```

### HTTPRequestStepResponseDto

```go
workflowResponseDtoSteps := components.CreateWorkflowResponseDtoStepsHTTPRequest(components.HTTPRequestStepResponseDto{/* values here */})
```

### ToolStepResponseDto

```go
workflowResponseDtoSteps := components.CreateWorkflowResponseDtoStepsTool(components.ToolStepResponseDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch workflowResponseDtoSteps.Type {
	case components.WorkflowResponseDtoStepsTypeInApp:
		// workflowResponseDtoSteps.InAppStepResponseDto is populated
	case components.WorkflowResponseDtoStepsTypeEmail:
		// workflowResponseDtoSteps.EmailStepResponseDto is populated
	case components.WorkflowResponseDtoStepsTypeSms:
		// workflowResponseDtoSteps.SmsStepResponseDto is populated
	case components.WorkflowResponseDtoStepsTypePush:
		// workflowResponseDtoSteps.PushStepResponseDto is populated
	case components.WorkflowResponseDtoStepsTypeChat:
		// workflowResponseDtoSteps.ChatStepResponseDto is populated
	case components.WorkflowResponseDtoStepsTypeDelay:
		// workflowResponseDtoSteps.DelayStepResponseDto is populated
	case components.WorkflowResponseDtoStepsTypeDigest:
		// workflowResponseDtoSteps.DigestStepResponseDto is populated
	case components.WorkflowResponseDtoStepsTypeCustom:
		// workflowResponseDtoSteps.CustomStepResponseDto is populated
	case components.WorkflowResponseDtoStepsTypeThrottle:
		// workflowResponseDtoSteps.ThrottleStepResponseDto is populated
	case components.WorkflowResponseDtoStepsTypeHTTPRequest:
		// workflowResponseDtoSteps.HTTPRequestStepResponseDto is populated
	case components.WorkflowResponseDtoStepsTypeTool:
		// workflowResponseDtoSteps.ToolStepResponseDto is populated
}
```
