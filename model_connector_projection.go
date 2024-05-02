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
// ConnectorProjection Connector
type ConnectorProjection struct {
	Enabled bool `json:"enabled,omitempty"`
	UserId string `json:"userId"`
	EmailAddress string `json:"emailAddress,omitempty"`
	InboxId string `json:"inboxId"`
	SyncEnabled bool `json:"syncEnabled,omitempty"`
	SyncScheduleType string `json:"syncScheduleType"`
	SyncInterval int32 `json:"syncInterval,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	Name string `json:"name,omitempty"`
	Id string `json:"id,omitempty"`
}
