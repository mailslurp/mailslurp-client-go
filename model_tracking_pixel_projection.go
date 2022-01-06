/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository
 *
 * API version: 6.5.2
 * Contact: contact@mailslurp.dev
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
import (
	"time"
)
// TrackingPixelProjection struct for TrackingPixelProjection
type TrackingPixelProjection struct {
	Name string `json:"name,omitempty"`
	Id string `json:"id,omitempty"`
	UserId string `json:"userId,omitempty"`
	InboxId string `json:"inboxId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Seen bool `json:"seen,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	SeenAt time.Time `json:"seenAt,omitempty"`
	SentEmailId string `json:"sentEmailId,omitempty"`
}
