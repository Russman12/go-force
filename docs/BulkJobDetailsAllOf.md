# BulkJobDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **float32** | The API version that the job was created in. | 
**ConcurrencyMode** | **string** | For future use. How the request was processed. Currently only parallel mode is supported. (When other modes are added, the mode will be chosen automatically by the API and will not be user configurable.) | 
**Id** | **string** | Unique ID for this job. | 
**JobType** | **string** | The job’s type. | 

## Methods

### NewBulkJobDetailsAllOf

`func NewBulkJobDetailsAllOf(apiVersion float32, concurrencyMode string, id string, jobType string, ) *BulkJobDetailsAllOf`

NewBulkJobDetailsAllOf instantiates a new BulkJobDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkJobDetailsAllOfWithDefaults

`func NewBulkJobDetailsAllOfWithDefaults() *BulkJobDetailsAllOf`

NewBulkJobDetailsAllOfWithDefaults instantiates a new BulkJobDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *BulkJobDetailsAllOf) GetApiVersion() float32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BulkJobDetailsAllOf) GetApiVersionOk() (*float32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BulkJobDetailsAllOf) SetApiVersion(v float32)`

SetApiVersion sets ApiVersion field to given value.


### GetConcurrencyMode

`func (o *BulkJobDetailsAllOf) GetConcurrencyMode() string`

GetConcurrencyMode returns the ConcurrencyMode field if non-nil, zero value otherwise.

### GetConcurrencyModeOk

`func (o *BulkJobDetailsAllOf) GetConcurrencyModeOk() (*string, bool)`

GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyMode

`func (o *BulkJobDetailsAllOf) SetConcurrencyMode(v string)`

SetConcurrencyMode sets ConcurrencyMode field to given value.


### GetId

`func (o *BulkJobDetailsAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkJobDetailsAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkJobDetailsAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetJobType

`func (o *BulkJobDetailsAllOf) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BulkJobDetailsAllOf) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BulkJobDetailsAllOf) SetJobType(v string)`

SetJobType sets JobType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


