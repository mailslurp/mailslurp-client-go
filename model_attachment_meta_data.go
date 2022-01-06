/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// AttachmentMetaData Meta data associated with an attachment. Attachments are stored as byte blobs so the meta data is stored separately.
type AttachmentMetaData struct {
	// Name of attachment if given
	Name string `json:"name,omitempty"`
	// Content type of attachment such as `image/png`
	ContentType string `json:"contentType,omitempty"`
	// Size of attachment in bytes
	ContentLength int64 `json:"contentLength,omitempty"`
	// ID of attachment
	Id string `json:"id,omitempty"`
}
