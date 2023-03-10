# OrgApi

All URIs are relative to *https://myorg.lightning.force.com/services/data/v56.0*

| Method        | HTTP request  | Summary       |
| ------------- | ------------- | ------------- |
| [**GetLimits**](#getlimits) | **Get** /limits | List Org Limits |
| [**GetResources**](#getresources) | **Get** / | List Available REST Resources |



## GetLimits

> interface{} GetLimits(ctx).AcceptEncoding(acceptEncoding).Execute()

Lists information about limits in your org. For each limit, this resource returns the maximum allocation and the remaining allocation based on usage. This resource is available in REST API version 29.0 and later for API users with the View Setup and Configuration permission.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_limits.htm)

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

    resp, r, err := apiClient.OrgApi.GetLimits(context.Background()).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.GetLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLimits`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrgApi.GetLimits`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiGetLimitsRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  **acceptEncoding** | [**EncodingType**](EncodingType.md) |  |  |

### Return type

**interface{}**

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResources

> interface{} GetResources(ctx).AcceptEncoding(acceptEncoding).Execute()

Lists available resources for the specified API version, including resource name and URI.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_discoveryresource.htm)

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

    resp, r, err := apiClient.OrgApi.GetResources(context.Background()).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.GetResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResources`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrgApi.GetResources`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcesRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  **acceptEncoding** | [**EncodingType**](EncodingType.md) |  |  |

### Return type

**interface{}**

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

