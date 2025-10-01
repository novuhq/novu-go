# UploadTranslationsResponseDto


## Fields

| Field                                     | Type                                      | Required                                  | Description                               | Example                                   |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `TotalFiles`                              | *float64*                                 | :heavy_check_mark:                        | Total number of files processed           | 3                                         |
| `SuccessfulUploads`                       | *float64*                                 | :heavy_check_mark:                        | Number of files successfully uploaded     | 2                                         |
| `FailedUploads`                           | *float64*                                 | :heavy_check_mark:                        | Number of files that failed to upload     | 1                                         |
| `Errors`                                  | []*string*                                | :heavy_check_mark:                        | List of error messages for failed uploads | [<br/>"Invalid JSON in file: es-ES.json"<br/>] |