/*
Salesforce Platform Bulk V2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 56.0
Contact: russell-laboe@outlook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bulkv2

import (
	"encoding/json"
	"fmt"
)

// CloseOrAbortJobRequest struct for CloseOrAbortJobRequest

type CloseOrAbortJobRequestState int

const (
	CLOSE_OR_ABORT_JOB_REQUEST_STATE_UPLOAD_COMPLETE CloseOrAbortJobRequestState = iota
	CLOSE_OR_ABORT_JOB_REQUEST_STATE_ABORTED
)

func (e *CloseOrAbortJobRequestState) String() string {
	s := []string{"UploadComplete", "Aborted"}[*e]
	return s
}
func (e CloseOrAbortJobRequestState) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *CloseOrAbortJobRequestState) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	e, err := CloseOrAbortJobRequestStateParse(str)

	return err
}
func CloseOrAbortJobRequestStateParse(s string) (*CloseOrAbortJobRequestState, error) {
	for i, e := range []string{"UploadComplete", "Aborted"} {
		if s == e {
			enum := CloseOrAbortJobRequestState(i)
			return &enum, nil
		}
	}
	return nil, fmt.Errorf("%q is not a valid CloseOrAbortJobRequestState", s)
}

type CloseOrAbortJobRequest struct {
	// The state to update the job to. Use UploadComplete to close a job, or Aborted to abort a job.
	State CloseOrAbortJobRequestState `json:"state"`
}

// NewCloseOrAbortJobRequest instantiates a new CloseOrAbortJobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloseOrAbortJobRequest(state CloseOrAbortJobRequestState) *CloseOrAbortJobRequest {
	this := CloseOrAbortJobRequest{}
	this.State = state
	return &this
}

// NewCloseOrAbortJobRequestWithDefaults instantiates a new CloseOrAbortJobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloseOrAbortJobRequestWithDefaults() *CloseOrAbortJobRequest {
	this := CloseOrAbortJobRequest{}
	return &this
}

// GetState returns the State field value
func (o *CloseOrAbortJobRequest) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State.String()
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CloseOrAbortJobRequest) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	s := o.State.String()
	return &s, true
}

// SetState sets field value
func (o *CloseOrAbortJobRequest) SetState(v CloseOrAbortJobRequestState) {
	o.State = v
}

func (o CloseOrAbortJobRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableCloseOrAbortJobRequest struct {
	value *CloseOrAbortJobRequest
	isSet bool
}

func (v NullableCloseOrAbortJobRequest) Get() *CloseOrAbortJobRequest {
	return v.value
}

func (v *NullableCloseOrAbortJobRequest) Set(val *CloseOrAbortJobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloseOrAbortJobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloseOrAbortJobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloseOrAbortJobRequest(val *CloseOrAbortJobRequest) *NullableCloseOrAbortJobRequest {
	return &NullableCloseOrAbortJobRequest{value: val, isSet: true}
}

func (v NullableCloseOrAbortJobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloseOrAbortJobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
