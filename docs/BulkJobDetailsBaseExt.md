# BulkJobDetailsBaseExt

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

## Methods

### NewBulkJobDetailsBaseExt

`func NewBulkJobDetailsBaseExt() *BulkJobDetailsBaseExt`

NewBulkJobDetailsBaseExt instantiates a new BulkJobDetailsBaseExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkJobDetailsBaseExtWithDefaults

`func NewBulkJobDetailsBaseExtWithDefaults() *BulkJobDetailsBaseExt`

NewBulkJobDetailsBaseExtWithDefaults instantiates a new BulkJobDetailsBaseExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *BulkJobDetailsBaseExt) GetApiVersion() float32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BulkJobDetailsBaseExt) GetApiVersionOk() (*float32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BulkJobDetailsBaseExt) SetApiVersion(v float32)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BulkJobDetailsBaseExt) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetConcurrencyMode

`func (o *BulkJobDetailsBaseExt) GetConcurrencyMode() string`

GetConcurrencyMode returns the ConcurrencyMode field if non-nil, zero value otherwise.

### GetConcurrencyModeOk

`func (o *BulkJobDetailsBaseExt) GetConcurrencyModeOk() (*string, bool)`

GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyMode

`func (o *BulkJobDetailsBaseExt) SetConcurrencyMode(v string)`

SetConcurrencyMode sets ConcurrencyMode field to given value.

### HasConcurrencyMode

`func (o *BulkJobDetailsBaseExt) HasConcurrencyMode() bool`

HasConcurrencyMode returns a boolean if a field has been set.

### GetContentUrl

`func (o *BulkJobDetailsBaseExt) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *BulkJobDetailsBaseExt) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *BulkJobDetailsBaseExt) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.

### HasContentUrl

`func (o *BulkJobDetailsBaseExt) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### GetCreatedById

`func (o *BulkJobDetailsBaseExt) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *BulkJobDetailsBaseExt) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *BulkJobDetailsBaseExt) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *BulkJobDetailsBaseExt) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedDate

`func (o *BulkJobDetailsBaseExt) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *BulkJobDetailsBaseExt) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *BulkJobDetailsBaseExt) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *BulkJobDetailsBaseExt) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetId

`func (o *BulkJobDetailsBaseExt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkJobDetailsBaseExt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkJobDetailsBaseExt) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkJobDetailsBaseExt) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobType

`func (o *BulkJobDetailsBaseExt) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BulkJobDetailsBaseExt) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BulkJobDetailsBaseExt) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *BulkJobDetailsBaseExt) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetSystemModstamp

`func (o *BulkJobDetailsBaseExt) GetSystemModstamp() string`

GetSystemModstamp returns the SystemModstamp field if non-nil, zero value otherwise.

### GetSystemModstampOk

`func (o *BulkJobDetailsBaseExt) GetSystemModstampOk() (*string, bool)`

GetSystemModstampOk returns a tuple with the SystemModstamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModstamp

`func (o *BulkJobDetailsBaseExt) SetSystemModstamp(v string)`

SetSystemModstamp sets SystemModstamp field to given value.

### HasSystemModstamp

`func (o *BulkJobDetailsBaseExt) HasSystemModstamp() bool`

HasSystemModstamp returns a boolean if a field has been set.

### GetState

`func (o *BulkJobDetailsBaseExt) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BulkJobDetailsBaseExt) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BulkJobDetailsBaseExt) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BulkJobDetailsBaseExt) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


