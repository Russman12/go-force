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

// SObjectResult struct for SObjectResult
type SObjectResult struct {
	ObjectDescribe *SObjectDescribe `json:"objectDescribe,omitempty"`
	RecentItems []map[string]interface{} `json:"recentItems,omitempty"`
}

// NewSObjectResult instantiates a new SObjectResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSObjectResult() *SObjectResult {
	this := SObjectResult{}
	return &this
}

// NewSObjectResultWithDefaults instantiates a new SObjectResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSObjectResultWithDefaults() *SObjectResult {
	this := SObjectResult{}
	return &this
}

// GetObjectDescribe returns the ObjectDescribe field value if set, zero value otherwise.
func (o *SObjectResult) GetObjectDescribe() SObjectDescribe {
	if o == nil || isNil(o.ObjectDescribe) {
		var ret SObjectDescribe
		return ret
	}
	return *o.ObjectDescribe
}

// GetObjectDescribeOk returns a tuple with the ObjectDescribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectResult) GetObjectDescribeOk() (*SObjectDescribe, bool) {
	if o == nil || isNil(o.ObjectDescribe) {
    return nil, false
	}
	return o.ObjectDescribe, true
}

// HasObjectDescribe returns a boolean if a field has been set.
func (o *SObjectResult) HasObjectDescribe() bool {
	if o != nil && !isNil(o.ObjectDescribe) {
		return true
	}

	return false
}

// SetObjectDescribe gets a reference to the given SObjectDescribe and assigns it to the ObjectDescribe field.
func (o *SObjectResult) SetObjectDescribe(v SObjectDescribe) {
	o.ObjectDescribe = &v
}

// GetRecentItems returns the RecentItems field value if set, zero value otherwise.
func (o *SObjectResult) GetRecentItems() []map[string]interface{} {
	if o == nil || isNil(o.RecentItems) {
		var ret []map[string]interface{}
		return ret
	}
	return o.RecentItems
}

// GetRecentItemsOk returns a tuple with the RecentItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectResult) GetRecentItemsOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.RecentItems) {
    return nil, false
	}
	return o.RecentItems, true
}

// HasRecentItems returns a boolean if a field has been set.
func (o *SObjectResult) HasRecentItems() bool {
	if o != nil && !isNil(o.RecentItems) {
		return true
	}

	return false
}

// SetRecentItems gets a reference to the given []map[string]interface{} and assigns it to the RecentItems field.
func (o *SObjectResult) SetRecentItems(v []map[string]interface{}) {
	o.RecentItems = v
}

func (o SObjectResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ObjectDescribe) {
		toSerialize["objectDescribe"] = o.ObjectDescribe
	}
	if !isNil(o.RecentItems) {
		toSerialize["recentItems"] = o.RecentItems
	}
	return json.Marshal(toSerialize)
}

type NullableSObjectResult struct {
	value *SObjectResult
	isSet bool
}

func (v NullableSObjectResult) Get() *SObjectResult {
	return v.value
}

func (v *NullableSObjectResult) Set(val *SObjectResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSObjectResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSObjectResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSObjectResult(val *SObjectResult) *NullableSObjectResult {
	return &NullableSObjectResult{value: val, isSet: true}
}

func (v NullableSObjectResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSObjectResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


