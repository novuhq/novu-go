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

### CreateWebexRoomEndpointDto

```go
channelEndpointsControllerCreateChannelEndpointRequestBody := operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodyWebexRoom(components.CreateWebexRoomEndpointDto{/* values here */})
```

### CreateWebexPersonEndpointDto

```go
channelEndpointsControllerCreateChannelEndpointRequestBody := operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodyWebexPerson(components.CreateWebexPersonEndpointDto{/* values here */})
```

### CreateLineUserEndpointDto

```go
channelEndpointsControllerCreateChannelEndpointRequestBody := operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodyLineUser(components.CreateLineUserEndpointDto{/* values here */})
```

### CreatePagerDutyServiceEndpointDto

```go
channelEndpointsControllerCreateChannelEndpointRequestBody := operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodyPagerdutyService(components.CreatePagerDutyServiceEndpointDto{/* values here */})
```

### CreateOpsgenieIntegrationEndpointDto

```go
channelEndpointsControllerCreateChannelEndpointRequestBody := operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodyOpsgenieIntegration(components.CreateOpsgenieIntegrationEndpointDto{/* values here */})
```

### CreateGrafanaOnCallIntegrationEndpointDto

```go
channelEndpointsControllerCreateChannelEndpointRequestBody := operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodyGrafanaOncallIntegration(components.CreateGrafanaOnCallIntegrationEndpointDto{/* values here */})
```

### CreateToolWebhookEndpointDto

```go
channelEndpointsControllerCreateChannelEndpointRequestBody := operations.CreateChannelEndpointsControllerCreateChannelEndpointRequestBodyToolWebhook(components.CreateToolWebhookEndpointDto{/* values here */})
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
	case operations.ChannelEndpointsControllerCreateChannelEndpointRequestBodyTypeWebexRoom:
		// channelEndpointsControllerCreateChannelEndpointRequestBody.CreateWebexRoomEndpointDto is populated
	case operations.ChannelEndpointsControllerCreateChannelEndpointRequestBodyTypeWebexPerson:
		// channelEndpointsControllerCreateChannelEndpointRequestBody.CreateWebexPersonEndpointDto is populated
	case operations.ChannelEndpointsControllerCreateChannelEndpointRequestBodyTypeLineUser:
		// channelEndpointsControllerCreateChannelEndpointRequestBody.CreateLineUserEndpointDto is populated
	case operations.ChannelEndpointsControllerCreateChannelEndpointRequestBodyTypePagerdutyService:
		// channelEndpointsControllerCreateChannelEndpointRequestBody.CreatePagerDutyServiceEndpointDto is populated
	case operations.ChannelEndpointsControllerCreateChannelEndpointRequestBodyTypeOpsgenieIntegration:
		// channelEndpointsControllerCreateChannelEndpointRequestBody.CreateOpsgenieIntegrationEndpointDto is populated
	case operations.ChannelEndpointsControllerCreateChannelEndpointRequestBodyTypeGrafanaOncallIntegration:
		// channelEndpointsControllerCreateChannelEndpointRequestBody.CreateGrafanaOnCallIntegrationEndpointDto is populated
	case operations.ChannelEndpointsControllerCreateChannelEndpointRequestBodyTypeToolWebhook:
		// channelEndpointsControllerCreateChannelEndpointRequestBody.CreateToolWebhookEndpointDto is populated
}
```
