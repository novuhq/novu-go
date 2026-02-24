# UserAll

A preference for the workflow. The values specified here will be used if no preference is specified for a channel.


## Supported Types

### WorkflowPreferenceDto

```go
userAll := components.CreateUserAllWorkflowPreferenceDto(components.WorkflowPreferenceDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch userAll.Type {
	case components.UserAllTypeWorkflowPreferenceDto:
		// userAll.WorkflowPreferenceDto is populated
}
```
