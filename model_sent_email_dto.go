/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
import (
	"time"
)
// SentEmailDto Sent email details
type SentEmailDto struct {
	// Array of IDs of attachments that were sent with this email
	Attachments []string `json:"attachments,omitempty"`
	Bcc []string `json:"bcc,omitempty"`
	Body string `json:"body,omitempty"`
	// MD5 Hash
	BodyMD5Hash string `json:"bodyMD5Hash,omitempty"`
	Cc []string `json:"cc,omitempty"`
	Charset string `json:"charset,omitempty"`
	From string `json:"from,omitempty"`
	// ID of sent email
	Id string `json:"id,omitempty"`
	// Inbox ID email was sent from
	InboxId string `json:"inboxId,omitempty"`
	IsHTML bool `json:"isHTML,omitempty"`
	ReplyTo string `json:"replyTo,omitempty"`
	SentAt time.Time `json:"sentAt"`
	Subject string `json:"subject,omitempty"`
	// Recipients email was sent to
	To []string `json:"to,omitempty"`
	// User ID
	UserId string `json:"userId,omitempty"`
}
