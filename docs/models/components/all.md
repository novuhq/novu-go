# All

A preference for the workflow. The values specified here will be used if no preference is specified for a channel.


## Supported Types

### WorkflowPreferenceDto

```go
all := components.CreateAllWorkflowPreferenceDto(components.WorkflowPreferenceDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch all.Type {
	case components.AllTypeWorkflowPreferenceDto:
		// all.WorkflowPreferenceDto is populated
}
```
