# TriggerEventResponseDto


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Acknowledged`                                                    | *bool*                                                            | :heavy_check_mark:                                                | Indicates whether the trigger was acknowledged or not             |
| `Status`                                                          | [components.Status](../../models/components/status.md)            | :heavy_check_mark:                                                | Status of the trigger                                             |
| `Error`                                                           | []*string*                                                        | :heavy_minus_sign:                                                | In case of an error, this field will contain the error message(s) |
| `TransactionID`                                                   | **string*                                                         | :heavy_minus_sign:                                                | The returned transaction ID of the trigger                        |