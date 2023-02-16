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
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// EmailVerificationControllerApiService EmailVerificationControllerApi service
type EmailVerificationControllerApiService service

// GetValidationRequestsOpts Optional parameters for the method 'GetValidationRequests'
type GetValidationRequestsOpts struct {
    Page optional.Int32
    Size optional.Int32
    Sort optional.String
    SearchFilter optional.String
    Since optional.Time
    Before optional.Time
    IsValid optional.Bool
}

/*
GetValidationRequests Validate a list of email addresses. Per unit billing. See your plan for pricing.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetValidationRequestsOpts - Optional Parameters:
 * @param "Page" (optional.Int32) -  Optional page index in list pagination
 * @param "Size" (optional.Int32) -  Optional page size for paginated result list.
 * @param "Sort" (optional.String) -  Optional createdAt sort direction ASC or DESC
 * @param "SearchFilter" (optional.String) -  Optional search filter
 * @param "Since" (optional.Time) -  Filter by created at after the given timestamp
 * @param "Before" (optional.Time) -  Filter by created at before the given timestamp
 * @param "IsValid" (optional.Bool) -  Filter where email is valid is true or false
@return PageEmailValidationRequest
*/
func (a *EmailVerificationControllerApiService) GetValidationRequests(ctx _context.Context, localVarOptionals *GetValidationRequestsOpts) (PageEmailValidationRequest, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PageEmailValidationRequest
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/email-verification/validation-requests"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchFilter.IsSet() {
		localVarQueryParams.Add("searchFilter", parameterToString(localVarOptionals.SearchFilter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Since.IsSet() {
		localVarQueryParams.Add("since", parameterToString(localVarOptionals.Since.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Before.IsSet() {
		localVarQueryParams.Add("before", parameterToString(localVarOptionals.Before.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IsValid.IsSet() {
		localVarQueryParams.Add("isValid", parameterToString(localVarOptionals.IsValid.Value(), ""))
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
ValidateEmailAddressList Validate a list of email addresses. Per unit billing. See your plan for pricing.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param validateEmailAddressListOptions
@return ValidateEmailAddressListResult
*/
func (a *EmailVerificationControllerApiService) ValidateEmailAddressList(ctx _context.Context, validateEmailAddressListOptions ValidateEmailAddressListOptions) (ValidateEmailAddressListResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ValidateEmailAddressListResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/email-verification/email-address-list"
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
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &validateEmailAddressListOptions
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
