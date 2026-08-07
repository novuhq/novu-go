# ChatRenderOutput


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Body`                                                                          | `*string`                                                                       | :heavy_minus_sign:                                                              | Body of the chat message. Mutually exclusive with `card`.                       |
| `Card`                                                                          | map[string]`any`                                                                | :heavy_minus_sign:                                                              | Rich Chat: compiled provider-agnostic card DSL. Mutually exclusive with `body`. |