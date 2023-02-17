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

// CodeLocation struct for CodeLocation
type CodeLocation struct {
	Column int32 `json:"column"`
	Line int32 `json:"line"`
	NumExecutions int32 `json:"numExecutions"`
	Time float64 `json:"time"`
}

// NewCodeLocation instantiates a new CodeLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeLocation(column int32, line int32, numExecutions int32, time float64) *CodeLocation {
	this := CodeLocation{}
	this.Column = column
	this.Line = line
	this.NumExecutions = numExecutions
	this.Time = time
	return &this
}

// NewCodeLocationWithDefaults instantiates a new CodeLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeLocationWithDefaults() *CodeLocation {
	this := CodeLocation{}
	return &this
}

// GetColumn returns the Column field value
func (o *CodeLocation) GetColumn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *CodeLocation) GetColumnOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value
func (o *CodeLocation) SetColumn(v int32) {
	o.Column = v
}

// GetLine returns the Line field value
func (o *CodeLocation) GetLine() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Line
}

// GetLineOk returns a tuple with the Line field value
// and a boolean to check if the value has been set.
func (o *CodeLocation) GetLineOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Line, true
}

// SetLine sets field value
func (o *CodeLocation) SetLine(v int32) {
	o.Line = v
}

// GetNumExecutions returns the NumExecutions field value
func (o *CodeLocation) GetNumExecutions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumExecutions
}

// GetNumExecutionsOk returns a tuple with the NumExecutions field value
// and a boolean to check if the value has been set.
func (o *CodeLocation) GetNumExecutionsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NumExecutions, true
}

// SetNumExecutions sets field value
func (o *CodeLocation) SetNumExecutions(v int32) {
	o.NumExecutions = v
}

// GetTime returns the Time field value
func (o *CodeLocation) GetTime() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *CodeLocation) GetTimeOk() (*float64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *CodeLocation) SetTime(v float64) {
	o.Time = v
}

func (o CodeLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["column"] = o.Column
	}
	if true {
		toSerialize["line"] = o.Line
	}
	if true {
		toSerialize["numExecutions"] = o.NumExecutions
	}
	if true {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableCodeLocation struct {
	value *CodeLocation
	isSet bool
}

func (v NullableCodeLocation) Get() *CodeLocation {
	return v.value
}

func (v *NullableCodeLocation) Set(val *CodeLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeLocation(val *CodeLocation) *NullableCodeLocation {
	return &NullableCodeLocation{value: val, isSet: true}
}

func (v NullableCodeLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

