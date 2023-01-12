# \QueryApi

All URIs are relative to *https://test.salesforce.com/services/data/v56.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryAbortJob**](QueryApi.md#QueryAbortJob) | **Patch** /jobs/query/{jobId} | Abort a Job Query
[**QueryCreateJob**](QueryApi.md#QueryCreateJob) | **Post** /jobs/query | Create job Query
[**QueryDeleteJob**](QueryApi.md#QueryDeleteJob) | **Delete** /jobs/query/{jobId} | Delete Job Query
[**QueryJobInfo**](QueryApi.md#QueryJobInfo) | **Get** /jobs/query/{jobId} | Get Job Info Query
[**QueryJobResults**](QueryApi.md#QueryJobResults) | **Get** /jobs/query/{jobId}/results | Get Job Query Result
[**QueryJobs**](QueryApi.md#QueryJobs) | **Get** /jobs/query | Get All Query Jobs



## QueryAbortJob

> QueryJobInfo QueryAbortJob(ctx, jobId).QueryJobInfo(queryJobInfo).Execute()

Abort a Job Query

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
    jobId := "jobId_example" // string | 
    queryJobInfo := *openapiclient.NewQueryJobInfo("Id_example", "Object_example", "CreatedById_example", "CreatedDate_example", "SystemModstamp_example", "State_example", "ConcurrencyMode_example", float32(123), "Operation_example", "ContentType_example") // QueryJobInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryApi.QueryAbortJob(context.Background(), jobId).QueryJobInfo(queryJobInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.QueryAbortJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryAbortJob`: QueryJobInfo
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.QueryAbortJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryAbortJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryJobInfo** | [**QueryJobInfo**](QueryJobInfo.md) |  | 

### Return type

[**QueryJobInfo**](QueryJobInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryCreateJob

> QueryJobInfo QueryCreateJob(ctx).ContentType(contentType).QueryJobCreateRequest(queryJobCreateRequest).Execute()

Create job Query

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
    contentType := "application/json" // string |  (optional)
    queryJobCreateRequest := *openapiclient.NewQueryJobCreateRequest("Operation_example", "Query_example") // QueryJobCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryApi.QueryCreateJob(context.Background()).ContentType(contentType).QueryJobCreateRequest(queryJobCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.QueryCreateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryCreateJob`: QueryJobInfo
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.QueryCreateJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryCreateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **queryJobCreateRequest** | [**QueryJobCreateRequest**](QueryJobCreateRequest.md) |  | 

### Return type

[**QueryJobInfo**](QueryJobInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryDeleteJob

> QueryDeleteJob(ctx, jobId).Execute()

Delete Job Query

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
    jobId := "jobId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryApi.QueryDeleteJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.QueryDeleteJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryDeleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryJobInfo

> QueryJobInfo QueryJobInfo(ctx, jobId).Execute()

Get Job Info Query

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
    jobId := "jobId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryApi.QueryJobInfo(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.QueryJobInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryJobInfo`: QueryJobInfo
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.QueryJobInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryJobInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryJobInfo**](QueryJobInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryJobResults

> io.ReadCloser QueryJobResults(ctx, jobId).ContentType(contentType).Accept(accept).Locator(locator).MaxRecords(maxRecords).Execute()

Get Job Query Result

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
    jobId := "jobId_example" // string | 
    contentType := "application/json" // string |  (optional)
    accept := "text/csv" // string |  (optional)
    locator := "locator_example" // string | A string that identifies a specific set of query results. Providing a value for this parameter returns only that set of results. Omitting this parameter returns the first set of results.  You can find the locator string for the next set of results in the response of each request. See Example and Rules and Guidelines.  As long as the associated job exists, the locator string for a set of results does not change. You can use the locator to retrieve a set of results multiple times.  (optional)
    maxRecords := int32(56) // int32 | The maximum number of records to retrieve per set of results for the query. The request is still subject to the size limits. If you are working with a very large number of query results, you may experience a timeout before receiving all the data from Salesforce. To prevent a timeout, specify the maximum number of records your client is expecting to receive in the maxRecords parameter. This splits the results into smaller sets with this value as the maximum size.  If you don’t provide a value for this parameter, the server uses a default value based on the service.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryApi.QueryJobResults(context.Background(), jobId).ContentType(contentType).Accept(accept).Locator(locator).MaxRecords(maxRecords).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.QueryJobResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryJobResults`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.QueryJobResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryJobResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **accept** | **string** |  | 
 **locator** | **string** | A string that identifies a specific set of query results. Providing a value for this parameter returns only that set of results. Omitting this parameter returns the first set of results.  You can find the locator string for the next set of results in the response of each request. See Example and Rules and Guidelines.  As long as the associated job exists, the locator string for a set of results does not change. You can use the locator to retrieve a set of results multiple times.  | 
 **maxRecords** | **int32** | The maximum number of records to retrieve per set of results for the query. The request is still subject to the size limits. If you are working with a very large number of query results, you may experience a timeout before receiving all the data from Salesforce. To prevent a timeout, specify the maximum number of records your client is expecting to receive in the maxRecords parameter. This splits the results into smaller sets with this value as the maximum size.  If you don’t provide a value for this parameter, the server uses a default value based on the service.  | 

### Return type

[**io.ReadCloser**](io.ReadCloser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryJobs

> QueryJobsInfos QueryJobs(ctx).IsPkChunkingEnabled(isPkChunkingEnabled).JobType(jobType).ConcurrencyMode(concurrencyMode).QueryLocator(queryLocator).Execute()

Get All Query Jobs

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
    isPkChunkingEnabled := false // bool | If set to true, the request only returns information about jobs where PK Chunking is enabled. (optional)
    jobType := "jobType_example" // string | Gets information only about jobs matching the specified job type. (optional)
    concurrencyMode := "concurrencyMode_example" // string | For future use. Gets information only about jobs matching the specified concurrency mode. Possible values are serial and parallel. (optional)
    queryLocator := "queryLocator_example" // string | use the value from the nextRecordsUrl from the previous set (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryApi.QueryJobs(context.Background()).IsPkChunkingEnabled(isPkChunkingEnabled).JobType(jobType).ConcurrencyMode(concurrencyMode).QueryLocator(queryLocator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.QueryJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryJobs`: QueryJobsInfos
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.QueryJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isPkChunkingEnabled** | **bool** | If set to true, the request only returns information about jobs where PK Chunking is enabled. | 
 **jobType** | **string** | Gets information only about jobs matching the specified job type. | 
 **concurrencyMode** | **string** | For future use. Gets information only about jobs matching the specified concurrency mode. Possible values are serial and parallel. | 
 **queryLocator** | **string** | use the value from the nextRecordsUrl from the previous set | 

### Return type

[**QueryJobsInfos**](QueryJobsInfos.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

