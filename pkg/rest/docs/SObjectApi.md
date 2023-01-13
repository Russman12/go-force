# \SobjectApi

All URIs are relative to *https://test.salesforce.com/services/data/v56.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSObjects**](SobjectApi.md#GetSObjects) | **Get** /sobjects | 
[**SObjectDescribe**](SobjectApi.md#SObjectDescribe) | **Get** /sobjects/{sObject}/describe | 



## GetSObjects

> SobjectDescribes GetSObjects(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SobjectApi.GetSObjects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SobjectApi.GetSObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSObjects`: SobjectDescribes
    fmt.Fprintf(os.Stdout, "Response from `SobjectApi.GetSObjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSObjectsRequest struct via the builder pattern


### Return type

[**SobjectDescribes**](SobjectDescribes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SObjectDescribe

> SobjectDescribe SObjectDescribe(ctx, sObject).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sObject := "Contact" // string | SObject name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SobjectApi.SObjectDescribe(context.Background(), sObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SobjectApi.SObjectDescribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SObjectDescribe`: SobjectDescribe
    fmt.Fprintf(os.Stdout, "Response from `SobjectApi.SObjectDescribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sObject** | **string** | SObject name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSObjectDescribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SobjectDescribe**](SobjectDescribe.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

