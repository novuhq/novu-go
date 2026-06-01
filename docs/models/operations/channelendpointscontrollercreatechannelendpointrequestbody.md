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

### CreateTelegramChatEndpointDto

```go
channelEndpointsControllerCreateChannelEndpointRequestBody := operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodyTelegramChat(components.CreateTelegramChatEndpointDto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch channelEndpointsControllerCreateChannelEndpointRequestBody.Type {
	case operations.ChannelEndpointsControllerCreateChannelEndpointRequestBodyTypeSlackChannel:
		// channelEndpointsControllerCreateChannelEndpointRequestBody.CreateSlackChannelEndpointDto is populated
	case operations.ChannelEndpointsControllerCreateChannelEndpointRequestBodyTypeSlackUser:
		// channelEndpointsControllerCreateChannelEndpointRequestBody.CreateSlackUserEndpointDto is populated
	case operations.ChannelEndpointsControllerCreateChannelEndpointRequestBodyTypeWebhook:
		// channelEndpointsControllerCreateChannelEndpointRequestBody.CreateWebhookEndpointDto is populated
	case operations.ChannelEndpointsControllerCreateChannelEndpointRequestBodyTypePhone:
		// channelEndpointsControllerCreateChannelEndpointRequestBody.CreatePhoneEndpointDto is populated
	case operations.ChannelEndpointsControllerCreateChannelEndpointRequestBodyTypeMsTeamsChannel:
		// channelEndpointsControllerCreateChannelEndpointRequestBody.CreateMsTeamsChannelEndpointDto is populated
	case operations.ChannelEndpointsControllerCreateChannelEndpointRequestBodyTypeMsTeamsUser:
		// channelEndpointsControllerCreateChannelEndpointRequestBody.CreateMsTeamsUserEndpointDto is populated
	case operations.ChannelEndpointsControllerCreateChannelEndpointRequestBodyTypeTelegramChat:
		// channelEndpointsControllerCreateChannelEndpointRequestBody.CreateTelegramChatEndpointDto is populated
}
```
