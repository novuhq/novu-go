# InAppRenderOutput


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Subject`                                                         | `*string`                                                         | :heavy_minus_sign:                                                | Subject of the in-app notification                                |
| `Body`                                                            | `string`                                                          | :heavy_check_mark:                                                | Body of the in-app notification                                   |
| `Avatar`                                                          | `*string`                                                         | :heavy_minus_sign:                                                | Avatar for the in-app notification                                |
| `PrimaryAction`                                                   | [*components.ActionDto](../../models/components/actiondto.md)     | :heavy_minus_sign:                                                | Primary action details                                            |
| `SecondaryAction`                                                 | [*components.ActionDto](../../models/components/actiondto.md)     | :heavy_minus_sign:                                                | Secondary action details                                          |
| `Data`                                                            | map[string]`any`                                                  | :heavy_minus_sign:                                                | Additional data                                                   |
| `Redirect`                                                        | [*components.RedirectDto](../../models/components/redirectdto.md) | :heavy_minus_sign:                                                | Redirect details                                                  |