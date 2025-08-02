# PatchWorkflowDto


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `Active`                                         | **bool*                                          | :heavy_minus_sign:                               | Activate or deactivate the workflow              |
| `Name`                                           | **string*                                        | :heavy_minus_sign:                               | New name for the workflow                        |
| `Description`                                    | **string*                                        | :heavy_minus_sign:                               | Updated description of the workflow              |
| `Tags`                                           | []*string*                                       | :heavy_minus_sign:                               | Tags associated with the workflow                |
| `PayloadSchema`                                  | map[string]*any*                                 | :heavy_minus_sign:                               | The payload JSON Schema for the workflow         |
| `ValidatePayload`                                | **bool*                                          | :heavy_minus_sign:                               | Enable or disable payload schema validation      |
| `IsTranslationEnabled`                           | **bool*                                          | :heavy_minus_sign:                               | Enable or disable translations for this workflow |