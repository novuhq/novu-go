# RuntimeIssueDto


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `IssueType`                                                                          | [components.WorkflowIssueTypeEnum](../../models/components/workflowissuetypeenum.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `VariableName`                                                                       | `*string`                                                                            | :heavy_minus_sign:                                                                   | Variable associated with the issue                                                   |
| `Message`                                                                            | `string`                                                                             | :heavy_check_mark:                                                                   | Human-readable issue message                                                         |