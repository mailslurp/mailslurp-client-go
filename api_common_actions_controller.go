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
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"github.com/antihax/optional"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// CommonActionsControllerApiService CommonActionsControllerApi service
type CommonActionsControllerApiService service

// CreateNewEmailAddressOpts Optional parameters for the method 'CreateNewEmailAddress'
type CreateNewEmailAddressOpts struct {
    AllowTeamAccess optional.Bool
    UseDomainPool optional.Bool
    ExpiresAt optional.Time
    ExpiresIn optional.Int64
    EmailAddress optional.String
    InboxType optional.String
    Description optional.String
    Name optional.String
    Tags optional.Interface
    Favourite optional.Bool
}

/*
CreateNewEmailAddress Create new random inbox
Returns an Inbox with an &#x60;id&#x60; and an &#x60;emailAddress&#x60;
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CreateNewEmailAddressOpts - Optional Parameters:
 * @param "AllowTeamAccess" (optional.Bool) - 
 * @param "UseDomainPool" (optional.Bool) - 
 * @param "ExpiresAt" (optional.Time) - 
 * @param "ExpiresIn" (optional.Int64) - 
 * @param "EmailAddress" (optional.String) - 
 * @param "InboxType" (optional.String) - 
 * @param "Description" (optional.String) - 
 * @param "Name" (optional.String) - 
 * @param "Tags" (optional.Interface of []string) - 
 * @param "Favourite" (optional.Bool) - 
@return InboxDto
*/
func (a *CommonActionsControllerApiService) CreateNewEmailAddress(ctx _context.Context, localVarOptionals *CreateNewEmailAddressOpts) (InboxDto, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InboxDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/createInbox"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.AllowTeamAccess.IsSet() {
		localVarQueryParams.Add("allowTeamAccess", parameterToString(localVarOptionals.AllowTeamAccess.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UseDomainPool.IsSet() {
		localVarQueryParams.Add("useDomainPool", parameterToString(localVarOptionals.UseDomainPool.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExpiresAt.IsSet() {
		localVarQueryParams.Add("expiresAt", parameterToString(localVarOptionals.ExpiresAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExpiresIn.IsSet() {
		localVarQueryParams.Add("expiresIn", parameterToString(localVarOptionals.ExpiresIn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EmailAddress.IsSet() {
		localVarQueryParams.Add("emailAddress", parameterToString(localVarOptionals.EmailAddress.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.InboxType.IsSet() {
		localVarQueryParams.Add("inboxType", parameterToString(localVarOptionals.InboxType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Description.IsSet() {
		localVarQueryParams.Add("description", parameterToString(localVarOptionals.Description.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Name.IsSet() {
		localVarQueryParams.Add("name", parameterToString(localVarOptionals.Name.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tags.IsSet() {
		t:=localVarOptionals.Tags.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("tags", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("tags", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.Favourite.IsSet() {
		localVarQueryParams.Add("favourite", parameterToString(localVarOptionals.Favourite.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// CreateNewEmailAddress1Opts Optional parameters for the method 'CreateNewEmailAddress1'
type CreateNewEmailAddress1Opts struct {
    AllowTeamAccess optional.Bool
    UseDomainPool optional.Bool
    ExpiresAt optional.Time
    ExpiresIn optional.Int64
    EmailAddress optional.String
    InboxType optional.String
    Description optional.String
    Name optional.String
    Tags optional.Interface
    Favourite optional.Bool
}

/*
CreateNewEmailAddress1 Create new random inbox
Returns an Inbox with an &#x60;id&#x60; and an &#x60;emailAddress&#x60;
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CreateNewEmailAddress1Opts - Optional Parameters:
 * @param "AllowTeamAccess" (optional.Bool) - 
 * @param "UseDomainPool" (optional.Bool) - 
 * @param "ExpiresAt" (optional.Time) - 
 * @param "ExpiresIn" (optional.Int64) - 
 * @param "EmailAddress" (optional.String) - 
 * @param "InboxType" (optional.String) - 
 * @param "Description" (optional.String) - 
 * @param "Name" (optional.String) - 
 * @param "Tags" (optional.Interface of []string) - 
 * @param "Favourite" (optional.Bool) - 
@return InboxDto
*/
func (a *CommonActionsControllerApiService) CreateNewEmailAddress1(ctx _context.Context, localVarOptionals *CreateNewEmailAddress1Opts) (InboxDto, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InboxDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/newEmailAddress"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.AllowTeamAccess.IsSet() {
		localVarQueryParams.Add("allowTeamAccess", parameterToString(localVarOptionals.AllowTeamAccess.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UseDomainPool.IsSet() {
		localVarQueryParams.Add("useDomainPool", parameterToString(localVarOptionals.UseDomainPool.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExpiresAt.IsSet() {
		localVarQueryParams.Add("expiresAt", parameterToString(localVarOptionals.ExpiresAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExpiresIn.IsSet() {
		localVarQueryParams.Add("expiresIn", parameterToString(localVarOptionals.ExpiresIn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EmailAddress.IsSet() {
		localVarQueryParams.Add("emailAddress", parameterToString(localVarOptionals.EmailAddress.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.InboxType.IsSet() {
		localVarQueryParams.Add("inboxType", parameterToString(localVarOptionals.InboxType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Description.IsSet() {
		localVarQueryParams.Add("description", parameterToString(localVarOptionals.Description.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Name.IsSet() {
		localVarQueryParams.Add("name", parameterToString(localVarOptionals.Name.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tags.IsSet() {
		t:=localVarOptionals.Tags.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("tags", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("tags", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.Favourite.IsSet() {
		localVarQueryParams.Add("favourite", parameterToString(localVarOptionals.Favourite.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
DeleteEmailAddress Delete inbox email address by inbox id
Deletes inbox email address
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param inboxId
*/
func (a *CommonActionsControllerApiService) DeleteEmailAddress(ctx _context.Context, inboxId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/deleteEmailAddress"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("inboxId", parameterToString(inboxId, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
EmptyInbox Delete all emails in an inbox
Deletes all emails
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param inboxId
*/
func (a *CommonActionsControllerApiService) EmptyInbox(ctx _context.Context, inboxId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/emptyInbox"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("inboxId", parameterToString(inboxId, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
SendEmailSimple Send an email
If no senderId or inboxId provided a random email address will be used to send from.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param simpleSendEmailOptions
*/
func (a *CommonActionsControllerApiService) SendEmailSimple(ctx _context.Context, simpleSendEmailOptions SimpleSendEmailOptions) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sendEmail"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &simpleSendEmailOptions
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
