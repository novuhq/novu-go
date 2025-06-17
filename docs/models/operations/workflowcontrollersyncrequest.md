# WorkflowControllerSyncRequest


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `WorkflowID`                                                             | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `IdempotencyKey`                                                         | **string*                                                                | :heavy_minus_sign:                                                       | A header for idempotency purposes                                        |
| `SyncWorkflowDto`                                                        | [components.SyncWorkflowDto](../../models/components/syncworkflowdto.md) | :heavy_check_mark:                                                       | Sync workflow details                                                    |