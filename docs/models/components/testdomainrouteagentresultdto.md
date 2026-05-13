# TestDomainRouteAgentResultDto


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AgentID`                                                       | `string`                                                        | :heavy_check_mark:                                              | N/A                                                             |
| `HTTPStatus`                                                    | `float64`                                                       | :heavy_check_mark:                                              | N/A                                                             |
| `AgentReply`                                                    | [*components.AgentReply](../../models/components/agentreply.md) | :heavy_minus_sign:                                              | Parsed JSON body from the agent webhook response when JSON.     |
| `LatencyMs`                                                     | `float64`                                                       | :heavy_check_mark:                                              | N/A                                                             |