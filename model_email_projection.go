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
// EmailProjection A compact representation of a full email. Used in list endpoints to keep response sizes low. Body and attachments are not included. To get all fields of the email use the `getEmail` method with the email projection's ID. See `EmailDto` for documentation on projection properties.
type EmailProjection struct {
	Attachments *[]string `json:"attachments,omitempty"`
	InboxId string `json:"inboxId"`
	To []string `json:"to"`
	DomainId *string `json:"domainId,omitempty"`
	Bcc *[]string `json:"bcc,omitempty"`
	Cc *[]string `json:"cc,omitempty"`
	Read bool `json:"read"`
	BodyExcerpt *string `json:"bodyExcerpt,omitempty"`
	TeamAccess bool `json:"teamAccess"`
	BodyMD5Hash *string `json:"bodyMD5Hash,omitempty"`
	TextExcerpt *string `json:"textExcerpt,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	Subject *string `json:"subject,omitempty"`
	Id string `json:"id"`
	From *string `json:"from,omitempty"`
}
