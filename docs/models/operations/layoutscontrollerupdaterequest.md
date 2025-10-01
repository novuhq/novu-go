# LayoutsControllerUpdateRequest


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `LayoutID`                                                               | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `IdempotencyKey`                                                         | **string*                                                                | :heavy_minus_sign:                                                       | A header for idempotency purposes                                        |
| `UpdateLayoutDto`                                                        | [components.UpdateLayoutDto](../../models/components/updatelayoutdto.md) | :heavy_check_mark:                                                       | Layout update details                                                    |