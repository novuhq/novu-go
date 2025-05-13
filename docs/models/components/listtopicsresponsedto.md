# ListTopicsResponseDto


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Data`                                                                          | [][components.TopicResponseDto](../../models/components/topicresponsedto.md)    | :heavy_check_mark:                                                              | List of returned Topics                                                         |
| `Next`                                                                          | *string*                                                                        | :heavy_check_mark:                                                              | The cursor for the next page of results, or null if there are no more pages.    |
| `Previous`                                                                      | *string*                                                                        | :heavy_check_mark:                                                              | The cursor for the previous page of results, or null if this is the first page. |