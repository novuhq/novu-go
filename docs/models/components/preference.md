# Preference


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Enabled`                                                                         | *bool*                                                                            | :heavy_check_mark:                                                                | Sets if the workflow is fully enabled for all channels or not for the subscriber. |
| `Channels`                                                                        | [components.PreferenceChannels](../../models/components/preferencechannels.md)    | :heavy_check_mark:                                                                | Subscriber preferences for the different channels regarding this workflow         |