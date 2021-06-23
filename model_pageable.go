/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// Pageable struct for Pageable
type Pageable struct {
	Offset int64 `json:"offset,omitempty"`
	PageNumber int32 `json:"pageNumber,omitempty"`
	PageSize int32 `json:"pageSize,omitempty"`
	Paged bool `json:"paged,omitempty"`
	Sort Sort `json:"sort,omitempty"`
	Unpaged bool `json:"unpaged,omitempty"`
}
