# Preferences


## Supported Types

### 

```go
preferences := components.CreatePreferencesStr(string{/* values here */})
```

### WorkflowPreferenceRequestDto

```go
preferences := components.CreatePreferencesWorkflowPreferenceRequestDto(components.WorkflowPreferenceRequestDto{/* values here */})
```

### GroupPreferenceFilterDto

```go
preferences := components.CreatePreferencesGroupPreferenceFilterDto(components.GroupPreferenceFilterDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch preferences.Type {
	case components.PreferencesTypeStr:
		// preferences.Str is populated
	case components.PreferencesTypeWorkflowPreferenceRequestDto:
		// preferences.WorkflowPreferenceRequestDto is populated
	case components.PreferencesTypeGroupPreferenceFilterDto:
		// preferences.GroupPreferenceFilterDto is populated
}
```
