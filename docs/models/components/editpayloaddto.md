# EditPayloadDto


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `MessageID`                                                              | `string`                                                                 | :heavy_check_mark:                                                       | Platform message id of the message to edit.                              | 1712345678.123456                                                        |
| `Content`                                                                | [components.Content](../../models/components/content.md)                 | :heavy_check_mark:                                                       | Replacement content. Exactly one of markdown, card, or toolApprovalCard. |                                                                          |