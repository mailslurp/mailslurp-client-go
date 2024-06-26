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
// ConnectorImapConnectionDto struct for ConnectorImapConnectionDto
type ConnectorImapConnectionDto struct {
	ConnectorId string `json:"connectorId"`
	ImapHost string `json:"imapHost,omitempty"`
	ImapPort int32 `json:"imapPort,omitempty"`
	ImapUsername string `json:"imapUsername,omitempty"`
	ImapPassword string `json:"imapPassword,omitempty"`
	ImapSsl bool `json:"imapSsl,omitempty"`
	SelectFolder string `json:"selectFolder,omitempty"`
	SearchTerms string `json:"searchTerms,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	Id string `json:"id"`
}
