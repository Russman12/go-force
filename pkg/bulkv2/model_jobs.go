/*
Salesforce Platform Bulk V2 APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 56.0
Contact: russell-laboe@outlook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bulkv2

import (
	"encoding/json"
)

// Jobs struct for Jobs
type Jobs struct {
	// Indicates whether there are more jobs to get. If false, use the nextRecordsUrl value to retrieve the next group of jobs.
	Done bool `json:"done"`
	// Contains information for each retrieved job.
	Records []JobInfo `json:"records"`
	// A URL that contains a query locator used to get the next set of results in a subsequent request if done isn’t true.
	NextRecordsUrl *string `json:"nextRecordsUrl,omitempty"`
}

// NewJobs instantiates a new Jobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobs(done bool, records []JobInfo) *Jobs {
	this := Jobs{}
	this.Done = done
	this.Records = records
	return &this
}

// NewJobsWithDefaults instantiates a new Jobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobsWithDefaults() *Jobs {
	this := Jobs{}
	return &this
}

// GetDone returns the Done field value
func (o *Jobs) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *Jobs) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *Jobs) SetDone(v bool) {
	o.Done = v
}

// GetRecords returns the Records field value
func (o *Jobs) GetRecords() []JobInfo {
	if o == nil {
		var ret []JobInfo
		return ret
	}

	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *Jobs) GetRecordsOk() ([]JobInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Records, true
}

// SetRecords sets field value
func (o *Jobs) SetRecords(v []JobInfo) {
	o.Records = v
}

// GetNextRecordsUrl returns the NextRecordsUrl field value if set, zero value otherwise.
func (o *Jobs) GetNextRecordsUrl() string {
	if o == nil || isNil(o.NextRecordsUrl) {
		var ret string
		return ret
	}
	return *o.NextRecordsUrl
}

// GetNextRecordsUrlOk returns a tuple with the NextRecordsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Jobs) GetNextRecordsUrlOk() (*string, bool) {
	if o == nil || isNil(o.NextRecordsUrl) {
		return nil, false
	}
	return o.NextRecordsUrl, true
}

// HasNextRecordsUrl returns a boolean if a field has been set.
func (o *Jobs) HasNextRecordsUrl() bool {
	if o != nil && !isNil(o.NextRecordsUrl) {
		return true
	}

	return false
}

// SetNextRecordsUrl gets a reference to the given string and assigns it to the NextRecordsUrl field.
func (o *Jobs) SetNextRecordsUrl(v string) {
	o.NextRecordsUrl = &v
}

func (o Jobs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["done"] = o.Done
	}
	if true {
		toSerialize["records"] = o.Records
	}
	if !isNil(o.NextRecordsUrl) {
		toSerialize["nextRecordsUrl"] = o.NextRecordsUrl
	}
	return json.Marshal(toSerialize)
}

type NullableJobs struct {
	value *Jobs
	isSet bool
}

func (v NullableJobs) Get() *Jobs {
	return v.value
}

func (v *NullableJobs) Set(val *Jobs) {
	v.value = val
	v.isSet = true
}

func (v NullableJobs) IsSet() bool {
	return v.isSet
}

func (v *NullableJobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobs(val *Jobs) *NullableJobs {
	return &NullableJobs{value: val, isSet: true}
}

func (v NullableJobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}