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
// CreateContactOptions Options for creating an email contact in address book
type CreateContactOptions struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	Company *string `json:"company,omitempty"`
	// Set of email addresses belonging to the contact
	EmailAddresses *[]string `json:"emailAddresses,omitempty"`
	// Tags that can be used to search and group contacts
	Tags *[]string `json:"tags,omitempty"`
	MetaData *map[string]interface{} `json:"metaData,omitempty"`
	// Has the user explicitly or implicitly opted out of being contacted? If so MailSlurp will ignore them in all actions.
	OptOut *bool `json:"optOut,omitempty"`
	// Group IDs that contact belongs to
	GroupId *string `json:"groupId,omitempty"`
	// Whether to validate contact email address exists
	VerifyEmailAddresses *bool `json:"verifyEmailAddresses,omitempty"`
}
