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
import (
	"time"
)
// WebhookNewEmailPayload NEW_EMAIL webhook payload. Sent to your webhook url endpoint via HTTP POST when an email is received by the inbox that your webhook is attached to. Use the email ID to fetch the full email body or attachments.
type WebhookNewEmailPayload struct {
	// Idempotent message ID. Store this ID locally or in a database to prevent message duplication.
	MessageId string `json:"messageId"`
	// ID of webhook entity being triggered
	WebhookId string `json:"webhookId"`
	// Name of the event type webhook is being triggered for.
	EventName string `json:"eventName"`
	// Name of the webhook being triggered
	WebhookName string `json:"webhookName,omitempty"`
	// Id of the inbox
	InboxId string `json:"inboxId"`
	// Id of the domain that received an email
	DomainId string `json:"domainId,omitempty"`
	// ID of the email that was received. Use this ID for fetching the email with the `EmailController`.
	EmailId string `json:"emailId"`
	// Date time of event creation
	CreatedAt time.Time `json:"createdAt"`
	// List of `To` recipient email addresses that the email was addressed to. See recipients object for names.
	To []string `json:"to"`
	// Who the email was sent from. An email address - see fromName for the sender name.
	From string `json:"from"`
	// List of `CC` recipients email addresses that the email was addressed to. See recipients object for names.
	Cc []string `json:"cc"`
	// List of `BCC` recipients email addresses that the email was addressed to. See recipients object for names.
	Bcc []string `json:"bcc"`
	// The subject line of the email message as specified by SMTP subject header
	Subject string `json:"subject,omitempty"`
	// List of attachment meta data objects if attachments present
	AttachmentMetaDatas []AttachmentMetaData `json:"attachmentMetaDatas"`
}
