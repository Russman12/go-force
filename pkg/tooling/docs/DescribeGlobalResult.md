# DescribeGlobalResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encoding** | Pointer to **string** |  | [optional] 
**MaxBatchSize** | Pointer to **int32** |  | [optional] 
**Sobjects** | Pointer to [**[]SObjectDescribe**](SObjectDescribe.md) |  | [optional] 

## Methods

### NewDescribeGlobalResult

`func NewDescribeGlobalResult() *DescribeGlobalResult`

NewDescribeGlobalResult instantiates a new DescribeGlobalResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGlobalResultWithDefaults

`func NewDescribeGlobalResultWithDefaults() *DescribeGlobalResult`

NewDescribeGlobalResultWithDefaults instantiates a new DescribeGlobalResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncoding

`func (o *DescribeGlobalResult) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *DescribeGlobalResult) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *DescribeGlobalResult) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *DescribeGlobalResult) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetMaxBatchSize

`func (o *DescribeGlobalResult) GetMaxBatchSize() int32`

GetMaxBatchSize returns the MaxBatchSize field if non-nil, zero value otherwise.

### GetMaxBatchSizeOk

`func (o *DescribeGlobalResult) GetMaxBatchSizeOk() (*int32, bool)`

GetMaxBatchSizeOk returns a tuple with the MaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBatchSize

`func (o *DescribeGlobalResult) SetMaxBatchSize(v int32)`

SetMaxBatchSize sets MaxBatchSize field to given value.

### HasMaxBatchSize

`func (o *DescribeGlobalResult) HasMaxBatchSize() bool`

HasMaxBatchSize returns a boolean if a field has been set.

### GetSobjects

`func (o *DescribeGlobalResult) GetSobjects() []SObjectDescribe`

GetSobjects returns the Sobjects field if non-nil, zero value otherwise.

### GetSobjectsOk

`func (o *DescribeGlobalResult) GetSobjectsOk() (*[]SObjectDescribe, bool)`

GetSobjectsOk returns a tuple with the Sobjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSobjects

`func (o *DescribeGlobalResult) SetSobjects(v []SObjectDescribe)`

SetSobjects sets Sobjects field to given value.

### HasSobjects

`func (o *DescribeGlobalResult) HasSobjects() bool`

HasSobjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


