/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// WaitForConditions Conditions that a `waitForXEmails` endpoint operates on. The methods wait until given conditions are met or a timeout is reached. If the conditions are met without needing to wait the results will be returned immediately.
type WaitForConditions struct {
	// Number of results that should match conditions. Either exactly or at least this amount based on the `countType`. If count condition is not met and the timeout has not been reached the `waitFor` method will retry the operation.
	Count int32 `json:"count,omitempty"`
	// How should the found count be compared to the expected count.
	CountType string `json:"countType,omitempty"`
	// ID of inbox to search within and apply conditions to. Essentially filtering the emails found to give a count.
	InboxId string `json:"inboxId,omitempty"`
	// Conditions that should be matched for an email to qualify for results. Each condition will be applied in order to each email within an inbox to filter a result list of matching emails you are waiting for.
	Matches []MatchOption `json:"matches,omitempty"`
	// Direction to sort matching emails by created time
	SortDirection string `json:"sortDirection,omitempty"`
	// Max time in milliseconds to retry the `waitFor` operation until conditions are met.
	Timeout int64 `json:"timeout,omitempty"`
	// Apply conditions only to **unread** emails. All emails begin with `read=false`. An email is marked `read=true` when an `EmailDto` representation of it has been returned to the user at least once. For example you have called `getEmail` or `waitForLatestEmail` etc., or you have viewed the email in the dashboard. 
	UnreadOnly bool `json:"unreadOnly,omitempty"`
}
