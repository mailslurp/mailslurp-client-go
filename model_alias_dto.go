/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
import (
	"time"
)
// AliasDto Email alias representation
type AliasDto struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// The alias's email address for receiving email
	EmailAddress string `json:"emailAddress,omitempty"`
	Id string `json:"id"`
	// Inbox that is associated with the alias
	InboxId string `json:"inboxId,omitempty"`
	// Has the alias been verified. You must verify an alias if the masked email address has not yet been verified by your account
	IsVerified bool `json:"isVerified,omitempty"`
	// The underlying email address that is hidden and will received forwarded email
	MaskedEmailAddress string `json:"maskedEmailAddress,omitempty"`
	Name string `json:"name,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// If alias will generate response threads or not when email are received by it
	UseThreads bool `json:"useThreads,omitempty"`
	UserId string `json:"userId"`
}
