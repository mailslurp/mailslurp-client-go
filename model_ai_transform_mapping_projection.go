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
// AiTransformMappingProjection struct for AiTransformMappingProjection
type AiTransformMappingProjection struct {
	UserId string `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	AiTransformId string `json:"aiTransformId"`
	EntityId string `json:"entityId,omitempty"`
	EntityType string `json:"entityType"`
	ContentSelector string `json:"contentSelector,omitempty"`
	TriggerSelector string `json:"triggerSelector,omitempty"`
	Name string `json:"name,omitempty"`
	Id string `json:"id"`
}
