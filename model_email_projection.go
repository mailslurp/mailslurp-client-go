/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: d1659dc1567a5b62f65d0612107a50aace03e085
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mailslurp
import (
	"time"
)
// EmailProjection A compact representation of a full email. Used in list endpoints to keep response sizes low. Body and attachments are not included. To get all fields of the email use the `getEmail` method with the email projection's ID. See `EmailDto` for documentation on projection properties.
type EmailProjection struct {
	Attachments []string `json:"attachments,omitempty"`
	Bcc []string `json:"bcc,omitempty"`
	Cc []string `json:"cc,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	Id string `json:"id"`
	InboxId string `json:"inboxId"`
	Read bool `json:"read,omitempty"`
	Subject string `json:"subject,omitempty"`
	To []string `json:"to"`
}