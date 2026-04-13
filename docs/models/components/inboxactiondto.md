# InboxActionDto


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Label`                                                           | `string`                                                          | :heavy_check_mark:                                                | Label of the action button                                        |
| `IsCompleted`                                                     | `bool`                                                            | :heavy_check_mark:                                                | Whether the action has been completed                             |
| `Redirect`                                                        | [*components.RedirectDto](../../models/components/redirectdto.md) | :heavy_minus_sign:                                                | Redirect configuration for the action                             |