# WorkflowControllerUpdateRequest


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `WorkflowID`                                                                 | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `IdempotencyKey`                                                             | **string*                                                                    | :heavy_minus_sign:                                                           | A header for idempotency purposes                                            |
| `UpdateWorkflowDto`                                                          | [components.UpdateWorkflowDto](../../models/components/updateworkflowdto.md) | :heavy_check_mark:                                                           | Workflow update details                                                      |