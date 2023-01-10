# BulkJobInfoExt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **float32** | The API version that the job was created in. | [optional] 
**ConcurrencyMode** | Pointer to **string** | For future use. How the request was processed. Currently only parallel mode is supported. (When other modes are added, the mode will be chosen automatically by the API and will not be user configurable.) | [optional] 
**ContentUrl** | Pointer to **string** | The URL to use for Upload Job Data requests for this job. Only valid if the job is in Open state. | [optional] 
**CreatedById** | Pointer to **string** | The ID of the user who created the job. | [optional] 
**CreatedDate** | Pointer to **string** | The date and time in the UTC time zone when the job was created. | [optional] 
**Id** | Pointer to **string** | Unique ID for this job. | [optional] 
**JobType** | Pointer to **string** | The job’s type. | [optional] 
**SystemModstamp** | Pointer to **string** | Date and time in the UTC time zone when the job finished. | [optional] 
**State** | Pointer to **string** | The current state of processing for the job. | [optional] 
**ApexProcessingTime** | Pointer to **int64** | The number of milliseconds taken to process triggers and other processes related to the job data. This doesn&#39;t include the time used for processing asynchronous and batch Apex operations. If there are no triggers, the value is 0. | [optional] 
**ApiActiveProcessingTime** | Pointer to **int64** | The number of milliseconds taken to actively process the job and includes apexProcessingTime, but doesn&#39;t include the time the job waited in the queue to be processed or the time required for serialization and deserialization. | [optional] 
**ErrorMessage** | Pointer to **string** | The error message shown for jobs with errors. | [optional] 
**NumberRecordsFailed** | Pointer to **int64** | The number of records that were not processed successfully in this job. | [optional] 
**NumberRecordsProcessed** | Pointer to **int64** | The number of records already processed. | [optional] 
**Retries** | Pointer to **int32** | The number of times that Salesforce attempted to save the results of an operation. The repeated attempts are due to a problem, such as a lock contention. | [optional] 
**TotalProcessingTime** | Pointer to **int64** | The number of milliseconds taken to process the job. | [optional] 

## Methods

### NewBulkJobInfoExt

`func NewBulkJobInfoExt() *BulkJobInfoExt`

NewBulkJobInfoExt instantiates a new BulkJobInfoExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkJobInfoExtWithDefaults

`func NewBulkJobInfoExtWithDefaults() *BulkJobInfoExt`

NewBulkJobInfoExtWithDefaults instantiates a new BulkJobInfoExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *BulkJobInfoExt) GetApiVersion() float32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BulkJobInfoExt) GetApiVersionOk() (*float32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BulkJobInfoExt) SetApiVersion(v float32)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BulkJobInfoExt) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetConcurrencyMode

`func (o *BulkJobInfoExt) GetConcurrencyMode() string`

GetConcurrencyMode returns the ConcurrencyMode field if non-nil, zero value otherwise.

### GetConcurrencyModeOk

`func (o *BulkJobInfoExt) GetConcurrencyModeOk() (*string, bool)`

GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyMode

`func (o *BulkJobInfoExt) SetConcurrencyMode(v string)`

SetConcurrencyMode sets ConcurrencyMode field to given value.

### HasConcurrencyMode

`func (o *BulkJobInfoExt) HasConcurrencyMode() bool`

HasConcurrencyMode returns a boolean if a field has been set.

### GetContentUrl

`func (o *BulkJobInfoExt) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *BulkJobInfoExt) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *BulkJobInfoExt) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.

### HasContentUrl

`func (o *BulkJobInfoExt) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### GetCreatedById

`func (o *BulkJobInfoExt) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *BulkJobInfoExt) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *BulkJobInfoExt) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *BulkJobInfoExt) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedDate

`func (o *BulkJobInfoExt) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *BulkJobInfoExt) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *BulkJobInfoExt) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *BulkJobInfoExt) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetId

`func (o *BulkJobInfoExt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkJobInfoExt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkJobInfoExt) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkJobInfoExt) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobType

`func (o *BulkJobInfoExt) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BulkJobInfoExt) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BulkJobInfoExt) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *BulkJobInfoExt) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetSystemModstamp

`func (o *BulkJobInfoExt) GetSystemModstamp() string`

GetSystemModstamp returns the SystemModstamp field if non-nil, zero value otherwise.

### GetSystemModstampOk

`func (o *BulkJobInfoExt) GetSystemModstampOk() (*string, bool)`

GetSystemModstampOk returns a tuple with the SystemModstamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModstamp

`func (o *BulkJobInfoExt) SetSystemModstamp(v string)`

SetSystemModstamp sets SystemModstamp field to given value.

