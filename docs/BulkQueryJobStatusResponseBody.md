# BulkQueryJobStatusResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **string** |  | [optional] 
**SystemModstamp** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**ConcurrencyMode** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **float32** |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**LineEnding** | Pointer to **string** |  | [optional] 
**ColumnDelimiter** | Pointer to **string** |  | [optional] 
**NumberRecordsProcessed** | Pointer to **int32** |  | [optional] 
**Retries** | Pointer to **int32** |  | [optional] 
**TotalProcessingTime** | Pointer to **int32** |  | [optional] 

## Methods

### NewBulkQueryJobStatusResponseBody

`func NewBulkQueryJobStatusResponseBody() *BulkQueryJobStatusResponseBody`

NewBulkQueryJobStatusResponseBody instantiates a new BulkQueryJobStatusResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkQueryJobStatusResponseBodyWithDefaults

`func NewBulkQueryJobStatusResponseBodyWithDefaults() *BulkQueryJobStatusResponseBody`

NewBulkQueryJobStatusResponseBodyWithDefaults instantiates a new BulkQueryJobStatusResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkQueryJobStatusResponseBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkQueryJobStatusResponseBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkQueryJobStatusResponseBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkQueryJobStatusResponseBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOperation

`func (o *BulkQueryJobStatusResponseBody) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BulkQueryJobStatusResponseBody) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BulkQueryJobStatusResponseBody) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *BulkQueryJobStatusResponseBody) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetObject

`func (o *BulkQueryJobStatusResponseBody) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BulkQueryJobStatusResponseBody) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BulkQueryJobStatusResponseBody) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *BulkQueryJobStatusResponseBody) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedById

`func (o *BulkQueryJobStatusResponseBody) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *BulkQueryJobStatusResponseBody) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *BulkQueryJobStatusResponseBody) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *BulkQueryJobStatusResponseBody) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedDate

`func (o *BulkQueryJobStatusResponseBody) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *BulkQueryJobStatusResponseBody) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *BulkQueryJobStatusResponseBody) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *BulkQueryJobStatusResponseBody) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetSystemModstamp

`func (o *BulkQueryJobStatusResponseBody) GetSystemModstamp() string`

GetSystemModstamp returns the SystemModstamp field if non-nil, zero value otherwise.

### GetSystemModstampOk

`func (o *BulkQueryJobStatusResponseBody) GetSystemModstampOk() (*string, bool)`

GetSystemModstampOk returns a tuple with the SystemModstamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModstamp

`func (o *BulkQueryJobStatusResponseBody) SetSystemModstamp(v string)`

SetSystemModstamp sets SystemModstamp field to given value.

### HasSystemModstamp

`func (o *BulkQueryJobStatusResponseBody) HasSystemModstamp() bool`

HasSystemModstamp returns a boolean if a field has been set.

### GetState

`func (o *BulkQueryJobStatusResponseBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BulkQueryJobStatusResponseBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BulkQueryJobStatusResponseBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BulkQueryJobStatusResponseBody) HasState() bool`

HasState returns a boolean if a field has been set.

### GetConcurrencyMode

`func (o *BulkQueryJobStatusResponseBody) GetConcurrencyMode() string`

GetConcurrencyMode returns the ConcurrencyMode field if non-nil, zero value otherwise.

### GetConcurrencyModeOk

`func (o *BulkQueryJobStatusResponseBody) GetConcurrencyModeOk() (*string, bool)`

GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyMode

`func (o *BulkQueryJobStatusResponseBody) SetConcurrencyMode(v string)`

SetConcurrencyMode sets ConcurrencyMode field to given value.

### HasConcurrencyMode

`func (o *BulkQueryJobStatusResponseBody) HasConcurrencyMode() bool`

HasConcurrencyMode returns a boolean if a field has been set.

### GetContentType

`func (o *BulkQueryJobStatusResponseBody) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BulkQueryJobStatusResponseBody) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BulkQueryJobStatusResponseBody) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *BulkQueryJobStatusResponseBody) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetApiVersion

`func (o *BulkQueryJobStatusResponseBody) GetApiVersion() float32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BulkQueryJobStatusResponseBody) GetApiVersionOk() (*float32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BulkQueryJobStatusResponseBody) SetApiVersion(v float32)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BulkQueryJobStatusResponseBody) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetJobType

`func (o *BulkQueryJobStatusResponseBody) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BulkQueryJobStatusResponseBody) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BulkQueryJobStatusResponseBody) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *BulkQueryJobStatusResponseBody) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetLineEnding

`func (o *BulkQueryJobStatusResponseBody) GetLineEnding() string`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *BulkQueryJobStatusResponseBody) GetLineEndingOk() (*string, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *BulkQueryJobStatusResponseBody) SetLineEnding(v string)`

SetLineEnding sets LineEnding field to given value.

### HasLineEnding

`func (o *BulkQueryJobStatusResponseBody) HasLineEnding() bool`

HasLineEnding returns a boolean if a field has been set.

### GetColumnDelimiter

`func (o *BulkQueryJobStatusResponseBody) GetColumnDelimiter() string`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *BulkQueryJobStatusResponseBody) GetColumnDelimiterOk() (*string, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *BulkQueryJobStatusResponseBody) SetColumnDelimiter(v string)`

SetColumnDelimiter sets ColumnDelimiter field to given value.

### HasColumnDelimiter

`func (o *BulkQueryJobStatusResponseBody) HasColumnDelimiter() bool`

HasColumnDelimiter returns a boolean if a field has been set.

### GetNumberRecordsProcessed

`func (o *BulkQueryJobStatusResponseBody) GetNumberRecordsProcessed() int32`

GetNumberRecordsProcessed returns the NumberRecordsProcessed field if non-nil, zero value otherwise.

### GetNumberRecordsProcessedOk

`func (o *BulkQueryJobStatusResponseBody) GetNumberRecordsProcessedOk() (*int32, bool)`

GetNumberRecordsProcessedOk returns a tuple with the NumberRecordsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberRecordsProcessed

`func (o *BulkQueryJobStatusResponseBody) SetNumberRecordsProcessed(v int32)`

SetNumberRecordsProcessed sets NumberRecordsProcessed field to given value.

### HasNumberRecordsProcessed

`func (o *BulkQueryJobStatusResponseBody) HasNumberRecordsProcessed() bool`

HasNumberRecordsProcessed returns a boolean if a field has been set.

### GetRetries

`func (o *BulkQueryJobStatusResponseBody) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *BulkQueryJobStatusResponseBody) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *BulkQueryJobStatusResponseBody) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *BulkQueryJobStatusResponseBody) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetTotalProcessingTime

`func (o *BulkQueryJobStatusResponseBody) GetTotalProcessingTime() int32`

GetTotalProcessingTime returns the TotalProcessingTime field if non-nil, zero value otherwise.

### GetTotalProcessingTimeOk

`func (o *BulkQueryJobStatusResponseBody) GetTotalProcessingTimeOk() (*int32, bool)`

GetTotalProcessingTimeOk returns a tuple with the TotalProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProcessingTime

`func (o *BulkQueryJobStatusResponseBody) SetTotalProcessingTime(v int32)`

SetTotalProcessingTime sets TotalProcessingTime field to given value.

### HasTotalProcessingTime

`func (o *BulkQueryJobStatusResponseBody) HasTotalProcessingTime() bool`

HasTotalProcessingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


