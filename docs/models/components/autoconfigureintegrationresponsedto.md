# AutoConfigureIntegrationResponseDto


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Success`                                                          | *bool*                                                             | :heavy_check_mark:                                                 | Indicates whether the auto-configuration was successful            |
| `Message`                                                          | **string*                                                          | :heavy_minus_sign:                                                 | Optional message describing the result or any errors that occurred |
| `Integration`                                                      | [*components.Integration](../../models/components/integration.md)  | :heavy_minus_sign:                                                 | The updated configurations after auto-configuration                |