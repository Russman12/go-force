# DefaultApi

All URIs are relative to *https://myorg.lightning.force.com/services/data/v56.0*

| Method        | HTTP request  | Summary       |
| ------------- | ------------- | ------------- |
| [**SObjectDescribe**](#sobjectdescribe) | **Get** /sobjects/{sObject}/describe | sObject Describe |



## SObjectDescribe

> SObjectDescribe SObjectDescribe(ctx, sObject).Execute()

Completely describes the individual metadata at all levels for the specified object. For example, this can be used to retrieve the fields, URLs, and child relationships for the Account object.
For more information about the metadata that is retrieved, see [DescribesObjectResult](https://developer.salesforce.com/docs/atlas.en-us.242.0.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm) in the SOAP API Developers Guide.
You can use the If-Modified-Since or If-Unmodified-Since header with this resource. When using the If-Modified-Since header, if no available object’s metadata has changed since the provided date, a 304 Not Modified status code is returned with no response body.

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

    sObject := rest.Test // string | SObject name

    resp, r, err := apiClient.DefaultApi.SObjectDescribe(context.Background(), sObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SObjectDescribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SObjectDescribe`: SObjectDescribe
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SObjectDescribe`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type          | Description   | Notes        |
| ------------ | ------------- | ------------- | ------------ |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. | |
| **sObject** | **string** | SObject name |  |

### Other Parameters

Other parameters are passed through a pointer to a apiSObjectDescribeRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  |

### Return type

[**SObjectDescribe**](SObjectDescribe.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (implicit)](../README.md#oauth--implicit-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

