/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.  ## Resources  - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// DnsLookupOptions Options for DNS query. 
type DnsLookupOptions struct {
	// List of record types you wish to query such as MX, DNS, TXT, NS, A etc.
	Hostname string `json:"hostname"`
	// List of record types you wish to query such as MX, DNS, TXT, NS, A etc.
	RecordTypes []string `json:"recordTypes"`
	// Optionally control whether to omit the final dot in full DNS name values.
	OmitFinalDNSDot bool `json:"omitFinalDNSDot,omitempty"`
}
