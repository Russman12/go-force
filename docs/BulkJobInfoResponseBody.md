# BulkJobInfoResponseBody

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
**ApiVersion** | Pointer to **int64** |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**ContentUrl** | Pointer to **string** |  | [optional] 
**LineEnding** | Pointer to **string** |  | [optional] 
**ColumnDelimiter** | Pointer to **string** |  | [optional] 
**Retries** | Pointer to **int32** |  | [optional] 
**TotalProcessingTime** | Pointer to **int32** |  | [optional] 
**ApiActiveProcessingTime** | Pointer to **int32** |  | [optional] 
**ApexProcessingTime** | Pointer to **int32** |  | [optional] 
**NumberRecordsProcessed** | Pointer to **int64** |  | [optional] 

## Methods

### NewBulkJobInfoResponseBody

`func NewBulkJobInfoResponseBody() *BulkJobInfoResponseBody`

NewBulkJobInfoResponseBody instantiates a new BulkJobInfoResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkJobInfoResponseBodyWithDefaults

`func NewBulkJobInfoResponseBodyWithDefaults() *BulkJobInfoResponseBody`

NewBulkJobInfoResponseBodyWithDefaults instantiates a new BulkJobInfoResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkJobInfoResponseBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkJobInfoResponseBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkJobInfoResponseBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkJobInfoResponseBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOperation

`func (o *BulkJobInfoResponseBody) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BulkJobInfoResponseBody) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BulkJobInfoResponseBody) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *BulkJobInfoResponseBody) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetObject

`func (o *BulkJobInfoResponseBody) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BulkJobInfoResponseBody) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BulkJobInfoResponseBody) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *BulkJobInfoResponseBody) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedById

`func (o *BulkJobInfoResponseBody) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *BulkJobInfoResponseBody) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *BulkJobInfoResponseBody) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *BulkJobInfoResponseBody) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedDate

`func (o *BulkJobInfoResponseBody) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *BulkJobInfoResponseBody) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *BulkJobInfoResponseBody) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *BulkJobInfoResponseBody) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetSystemModstamp

`func (o *BulkJobInfoResponseBody) GetSystemModstamp() string`

GetSystemModstamp returns the SystemModstamp field if non-nil, zero value otherwise.

### GetSystemModstampOk

`func (o *BulkJobInfoResponseBody) GetSystemModstampOk() (*string, bool)`

GetSystemModstampOk returns a tuple with the SystemModstamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModstamp

`func (o *BulkJobInfoResponseBody) SetSystemModstamp(v string)`

SetSystemModstamp sets SystemModstamp field to given value.

### HasSystemModstamp

`func (o *BulkJobInfoResponseBody) HasSystemModstamp() bool`

HasSystemModstamp returns a boolean if a field has been set.

### GetState

`func (o *BulkJobInfoResponseBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BulkJobInfoResponseBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BulkJobInfoResponseBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BulkJobInfoResponseBody) HasState() bool`

HasState returns a boolean if a field has been set.

### GetConcurrencyMode

`func (o *BulkJobInfoResponseBody) GetConcurrencyMode() string`

GetConcurrencyMode returns the ConcurrencyMode field if non-nil, zero value otherwise.

### GetConcurrencyModeOk

`func (o *BulkJobInfoResponseBody) GetConcurrencyModeOk() (*string, bool)`

GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyMode

`func (o *BulkJobInfoResponseBody) SetConcurrencyMode(v string)`

SetConcurrencyMode sets ConcurrencyMode field to given value.

### HasConcurrencyMode

`func (o *BulkJobInfoResponseBody) HasConcurrencyMode() bool`

HasConcurrencyMode returns a boolean if a field has been set.

### GetContentType

`func (o *BulkJobInfoResponseBody) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BulkJobInfoResponseBody) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BulkJobInfoResponseBody) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *BulkJobInfoResponseBody) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetApiVersion

`func (o *BulkJobInfoResponseBody) GetApiVersion() int64`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BulkJobInfoResponseBody) GetApiVersionOk() (*int64, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BulkJobInfoResponseBody) SetApiVersion(v int64)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BulkJobInfoResponseBody) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetJobType

`func (o *BulkJobInfoResponseBody) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BulkJobInfoResponseBody) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BulkJobInfoResponseBody) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *BulkJobInfoResponseBody) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetContentUrl

`func (o *BulkJobInfoResponseBody) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *BulkJobInfoResponseBody) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *BulkJobInfoResponseBody) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.

### HasContentUrl

`func (o *BulkJobInfoResponseBody) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### GetLineEnding

