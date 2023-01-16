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
)

// OrgApiService OrgApi service
type OrgApiService service

type ApiGetLimitsRequest struct {
	ctx        context.Context
	ApiService *OrgApiService
}

func (r ApiGetLimitsRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.GetLimitsExecute(r)
}

/*
GetLimits Limits

Lists information about limits in your org. For each limit, this resource returns the maximum allocation and the remaining allocation based on usage. This resource is available in REST API version 29.0 and later for API users with the View Setup and Configuration permission.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetLimitsRequest
*/
func (a *OrgApiService) GetLimits(ctx context.Context) ApiGetLimitsRequest {
	return ApiGetLimitsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return interface{}
func (a *OrgApiService) GetLimitsExecute(r ApiGetLimitsRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue interface{}
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

	localVarPath := localBasePath + "/limits"

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

type ApiGetResourcesRequest struct {
	ctx        context.Context
	ApiService *OrgApiService
}

func (r ApiGetResourcesRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.GetResourcesExecute(r)
}

/*
GetResources Resources by Version

Lists available resources for the specified API version, including resource name and URI.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetResourcesRequest
*/
func (a *OrgApiService) GetResources(ctx context.Context) ApiGetResourcesRequest {
	return ApiGetResourcesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return interface{}
func (a *OrgApiService) GetResourcesExecute(r ApiGetResourcesRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue interface{}
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

	localVarPath := localBasePath + "/"

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
