# SendEmailOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToContacts** | Pointer to **[]string** | Optional list of contact IDs to send email to. Manage your contacts via the API or dashboard. When contacts are used the email is sent to each contact separately so they will not see other recipients. | [optional] 
**ToGroup** | Pointer to **string** | Optional contact group ID to send email to. You can create contacts and contact groups in the API or dashboard and use them for email campaigns. When contact groups are used the email is sent to each contact separately so they will not see other recipients | [optional] 
**To** | Pointer to **[]string** | List of destination email addresses. Each email address must be RFC 5322 format. Even single recipients must be in array form. Maximum recipients per email depends on your plan. If you need to send many emails try using contacts or contact groups or use a non standard sendStrategy to ensure that spam filters are not triggered (many recipients in one email can affect your spam rating). Be cautious when sending emails that your recipients exist. High bounce rates (meaning a high percentage of emails cannot be delivered because an address does not exist) can result in account freezing. | [optional] 
**From** | Pointer to **string** | Optional from address. Email address is RFC 5322 format and may include a display name and email in angle brackets (&#x60;my@address.com&#x60; or &#x60;My inbox &lt;my@address.com&gt;&#x60;). If no sender is set the source inbox address will be used for this field. If you set &#x60;useInboxName&#x60; to &#x60;true&#x60; the from field will include the inbox name as a display name: &#x60;inbox_name &lt;inbox@address.com&gt;&#x60;. For this to work use the name field when creating an inbox. Beware of potential spam penalties when setting the from field to an address not used by the inbox. Your emails may get blocked by services if you impersonate another address. To use a custom email addresses use a custom domain. You can create domains with the DomainController. The domain must be verified in the dashboard before it can be used. | [optional] 
**FromName** | Pointer to **string** | Optional from name if not passed with address. If you set &#x60;useInboxName&#x60; to &#x60;true&#x60; the from field will include the inbox name as a display name | [optional] 
**Cc** | Pointer to **[]string** | Optional list of cc destination email addresses | [optional] 
**Bcc** | Pointer to **[]string** | Optional list of bcc destination email addresses | [optional] 
**Subject** | Pointer to **string** | Optional email subject line | [optional] 
**ReplyTo** | Pointer to **string** | Optional replyTo header | [optional] 
**CustomHeaders** | Pointer to **map[string]string** | Optional custom headers | [optional] 
**Body** | Pointer to **string** | Optional contents of email. If body contains HTML then set &#x60;isHTML&#x60; to true to ensure that email clients render it correctly. You can use moustache template syntax in the email body in conjunction with &#x60;toGroup&#x60; contact variables or &#x60;templateVariables&#x60; data. If you need more templating control consider creating a template and using the &#x60;template&#x60; property instead of the body. | [optional] 
**Html** | Pointer to **bool** | Optional HTML flag to indicate that contents is HTML. Set&#39;s a &#x60;content-type: text/html&#x60; for email. (Deprecated: use &#x60;isHTML&#x60; instead.) | [optional] 
**IsHTML** | Pointer to **bool** | Optional HTML flag. If true the &#x60;content-type&#x60; of the email will be &#x60;text/html&#x60;. Set to true when sending HTML to ensure proper rending on email clients | [optional] 
**Charset** | Pointer to **string** | Optional charset | [optional] 
**Attachments** | Pointer to **[]string** | Optional list of attachment IDs to send with this email. You must first upload each attachment separately via method call or dashboard in order to obtain attachment IDs. This way you can reuse attachments with different emails once uploaded. There are several ways to upload that support &#x60;multi-part form&#x60;, &#x60;base64 file encoding&#x60;, and octet stream binary uploads. See the &#x60;UploadController&#x60; for available methods.  | [optional] 
**TemplateVariables** | Pointer to **map[string]map[string]interface{}** | Optional map of template variables. Will replace moustache syntax variables in subject and body or template with the associated values if found. | [optional] 
**Template** | Pointer to **string** | Optional template ID to use for body. Will override body if provided. When using a template make sure you pass the corresponding map of &#x60;templateVariables&#x60;. You can find which variables are needed by fetching the template itself or viewing it in the dashboard. | [optional] 
**SendStrategy** | Pointer to **string** | How an email should be sent based on its recipients | [optional] 
**UseInboxName** | Pointer to **bool** | Use name of inbox as sender email address name. Will construct RFC 5322 email address with &#x60;Inbox name &lt;inbox@address.com&gt;&#x60; if the inbox has a name. | [optional] 
**AddTrackingPixel** | Pointer to **bool** | Add tracking pixel to email | [optional] 
**FilterBouncedRecipients** | Pointer to **bool** | Filter recipients to remove any bounced recipients from to, bcc, and cc before sending | [optional] 
**ValidateEmailAddresses** | Pointer to **string** | Validate recipient email addresses before sending | [optional] 
**IgnoreEmptyRecipients** | Pointer to **bool** | Ignore empty recipients after validation removes all recipients as invalid and fail silently | [optional] 
**IsXAmpHtml** | Pointer to **bool** | Is content AMP4EMAIL compatible. If set will send as x-amp-html part. | [optional] 
**BodyParts** | Pointer to [**[]SendEmailBodyPart**](SendEmailBodyPart) | Email body content parts for multipart mime message. Will override body. | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


