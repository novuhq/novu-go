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

