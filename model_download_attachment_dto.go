/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// DownloadAttachmentDto Content of attachment
type DownloadAttachmentDto struct {
	// Base64 encoded string of attachment bytes. Decode the base64 string to get the raw file bytes
	Base64FileContents string `json:"base64FileContents,omitempty"`
	// Content type of attachment
	ContentType string `json:"contentType,omitempty"`
	// Size in bytes of attachment
	SizeBytes int64 `json:"sizeBytes,omitempty"`
}