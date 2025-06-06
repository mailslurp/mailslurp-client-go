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
// SearchInboxesOptions struct for SearchInboxesOptions
type SearchInboxesOptions struct {
	// Optional page index in list pagination
	PageIndex *int32 `json:"pageIndex,omitempty"`
	// Optional page size in list pagination
	PageSize *int32 `json:"pageSize,omitempty"`
	// Optional createdAt sort direction ASC or DESC
	SortDirection *string `json:"sortDirection,omitempty"`
	// Optionally filter results for favourites only
	Favourite *bool `json:"favourite,omitempty"`
	// Optionally filter by search words partial matching ID, tags, name, and email address
	Search *string `json:"search,omitempty"`
	// Optionally filter by tags. Will return inboxes that include given tags
	Tag *string `json:"tag,omitempty"`
	// Optional filter by created after given date time
	Since *time.Time `json:"since,omitempty"`
	// Optional filter by created before given date time
	Before *time.Time `json:"before,omitempty"`
	// Type of inbox. HTTP inboxes are faster and better for most cases. SMTP inboxes are more suited for public facing inbound messages (but cannot send).
	InboxType *string `json:"inboxType,omitempty"`
	// Optional filter by inbox function
	InboxFunction *string `json:"inboxFunction,omitempty"`
	// Optional domain ID filter
	DomainId *string `json:"domainId,omitempty"`
}
