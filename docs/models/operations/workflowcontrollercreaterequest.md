# WorkflowControllerCreateRequest


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `IdempotencyKey`                                                             | **string*                                                                    | :heavy_minus_sign:                                                           | A header for idempotency purposes                                            |
| `CreateWorkflowDto`                                                          | [components.CreateWorkflowDto](../../models/components/createworkflowdto.md) | :heavy_check_mark:                                                           | Workflow creation details                                                    |