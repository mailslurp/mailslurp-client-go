/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// ConditionOption Options for matching emails in an inbox based on a condition such as `HAS_ATTACHMENTS=TRUE`
type ConditionOption struct {
	// Condition of an email object that can be used to filter results
	Condition string `json:"condition,omitempty"`
	// Expected condition value
	Value string `json:"value,omitempty"`
}
