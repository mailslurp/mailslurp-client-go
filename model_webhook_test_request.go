/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// WebhookTestRequest struct for WebhookTestRequest
type WebhookTestRequest struct {
	Url string `json:"url,omitempty"`
	Method string `json:"method,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Payload string `json:"payload,omitempty"`
}
