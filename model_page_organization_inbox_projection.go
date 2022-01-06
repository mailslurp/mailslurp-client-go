/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// PageOrganizationInboxProjection Paginated organization inbox results. Page index starts at zero. Projection results may omit larger entity fields. For fetching a full entity use the projection ID with individual method calls.
type PageOrganizationInboxProjection struct {
	Content []OrganizationInboxProjection `json:"content,omitempty"`
	Pageable Pageable `json:"pageable,omitempty"`
	TotalElements int64 `json:"totalElements,omitempty"`
	TotalPages int32 `json:"totalPages,omitempty"`
	Last bool `json:"last,omitempty"`
	Size int32 `json:"size,omitempty"`
	Number int32 `json:"number,omitempty"`
	Sort Sort `json:"sort,omitempty"`
	NumberOfElements int32 `json:"numberOfElements,omitempty"`
	First bool `json:"first,omitempty"`
	Empty bool `json:"empty,omitempty"`
}
