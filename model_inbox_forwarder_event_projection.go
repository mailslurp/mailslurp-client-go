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
// InboxForwarderEventProjection Inbox forwarder event
type InboxForwarderEventProjection struct {
	Message *string `json:"message,omitempty"`
	Id *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	EmailId *string `json:"emailId,omitempty"`
	InboxId *string `json:"inboxId,omitempty"`
	UserId *string `json:"userId,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	ForwarderId *string `json:"forwarderId,omitempty"`
}
