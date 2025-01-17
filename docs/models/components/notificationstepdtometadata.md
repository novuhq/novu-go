# NotificationStepDtoMetadata

Metadata associated with the workflow step. Can vary based on the type of step.


## Supported Types

### DigestRegularMetadata

```go
notificationStepDtoMetadata := components.CreateNotificationStepDtoMetadataDigestRegularMetadata(components.DigestRegularMetadata{/* values here */})
```

### DigestTimedMetadata

```go
notificationStepDtoMetadata := components.CreateNotificationStepDtoMetadataDigestTimedMetadata(components.DigestTimedMetadata{/* values here */})
```

### DelayRegularMetadata

```go
notificationStepDtoMetadata := components.CreateNotificationStepDtoMetadataDelayRegularMetadata(components.DelayRegularMetadata{/* values here */})
```

### DelayScheduledMetadata

```go
notificationStepDtoMetadata := components.CreateNotificationStepDtoMetadataDelayScheduledMetadata(components.DelayScheduledMetadata{/* values here */})
```

