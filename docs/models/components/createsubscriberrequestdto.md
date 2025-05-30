# CreateSubscriberRequestDto


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `SubscriberID`                            | *string*                                  | :heavy_check_mark:                        | Unique identifier of the subscriber       |
| `FirstName`                               | **string*                                 | :heavy_minus_sign:                        | First name of the subscriber              |
| `LastName`                                | **string*                                 | :heavy_minus_sign:                        | Last name of the subscriber               |
| `Email`                                   | **string*                                 | :heavy_minus_sign:                        | Email address of the subscriber           |
| `Phone`                                   | **string*                                 | :heavy_minus_sign:                        | Phone number of the subscriber            |
| `Avatar`                                  | **string*                                 | :heavy_minus_sign:                        | Avatar URL or identifier                  |
| `Timezone`                                | **string*                                 | :heavy_minus_sign:                        | Timezone of the subscriber                |
| `Locale`                                  | **string*                                 | :heavy_minus_sign:                        | Locale of the subscriber                  |
| `Data`                                    | map[string]*any*                          | :heavy_minus_sign:                        | Additional custom data for the subscriber |