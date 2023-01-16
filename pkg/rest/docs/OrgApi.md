# OrgApi

All URIs are relative to *https://myorg.lightning.force.com/services/data/v56.0*

| Method        | HTTP request  | Summary       |
| ------------- | ------------- | ------------- |
| [**GetLimits**](#getlimits) | **Get** /limits | Limits |
| [**GetResources**](#getresources) | **Get** / | Resources by Version |



## GetLimits

> interface{} GetLimits(ctx).Execute()

Lists information about limits in your org. For each limit, this resource returns the maximum allocation and the remaining allocation based on usage. This resource is available in REST API version 29.0 and later for API users with the View Setup and Configuration permission.

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
    resp, r, err := apiClient.OrgApi.GetLimits(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.GetLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLimits`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrgApi.GetLimits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetLimitsRequest struct via the builder pattern
 |

### Return type

**interface{}**

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (implicit)](../README.md#oauth--implicit-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResources

> interface{} GetResources(ctx).Execute()

Lists available resources for the specified API version, including resource name and URI.

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
    resp, r, err := apiClient.OrgApi.GetResources(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.GetResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResources`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrgApi.GetResources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcesRequest struct via the builder pattern
 |

### Return type

**interface{}**

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (implicit)](../README.md#oauth--implicit-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

