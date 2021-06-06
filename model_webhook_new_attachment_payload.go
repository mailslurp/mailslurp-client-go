/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// WebhookNewAttachmentPayload NEW_ATTACHMENT webhook payload
type WebhookNewAttachmentPayload struct {
	// ID of attachment. Use the `AttachmentController` to
	AttachmentId string `json:"attachmentId,omitempty"`
	// Size of attachment in bytes
	ContentLength int64 `json:"contentLength,omitempty"`
	// Content type of attachment such as 'image/png' or 'application/pdf
	ContentType string `json:"contentType,omitempty"`
	// Name of the event type webhook is being triggered for.
	EventName string `json:"eventName,omitempty"`
	// Idempotent message ID. Store this ID locally or in a database to prevent message duplication.
	MessageId string `json:"messageId,omitempty"`
	// Filename of the attachment if present
	Name string `json:"name,omitempty"`
	// ID of webhook entity being triggered
	WebhookId string `json:"webhookId,omitempty"`
	// Name of the webhook being triggered
	WebhookName string `json:"webhookName,omitempty"`
}