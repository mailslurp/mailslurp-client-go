/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
import (
	"time"
)
// DomainDto Domain plus verification records and status
type DomainDto struct {
	CreatedAt time.Time `json:"createdAt"`
	// DNS records for DKIM approval
	DkimTokens []string `json:"dkimTokens,omitempty"`
	// Custom domain name
	Domain string `json:"domain,omitempty"`
	Id string `json:"id"`
	// Whether domain has been verified or not. If the domain is not verified after 72 hours there is most likely an issue with the domains DNS records.
	IsVerified bool `json:"isVerified,omitempty"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserId string `json:"userId"`
	// A TXT record that you must place in the DNS settings of the domain to complete domain verification
	VerificationToken string `json:"verificationToken,omitempty"`
}
