# BulkUpdateSubscriberPreferencesDtoContext


## Supported Types

### 

```go
bulkUpdateSubscriberPreferencesDtoContext := components.CreateBulkUpdateSubscriberPreferencesDtoContextStr(string{/* values here */})
```

### Context2

```go
bulkUpdateSubscriberPreferencesDtoContext := components.CreateBulkUpdateSubscriberPreferencesDtoContextContext2(components.Context2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch bulkUpdateSubscriberPreferencesDtoContext.Type {
	case components.BulkUpdateSubscriberPreferencesDtoContextTypeStr:
		// bulkUpdateSubscriberPreferencesDtoContext.Str is populated
	case components.BulkUpdateSubscriberPreferencesDtoContextTypeContext2:
		// bulkUpdateSubscriberPreferencesDtoContext.Context2 is populated
}
```
