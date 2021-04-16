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
// ThreadProjection A thread is a message thread created for a message received by an alias
type ThreadProjection struct {
	AliasId string `json:"aliasId"`
	Bcc []string `json:"bcc,omitempty"`
	Cc []string `json:"cc,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	Id string `json:"id"`
	InboxId string `json:"inboxId"`
	Name string `json:"name,omitempty"`
	Subject string `json:"subject,omitempty"`
	To []string `json:"to"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserId string `json:"userId"`
}
