# BulkQueryJobInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID for this job. | 
**Object** | **string** | The object type being queried. | 
**CreatedById** | **string** | The ID of the user who created the job. | 
**CreatedDate** | **string** | The UTC date and time when the job was created. | 
**SystemModstamp** | **string** | The UTC date and time when the API last updated the job information. | 
**State** | **string** | The current state of processing for the job. | 
**ConcurrencyMode** | **string** | Reserved for future use. How the request is processed. Currently only parallel mode is supported. (When other modes are added, the API chooses the mode automatically. The mode isn’t user configurable.) | 
**ApiVersion** | **float32** | The API version that the job was created in. | 
**JobType** | Pointer to **string** | The job’s type. For a query job, the type is always V2Query. | [optional] 
**NumberRecordsProcessed** | Pointer to **int64** | The number of records processed in this job. | [optional] 
**Retries** | Pointer to **int32** | The number of times that Salesforce attempted to save the results of an operation. Repeated attempts indicate a problem such as a lock contention. | [optional] 
**TotalProcessingTime** | Pointer to **int64** | The number of milliseconds taken to process the job. | [optional] 
**Operation** | **string** | The type of query. | 
**ContentType** | **string** | The format that is used for the results. Currently the only supported value is CSV. | 
**ColumnDelimiter** | Pointer to **string** | The column delimiter used for CSV job data. The default value is COMMA. | [optional] 
**LineEnding** | Pointer to **string** | The line ending used for CSV job data, marking the end of a data row. The default is LF. | [optional] 

## Methods

### NewBulkQueryJobInfo

`func NewBulkQueryJobInfo(id string, object string, createdById string, createdDate string, systemModstamp string, state string, concurrencyMode string, apiVersion float32, operation string, contentType string, ) *BulkQueryJobInfo`

NewBulkQueryJobInfo instantiates a new BulkQueryJobInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkQueryJobInfoWithDefaults

`func NewBulkQueryJobInfoWithDefaults() *BulkQueryJobInfo`

NewBulkQueryJobInfoWithDefaults instantiates a new BulkQueryJobInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkQueryJobInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkQueryJobInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkQueryJobInfo) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *BulkQueryJobInfo) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BulkQueryJobInfo) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BulkQueryJobInfo) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreatedById

`func (o *BulkQueryJobInfo) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *BulkQueryJobInfo) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *BulkQueryJobInfo) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetCreatedDate

`func (o *BulkQueryJobInfo) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *BulkQueryJobInfo) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *BulkQueryJobInfo) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetSystemModstamp

`func (o *BulkQueryJobInfo) GetSystemModstamp() string`

GetSystemModstamp returns the SystemModstamp field if non-nil, zero value otherwise.

### GetSystemModstampOk

`func (o *BulkQueryJobInfo) GetSystemModstampOk() (*string, bool)`

GetSystemModstampOk returns a tuple with the SystemModstamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModstamp

`func (o *BulkQueryJobInfo) SetSystemModstamp(v string)`

SetSystemModstamp sets SystemModstamp field to given value.


### GetState

`func (o *BulkQueryJobInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BulkQueryJobInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BulkQueryJobInfo) SetState(v string)`

SetState sets State field to given value.


### GetConcurrencyMode

`func (o *BulkQueryJobInfo) GetConcurrencyMode() string`

GetConcurrencyMode returns the ConcurrencyMode field if non-nil, zero value otherwise.

### GetConcurrencyModeOk

`func (o *BulkQueryJobInfo) GetConcurrencyModeOk() (*string, bool)`

GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyMode

`func (o *BulkQueryJobInfo) SetConcurrencyMode(v string)`

SetConcurrencyMode sets ConcurrencyMode field to given value.


### GetApiVersion

`func (o *BulkQueryJobInfo) GetApiVersion() float32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BulkQueryJobInfo) GetApiVersionOk() (*float32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BulkQueryJobInfo) SetApiVersion(v float32)`

SetApiVersion sets ApiVersion field to given value.


### GetJobType

`func (o *BulkQueryJobInfo) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BulkQueryJobInfo) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BulkQueryJobInfo) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *BulkQueryJobInfo) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetNumberRecordsProcessed

`func (o *BulkQueryJobInfo) GetNumberRecordsProcessed() int64`

GetNumberRecordsProcessed returns the NumberRecordsProcessed field if non-nil, zero value otherwise.

### GetNumberRecordsProcessedOk

`func (o *BulkQueryJobInfo) GetNumberRecordsProcessedOk() (*int64, bool)`

GetNumberRecordsProcessedOk returns a tuple with the NumberRecordsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberRecordsProcessed

`func (o *BulkQueryJobInfo) SetNumberRecordsProcessed(v int64)`

SetNumberRecordsProcessed sets NumberRecordsProcessed field to given value.

### HasNumberRecordsProcessed

`func (o *BulkQueryJobInfo) HasNumberRecordsProcessed() bool`

HasNumberRecordsProcessed returns a boolean if a field has been set.

### GetRetries

`func (o *BulkQueryJobInfo) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *BulkQueryJobInfo) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *BulkQueryJobInfo) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *BulkQueryJobInfo) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetTotalProcessingTime

`func (o *BulkQueryJobInfo) GetTotalProcessingTime() int64`

GetTotalProcessingTime returns the TotalProcessingTime field if non-nil, zero value otherwise.

### GetTotalProcessingTimeOk

`func (o *BulkQueryJobInfo) GetTotalProcessingTimeOk() (*int64, bool)`

GetTotalProcessingTimeOk returns a tuple with the TotalProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProcessingTime

`func (o *BulkQueryJobInfo) SetTotalProcessingTime(v int64)`

SetTotalProcessingTime sets TotalProcessingTime field to given value.

### HasTotalProcessingTime

`func (o *BulkQueryJobInfo) HasTotalProcessingTime() bool`

HasTotalProcessingTime returns a boolean if a field has been set.

### GetOperation

`func (o *BulkQueryJobInfo) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BulkQueryJobInfo) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BulkQueryJobInfo) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetContentType

`func (o *BulkQueryJobInfo) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BulkQueryJobInfo) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BulkQueryJobInfo) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetColumnDelimiter

`func (o *BulkQueryJobInfo) GetColumnDelimiter() string`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *BulkQueryJobInfo) GetColumnDelimiterOk() (*string, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *BulkQueryJobInfo) SetColumnDelimiter(v string)`

SetColumnDelimiter sets ColumnDelimiter field to given value.

### HasColumnDelimiter

`func (o *BulkQueryJobInfo) HasColumnDelimiter() bool`

HasColumnDelimiter returns a boolean if a field has been set.

### GetLineEnding

`func (o *BulkQueryJobInfo) GetLineEnding() string`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *BulkQueryJobInfo) GetLineEndingOk() (*string, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *BulkQueryJobInfo) SetLineEnding(v string)`

SetLineEnding sets LineEnding field to given value.

### HasLineEnding

`func (o *BulkQueryJobInfo) HasLineEnding() bool`

HasLineEnding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


