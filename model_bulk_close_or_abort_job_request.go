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

// BulkCloseOrAbortJobRequest struct for BulkCloseOrAbortJobRequest
type BulkCloseOrAbortJobRequest struct {
	// The state to update the job to. Use UploadComplete to close a job, or Aborted to abort a job.
	State string `json:"state"`
}

// NewBulkCloseOrAbortJobRequest instantiates a new BulkCloseOrAbortJobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkCloseOrAbortJobRequest(state string) *BulkCloseOrAbortJobRequest {
	this := BulkCloseOrAbortJobRequest{}
	this.State = state
	return &this
}

// NewBulkCloseOrAbortJobRequestWithDefaults instantiates a new BulkCloseOrAbortJobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkCloseOrAbortJobRequestWithDefaults() *BulkCloseOrAbortJobRequest {
	this := BulkCloseOrAbortJobRequest{}
	return &this
}

// GetState returns the State field value
func (o *BulkCloseOrAbortJobRequest) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *BulkCloseOrAbortJobRequest) GetStateOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *BulkCloseOrAbortJobRequest) SetState(v string) {
	o.State = v
}

func (o BulkCloseOrAbortJobRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableBulkCloseOrAbortJobRequest struct {
	value *BulkCloseOrAbortJobRequest
	isSet bool
}

func (v NullableBulkCloseOrAbortJobRequest) Get() *BulkCloseOrAbortJobRequest {
	return v.value
}

func (v *NullableBulkCloseOrAbortJobRequest) Set(val *BulkCloseOrAbortJobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkCloseOrAbortJobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkCloseOrAbortJobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkCloseOrAbortJobRequest(val *BulkCloseOrAbortJobRequest) *NullableBulkCloseOrAbortJobRequest {
	return &NullableBulkCloseOrAbortJobRequest{value: val, isSet: true}
}

func (v NullableBulkCloseOrAbortJobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkCloseOrAbortJobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


