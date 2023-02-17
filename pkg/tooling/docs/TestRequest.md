# TestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tests** | Pointer to [**[]TestsArrayItem**](TestsArrayItem.md) |  | [optional] 

## Methods

### NewTestRequest

`func NewTestRequest() *TestRequest`

NewTestRequest instantiates a new TestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRequestWithDefaults

`func NewTestRequestWithDefaults() *TestRequest`

NewTestRequestWithDefaults instantiates a new TestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTests

`func (o *TestRequest) GetTests() []TestsArrayItem`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *TestRequest) GetTestsOk() (*[]TestsArrayItem, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *TestRequest) SetTests(v []TestsArrayItem)`

SetTests sets Tests field to given value.

### HasTests

`func (o *TestRequest) HasTests() bool`

HasTests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