### HasSystemModstamp

`func (o *BulkJobInfoExt) HasSystemModstamp() bool`

HasSystemModstamp returns a boolean if a field has been set.

### GetState

`func (o *BulkJobInfoExt) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BulkJobInfoExt) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BulkJobInfoExt) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BulkJobInfoExt) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApexProcessingTime

`func (o *BulkJobInfoExt) GetApexProcessingTime() int64`

GetApexProcessingTime returns the ApexProcessingTime field if non-nil, zero value otherwise.

### GetApexProcessingTimeOk

`func (o *BulkJobInfoExt) GetApexProcessingTimeOk() (*int64, bool)`

GetApexProcessingTimeOk returns a tuple with the ApexProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApexProcessingTime

`func (o *BulkJobInfoExt) SetApexProcessingTime(v int64)`

SetApexProcessingTime sets ApexProcessingTime field to given value.

### HasApexProcessingTime

`func (o *BulkJobInfoExt) HasApexProcessingTime() bool`

HasApexProcessingTime returns a boolean if a field has been set.

### GetApiActiveProcessingTime

`func (o *BulkJobInfoExt) GetApiActiveProcessingTime() int64`

GetApiActiveProcessingTime returns the ApiActiveProcessingTime field if non-nil, zero value otherwise.

### GetApiActiveProcessingTimeOk

`func (o *BulkJobInfoExt) GetApiActiveProcessingTimeOk() (*int64, bool)`

GetApiActiveProcessingTimeOk returns a tuple with the ApiActiveProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiActiveProcessingTime

`func (o *BulkJobInfoExt) SetApiActiveProcessingTime(v int64)`

SetApiActiveProcessingTime sets ApiActiveProcessingTime field to given value.

### HasApiActiveProcessingTime

`func (o *BulkJobInfoExt) HasApiActiveProcessingTime() bool`

HasApiActiveProcessingTime returns a boolean if a field has been set.

### GetErrorMessage

`func (o *BulkJobInfoExt) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BulkJobInfoExt) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BulkJobInfoExt) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BulkJobInfoExt) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetNumberRecordsFailed

`func (o *BulkJobInfoExt) GetNumberRecordsFailed() int64`

GetNumberRecordsFailed returns the NumberRecordsFailed field if non-nil, zero value otherwise.

### GetNumberRecordsFailedOk

`func (o *BulkJobInfoExt) GetNumberRecordsFailedOk() (*int64, bool)`

GetNumberRecordsFailedOk returns a tuple with the NumberRecordsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberRecordsFailed

`func (o *BulkJobInfoExt) SetNumberRecordsFailed(v int64)`

SetNumberRecordsFailed sets NumberRecordsFailed field to given value.

### HasNumberRecordsFailed

`func (o *BulkJobInfoExt) HasNumberRecordsFailed() bool`

HasNumberRecordsFailed returns a boolean if a field has been set.

### GetNumberRecordsProcessed

`func (o *BulkJobInfoExt) GetNumberRecordsProcessed() int64`

GetNumberRecordsProcessed returns the NumberRecordsProcessed field if non-nil, zero value otherwise.

### GetNumberRecordsProcessedOk

`func (o *BulkJobInfoExt) GetNumberRecordsProcessedOk() (*int64, bool)`

GetNumberRecordsProcessedOk returns a tuple with the NumberRecordsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberRecordsProcessed

`func (o *BulkJobInfoExt) SetNumberRecordsProcessed(v int64)`

SetNumberRecordsProcessed sets NumberRecordsProcessed field to given value.

### HasNumberRecordsProcessed

`func (o *BulkJobInfoExt) HasNumberRecordsProcessed() bool`

HasNumberRecordsProcessed returns a boolean if a field has been set.

### GetRetries

`func (o *BulkJobInfoExt) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *BulkJobInfoExt) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *BulkJobInfoExt) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *BulkJobInfoExt) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetTotalProcessingTime

`func (o *BulkJobInfoExt) GetTotalProcessingTime() int64`

GetTotalProcessingTime returns the TotalProcessingTime field if non-nil, zero value otherwise.

### GetTotalProcessingTimeOk

`func (o *BulkJobInfoExt) GetTotalProcessingTimeOk() (*int64, bool)`

GetTotalProcessingTimeOk returns a tuple with the TotalProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProcessingTime

`func (o *BulkJobInfoExt) SetTotalProcessingTime(v int64)`

SetTotalProcessingTime sets TotalProcessingTime field to given value.

### HasTotalProcessingTime

`func (o *BulkJobInfoExt) HasTotalProcessingTime() bool`

HasTotalProcessingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


