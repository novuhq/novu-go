# UpdateAgentBridgeRequestDto


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `BridgeURL`                                  | `*string`                                    | :heavy_minus_sign:                           | Production bridge URL for this agent         |
| `DevBridgeURL`                               | `*string`                                    | :heavy_minus_sign:                           | Development bridge URL (set by npx novu dev) |
| `DevBridgeActive`                            | `*bool`                                      | :heavy_minus_sign:                           | Whether the dev bridge override is active    |