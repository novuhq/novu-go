# Thursday

Thursday schedule


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `IsEnabled`                                                          | *bool*                                                               | :heavy_check_mark:                                                   | Day schedule enabled                                                 | true                                                                 |
| `Hours`                                                              | [][components.TimeRangeDto](../../models/components/timerangedto.md) | :heavy_minus_sign:                                                   | Hours                                                                | [<br/>{<br/>"start": "09:00 AM",<br/>"end": "05:00 PM"<br/>}<br/>]   |