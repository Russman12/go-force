/*
Salesforce Platform REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 56.0
Contact: russell-laboe@outlook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tooling

import (
	"encoding/json"
)

// FlowCoverageResult struct for FlowCoverageResult
type FlowCoverageResult struct {
	ElementsNotCovered []string `json:"elementsNotCovered,omitempty"`
	FlowId string `json:"flowId"`
	FlowName string `json:"flowName"`
	FlowNamespace *string `json:"flowNamespace,omitempty"`
	NumElements int32 `json:"numElements"`
	NumElementsNotCovered int32 `json:"numElementsNotCovered"`
	ProcessType *FlowProcessType `json:"processType,omitempty"`
}

// NewFlowCoverageResult instantiates a new FlowCoverageResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowCoverageResult(flowId string, flowName string, numElements int32, numElementsNotCovered int32) *FlowCoverageResult {
	this := FlowCoverageResult{}
	this.FlowId = flowId
	this.FlowName = flowName
	this.NumElements = numElements
	this.NumElementsNotCovered = numElementsNotCovered
	return &this
}

// NewFlowCoverageResultWithDefaults instantiates a new FlowCoverageResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowCoverageResultWithDefaults() *FlowCoverageResult {
	this := FlowCoverageResult{}
	return &this
}

// GetElementsNotCovered returns the ElementsNotCovered field value if set, zero value otherwise.
func (o *FlowCoverageResult) GetElementsNotCovered() []string {
	if o == nil || isNil(o.ElementsNotCovered) {
		var ret []string
		return ret
	}
	return o.ElementsNotCovered
}

// GetElementsNotCoveredOk returns a tuple with the ElementsNotCovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowCoverageResult) GetElementsNotCoveredOk() ([]string, bool) {
	if o == nil || isNil(o.ElementsNotCovered) {
    return nil, false
	}
	return o.ElementsNotCovered, true
}

// HasElementsNotCovered returns a boolean if a field has been set.
func (o *FlowCoverageResult) HasElementsNotCovered() bool {
	if o != nil && !isNil(o.ElementsNotCovered) {
		return true
	}

	return false
}

// SetElementsNotCovered gets a reference to the given []string and assigns it to the ElementsNotCovered field.
func (o *FlowCoverageResult) SetElementsNotCovered(v []string) {
	o.ElementsNotCovered = v
}

// GetFlowId returns the FlowId field value
func (o *FlowCoverageResult) GetFlowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlowId
}

// GetFlowIdOk returns a tuple with the FlowId field value
// and a boolean to check if the value has been set.
func (o *FlowCoverageResult) GetFlowIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FlowId, true
}

// SetFlowId sets field value
func (o *FlowCoverageResult) SetFlowId(v string) {
	o.FlowId = v
}

// GetFlowName returns the FlowName field value
func (o *FlowCoverageResult) GetFlowName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlowName
}

// GetFlowNameOk returns a tuple with the FlowName field value
// and a boolean to check if the value has been set.
func (o *FlowCoverageResult) GetFlowNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FlowName, true
}

// SetFlowName sets field value
func (o *FlowCoverageResult) SetFlowName(v string) {
	o.FlowName = v
}

// GetFlowNamespace returns the FlowNamespace field value if set, zero value otherwise.
func (o *FlowCoverageResult) GetFlowNamespace() string {
	if o == nil || isNil(o.FlowNamespace) {
		var ret string
		return ret
	}
	return *o.FlowNamespace
}

// GetFlowNamespaceOk returns a tuple with the FlowNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowCoverageResult) GetFlowNamespaceOk() (*string, bool) {
	if o == nil || isNil(o.FlowNamespace) {
    return nil, false
	}
	return o.FlowNamespace, true
}

// HasFlowNamespace returns a boolean if a field has been set.
func (o *FlowCoverageResult) HasFlowNamespace() bool {
	if o != nil && !isNil(o.FlowNamespace) {
		return true
	}

	return false
}

// SetFlowNamespace gets a reference to the given string and assigns it to the FlowNamespace field.
func (o *FlowCoverageResult) SetFlowNamespace(v string) {
	o.FlowNamespace = &v
}

// GetNumElements returns the NumElements field value
func (o *FlowCoverageResult) GetNumElements() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumElements
}

// GetNumElementsOk returns a tuple with the NumElements field value
// and a boolean to check if the value has been set.
func (o *FlowCoverageResult) GetNumElementsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NumElements, true
}

// SetNumElements sets field value
func (o *FlowCoverageResult) SetNumElements(v int32) {
	o.NumElements = v
}

// GetNumElementsNotCovered returns the NumElementsNotCovered field value
func (o *FlowCoverageResult) GetNumElementsNotCovered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumElementsNotCovered
}

// GetNumElementsNotCoveredOk returns a tuple with the NumElementsNotCovered field value
// and a boolean to check if the value has been set.
func (o *FlowCoverageResult) GetNumElementsNotCoveredOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NumElementsNotCovered, true
}

// SetNumElementsNotCovered sets field value
func (o *FlowCoverageResult) SetNumElementsNotCovered(v int32) {
	o.NumElementsNotCovered = v
}

// GetProcessType returns the ProcessType field value if set, zero value otherwise.
func (o *FlowCoverageResult) GetProcessType() FlowProcessType {
	if o == nil || isNil(o.ProcessType) {
		var ret FlowProcessType
		return ret
	}
	return *o.ProcessType
}

// GetProcessTypeOk returns a tuple with the ProcessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowCoverageResult) GetProcessTypeOk() (*FlowProcessType, bool) {
	if o == nil || isNil(o.ProcessType) {
    return nil, false
	}
	return o.ProcessType, true
}

// HasProcessType returns a boolean if a field has been set.
func (o *FlowCoverageResult) HasProcessType() bool {
	if o != nil && !isNil(o.ProcessType) {
		return true
	}

	return false
}

// SetProcessType gets a reference to the given FlowProcessType and assigns it to the ProcessType field.
func (o *FlowCoverageResult) SetProcessType(v FlowProcessType) {
	o.ProcessType = &v
}

func (o FlowCoverageResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ElementsNotCovered) {
		toSerialize["elementsNotCovered"] = o.ElementsNotCovered
	}
	if true {
		toSerialize["flowId"] = o.FlowId
	}
	if true {
		toSerialize["flowName"] = o.FlowName
	}
	if !isNil(o.FlowNamespace) {
		toSerialize["flowNamespace"] = o.FlowNamespace
	}
	if true {
		toSerialize["numElements"] = o.NumElements
	}
	if true {
		toSerialize["numElementsNotCovered"] = o.NumElementsNotCovered
	}
	if !isNil(o.ProcessType) {
		toSerialize["processType"] = o.ProcessType
	}
	return json.Marshal(toSerialize)
}

type NullableFlowCoverageResult struct {
	value *FlowCoverageResult
	isSet bool
}

func (v NullableFlowCoverageResult) Get() *FlowCoverageResult {
	return v.value
}

func (v *NullableFlowCoverageResult) Set(val *FlowCoverageResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowCoverageResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowCoverageResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowCoverageResult(val *FlowCoverageResult) *NullableFlowCoverageResult {
	return &NullableFlowCoverageResult{value: val, isSet: true}
}

func (v NullableFlowCoverageResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowCoverageResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


