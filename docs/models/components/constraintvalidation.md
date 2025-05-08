# ConstraintValidation


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Messages`                                            | []*string*                                            | :heavy_check_mark:                                    | List of validation error messages                     | [<br/>"Field is required",<br/>"Invalid format"<br/>] |
| `Value`                                               | [*components.Value](../../models/components/value.md) | :heavy_minus_sign:                                    | Value that failed validation                          | xx xx xx                                              |