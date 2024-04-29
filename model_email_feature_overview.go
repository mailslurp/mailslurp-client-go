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
// EmailFeatureOverview struct for EmailFeatureOverview
type EmailFeatureOverview struct {
	Feature string `json:"feature"`
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Category string `json:"category,omitempty"`
	Notes string `json:"notes,omitempty"`
	NotesNumbers map[string]string `json:"notesNumbers,omitempty"`
	FeatureStatistics []EmailFeatureFamilyStatistics `json:"featureStatistics,omitempty"`
	Statuses []string `json:"statuses"`
}