# \SObjectApi

All URIs are relative to *https://test.salesforce.com/services/data/v56.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SObjectDescribe**](SObjectApi.md#SObjectDescribe) | **Get** /sobjects/{sObject}/describe | 
[**SObjects**](SObjectApi.md#SObjects) | **Get** /sobjects | 



## SObjectDescribe

> SObjectDescribe SObjectDescribe(ctx, sObject).Execute()



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
    resp, r, err := apiClient.SObjectApi.SObjectDescribe(context.Background(), sObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SObjectApi.SObjectDescribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SObjectDescribe`: SObjectDescribe
    fmt.Fprintf(os.Stdout, "Response from `SObjectApi.SObjectDescribe`: %v\n", resp)
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

[**SObjectDescribe**](SObjectDescribe.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SObjects

> SObjectDescribes SObjects(ctx).Execute()



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
    resp, r, err := apiClient.SObjectApi.SObjects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SObjectApi.SObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SObjects`: SObjectDescribes
    fmt.Fprintf(os.Stdout, "Response from `SObjectApi.SObjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSObjectsRequest struct via the builder pattern


### Return type

[**SObjectDescribes**](SObjectDescribes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

