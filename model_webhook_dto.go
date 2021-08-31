/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
import (
	"time"
)
// WebhookDto Representation of a webhook for an inbox. The URL specified will be using by MailSlurp whenever an email is received by the attached inbox. A webhook entity should have a URL that points to your server. Your server should accept HTTP/S POST requests and return a success 200. MailSlurp will retry your webhooks if they fail. See https://api.mailslurp.com/schemas/webhook-payload for the payload schema.
type WebhookDto struct {
	// Does webhook expect basic authentication? If true it means you created this webhook with a username and password. MailSlurp will use these in the URL to authenticate itself.
	BasicAuth bool `json:"basicAuth,omitempty"`
	// When the webhook was created
	CreatedAt time.Time `json:"createdAt,omitempty"`
	EventName string `json:"eventName,omitempty"`
	// ID of the Webhook
	Id string `json:"id,omitempty"`
	// The inbox that the Webhook will be triggered by
	InboxId string `json:"inboxId,omitempty"`
	// HTTP method that your server endpoint must listen for
	Method string `json:"method,omitempty"`
	// Name of the webhook
	Name string `json:"name,omitempty"`
	// JSON Schema for the payload that will be sent to your URL via the HTTP method described.
	PayloadJsonSchema string `json:"payloadJsonSchema,omitempty"`
	UpdatedAt time.Time `json:"updatedAt"`
	// URL of your server that the webhook will be sent to. The schema of the JSON that is sent is described by the payloadJsonSchema.
	Url string `json:"url,omitempty"`
	// User ID of the Webhook
	UserId string `json:"userId,omitempty"`
}
