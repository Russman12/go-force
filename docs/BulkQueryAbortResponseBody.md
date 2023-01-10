# BulkQueryAbortResponseBody

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
**ApiVersion** | Pointer to **int32** |  | [optional] 

## Methods

### NewBulkQueryAbortResponseBody

`func NewBulkQueryAbortResponseBody() *BulkQueryAbortResponseBody`

NewBulkQueryAbortResponseBody instantiates a new BulkQueryAbortResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkQueryAbortResponseBodyWithDefaults

`func NewBulkQueryAbortResponseBodyWithDefaults() *BulkQueryAbortResponseBody`

NewBulkQueryAbortResponseBodyWithDefaults instantiates a new BulkQueryAbortResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkQueryAbortResponseBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkQueryAbortResponseBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkQueryAbortResponseBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkQueryAbortResponseBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOperation

`func (o *BulkQueryAbortResponseBody) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BulkQueryAbortResponseBody) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BulkQueryAbortResponseBody) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *BulkQueryAbortResponseBody) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetObject

`func (o *BulkQueryAbortResponseBody) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BulkQueryAbortResponseBody) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BulkQueryAbortResponseBody) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *BulkQueryAbortResponseBody) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedById

`func (o *BulkQueryAbortResponseBody) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *BulkQueryAbortResponseBody) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *BulkQueryAbortResponseBody) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *BulkQueryAbortResponseBody) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedDate

`func (o *BulkQueryAbortResponseBody) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *BulkQueryAbortResponseBody) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *BulkQueryAbortResponseBody) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *BulkQueryAbortResponseBody) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetSystemModstamp

`func (o *BulkQueryAbortResponseBody) GetSystemModstamp() string`

GetSystemModstamp returns the SystemModstamp field if non-nil, zero value otherwise.

### GetSystemModstampOk

`func (o *BulkQueryAbortResponseBody) GetSystemModstampOk() (*string, bool)`

GetSystemModstampOk returns a tuple with the SystemModstamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModstamp

`func (o *BulkQueryAbortResponseBody) SetSystemModstamp(v string)`

SetSystemModstamp sets SystemModstamp field to given value.

### HasSystemModstamp

`func (o *BulkQueryAbortResponseBody) HasSystemModstamp() bool`

HasSystemModstamp returns a boolean if a field has been set.

### GetState

`func (o *BulkQueryAbortResponseBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BulkQueryAbortResponseBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BulkQueryAbortResponseBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BulkQueryAbortResponseBody) HasState() bool`

HasState returns a boolean if a field has been set.

### GetConcurrencyMode

`func (o *BulkQueryAbortResponseBody) GetConcurrencyMode() string`

GetConcurrencyMode returns the ConcurrencyMode field if non-nil, zero value otherwise.

### GetConcurrencyModeOk

`func (o *BulkQueryAbortResponseBody) GetConcurrencyModeOk() (*string, bool)`

GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyMode

`func (o *BulkQueryAbortResponseBody) SetConcurrencyMode(v string)`

SetConcurrencyMode sets ConcurrencyMode field to given value.

### HasConcurrencyMode

`func (o *BulkQueryAbortResponseBody) HasConcurrencyMode() bool`

HasConcurrencyMode returns a boolean if a field has been set.

### GetContentType

`func (o *BulkQueryAbortResponseBody) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BulkQueryAbortResponseBody) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BulkQueryAbortResponseBody) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *BulkQueryAbortResponseBody) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetApiVersion

`func (o *BulkQueryAbortResponseBody) GetApiVersion() int32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BulkQueryAbortResponseBody) GetApiVersionOk() (*int32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BulkQueryAbortResponseBody) SetApiVersion(v int32)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BulkQueryAbortResponseBody) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


