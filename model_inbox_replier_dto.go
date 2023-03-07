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
// InboxReplierDto Inbox replier. Will automatically reply to inbound emails that match given field for an inbox.
type InboxReplierDto struct {
	Id string `json:"id"`
	InboxId string `json:"inboxId"`
	Name string `json:"name,omitempty"`
	Field string `json:"field"`
	Match string `json:"match"`
	ReplyTo string `json:"replyTo,omitempty"`
	Subject string `json:"subject,omitempty"`
	From string `json:"from,omitempty"`
	Charset string `json:"charset,omitempty"`
	IsHTML bool `json:"isHTML"`
	TemplateId string `json:"templateId,omitempty"`
	TemplateVariables map[string]map[string]interface{} `json:"templateVariables,omitempty"`
	IgnoreReplyTo bool `json:"ignoreReplyTo"`
	CreatedAt time.Time `json:"createdAt"`
}
