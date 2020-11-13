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
// ContactProjection struct for ContactProjection
type ContactProjection struct {
	Company string `json:"company,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	FirstName string `json:"firstName,omitempty"`
	GroupId string `json:"groupId,omitempty"`
	Id string `json:"id"`
	LastName string `json:"lastName,omitempty"`
	OptOut bool `json:"optOut,omitempty"`
}
