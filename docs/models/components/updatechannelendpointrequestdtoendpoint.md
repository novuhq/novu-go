# UpdateChannelEndpointRequestDtoEndpoint

Updated endpoint data. The structure must match the existing channel endpoint type.


## Supported Types

### SlackChannelEndpointDto

```go
updateChannelEndpointRequestDtoEndpoint := components.CreateUpdateChannelEndpointRequestDtoEndpointSlackChannelEndpointDto(components.SlackChannelEndpointDto{/* values here */})
```

### SlackUserEndpointDto

```go
updateChannelEndpointRequestDtoEndpoint := components.CreateUpdateChannelEndpointRequestDtoEndpointSlackUserEndpointDto(components.SlackUserEndpointDto{/* values here */})
```

### WebhookEndpointDto

```go
updateChannelEndpointRequestDtoEndpoint := components.CreateUpdateChannelEndpointRequestDtoEndpointWebhookEndpointDto(components.WebhookEndpointDto{/* values here */})
```

### PhoneEndpointDto

```go
updateChannelEndpointRequestDtoEndpoint := components.CreateUpdateChannelEndpointRequestDtoEndpointPhoneEndpointDto(components.PhoneEndpointDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateChannelEndpointRequestDtoEndpoint.Type {
	case components.UpdateChannelEndpointRequestDtoEndpointTypeSlackChannelEndpointDto:
		// updateChannelEndpointRequestDtoEndpoint.SlackChannelEndpointDto is populated
	case components.UpdateChannelEndpointRequestDtoEndpointTypeSlackUserEndpointDto:
		// updateChannelEndpointRequestDtoEndpoint.SlackUserEndpointDto is populated
	case components.UpdateChannelEndpointRequestDtoEndpointTypeWebhookEndpointDto:
		// updateChannelEndpointRequestDtoEndpoint.WebhookEndpointDto is populated
	case components.UpdateChannelEndpointRequestDtoEndpointTypePhoneEndpointDto:
		// updateChannelEndpointRequestDtoEndpoint.PhoneEndpointDto is populated
}
```
