# QueryApi

All URIs are relative to *https://test.salesforce.com/services/data/v56.0*

| Method        | HTTP request  | Summary       |
| ------------- | ------------- | ------------- |
| [**AbortQueryJob**](#abortqueryjob) | **Patch** /jobs/query/{jobId} | Abort a Query Job |
| [**CreateQueryJob**](#createqueryjob) | **Post** /jobs/query | Create a Query Job |
| [**DeleteQueryJob**](#deletequeryjob) | **Delete** /jobs/query/{jobId} | Delete a Query Job |
| [**GetJobResults**](#getjobresults) | **Get** /jobs/query/{jobId}/results | Get Results for a Query Job |
| [**GetQueryJobInfo**](#getqueryjobinfo) | **Get** /jobs/query/{jobId} | Get Information About a Query Job |
| [**GetQueryJobs**](#getqueryjobs) | **Get** /jobs/query | Get Information About All Query Jobs |



## AbortQueryJob

> QueryJobInfo AbortQueryJob(ctx, jobId).AbortQueryJobRequest(abortQueryJobRequest).Execute()

Aborts a query job.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/query_abort_job.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/bulkv2"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := bulkv2.NewConfiguration()
    apiClient := bulkv2.NewAPIClient(configuration, tokenSrc)

    jobId := "jobId_example" // string | 
    abortQueryJobRequest := AbortQueryJobRequest{} // AbortQueryJobRequest | 

    resp, r, err := apiClient.QueryApi.AbortQueryJob(context.Background(), jobId).AbortQueryJobRequest(abortQueryJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.AbortQueryJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AbortQueryJob`: QueryJobInfo
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.AbortQueryJob`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type          | Description   | Notes        |
| ------------ | ------------- | ------------- | ------------ |
| **ctx** | **context.Context** | context for logging, cancellation, deadlines, tracing, etc. | |
| **jobId** | **string** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAbortQueryJobRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **abortQueryJobRequest** | [**AbortQueryJobRequest**](AbortQueryJobRequest.md) |  |  |

### Return type

[**QueryJobInfo**](QueryJobInfo.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQueryJob

> QueryJobInfo CreateQueryJob(ctx).CreateQueryJobRequest(createQueryJobRequest).SforceCallOptions(sforceCallOptions).Execute()

Creates a query job.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/query_create_job.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/bulkv2"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := bulkv2.NewConfiguration()
    apiClient := bulkv2.NewAPIClient(configuration, tokenSrc)

    createQueryJobRequest := CreateQueryJobRequest{} // CreateQueryJobRequest | 
    sforceCallOptions := "client=caseSensitiveToken; defaultNamespace=battle" // string |  (optional)

    resp, r, err := apiClient.QueryApi.CreateQueryJob(context.Background()).CreateQueryJobRequest(createQueryJobRequest).SforceCallOptions(sforceCallOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.CreateQueryJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQueryJob`: QueryJobInfo
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.CreateQueryJob`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateQueryJobRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  **createQueryJobRequest** | [**CreateQueryJobRequest**](CreateQueryJobRequest.md) |  | 
|  **sforceCallOptions** | **string** |  |  |

### Return type

[**QueryJobInfo**](QueryJobInfo.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQueryJob

> DeleteQueryJob(ctx, jobId).Execute()

Deletes a query job. When a job is deleted, job data stored by Salesforce is deleted and job metadata information is removed. The job no longer displays in the Bulk Data Load Jobs page in Salesforce.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/query_delete_job.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/bulkv2"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := bulkv2.NewConfiguration()
    apiClient := bulkv2.NewAPIClient(configuration, tokenSrc)

    jobId := "jobId_example" // string | 

    resp, r, err := apiClient.QueryApi.DeleteQueryJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.DeleteQueryJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


| Name         | Type          | Description   | Notes        |
| ------------ | ------------- | ------------- | ------------ |
| **ctx** | **context.Context** | context for logging, cancellation, deadlines, tracing, etc. | |
| **jobId** | **string** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueryJobRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  |

### Return type

 (empty response body)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobResults

> io.ReadCloser GetJobResults(ctx, jobId).AcceptEncoding(acceptEncoding).Locator(locator).MaxRecords(maxRecords).Execute()

Gets the results for a query job. The job must have the state `JobComplete`.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/query_get_job_results.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/bulkv2"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := bulkv2.NewConfiguration()
    apiClient := bulkv2.NewAPIClient(configuration, tokenSrc)

    jobId := "jobId_example" // string | 
    acceptEncoding := EncodingType{} // EncodingType |  (optional)
    locator := "locator_example" // string | A string that identifies a specific set of query results. Providing a value for this parameter returns only that set of results. Omitting this parameter returns the first set of results.  You can find the locator string for the next set of results in the response of each request. See Example and Rules and Guidelines.  As long as the associated job exists, the locator string for a set of results does not change. You can use the locator to retrieve a set of results multiple times.  (optional)
    maxRecords := int32{} // int32 | The maximum number of records to retrieve per set of results for the query. The request is still subject to the size limits. If you are working with a very large number of query results, you may experience a timeout before receiving all the data from Salesforce. To prevent a timeout, specify the maximum number of records your client is expecting to receive in the maxRecords parameter. This splits the results into smaller sets with this value as the maximum size.  If you don’t provide a value for this parameter, the server uses a default value based on the service.  (optional)

    resp, r, err := apiClient.QueryApi.GetJobResults(context.Background(), jobId).AcceptEncoding(acceptEncoding).Locator(locator).MaxRecords(maxRecords).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.GetJobResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobResults`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.GetJobResults`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type          | Description   | Notes        |
| ------------ | ------------- | ------------- | ------------ |
| **ctx** | **context.Context** | context for logging, cancellation, deadlines, tracing, etc. | |
| **jobId** | **string** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobResultsRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **acceptEncoding** | [**EncodingType**](EncodingType.md) |  | 
|  **locator** | **string** | A string that identifies a specific set of query results. Providing a value for this parameter returns only that set of results. Omitting this parameter returns the first set of results.  You can find the locator string for the next set of results in the response of each request. See Example and Rules and Guidelines.  As long as the associated job exists, the locator string for a set of results does not change. You can use the locator to retrieve a set of results multiple times.  | 
|  **maxRecords** | **int32** | The maximum number of records to retrieve per set of results for the query. The request is still subject to the size limits. If you are working with a very large number of query results, you may experience a timeout before receiving all the data from Salesforce. To prevent a timeout, specify the maximum number of records your client is expecting to receive in the maxRecords parameter. This splits the results into smaller sets with this value as the maximum size.  If you don’t provide a value for this parameter, the server uses a default value based on the service.  |  |

### Return type

[**io.ReadCloser**](io.ReadCloser.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueryJobInfo

> QueryJobInfo GetQueryJobInfo(ctx, jobId).Execute()

Gets information about one query job.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/query_get_one_job.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/bulkv2"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := bulkv2.NewConfiguration()
    apiClient := bulkv2.NewAPIClient(configuration, tokenSrc)

    jobId := "jobId_example" // string | 

    resp, r, err := apiClient.QueryApi.GetQueryJobInfo(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.GetQueryJobInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueryJobInfo`: QueryJobInfo
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.GetQueryJobInfo`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type          | Description   | Notes        |
| ------------ | ------------- | ------------- | ------------ |
| **ctx** | **context.Context** | context for logging, cancellation, deadlines, tracing, etc. | |
| **jobId** | **string** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryJobInfoRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  |

### Return type

[**QueryJobInfo**](QueryJobInfo.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueryJobs

> QueryJobInfos GetQueryJobs(ctx).IsPkChunkingEnabled(isPkChunkingEnabled).JobType(jobType).ConcurrencyMode(concurrencyMode).QueryLocator(queryLocator).Execute()

Gets information about all query jobs in the org. The information includes Bulk API 2.0 query jobs and all Bulk API jobs.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/query_get_all_jobs.htm)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "golang.org/x/oauth2"
    "github.com/russman12/go-force/pkg/bulkv2"
)

func main() {
    // auth against salesforce
    oAuthCfg := oauth2.Config{}
    token, err := oAuthCfg.PasswordCredentialsToken(context.Background(), "username", "password")
    if err != nil {
        panic(err)
    }
    tokenSrc := oAuthCfg.TokenSource(context.Background(), token)

    configuration := bulkv2.NewConfiguration()
    apiClient := bulkv2.NewAPIClient(configuration, tokenSrc)

    isPkChunkingEnabled := bool{} // bool | If set to true, the request only returns information about jobs where PK Chunking is enabled. (optional)
    jobType := "jobType_example" // string | Gets information only about jobs matching the specified job type. (optional)
    concurrencyMode := "concurrencyMode_example" // string | For future use. Gets information only about jobs matching the specified concurrency mode. Possible values are serial and parallel. (optional)
    queryLocator := "queryLocator_example" // string | use the value from the nextRecordsUrl from the previous set (optional)

    resp, r, err := apiClient.QueryApi.GetQueryJobs(context.Background()).IsPkChunkingEnabled(isPkChunkingEnabled).JobType(jobType).ConcurrencyMode(concurrencyMode).QueryLocator(queryLocator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.GetQueryJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueryJobs`: QueryJobInfos
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.GetQueryJobs`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryJobsRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  **isPkChunkingEnabled** | **bool** | If set to true, the request only returns information about jobs where PK Chunking is enabled. | 
|  **jobType** | **string** | Gets information only about jobs matching the specified job type. | 
|  **concurrencyMode** | **string** | For future use. Gets information only about jobs matching the specified concurrency mode. Possible values are serial and parallel. | 
|  **queryLocator** | **string** | use the value from the nextRecordsUrl from the previous set |  |

### Return type

[**QueryJobInfos**](QueryJobInfos.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

