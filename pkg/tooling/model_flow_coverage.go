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

// FlowCoverage struct for FlowCoverage
type FlowCoverage struct {
	ElementsNotCovered []string `json:"elementsNotCovered,omitempty"`
	FlowId *string `json:"flowId,omitempty"`
	FlowName *string `json:"flowName,omitempty"`
	FlowNamespace *string `json:"flowNamespace,omitempty"`
	NumElements *int32 `json:"numElements,omitempty"`
	NumElementsNotCovered *int32 `json:"numElementsNotCovered,omitempty"`
	ProcessType *string `json:"processType,omitempty"`
}

// NewFlowCoverage instantiates a new FlowCoverage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowCoverage() *FlowCoverage {
	this := FlowCoverage{}
	return &this
}

// NewFlowCoverageWithDefaults instantiates a new FlowCoverage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowCoverageWithDefaults() *FlowCoverage {
	this := FlowCoverage{}
	return &this
}

// GetElementsNotCovered returns the ElementsNotCovered field value if set, zero value otherwise.
func (o *FlowCoverage) GetElementsNotCovered() []string {
	if o == nil || isNil(o.ElementsNotCovered) {
		var ret []string
		return ret
	}
	return o.ElementsNotCovered
}

// GetElementsNotCoveredOk returns a tuple with the ElementsNotCovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowCoverage) GetElementsNotCoveredOk() ([]string, bool) {
	if o == nil || isNil(o.ElementsNotCovered) {
    return nil, false
	}
	return o.ElementsNotCovered, true
}

// HasElementsNotCovered returns a boolean if a field has been set.
func (o *FlowCoverage) HasElementsNotCovered() bool {
	if o != nil && !isNil(o.ElementsNotCovered) {
		return true
	}

	return false
}

// SetElementsNotCovered gets a reference to the given []string and assigns it to the ElementsNotCovered field.
func (o *FlowCoverage) SetElementsNotCovered(v []string) {
	o.ElementsNotCovered = v
}

// GetFlowId returns the FlowId field value if set, zero value otherwise.
func (o *FlowCoverage) GetFlowId() string {
	if o == nil || isNil(o.FlowId) {
		var ret string
		return ret
	}
	return *o.FlowId
}

// GetFlowIdOk returns a tuple with the FlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowCoverage) GetFlowIdOk() (*string, bool) {
	if o == nil || isNil(o.FlowId) {
    return nil, false
	}
	return o.FlowId, true
}

// HasFlowId returns a boolean if a field has been set.
func (o *FlowCoverage) HasFlowId() bool {
	if o != nil && !isNil(o.FlowId) {
		return true
	}

	return false
}

// SetFlowId gets a reference to the given string and assigns it to the FlowId field.
func (o *FlowCoverage) SetFlowId(v string) {
	o.FlowId = &v
}

// GetFlowName returns the FlowName field value if set, zero value otherwise.
func (o *FlowCoverage) GetFlowName() string {
	if o == nil || isNil(o.FlowName) {
		var ret string
		return ret
	}
	return *o.FlowName
}

// GetFlowNameOk returns a tuple with the FlowName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowCoverage) GetFlowNameOk() (*string, bool) {
	if o == nil || isNil(o.FlowName) {
    return nil, false
	}
	return o.FlowName, true
}

// HasFlowName returns a boolean if a field has been set.
func (o *FlowCoverage) HasFlowName() bool {
	if o != nil && !isNil(o.FlowName) {
		return true
	}

	return false
}

// SetFlowName gets a reference to the given string and assigns it to the FlowName field.
func (o *FlowCoverage) SetFlowName(v string) {
	o.FlowName = &v
}

// GetFlowNamespace returns the FlowNamespace field value if set, zero value otherwise.
func (o *FlowCoverage) GetFlowNamespace() string {
	if o == nil || isNil(o.FlowNamespace) {
		var ret string
		return ret
	}
	return *o.FlowNamespace
}

// GetFlowNamespaceOk returns a tuple with the FlowNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowCoverage) GetFlowNamespaceOk() (*string, bool) {
	if o == nil || isNil(o.FlowNamespace) {
    return nil, false
	}
	return o.FlowNamespace, true
}

// HasFlowNamespace returns a boolean if a field has been set.
func (o *FlowCoverage) HasFlowNamespace() bool {
	if o != nil && !isNil(o.FlowNamespace) {
		return true
	}

	return false
}

