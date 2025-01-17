# BulkCreateSubscriberResponseDto


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Updated`                                                                            | [][components.UpdatedSubscriberDto](../../models/components/updatedsubscriberdto.md) | :heavy_check_mark:                                                                   | An array of subscribers that were successfully updated.                              |
| `Created`                                                                            | [][components.CreatedSubscriberDto](../../models/components/createdsubscriberdto.md) | :heavy_check_mark:                                                                   | An array of subscribers that were successfully created.                              |
| `Failed`                                                                             | [][components.FailedOperationDto](../../models/components/failedoperationdto.md)     | :heavy_check_mark:                                                                   | An array of failed operations with error messages and optional subscriber IDs.       |