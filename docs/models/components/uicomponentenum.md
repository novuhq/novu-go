# UIComponentEnum

Component type for the UI Schema Property

## Example Usage

```go
import (
	"github.com/novuhq/novu-go/v3/models/components"
)

value := components.UIComponentEnumEmailEditorSelect

// Open enum: custom values can be created with a direct type cast
custom := components.UIComponentEnum("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `UIComponentEnumEmailEditorSelect`                  | EMAIL_EDITOR_SELECT                                 |
| `UIComponentEnumLayoutSelect`                       | LAYOUT_SELECT                                       |
| `UIComponentEnumBlockEditor`                        | BLOCK_EDITOR                                        |
| `UIComponentEnumEmailBody`                          | EMAIL_BODY                                          |
| `UIComponentEnumTextFullLine`                       | TEXT_FULL_LINE                                      |
| `UIComponentEnumTextInlineLabel`                    | TEXT_INLINE_LABEL                                   |
| `UIComponentEnumInAppBody`                          | IN_APP_BODY                                         |
| `UIComponentEnumInAppAvatar`                        | IN_APP_AVATAR                                       |
| `UIComponentEnumInAppPrimarySubject`                | IN_APP_PRIMARY_SUBJECT                              |
| `UIComponentEnumInAppButtonDropdown`                | IN_APP_BUTTON_DROPDOWN                              |
| `UIComponentEnumInAppDisableSanitizationSwitch`     | IN_APP_DISABLE_SANITIZATION_SWITCH                  |
| `UIComponentEnumDisableSanitizationSwitch`          | DISABLE_SANITIZATION_SWITCH                         |
| `UIComponentEnumURLTextBox`                         | URL_TEXT_BOX                                        |
| `UIComponentEnumDigestAmount`                       | DIGEST_AMOUNT                                       |
| `UIComponentEnumDigestUnit`                         | DIGEST_UNIT                                         |
| `UIComponentEnumDigestType`                         | DIGEST_TYPE                                         |
| `UIComponentEnumDigestKey`                          | DIGEST_KEY                                          |
| `UIComponentEnumDigestCron`                         | DIGEST_CRON                                         |
| `UIComponentEnumDelayAmount`                        | DELAY_AMOUNT                                        |
| `UIComponentEnumDelayUnit`                          | DELAY_UNIT                                          |
| `UIComponentEnumDelayType`                          | DELAY_TYPE                                          |
| `UIComponentEnumDelayCron`                          | DELAY_CRON                                          |
| `UIComponentEnumDelayDynamicKey`                    | DELAY_DYNAMIC_KEY                                   |
| `UIComponentEnumThrottleType`                       | THROTTLE_TYPE                                       |
| `UIComponentEnumThrottleWindow`                     | THROTTLE_WINDOW                                     |
| `UIComponentEnumThrottleUnit`                       | THROTTLE_UNIT                                       |
| `UIComponentEnumThrottleDynamicKey`                 | THROTTLE_DYNAMIC_KEY                                |
| `UIComponentEnumThrottleThreshold`                  | THROTTLE_THRESHOLD                                  |
| `UIComponentEnumThrottleKey`                        | THROTTLE_KEY                                        |
| `UIComponentEnumExtendToSchedule`                   | EXTEND_TO_SCHEDULE                                  |
| `UIComponentEnumSmsBody`                            | SMS_BODY                                            |
| `UIComponentEnumChatBody`                           | CHAT_BODY                                           |
| `UIComponentEnumPushBody`                           | PUSH_BODY                                           |
| `UIComponentEnumToolBody`                           | TOOL_BODY                                           |
| `UIComponentEnumPushSubject`                        | PUSH_SUBJECT                                        |
| `UIComponentEnumQueryEditor`                        | QUERY_EDITOR                                        |
| `UIComponentEnumData`                               | DATA                                                |
| `UIComponentEnumLayoutEmail`                        | LAYOUT_EMAIL                                        |
| `UIComponentEnumDestinationMethod`                  | DESTINATION_METHOD                                  |
| `UIComponentEnumDestinationURL`                     | DESTINATION_URL                                     |
| `UIComponentEnumDestinationHeaders`                 | DESTINATION_HEADERS                                 |
| `UIComponentEnumDestinationBody`                    | DESTINATION_BODY                                    |
| `UIComponentEnumDestinationResponseBodySchema`      | DESTINATION_RESPONSE_BODY_SCHEMA                    |
| `UIComponentEnumDestinationEnforceSchemaValidation` | DESTINATION_ENFORCE_SCHEMA_VALIDATION               |
| `UIComponentEnumDestinationContinueOnFailure`       | DESTINATION_CONTINUE_ON_FAILURE                     |
| `UIComponentEnumDestinationTimeout`                 | DESTINATION_TIMEOUT                                 |