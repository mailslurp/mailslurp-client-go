/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
import (
	"time"
)
// AttachmentProjection struct for AttachmentProjection
type AttachmentProjection struct {
	// Attachment ID
	AttachmentId string `json:"attachmentId,omitempty"`
	// Content length of attachment in bytes
	ContentLength int64 `json:"contentLength,omitempty"`
	// Content type of attachment.
	ContentType string `json:"contentType,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	Name string `json:"name,omitempty"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserId string `json:"userId"`
}