/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails and SMS from dynamically allocated email addresses and phone numbers. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://docs.mailslurp.com/) - [Examples](https://github.com/mailslurp/examples) repository
 *
 * API version: 6.5.2
 * Contact: contact@mailslurp.dev
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
import (
	"time"
)
// TrackingPixelProjection Tracking pixel data
type TrackingPixelProjection struct {
	InboxId string `json:"inboxId,omitempty"`
	UserId string `json:"userId"`
	SentEmailId string `json:"sentEmailId,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	Recipient string `json:"recipient,omitempty"`
	Seen bool `json:"seen"`
	SeenAt time.Time `json:"seenAt,omitempty"`
	Name string `json:"name,omitempty"`
	Id string `json:"id"`
}
