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
// WebhookNewContactPayload NEW_CONTACT webhook payload. Sent to your webhook url endpoint via HTTP POST when an email is received by the inbox that your webhook is attached to that contains a recipient that has not been saved as a contact.
type WebhookNewContactPayload struct {
	// Idempotent message ID. Store this ID locally or in a database to prevent message duplication.
	MessageId string `json:"messageId"`
	// ID of webhook entity being triggered
	WebhookId string `json:"webhookId"`
	// Name of the webhook being triggered
	WebhookName string `json:"webhookName,omitempty"`
	// Name of the event type webhook is being triggered for.
	EventName string `json:"eventName"`
	// Contact ID
	ContactId string `json:"contactId"`
	// Contact group ID
	GroupId *string `json:"groupId,omitempty"`
	// Contact first name
	FirstName *string `json:"firstName,omitempty"`
	// Contact last name
	LastName *string `json:"lastName,omitempty"`
	// Contact company name
	Company *string `json:"company,omitempty"`
	// Primary email address for contact
	PrimaryEmailAddress *string `json:"primaryEmailAddress,omitempty"`
	// Email addresses for contact
	EmailAddresses []string `json:"emailAddresses"`
	// Tags for contact
	Tags []string `json:"tags"`
	MetaData map[string]interface{} `json:"metaData,omitempty"`
	// Has contact opted out of emails
	OptOut bool `json:"optOut"`
	// Date time of event creation
	CreatedAt time.Time `json:"createdAt"`
}
