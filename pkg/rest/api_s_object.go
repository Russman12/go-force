/*
Salesforce Platform REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 56.0
Contact: russell-laboe@outlook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"bytes"
	"compress/gzip"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// SObjectApiService SObjectApi service
type SObjectApiService service

type ApiCreateRecordRequest struct {
	ctx        context.Context
	ApiService *SObjectApiService
	sObject    string
	body       *map[string]interface{}
}

// SObject record to insert
func (r ApiCreateRecordRequest) Body(body map[string]interface{}) ApiCreateRecordRequest {
	r.body = &body
	return r
}

func (r ApiCreateRecordRequest) Execute() (*CreateRecordResult, *http.Response, error) {
	return r.ApiService.CreateRecordExecute(r)
}

/*
CreateRecord Create Records Using sObject Basic Information

Creates a new record for a specified object based on field values in the request body. You must specify values for required fields in the request body. Specifying values for other fields is optional.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sObject SObject name
	@return ApiCreateRecordRequest
*/
func (a *SObjectApiService) CreateRecord(ctx context.Context, sObject string) ApiCreateRecordRequest {
	return ApiCreateRecordRequest{
		ApiService: a,
		ctx:        ctx,
		sObject:    sObject,
	}
}

// Execute executes the request
//
//	@return CreateRecordResult
func (a *SObjectApiService) CreateRecordExecute(r ApiCreateRecordRequest) (*CreateRecordResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateRecordResult
	)

	token, err := a.client.tokenSrc.Token()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	serverIdx, err := getServerIndex(r.ctx)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localBasePath, err := a.client.cfg.ServerURL(serverIdx, map[string]string{"instanceUrl": token.Extra("instance_url").(string)})
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sobjects/{sObject}"
	localVarPath = strings.Replace(localVarPath, "{"+"sObject"+"}", url.PathEscape(parameterToString(r.sObject, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

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
	localVarPostBody = r.body
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v []ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			// marshaling never returns err since v can only contain openapi valid types
			newErr.error, _ = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v []ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			// marshaling never returns err since v can only contain openapi valid types
			newErr.error, _ = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiGetBasicInfoRequest struct {
	ctx        context.Context
	ApiService *SObjectApiService
	sObject    string
}

func (r ApiGetBasicInfoRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetBasicInfoExecute(r)
}

/*
GetBasicInfo Retrieve Object Metadata Using sObject Basic Information

Retrieves basic metadata for a specified object, including some object properties, recent items, and URIs for other resources related to the object. To retrieve the complete metadata for an object, use the (sObject Describe)[https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_sobject_describe.htm] resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sObject SObject name
	@return ApiGetBasicInfoRequest
*/
func (a *SObjectApiService) GetBasicInfo(ctx context.Context, sObject string) ApiGetBasicInfoRequest {
	return ApiGetBasicInfoRequest{
		ApiService: a,
		ctx:        ctx,
		sObject:    sObject,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *SObjectApiService) GetBasicInfoExecute(r ApiGetBasicInfoRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	token, err := a.client.tokenSrc.Token()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	serverIdx, err := getServerIndex(r.ctx)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localBasePath, err := a.client.cfg.ServerURL(serverIdx, map[string]string{"instanceUrl": token.Extra("instance_url").(string)})
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sobjects/{sObject}"
	localVarPath = strings.Replace(localVarPath, "{"+"sObject"+"}", url.PathEscape(parameterToString(r.sObject, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v []ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			// marshaling never returns err since v can only contain openapi valid types
			newErr.error, _ = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v []ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			// marshaling never returns err since v can only contain openapi valid types
			newErr.error, _ = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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
	ctx        context.Context
	ApiService *SObjectApiService
}

func (r ApiGetSObjectsRequest) Execute() (*SObjectDescribes, *http.Response, error) {
	return r.ApiService.GetSObjectsExecute(r)
}

/*
GetSObjects Get a List of Objects

Lists the available objects and their metadata for your organization’s data. In addition, it provides the organization encoding, as well as the maximum batch size permitted in queries. For more information on encoding, see [Internationalization and Character Sets](https://developer.salesforce.com/docs/atlas.en-us.242.0.api.meta/api/implementation_considerations.htm#sforce_api_other_internationalization).
You can use the If-Modified-Since or If-Unmodified-Since header with this resource. When using the If-Modified-Since header, if no available object’s metadata has changed since the provided date, a 304 Not Modified status code is returned with no response body.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetSObjectsRequest
*/
func (a *SObjectApiService) GetSObjects(ctx context.Context) ApiGetSObjectsRequest {
	return ApiGetSObjectsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SObjectDescribes
func (a *SObjectApiService) GetSObjectsExecute(r ApiGetSObjectsRequest) (*SObjectDescribes, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SObjectDescribes
	)

	token, err := a.client.tokenSrc.Token()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	serverIdx, err := getServerIndex(r.ctx)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localBasePath, err := a.client.cfg.ServerURL(serverIdx, map[string]string{"instanceUrl": token.Extra("instance_url").(string)})
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sobjects"

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

type ApiSObjectDescribeRequest struct {
	ctx               context.Context
	ApiService        *SObjectApiService
	sObject           string
	ifModifiedSince   *string
	ifUnmodifiedSince *string
}

// An optional header specifying a date and time. The request returns records that have been modified after that date and time.
func (r ApiSObjectDescribeRequest) IfModifiedSince(ifModifiedSince string) ApiSObjectDescribeRequest {
	r.ifModifiedSince = &ifModifiedSince
	return r
}

// An optional header specifying a date and time. The request returns records that have not been modified after that date and time.
func (r ApiSObjectDescribeRequest) IfUnmodifiedSince(ifUnmodifiedSince string) ApiSObjectDescribeRequest {
	r.ifUnmodifiedSince = &ifUnmodifiedSince
	return r
}

func (r ApiSObjectDescribeRequest) Execute() (*SObjectDescribe, *http.Response, error) {
	return r.ApiService.SObjectDescribeExecute(r)
}

/*
SObjectDescribe Retrieve Metadata for an Object

Completely describes the individual metadata at all levels for the specified object. For example, this can be used to retrieve the fields, URLs, and child relationships for the Account object.
For more information about the metadata that is retrieved, see [DescribesObjectResult](https://developer.salesforce.com/docs/atlas.en-us.242.0.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm) in the SOAP API Developers Guide.
You can use the If-Modified-Since or If-Unmodified-Since header with this resource. When using the If-Modified-Since header, if no available object’s metadata has changed since the provided date, a 304 Not Modified status code is returned with no response body.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sObject SObject name
	@return ApiSObjectDescribeRequest
*/
func (a *SObjectApiService) SObjectDescribe(ctx context.Context, sObject string) ApiSObjectDescribeRequest {
	return ApiSObjectDescribeRequest{
		ApiService: a,
		ctx:        ctx,
		sObject:    sObject,
	}
}

// Execute executes the request
//
//	@return SObjectDescribe
func (a *SObjectApiService) SObjectDescribeExecute(r ApiSObjectDescribeRequest) (*SObjectDescribe, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SObjectDescribe
	)

	token, err := a.client.tokenSrc.Token()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	serverIdx, err := getServerIndex(r.ctx)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localBasePath, err := a.client.cfg.ServerURL(serverIdx, map[string]string{"instanceUrl": token.Extra("instance_url").(string)})
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sobjects/{sObject}/describe"
	localVarPath = strings.Replace(localVarPath, "{"+"sObject"+"}", url.PathEscape(parameterToString(r.sObject, "")), -1)

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
	if r.ifModifiedSince != nil {
		localVarHeaderParams["If-Modified-Since"] = parameterToString(*r.ifModifiedSince, "")
	}
	if r.ifUnmodifiedSince != nil {
		localVarHeaderParams["If-Unmodified-Since"] = parameterToString(*r.ifUnmodifiedSince, "")
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
