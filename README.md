# Go API client for goforce

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 56.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import goforce "github.com/russman12/go-force"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), goforce.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), goforce.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), goforce.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), goforce.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://test.salesforce.com/services/data/v56.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BulkV2Api* | [**BulkCloseOrAbortJob**](docs/BulkV2Api.md#bulkcloseorabortjob) | **Patch** /jobs/ingest/{jobId} | Close or Abort a Job
*BulkV2Api* | [**BulkDeleteJob**](docs/BulkV2Api.md#bulkdeletejob) | **Delete** /jobs/ingest/{jobId} | Delete Job
*BulkV2Api* | [**BulkJobFailedResults**](docs/BulkV2Api.md#bulkjobfailedresults) | **Get** /jobs/ingest/{jobId}/failedResults | Get Job Failed Record Results
*BulkV2Api* | [**BulkJobSuccessfulResults**](docs/BulkV2Api.md#bulkjobsuccessfulresults) | **Get** /jobs/ingest/{jobId}/successfulResults | Get Job Successful Record Results
*BulkV2Api* | [**BulkJobUnprocessedRecords**](docs/BulkV2Api.md#bulkjobunprocessedrecords) | **Get** /jobs/ingest/{jobId}/unprocessedrecords | Get Job Unprocessed Record Results
*BulkV2Api* | [**BulkJobs**](docs/BulkV2Api.md#bulkjobs) | **Get** /jobs/ingest | Get All Jobs
*BulkV2Api* | [**CreateJob**](docs/BulkV2Api.md#createjob) | **Post** /jobs/ingest | Create job
*BulkV2Api* | [**GetJobInfo**](docs/BulkV2Api.md#getjobinfo) | **Get** /jobs/ingest/{jobId} | Get Job Info
*BulkV2Api* | [**UploadJobData**](docs/BulkV2Api.md#uploadjobdata) | **Put** /jobs/ingest/{jobId}/batches | Upload Job Data
*BulkV2QueryApi* | [**BulkQueryAbortJob**](docs/BulkV2QueryApi.md#bulkqueryabortjob) | **Patch** /jobs/query/{jobId} | Abort a Job Query
*BulkV2QueryApi* | [**BulkQueryCreateJob**](docs/BulkV2QueryApi.md#bulkquerycreatejob) | **Post** /jobs/query | Create job Query
*BulkV2QueryApi* | [**BulkQueryDeleteJob**](docs/BulkV2QueryApi.md#bulkquerydeletejob) | **Delete** /jobs/query/{jobId} | Delete Job Query
*BulkV2QueryApi* | [**BulkQueryJobInfo**](docs/BulkV2QueryApi.md#bulkqueryjobinfo) | **Get** /jobs/query/{jobId} | Get Job Info Query
*BulkV2QueryApi* | [**BulkQueryJobResults**](docs/BulkV2QueryApi.md#bulkqueryjobresults) | **Get** /jobs/query/{jobId}/results | Get Job Query Result
*BulkV2QueryApi* | [**BulkQueryJobs**](docs/BulkV2QueryApi.md#bulkqueryjobs) | **Get** /jobs/query | Get All Query Jobs
*OAuthApi* | [**OAuthUserPass**](docs/OAuthApi.md#oauthuserpass) | **Post** /oauth2/token | Get OAuth2 token
*SObjectApi* | [**SObjectDescribe**](docs/SObjectApi.md#sobjectdescribe) | **Get** /sobjects/{sObject}/describe | 
*SObjectApi* | [**SObjects**](docs/SObjectApi.md#sobjects) | **Get** /sobjects | 


## Documentation For Models

 - [AuthResponse](docs/AuthResponse.md)
 - [BulkCloseOrAbortJobRequest](docs/BulkCloseOrAbortJobRequest.md)
 - [BulkJob](docs/BulkJob.md)
 - [BulkJobInfo](docs/BulkJobInfo.md)
 - [BulkJobInfoExt](docs/BulkJobInfoExt.md)
 - [BulkJobs](docs/BulkJobs.md)
 - [BulkQueryJob](docs/BulkQueryJob.md)
 - [BulkQueryJobCreateRequest](docs/BulkQueryJobCreateRequest.md)
 - [BulkQueryJobCreateRequestExt](docs/BulkQueryJobCreateRequestExt.md)
 - [BulkQueryJobInfo](docs/BulkQueryJobInfo.md)
 - [BulkQueryJobInfoExt](docs/BulkQueryJobInfoExt.md)
 - [BulkQueryJobsInfos](docs/BulkQueryJobsInfos.md)
 - [SObjectDescribe](docs/SObjectDescribe.md)
 - [SObjectDescribeActionOverrides](docs/SObjectDescribeActionOverrides.md)
 - [SObjectDescribeChildRelationships](docs/SObjectDescribeChildRelationships.md)
 - [SObjectDescribeSupportedScopes](docs/SObjectDescribeSupportedScopes.md)
 - [SObjectDescribes](docs/SObjectDescribes.md)
 - [SObjectFieldDescribe](docs/SObjectFieldDescribe.md)
 - [SObjectFieldDescribePicklist](docs/SObjectFieldDescribePicklist.md)
 - [SObjectRecordTypeInfo](docs/SObjectRecordTypeInfo.md)
 - [SObjectRecordTypeInfoUrls](docs/SObjectRecordTypeInfoUrls.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

russell-laboe@outlook.com
