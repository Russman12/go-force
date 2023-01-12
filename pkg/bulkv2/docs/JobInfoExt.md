# JobInfoExt

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

### NewJobInfoExt

`func NewJobInfoExt() *JobInfoExt`

NewJobInfoExt instantiates a new JobInfoExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobInfoExtWithDefaults

`func NewJobInfoExtWithDefaults() *JobInfoExt`

NewJobInfoExtWithDefaults instantiates a new JobInfoExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *JobInfoExt) GetApiVersion() float32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *JobInfoExt) GetApiVersionOk() (*float32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *JobInfoExt) SetApiVersion(v float32)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *JobInfoExt) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetConcurrencyMode

`func (o *JobInfoExt) GetConcurrencyMode() string`

GetConcurrencyMode returns the ConcurrencyMode field if non-nil, zero value otherwise.

### GetConcurrencyModeOk

`func (o *JobInfoExt) GetConcurrencyModeOk() (*string, bool)`

GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyMode

`func (o *JobInfoExt) SetConcurrencyMode(v string)`

SetConcurrencyMode sets ConcurrencyMode field to given value.

### HasConcurrencyMode

`func (o *JobInfoExt) HasConcurrencyMode() bool`

HasConcurrencyMode returns a boolean if a field has been set.

### GetContentUrl

`func (o *JobInfoExt) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *JobInfoExt) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *JobInfoExt) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.

### HasContentUrl

`func (o *JobInfoExt) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### GetCreatedById

`func (o *JobInfoExt) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *JobInfoExt) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *JobInfoExt) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *JobInfoExt) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedDate

`func (o *JobInfoExt) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *JobInfoExt) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *JobInfoExt) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *JobInfoExt) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetId

`func (o *JobInfoExt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobInfoExt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobInfoExt) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobInfoExt) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobType

`func (o *JobInfoExt) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *JobInfoExt) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *JobInfoExt) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *JobInfoExt) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetSystemModstamp

`func (o *JobInfoExt) GetSystemModstamp() string`

GetSystemModstamp returns the SystemModstamp field if non-nil, zero value otherwise.

### GetSystemModstampOk

`func (o *JobInfoExt) GetSystemModstampOk() (*string, bool)`

GetSystemModstampOk returns a tuple with the SystemModstamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModstamp

`func (o *JobInfoExt) SetSystemModstamp(v string)`

SetSystemModstamp sets SystemModstamp field to given value.

### HasSystemModstamp

`func (o *JobInfoExt) HasSystemModstamp() bool`

HasSystemModstamp returns a boolean if a field has been set.

### GetState

`func (o *JobInfoExt) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JobInfoExt) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JobInfoExt) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *JobInfoExt) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApexProcessingTime

`func (o *JobInfoExt) GetApexProcessingTime() int64`

GetApexProcessingTime returns the ApexProcessingTime field if non-nil, zero value otherwise.

### GetApexProcessingTimeOk

`func (o *JobInfoExt) GetApexProcessingTimeOk() (*int64, bool)`

GetApexProcessingTimeOk returns a tuple with the ApexProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApexProcessingTime

`func (o *JobInfoExt) SetApexProcessingTime(v int64)`

SetApexProcessingTime sets ApexProcessingTime field to given value.

### HasApexProcessingTime

`func (o *JobInfoExt) HasApexProcessingTime() bool`

HasApexProcessingTime returns a boolean if a field has been set.

### GetApiActiveProcessingTime

`func (o *JobInfoExt) GetApiActiveProcessingTime() int64`

GetApiActiveProcessingTime returns the ApiActiveProcessingTime field if non-nil, zero value otherwise.

### GetApiActiveProcessingTimeOk

`func (o *JobInfoExt) GetApiActiveProcessingTimeOk() (*int64, bool)`

GetApiActiveProcessingTimeOk returns a tuple with the ApiActiveProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiActiveProcessingTime

`func (o *JobInfoExt) SetApiActiveProcessingTime(v int64)`

SetApiActiveProcessingTime sets ApiActiveProcessingTime field to given value.

### HasApiActiveProcessingTime

`func (o *JobInfoExt) HasApiActiveProcessingTime() bool`

HasApiActiveProcessingTime returns a boolean if a field has been set.

### GetErrorMessage

`func (o *JobInfoExt) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *JobInfoExt) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *JobInfoExt) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *JobInfoExt) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetNumberRecordsFailed

`func (o *JobInfoExt) GetNumberRecordsFailed() int64`

GetNumberRecordsFailed returns the NumberRecordsFailed field if non-nil, zero value otherwise.

### GetNumberRecordsFailedOk

`func (o *JobInfoExt) GetNumberRecordsFailedOk() (*int64, bool)`

GetNumberRecordsFailedOk returns a tuple with the NumberRecordsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberRecordsFailed

`func (o *JobInfoExt) SetNumberRecordsFailed(v int64)`

SetNumberRecordsFailed sets NumberRecordsFailed field to given value.

### HasNumberRecordsFailed

`func (o *JobInfoExt) HasNumberRecordsFailed() bool`

HasNumberRecordsFailed returns a boolean if a field has been set.

### GetNumberRecordsProcessed

`func (o *JobInfoExt) GetNumberRecordsProcessed() int64`

GetNumberRecordsProcessed returns the NumberRecordsProcessed field if non-nil, zero value otherwise.

### GetNumberRecordsProcessedOk

`func (o *JobInfoExt) GetNumberRecordsProcessedOk() (*int64, bool)`

GetNumberRecordsProcessedOk returns a tuple with the NumberRecordsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberRecordsProcessed

`func (o *JobInfoExt) SetNumberRecordsProcessed(v int64)`

SetNumberRecordsProcessed sets NumberRecordsProcessed field to given value.

### HasNumberRecordsProcessed

`func (o *JobInfoExt) HasNumberRecordsProcessed() bool`

HasNumberRecordsProcessed returns a boolean if a field has been set.

### GetRetries

`func (o *JobInfoExt) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *JobInfoExt) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *JobInfoExt) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *JobInfoExt) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetTotalProcessingTime

`func (o *JobInfoExt) GetTotalProcessingTime() int64`

GetTotalProcessingTime returns the TotalProcessingTime field if non-nil, zero value otherwise.

### GetTotalProcessingTimeOk

`func (o *JobInfoExt) GetTotalProcessingTimeOk() (*int64, bool)`

GetTotalProcessingTimeOk returns a tuple with the TotalProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProcessingTime

`func (o *JobInfoExt) SetTotalProcessingTime(v int64)`

SetTotalProcessingTime sets TotalProcessingTime field to given value.

### HasTotalProcessingTime

`func (o *JobInfoExt) HasTotalProcessingTime() bool`

HasTotalProcessingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


