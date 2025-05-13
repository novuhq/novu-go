# CreateTopicSubscriptionsResponseDto


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Data`                                                                               | [][components.SubscriptionDto](../../models/components/subscriptiondto.md)           | :heavy_check_mark:                                                                   | The list of successfully created subscriptions                                       |
| `Meta`                                                                               | [components.MetaDto](../../models/components/metadto.md)                             | :heavy_check_mark:                                                                   | Metadata about the operation                                                         |
| `Errors`                                                                             | [][components.SubscriptionErrorDto](../../models/components/subscriptionerrordto.md) | :heavy_minus_sign:                                                                   | The list of errors for failed subscription attempts                                  |