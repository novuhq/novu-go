# GetContextResponseDto


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Type`                                                               | `string`                                                             | :heavy_check_mark:                                                   | Context type (e.g., tenant, app, workspace)                          |
| `ID`                                                                 | `string`                                                             | :heavy_check_mark:                                                   | Unique identifier for this context                                   |
| `Data`                                                               | map[string]`any`                                                     | :heavy_check_mark:                                                   | Custom data associated with this context                             |
| `BridgeURL`                                                          | `*string`                                                            | :heavy_minus_sign:                                                   | Bridge URL override for agent connect, if configured on this context |
| `CreatedAt`                                                          | `string`                                                             | :heavy_check_mark:                                                   | Creation timestamp                                                   |
| `UpdatedAt`                                                          | `string`                                                             | :heavy_check_mark:                                                   | Last update timestamp                                                |