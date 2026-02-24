# User

User workflow preferences


## Supported Types

### UserWorkflowPreferencesDto

```go
user := components.CreateUserUserWorkflowPreferencesDto(components.UserWorkflowPreferencesDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch user.Type {
	case components.UserTypeUserWorkflowPreferencesDto:
		// user.UserWorkflowPreferencesDto is populated
}
```
