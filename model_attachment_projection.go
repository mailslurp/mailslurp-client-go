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
// AttachmentProjection Email attachment data
type AttachmentProjection struct {
	Name *string `json:"name,omitempty"`
	// Content length of attachment in bytes
	ContentLength *int64 `json:"contentLength,omitempty"`
	// Content type of attachment.
	ContentType *string `json:"contentType,omitempty"`
	UserId string `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// Attachment ID
	AttachmentId string `json:"attachmentId"`
}
