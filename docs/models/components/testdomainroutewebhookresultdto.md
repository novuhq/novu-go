# TestDomainRouteWebhookResultDto


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Skipped`                                                                            | `*bool`                                                                              | :heavy_minus_sign:                                                                   | True when outbound webhooks are disabled for this environment (nothing was emitted). |
| `LatencyMs`                                                                          | `float64`                                                                            | :heavy_check_mark:                                                                   | N/A                                                                                  |