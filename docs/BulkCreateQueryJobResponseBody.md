# BulkCreateQueryJobResponseBody

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
**LineEnding** | Pointer to **string** |  | [optional] 
**ColumnDelimiter** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkCreateQueryJobResponseBody

`func NewBulkCreateQueryJobResponseBody() *BulkCreateQueryJobResponseBody`

NewBulkCreateQueryJobResponseBody instantiates a new BulkCreateQueryJobResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCreateQueryJobResponseBodyWithDefaults

`func NewBulkCreateQueryJobResponseBodyWithDefaults() *BulkCreateQueryJobResponseBody`

NewBulkCreateQueryJobResponseBodyWithDefaults instantiates a new BulkCreateQueryJobResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkCreateQueryJobResponseBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkCreateQueryJobResponseBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkCreateQueryJobResponseBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkCreateQueryJobResponseBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOperation

`func (o *BulkCreateQueryJobResponseBody) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BulkCreateQueryJobResponseBody) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BulkCreateQueryJobResponseBody) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *BulkCreateQueryJobResponseBody) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetObject

`func (o *BulkCreateQueryJobResponseBody) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BulkCreateQueryJobResponseBody) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BulkCreateQueryJobResponseBody) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *BulkCreateQueryJobResponseBody) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedById

`func (o *BulkCreateQueryJobResponseBody) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *BulkCreateQueryJobResponseBody) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *BulkCreateQueryJobResponseBody) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *BulkCreateQueryJobResponseBody) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedDate

`func (o *BulkCreateQueryJobResponseBody) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *BulkCreateQueryJobResponseBody) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *BulkCreateQueryJobResponseBody) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *BulkCreateQueryJobResponseBody) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetSystemModstamp

`func (o *BulkCreateQueryJobResponseBody) GetSystemModstamp() string`

GetSystemModstamp returns the SystemModstamp field if non-nil, zero value otherwise.

### GetSystemModstampOk

`func (o *BulkCreateQueryJobResponseBody) GetSystemModstampOk() (*string, bool)`

GetSystemModstampOk returns a tuple with the SystemModstamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModstamp

`func (o *BulkCreateQueryJobResponseBody) SetSystemModstamp(v string)`

SetSystemModstamp sets SystemModstamp field to given value.

### HasSystemModstamp

`func (o *BulkCreateQueryJobResponseBody) HasSystemModstamp() bool`

HasSystemModstamp returns a boolean if a field has been set.

### GetState

`func (o *BulkCreateQueryJobResponseBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BulkCreateQueryJobResponseBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BulkCreateQueryJobResponseBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BulkCreateQueryJobResponseBody) HasState() bool`

HasState returns a boolean if a field has been set.

### GetConcurrencyMode

`func (o *BulkCreateQueryJobResponseBody) GetConcurrencyMode() string`

GetConcurrencyMode returns the ConcurrencyMode field if non-nil, zero value otherwise.

### GetConcurrencyModeOk

`func (o *BulkCreateQueryJobResponseBody) GetConcurrencyModeOk() (*string, bool)`

GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyMode

`func (o *BulkCreateQueryJobResponseBody) SetConcurrencyMode(v string)`

SetConcurrencyMode sets ConcurrencyMode field to given value.

### HasConcurrencyMode

`func (o *BulkCreateQueryJobResponseBody) HasConcurrencyMode() bool`

HasConcurrencyMode returns a boolean if a field has been set.

### GetContentType

`func (o *BulkCreateQueryJobResponseBody) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BulkCreateQueryJobResponseBody) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BulkCreateQueryJobResponseBody) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *BulkCreateQueryJobResponseBody) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetApiVersion

`func (o *BulkCreateQueryJobResponseBody) GetApiVersion() float32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BulkCreateQueryJobResponseBody) GetApiVersionOk() (*float32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BulkCreateQueryJobResponseBody) SetApiVersion(v float32)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BulkCreateQueryJobResponseBody) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetLineEnding

`func (o *BulkCreateQueryJobResponseBody) GetLineEnding() string`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *BulkCreateQueryJobResponseBody) GetLineEndingOk() (*string, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *BulkCreateQueryJobResponseBody) SetLineEnding(v string)`

SetLineEnding sets LineEnding field to given value.

### HasLineEnding

`func (o *BulkCreateQueryJobResponseBody) HasLineEnding() bool`

HasLineEnding returns a boolean if a field has been set.

### GetColumnDelimiter

`func (o *BulkCreateQueryJobResponseBody) GetColumnDelimiter() string`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *BulkCreateQueryJobResponseBody) GetColumnDelimiterOk() (*string, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *BulkCreateQueryJobResponseBody) SetColumnDelimiter(v string)`

SetColumnDelimiter sets ColumnDelimiter field to given value.

### HasColumnDelimiter

`func (o *BulkCreateQueryJobResponseBody) HasColumnDelimiter() bool`

HasColumnDelimiter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


