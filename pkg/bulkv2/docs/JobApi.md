# \JobApi

All URIs are relative to *https://test.salesforce.com/services/data/v56.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseOrAbortJob**](JobApi.md#CloseOrAbortJob) | **Patch** /jobs/ingest/{jobId} | Close or Abort a Job
[**CreateJob**](JobApi.md#CreateJob) | **Post** /jobs/ingest | Create job
[**DeleteJob**](JobApi.md#DeleteJob) | **Delete** /jobs/ingest/{jobId} | Delete Job
[**GetJobFailedResults**](JobApi.md#GetJobFailedResults) | **Get** /jobs/ingest/{jobId}/failedResults | Get Job Failed Record Results
[**GetJobInfo**](JobApi.md#GetJobInfo) | **Get** /jobs/ingest/{jobId} | Get Job Info
[**GetJobSuccessfulResults**](JobApi.md#GetJobSuccessfulResults) | **Get** /jobs/ingest/{jobId}/successfulResults | Get Job Successful Record Results
[**GetJobUnprocessedRecords**](JobApi.md#GetJobUnprocessedRecords) | **Get** /jobs/ingest/{jobId}/unprocessedrecords | Get Job Unprocessed Record Results
[**GetJobs**](JobApi.md#GetJobs) | **Get** /jobs/ingest | Get All Jobs
[**UploadJobData**](JobApi.md#UploadJobData) | **Put** /jobs/ingest/{jobId}/batches | Upload Job Data



## CloseOrAbortJob

> JobInfo CloseOrAbortJob(ctx, jobId).CloseOrAbortJobRequest(closeOrAbortJobRequest).Execute()

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
    closeOrAbortJobRequest := *openapiclient.NewCloseOrAbortJobRequest(openapiclient.JobCloseAbortState("UploadComplete")) // CloseOrAbortJobRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobApi.CloseOrAbortJob(context.Background(), jobId).CloseOrAbortJobRequest(closeOrAbortJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.CloseOrAbortJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloseOrAbortJob`: JobInfo
    fmt.Fprintf(os.Stdout, "Response from `JobApi.CloseOrAbortJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseOrAbortJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **closeOrAbortJobRequest** | [**CloseOrAbortJobRequest**](CloseOrAbortJobRequest.md) |  | 

### Return type

[**JobInfo**](JobInfo.md)

### Authorization

[oAuth](../README.md#oAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJob

> JobInfo CreateJob(ctx).CreateJobRequest(createJobRequest).Execute()

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
    createJobRequest := *openapiclient.NewCreateJobRequest("Object_example", openapiclient.JobOperation("insert")) // CreateJobRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobApi.CreateJob(context.Background()).CreateJobRequest(createJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.CreateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJob`: JobInfo
    fmt.Fprintf(os.Stdout, "Response from `JobApi.CreateJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createJobRequest** | [**CreateJobRequest**](CreateJobRequest.md) |  | 

### Return type

[**JobInfo**](JobInfo.md)

### Authorization

[oAuth](../README.md#oAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJob

> DeleteJob(ctx, jobId).Execute()

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
    resp, r, err := apiClient.JobApi.DeleteJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.DeleteJob``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobFailedResults

> io.ReadCloser GetJobFailedResults(ctx, jobId).Execute()

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
    resp, r, err := apiClient.JobApi.GetJobFailedResults(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.GetJobFailedResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobFailedResults`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `JobApi.GetJobFailedResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobFailedResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**io.ReadCloser**](io.ReadCloser.md)

### Authorization

[oAuth](../README.md#oAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobInfo

> JobInfo GetJobInfo(ctx, jobId).Execute()

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
    resp, r, err := apiClient.JobApi.GetJobInfo(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.GetJobInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobInfo`: JobInfo
    fmt.Fprintf(os.Stdout, "Response from `JobApi.GetJobInfo`: %v\n", resp)
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

[**JobInfo**](JobInfo.md)

### Authorization

[oAuth](../README.md#oAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobSuccessfulResults

> io.ReadCloser GetJobSuccessfulResults(ctx, jobId).Execute()

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
    resp, r, err := apiClient.JobApi.GetJobSuccessfulResults(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.GetJobSuccessfulResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobSuccessfulResults`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `JobApi.GetJobSuccessfulResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobSuccessfulResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**io.ReadCloser**](io.ReadCloser.md)

### Authorization

[oAuth](../README.md#oAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobUnprocessedRecords

> io.ReadCloser GetJobUnprocessedRecords(ctx, jobId).Execute()

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
    resp, r, err := apiClient.JobApi.GetJobUnprocessedRecords(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.GetJobUnprocessedRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobUnprocessedRecords`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `JobApi.GetJobUnprocessedRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobUnprocessedRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**io.ReadCloser**](io.ReadCloser.md)

### Authorization

[oAuth](../README.md#oAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobs

> map[string]interface{} GetJobs(ctx).IsPkChunkingEnabled(isPkChunkingEnabled).JobType(jobType).QueryLocator(queryLocator).Execute()

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
    resp, r, err := apiClient.JobApi.GetJobs(context.Background()).IsPkChunkingEnabled(isPkChunkingEnabled).JobType(jobType).QueryLocator(queryLocator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.GetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `JobApi.GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isPkChunkingEnabled** | **bool** |  | 
 **jobType** | **string** |  | 
 **queryLocator** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[oAuth](../README.md#oAuth), [oAuth](../README.md#oAuth)

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
    resp, r, err := apiClient.JobApi.UploadJobData(context.Background(), jobId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.UploadJobData``: %v\n", err)
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

[oAuth](../README.md#oAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: text/csv
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

