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

// BulkQueryJobsInfos struct for BulkQueryJobsInfos
type BulkQueryJobsInfos struct {
	// This is true if this is the last (or only) set of results. It is false if there are more records to fetch.
	Done bool `json:"done"`
	// An array of record objects.
	Records []BulkQueryJobInfo `json:"records"`
	// The URI to get the next set of records (if there are any). This method returns up to 1,000 result rows per request. If there are more than 1,000 records, use the nextRecordsUrl to get the next set of records. This parameter is null if there are no more records to fetch. 
	NextRecordsUrl *string `json:"nextRecordsUrl,omitempty"`
}

// NewBulkQueryJobsInfos instantiates a new BulkQueryJobsInfos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkQueryJobsInfos(done bool, records []BulkQueryJobInfo) *BulkQueryJobsInfos {
	this := BulkQueryJobsInfos{}
	this.Done = done
	this.Records = records
	return &this
}

// NewBulkQueryJobsInfosWithDefaults instantiates a new BulkQueryJobsInfos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkQueryJobsInfosWithDefaults() *BulkQueryJobsInfos {
	this := BulkQueryJobsInfos{}
	return &this
}

// GetDone returns the Done field value
func (o *BulkQueryJobsInfos) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *BulkQueryJobsInfos) GetDoneOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *BulkQueryJobsInfos) SetDone(v bool) {
	o.Done = v
}

// GetRecords returns the Records field value
func (o *BulkQueryJobsInfos) GetRecords() []BulkQueryJobInfo {
	if o == nil {
		var ret []BulkQueryJobInfo
		return ret
	}

	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *BulkQueryJobsInfos) GetRecordsOk() ([]BulkQueryJobInfo, bool) {
	if o == nil {
    return nil, false
	}
	return o.Records, true
}

// SetRecords sets field value
func (o *BulkQueryJobsInfos) SetRecords(v []BulkQueryJobInfo) {
	o.Records = v
}

// GetNextRecordsUrl returns the NextRecordsUrl field value if set, zero value otherwise.
func (o *BulkQueryJobsInfos) GetNextRecordsUrl() string {
	if o == nil || isNil(o.NextRecordsUrl) {
		var ret string
		return ret
	}
	return *o.NextRecordsUrl
}

// GetNextRecordsUrlOk returns a tuple with the NextRecordsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQueryJobsInfos) GetNextRecordsUrlOk() (*string, bool) {
	if o == nil || isNil(o.NextRecordsUrl) {
    return nil, false
	}
	return o.NextRecordsUrl, true
}

// HasNextRecordsUrl returns a boolean if a field has been set.
func (o *BulkQueryJobsInfos) HasNextRecordsUrl() bool {
	if o != nil && !isNil(o.NextRecordsUrl) {
		return true
	}

	return false
}

// SetNextRecordsUrl gets a reference to the given string and assigns it to the NextRecordsUrl field.
func (o *BulkQueryJobsInfos) SetNextRecordsUrl(v string) {
	o.NextRecordsUrl = &v
}

func (o BulkQueryJobsInfos) MarshalJSON() ([]byte, error) {
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

type NullableBulkQueryJobsInfos struct {
	value *BulkQueryJobsInfos
	isSet bool
}

func (v NullableBulkQueryJobsInfos) Get() *BulkQueryJobsInfos {
	return v.value
}

func (v *NullableBulkQueryJobsInfos) Set(val *BulkQueryJobsInfos) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkQueryJobsInfos) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkQueryJobsInfos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkQueryJobsInfos(val *BulkQueryJobsInfos) *NullableBulkQueryJobsInfos {
	return &NullableBulkQueryJobsInfos{value: val, isSet: true}
}

func (v NullableBulkQueryJobsInfos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkQueryJobsInfos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

