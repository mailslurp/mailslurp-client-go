/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// ReplyToEmailOptions struct for ReplyToEmailOptions
type ReplyToEmailOptions struct {
	// List of uploaded attachments to send with the reply. Optional.
	Attachments []string `json:"attachments,omitempty"`
	// Body of the reply email you want to send
	Body string `json:"body,omitempty"`
	// The charset that your message should be sent with. Optional. Default is UTF-8
	Charset string `json:"charset,omitempty"`
	// Is the reply HTML
	IsHTML bool `json:"isHTML,omitempty"`
	// The replyTo header that should be used. Optional
	ReplyTo string `json:"replyTo,omitempty"`
	// When to send the email. Typically immediately
	SendStrategy string `json:"sendStrategy,omitempty"`
	// Template ID to use instead of body. Will use template variable map to fill defined variable slots.
	Template string `json:"template,omitempty"`
	// Template variables if using a template
	TemplateVariables map[string]interface{} `json:"templateVariables,omitempty"`
}
