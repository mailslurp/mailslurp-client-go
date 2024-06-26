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
// SentEmailDto Sent email details
type SentEmailDto struct {
	// ID of sent email
	Id string `json:"id"`
	// User ID
	UserId string `json:"userId"`
	// Inbox ID email was sent from
	InboxId string `json:"inboxId"`
	// Domain ID
	DomainId *string `json:"domainId,omitempty"`
	// Recipients email was sent to
	To *[]string `json:"to,omitempty"`
	// Sent from address
	From *string `json:"from,omitempty"`
	ReplyTo *string `json:"replyTo,omitempty"`
	Cc *[]string `json:"cc,omitempty"`
	Bcc *[]string `json:"bcc,omitempty"`
	// Array of IDs of attachments that were sent with this email
	Attachments *[]string `json:"attachments,omitempty"`
	Subject *string `json:"subject,omitempty"`
	// MD5 Hash
	BodyMD5Hash *string `json:"bodyMD5Hash,omitempty"`
	// Sent email body
	Body *string `json:"body,omitempty"`
	ToContacts *[]string `json:"toContacts,omitempty"`
	ToGroup *string `json:"toGroup,omitempty"`
	Charset *string `json:"charset,omitempty"`
	IsHTML *bool `json:"isHTML,omitempty"`
	SentAt time.Time `json:"sentAt"`
	PixelIds *[]string `json:"pixelIds,omitempty"`
	MessageId *string `json:"messageId,omitempty"`
	MessageIds *[]string `json:"messageIds,omitempty"`
	VirtualSend *bool `json:"virtualSend,omitempty"`
	TemplateId *string `json:"templateId,omitempty"`
	TemplateVariables *map[string]map[string]interface{} `json:"templateVariables,omitempty"`
	Headers *map[string]string `json:"headers,omitempty"`
	Html bool `json:"html,omitempty"`
}
