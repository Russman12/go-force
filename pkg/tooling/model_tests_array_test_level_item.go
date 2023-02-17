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

// TestsArrayTestLevelItem struct for TestsArrayTestLevelItem
type TestsArrayTestLevelItem struct {
	TestLevel *TestLevel `json:"testLevel,omitempty"`
}

// NewTestsArrayTestLevelItem instantiates a new TestsArrayTestLevelItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestsArrayTestLevelItem() *TestsArrayTestLevelItem {
	this := TestsArrayTestLevelItem{}
	return &this
}

// NewTestsArrayTestLevelItemWithDefaults instantiates a new TestsArrayTestLevelItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestsArrayTestLevelItemWithDefaults() *TestsArrayTestLevelItem {
	this := TestsArrayTestLevelItem{}
	return &this
}

// GetTestLevel returns the TestLevel field value if set, zero value otherwise.
func (o *TestsArrayTestLevelItem) GetTestLevel() TestLevel {
	if o == nil || isNil(o.TestLevel) {
		var ret TestLevel
		return ret
	}
	return *o.TestLevel
}

// GetTestLevelOk returns a tuple with the TestLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestsArrayTestLevelItem) GetTestLevelOk() (*TestLevel, bool) {
	if o == nil || isNil(o.TestLevel) {
    return nil, false
	}
	return o.TestLevel, true
}

// HasTestLevel returns a boolean if a field has been set.
func (o *TestsArrayTestLevelItem) HasTestLevel() bool {
	if o != nil && !isNil(o.TestLevel) {
		return true
	}

	return false
}

// SetTestLevel gets a reference to the given TestLevel and assigns it to the TestLevel field.
func (o *TestsArrayTestLevelItem) SetTestLevel(v TestLevel) {
	o.TestLevel = &v
}

func (o TestsArrayTestLevelItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TestLevel) {
		toSerialize["testLevel"] = o.TestLevel
	}
	return json.Marshal(toSerialize)
}

type NullableTestsArrayTestLevelItem struct {
	value *TestsArrayTestLevelItem
	isSet bool
}

func (v NullableTestsArrayTestLevelItem) Get() *TestsArrayTestLevelItem {
	return v.value
}

func (v *NullableTestsArrayTestLevelItem) Set(val *TestsArrayTestLevelItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTestsArrayTestLevelItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTestsArrayTestLevelItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestsArrayTestLevelItem(val *TestsArrayTestLevelItem) *NullableTestsArrayTestLevelItem {
	return &NullableTestsArrayTestLevelItem{value: val, isSet: true}
}

func (v NullableTestsArrayTestLevelItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestsArrayTestLevelItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


