/*
Salesforce Platform REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 56.0
Contact: russell-laboe@outlook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tooling

import (
	"bytes"
	"context"
    "compress/gzip"
	"io/ioutil"
	"net/http"
	"net/url"
    "strings"
)


type DefaultApi interface {

	/*
	ExecuteAnonymous Executes anonymous Apex

	Executes Apex code anonymously. Available from API version 29.0 or later.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExecuteAnonymousRequest
	*/
	ExecuteAnonymous(ctx context.Context) ApiExecuteAnonymousRequest

	// ExecuteAnonymousExecute executes the request
	//  @return ExecuteAnonymousResult
	ExecuteAnonymousExecute(r ApiExecuteAnonymousRequest) (*ExecuteAnonymousResult, *http.Response, error)

	/*
	GetCompletions Retrieve available code completions

	Retrieves available code completions of the referenced type for Apex system method symbols (type=apex). Available from API version 28.0 or later. Retrieves available code completions of the referenced type for Visualforce markup (type=visualforce). Available from API version 38.0 or later.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCompletionsRequest
	*/
	GetCompletions(ctx context.Context) ApiGetCompletionsRequest

	// GetCompletionsExecute executes the request
	//  @return interface{}
	GetCompletionsExecute(r ApiGetCompletionsRequest) (interface{}, *http.Response, error)

	/*
	GetSObject Describe SObject metadata

	Describes the individual metadata for the specified object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sObjectName
	@return ApiGetSObjectRequest
	*/
	GetSObject(ctx context.Context, sObjectName string) ApiGetSObjectRequest

	// GetSObjectExecute executes the request
	//  @return SObjectResult
	GetSObjectExecute(r ApiGetSObjectRequest) (*SObjectResult, *http.Response, error)

	/*
	GetSObjects List SObjects

	Lists the available Tooling API objects and their metadata.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetSObjectsRequest
	*/
	GetSObjects(ctx context.Context) ApiGetSObjectsRequest

	// GetSObjectsExecute executes the request
	//  @return DescribeGlobalResult
	GetSObjectsExecute(r ApiGetSObjectsRequest) (*DescribeGlobalResult, *http.Response, error)

	/*
	Query Executes query

	Executes a query against an object and returns data that matches the specified criteria. Tooling API exposes objects like EntityDefinition and FieldDefinition that use the external object framework. That is, they don’t exist in the database but are constructed dynamically. Special query rules apply to virtual entities. If the query result is too large, it’s broken up into batches. The response contains the first batch of results and a query identifier. The identifier can be used in a request to retrieve the next batch.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiQueryRequest
	*/
	Query(ctx context.Context) ApiQueryRequest

	// QueryExecute executes the request
	//  @return QueryResult
	QueryExecute(r ApiQueryRequest) (*QueryResult, *http.Response, error)

	/*
	RunTestsAsync Run tests asynchronously

	Runs one or more methods within one or more Apex classes, using the asynchronous test execution mechanism. In the request body, you can specify test class names and IDs, suite names and IDs, the maximum number of failed tests to allow, and the test level, as comma-separated lists or as an array. You can also indicate whether to opt out of collecting code coverage information during the test run (available in API version 43.0 and later).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRunTestsAsyncRequest
	*/
	RunTestsAsync(ctx context.Context) ApiRunTestsAsyncRequest

	// RunTestsAsyncExecute executes the request
	//  @return string
	RunTestsAsyncExecute(r ApiRunTestsAsyncRequest) (string, *http.Response, error)

	/*
	RunTestsSync Run tests synchronously

	Runs one or more methods within an Apex class, using the synchronous test execution mechanism. All test methods in a synchronous test run must be in the same class.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRunTestsSyncRequest
	*/
	RunTestsSync(ctx context.Context) ApiRunTestsSyncRequest

	// RunTestsSyncExecute executes the request
	//  @return RunTestsResult
	RunTestsSyncExecute(r ApiRunTestsSyncRequest) (*RunTestsResult, *http.Response, error)

	/*
	Search Executes SOSL

	Search for records containing specified values.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchRequest
	*/
	Search(ctx context.Context) ApiSearchRequest

	// SearchExecute executes the request
	//  @return SOSLResult
	SearchExecute(r ApiSearchRequest) (*SOSLResult, *http.Response, error)
}

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiExecuteAnonymousRequest struct {
	ctx context.Context
	ApiService DefaultApi
	anonymousBody *string
}

