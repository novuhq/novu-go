# LayoutsControllerDuplicateRequest


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `LayoutID`                                                                     | *string*                                                                       | :heavy_check_mark:                                                             | N/A                                                                            |
| `IdempotencyKey`                                                               | **string*                                                                      | :heavy_minus_sign:                                                             | A header for idempotency purposes                                              |
| `DuplicateLayoutDto`                                                           | [components.DuplicateLayoutDto](../../models/components/duplicatelayoutdto.md) | :heavy_check_mark:                                                             | N/A                                                                            |