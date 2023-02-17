# TestResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApexLogId** | Pointer to **string** |  | [optional] 
**CodeCoverage** | Pointer to [**[]CodeCoverage**](CodeCoverage.md) |  | [optional] 
**CodeCoverageWarnings** | Pointer to [**[]CodeCoverageWarning**](CodeCoverageWarning.md) |  | [optional] 
**Failures** | Pointer to **[]map[string]interface{}** |  | [optional] 
**FlowCoverage** | Pointer to [**[]FlowCoverage**](FlowCoverage.md) |  | [optional] 
**FlowCoverageWarnings** | Pointer to **[]map[string]interface{}** |  | [optional] 
**NumFailures** | Pointer to **int32** |  | [optional] 
**NumTestsRun** | Pointer to **int32** |  | [optional] 
**Successes** | Pointer to [**[]TestResultsSuccess**](TestResultsSuccess.md) |  | [optional] 
**TotalTime** | Pointer to **int32** |  | [optional] 

## Methods

### NewTestResults

`func NewTestResults() *TestResults`

NewTestResults instantiates a new TestResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultsWithDefaults

`func NewTestResultsWithDefaults() *TestResults`

NewTestResultsWithDefaults instantiates a new TestResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApexLogId

`func (o *TestResults) GetApexLogId() string`

GetApexLogId returns the ApexLogId field if non-nil, zero value otherwise.

### GetApexLogIdOk

`func (o *TestResults) GetApexLogIdOk() (*string, bool)`

GetApexLogIdOk returns a tuple with the ApexLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApexLogId

`func (o *TestResults) SetApexLogId(v string)`

SetApexLogId sets ApexLogId field to given value.

### HasApexLogId

`func (o *TestResults) HasApexLogId() bool`

HasApexLogId returns a boolean if a field has been set.

### GetCodeCoverage

`func (o *TestResults) GetCodeCoverage() []CodeCoverage`

GetCodeCoverage returns the CodeCoverage field if non-nil, zero value otherwise.

### GetCodeCoverageOk

`func (o *TestResults) GetCodeCoverageOk() (*[]CodeCoverage, bool)`

GetCodeCoverageOk returns a tuple with the CodeCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeCoverage

`func (o *TestResults) SetCodeCoverage(v []CodeCoverage)`

SetCodeCoverage sets CodeCoverage field to given value.

### HasCodeCoverage

`func (o *TestResults) HasCodeCoverage() bool`

HasCodeCoverage returns a boolean if a field has been set.

### GetCodeCoverageWarnings

`func (o *TestResults) GetCodeCoverageWarnings() []CodeCoverageWarning`

GetCodeCoverageWarnings returns the CodeCoverageWarnings field if non-nil, zero value otherwise.

### GetCodeCoverageWarningsOk

`func (o *TestResults) GetCodeCoverageWarningsOk() (*[]CodeCoverageWarning, bool)`

GetCodeCoverageWarningsOk returns a tuple with the CodeCoverageWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeCoverageWarnings

`func (o *TestResults) SetCodeCoverageWarnings(v []CodeCoverageWarning)`

SetCodeCoverageWarnings sets CodeCoverageWarnings field to given value.

### HasCodeCoverageWarnings

`func (o *TestResults) HasCodeCoverageWarnings() bool`

HasCodeCoverageWarnings returns a boolean if a field has been set.

### GetFailures

`func (o *TestResults) GetFailures() []map[string]interface{}`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *TestResults) GetFailuresOk() (*[]map[string]interface{}, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *TestResults) SetFailures(v []map[string]interface{})`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *TestResults) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetFlowCoverage

`func (o *TestResults) GetFlowCoverage() []FlowCoverage`

GetFlowCoverage returns the FlowCoverage field if non-nil, zero value otherwise.

### GetFlowCoverageOk

`func (o *TestResults) GetFlowCoverageOk() (*[]FlowCoverage, bool)`

GetFlowCoverageOk returns a tuple with the FlowCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowCoverage

`func (o *TestResults) SetFlowCoverage(v []FlowCoverage)`

SetFlowCoverage sets FlowCoverage field to given value.

### HasFlowCoverage

`func (o *TestResults) HasFlowCoverage() bool`

HasFlowCoverage returns a boolean if a field has been set.

### GetFlowCoverageWarnings

`func (o *TestResults) GetFlowCoverageWarnings() []map[string]interface{}`

GetFlowCoverageWarnings returns the FlowCoverageWarnings field if non-nil, zero value otherwise.

### GetFlowCoverageWarningsOk

`func (o *TestResults) GetFlowCoverageWarningsOk() (*[]map[string]interface{}, bool)`

GetFlowCoverageWarningsOk returns a tuple with the FlowCoverageWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowCoverageWarnings

`func (o *TestResults) SetFlowCoverageWarnings(v []map[string]interface{})`

SetFlowCoverageWarnings sets FlowCoverageWarnings field to given value.

### HasFlowCoverageWarnings

`func (o *TestResults) HasFlowCoverageWarnings() bool`

HasFlowCoverageWarnings returns a boolean if a field has been set.

### GetNumFailures

`func (o *TestResults) GetNumFailures() int32`

GetNumFailures returns the NumFailures field if non-nil, zero value otherwise.

### GetNumFailuresOk

`func (o *TestResults) GetNumFailuresOk() (*int32, bool)`

GetNumFailuresOk returns a tuple with the NumFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailures

`func (o *TestResults) SetNumFailures(v int32)`

SetNumFailures sets NumFailures field to given value.

### HasNumFailures

`func (o *TestResults) HasNumFailures() bool`

HasNumFailures returns a boolean if a field has been set.

### GetNumTestsRun

`func (o *TestResults) GetNumTestsRun() int32`

GetNumTestsRun returns the NumTestsRun field if non-nil, zero value otherwise.

### GetNumTestsRunOk

`func (o *TestResults) GetNumTestsRunOk() (*int32, bool)`

GetNumTestsRunOk returns a tuple with the NumTestsRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTestsRun

`func (o *TestResults) SetNumTestsRun(v int32)`

SetNumTestsRun sets NumTestsRun field to given value.

### HasNumTestsRun

`func (o *TestResults) HasNumTestsRun() bool`

HasNumTestsRun returns a boolean if a field has been set.

### GetSuccesses

`func (o *TestResults) GetSuccesses() []TestResultsSuccess`

GetSuccesses returns the Successes field if non-nil, zero value otherwise.

### GetSuccessesOk

`func (o *TestResults) GetSuccessesOk() (*[]TestResultsSuccess, bool)`

GetSuccessesOk returns a tuple with the Successes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccesses

`func (o *TestResults) SetSuccesses(v []TestResultsSuccess)`

SetSuccesses sets Successes field to given value.

### HasSuccesses

`func (o *TestResults) HasSuccesses() bool`

HasSuccesses returns a boolean if a field has been set.

### GetTotalTime

`func (o *TestResults) GetTotalTime() int32`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *TestResults) GetTotalTimeOk() (*int32, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *TestResults) SetTotalTime(v int32)`

SetTotalTime sets TotalTime field to given value.

### HasTotalTime

`func (o *TestResults) HasTotalTime() bool`

HasTotalTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


