# AsyncTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassNames** | Pointer to **string** |  | [optional] 
**Classids** | Pointer to **string** |  | [optional] 
**SuiteNames** | Pointer to **string** |  | [optional] 
**Suiteids** | Pointer to **string** |  | [optional] 
**MaxFailedTests** | Pointer to **string** |  | [optional] 
**TestLevel** | Pointer to **string** |  | [optional] 
**SkipCodeCoverage** | Pointer to **string** |  | [optional] 
**Tests** | Pointer to [**[]TestsArrayItem**](TestsArrayItem.md) |  | [optional] 

## Methods

### NewAsyncTestRequest

`func NewAsyncTestRequest() *AsyncTestRequest`

NewAsyncTestRequest instantiates a new AsyncTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncTestRequestWithDefaults

`func NewAsyncTestRequestWithDefaults() *AsyncTestRequest`

NewAsyncTestRequestWithDefaults instantiates a new AsyncTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassNames

`func (o *AsyncTestRequest) GetClassNames() string`

GetClassNames returns the ClassNames field if non-nil, zero value otherwise.

### GetClassNamesOk

`func (o *AsyncTestRequest) GetClassNamesOk() (*string, bool)`

GetClassNamesOk returns a tuple with the ClassNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassNames

`func (o *AsyncTestRequest) SetClassNames(v string)`

SetClassNames sets ClassNames field to given value.

### HasClassNames

`func (o *AsyncTestRequest) HasClassNames() bool`

HasClassNames returns a boolean if a field has been set.

### GetClassids

`func (o *AsyncTestRequest) GetClassids() string`

GetClassids returns the Classids field if non-nil, zero value otherwise.

### GetClassidsOk

`func (o *AsyncTestRequest) GetClassidsOk() (*string, bool)`

GetClassidsOk returns a tuple with the Classids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassids

`func (o *AsyncTestRequest) SetClassids(v string)`

SetClassids sets Classids field to given value.

### HasClassids

`func (o *AsyncTestRequest) HasClassids() bool`

HasClassids returns a boolean if a field has been set.

### GetSuiteNames

`func (o *AsyncTestRequest) GetSuiteNames() string`

GetSuiteNames returns the SuiteNames field if non-nil, zero value otherwise.

### GetSuiteNamesOk

`func (o *AsyncTestRequest) GetSuiteNamesOk() (*string, bool)`

GetSuiteNamesOk returns a tuple with the SuiteNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteNames

`func (o *AsyncTestRequest) SetSuiteNames(v string)`

SetSuiteNames sets SuiteNames field to given value.

### HasSuiteNames

`func (o *AsyncTestRequest) HasSuiteNames() bool`

HasSuiteNames returns a boolean if a field has been set.

### GetSuiteids

`func (o *AsyncTestRequest) GetSuiteids() string`

GetSuiteids returns the Suiteids field if non-nil, zero value otherwise.

### GetSuiteidsOk

`func (o *AsyncTestRequest) GetSuiteidsOk() (*string, bool)`

GetSuiteidsOk returns a tuple with the Suiteids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteids

`func (o *AsyncTestRequest) SetSuiteids(v string)`

SetSuiteids sets Suiteids field to given value.

### HasSuiteids

`func (o *AsyncTestRequest) HasSuiteids() bool`

HasSuiteids returns a boolean if a field has been set.

### GetMaxFailedTests

`func (o *AsyncTestRequest) GetMaxFailedTests() string`

GetMaxFailedTests returns the MaxFailedTests field if non-nil, zero value otherwise.

### GetMaxFailedTestsOk

`func (o *AsyncTestRequest) GetMaxFailedTestsOk() (*string, bool)`

GetMaxFailedTestsOk returns a tuple with the MaxFailedTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailedTests

`func (o *AsyncTestRequest) SetMaxFailedTests(v string)`

SetMaxFailedTests sets MaxFailedTests field to given value.

### HasMaxFailedTests

`func (o *AsyncTestRequest) HasMaxFailedTests() bool`

HasMaxFailedTests returns a boolean if a field has been set.

### GetTestLevel

`func (o *AsyncTestRequest) GetTestLevel() string`

GetTestLevel returns the TestLevel field if non-nil, zero value otherwise.

### GetTestLevelOk

`func (o *AsyncTestRequest) GetTestLevelOk() (*string, bool)`

GetTestLevelOk returns a tuple with the TestLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestLevel

`func (o *AsyncTestRequest) SetTestLevel(v string)`

SetTestLevel sets TestLevel field to given value.

### HasTestLevel

`func (o *AsyncTestRequest) HasTestLevel() bool`

HasTestLevel returns a boolean if a field has been set.

### GetSkipCodeCoverage

`func (o *AsyncTestRequest) GetSkipCodeCoverage() string`

GetSkipCodeCoverage returns the SkipCodeCoverage field if non-nil, zero value otherwise.

### GetSkipCodeCoverageOk

`func (o *AsyncTestRequest) GetSkipCodeCoverageOk() (*string, bool)`

GetSkipCodeCoverageOk returns a tuple with the SkipCodeCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCodeCoverage

`func (o *AsyncTestRequest) SetSkipCodeCoverage(v string)`

SetSkipCodeCoverage sets SkipCodeCoverage field to given value.

### HasSkipCodeCoverage

`func (o *AsyncTestRequest) HasSkipCodeCoverage() bool`

HasSkipCodeCoverage returns a boolean if a field has been set.

### GetTests

`func (o *AsyncTestRequest) GetTests() []TestsArrayItem`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *AsyncTestRequest) GetTestsOk() (*[]TestsArrayItem, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *AsyncTestRequest) SetTests(v []TestsArrayItem)`

SetTests sets Tests field to given value.

### HasTests

`func (o *AsyncTestRequest) HasTests() bool`

HasTests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


