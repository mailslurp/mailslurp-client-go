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
// ImapAccessDetails Access details for inbox using IMAP
type ImapAccessDetails struct {
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
}
