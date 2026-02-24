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
}
```
