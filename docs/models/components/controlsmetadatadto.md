# ControlsMetadataDto


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `DataSchema`                                                | map[string]`any`                                            | :heavy_minus_sign:                                          | JSON Schema for data                                        |
| `UISchema`                                                  | [*components.UISchema](../../models/components/uischema.md) | :heavy_minus_sign:                                          | UI Schema for rendering                                     |
| `Values`                                                    | map[string]`any`                                            | :heavy_minus_sign:                                          | Resolved control values for the step                        |