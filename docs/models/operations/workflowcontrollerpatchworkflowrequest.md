# WorkflowControllerPatchWorkflowRequest


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `WorkflowID`                                                               | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        |
| `IdempotencyKey`                                                           | **string*                                                                  | :heavy_minus_sign:                                                         | A header for idempotency purposes                                          |
| `PatchWorkflowDto`                                                         | [components.PatchWorkflowDto](../../models/components/patchworkflowdto.md) | :heavy_check_mark:                                                         | Workflow patch details                                                     |