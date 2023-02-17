# FlowCoverageWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | **string** |  | 
**FlowName** | **string** |  | 
**FlowNamespace** | Pointer to **NullableString** |  | [optional] 
**Message** | **string** |  | 

## Methods

### NewFlowCoverageWarning

`func NewFlowCoverageWarning(flowId string, flowName string, message string, ) *FlowCoverageWarning`

NewFlowCoverageWarning instantiates a new FlowCoverageWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowCoverageWarningWithDefaults

`func NewFlowCoverageWarningWithDefaults() *FlowCoverageWarning`

NewFlowCoverageWarningWithDefaults instantiates a new FlowCoverageWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *FlowCoverageWarning) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FlowCoverageWarning) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FlowCoverageWarning) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetFlowName

`func (o *FlowCoverageWarning) GetFlowName() string`

GetFlowName returns the FlowName field if non-nil, zero value otherwise.

### GetFlowNameOk

`func (o *FlowCoverageWarning) GetFlowNameOk() (*string, bool)`

GetFlowNameOk returns a tuple with the FlowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowName

`func (o *FlowCoverageWarning) SetFlowName(v string)`

SetFlowName sets FlowName field to given value.


### GetFlowNamespace

`func (o *FlowCoverageWarning) GetFlowNamespace() string`

GetFlowNamespace returns the FlowNamespace field if non-nil, zero value otherwise.

### GetFlowNamespaceOk

`func (o *FlowCoverageWarning) GetFlowNamespaceOk() (*string, bool)`

GetFlowNamespaceOk returns a tuple with the FlowNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowNamespace

`func (o *FlowCoverageWarning) SetFlowNamespace(v string)`

SetFlowNamespace sets FlowNamespace field to given value.

### HasFlowNamespace

`func (o *FlowCoverageWarning) HasFlowNamespace() bool`

HasFlowNamespace returns a boolean if a field has been set.

### SetFlowNamespaceNil

`func (o *FlowCoverageWarning) SetFlowNamespaceNil(b bool)`

 SetFlowNamespaceNil sets the value for FlowNamespace to be an explicit nil

### UnsetFlowNamespace
`func (o *FlowCoverageWarning) UnsetFlowNamespace()`

UnsetFlowNamespace ensures that no value is present for FlowNamespace, not even an explicit nil
### GetMessage

`func (o *FlowCoverageWarning) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FlowCoverageWarning) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FlowCoverageWarning) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


