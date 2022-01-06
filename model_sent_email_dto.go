/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository
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
	// ID of sent email
	Id string `json:"id,omitempty"`
	// User ID
	UserId string `json:"userId,omitempty"`
	// Inbox ID email was sent from
	InboxId string `json:"inboxId,omitempty"`
	// Recipients email was sent to
	To []string `json:"to,omitempty"`
	From string `json:"from,omitempty"`
	ReplyTo string `json:"replyTo,omitempty"`
	Cc []string `json:"cc,omitempty"`
	Bcc []string `json:"bcc,omitempty"`
	// Array of IDs of attachments that were sent with this email
	Attachments []string `json:"attachments,omitempty"`
	Subject string `json:"subject,omitempty"`
	// MD5 Hash
	BodyMD5Hash string `json:"bodyMD5Hash,omitempty"`
	Body string `json:"body,omitempty"`
	Charset string `json:"charset,omitempty"`
	SentAt time.Time `json:"sentAt,omitempty"`
	PixelIds []string `json:"pixelIds,omitempty"`
	Html bool `json:"html,omitempty"`
}
