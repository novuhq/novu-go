# CreateDomainDto


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `Name`                                                                                  | `string`                                                                                | :heavy_check_mark:                                                                      | The domain name (e.g. "recent.dev")                                                     |
| `Data`                                                                                  | map[string]`string`                                                                     | :heavy_minus_sign:                                                                      | Optional string key-value metadata (max 10 keys, 500 characters total for keys+values). |