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
// SentEmailProjection Collection of items
type SentEmailProjection struct {
	Id string `json:"id"`
	From string `json:"from,omitempty"`
	UserId string `json:"userId"`
	Subject string `json:"subject,omitempty"`
	Attachments []string `json:"attachments"`
	InboxId string `json:"inboxId"`
	To []string `json:"to"`
	CreatedAt time.Time `json:"createdAt"`
	Bcc []string `json:"bcc"`
	Cc []string `json:"cc"`
	BodyMD5Hash string `json:"bodyMD5Hash,omitempty"`
	VirtualSend bool `json:"virtualSend"`
}
