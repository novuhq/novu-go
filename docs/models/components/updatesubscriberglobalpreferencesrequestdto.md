# UpdateSubscriberGlobalPreferencesRequestDto


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Enabled`                                                                      | **bool*                                                                        | :heavy_minus_sign:                                                             | Enable or disable the subscriber global preferences.                           |
| `Preferences`                                                                  | [][components.ChannelPreference](../../models/components/channelpreference.md) | :heavy_minus_sign:                                                             | The subscriber global preferences for every ChannelTypeEnum.                   |