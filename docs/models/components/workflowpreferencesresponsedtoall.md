# WorkflowPreferencesResponseDtoAll

A preference for the workflow. The values specified here will be used if no preference is specified for a channel.


## Supported Types

### WorkflowPreferenceDto

```go
workflowPreferencesResponseDtoAll := components.CreateWorkflowPreferencesResponseDtoAllWorkflowPreferenceDto(components.WorkflowPreferenceDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch workflowPreferencesResponseDtoAll.Type {
	case components.WorkflowPreferencesResponseDtoAllTypeWorkflowPreferenceDto:
		// workflowPreferencesResponseDtoAll.WorkflowPreferenceDto is populated
}
```
