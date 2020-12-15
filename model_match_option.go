/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// MatchOption Options for matching emails in an inbox. Each match option object contains a `field`, `should` and `value` property. Together they form logical conditions such as `SUBJECT` should `CONTAIN` value.
type MatchOption struct {
	// The email property to match on. One of SUBJECT, TO, BCC, CC or FROM
	Field string `json:"field,omitempty"`
	// What criteria to apply. CONTAIN or EQUAL. Note CONTAIN is recommended due to some SMTP servers adding new lines to fields and body content.
	Should string `json:"should,omitempty"`
	// The value you wish to compare with the value of the field specified using the `should` value passed. For example `BODY` should `CONTAIN` a value passed.
	Value string `json:"value,omitempty"`
}
