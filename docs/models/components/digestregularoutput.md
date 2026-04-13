# DigestRegularOutput


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Amount`                                                                | `float64`                                                               | :heavy_check_mark:                                                      | Amount of time units                                                    |
| `Unit`                                                                  | [components.TimeUnitEnum](../../models/components/timeunitenum.md)      | :heavy_check_mark:                                                      | Time unit                                                               |
| `DigestKey`                                                             | `*string`                                                               | :heavy_minus_sign:                                                      | Optional digest key                                                     |
| `LookBackWindow`                                                        | [*components.LookBackWindow](../../models/components/lookbackwindow.md) | :heavy_minus_sign:                                                      | Look back window configuration                                          |