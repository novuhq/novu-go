# AssignSubscriberToTopicDto


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Succeeded`                                                                         | []*string*                                                                          | :heavy_check_mark:                                                                  | List of successfully assigned subscriber IDs                                        |
| `Failed`                                                                            | [*components.FailedAssignmentsDto](../../models/components/failedassignmentsdto.md) | :heavy_minus_sign:                                                                  | Details about failed assignments                                                    |