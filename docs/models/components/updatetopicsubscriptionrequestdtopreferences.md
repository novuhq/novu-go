# UpdateTopicSubscriptionRequestDtoPreferences


## Supported Types

### 

```go
updateTopicSubscriptionRequestDtoPreferences := components.CreateUpdateTopicSubscriptionRequestDtoPreferencesStr(string{/* values here */})
```

### WorkflowPreferenceRequestDto

```go
updateTopicSubscriptionRequestDtoPreferences := components.CreateUpdateTopicSubscriptionRequestDtoPreferencesWorkflowPreferenceRequestDto(components.WorkflowPreferenceRequestDto{/* values here */})
```

### GroupPreferenceFilterDto

```go
updateTopicSubscriptionRequestDtoPreferences := components.CreateUpdateTopicSubscriptionRequestDtoPreferencesGroupPreferenceFilterDto(components.GroupPreferenceFilterDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateTopicSubscriptionRequestDtoPreferences.Type {
	case components.UpdateTopicSubscriptionRequestDtoPreferencesTypeStr:
		// updateTopicSubscriptionRequestDtoPreferences.Str is populated
	case components.UpdateTopicSubscriptionRequestDtoPreferencesTypeWorkflowPreferenceRequestDto:
		// updateTopicSubscriptionRequestDtoPreferences.WorkflowPreferenceRequestDto is populated
	case components.UpdateTopicSubscriptionRequestDtoPreferencesTypeGroupPreferenceFilterDto:
		// updateTopicSubscriptionRequestDtoPreferences.GroupPreferenceFilterDto is populated
}
```
