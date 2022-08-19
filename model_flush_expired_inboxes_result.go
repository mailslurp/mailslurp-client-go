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
// FlushExpiredInboxesResult Result from calling expire on any inboxes that have applicable expiration dates given current time.
type FlushExpiredInboxesResult struct {
	// Inbox IDs affected by expiration
	InboxIds []string `json:"inboxIds"`
	// DateTime to filter inboxes so that those expiring before this time are expired
	ExpireBefore time.Time `json:"expireBefore"`
}
