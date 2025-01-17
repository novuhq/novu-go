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

