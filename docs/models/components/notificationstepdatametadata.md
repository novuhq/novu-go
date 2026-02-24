# NotificationStepDataMetadata

Metadata associated with the workflow step. Can vary based on the type of step.


## Supported Types

### DigestRegularMetadata

```go
notificationStepDataMetadata := components.CreateNotificationStepDataMetadataDigestRegularMetadata(components.DigestRegularMetadata{/* values here */})
```

### DigestTimedMetadata

```go
notificationStepDataMetadata := components.CreateNotificationStepDataMetadataDigestTimedMetadata(components.DigestTimedMetadata{/* values here */})
```

### DelayRegularMetadata

```go
notificationStepDataMetadata := components.CreateNotificationStepDataMetadataDelayRegularMetadata(components.DelayRegularMetadata{/* values here */})
```

### DelayScheduledMetadata

```go
notificationStepDataMetadata := components.CreateNotificationStepDataMetadataDelayScheduledMetadata(components.DelayScheduledMetadata{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch notificationStepDataMetadata.Type {
	case components.NotificationStepDataMetadataTypeDigestRegularMetadata:
		// notificationStepDataMetadata.DigestRegularMetadata is populated
	case components.NotificationStepDataMetadataTypeDigestTimedMetadata:
		// notificationStepDataMetadata.DigestTimedMetadata is populated
	case components.NotificationStepDataMetadataTypeDelayRegularMetadata:
		// notificationStepDataMetadata.DelayRegularMetadata is populated
	case components.NotificationStepDataMetadataTypeDelayScheduledMetadata:
		// notificationStepDataMetadata.DelayScheduledMetadata is populated
}
```
