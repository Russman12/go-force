# CloseOrAbortJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**JobCloseAbortState**](JobCloseAbortState.md) |  | 

## Methods

### NewCloseOrAbortJobRequest

`func NewCloseOrAbortJobRequest(state JobCloseAbortState, ) *CloseOrAbortJobRequest`

NewCloseOrAbortJobRequest instantiates a new CloseOrAbortJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloseOrAbortJobRequestWithDefaults

`func NewCloseOrAbortJobRequestWithDefaults() *CloseOrAbortJobRequest`

NewCloseOrAbortJobRequestWithDefaults instantiates a new CloseOrAbortJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CloseOrAbortJobRequest) GetState() JobCloseAbortState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloseOrAbortJobRequest) GetStateOk() (*JobCloseAbortState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloseOrAbortJobRequest) SetState(v JobCloseAbortState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


