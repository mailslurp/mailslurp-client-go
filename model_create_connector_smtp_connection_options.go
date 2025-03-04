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
// CreateConnectorSmtpConnectionOptions struct for CreateConnectorSmtpConnectionOptions
type CreateConnectorSmtpConnectionOptions struct {
	Enabled *bool `json:"enabled,omitempty"`
	SmtpHost string `json:"smtpHost"`
	SmtpPort *int32 `json:"smtpPort,omitempty"`
	SmtpSsl *bool `json:"smtpSsl,omitempty"`
	SmtpUsername *string `json:"smtpUsername,omitempty"`
	SmtpPassword *string `json:"smtpPassword,omitempty"`
	SmtpMechanisms *[]string `json:"smtpMechanisms,omitempty"`
	StartTls *bool `json:"startTls,omitempty"`
	LocalHostName *string `json:"localHostName,omitempty"`
	ProxyHost *string `json:"proxyHost,omitempty"`
	ProxyPort *int32 `json:"proxyPort,omitempty"`
	ProxyEnabled *bool `json:"proxyEnabled,omitempty"`
	SslTrust *string `json:"sslTrust,omitempty"`
	SslProtocols *[]string `json:"sslProtocols,omitempty"`
}
