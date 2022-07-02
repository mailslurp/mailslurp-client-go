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
// PageSentEmailWithQueueProjection Paginated sent email results for emails sent with queue. Page index starts at zero. Projection results may omit larger entity fields. For fetching a full sent email entity use the projection ID with individual method calls.
type PageSentEmailWithQueueProjection struct {
	// Collection of items
	Content []SendWithQueueResult `json:"content"`
	Pageable PageableObject `json:"pageable,omitempty"`
	Total int64 `json:"total,omitempty"`
	// Size of page requested
	Size int32 `json:"size"`
	// Page number starting at 0
	Number int32 `json:"number"`
	// Number of items returned
	NumberOfElements int32 `json:"numberOfElements"`
	// Total number of pages available
	TotalPages int32 `json:"totalPages"`
	// Total number of items available for querying
	TotalElements int64 `json:"totalElements"`
	Last bool `json:"last,omitempty"`
	Sort Sort `json:"sort,omitempty"`
	First bool `json:"first,omitempty"`
	Empty bool `json:"empty,omitempty"`
}
