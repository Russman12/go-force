# SObjectApi

All URIs are relative to *https://myorg.lightning.force.com/services/data/v56.0*

| Method        | HTTP request  | Summary       |
| ------------- | ------------- | ------------- |
| [**CreateRecord**](#createrecord) | **Post** /sobjects/{sObject} | Create a Record |
| [**GetSObjects**](#getsobjects) | **Get** /sobjects | Describe Global |



## CreateRecord

> CreateRecordResult CreateRecord(ctx, sObject).Body(body).Execute()

Use the [sObject Basic Information](https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_sobject_basic_info.htm) resource to create new records. You supply the required field values in the request data, and send the request using the POST HTTP method. The response body contains the ID of the new record if the call is successful.

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
    body := map[string]interface{}{ ... } // map[string]interface{} | SObject record to insert (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SObjectApi.CreateRecord(context.Background(), sObject).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SObjectApi.CreateRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRecord`: CreateRecordResult
    fmt.Fprintf(os.Stdout, "Response from `SObjectApi.CreateRecord`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type          | Description   | Notes        |
| ------------ | ------------- | ------------- | ------------ |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. | |
| **sObject** | **string** | SObject name |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecordRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **body** | **map[string]interface{}** | SObject record to insert |  |

### Return type

[**CreateRecordResult**](CreateRecordResult.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (implicit)](../README.md#oauth--implicit-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSObjects

> SObjectDescribes GetSObjects(ctx).Execute()

Lists the available objects and their metadata for your organization’s data. In addition, it provides the organization encoding, as well as the maximum batch size permitted in queries. For more information on encoding, see [Internationalization and Character Sets](https://developer.salesforce.com/docs/atlas.en-us.242.0.api.meta/api/implementation_considerations.htm#sforce_api_other_internationalization).
You can use the If-Modified-Since or If-Unmodified-Since header with this resource. When using the If-Modified-Since header, if no available object’s metadata has changed since the provided date, a 304 Not Modified status code is returned with no response body.

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
    resp, r, err := apiClient.SObjectApi.GetSObjects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SObjectApi.GetSObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSObjects`: SObjectDescribes
    fmt.Fprintf(os.Stdout, "Response from `SObjectApi.GetSObjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSObjectsRequest struct via the builder pattern
 |

### Return type

[**SObjectDescribes**](SObjectDescribes.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (implicit)](../README.md#oauth--implicit-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

