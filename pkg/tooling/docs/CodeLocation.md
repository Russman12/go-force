# CodeLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | **int32** |  | 
**Line** | **int32** |  | 
**NumExecutions** | **int32** |  | 
**Time** | **float64** |  | 

## Methods

### NewCodeLocation

`func NewCodeLocation(column int32, line int32, numExecutions int32, time float64, ) *CodeLocation`

NewCodeLocation instantiates a new CodeLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeLocationWithDefaults

`func NewCodeLocationWithDefaults() *CodeLocation`

NewCodeLocationWithDefaults instantiates a new CodeLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *CodeLocation) GetColumn() int32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *CodeLocation) GetColumnOk() (*int32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *CodeLocation) SetColumn(v int32)`

SetColumn sets Column field to given value.


### GetLine

`func (o *CodeLocation) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *CodeLocation) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *CodeLocation) SetLine(v int32)`

SetLine sets Line field to given value.


### GetNumExecutions

`func (o *CodeLocation) GetNumExecutions() int32`

GetNumExecutions returns the NumExecutions field if non-nil, zero value otherwise.

### GetNumExecutionsOk

`func (o *CodeLocation) GetNumExecutionsOk() (*int32, bool)`

GetNumExecutionsOk returns a tuple with the NumExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumExecutions

`func (o *CodeLocation) SetNumExecutions(v int32)`

SetNumExecutions sets NumExecutions field to given value.


### GetTime

`func (o *CodeLocation) GetTime() float64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CodeLocation) GetTimeOk() (*float64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CodeLocation) SetTime(v float64)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


