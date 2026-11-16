# StepListResponseDto


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Slug`                                                                | `string`                                                              | :heavy_check_mark:                                                    | Slug of the step                                                      |
| `Type`                                                                | [components.StepTypeEnum](../../models/components/steptypeenum.md)    | :heavy_check_mark:                                                    | Channels of the notification                                          |
| `Issues`                                                              | [*components.StepIssuesDto](../../models/components/stepissuesdto.md) | :heavy_minus_sign:                                                    | Issues associated with the step                                       |