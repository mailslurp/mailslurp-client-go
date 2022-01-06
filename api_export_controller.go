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
)

// Linger please
var (
	_ _context.Context
)

// ExportControllerApiService ExportControllerApi service
type ExportControllerApiService service

// ExportEntitiesOpts Optional parameters for the method 'ExportEntities'
type ExportEntitiesOpts struct {
    Filter optional.String
    ListSeparatorToken optional.String
    ExcludePreviouslyExported optional.Bool
    CreatedEarliestTime optional.Time
    CreatedOldestTime optional.Time
}

/*
ExportEntities Export inboxes link callable via browser
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param exportType
 * @param apiKey
 * @param outputFormat
 * @param optional nil or *ExportEntitiesOpts - Optional Parameters:
 * @param "Filter" (optional.String) - 
 * @param "ListSeparatorToken" (optional.String) - 
 * @param "ExcludePreviouslyExported" (optional.Bool) - 
 * @param "CreatedEarliestTime" (optional.Time) - 
 * @param "CreatedOldestTime" (optional.Time) - 
@return []string
*/
func (a *ExportControllerApiService) ExportEntities(ctx _context.Context, exportType string, apiKey string, outputFormat string, localVarOptionals *ExportEntitiesOpts) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/export"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("exportType", parameterToString(exportType, ""))
	localVarQueryParams.Add("apiKey", parameterToString(apiKey, ""))
	localVarQueryParams.Add("outputFormat", parameterToString(outputFormat, ""))
	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ListSeparatorToken.IsSet() {
		localVarQueryParams.Add("listSeparatorToken", parameterToString(localVarOptionals.ListSeparatorToken.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludePreviouslyExported.IsSet() {
		localVarQueryParams.Add("excludePreviouslyExported", parameterToString(localVarOptionals.ExcludePreviouslyExported.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreatedEarliestTime.IsSet() {
		localVarQueryParams.Add("createdEarliestTime", parameterToString(localVarOptionals.CreatedEarliestTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreatedOldestTime.IsSet() {
		localVarQueryParams.Add("createdOldestTime", parameterToString(localVarOptionals.CreatedOldestTime.Value(), ""))
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

// GetExportLinkOpts Optional parameters for the method 'GetExportLink'
type GetExportLinkOpts struct {
    ApiKey optional.String
}

/*
GetExportLink Get export link
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param exportType
 * @param exportOptions
 * @param optional nil or *GetExportLinkOpts - Optional Parameters:
 * @param "ApiKey" (optional.String) - 
@return ExportLink
*/
func (a *ExportControllerApiService) GetExportLink(ctx _context.Context, exportType string, exportOptions ExportOptions, localVarOptionals *GetExportLinkOpts) (ExportLink, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ExportLink
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/export"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("exportType", parameterToString(exportType, ""))
	if localVarOptionals != nil && localVarOptionals.ApiKey.IsSet() {
		localVarQueryParams.Add("apiKey", parameterToString(localVarOptionals.ApiKey.Value(), ""))
	}
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
	localVarPostBody = &exportOptions
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
