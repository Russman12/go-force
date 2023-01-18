# JobApi

All URIs are relative to *https://test.salesforce.com/services/data/v56.0*

| Method        | HTTP request  | Summary       |
| ------------- | ------------- | ------------- |
| [**CloseOrAbortJob**](#closeorabortjob) | **Patch** /jobs/ingest/{jobId} | Close or Abort a Job |
| [**CreateJob**](#createjob) | **Post** /jobs/ingest | Create a job |
| [**DeleteJob**](#deletejob) | **Delete** /jobs/ingest/{jobId} | Delete a Job |
| [**GetJobFailedResults**](#getjobfailedresults) | **Get** /jobs/ingest/{jobId}/failedResults | Get Job Failed Record Results |
| [**GetJobInfo**](#getjobinfo) | **Get** /jobs/ingest/{jobId} | Get Job Info |
| [**GetJobSuccessfulResults**](#getjobsuccessfulresults) | **Get** /jobs/ingest/{jobId}/successfulResults | Get Job Successful Record Results |
| [**GetJobUnprocessedRecords**](#getjobunprocessedrecords) | **Get** /jobs/ingest/{jobId}/unprocessedrecords | Get Job Unprocessed Record Results |
| [**GetJobs**](#getjobs) | **Get** /jobs/ingest | Get All Jobs |
| [**UploadJobData**](#uploadjobdata) | **Put** /jobs/ingest/{jobId}/batches | Upload Job Data |



## CloseOrAbortJob

> JobInfo CloseOrAbortJob(ctx, jobId).CloseOrAbortJobRequest(closeOrAbortJobRequest).ContentEncoding(contentEncoding).Execute()

Closes or aborts a job. If you close a job, Salesforce queues the job and uploaded data for processing, and you can’t add any more job data. If you abort a job, the job doesn’t get queued or processed.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/close_job.htm)

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
    closeOrAbortJobRequest := CloseOrAbortJobRequest{} // CloseOrAbortJobRequest | 
    contentEncoding := EncodingType{} // EncodingType |  (optional)

    resp, r, err := apiClient.JobApi.CloseOrAbortJob(context.Background(), jobId).CloseOrAbortJobRequest(closeOrAbortJobRequest).ContentEncoding(contentEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.CloseOrAbortJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloseOrAbortJob`: JobInfo
    fmt.Fprintf(os.Stdout, "Response from `JobApi.CloseOrAbortJob`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type          | Description   | Notes        |
| ------------ | ------------- | ------------- | ------------ |
| **ctx** | **context.Context** | context for logging, cancellation, deadlines, tracing, etc. | |
| **jobId** | **string** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCloseOrAbortJobRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **closeOrAbortJobRequest** | [**CloseOrAbortJobRequest**](CloseOrAbortJobRequest.md) |  | 
|  **contentEncoding** | [**EncodingType**](EncodingType.md) |  |  |

### Return type

[**JobInfo**](JobInfo.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJob

> JobInfo CreateJob(ctx).CreateJobRequest(createJobRequest).SforceCallOptions(sforceCallOptions).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()

Creates a job representing a bulk operation and its associated data that is sent to Salesforce for asynchronous processing. Provide job data via an Upload Job Data request or as part of a multipart create job request.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/create_job.htm)

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

    createJobRequest := CreateJobRequest{} // CreateJobRequest | 
    sforceCallOptions := "client=caseSensitiveToken; defaultNamespace=battle" // string |  (optional)
    contentEncoding := EncodingType{} // EncodingType |  (optional)
    acceptEncoding := EncodingType{} // EncodingType |  (optional)

    resp, r, err := apiClient.JobApi.CreateJob(context.Background()).CreateJobRequest(createJobRequest).SforceCallOptions(sforceCallOptions).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.CreateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJob`: JobInfo
    fmt.Fprintf(os.Stdout, "Response from `JobApi.CreateJob`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  **createJobRequest** | [**CreateJobRequest**](CreateJobRequest.md) |  | 
|  **sforceCallOptions** | **string** |  | 
|  **contentEncoding** | [**EncodingType**](EncodingType.md) |  | 
|  **acceptEncoding** | [**EncodingType**](EncodingType.md) |  |  |

### Return type

[**JobInfo**](JobInfo.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJob

> DeleteJob(ctx, jobId).Execute()

Deletes a job. To be deleted, a job must have a state of `UploadComplete`, `JobComplete`, `Aborted`, or `Failed`.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/delete_job.htm)

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

    resp, r, err := apiClient.JobApi.DeleteJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.DeleteJob``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteJobRequest struct via the builder pattern


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


## GetJobFailedResults

> io.ReadCloser GetJobFailedResults(ctx, jobId).AcceptEncoding(acceptEncoding).Execute()

Retrieves a list of failed records for a completed insert, delete, update, or upsert job.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/get_job_failed_results.htm)

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

    resp, r, err := apiClient.JobApi.GetJobFailedResults(context.Background(), jobId).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.GetJobFailedResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobFailedResults`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `JobApi.GetJobFailedResults`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type          | Description   | Notes        |
| ------------ | ------------- | ------------- | ------------ |
| **ctx** | **context.Context** | context for logging, cancellation, deadlines, tracing, etc. | |
| **jobId** | **string** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobFailedResultsRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **acceptEncoding** | [**EncodingType**](EncodingType.md) |  |  |

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


## GetJobInfo

> JobInfo GetJobInfo(ctx, jobId).AcceptEncoding(acceptEncoding).Execute()

Retrieves detailed information about a job.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/get_job_info.htm)

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

    resp, r, err := apiClient.JobApi.GetJobInfo(context.Background(), jobId).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.GetJobInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobInfo`: JobInfo
    fmt.Fprintf(os.Stdout, "Response from `JobApi.GetJobInfo`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type          | Description   | Notes        |
| ------------ | ------------- | ------------- | ------------ |
| **ctx** | **context.Context** | context for logging, cancellation, deadlines, tracing, etc. | |
| **jobId** | **string** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobInfoRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **acceptEncoding** | [**EncodingType**](EncodingType.md) |  |  |

### Return type

[**JobInfo**](JobInfo.md)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobSuccessfulResults

> io.ReadCloser GetJobSuccessfulResults(ctx, jobId).AcceptEncoding(acceptEncoding).Execute()

Retrieves a list of successfully processed records for a completed job.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/get_job_successful_results.htm)

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

    resp, r, err := apiClient.JobApi.GetJobSuccessfulResults(context.Background(), jobId).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.GetJobSuccessfulResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobSuccessfulResults`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `JobApi.GetJobSuccessfulResults`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type          | Description   | Notes        |
| ------------ | ------------- | ------------- | ------------ |
| **ctx** | **context.Context** | context for logging, cancellation, deadlines, tracing, etc. | |
| **jobId** | **string** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobSuccessfulResultsRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **acceptEncoding** | [**EncodingType**](EncodingType.md) |  |  |

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


## GetJobUnprocessedRecords

> io.ReadCloser GetJobUnprocessedRecords(ctx, jobId).AcceptEncoding(acceptEncoding).Execute()

Retrieves a list of unprocessed records for failed or aborted jobs.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/get_job_unprocessed_results.htm)

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

    resp, r, err := apiClient.JobApi.GetJobUnprocessedRecords(context.Background(), jobId).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.GetJobUnprocessedRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobUnprocessedRecords`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `JobApi.GetJobUnprocessedRecords`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type          | Description   | Notes        |
| ------------ | ------------- | ------------- | ------------ |
| **ctx** | **context.Context** | context for logging, cancellation, deadlines, tracing, etc. | |
| **jobId** | **string** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobUnprocessedRecordsRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **acceptEncoding** | [**EncodingType**](EncodingType.md) |  |  |

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


## GetJobs

> map[string]interface{} GetJobs(ctx).AcceptEncoding(acceptEncoding).IsPkChunkingEnabled(isPkChunkingEnabled).JobType(jobType).QueryLocator(queryLocator).Execute()

Retrieves all jobs in the org.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/get_all_jobs.htm)

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

    acceptEncoding := EncodingType{} // EncodingType |  (optional)
    isPkChunkingEnabled := bool{} // bool |  (optional)
    jobType := "jobType_example" // string |  (optional)
    queryLocator := "queryLocator_example" // string |  (optional)

    resp, r, err := apiClient.JobApi.GetJobs(context.Background()).AcceptEncoding(acceptEncoding).IsPkChunkingEnabled(isPkChunkingEnabled).JobType(jobType).QueryLocator(queryLocator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.GetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `JobApi.GetJobs`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
|  **acceptEncoding** | [**EncodingType**](EncodingType.md) |  | 
|  **isPkChunkingEnabled** | **bool** |  | 
|  **jobType** | **string** |  | 
|  **queryLocator** | **string** |  |  |

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


## UploadJobData

> UploadJobData(ctx, jobId).Body(body).ContentEncoding(contentEncoding).Execute()

Uploads data for a job using CSV data you provide.

For more details see [Salesforce Documentation](https://developer.salesforce.com/docs/atlas.en-us.api_asynch.meta/api_asynch/upload_job_data.htm)

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
    body := io.ReadCloser{} // io.ReadCloser | 
    contentEncoding := EncodingType{} // EncodingType |  (optional)

    resp, r, err := apiClient.JobApi.UploadJobData(context.Background(), jobId).Body(body).ContentEncoding(contentEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.UploadJobData``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUploadJobDataRequest struct via the builder pattern


| Name          | Type          | Description   | Notes         |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **body** | **io.ReadCloser** |  | 
|  **contentEncoding** | [**EncodingType**](EncodingType.md) |  |  |

### Return type

 (empty response body)

### Authorization

[oAuth (password)](../README.md#oauth--password-), [oAuth (application)](../README.md#oauth--application-)

### HTTP request headers

- **Content-Type**: text/csv
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

