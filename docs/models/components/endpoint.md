# Endpoint

Endpoint data specific to the channel type


## Supported Types

### SlackChannelEndpointDto

```go
endpoint := components.CreateEndpointSlackChannelEndpointDto(components.SlackChannelEndpointDto{/* values here */})
```

### SlackUserEndpointDto

```go
endpoint := components.CreateEndpointSlackUserEndpointDto(components.SlackUserEndpointDto{/* values here */})
```

### WebhookEndpointDto

```go
endpoint := components.CreateEndpointWebhookEndpointDto(components.WebhookEndpointDto{/* values here */})
```

### PhoneEndpointDto

```go
endpoint := components.CreateEndpointPhoneEndpointDto(components.PhoneEndpointDto{/* values here */})
```

### MsTeamsChannelEndpointDto

```go
endpoint := components.CreateEndpointMsTeamsChannelEndpointDto(components.MsTeamsChannelEndpointDto{/* values here */})
```

### MsTeamsUserEndpointDto

```go
endpoint := components.CreateEndpointMsTeamsUserEndpointDto(components.MsTeamsUserEndpointDto{/* values here */})
```

### TelegramChatEndpointDto

```go
endpoint := components.CreateEndpointTelegramChatEndpointDto(components.TelegramChatEndpointDto{/* values here */})
```

### WebexRoomEndpointDto

```go
endpoint := components.CreateEndpointWebexRoomEndpointDto(components.WebexRoomEndpointDto{/* values here */})
```

### WebexPersonEndpointDto

```go
endpoint := components.CreateEndpointWebexPersonEndpointDto(components.WebexPersonEndpointDto{/* values here */})
```

### LineUserEndpointDto

```go
endpoint := components.CreateEndpointLineUserEndpointDto(components.LineUserEndpointDto{/* values here */})
```

### PagerDutyServiceEndpointDto

```go
endpoint := components.CreateEndpointPagerDutyServiceEndpointDto(components.PagerDutyServiceEndpointDto{/* values here */})
```

### OpsgenieIntegrationEndpointDto

```go
endpoint := components.CreateEndpointOpsgenieIntegrationEndpointDto(components.OpsgenieIntegrationEndpointDto{/* values here */})
```

### GrafanaOnCallIntegrationEndpointDto

```go
endpoint := components.CreateEndpointGrafanaOnCallIntegrationEndpointDto(components.GrafanaOnCallIntegrationEndpointDto{/* values here */})
```

### ToolWebhookEndpointDto

```go
endpoint := components.CreateEndpointToolWebhookEndpointDto(components.ToolWebhookEndpointDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch endpoint.Type {
	case components.EndpointTypeSlackChannelEndpointDto:
		// endpoint.SlackChannelEndpointDto is populated
	case components.EndpointTypeSlackUserEndpointDto:
		// endpoint.SlackUserEndpointDto is populated
	case components.EndpointTypeWebhookEndpointDto:
		// endpoint.WebhookEndpointDto is populated
	case components.EndpointTypePhoneEndpointDto:
		// endpoint.PhoneEndpointDto is populated
	case components.EndpointTypeMsTeamsChannelEndpointDto:
		// endpoint.MsTeamsChannelEndpointDto is populated
	case components.EndpointTypeMsTeamsUserEndpointDto:
		// endpoint.MsTeamsUserEndpointDto is populated
	case components.EndpointTypeTelegramChatEndpointDto:
		// endpoint.TelegramChatEndpointDto is populated
	case components.EndpointTypeWebexRoomEndpointDto:
		// endpoint.WebexRoomEndpointDto is populated
	case components.EndpointTypeWebexPersonEndpointDto:
		// endpoint.WebexPersonEndpointDto is populated
	case components.EndpointTypeLineUserEndpointDto:
		// endpoint.LineUserEndpointDto is populated
	case components.EndpointTypePagerDutyServiceEndpointDto:
		// endpoint.PagerDutyServiceEndpointDto is populated
	case components.EndpointTypeOpsgenieIntegrationEndpointDto:
		// endpoint.OpsgenieIntegrationEndpointDto is populated
	case components.EndpointTypeGrafanaOnCallIntegrationEndpointDto:
		// endpoint.GrafanaOnCallIntegrationEndpointDto is populated
	case components.EndpointTypeToolWebhookEndpointDto:
		// endpoint.ToolWebhookEndpointDto is populated
}
```
