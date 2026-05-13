# DomainsControllerUpdateDomainRequest


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Domain`                                                                 | `string`                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `IdempotencyKey`                                                         | `*string`                                                                | :heavy_minus_sign:                                                       | A header for idempotency purposes                                        |
| `UpdateDomainDto`                                                        | [components.UpdateDomainDto](../../models/components/updatedomaindto.md) | :heavy_check_mark:                                                       | N/A                                                                      |