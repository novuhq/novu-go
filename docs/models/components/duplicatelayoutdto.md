# DuplicateLayoutDto


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Name`                                                                           | `string`                                                                         | :heavy_check_mark:                                                               | Name of the layout                                                               |
| `LayoutID`                                                                       | `*string`                                                                        | :heavy_minus_sign:                                                               | Identifier for the duplicated layout. When omitted, it is derived from the name. |
| `IsTranslationEnabled`                                                           | `*bool`                                                                          | :heavy_minus_sign:                                                               | Enable or disable translations for this layout                                   |