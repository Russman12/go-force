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
	body       *interface{}
}

// SObject record to insert
func (r ApiCreateRecordRequest) Body(body interface{}) ApiCreateRecordRequest {
	r.body = &body
	return r
}

func (r ApiCreateRecordRequest) Execute() (*CreateRecordResult, *http.Response, error) {
	return r.ApiService.CreateRecordExecute(r)
}

/*
CreateRecord Method for CreateRecord

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

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
		localVarHTTPResponse.Body.Close()
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
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
GetSObjects Method for GetSObjects

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

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
		localVarHTTPResponse.Body.Close()
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
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
	ctx        context.Context
	ApiService *SObjectApiService
	sObject    string
}

func (r ApiSObjectDescribeRequest) Execute() (*SObjectDescribe, *http.Response, error) {
	return r.ApiService.SObjectDescribeExecute(r)
}

/*
SObjectDescribe Method for SObjectDescribe

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	token.SetAuthHeader(req)

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
		localVarHTTPResponse.Body.Close()
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
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