func (r ApiExecuteAnonymousRequest) AnonymousBody(anonymousBody string) ApiExecuteAnonymousRequest {
	r.anonymousBody = &anonymousBody
	return r
}

func (r ApiExecuteAnonymousRequest) Execute() (*ExecuteAnonymousResult, *http.Response, error) {
	return r.ApiService.ExecuteAnonymousExecute(r)
}

/*
ExecuteAnonymous Executes anonymous Apex

Executes Apex code anonymously. Available from API version 29.0 or later.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExecuteAnonymousRequest
*/
func (a *DefaultApiService) ExecuteAnonymous(ctx context.Context) ApiExecuteAnonymousRequest {
	return ApiExecuteAnonymousRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ExecuteAnonymousResult
func (a *DefaultApiService) ExecuteAnonymousExecute(r ApiExecuteAnonymousRequest) (*ExecuteAnonymousResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExecuteAnonymousResult
	)

    token, err := a.client.TokenSrc.Token()
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    activeServer := a.client.cfg.GetActiveServer()
    activeServer.SetServerVariable("instanceUrl", token.Extra("instance_url").(string))

	localVarPath := activeServer.GetURL() + "/executeAnonymous"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.anonymousBody == nil {
		return localVarReturnValue, nil, reportError("anonymousBody is required and must be specified")
	}

	localVarQueryParams.Add("anonymousBody", parameterToString(*r.anonymousBody, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

    token.SetAuthHeader(req)

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}


    respBodyReadCloser := localVarHTTPResponse.Body
    if localVarHTTPResponse.Header.Get("Content-Encoding") == "gzip" {
        respBodyReadCloser, err = gzip.NewReader(localVarHTTPResponse.Body)
    }

	if localVarHTTPResponse.StatusCode >= 300 {
        localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
        localVarHTTPResponse.Body.Close()
        respBodyReadCloser.Close()
        localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
        if err != nil {
            return localVarReturnValue, localVarHTTPResponse, err
        }
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
    localVarHTTPResponse.Body.Close()
    respBodyReadCloser.Close()
    localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCompletionsRequest struct {
	ctx context.Context
	ApiService DefaultApi
	type_ *CompletionType
}

func (r ApiGetCompletionsRequest) Type_(type_ CompletionType) ApiGetCompletionsRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetCompletionsRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.GetCompletionsExecute(r)
}

/*
GetCompletions Retrieve available code completions

Retrieves available code completions of the referenced type for Apex system method symbols (type=apex). Available from API version 28.0 or later. Retrieves available code completions of the referenced type for Visualforce markup (type=visualforce). Available from API version 38.0 or later.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCompletionsRequest
*/
func (a *DefaultApiService) GetCompletions(ctx context.Context) ApiGetCompletionsRequest {
	return ApiGetCompletionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *DefaultApiService) GetCompletionsExecute(r ApiGetCompletionsRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

    token, err := a.client.TokenSrc.Token()
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    activeServer := a.client.cfg.GetActiveServer()
    activeServer.SetServerVariable("instanceUrl", token.Extra("instance_url").(string))

	localVarPath := activeServer.GetURL() + "/completions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.type_ == nil {
		return localVarReturnValue, nil, reportError("type_ is required and must be specified")
	}

	localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

    token.SetAuthHeader(req)

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}


    respBodyReadCloser := localVarHTTPResponse.Body
    if localVarHTTPResponse.Header.Get("Content-Encoding") == "gzip" {
        respBodyReadCloser, err = gzip.NewReader(localVarHTTPResponse.Body)
    }

	if localVarHTTPResponse.StatusCode >= 300 {
        localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
        localVarHTTPResponse.Body.Close()
        respBodyReadCloser.Close()
        localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
        if err != nil {
            return localVarReturnValue, localVarHTTPResponse, err
        }
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
    localVarHTTPResponse.Body.Close()
    respBodyReadCloser.Close()
    localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSObjectRequest struct {
	ctx context.Context
	ApiService DefaultApi
	sObjectName string
}

func (r ApiGetSObjectRequest) Execute() (*SObjectResult, *http.Response, error) {
	return r.ApiService.GetSObjectExecute(r)
}

/*
GetSObject Describe SObject metadata

Describes the individual metadata for the specified object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sObjectName
 @return ApiGetSObjectRequest
*/
func (a *DefaultApiService) GetSObject(ctx context.Context, sObjectName string) ApiGetSObjectRequest {
	return ApiGetSObjectRequest{
		ApiService: a,
		ctx: ctx,
		sObjectName: sObjectName,
	}
}

// Execute executes the request
//  @return SObjectResult
func (a *DefaultApiService) GetSObjectExecute(r ApiGetSObjectRequest) (*SObjectResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SObjectResult
	)

    token, err := a.client.TokenSrc.Token()
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    activeServer := a.client.cfg.GetActiveServer()
    activeServer.SetServerVariable("instanceUrl", token.Extra("instance_url").(string))

	localVarPath := activeServer.GetURL() + "/sobjects/{SObjectName}"
    pathParams := map[string]string {
        "SObjectName": url.PathEscape(parameterToString(strings.Trim(r.sObjectName, " "), "")),
    }
    localVarPath = injectUrlVars(localVarPath, pathParams)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

    token.SetAuthHeader(req)

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}


    respBodyReadCloser := localVarHTTPResponse.Body
    if localVarHTTPResponse.Header.Get("Content-Encoding") == "gzip" {
        respBodyReadCloser, err = gzip.NewReader(localVarHTTPResponse.Body)
    }

	if localVarHTTPResponse.StatusCode >= 300 {
        localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
        localVarHTTPResponse.Body.Close()
        respBodyReadCloser.Close()
        localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
        if err != nil {
            return localVarReturnValue, localVarHTTPResponse, err
        }
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
    localVarHTTPResponse.Body.Close()
    respBodyReadCloser.Close()
    localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSObjectsRequest struct {
	ctx context.Context
	ApiService DefaultApi
}

func (r ApiGetSObjectsRequest) Execute() (*DescribeGlobalResult, *http.Response, error) {
	return r.ApiService.GetSObjectsExecute(r)
}

/*
GetSObjects List SObjects

Lists the available Tooling API objects and their metadata.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSObjectsRequest
*/
func (a *DefaultApiService) GetSObjects(ctx context.Context) ApiGetSObjectsRequest {
	return ApiGetSObjectsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DescribeGlobalResult
func (a *DefaultApiService) GetSObjectsExecute(r ApiGetSObjectsRequest) (*DescribeGlobalResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DescribeGlobalResult
	)

    token, err := a.client.TokenSrc.Token()
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    activeServer := a.client.cfg.GetActiveServer()
    activeServer.SetServerVariable("instanceUrl", token.Extra("instance_url").(string))

	localVarPath := activeServer.GetURL() + "/sobjects"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

    token.SetAuthHeader(req)

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}


    respBodyReadCloser := localVarHTTPResponse.Body
    if localVarHTTPResponse.Header.Get("Content-Encoding") == "gzip" {
        respBodyReadCloser, err = gzip.NewReader(localVarHTTPResponse.Body)
    }

	if localVarHTTPResponse.StatusCode >= 300 {
        localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
        localVarHTTPResponse.Body.Close()
        respBodyReadCloser.Close()
        localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
        if err != nil {
            return localVarReturnValue, localVarHTTPResponse, err
        }
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
    localVarHTTPResponse.Body.Close()
    respBodyReadCloser.Close()
    localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiQueryRequest struct {
	ctx context.Context
	ApiService DefaultApi
	q *string
}

// SOQL query statement
func (r ApiQueryRequest) Q(q string) ApiQueryRequest {
	r.q = &q
	return r
}

func (r ApiQueryRequest) Execute() (*QueryResult, *http.Response, error) {
	return r.ApiService.QueryExecute(r)
}

/*
Query Executes query

Executes a query against an object and returns data that matches the specified criteria. Tooling API exposes objects like EntityDefinition and FieldDefinition that use the external object framework. That is, they don’t exist in the database but are constructed dynamically. Special query rules apply to virtual entities. If the query result is too large, it’s broken up into batches. The response contains the first batch of results and a query identifier. The identifier can be used in a request to retrieve the next batch.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryRequest
*/
func (a *DefaultApiService) Query(ctx context.Context) ApiQueryRequest {
	return ApiQueryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QueryResult
func (a *DefaultApiService) QueryExecute(r ApiQueryRequest) (*QueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QueryResult
	)

    token, err := a.client.TokenSrc.Token()
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    activeServer := a.client.cfg.GetActiveServer()
    activeServer.SetServerVariable("instanceUrl", token.Extra("instance_url").(string))

	localVarPath := activeServer.GetURL() + "/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.q == nil {
		return localVarReturnValue, nil, reportError("q is required and must be specified")
	}

	localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

    token.SetAuthHeader(req)

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}


    respBodyReadCloser := localVarHTTPResponse.Body
    if localVarHTTPResponse.Header.Get("Content-Encoding") == "gzip" {
        respBodyReadCloser, err = gzip.NewReader(localVarHTTPResponse.Body)
    }

	if localVarHTTPResponse.StatusCode >= 300 {
        localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
        localVarHTTPResponse.Body.Close()
        respBodyReadCloser.Close()
        localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
        if err != nil {
            return localVarReturnValue, localVarHTTPResponse, err
        }
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
    localVarHTTPResponse.Body.Close()
    respBodyReadCloser.Close()
    localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRunTestsAsyncRequest struct {
	ctx context.Context
	ApiService DefaultApi
	asyncTestRequest *AsyncTestRequest
}

func (r ApiRunTestsAsyncRequest) AsyncTestRequest(asyncTestRequest AsyncTestRequest) ApiRunTestsAsyncRequest {
	r.asyncTestRequest = &asyncTestRequest
	return r
}

func (r ApiRunTestsAsyncRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.RunTestsAsyncExecute(r)
}

/*
RunTestsAsync Run tests asynchronously

Runs one or more methods within one or more Apex classes, using the asynchronous test execution mechanism. In the request body, you can specify test class names and IDs, suite names and IDs, the maximum number of failed tests to allow, and the test level, as comma-separated lists or as an array. You can also indicate whether to opt out of collecting code coverage information during the test run (available in API version 43.0 and later).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRunTestsAsyncRequest
*/
func (a *DefaultApiService) RunTestsAsync(ctx context.Context) ApiRunTestsAsyncRequest {
	return ApiRunTestsAsyncRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *DefaultApiService) RunTestsAsyncExecute(r ApiRunTestsAsyncRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

    token, err := a.client.TokenSrc.Token()
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    activeServer := a.client.cfg.GetActiveServer()
    activeServer.SetServerVariable("instanceUrl", token.Extra("instance_url").(string))

	localVarPath := activeServer.GetURL() + "/runTestsAsynchronous"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.asyncTestRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

    token.SetAuthHeader(req)

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}


    respBodyReadCloser := localVarHTTPResponse.Body
    if localVarHTTPResponse.Header.Get("Content-Encoding") == "gzip" {
        respBodyReadCloser, err = gzip.NewReader(localVarHTTPResponse.Body)
    }

	if localVarHTTPResponse.StatusCode >= 300 {
        localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
        localVarHTTPResponse.Body.Close()
        respBodyReadCloser.Close()
        localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
        if err != nil {
            return localVarReturnValue, localVarHTTPResponse, err
        }
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
    localVarHTTPResponse.Body.Close()
    respBodyReadCloser.Close()
    localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRunTestsSyncRequest struct {
	ctx context.Context
	ApiService DefaultApi
	testRequest *TestRequest
}

func (r ApiRunTestsSyncRequest) TestRequest(testRequest TestRequest) ApiRunTestsSyncRequest {
	r.testRequest = &testRequest
	return r
}

func (r ApiRunTestsSyncRequest) Execute() (*RunTestsResult, *http.Response, error) {
	return r.ApiService.RunTestsSyncExecute(r)
}

/*
RunTestsSync Run tests synchronously

Runs one or more methods within an Apex class, using the synchronous test execution mechanism. All test methods in a synchronous test run must be in the same class.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRunTestsSyncRequest
*/
func (a *DefaultApiService) RunTestsSync(ctx context.Context) ApiRunTestsSyncRequest {
	return ApiRunTestsSyncRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RunTestsResult
func (a *DefaultApiService) RunTestsSyncExecute(r ApiRunTestsSyncRequest) (*RunTestsResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RunTestsResult
	)

    token, err := a.client.TokenSrc.Token()
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    activeServer := a.client.cfg.GetActiveServer()
    activeServer.SetServerVariable("instanceUrl", token.Extra("instance_url").(string))

	localVarPath := activeServer.GetURL() + "/runTestsSynchronous"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.testRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

    token.SetAuthHeader(req)

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}


    respBodyReadCloser := localVarHTTPResponse.Body
    if localVarHTTPResponse.Header.Get("Content-Encoding") == "gzip" {
        respBodyReadCloser, err = gzip.NewReader(localVarHTTPResponse.Body)
    }

	if localVarHTTPResponse.StatusCode >= 300 {
        localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
        localVarHTTPResponse.Body.Close()
        respBodyReadCloser.Close()
        localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
        if err != nil {
            return localVarReturnValue, localVarHTTPResponse, err
        }
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
    localVarHTTPResponse.Body.Close()
    respBodyReadCloser.Close()
    localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSearchRequest struct {
	ctx context.Context
	ApiService DefaultApi
	q *string
}

// SOSL search statement
func (r ApiSearchRequest) Q(q string) ApiSearchRequest {
	r.q = &q
	return r
}

func (r ApiSearchRequest) Execute() (*SOSLResult, *http.Response, error) {
	return r.ApiService.SearchExecute(r)
}

/*
Search Executes SOSL

Search for records containing specified values.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchRequest
*/
func (a *DefaultApiService) Search(ctx context.Context) ApiSearchRequest {
	return ApiSearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SOSLResult
func (a *DefaultApiService) SearchExecute(r ApiSearchRequest) (*SOSLResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SOSLResult
	)

    token, err := a.client.TokenSrc.Token()
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    activeServer := a.client.cfg.GetActiveServer()
    activeServer.SetServerVariable("instanceUrl", token.Extra("instance_url").(string))

	localVarPath := activeServer.GetURL() + "/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.q == nil {
		return localVarReturnValue, nil, reportError("q is required and must be specified")
	}

	localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

    token.SetAuthHeader(req)

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}


    respBodyReadCloser := localVarHTTPResponse.Body
    if localVarHTTPResponse.Header.Get("Content-Encoding") == "gzip" {
        respBodyReadCloser, err = gzip.NewReader(localVarHTTPResponse.Body)
    }

	if localVarHTTPResponse.StatusCode >= 300 {
        localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
        localVarHTTPResponse.Body.Close()
        respBodyReadCloser.Close()
        localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
        if err != nil {
            return localVarReturnValue, localVarHTTPResponse, err
        }
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    localVarBody, err := ioutil.ReadAll(respBodyReadCloser)
    localVarHTTPResponse.Body.Close()
    respBodyReadCloser.Close()
    localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

    return localVarReturnValue, localVarHTTPResponse, nil
}
