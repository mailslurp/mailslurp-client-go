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
import (
	"time"
)
// InboxDto Representation of a MailSlurp inbox. An inbox has an ID and a real email address. Emails can be sent to or from this email address. Inboxes are either `SMTP` or `HTTP` mailboxes. The default, `HTTP` inboxes, use AWS SES to process emails and are best suited as test email accounts and do not support IMAP or POP3. `SMTP` inboxes use a custom mail server at `mxslurp.click` and support SMTP login, IMAP and POP3. Use the `EmailController` or the `InboxController` methods to send and receive emails and attachments. Inboxes may have a description, name, and tags for display purposes. You can also favourite an inbox for easier searching.
type InboxDto struct {
	// ID of the inbox. The ID is a UUID-V4 format string. Use the inboxId for calls to Inbox and Email Controller endpoints. See the emailAddress property for the email address or the inbox. To get emails in an inbox use the WaitFor and Inbox Controller methods `waitForLatestEmail` and `getEmails` methods respectively. Inboxes can be used with aliases to forward emails automatically.
	Id string `json:"id"`
	// ID of user that inbox belongs to
	UserId *string `json:"userId"`
	// When the inbox was created. Time stamps are in ISO DateTime Format `yyyy-MM-dd'T'HH:mm:ss.SSSXXX` e.g. `2000-10-31T01:30:00.000-05:00`.
	CreatedAt time.Time `json:"createdAt"`
	// Name of the inbox and used as the sender name when sending emails .Displayed in the dashboard for easier search
	Name *string `json:"name,omitempty"`
	// ID of custom domain used by the inbox if any
	DomainId *string `json:"domainId,omitempty"`
	// Description of an inbox for labelling and searching purposes
	Description *string `json:"description,omitempty"`
	// The inbox's email address. Inbox projections and previews may not include the email address. To view the email address fetch the inbox entity directly. Send an email to this address and the inbox will receive and store it for you. Note the email address in MailSlurp match characters exactly and are case sensitive so `+123` additions are considered different addresses. To retrieve the email use the Inbox and Email Controller endpoints with the inbox ID.
	EmailAddress string `json:"emailAddress"`
	// Inbox expiration time. When, if ever, the inbox should expire and be deleted. If null then this inbox is permanent and the emails in it won't be deleted. This is the default behavior unless expiration date is set. If an expiration date is set and the time is reached MailSlurp will expire the inbox and move it to an expired inbox entity. You can still access the emails belonging to it but it can no longer send or receive email.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Is the inbox a favorite inbox. Make an inbox a favorite is typically done in the dashboard for quick access or filtering
	Favourite bool `json:"favourite"`
	// Tags that inbox has been tagged with. Tags can be added to inboxes to group different inboxes within an account. You can also search for inboxes by tag in the dashboard UI.
	Tags *[]string `json:"tags,omitempty"`
	// Type of inbox. HTTP inboxes are faster and better for most cases. SMTP inboxes are more suited for public facing inbound messages (but cannot send).
	InboxType *string `json:"inboxType,omitempty"`
	// Is the inbox readOnly for the caller. Read only means can not be deleted or modified. This flag is present when using team accounts and shared inboxes.
	ReadOnly bool `json:"readOnly"`
	// Virtual inbox can receive email but will not send emails to real recipients. Will save sent email record but never send an actual email. Perfect for testing mail server actions.
	VirtualInbox bool `json:"virtualInbox"`
	// Inbox function if used as a primitive for another system.
	FunctionsAs *string `json:"functionsAs,omitempty"`
	// Local part of email addresses before the @ symbol
	LocalPart *string `json:"localPart,omitempty"`
	// Domain name of the email address
	Domain *string `json:"domain,omitempty"`
	// Region of the inbox
	AccountRegion *string `json:"accountRegion,omitempty"`
}
