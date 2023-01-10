# BulkQueryJobInfoExt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID for this job. | [optional] 
**Object** | Pointer to **string** | The object type being queried. | [optional] 
**CreatedById** | Pointer to **string** | The ID of the user who created the job. | [optional] 
**CreatedDate** | Pointer to **string** | The UTC date and time when the job was created. | [optional] 
**SystemModstamp** | Pointer to **string** | The UTC date and time when the API last updated the job information. | [optional] 
**State** | Pointer to **string** | The current state of processing for the job. | [optional] 
**ConcurrencyMode** | Pointer to **string** | Reserved for future use. How the request is processed. Currently only parallel mode is supported. (When other modes are added, the API chooses the mode automatically. The mode isn’t user configurable.) | [optional] 
**ApiVersion** | Pointer to **float32** | The API version that the job was created in. | [optional] 
**JobType** | Pointer to **string** | The job’s type. For a query job, the type is always V2Query. | [optional] 
**NumberRecordsProcessed** | Pointer to **int64** | The number of records processed in this job. | [optional] 
**Retries** | Pointer to **int32** | The number of times that Salesforce attempted to save the results of an operation. Repeated attempts indicate a problem such as a lock contention. | [optional] 
**TotalProcessingTime** | Pointer to **int64** | The number of milliseconds taken to process the job. | [optional] 

## Methods

### NewBulkQueryJobInfoExt

`func NewBulkQueryJobInfoExt() *BulkQueryJobInfoExt`

NewBulkQueryJobInfoExt instantiates a new BulkQueryJobInfoExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkQueryJobInfoExtWithDefaults

`func NewBulkQueryJobInfoExtWithDefaults() *BulkQueryJobInfoExt`

NewBulkQueryJobInfoExtWithDefaults instantiates a new BulkQueryJobInfoExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkQueryJobInfoExt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkQueryJobInfoExt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkQueryJobInfoExt) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkQueryJobInfoExt) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *BulkQueryJobInfoExt) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BulkQueryJobInfoExt) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BulkQueryJobInfoExt) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *BulkQueryJobInfoExt) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedById

`func (o *BulkQueryJobInfoExt) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *BulkQueryJobInfoExt) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *BulkQueryJobInfoExt) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *BulkQueryJobInfoExt) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedDate

`func (o *BulkQueryJobInfoExt) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *BulkQueryJobInfoExt) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *BulkQueryJobInfoExt) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *BulkQueryJobInfoExt) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetSystemModstamp

`func (o *BulkQueryJobInfoExt) GetSystemModstamp() string`

GetSystemModstamp returns the SystemModstamp field if non-nil, zero value otherwise.

### GetSystemModstampOk

`func (o *BulkQueryJobInfoExt) GetSystemModstampOk() (*string, bool)`

GetSystemModstampOk returns a tuple with the SystemModstamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModstamp

`func (o *BulkQueryJobInfoExt) SetSystemModstamp(v string)`

SetSystemModstamp sets SystemModstamp field to given value.

### HasSystemModstamp

`func (o *BulkQueryJobInfoExt) HasSystemModstamp() bool`

HasSystemModstamp returns a boolean if a field has been set.

### GetState

`func (o *BulkQueryJobInfoExt) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BulkQueryJobInfoExt) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BulkQueryJobInfoExt) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BulkQueryJobInfoExt) HasState() bool`

HasState returns a boolean if a field has been set.

### GetConcurrencyMode

`func (o *BulkQueryJobInfoExt) GetConcurrencyMode() string`

GetConcurrencyMode returns the ConcurrencyMode field if non-nil, zero value otherwise.

### GetConcurrencyModeOk

`func (o *BulkQueryJobInfoExt) GetConcurrencyModeOk() (*string, bool)`

GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyMode

`func (o *BulkQueryJobInfoExt) SetConcurrencyMode(v string)`

SetConcurrencyMode sets ConcurrencyMode field to given value.

### HasConcurrencyMode

`func (o *BulkQueryJobInfoExt) HasConcurrencyMode() bool`

HasConcurrencyMode returns a boolean if a field has been set.

### GetApiVersion

`func (o *BulkQueryJobInfoExt) GetApiVersion() float32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BulkQueryJobInfoExt) GetApiVersionOk() (*float32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BulkQueryJobInfoExt) SetApiVersion(v float32)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BulkQueryJobInfoExt) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetJobType

`func (o *BulkQueryJobInfoExt) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BulkQueryJobInfoExt) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BulkQueryJobInfoExt) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *BulkQueryJobInfoExt) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetNumberRecordsProcessed

`func (o *BulkQueryJobInfoExt) GetNumberRecordsProcessed() int64`

GetNumberRecordsProcessed returns the NumberRecordsProcessed field if non-nil, zero value otherwise.

### GetNumberRecordsProcessedOk

`func (o *BulkQueryJobInfoExt) GetNumberRecordsProcessedOk() (*int64, bool)`

GetNumberRecordsProcessedOk returns a tuple with the NumberRecordsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberRecordsProcessed

`func (o *BulkQueryJobInfoExt) SetNumberRecordsProcessed(v int64)`

SetNumberRecordsProcessed sets NumberRecordsProcessed field to given value.

### HasNumberRecordsProcessed

`func (o *BulkQueryJobInfoExt) HasNumberRecordsProcessed() bool`

HasNumberRecordsProcessed returns a boolean if a field has been set.

### GetRetries

`func (o *BulkQueryJobInfoExt) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *BulkQueryJobInfoExt) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *BulkQueryJobInfoExt) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *BulkQueryJobInfoExt) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetTotalProcessingTime

`func (o *BulkQueryJobInfoExt) GetTotalProcessingTime() int64`

GetTotalProcessingTime returns the TotalProcessingTime field if non-nil, zero value otherwise.

### GetTotalProcessingTimeOk

`func (o *BulkQueryJobInfoExt) GetTotalProcessingTimeOk() (*int64, bool)`

GetTotalProcessingTimeOk returns a tuple with the TotalProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProcessingTime

`func (o *BulkQueryJobInfoExt) SetTotalProcessingTime(v int64)`

SetTotalProcessingTime sets TotalProcessingTime field to given value.

### HasTotalProcessingTime

`func (o *BulkQueryJobInfoExt) HasTotalProcessingTime() bool`

HasTotalProcessingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


