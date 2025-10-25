# InboundWebhooksControllerHandleWebhookRequest


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `EnvironmentID`                                      | *string*                                             | :heavy_check_mark:                                   | The environment identifier                           |
| `IntegrationID`                                      | *string*                                             | :heavy_check_mark:                                   | The integration identifier for the delivery provider |
| `IdempotencyKey`                                     | **string*                                            | :heavy_minus_sign:                                   | A header for idempotency purposes                    |
| `RequestBody`                                        | *any*                                                | :heavy_check_mark:                                   | Webhook event payload from the delivery provider     |