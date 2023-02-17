# QueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int32** |  | 
**TotalSize** | **int32** |  | 
**Done** | **bool** |  | 
**QueryLocator** | **NullableString** |  | 
**EntityTypeName** | **string** |  | 
**Records** | **[]map[string]interface{}** |  | 

## Methods

### NewQueryResult

`func NewQueryResult(size int32, totalSize int32, done bool, queryLocator NullableString, entityTypeName string, records []map[string]interface{}, ) *QueryResult`

NewQueryResult instantiates a new QueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResultWithDefaults

`func NewQueryResultWithDefaults() *QueryResult`

NewQueryResultWithDefaults instantiates a new QueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *QueryResult) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *QueryResult) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *QueryResult) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTotalSize

`func (o *QueryResult) GetTotalSize() int32`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *QueryResult) GetTotalSizeOk() (*int32, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *QueryResult) SetTotalSize(v int32)`

SetTotalSize sets TotalSize field to given value.


### GetDone

`func (o *QueryResult) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *QueryResult) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *QueryResult) SetDone(v bool)`

SetDone sets Done field to given value.


### GetQueryLocator

`func (o *QueryResult) GetQueryLocator() string`

GetQueryLocator returns the QueryLocator field if non-nil, zero value otherwise.

### GetQueryLocatorOk

`func (o *QueryResult) GetQueryLocatorOk() (*string, bool)`

GetQueryLocatorOk returns a tuple with the QueryLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryLocator

`func (o *QueryResult) SetQueryLocator(v string)`

SetQueryLocator sets QueryLocator field to given value.


### SetQueryLocatorNil

`func (o *QueryResult) SetQueryLocatorNil(b bool)`

 SetQueryLocatorNil sets the value for QueryLocator to be an explicit nil

### UnsetQueryLocator
`func (o *QueryResult) UnsetQueryLocator()`

UnsetQueryLocator ensures that no value is present for QueryLocator, not even an explicit nil
### GetEntityTypeName

`func (o *QueryResult) GetEntityTypeName() string`

GetEntityTypeName returns the EntityTypeName field if non-nil, zero value otherwise.

### GetEntityTypeNameOk

`func (o *QueryResult) GetEntityTypeNameOk() (*string, bool)`

GetEntityTypeNameOk returns a tuple with the EntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeName

`func (o *QueryResult) SetEntityTypeName(v string)`

SetEntityTypeName sets EntityTypeName field to given value.


### GetRecords

`func (o *QueryResult) GetRecords() []map[string]interface{}`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *QueryResult) GetRecordsOk() (*[]map[string]interface{}, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *QueryResult) SetRecords(v []map[string]interface{})`

SetRecords sets Records field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


