# BulkCloseOrAbortJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The state to update the job to. Use UploadComplete to close a job, or Aborted to abort a job. | 

## Methods

### NewBulkCloseOrAbortJobRequest

`func NewBulkCloseOrAbortJobRequest(state string, ) *BulkCloseOrAbortJobRequest`

NewBulkCloseOrAbortJobRequest instantiates a new BulkCloseOrAbortJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCloseOrAbortJobRequestWithDefaults

`func NewBulkCloseOrAbortJobRequestWithDefaults() *BulkCloseOrAbortJobRequest`

NewBulkCloseOrAbortJobRequestWithDefaults instantiates a new BulkCloseOrAbortJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *BulkCloseOrAbortJobRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BulkCloseOrAbortJobRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BulkCloseOrAbortJobRequest) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


