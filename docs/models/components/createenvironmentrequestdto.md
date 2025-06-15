# CreateEnvironmentRequestDto


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Name`                                                | *string*                                              | :heavy_check_mark:                                    | Name of the environment to be created                 | Production Environment                                |
| `ParentID`                                            | **string*                                             | :heavy_minus_sign:                                    | MongoDB ObjectId of the parent environment (optional) | 60d5ecb8b3b3a30015f3e1a1                              |
| `Color`                                               | *string*                                              | :heavy_check_mark:                                    | Hex color code for the environment                    | #3498db                                               |