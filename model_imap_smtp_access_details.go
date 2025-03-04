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
// ImapSmtpAccessDetails Access details for inbox using SMTP or IMAP
type ImapSmtpAccessDetails struct {
	// Email address for SMTP/IMAP login
	EmailAddress string `json:"emailAddress"`
	// Secure TLS SMTP server host domain
	SecureSmtpServerHost string `json:"secureSmtpServerHost"`
	// Secure TLS SMTP server host port
	SecureSmtpServerPort int32 `json:"secureSmtpServerPort"`
	// Secure TLS SMTP username for login
	SecureSmtpUsername string `json:"secureSmtpUsername"`
	// Secure TLS SMTP password for login
	SecureSmtpPassword string `json:"secureSmtpPassword"`
	// SMTP server host domain
	SmtpServerHost string `json:"smtpServerHost"`
	// SMTP server host port
	SmtpServerPort int32 `json:"smtpServerPort"`
	// SMTP username for login
	SmtpUsername string `json:"smtpUsername"`
	// SMTP password for login
	SmtpPassword string `json:"smtpPassword"`
	// Secure TLS IMAP server host domain
	SecureImapServerHost string `json:"secureImapServerHost"`
	// Secure TLS IMAP server host port
	SecureImapServerPort int32 `json:"secureImapServerPort"`
	// Secure TLS IMAP username for login
	SecureImapUsername string `json:"secureImapUsername"`
	// Secure TLS IMAP password for login
	SecureImapPassword string `json:"secureImapPassword"`
	// IMAP server host domain
	ImapServerHost string `json:"imapServerHost"`
	// IMAP server host port
	ImapServerPort int32 `json:"imapServerPort"`
	// IMAP username for login
	ImapUsername string `json:"imapUsername"`
	// IMAP password for login
	ImapPassword string `json:"imapPassword"`
	// IMAP mailbox to SELECT
	ImapMailbox string `json:"imapMailbox"`
	// Mail from domain or SMTP HELO value
	MailFromDomain *string `json:"mailFromDomain,omitempty"`
}
