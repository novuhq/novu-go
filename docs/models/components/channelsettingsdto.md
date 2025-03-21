# ChannelSettingsDto


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ProviderID`                                                                   | [components.ProviderID](../../models/components/providerid.md)                 | :heavy_check_mark:                                                             | The provider identifier for the credentials                                    |
| `IntegrationIdentifier`                                                        | **string*                                                                      | :heavy_minus_sign:                                                             | The integration identifier                                                     |
| `Credentials`                                                                  | [components.ChannelCredentials](../../models/components/channelcredentials.md) | :heavy_check_mark:                                                             | Credentials payload for the specified provider                                 |
| `IntegrationID`                                                                | *string*                                                                       | :heavy_check_mark:                                                             | The unique identifier of the integration associated with this channel.         |