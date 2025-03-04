/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails and SMS from dynamically allocated email addresses and phone numbers. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://docs.mailslurp.com/) - [Examples](https://github.com/mailslurp/examples) repository
 *
 * API version: 6.5.2
 * Contact: contact@mailslurp.dev
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
import (
	"time"
)
// WaitForSmsConditions Conditions to apply to emails that you are waiting for
type WaitForSmsConditions struct {
	// ID of phone number to search within and apply conditions to. Essentially filtering the SMS found to give a count.
	PhoneNumberId string `json:"phoneNumberId"`
	// Limit results
	Limit *int32 `json:"limit,omitempty"`
	// Number of results that should match conditions. Either exactly or at least this amount based on the `countType`. If count condition is not met and the timeout has not been reached the `waitFor` method will retry the operation.
	Count int64 `json:"count"`
	// Max time in milliseconds to wait between retries if a `timeout` is specified.
	DelayTimeout *int64 `json:"delayTimeout,omitempty"`
	// Max time in milliseconds to retry the `waitFor` operation until conditions are met.
	Timeout int64 `json:"timeout"`
	// Apply conditions only to **unread** SMS. All SMS messages begin with `read=false`. An SMS is marked `read=true` when an `SMS` has been returned to the user at least once. For example you have called `getSms` or `waitForSms` etc., or you have viewed the SMS in the dashboard.
	UnreadOnly *bool `json:"unreadOnly,omitempty"`
	// How result size should be compared with the expected size. Exactly or at-least matching result?
	CountType *string `json:"countType,omitempty"`
	// Conditions that should be matched for an SMS to qualify for results. Each condition will be applied in order to each SMS within a phone number to filter a result list of matching SMSs you are waiting for.
	Matches *[]SmsMatchOption `json:"matches,omitempty"`
	// Direction to sort matching SMSs by created time
	SortDirection *string `json:"sortDirection,omitempty"`
	// ISO Date Time earliest time of SMS to consider. Filter for matching SMSs that were received after this date
	Since *time.Time `json:"since,omitempty"`
	// ISO Date Time latest time of SMS to consider. Filter for matching SMSs that were received before this date
	Before *time.Time `json:"before,omitempty"`
}
