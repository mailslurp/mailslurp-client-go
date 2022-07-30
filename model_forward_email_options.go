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
// ForwardEmailOptions Options for forwarding an email
type ForwardEmailOptions struct {
	// To recipients for forwarded email
	To []string `json:"to"`
	// Subject for forwarded email
	Subject string `json:"subject,omitempty"`
	// Optional cc recipients
	Cc []string `json:"cc,omitempty"`
	// Optional bcc recipients
	Bcc []string `json:"bcc,omitempty"`
	// Optional from override
	From string `json:"from,omitempty"`
	// Optionally use inbox name as display name for sender email address
	UseInboxName bool `json:"useInboxName,omitempty"`
}
