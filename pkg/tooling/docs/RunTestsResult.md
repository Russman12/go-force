# RunTestsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApexLogId** | **string** |  | 
**CodeCoverage** | Pointer to [**[]CodeCoverageResult**](CodeCoverageResult.md) |  | [optional] 
**CodeCoverageWarnings** | Pointer to [**[]CodeCoverageWarning**](CodeCoverageWarning.md) |  | [optional] 
**Failures** | Pointer to **[]map[string]interface{}** |  | [optional] 
**FlowCoverage** | Pointer to [**[]FlowCoverageResult**](FlowCoverageResult.md) |  | [optional] 
**FlowCoverageWarnings** | Pointer to [**[]FlowCoverageWarning**](FlowCoverageWarning.md) |  | [optional] 
**NumFailures** | **int32** |  | 
**NumTestsRun** | **int32** |  | 
**Successes** | Pointer to [**[]TestResultsSuccess**](TestResultsSuccess.md) |  | [optional] 
**TotalTime** | **float64** |  | 

## Methods

### NewRunTestsResult

`func NewRunTestsResult(apexLogId string, numFailures int32, numTestsRun int32, totalTime float64, ) *RunTestsResult`

NewRunTestsResult instantiates a new RunTestsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunTestsResultWithDefaults

`func NewRunTestsResultWithDefaults() *RunTestsResult`

NewRunTestsResultWithDefaults instantiates a new RunTestsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApexLogId

`func (o *RunTestsResult) GetApexLogId() string`

GetApexLogId returns the ApexLogId field if non-nil, zero value otherwise.

### GetApexLogIdOk

`func (o *RunTestsResult) GetApexLogIdOk() (*string, bool)`

GetApexLogIdOk returns a tuple with the ApexLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApexLogId

`func (o *RunTestsResult) SetApexLogId(v string)`

SetApexLogId sets ApexLogId field to given value.


### GetCodeCoverage

`func (o *RunTestsResult) GetCodeCoverage() []CodeCoverageResult`

GetCodeCoverage returns the CodeCoverage field if non-nil, zero value otherwise.

### GetCodeCoverageOk

`func (o *RunTestsResult) GetCodeCoverageOk() (*[]CodeCoverageResult, bool)`

GetCodeCoverageOk returns a tuple with the CodeCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeCoverage

`func (o *RunTestsResult) SetCodeCoverage(v []CodeCoverageResult)`

SetCodeCoverage sets CodeCoverage field to given value.

### HasCodeCoverage

`func (o *RunTestsResult) HasCodeCoverage() bool`

HasCodeCoverage returns a boolean if a field has been set.

### GetCodeCoverageWarnings

`func (o *RunTestsResult) GetCodeCoverageWarnings() []CodeCoverageWarning`

GetCodeCoverageWarnings returns the CodeCoverageWarnings field if non-nil, zero value otherwise.

### GetCodeCoverageWarningsOk

`func (o *RunTestsResult) GetCodeCoverageWarningsOk() (*[]CodeCoverageWarning, bool)`

GetCodeCoverageWarningsOk returns a tuple with the CodeCoverageWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeCoverageWarnings

`func (o *RunTestsResult) SetCodeCoverageWarnings(v []CodeCoverageWarning)`

SetCodeCoverageWarnings sets CodeCoverageWarnings field to given value.

### HasCodeCoverageWarnings

`func (o *RunTestsResult) HasCodeCoverageWarnings() bool`

HasCodeCoverageWarnings returns a boolean if a field has been set.

### GetFailures

`func (o *RunTestsResult) GetFailures() []map[string]interface{}`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *RunTestsResult) GetFailuresOk() (*[]map[string]interface{}, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *RunTestsResult) SetFailures(v []map[string]interface{})`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *RunTestsResult) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetFlowCoverage

`func (o *RunTestsResult) GetFlowCoverage() []FlowCoverageResult`

GetFlowCoverage returns the FlowCoverage field if non-nil, zero value otherwise.

### GetFlowCoverageOk

`func (o *RunTestsResult) GetFlowCoverageOk() (*[]FlowCoverageResult, bool)`

GetFlowCoverageOk returns a tuple with the FlowCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowCoverage

`func (o *RunTestsResult) SetFlowCoverage(v []FlowCoverageResult)`

SetFlowCoverage sets FlowCoverage field to given value.

### HasFlowCoverage

`func (o *RunTestsResult) HasFlowCoverage() bool`

HasFlowCoverage returns a boolean if a field has been set.

### GetFlowCoverageWarnings

`func (o *RunTestsResult) GetFlowCoverageWarnings() []FlowCoverageWarning`

GetFlowCoverageWarnings returns the FlowCoverageWarnings field if non-nil, zero value otherwise.

### GetFlowCoverageWarningsOk

`func (o *RunTestsResult) GetFlowCoverageWarningsOk() (*[]FlowCoverageWarning, bool)`

GetFlowCoverageWarningsOk returns a tuple with the FlowCoverageWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowCoverageWarnings

`func (o *RunTestsResult) SetFlowCoverageWarnings(v []FlowCoverageWarning)`

SetFlowCoverageWarnings sets FlowCoverageWarnings field to given value.

### HasFlowCoverageWarnings

`func (o *RunTestsResult) HasFlowCoverageWarnings() bool`

HasFlowCoverageWarnings returns a boolean if a field has been set.

### GetNumFailures

`func (o *RunTestsResult) GetNumFailures() int32`

GetNumFailures returns the NumFailures field if non-nil, zero value otherwise.

### GetNumFailuresOk

`func (o *RunTestsResult) GetNumFailuresOk() (*int32, bool)`

GetNumFailuresOk returns a tuple with the NumFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailures

`func (o *RunTestsResult) SetNumFailures(v int32)`

SetNumFailures sets NumFailures field to given value.


### GetNumTestsRun

`func (o *RunTestsResult) GetNumTestsRun() int32`

GetNumTestsRun returns the NumTestsRun field if non-nil, zero value otherwise.

### GetNumTestsRunOk

`func (o *RunTestsResult) GetNumTestsRunOk() (*int32, bool)`

GetNumTestsRunOk returns a tuple with the NumTestsRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTestsRun

`func (o *RunTestsResult) SetNumTestsRun(v int32)`

SetNumTestsRun sets NumTestsRun field to given value.


### GetSuccesses

`func (o *RunTestsResult) GetSuccesses() []TestResultsSuccess`

GetSuccesses returns the Successes field if non-nil, zero value otherwise.

### GetSuccessesOk

`func (o *RunTestsResult) GetSuccessesOk() (*[]TestResultsSuccess, bool)`

GetSuccessesOk returns a tuple with the Successes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccesses

`func (o *RunTestsResult) SetSuccesses(v []TestResultsSuccess)`

SetSuccesses sets Successes field to given value.

### HasSuccesses

`func (o *RunTestsResult) HasSuccesses() bool`

HasSuccesses returns a boolean if a field has been set.

### GetTotalTime

`func (o *RunTestsResult) GetTotalTime() float64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *RunTestsResult) GetTotalTimeOk() (*float64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *RunTestsResult) SetTotalTime(v float64)`

SetTotalTime sets TotalTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


