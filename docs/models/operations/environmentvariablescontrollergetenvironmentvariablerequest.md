# EnvironmentVariablesControllerGetEnvironmentVariableRequest


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `VariableKey`                                              | `string`                                                   | :heavy_check_mark:                                         | The unique key of the environment variable (e.g. BASE_URL) | BASE_URL                                                   |
| `IdempotencyKey`                                           | `*string`                                                  | :heavy_minus_sign:                                         | A header for idempotency purposes                          |                                                            |