# DefaultApi

All URIs are relative to *https://myorg.lightning.force.com/services/data/v56.0/tooling*

| Method        | HTTP request  | Summary       |
| ------------- | ------------- | ------------- |
| [**ExecuteAnonymous**](#executeanonymous) | **Get** /executeAnonymous | Executes anonymous Apex |
| [**GetCompletions**](#getcompletions) | **Get** /completions | Retrieve available code completions |
| [**GetSObjects**](#getsobjects) | **Get** /sobjects | List SObjects |
| [**Query**](#query) | **Get** /query | Executes query |
| [**RunTestsAsync**](#runtestsasync) | **Post** /runTestsAsynchronous | Run tests asynchronously |
| [**RunTestsSync**](#runtestssync) | **Post** /runTestsSynchronous | Run tests synchronously |
| [**Search**](#search) | **Get** /search | Executes SOSL |



## ExecuteAnonymous

> ExecuteAnonymousResult ExecuteAnonymous(ctx).AnonymousBody(anonymousBody).Execute()

Executes Apex code anonymously. Available from API version 29.0 or later.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_tooling.meta/api_tooling/intro_rest_resources.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/tooling"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := tooling.NewConfiguration()
    apiClient := tooling.NewAPIClient(configuration, tokenSrc)

    anonymousBody := "anonymousBody_example" // string | 

    resp, r, err := apiClient.DefaultApi.ExecuteAnonymous(context.Background()).AnonymousBody(anonymousBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExecuteAnonymous``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteAnonymous`: ExecuteAnonymousResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ExecuteAnonymous`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteAnonymousRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  **anonymousBody** | **string** |  |  |

### Return type

[**ExecuteAnonymousResult**](ExecuteAnonymousResult.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompletions

> interface{} GetCompletions(ctx).Type_(type_).Execute()

Retrieves available code completions of the referenced type for Apex system method symbols (type=apex). Available from API version 28.0 or later. Retrieves available code completions of the referenced type for Visualforce markup (type=visualforce). Available from API version 38.0 or later.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_tooling.meta/api_tooling/intro_rest_resources.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/tooling"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := tooling.NewConfiguration()
    apiClient := tooling.NewAPIClient(configuration, tokenSrc)

    type_ := CompletionType{} // CompletionType | 

    resp, r, err := apiClient.DefaultApi.GetCompletions(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCompletions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompletions`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCompletions`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompletionsRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  **type_** | [**CompletionType**](CompletionType.md) |  |  |

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


## GetSObjects

> DescribeGlobalResult GetSObjects(ctx).Execute()

Lists the available Tooling API objects and their metadata.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_tooling.meta/api_tooling/intro_rest_resources.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/tooling"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := tooling.NewConfiguration()
    apiClient := tooling.NewAPIClient(configuration, tokenSrc)


    resp, r, err := apiClient.DefaultApi.GetSObjects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSObjects`: DescribeGlobalResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSObjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSObjectsRequest struct via the builder pattern
 |

### Return type

[**DescribeGlobalResult**](DescribeGlobalResult.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Query

> QueryResult Query(ctx).Q(q).Execute()

Executes a query against an object and returns data that matches the specified criteria. Tooling API exposes objects like EntityDefinition and FieldDefinition that use the external object framework. That is, they don’t exist in the database but are constructed dynamically. Special query rules apply to virtual entities. If the query result is too large, it’s broken up into batches. The response contains the first batch of results and a query identifier. The identifier can be used in a request to retrieve the next batch.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_tooling.meta/api_tooling/intro_rest_resources.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/tooling"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := tooling.NewConfiguration()
    apiClient := tooling.NewAPIClient(configuration, tokenSrc)

    q := "q_example" // string | SOQL query statement

    resp, r, err := apiClient.DefaultApi.Query(context.Background()).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Query``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Query`: QueryResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Query`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiQueryRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  **q** | **string** | SOQL query statement |  |

### Return type

[**QueryResult**](QueryResult.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunTestsAsync

> string RunTestsAsync(ctx).AsyncTestRequest(asyncTestRequest).Execute()

Runs one or more methods within one or more Apex classes, using the asynchronous test execution mechanism. In the request body, you can specify test class names and IDs, suite names and IDs, the maximum number of failed tests to allow, and the test level, as comma-separated lists or as an array. You can also indicate whether to opt out of collecting code coverage information during the test run (available in API version 43.0 and later).

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_tooling.meta/api_tooling/intro_rest_resources.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/tooling"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := tooling.NewConfiguration()
    apiClient := tooling.NewAPIClient(configuration, tokenSrc)

    asyncTestRequest := AsyncTestRequest{} // AsyncTestRequest |  (optional)

    resp, r, err := apiClient.DefaultApi.RunTestsAsync(context.Background()).AsyncTestRequest(asyncTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RunTestsAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunTestsAsync`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RunTestsAsync`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiRunTestsAsyncRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  **asyncTestRequest** | [**AsyncTestRequest**](AsyncTestRequest.md) |  |  |

### Return type

**string**

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunTestsSync

> RunTestsResult RunTestsSync(ctx).TestRequest(testRequest).Execute()

Runs one or more methods within an Apex class, using the synchronous test execution mechanism. All test methods in a synchronous test run must be in the same class.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_tooling.meta/api_tooling/intro_rest_resources.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/tooling"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := tooling.NewConfiguration()
    apiClient := tooling.NewAPIClient(configuration, tokenSrc)

    testRequest := TestRequest{} // TestRequest |  (optional)

    resp, r, err := apiClient.DefaultApi.RunTestsSync(context.Background()).TestRequest(testRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RunTestsSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunTestsSync`: RunTestsResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RunTestsSync`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiRunTestsSyncRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  **testRequest** | [**TestRequest**](TestRequest.md) |  |  |

### Return type

[**RunTestsResult**](RunTestsResult.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> SOSLResult Search(ctx).Q(q).Execute()

Search for records containing specified values.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_tooling.meta/api_tooling/intro_rest_resources.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/tooling"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := tooling.NewConfiguration()
    apiClient := tooling.NewAPIClient(configuration, tokenSrc)

    q := "q_example" // string | SOSL search statement

    resp, r, err := apiClient.DefaultApi.Search(context.Background()).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: SOSLResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Search`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  **q** | **string** | SOSL search statement |  |

### Return type

[**SOSLResult**](SOSLResult.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

