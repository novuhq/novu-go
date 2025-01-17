# SubscribersControllerGetUnseenCountRequest


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `SubscriberID`                                 | *string*                                       | :heavy_check_mark:                             | N/A                                            |
| `Seen`                                         | **bool*                                        | :heavy_minus_sign:                             | Indicates whether to count seen notifications. |
| `Limit`                                        | **float64*                                     | :heavy_minus_sign:                             | The maximum number of notifications to return. |