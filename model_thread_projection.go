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
// ThreadProjection A thread is a message thread created for a message received by an alias
type ThreadProjection struct {
	// Created at DateTime
	CreatedAt time.Time `json:"createdAt"`
	// Updated at DateTime
	UpdatedAt time.Time `json:"updatedAt"`
	// Inbox ID
	InboxId string `json:"inboxId"`
	// User ID
	UserId string `json:"userId"`
	// To recipients
	To []string `json:"to"`
	// BCC recipients
	Bcc []string `json:"bcc,omitempty"`
	// CC recipients
	Cc []string `json:"cc,omitempty"`
	// Alias ID
	AliasId string `json:"aliasId"`
	// Thread subject
	Subject string `json:"subject,omitempty"`
	// Name of thread
	Name string `json:"name,omitempty"`
	// ID of email thread
	Id string `json:"id"`
}
