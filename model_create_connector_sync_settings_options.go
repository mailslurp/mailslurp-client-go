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
// CreateConnectorSyncSettingsOptions Options for creating automatic syncing between an inbox connection and an external mail provider
type CreateConnectorSyncSettingsOptions struct {
	// Enable automatic background sync
	SyncEnabled *bool `json:"syncEnabled,omitempty"`
	// Sync schedule type
	SyncScheduleType *string `json:"syncScheduleType,omitempty"`
	// Sync interval in minutes
	SyncInterval *int32 `json:"syncInterval,omitempty"`
}
