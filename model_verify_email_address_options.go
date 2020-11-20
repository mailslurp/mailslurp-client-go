/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// VerifyEmailAddressOptions Options for verifying that an email address exists at a remote mail server.
type VerifyEmailAddressOptions struct {
	MailServerDomain string `json:"mailServerDomain,omitempty"`
	EmailAddress string `json:"emailAddress"`
	SenderEmailAddress string `json:"senderEmailAddress,omitempty"`
	Port int32 `json:"port,omitempty"`
}
