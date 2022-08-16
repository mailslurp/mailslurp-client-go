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
// EmergencyAddress struct for EmergencyAddress
type EmergencyAddress struct {
	Id string `json:"id,omitempty"`
	Sid string `json:"sid"`
	UserId string `json:"userId"`
	DisplayName string `json:"displayName"`
	CustomerName string `json:"customerName"`
	Address1 string `json:"address1"`
	City string `json:"city"`
	Region string `json:"region"`
	PostalCode string `json:"postalCode"`
	PhoneCountry string `json:"phoneCountry"`
	AccountSid string `json:"accountSid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
