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
// GenerateStructuredContentSmsOptions Options for generating structured content output from an SMS
type GenerateStructuredContentSmsOptions struct {
	// SMS ID to read and pass to AI
	SmsId string `json:"smsId"`
	// Optional instructions for the AI to follow. Try to be precise and clear. You can include examples and hints.
	Instructions *string `json:"instructions,omitempty"`
	OutputSchema *StructuredOutputSchema `json:"outputSchema,omitempty"`
	// ID of transformer to apply
	TransformId *string `json:"transformId,omitempty"`
}
