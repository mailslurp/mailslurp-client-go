/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://docs.mailslurp.com/) - [Examples](https://github.com/mailslurp/examples) repository
 *
 * API version: 6.5.2
 * Contact: contact@mailslurp.dev
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// SmsMatchOption Options for matching SMS messages in a phone number. Each match option object contains a `field`, `should` and `value` property. Together they form logical conditions such as `BODY` should `CONTAIN` value.
type SmsMatchOption struct {
	// Fields of an SMS object that can be used to filter results
	Field string `json:"field"`
	// How the value of the email field specified should be compared to the value given in the match options.
	Should string `json:"should"`
	// The value you wish to compare with the value of the field specified using the `should` value passed. For example `BODY` should `CONTAIN` a value passed.
	Value string `json:"value"`
}
