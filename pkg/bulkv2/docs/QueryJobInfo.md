# QueryJobInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID for this job. | [optional] 
**Object** | Pointer to **string** | The object type being queried. | [optional] 
**CreatedById** | Pointer to **string** | The ID of the user who created the job. | [optional] 
**CreatedDate** | Pointer to **string** | The UTC date and time when the job was created. | [optional] 
**SystemModstamp** | Pointer to **string** | The UTC date and time when the API last updated the job information. | [optional] 
**State** | Pointer to [**QueryJobState**](QueryJobState.md) |  | [optional] 
**ConcurrencyMode** | Pointer to **string** | Reserved for future use. How the request is processed. Currently only parallel mode is supported. (When other modes are added, the API chooses the mode automatically. The mode isn’t user configurable.) | [optional] 
**ApiVersion** | Pointer to **float32** | The API version that the job was created in. | [optional] 
**JobType** | Pointer to [**QueryJobType**](QueryJobType.md) |  | [optional] 
**NumberRecordsProcessed** | Pointer to **int64** | The number of records processed in this job. | [optional] 
**Retries** | Pointer to **int32** | The number of times that Salesforce attempted to save the results of an operation. Repeated attempts indicate a problem such as a lock contention. | [optional] 
**TotalProcessingTime** | Pointer to **int64** | The number of milliseconds taken to process the job. | [optional] 
**Operation** | Pointer to [**QueryOperation**](QueryOperation.md) |  | [optional] 
**ContentType** | Pointer to [**ContentType**](ContentType.md) |  | [optional] 
**ColumnDelimiter** | Pointer to [**ColumnDelimiter**](ColumnDelimiter.md) |  | [optional] 
**LineEnding** | Pointer to [**LineEnding**](LineEnding.md) |  | [optional] 

## Methods

### NewQueryJobInfo

`func NewQueryJobInfo() *QueryJobInfo`

NewQueryJobInfo instantiates a new QueryJobInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryJobInfoWithDefaults

`func NewQueryJobInfoWithDefaults() *QueryJobInfo`

NewQueryJobInfoWithDefaults instantiates a new QueryJobInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QueryJobInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueryJobInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueryJobInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QueryJobInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *QueryJobInfo) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *QueryJobInfo) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *QueryJobInfo) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *QueryJobInfo) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedById

`func (o *QueryJobInfo) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *QueryJobInfo) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *QueryJobInfo) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *QueryJobInfo) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedDate

`func (o *QueryJobInfo) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *QueryJobInfo) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *QueryJobInfo) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *QueryJobInfo) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetSystemModstamp

`func (o *QueryJobInfo) GetSystemModstamp() string`

GetSystemModstamp returns the SystemModstamp field if non-nil, zero value otherwise.

### GetSystemModstampOk

`func (o *QueryJobInfo) GetSystemModstampOk() (*string, bool)`

GetSystemModstampOk returns a tuple with the SystemModstamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModstamp

`func (o *QueryJobInfo) SetSystemModstamp(v string)`

SetSystemModstamp sets SystemModstamp field to given value.

### HasSystemModstamp

`func (o *QueryJobInfo) HasSystemModstamp() bool`

HasSystemModstamp returns a boolean if a field has been set.

### GetState

`func (o *QueryJobInfo) GetState() QueryJobState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QueryJobInfo) GetStateOk() (*QueryJobState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QueryJobInfo) SetState(v QueryJobState)`

SetState sets State field to given value.

### HasState

`func (o *QueryJobInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetConcurrencyMode

`func (o *QueryJobInfo) GetConcurrencyMode() string`

GetConcurrencyMode returns the ConcurrencyMode field if non-nil, zero value otherwise.

### GetConcurrencyModeOk

`func (o *QueryJobInfo) GetConcurrencyModeOk() (*string, bool)`

GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyMode

`func (o *QueryJobInfo) SetConcurrencyMode(v string)`

SetConcurrencyMode sets ConcurrencyMode field to given value.

### HasConcurrencyMode

`func (o *QueryJobInfo) HasConcurrencyMode() bool`

HasConcurrencyMode returns a boolean if a field has been set.

### GetApiVersion

`func (o *QueryJobInfo) GetApiVersion() float32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *QueryJobInfo) GetApiVersionOk() (*float32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *QueryJobInfo) SetApiVersion(v float32)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *QueryJobInfo) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetJobType

`func (o *QueryJobInfo) GetJobType() QueryJobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *QueryJobInfo) GetJobTypeOk() (*QueryJobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *QueryJobInfo) SetJobType(v QueryJobType)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *QueryJobInfo) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetNumberRecordsProcessed

`func (o *QueryJobInfo) GetNumberRecordsProcessed() int64`

GetNumberRecordsProcessed returns the NumberRecordsProcessed field if non-nil, zero value otherwise.

### GetNumberRecordsProcessedOk

`func (o *QueryJobInfo) GetNumberRecordsProcessedOk() (*int64, bool)`

GetNumberRecordsProcessedOk returns a tuple with the NumberRecordsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberRecordsProcessed

`func (o *QueryJobInfo) SetNumberRecordsProcessed(v int64)`

SetNumberRecordsProcessed sets NumberRecordsProcessed field to given value.

### HasNumberRecordsProcessed

`func (o *QueryJobInfo) HasNumberRecordsProcessed() bool`

HasNumberRecordsProcessed returns a boolean if a field has been set.

### GetRetries

`func (o *QueryJobInfo) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *QueryJobInfo) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *QueryJobInfo) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *QueryJobInfo) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetTotalProcessingTime

`func (o *QueryJobInfo) GetTotalProcessingTime() int64`

GetTotalProcessingTime returns the TotalProcessingTime field if non-nil, zero value otherwise.

### GetTotalProcessingTimeOk

`func (o *QueryJobInfo) GetTotalProcessingTimeOk() (*int64, bool)`

GetTotalProcessingTimeOk returns a tuple with the TotalProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProcessingTime

`func (o *QueryJobInfo) SetTotalProcessingTime(v int64)`

SetTotalProcessingTime sets TotalProcessingTime field to given value.

### HasTotalProcessingTime

`func (o *QueryJobInfo) HasTotalProcessingTime() bool`

HasTotalProcessingTime returns a boolean if a field has been set.

### GetOperation

`func (o *QueryJobInfo) GetOperation() QueryOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *QueryJobInfo) GetOperationOk() (*QueryOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *QueryJobInfo) SetOperation(v QueryOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *QueryJobInfo) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetContentType

`func (o *QueryJobInfo) GetContentType() ContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *QueryJobInfo) GetContentTypeOk() (*ContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *QueryJobInfo) SetContentType(v ContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *QueryJobInfo) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetColumnDelimiter

`func (o *QueryJobInfo) GetColumnDelimiter() ColumnDelimiter`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *QueryJobInfo) GetColumnDelimiterOk() (*ColumnDelimiter, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *QueryJobInfo) SetColumnDelimiter(v ColumnDelimiter)`

SetColumnDelimiter sets ColumnDelimiter field to given value.

### HasColumnDelimiter

`func (o *QueryJobInfo) HasColumnDelimiter() bool`

HasColumnDelimiter returns a boolean if a field has been set.

### GetLineEnding

`func (o *QueryJobInfo) GetLineEnding() LineEnding`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *QueryJobInfo) GetLineEndingOk() (*LineEnding, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *QueryJobInfo) SetLineEnding(v LineEnding)`

SetLineEnding sets LineEnding field to given value.

### HasLineEnding

`func (o *QueryJobInfo) HasLineEnding() bool`

HasLineEnding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


