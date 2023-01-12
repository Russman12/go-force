# SObjectDescribes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encoding** | **string** |  | 
**MaxBatchSize** | **int32** |  | 
**Sobjects** | [**[]SObjectDescribe**](SObjectDescribe.md) |  | 

## Methods

### NewSObjectDescribes

`func NewSObjectDescribes(encoding string, maxBatchSize int32, sobjects []SObjectDescribe, ) *SObjectDescribes`

NewSObjectDescribes instantiates a new SObjectDescribes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSObjectDescribesWithDefaults

`func NewSObjectDescribesWithDefaults() *SObjectDescribes`

NewSObjectDescribesWithDefaults instantiates a new SObjectDescribes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncoding

`func (o *SObjectDescribes) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *SObjectDescribes) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *SObjectDescribes) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.


### GetMaxBatchSize

`func (o *SObjectDescribes) GetMaxBatchSize() int32`

GetMaxBatchSize returns the MaxBatchSize field if non-nil, zero value otherwise.

### GetMaxBatchSizeOk

`func (o *SObjectDescribes) GetMaxBatchSizeOk() (*int32, bool)`

GetMaxBatchSizeOk returns a tuple with the MaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBatchSize

`func (o *SObjectDescribes) SetMaxBatchSize(v int32)`

SetMaxBatchSize sets MaxBatchSize field to given value.


### GetSobjects

`func (o *SObjectDescribes) GetSobjects() []SObjectDescribe`

GetSobjects returns the Sobjects field if non-nil, zero value otherwise.

### GetSobjectsOk

`func (o *SObjectDescribes) GetSobjectsOk() (*[]SObjectDescribe, bool)`

GetSobjectsOk returns a tuple with the Sobjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSobjects

`func (o *SObjectDescribes) SetSobjects(v []SObjectDescribe)`

SetSobjects sets Sobjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


