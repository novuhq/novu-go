# SubscribersControllerChatAccessOauthRequest


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `SubscriberID`                                        | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `ProviderID`                                          | *any*                                                 | :heavy_check_mark:                                    | N/A                                                   |
| `HmacHash`                                            | *string*                                              | :heavy_check_mark:                                    | HMAC hash for the request                             |
| `EnvironmentID`                                       | *string*                                              | :heavy_check_mark:                                    | The ID of the environment, must be a valid MongoDB ID |
| `IntegrationIdentifier`                               | **string*                                             | :heavy_minus_sign:                                    | Optional integration identifier                       |