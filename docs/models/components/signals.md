# Signals


## Supported Types

### MetadataSetSignalDto

```go
signals := components.CreateSignalsMetadataSetSignalDto(components.MetadataSetSignalDto{/* values here */})
```

### MetadataDeleteSignalDto

```go
signals := components.CreateSignalsMetadataDeleteSignalDto(components.MetadataDeleteSignalDto{/* values here */})
```

### MetadataClearSignalDto

```go
signals := components.CreateSignalsMetadataClearSignalDto(components.MetadataClearSignalDto{/* values here */})
```

### TriggerSignalDto

```go
signals := components.CreateSignalsTriggerSignalDto(components.TriggerSignalDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch signals.Type {
	case components.SignalsTypeMetadataSetSignalDto:
		// signals.MetadataSetSignalDto is populated
	case components.SignalsTypeMetadataDeleteSignalDto:
		// signals.MetadataDeleteSignalDto is populated
	case components.SignalsTypeMetadataClearSignalDto:
		// signals.MetadataClearSignalDto is populated
	case components.SignalsTypeTriggerSignalDto:
		// signals.TriggerSignalDto is populated
}
```
