# RedirectDto


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `URL`                                                                       | **string*                                                                   | :heavy_minus_sign:                                                          | URL for redirection. Must be a valid URL or start with / or {{ variable }}. |
| `Target`                                                                    | [*components.Target](../../models/components/target.md)                     | :heavy_minus_sign:                                                          | Target window for the redirection.                                          |