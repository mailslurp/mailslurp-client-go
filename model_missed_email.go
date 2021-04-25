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
// MissedEmail struct for MissedEmail
type MissedEmail struct {
	AttachmentCount int32 `json:"attachmentCount"`
	Bcc []string `json:"bcc"`
	BodyExcerpt string `json:"bodyExcerpt,omitempty"`
	Cc []string `json:"cc"`
	CreatedAt time.Time `json:"createdAt"`
	From string `json:"from,omitempty"`
	Id string `json:"id,omitempty"`
	InboxIds []string `json:"inboxIds"`
	Subject string `json:"subject,omitempty"`
	To []string `json:"to"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserId string `json:"userId"`
}
