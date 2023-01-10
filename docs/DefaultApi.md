# \DefaultApi

All URIs are relative to *https://test.salesforce.com/services/data/v56.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SObjects**](DefaultApi.md#SObjects) | **Get** /sobjects | 



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
    resp, r, err := apiClient.DefaultApi.SObjects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SObjects`: SObjectDescribes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SObjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSObjectsRequest struct via the builder pattern


### Return type

[**SObjectDescribes**](SObjectDescribes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

