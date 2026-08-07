# AgentsControllerCreateAgentRequest


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `NovuAnalyticsSource`                                                                | `string`                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `IdempotencyKey`                                                                     | `*string`                                                                            | :heavy_minus_sign:                                                                   | A header for idempotency purposes                                                    |
| `CreateAgentRequestDto`                                                              | [components.CreateAgentRequestDto](../../models/components/createagentrequestdto.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |