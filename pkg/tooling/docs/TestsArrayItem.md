# TestsArrayItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | **string** |  | 
**TestMethods** | Pointer to **[]string** |  | [optional] 
**ClassId** | Pointer to **string** |  | [optional] 
**MaxFailedTests** | Pointer to **int32** |  | [optional] 
**TestLevel** | Pointer to [**TestLevel**](TestLevel.md) |  | [optional] 

## Methods

### NewTestsArrayItem

`func NewTestsArrayItem(className string, ) *TestsArrayItem`

NewTestsArrayItem instantiates a new TestsArrayItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestsArrayItemWithDefaults

`func NewTestsArrayItemWithDefaults() *TestsArrayItem`

NewTestsArrayItemWithDefaults instantiates a new TestsArrayItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *TestsArrayItem) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *TestsArrayItem) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *TestsArrayItem) SetClassName(v string)`

SetClassName sets ClassName field to given value.


### GetTestMethods

`func (o *TestsArrayItem) GetTestMethods() []string`

GetTestMethods returns the TestMethods field if non-nil, zero value otherwise.

### GetTestMethodsOk

`func (o *TestsArrayItem) GetTestMethodsOk() (*[]string, bool)`

GetTestMethodsOk returns a tuple with the TestMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestMethods

`func (o *TestsArrayItem) SetTestMethods(v []string)`

SetTestMethods sets TestMethods field to given value.

### HasTestMethods

`func (o *TestsArrayItem) HasTestMethods() bool`

HasTestMethods returns a boolean if a field has been set.

### GetClassId

`func (o *TestsArrayItem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TestsArrayItem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TestsArrayItem) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *TestsArrayItem) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetMaxFailedTests

`func (o *TestsArrayItem) GetMaxFailedTests() int32`

GetMaxFailedTests returns the MaxFailedTests field if non-nil, zero value otherwise.

### GetMaxFailedTestsOk

`func (o *TestsArrayItem) GetMaxFailedTestsOk() (*int32, bool)`

GetMaxFailedTestsOk returns a tuple with the MaxFailedTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailedTests

`func (o *TestsArrayItem) SetMaxFailedTests(v int32)`

SetMaxFailedTests sets MaxFailedTests field to given value.

### HasMaxFailedTests

`func (o *TestsArrayItem) HasMaxFailedTests() bool`

HasMaxFailedTests returns a boolean if a field has been set.

### GetTestLevel

`func (o *TestsArrayItem) GetTestLevel() TestLevel`

GetTestLevel returns the TestLevel field if non-nil, zero value otherwise.

### GetTestLevelOk

`func (o *TestsArrayItem) GetTestLevelOk() (*TestLevel, bool)`

GetTestLevelOk returns a tuple with the TestLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestLevel

`func (o *TestsArrayItem) SetTestLevel(v TestLevel)`

SetTestLevel sets TestLevel field to given value.

### HasTestLevel

`func (o *TestsArrayItem) HasTestLevel() bool`

HasTestLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


