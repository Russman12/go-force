# FlowCoverageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElementsNotCovered** | Pointer to **[]string** |  | [optional] 
**FlowId** | **string** |  | 
**FlowName** | **string** |  | 
**FlowNamespace** | Pointer to **string** |  | [optional] 
**NumElements** | **int32** |  | 
**NumElementsNotCovered** | **int32** |  | 
**ProcessType** | Pointer to [**FlowProcessType**](FlowProcessType.md) |  | [optional] 

## Methods

### NewFlowCoverageResult

`func NewFlowCoverageResult(flowId string, flowName string, numElements int32, numElementsNotCovered int32, ) *FlowCoverageResult`

NewFlowCoverageResult instantiates a new FlowCoverageResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowCoverageResultWithDefaults

`func NewFlowCoverageResultWithDefaults() *FlowCoverageResult`

NewFlowCoverageResultWithDefaults instantiates a new FlowCoverageResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElementsNotCovered

`func (o *FlowCoverageResult) GetElementsNotCovered() []string`

GetElementsNotCovered returns the ElementsNotCovered field if non-nil, zero value otherwise.

### GetElementsNotCoveredOk

`func (o *FlowCoverageResult) GetElementsNotCoveredOk() (*[]string, bool)`

GetElementsNotCoveredOk returns a tuple with the ElementsNotCovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementsNotCovered

`func (o *FlowCoverageResult) SetElementsNotCovered(v []string)`

SetElementsNotCovered sets ElementsNotCovered field to given value.

### HasElementsNotCovered

`func (o *FlowCoverageResult) HasElementsNotCovered() bool`

HasElementsNotCovered returns a boolean if a field has been set.

### GetFlowId

`func (o *FlowCoverageResult) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FlowCoverageResult) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FlowCoverageResult) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetFlowName

`func (o *FlowCoverageResult) GetFlowName() string`

GetFlowName returns the FlowName field if non-nil, zero value otherwise.

### GetFlowNameOk

`func (o *FlowCoverageResult) GetFlowNameOk() (*string, bool)`

GetFlowNameOk returns a tuple with the FlowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowName

`func (o *FlowCoverageResult) SetFlowName(v string)`

SetFlowName sets FlowName field to given value.


### GetFlowNamespace

`func (o *FlowCoverageResult) GetFlowNamespace() string`

GetFlowNamespace returns the FlowNamespace field if non-nil, zero value otherwise.

### GetFlowNamespaceOk

`func (o *FlowCoverageResult) GetFlowNamespaceOk() (*string, bool)`

GetFlowNamespaceOk returns a tuple with the FlowNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowNamespace

`func (o *FlowCoverageResult) SetFlowNamespace(v string)`

SetFlowNamespace sets FlowNamespace field to given value.

### HasFlowNamespace

`func (o *FlowCoverageResult) HasFlowNamespace() bool`

HasFlowNamespace returns a boolean if a field has been set.

### GetNumElements

`func (o *FlowCoverageResult) GetNumElements() int32`

GetNumElements returns the NumElements field if non-nil, zero value otherwise.

### GetNumElementsOk

`func (o *FlowCoverageResult) GetNumElementsOk() (*int32, bool)`

GetNumElementsOk returns a tuple with the NumElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumElements

`func (o *FlowCoverageResult) SetNumElements(v int32)`

SetNumElements sets NumElements field to given value.


### GetNumElementsNotCovered

`func (o *FlowCoverageResult) GetNumElementsNotCovered() int32`

GetNumElementsNotCovered returns the NumElementsNotCovered field if non-nil, zero value otherwise.

### GetNumElementsNotCoveredOk

`func (o *FlowCoverageResult) GetNumElementsNotCoveredOk() (*int32, bool)`

GetNumElementsNotCoveredOk returns a tuple with the NumElementsNotCovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumElementsNotCovered

`func (o *FlowCoverageResult) SetNumElementsNotCovered(v int32)`

SetNumElementsNotCovered sets NumElementsNotCovered field to given value.


### GetProcessType

`func (o *FlowCoverageResult) GetProcessType() FlowProcessType`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *FlowCoverageResult) GetProcessTypeOk() (*FlowProcessType, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *FlowCoverageResult) SetProcessType(v FlowProcessType)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *FlowCoverageResult) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


