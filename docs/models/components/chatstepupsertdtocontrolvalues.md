# ChatStepUpsertDtoControlValues

Control values for the Chat step.


## Supported Types

### ChatControlDto

```go
chatStepUpsertDtoControlValues := components.CreateChatStepUpsertDtoControlValuesChatControlDto(components.ChatControlDto{/* values here */})
```

### 

```go
chatStepUpsertDtoControlValues := components.CreateChatStepUpsertDtoControlValuesMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatStepUpsertDtoControlValues.Type {
	case components.ChatStepUpsertDtoControlValuesTypeChatControlDto:
		// chatStepUpsertDtoControlValues.ChatControlDto is populated
	case components.ChatStepUpsertDtoControlValuesTypeMapOfAny:
		// chatStepUpsertDtoControlValues.MapOfAny is populated
}
```