`func (o *BulkJobInfoResponseBody) GetLineEnding() string`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *BulkJobInfoResponseBody) GetLineEndingOk() (*string, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *BulkJobInfoResponseBody) SetLineEnding(v string)`

SetLineEnding sets LineEnding field to given value.

### HasLineEnding

`func (o *BulkJobInfoResponseBody) HasLineEnding() bool`

HasLineEnding returns a boolean if a field has been set.

### GetColumnDelimiter

`func (o *BulkJobInfoResponseBody) GetColumnDelimiter() string`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *BulkJobInfoResponseBody) GetColumnDelimiterOk() (*string, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *BulkJobInfoResponseBody) SetColumnDelimiter(v string)`

SetColumnDelimiter sets ColumnDelimiter field to given value.

### HasColumnDelimiter

`func (o *BulkJobInfoResponseBody) HasColumnDelimiter() bool`

HasColumnDelimiter returns a boolean if a field has been set.

### GetRetries

`func (o *BulkJobInfoResponseBody) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *BulkJobInfoResponseBody) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *BulkJobInfoResponseBody) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *BulkJobInfoResponseBody) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetTotalProcessingTime

`func (o *BulkJobInfoResponseBody) GetTotalProcessingTime() int32`

GetTotalProcessingTime returns the TotalProcessingTime field if non-nil, zero value otherwise.

### GetTotalProcessingTimeOk

`func (o *BulkJobInfoResponseBody) GetTotalProcessingTimeOk() (*int32, bool)`

GetTotalProcessingTimeOk returns a tuple with the TotalProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProcessingTime

`func (o *BulkJobInfoResponseBody) SetTotalProcessingTime(v int32)`

SetTotalProcessingTime sets TotalProcessingTime field to given value.

### HasTotalProcessingTime

`func (o *BulkJobInfoResponseBody) HasTotalProcessingTime() bool`

HasTotalProcessingTime returns a boolean if a field has been set.

### GetApiActiveProcessingTime

`func (o *BulkJobInfoResponseBody) GetApiActiveProcessingTime() int32`

GetApiActiveProcessingTime returns the ApiActiveProcessingTime field if non-nil, zero value otherwise.

### GetApiActiveProcessingTimeOk

`func (o *BulkJobInfoResponseBody) GetApiActiveProcessingTimeOk() (*int32, bool)`

GetApiActiveProcessingTimeOk returns a tuple with the ApiActiveProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiActiveProcessingTime

`func (o *BulkJobInfoResponseBody) SetApiActiveProcessingTime(v int32)`

SetApiActiveProcessingTime sets ApiActiveProcessingTime field to given value.

### HasApiActiveProcessingTime

`func (o *BulkJobInfoResponseBody) HasApiActiveProcessingTime() bool`

HasApiActiveProcessingTime returns a boolean if a field has been set.

### GetApexProcessingTime

`func (o *BulkJobInfoResponseBody) GetApexProcessingTime() int32`

GetApexProcessingTime returns the ApexProcessingTime field if non-nil, zero value otherwise.

### GetApexProcessingTimeOk

`func (o *BulkJobInfoResponseBody) GetApexProcessingTimeOk() (*int32, bool)`

GetApexProcessingTimeOk returns a tuple with the ApexProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApexProcessingTime

`func (o *BulkJobInfoResponseBody) SetApexProcessingTime(v int32)`

SetApexProcessingTime sets ApexProcessingTime field to given value.

### HasApexProcessingTime

`func (o *BulkJobInfoResponseBody) HasApexProcessingTime() bool`

HasApexProcessingTime returns a boolean if a field has been set.

### GetNumberRecordsProcessed

`func (o *BulkJobInfoResponseBody) GetNumberRecordsProcessed() int64`

GetNumberRecordsProcessed returns the NumberRecordsProcessed field if non-nil, zero value otherwise.

### GetNumberRecordsProcessedOk

`func (o *BulkJobInfoResponseBody) GetNumberRecordsProcessedOk() (*int64, bool)`

GetNumberRecordsProcessedOk returns a tuple with the NumberRecordsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberRecordsProcessed

`func (o *BulkJobInfoResponseBody) SetNumberRecordsProcessed(v int64)`

SetNumberRecordsProcessed sets NumberRecordsProcessed field to given value.

### HasNumberRecordsProcessed

`func (o *BulkJobInfoResponseBody) HasNumberRecordsProcessed() bool`

HasNumberRecordsProcessed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


