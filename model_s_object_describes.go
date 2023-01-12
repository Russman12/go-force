/*
Salesforce Platform APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 56.0
Contact: russell-laboe@outlook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goforce

import (
	"encoding/json"
)

// SObjectDescribes struct for SObjectDescribes
type SObjectDescribes struct {
	Encoding     string            `json:"encoding"`
	MaxBatchSize int32             `json:"maxBatchSize"`
	Sobjects     []SObjectDescribe `json:"sobjects"`
}

// NewSObjectDescribes instantiates a new SObjectDescribes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSObjectDescribes(encoding string, maxBatchSize int32, sobjects []SObjectDescribe) *SObjectDescribes {
	this := SObjectDescribes{}
	this.Encoding = encoding
	this.MaxBatchSize = maxBatchSize
	this.Sobjects = sobjects
	return &this
}

// NewSObjectDescribesWithDefaults instantiates a new SObjectDescribes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSObjectDescribesWithDefaults() *SObjectDescribes {
	this := SObjectDescribes{}
	return &this
}

// GetEncoding returns the Encoding field value
func (o *SObjectDescribes) GetEncoding() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *SObjectDescribes) GetEncodingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value
func (o *SObjectDescribes) SetEncoding(v string) {
	o.Encoding = v
}

// GetMaxBatchSize returns the MaxBatchSize field value
func (o *SObjectDescribes) GetMaxBatchSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxBatchSize
}

// GetMaxBatchSizeOk returns a tuple with the MaxBatchSize field value
// and a boolean to check if the value has been set.
func (o *SObjectDescribes) GetMaxBatchSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxBatchSize, true
}

// SetMaxBatchSize sets field value
func (o *SObjectDescribes) SetMaxBatchSize(v int32) {
	o.MaxBatchSize = v
}

// GetSobjects returns the Sobjects field value
func (o *SObjectDescribes) GetSobjects() []SObjectDescribe {
	if o == nil {
		var ret []SObjectDescribe
		return ret
	}

	return o.Sobjects
}

// GetSobjectsOk returns a tuple with the Sobjects field value
// and a boolean to check if the value has been set.
func (o *SObjectDescribes) GetSobjectsOk() ([]SObjectDescribe, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sobjects, true
}

// SetSobjects sets field value
func (o *SObjectDescribes) SetSobjects(v []SObjectDescribe) {
	o.Sobjects = v
}

func (o SObjectDescribes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["encoding"] = o.Encoding
	}
	if true {
		toSerialize["maxBatchSize"] = o.MaxBatchSize
	}
	if true {
		toSerialize["sobjects"] = o.Sobjects
	}
	return json.Marshal(toSerialize)
}

type NullableSObjectDescribes struct {
	value *SObjectDescribes
	isSet bool
}

func (v NullableSObjectDescribes) Get() *SObjectDescribes {
	return v.value
}

func (v *NullableSObjectDescribes) Set(val *SObjectDescribes) {
	v.value = val
	v.isSet = true
}

func (v NullableSObjectDescribes) IsSet() bool {
	return v.isSet
}

func (v *NullableSObjectDescribes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSObjectDescribes(val *SObjectDescribes) *NullableSObjectDescribes {
	return &NullableSObjectDescribes{value: val, isSet: true}
}

func (v NullableSObjectDescribes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSObjectDescribes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
