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
)

// QueryJobInfos struct for QueryJobInfos
type QueryJobInfos struct {
	// This is true if this is the last (or only) set of results. It is false if there are more records to fetch.
	Done *bool `json:"done,omitempty"`
	// An array of record objects.
	Records []QueryJobInfo `json:"records,omitempty"`
	// The URI to get the next set of records (if there are any). This method returns up to 1,000 result rows per request. If there are more than 1,000 records, use the nextRecordsUrl to get the next set of records. This parameter is null if there are no more records to fetch.
	NextRecordsUrl *string `json:"nextRecordsUrl,omitempty"`
}

// NewQueryJobInfos instantiates a new QueryJobInfos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryJobInfos() *QueryJobInfos {
	this := QueryJobInfos{}
	return &this
}

// NewQueryJobInfosWithDefaults instantiates a new QueryJobInfos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryJobInfosWithDefaults() *QueryJobInfos {
	this := QueryJobInfos{}
	return &this
}

// GetDone returns the Done field value if set, zero value otherwise.
func (o *QueryJobInfos) GetDone() bool {
	if o == nil || isNil(o.Done) {
		var ret bool
		return ret
	}
	return *o.Done
}

// GetDoneOk returns a tuple with the Done field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryJobInfos) GetDoneOk() (*bool, bool) {
	if o == nil || isNil(o.Done) {
		return nil, false
	}
	return o.Done, true
}

// HasDone returns a boolean if a field has been set.
func (o *QueryJobInfos) HasDone() bool {
	if o != nil && !isNil(o.Done) {
		return true
	}

	return false
}

// SetDone gets a reference to the given bool and assigns it to the Done field.
func (o *QueryJobInfos) SetDone(v bool) {
	o.Done = &v
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *QueryJobInfos) GetRecords() []QueryJobInfo {
	if o == nil || isNil(o.Records) {
		var ret []QueryJobInfo
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryJobInfos) GetRecordsOk() ([]QueryJobInfo, bool) {
	if o == nil || isNil(o.Records) {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *QueryJobInfos) HasRecords() bool {
	if o != nil && !isNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []QueryJobInfo and assigns it to the Records field.
func (o *QueryJobInfos) SetRecords(v []QueryJobInfo) {
	o.Records = v
}

// GetNextRecordsUrl returns the NextRecordsUrl field value if set, zero value otherwise.
func (o *QueryJobInfos) GetNextRecordsUrl() string {
	if o == nil || isNil(o.NextRecordsUrl) {
		var ret string
		return ret
	}
	return *o.NextRecordsUrl
}

// GetNextRecordsUrlOk returns a tuple with the NextRecordsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryJobInfos) GetNextRecordsUrlOk() (*string, bool) {
	if o == nil || isNil(o.NextRecordsUrl) {
		return nil, false
	}
	return o.NextRecordsUrl, true
}

// HasNextRecordsUrl returns a boolean if a field has been set.
func (o *QueryJobInfos) HasNextRecordsUrl() bool {
	if o != nil && !isNil(o.NextRecordsUrl) {
		return true
	}

	return false
}

// SetNextRecordsUrl gets a reference to the given string and assigns it to the NextRecordsUrl field.
func (o *QueryJobInfos) SetNextRecordsUrl(v string) {
	o.NextRecordsUrl = &v
}

func (o QueryJobInfos) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Done) {
		toSerialize["done"] = o.Done
	}
	if !isNil(o.Records) {
		toSerialize["records"] = o.Records
	}
	if !isNil(o.NextRecordsUrl) {
		toSerialize["nextRecordsUrl"] = o.NextRecordsUrl
	}
	return json.Marshal(toSerialize)
}

type NullableQueryJobInfos struct {
	value *QueryJobInfos
	isSet bool
}

func (v NullableQueryJobInfos) Get() *QueryJobInfos {
	return v.value
}

func (v *NullableQueryJobInfos) Set(val *QueryJobInfos) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryJobInfos) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryJobInfos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryJobInfos(val *QueryJobInfos) *NullableQueryJobInfos {
	return &NullableQueryJobInfos{value: val, isSet: true}
}

func (v NullableQueryJobInfos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryJobInfos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
