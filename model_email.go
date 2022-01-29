/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository
 *
 * API version: 6.5.2
 * Contact: contact@mailslurp.dev
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
import (
	"time"
)
// Email Email entity (also known as EmailDto). When an SMTP email message is received by MailSlurp it is parsed. The body and attachments are written to disk and the fields such as to, from, subject etc are stored in a database. The `body` contains the email content. If you want the original SMTP message see the `getRawEmail` endpoints. The attachments can be fetched using the AttachmentController
type Email struct {
	// ID of the email entity
	Id string `json:"id"`
	// ID of user that email belongs to
	UserId string `json:"userId"`
	// ID of the inbox that received the email
	InboxId string `json:"inboxId"`
	// List of `To` recipient email addresses that the email was addressed to. See recipients object for names.
	To []string `json:"to"`
	// Who the email was sent from. An email address - see fromName for the sender name.
	From string `json:"from,omitempty"`
	Sender Sender `json:"sender,omitempty"`
	Recipients EmailRecipients `json:"recipients,omitempty"`
	// The `replyTo` field on the received email message
	ReplyTo string `json:"replyTo,omitempty"`
	// List of `CC` recipients email addresses that the email was addressed to. See recipients object for names.
	Cc []string `json:"cc,omitempty"`
	// List of `BCC` recipients email addresses that the email was addressed to. See recipients object for names.
	Bcc []string `json:"bcc,omitempty"`
	// Collection of SMTP headers attached to email
	Headers map[string]string `json:"headers,omitempty"`
	// List of IDs of attachments found in the email. Use these IDs with the Inbox and Email Controllers to download attachments and attachment meta data such as filesize, name, extension.
	Attachments []string `json:"attachments,omitempty"`
	// The subject line of the email message as specified by SMTP subject header
	Subject string `json:"subject,omitempty"`
	// The body of the email message as text parsed from the SMTP message body (does not include attachments). Fetch the raw content to access the SMTP message and use the attachments property to access attachments. The body is stored separately to the email entity so the body is not returned in paginated results only in full single email or wait requests.
	Body string `json:"body,omitempty"`
	// An excerpt of the body of the email message for quick preview .
	BodyExcerpt string `json:"bodyExcerpt,omitempty"`
	// A hash signature of the email message using MD5. Useful for comparing emails without fetching full body.
	BodyMD5Hash string `json:"bodyMD5Hash,omitempty"`
	// Is the email body content type HTML?
	IsHTML bool `json:"isHTML,omitempty"`
	// Detected character set of the email body such as UTF-8
	Charset string `json:"charset,omitempty"`
	Analysis EmailAnalysis `json:"analysis,omitempty"`
	// When was the email received by MailSlurp
	CreatedAt time.Time `json:"createdAt"`
	// When was the email last updated
	UpdatedAt time.Time `json:"updatedAt"`
	// Read flag. Has the email ever been viewed in the dashboard or fetched via the API with a hydrated body? If so the email is marked as read. Paginated results do not affect read status. Read status is different to email opened event as it depends on your own account accessing the email. Email opened is determined by tracking pixels sent to other uses if enable during sending. You can listened for both email read and email opened events using webhooks.
	Read bool `json:"read"`
	// Can the email be accessed by organization team members
	TeamAccess bool `json:"teamAccess"`
}
