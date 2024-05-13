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
// PageConnector Paginated inbox connectors. Page index starts at zero. Projection results may omit larger entity fields. For fetching a full entity use the projection ID with individual method calls.
type PageConnector struct {
	Content []ConnectorProjection `json:"content,omitempty"`
	Pageable PageableObject `json:"pageable,omitempty"`
	TotalPages int32 `json:"totalPages"`
	TotalElements int64 `json:"totalElements"`
	Last bool `json:"last,omitempty"`
	NumberOfElements int32 `json:"numberOfElements,omitempty"`
	Size int32 `json:"size,omitempty"`
	Number int32 `json:"number,omitempty"`
	Sort SortObject `json:"sort,omitempty"`
	First bool `json:"first,omitempty"`
	Empty bool `json:"empty,omitempty"`
}
