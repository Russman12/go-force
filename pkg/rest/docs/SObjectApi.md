# SObjectApi

All URIs are relative to *https://myorg.lightning.force.com/services/data/v56.0*

| Method        | HTTP request  | Summary       |
| ------------- | ------------- | ------------- |
| [**CreateRecord**](#createrecord) | **Post** /sobjects/{sObject} | Create Records Using sObject Basic Information |
| [**GetBasicInfo**](#getbasicinfo) | **Get** /sobjects/{sObject} | Retrieve Object Metadata Using sObject Basic Information |
| [**GetSObjects**](#getsobjects) | **Get** /sobjects | Get a List of Objects |
| [**RetrieveRecord**](#retrieverecord) | **Get** /sobjects/{sObject}/{recordId} | Retrieve Records Using sObject Rows |
| [**SObjectDescribe**](#sobjectdescribe) | **Get** /sobjects/{sObject}/describe | Retrieve Metadata for an Object |



## CreateRecord

> CreateRecordResult CreateRecord(ctx, sObject).Body(body).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()

Creates a new record for a specified object based on field values in the request body. You must specify values for required fields in the request body. Specifying values for other fields is optional.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_sobject_basic_info_post.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/rest"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := rest.NewConfiguration()
    apiClient := rest.NewAPIClient(configuration, tokenSrc)

    sObject := "Contact" // string | The name of the object.
    body := map[string]interface{}{} // map[string]interface{} | SObject record to insert
    contentEncoding := EncodingType{} // EncodingType |  (optional)
    acceptEncoding := EncodingType{} // EncodingType |  (optional)

    resp, r, err := apiClient.SObjectApi.CreateRecord(context.Background(), sObject).Body(body).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()
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
| **ctx** | **context.Context** | context for logging, cancellation, deadlines, tracing, etc. | |
| **sObject** | **string** | The name of the object. |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecordRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **body** | **map[string]interface{}** | SObject record to insert | 
|  **contentEncoding** | [**EncodingType**](EncodingType.md) |  | 
|  **acceptEncoding** | [**EncodingType**](EncodingType.md) |  |  |

### Return type

[**CreateRecordResult**](CreateRecordResult.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBasicInfo

> map[string]interface{} GetBasicInfo(ctx, sObject).AcceptEncoding(acceptEncoding).Execute()

Retrieves basic metadata for a specified object, including some object properties, recent items, and URIs for other resources related to the object. To retrieve the complete metadata for an object, use the (sObject Describe)[https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_sobject_describe.htm] resource.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_sobject_basic_info_get.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/rest"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := rest.NewConfiguration()
    apiClient := rest.NewAPIClient(configuration, tokenSrc)

    sObject := "Contact" // string | The name of the object.
    acceptEncoding := EncodingType{} // EncodingType |  (optional)

    resp, r, err := apiClient.SObjectApi.GetBasicInfo(context.Background(), sObject).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SObjectApi.GetBasicInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBasicInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SObjectApi.GetBasicInfo`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type          | Description   | Notes        |
| ------------ | ------------- | ------------- | ------------ |
| **ctx** | **context.Context** | context for logging, cancellation, deadlines, tracing, etc. | |
| **sObject** | **string** | The name of the object. |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetBasicInfoRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **acceptEncoding** | [**EncodingType**](EncodingType.md) |  |  |

### Return type

**map[string]interface{}**

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSObjects

> SObjectDescribes GetSObjects(ctx).AcceptEncoding(acceptEncoding).Execute()

Lists the available objects and their metadata for your organization???s data. In addition, it provides the organization encoding, as well as the maximum batch size permitted in queries. For more information on encoding, see [Internationalization and Character Sets](https://developer.salesforce.com/docs/atlas.en-us.242.0.api.meta/api/implementation_considerations.htm#sforce_api_other_internationalization).
You can use the If-Modified-Since or If-Unmodified-Since header with this resource. When using the If-Modified-Since header, if no available object???s metadata has changed since the provided date, a 304 Not Modified status code is returned with no response body.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_describeGlobal.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/rest"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := rest.NewConfiguration()
    apiClient := rest.NewAPIClient(configuration, tokenSrc)

    acceptEncoding := EncodingType{} // EncodingType |  (optional)

    resp, r, err := apiClient.SObjectApi.GetSObjects(context.Background()).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SObjectApi.GetSObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSObjects`: SObjectDescribes
    fmt.Fprintf(os.Stdout, "Response from `SObjectApi.GetSObjects`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSObjectsRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  **acceptEncoding** | [**EncodingType**](EncodingType.md) |  |  |

### Return type

[**SObjectDescribes**](SObjectDescribes.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRecord

> map[string]interface{} RetrieveRecord(ctx, sObject, recordId).AcceptEncoding(acceptEncoding).Fields(fields).Execute()

Retrieves a record based on the specified object and record ID. The fields and field values of the record are returned in the response body.


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/rest"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := rest.NewConfiguration()
    apiClient := rest.NewAPIClient(configuration, tokenSrc)

    sObject := "Contact" // string | The name of the object.
    recordId := "001R0000005hDFYIA2" // string | The identifier of the object.
    acceptEncoding := EncodingType{} // EncodingType |  (optional)
    fields := []string{} // []string | A comma-delimited list of fields specifying the fields and values returned in the response body. (optional)

    resp, r, err := apiClient.SObjectApi.RetrieveRecord(context.Background(), sObject, recordId).AcceptEncoding(acceptEncoding).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SObjectApi.RetrieveRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRecord`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SObjectApi.RetrieveRecord`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type          | Description   | Notes        |
| ------------ | ------------- | ------------- | ------------ |
| **ctx** | **context.Context** | context for logging, cancellation, deadlines, tracing, etc. | |
| **sObject** | **string** | The name of the object. |  | |
| **recordId** | **string** | The identifier of the object. |  |

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRecordRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
| 
| 
|  **acceptEncoding** | [**EncodingType**](EncodingType.md) |  | 
|  **fields** | **[]string** | A comma-delimited list of fields specifying the fields and values returned in the response body. |  |

### Return type

**map[string]interface{}**

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SObjectDescribe

> SObjectDescribe SObjectDescribe(ctx, sObject).IfModifiedSince(ifModifiedSince).IfUnmodifiedSince(ifUnmodifiedSince).AcceptEncoding(acceptEncoding).Execute()

Completely describes the individual metadata at all levels for the specified object. For example, this can be used to retrieve the fields, URLs, and child relationships for the Account object.
For more information about the metadata that is retrieved, see [DescribesObjectResult](https://developer.salesforce.com/docs/atlas.en-us.242.0.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm) in the SOAP API Developers Guide.
You can use the If-Modified-Since or If-Unmodified-Since header with this resource. When using the If-Modified-Since header, if no available object???s metadata has changed since the provided date, a 304 Not Modified status code is returned with no response body.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_sobject_describe.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/rest"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := rest.NewConfiguration()
    apiClient := rest.NewAPIClient(configuration, tokenSrc)

    sObject := "Contact" // string | The name of the object.
    ifModifiedSince := "Mon, 30 Nov 2020 08:34:54 MST" // string | An optional header specifying a date and time. The request returns records that have been modified after that date and time. (optional)
    ifUnmodifiedSince := "Mon, 30 Nov 2020 08:34:54 MST" // string | An optional header specifying a date and time. The request returns records that have not been modified after that date and time. (optional)
    acceptEncoding := EncodingType{} // EncodingType |  (optional)

    resp, r, err := apiClient.SObjectApi.SObjectDescribe(context.Background(), sObject).IfModifiedSince(ifModifiedSince).IfUnmodifiedSince(ifUnmodifiedSince).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SObjectApi.SObjectDescribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SObjectDescribe`: SObjectDescribe
    fmt.Fprintf(os.Stdout, "Response from `SObjectApi.SObjectDescribe`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type          | Description   | Notes        |
| ------------ | ------------- | ------------- | ------------ |
| **ctx** | **context.Context** | context for logging, cancellation, deadlines, tracing, etc. | |
| **sObject** | **string** | The name of the object. |  |

### Other Parameters

Other parameters are passed through a pointer to a apiSObjectDescribeRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **ifModifiedSince** | **string** | An optional header specifying a date and time. The request returns records that have been modified after that date and time. | 
|  **ifUnmodifiedSince** | **string** | An optional header specifying a date and time. The request returns records that have not been modified after that date and time. | 
|  **acceptEncoding** | [**EncodingType**](EncodingType.md) |  |  |

### Return type

[**SObjectDescribe**](SObjectDescribe.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

