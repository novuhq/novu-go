# MessagesControllerDeleteMessagesByTransactionIDRequest


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Channel`                                                 | [*operations.Channel](../../models/operations/channel.md) | :heavy_minus_sign:                                        | The channel of the message to be deleted                  |                                                           |
| `TransactionID`                                           | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       | 507f1f77bcf86cd799439011                                  |
| `IdempotencyKey`                                          | **string*                                                 | :heavy_minus_sign:                                        | A header for idempotency purposes                         |                                                           |