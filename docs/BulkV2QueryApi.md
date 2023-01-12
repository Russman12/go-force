# \BulkV2QueryApi

All URIs are relative to *https://test.salesforce.com/services/data/v56.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkQueryAbortJob**](BulkV2QueryApi.md#BulkQueryAbortJob) | **Patch** /jobs/query/{jobId} | Abort a Job Query
[**BulkQueryCreateJob**](BulkV2QueryApi.md#BulkQueryCreateJob) | **Post** /jobs/query | Create job Query
[**BulkQueryDeleteJob**](BulkV2QueryApi.md#BulkQueryDeleteJob) | **Delete** /jobs/query/{jobId} | Delete Job Query
[**BulkQueryJobInfo**](BulkV2QueryApi.md#BulkQueryJobInfo) | **Get** /jobs/query/{jobId} | Get Job Info Query
[**BulkQueryJobResults**](BulkV2QueryApi.md#BulkQueryJobResults) | **Get** /jobs/query/{jobId}/results | Get Job Query Result
[**BulkQueryJobs**](BulkV2QueryApi.md#BulkQueryJobs) | **Get** /jobs/query | Get All Query Jobs



## BulkQueryAbortJob

> BulkQueryJobInfo BulkQueryAbortJob(ctx, jobId).BulkQueryJobInfo(bulkQueryJobInfo).Execute()

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
    bulkQueryJobInfo := *openapiclient.NewBulkQueryJobInfo("Id_example", "Object_example", "CreatedById_example", "CreatedDate_example", "SystemModstamp_example", "State_example", "ConcurrencyMode_example", float32(123), "Operation_example", "ContentType_example") // BulkQueryJobInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkV2QueryApi.BulkQueryAbortJob(context.Background(), jobId).BulkQueryJobInfo(bulkQueryJobInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2QueryApi.BulkQueryAbortJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkQueryAbortJob`: BulkQueryJobInfo
    fmt.Fprintf(os.Stdout, "Response from `BulkV2QueryApi.BulkQueryAbortJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkQueryAbortJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkQueryJobInfo** | [**BulkQueryJobInfo**](BulkQueryJobInfo.md) |  | 

### Return type

[**BulkQueryJobInfo**](BulkQueryJobInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkQueryCreateJob

> BulkQueryJobInfo BulkQueryCreateJob(ctx).ContentType(contentType).BulkQueryJobCreateRequest(bulkQueryJobCreateRequest).Execute()

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
    bulkQueryJobCreateRequest := *openapiclient.NewBulkQueryJobCreateRequest("Operation_example", "Query_example") // BulkQueryJobCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkV2QueryApi.BulkQueryCreateJob(context.Background()).ContentType(contentType).BulkQueryJobCreateRequest(bulkQueryJobCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2QueryApi.BulkQueryCreateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkQueryCreateJob`: BulkQueryJobInfo
    fmt.Fprintf(os.Stdout, "Response from `BulkV2QueryApi.BulkQueryCreateJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkQueryCreateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **bulkQueryJobCreateRequest** | [**BulkQueryJobCreateRequest**](BulkQueryJobCreateRequest.md) |  | 

### Return type

[**BulkQueryJobInfo**](BulkQueryJobInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkQueryDeleteJob

> BulkQueryDeleteJob(ctx, jobId).Execute()

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
    resp, r, err := apiClient.BulkV2QueryApi.BulkQueryDeleteJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2QueryApi.BulkQueryDeleteJob``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBulkQueryDeleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkQueryJobInfo

> BulkQueryJobInfo BulkQueryJobInfo(ctx, jobId).Execute()

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
    resp, r, err := apiClient.BulkV2QueryApi.BulkQueryJobInfo(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2QueryApi.BulkQueryJobInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkQueryJobInfo`: BulkQueryJobInfo
    fmt.Fprintf(os.Stdout, "Response from `BulkV2QueryApi.BulkQueryJobInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkQueryJobInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BulkQueryJobInfo**](BulkQueryJobInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkQueryJobResults

> io.ReadCloser BulkQueryJobResults(ctx, jobId).ContentType(contentType).Accept(accept).Locator(locator).MaxRecords(maxRecords).Execute()

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
    resp, r, err := apiClient.BulkV2QueryApi.BulkQueryJobResults(context.Background(), jobId).ContentType(contentType).Accept(accept).Locator(locator).MaxRecords(maxRecords).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2QueryApi.BulkQueryJobResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkQueryJobResults`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `BulkV2QueryApi.BulkQueryJobResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkQueryJobResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **accept** | **string** |  | 
 **locator** | **string** | A string that identifies a specific set of query results. Providing a value for this parameter returns only that set of results. Omitting this parameter returns the first set of results.  You can find the locator string for the next set of results in the response of each request. See Example and Rules and Guidelines.  As long as the associated job exists, the locator string for a set of results does not change. You can use the locator to retrieve a set of results multiple times.  | 
 **maxRecords** | **int32** | The maximum number of records to retrieve per set of results for the query. The request is still subject to the size limits. If you are working with a very large number of query results, you may experience a timeout before receiving all the data from Salesforce. To prevent a timeout, specify the maximum number of records your client is expecting to receive in the maxRecords parameter. This splits the results into smaller sets with this value as the maximum size.  If you don’t provide a value for this parameter, the server uses a default value based on the service.  | 

### Return type

[**io.ReadCloser**](io.ReadCloser.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkQueryJobs

> BulkQueryJobsInfos BulkQueryJobs(ctx).IsPkChunkingEnabled(isPkChunkingEnabled).JobType(jobType).ConcurrencyMode(concurrencyMode).QueryLocator(queryLocator).Execute()

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
    resp, r, err := apiClient.BulkV2QueryApi.BulkQueryJobs(context.Background()).IsPkChunkingEnabled(isPkChunkingEnabled).JobType(jobType).ConcurrencyMode(concurrencyMode).QueryLocator(queryLocator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2QueryApi.BulkQueryJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkQueryJobs`: BulkQueryJobsInfos
    fmt.Fprintf(os.Stdout, "Response from `BulkV2QueryApi.BulkQueryJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkQueryJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isPkChunkingEnabled** | **bool** | If set to true, the request only returns information about jobs where PK Chunking is enabled. | 
 **jobType** | **string** | Gets information only about jobs matching the specified job type. | 
 **concurrencyMode** | **string** | For future use. Gets information only about jobs matching the specified concurrency mode. Possible values are serial and parallel. | 
 **queryLocator** | **string** | use the value from the nextRecordsUrl from the previous set | 

### Return type

[**BulkQueryJobsInfos**](BulkQueryJobsInfos.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

