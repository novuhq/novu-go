# TopicsControllerUpdateTopicRequest


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `TopicKey`                                                                           | *string*                                                                             | :heavy_check_mark:                                                                   | The key identifier of the topic                                                      |
| `IdempotencyKey`                                                                     | **string*                                                                            | :heavy_minus_sign:                                                                   | A header for idempotency purposes                                                    |
| `UpdateTopicRequestDto`                                                              | [components.UpdateTopicRequestDto](../../models/components/updatetopicrequestdto.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |