# Context2

Rich context object with id and optional data


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    | Example                                        |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `ID`                                           | *string*                                       | :heavy_check_mark:                             | N/A                                            | org-acme                                       |
| `Data`                                         | map[string]*any*                               | :heavy_minus_sign:                             | Optional additional context data               | {<br/>"name": "Acme Corp",<br/>"region": "us-east-1"<br/>} |