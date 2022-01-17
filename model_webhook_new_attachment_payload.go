/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository
 *
 * API version: 6.5.2
 * Contact: contact@mailslurp.dev
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// WebhookNewAttachmentPayload NEW_ATTACHMENT webhook payload. Sent to your webhook url endpoint via HTTP POST when an email is received by the inbox that your webhook is attached to that contains an attachment. You can use the attachmentId to download the attachment.
type WebhookNewAttachmentPayload struct {
	// Idempotent message ID. Store this ID locally or in a database to prevent message duplication.
	MessageId string `json:"messageId"`
	// ID of webhook entity being triggered
	WebhookId string `json:"webhookId"`
	// Name of the webhook being triggered
	WebhookName string `json:"webhookName,omitempty"`
	// Name of the event type webhook is being triggered for.
	EventName string `json:"eventName"`
	// ID of attachment. Use the `AttachmentController` to
	AttachmentId string `json:"attachmentId"`
	// Filename of the attachment if present
	Name string `json:"name"`
	// Content type of attachment such as 'image/png' or 'application/pdf
	ContentType string `json:"contentType"`
	// Size of attachment in bytes
	ContentLength int64 `json:"contentLength"`
}
