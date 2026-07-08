# ConfigureTelegramWebhookResponseDto


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `WebhookURL`                                           | `string`                                               | :heavy_check_mark:                                     | URL Novu registered with Telegram for incoming updates |
| `ConfiguredAt`                                         | `string`                                               | :heavy_check_mark:                                     | ISO-8601 timestamp the webhook was configured at       |
| `BotUsername`                                          | `string`                                               | :heavy_check_mark:                                     | Resolved bot username from getMe                       |