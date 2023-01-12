# \BulkV2Api

All URIs are relative to *https://test.salesforce.com/services/data/v56.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCloseOrAbortJob**](BulkV2Api.md#BulkCloseOrAbortJob) | **Patch** /jobs/ingest/{jobId} | Close or Abort a Job
[**BulkDeleteJob**](BulkV2Api.md#BulkDeleteJob) | **Delete** /jobs/ingest/{jobId} | Delete Job
[**BulkJobFailedResults**](BulkV2Api.md#BulkJobFailedResults) | **Get** /jobs/ingest/{jobId}/failedResults | Get Job Failed Record Results
[**BulkJobSuccessfulResults**](BulkV2Api.md#BulkJobSuccessfulResults) | **Get** /jobs/ingest/{jobId}/successfulResults | Get Job Successful Record Results
[**BulkJobUnprocessedRecords**](BulkV2Api.md#BulkJobUnprocessedRecords) | **Get** /jobs/ingest/{jobId}/unprocessedrecords | Get Job Unprocessed Record Results
[**BulkJobs**](BulkV2Api.md#BulkJobs) | **Get** /jobs/ingest | Get All Jobs
[**CreateJob**](BulkV2Api.md#CreateJob) | **Post** /jobs/ingest | Create job
[**GetJobInfo**](BulkV2Api.md#GetJobInfo) | **Get** /jobs/ingest/{jobId} | Get Job Info
[**UploadJobData**](BulkV2Api.md#UploadJobData) | **Put** /jobs/ingest/{jobId}/batches | Upload Job Data



## BulkCloseOrAbortJob

> BulkJobInfo BulkCloseOrAbortJob(ctx, jobId).BulkCloseOrAbortJobRequest(bulkCloseOrAbortJobRequest).Execute()

Close or Abort a Job

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
    bulkCloseOrAbortJobRequest := *openapiclient.NewBulkCloseOrAbortJobRequest("State_example") // BulkCloseOrAbortJobRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkV2Api.BulkCloseOrAbortJob(context.Background(), jobId).BulkCloseOrAbortJobRequest(bulkCloseOrAbortJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2Api.BulkCloseOrAbortJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCloseOrAbortJob`: BulkJobInfo
    fmt.Fprintf(os.Stdout, "Response from `BulkV2Api.BulkCloseOrAbortJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkCloseOrAbortJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkCloseOrAbortJobRequest** | [**BulkCloseOrAbortJobRequest**](BulkCloseOrAbortJobRequest.md) |  | 

### Return type

[**BulkJobInfo**](BulkJobInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDeleteJob

> BulkDeleteJob(ctx, jobId).Execute()

Delete Job

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
    resp, r, err := apiClient.BulkV2Api.BulkDeleteJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2Api.BulkDeleteJob``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBulkDeleteJobRequest struct via the builder pattern


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


## BulkJobFailedResults

> io.ReadCloser BulkJobFailedResults(ctx, jobId).Execute()

Get Job Failed Record Results

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
    resp, r, err := apiClient.BulkV2Api.BulkJobFailedResults(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2Api.BulkJobFailedResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkJobFailedResults`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `BulkV2Api.BulkJobFailedResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkJobFailedResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BulkJobSuccessfulResults

> io.ReadCloser BulkJobSuccessfulResults(ctx, jobId).Execute()

Get Job Successful Record Results

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
    resp, r, err := apiClient.BulkV2Api.BulkJobSuccessfulResults(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2Api.BulkJobSuccessfulResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkJobSuccessfulResults`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `BulkV2Api.BulkJobSuccessfulResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkJobSuccessfulResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BulkJobUnprocessedRecords

> io.ReadCloser BulkJobUnprocessedRecords(ctx, jobId).Execute()

Get Job Unprocessed Record Results

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
    resp, r, err := apiClient.BulkV2Api.BulkJobUnprocessedRecords(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2Api.BulkJobUnprocessedRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkJobUnprocessedRecords`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `BulkV2Api.BulkJobUnprocessedRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkJobUnprocessedRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BulkJobs

> map[string]interface{} BulkJobs(ctx).IsPkChunkingEnabled(isPkChunkingEnabled).JobType(jobType).QueryLocator(queryLocator).Execute()

Get All Jobs

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
    isPkChunkingEnabled := true // bool |  (optional)
    jobType := "jobType_example" // string |  (optional)
    queryLocator := "queryLocator_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkV2Api.BulkJobs(context.Background()).IsPkChunkingEnabled(isPkChunkingEnabled).JobType(jobType).QueryLocator(queryLocator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2Api.BulkJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkJobs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BulkV2Api.BulkJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isPkChunkingEnabled** | **bool** |  | 
 **jobType** | **string** |  | 
 **queryLocator** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJob

> BulkJobInfo CreateJob(ctx).BulkJob(bulkJob).Execute()

Create job

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
    bulkJob := *openapiclient.NewBulkJob("Object_example", "Operation_example") // BulkJob |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkV2Api.CreateJob(context.Background()).BulkJob(bulkJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2Api.CreateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJob`: BulkJobInfo
    fmt.Fprintf(os.Stdout, "Response from `BulkV2Api.CreateJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkJob** | [**BulkJob**](BulkJob.md) |  | 

### Return type

[**BulkJobInfo**](BulkJobInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobInfo

> BulkJobInfo GetJobInfo(ctx, jobId).Execute()

Get Job Info

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
    resp, r, err := apiClient.BulkV2Api.GetJobInfo(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2Api.GetJobInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobInfo`: BulkJobInfo
    fmt.Fprintf(os.Stdout, "Response from `BulkV2Api.GetJobInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BulkJobInfo**](BulkJobInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadJobData

> UploadJobData(ctx, jobId).Body(body).Execute()

Upload Job Data

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
    body := os.NewFile(1234, "some_file") // io.ReadCloser |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkV2Api.UploadJobData(context.Background(), jobId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkV2Api.UploadJobData``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUploadJobDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **io.ReadCloser** |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: text/csv
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

