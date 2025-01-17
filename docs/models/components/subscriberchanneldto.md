# SubscriberChannelDto


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ProviderID`                                                                         | [components.ProviderID](../../models/components/providerid.md)                       | :heavy_check_mark:                                                                   | The ID of the chat or push provider.                                                 |
| `IntegrationIdentifier`                                                              | **string*                                                                            | :heavy_minus_sign:                                                                   | An optional identifier for the integration.                                          |
| `Credentials`                                                                        | [components.ChannelCredentialsDto](../../models/components/channelcredentialsdto.md) | :heavy_check_mark:                                                                   | Credentials for the channel.                                                         |