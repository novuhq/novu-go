# SubscribersV1ControllerGetSubscriberRequest


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `SubscriberID`                                     | *string*                                           | :heavy_check_mark:                                 | N/A                                                |
| `IncludeTopics`                                    | **bool*                                            | :heavy_minus_sign:                                 | Includes the topics associated with the subscriber |
| `IdempotencyKey`                                   | **string*                                          | :heavy_minus_sign:                                 | A header for idempotency purposes                  |