// SetFlowNamespace gets a reference to the given string and assigns it to the FlowNamespace field.
func (o *FlowCoverage) SetFlowNamespace(v string) {
	o.FlowNamespace = &v
}

// GetNumElements returns the NumElements field value if set, zero value otherwise.
func (o *FlowCoverage) GetNumElements() int32 {
	if o == nil || isNil(o.NumElements) {
		var ret int32
		return ret
	}
	return *o.NumElements
}

// GetNumElementsOk returns a tuple with the NumElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowCoverage) GetNumElementsOk() (*int32, bool) {
	if o == nil || isNil(o.NumElements) {
    return nil, false
	}
	return o.NumElements, true
}

// HasNumElements returns a boolean if a field has been set.
func (o *FlowCoverage) HasNumElements() bool {
	if o != nil && !isNil(o.NumElements) {
		return true
	}

	return false
}

// SetNumElements gets a reference to the given int32 and assigns it to the NumElements field.
func (o *FlowCoverage) SetNumElements(v int32) {
	o.NumElements = &v
}

// GetNumElementsNotCovered returns the NumElementsNotCovered field value if set, zero value otherwise.
func (o *FlowCoverage) GetNumElementsNotCovered() int32 {
	if o == nil || isNil(o.NumElementsNotCovered) {
		var ret int32
		return ret
	}
	return *o.NumElementsNotCovered
}

// GetNumElementsNotCoveredOk returns a tuple with the NumElementsNotCovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowCoverage) GetNumElementsNotCoveredOk() (*int32, bool) {
	if o == nil || isNil(o.NumElementsNotCovered) {
    return nil, false
	}
	return o.NumElementsNotCovered, true
}

// HasNumElementsNotCovered returns a boolean if a field has been set.
func (o *FlowCoverage) HasNumElementsNotCovered() bool {
	if o != nil && !isNil(o.NumElementsNotCovered) {
		return true
	}

	return false
}

// SetNumElementsNotCovered gets a reference to the given int32 and assigns it to the NumElementsNotCovered field.
func (o *FlowCoverage) SetNumElementsNotCovered(v int32) {
	o.NumElementsNotCovered = &v
}

// GetProcessType returns the ProcessType field value if set, zero value otherwise.
func (o *FlowCoverage) GetProcessType() string {
	if o == nil || isNil(o.ProcessType) {
		var ret string
		return ret
	}
	return *o.ProcessType
}

// GetProcessTypeOk returns a tuple with the ProcessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowCoverage) GetProcessTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProcessType) {
    return nil, false
	}
	return o.ProcessType, true
}

// HasProcessType returns a boolean if a field has been set.
func (o *FlowCoverage) HasProcessType() bool {
	if o != nil && !isNil(o.ProcessType) {
		return true
	}

	return false
}

// SetProcessType gets a reference to the given string and assigns it to the ProcessType field.
func (o *FlowCoverage) SetProcessType(v string) {
	o.ProcessType = &v
}

func (o FlowCoverage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ElementsNotCovered) {
		toSerialize["elementsNotCovered"] = o.ElementsNotCovered
	}
	if !isNil(o.FlowId) {
		toSerialize["flowId"] = o.FlowId
	}
	if !isNil(o.FlowName) {
		toSerialize["flowName"] = o.FlowName
	}
	if !isNil(o.FlowNamespace) {
		toSerialize["flowNamespace"] = o.FlowNamespace
	}
	if !isNil(o.NumElements) {
		toSerialize["numElements"] = o.NumElements
	}
	if !isNil(o.NumElementsNotCovered) {
		toSerialize["numElementsNotCovered"] = o.NumElementsNotCovered
	}
	if !isNil(o.ProcessType) {
		toSerialize["processType"] = o.ProcessType
	}
	return json.Marshal(toSerialize)
}

type NullableFlowCoverage struct {
	value *FlowCoverage
	isSet bool
}

func (v NullableFlowCoverage) Get() *FlowCoverage {
	return v.value
}

func (v *NullableFlowCoverage) Set(val *FlowCoverage) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowCoverage) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowCoverage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowCoverage(val *FlowCoverage) *NullableFlowCoverage {
	return &NullableFlowCoverage{value: val, isSet: true}
}

func (v NullableFlowCoverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowCoverage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


