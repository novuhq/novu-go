# LayoutsControllerCreateRequest


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `IdempotencyKey`                                                         | **string*                                                                | :heavy_minus_sign:                                                       | A header for idempotency purposes                                        |
| `CreateLayoutDto`                                                        | [components.CreateLayoutDto](../../models/components/createlayoutdto.md) | :heavy_check_mark:                                                       | Layout creation details                                                  |