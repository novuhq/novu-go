# InAppStepUpsertDto


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ID`                                                                  | **string*                                                             | :heavy_minus_sign:                                                    | Unique identifier of the step                                         |
| `Name`                                                                | *string*                                                              | :heavy_check_mark:                                                    | Name of the step                                                      |
| `Type`                                                                | [components.StepTypeEnum](../../models/components/steptypeenum.md)    | :heavy_check_mark:                                                    | Type of the step                                                      |
| `ControlValues`                                                       | [*components.ControlValues](../../models/components/controlvalues.md) | :heavy_minus_sign:                                                    | Control values for the In-App step                                    |