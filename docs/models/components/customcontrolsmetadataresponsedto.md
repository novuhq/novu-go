# CustomControlsMetadataResponseDto


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `DataSchema`                                                               | map[string]*any*                                                           | :heavy_minus_sign:                                                         | JSON Schema for data                                                       |
| `UISchema`                                                                 | [*components.UISchema](../../models/components/uischema.md)                | :heavy_minus_sign:                                                         | UI Schema for rendering                                                    |
| `Values`                                                                   | [components.CustomControlDto](../../models/components/customcontroldto.md) | :heavy_check_mark:                                                         | Control values specific to Custom step                                     |