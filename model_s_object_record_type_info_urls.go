/*
Salesforce Platform APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 56.0
Contact: russell-laboe@outlook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"

)

// SObjectRecordTypeInfoUrls struct for SObjectRecordTypeInfoUrls
type SObjectRecordTypeInfoUrls struct {
	Layout *string `json:"layout,omitempty"`
}

// NewSObjectRecordTypeInfoUrls instantiates a new SObjectRecordTypeInfoUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSObjectRecordTypeInfoUrls() *SObjectRecordTypeInfoUrls {
	this := SObjectRecordTypeInfoUrls{}
	return &this
}

// NewSObjectRecordTypeInfoUrlsWithDefaults instantiates a new SObjectRecordTypeInfoUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSObjectRecordTypeInfoUrlsWithDefaults() *SObjectRecordTypeInfoUrls {
	this := SObjectRecordTypeInfoUrls{}
	return &this
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *SObjectRecordTypeInfoUrls) GetLayout() string {
	if o == nil || isNil(o.Layout) {
		var ret string
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectRecordTypeInfoUrls) GetLayoutOk() (*string, bool) {
	if o == nil || isNil(o.Layout) {
    return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *SObjectRecordTypeInfoUrls) HasLayout() bool {
	if o != nil && !isNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given string and assigns it to the Layout field.
func (o *SObjectRecordTypeInfoUrls) SetLayout(v string) {
	o.Layout = &v
}

func (o SObjectRecordTypeInfoUrls) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Layout) {
		toSerialize["layout"] = o.Layout
	}
	return json.Marshal(toSerialize)
}

type NullableSObjectRecordTypeInfoUrls struct {
	value *SObjectRecordTypeInfoUrls
	isSet bool
}

func (v NullableSObjectRecordTypeInfoUrls) Get() *SObjectRecordTypeInfoUrls {
	return v.value
}

func (v *NullableSObjectRecordTypeInfoUrls) Set(val *SObjectRecordTypeInfoUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableSObjectRecordTypeInfoUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableSObjectRecordTypeInfoUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSObjectRecordTypeInfoUrls(val *SObjectRecordTypeInfoUrls) *NullableSObjectRecordTypeInfoUrls {
	return &NullableSObjectRecordTypeInfoUrls{value: val, isSet: true}
}

func (v NullableSObjectRecordTypeInfoUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSObjectRecordTypeInfoUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


