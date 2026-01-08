# ChannelEndpointsControllerCreateChannelEndpointRequestBody

Channel endpoint creation request. The structure varies based on the type field.


## Supported Types

### CreateSlackChannelEndpointDto

```go
channelEndpointsControllerCreateChannelEndpointRequestBody := operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodySlackChannel(components.CreateSlackChannelEndpointDto{/* values here */})
```

### CreateSlackUserEndpointDto

```go
channelEndpointsControllerCreateChannelEndpointRequestBody := operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodySlackUser(components.CreateSlackUserEndpointDto{/* values here */})
```

### CreateWebhookEndpointDto

```go
channelEndpointsControllerCreateChannelEndpointRequestBody := operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodyWebhook(components.CreateWebhookEndpointDto{/* values here */})
```

### CreatePhoneEndpointDto

```go
channelEndpointsControllerCreateChannelEndpointRequestBody := operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodyPhone(components.CreatePhoneEndpointDto{/* values here */})
```

### CreateMsTeamsChannelEndpointDto

```go
channelEndpointsControllerCreateChannelEndpointRequestBody := operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodyMsTeamsChannel(components.CreateMsTeamsChannelEndpointDto{/* values here */})
```

### CreateMsTeamsUserEndpointDto

```go
channelEndpointsControllerCreateChannelEndpointRequestBody := operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodyMsTeamsUser(components.CreateMsTeamsUserEndpointDto{/* values here */})
```

