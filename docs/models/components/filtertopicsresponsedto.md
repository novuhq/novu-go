# FilterTopicsResponseDto


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Data`                                                       | [][components.TopicDto](../../models/components/topicdto.md) | :heavy_check_mark:                                           | The list of topics                                           | []                                                           |
| `Page`                                                       | *float64*                                                    | :heavy_check_mark:                                           | The current page number                                      | 1                                                            |
| `PageSize`                                                   | *float64*                                                    | :heavy_check_mark:                                           | The number of items per page                                 | 10                                                           |
| `TotalCount`                                                 | *float64*                                                    | :heavy_check_mark:                                           | The total number of items                                    | 10                                                           |