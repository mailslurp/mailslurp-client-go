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
// WebhookDto Representation of a webhook for an inbox. The URL specified will be using by MailSlurp whenever an email is received by the attached inbox. A webhook entity should have a URL that points to your server. Your server should accept HTTP/S POST requests and return a success 200. MailSlurp will retry your webhooks if they fail. See https://golang.api.mailslurp.com/schemas/webhook-payload for the payload schema.
type WebhookDto struct {
	// ID of the Webhook
	Id string `json:"id"`
	// User ID of the Webhook
	UserId string `json:"userId"`
	// Does webhook expect basic authentication? If true it means you created this webhook with a username and password. MailSlurp will use these in the URL to authenticate itself.
	BasicAuth bool `json:"basicAuth"`
	// Name of the webhook
	Name *string `json:"name,omitempty"`
	// The phoneNumberId that the Webhook will be triggered by. If null then webhook triggered at account level or inbox level if inboxId set
	PhoneId *string `json:"phoneId,omitempty"`
	// The inbox that the Webhook will be triggered by. If null then webhook triggered at account level or phone level if phoneId set
	InboxId *string `json:"inboxId,omitempty"`
	// Request body template for HTTP request that will be sent for the webhook. Use Moustache style template variables to insert values from the original event payload.
	RequestBodyTemplate *string `json:"requestBodyTemplate,omitempty"`
	// URL of your server that the webhook will be sent to. The schema of the JSON that is sent is described by the payloadJsonSchema.
	Url string `json:"url"`
	// HTTP method that your server endpoint must listen for
	Method string `json:"method"`
	// Deprecated. Fetch JSON Schema for webhook using the getJsonSchemaForWebhookPayload method
	PayloadJsonSchema string `json:"payloadJsonSchema"`
	// When the webhook was created
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// Webhook trigger event name
	EventName *string `json:"eventName,omitempty"`
	RequestHeaders WebhookHeaders `json:"requestHeaders,omitempty"`
	// Should notifier ignore insecure SSL certificates
	IgnoreInsecureSslCertificates *bool `json:"ignoreInsecureSslCertificates,omitempty"`
}
