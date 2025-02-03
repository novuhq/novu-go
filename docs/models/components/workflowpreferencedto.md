# WorkflowPreferenceDto


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Enabled`                                                                      | *bool*                                                                         | :heavy_check_mark:                                                             | Whether notifications are enabled for this workflow                            |
| `Channels`                                                                     | [components.PreferenceChannels](../../models/components/preferencechannels.md) | :heavy_check_mark:                                                             | Channel-specific preference settings for this workflow                         |
| `Overrides`                                                                    | [][components.Overrides](../../models/components/overrides.md)                 | :heavy_check_mark:                                                             | List of preference overrides                                                   |
| `Workflow`                                                                     | [components.WorkflowInfoDto](../../models/components/workflowinfodto.md)       | :heavy_check_mark:                                                             | Workflow information                                                           |