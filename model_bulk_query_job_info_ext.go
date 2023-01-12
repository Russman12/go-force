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

// BulkQueryJobInfoExt struct for BulkQueryJobInfoExt
type BulkQueryJobInfoExt struct {
	// The unique ID for this job.
	Id *string `json:"id,omitempty"`
	// The object type being queried.
	Object *string `json:"object,omitempty"`
	// The ID of the user who created the job.
	CreatedById *string `json:"createdById,omitempty"`
	// The UTC date and time when the job was created.
	CreatedDate *string `json:"createdDate,omitempty"`
	// The UTC date and time when the API last updated the job information.
	SystemModstamp *string `json:"systemModstamp,omitempty"`
	// The current state of processing for the job.
	State *string `json:"state,omitempty"`
	// Reserved for future use. How the request is processed. Currently only parallel mode is supported. (When other modes are added, the API chooses the mode automatically. The mode isn’t user configurable.)
	ConcurrencyMode *string `json:"concurrencyMode,omitempty"`
	// The API version that the job was created in.
	ApiVersion *float32 `json:"apiVersion,omitempty"`
	// The job’s type. For a query job, the type is always V2Query.
	JobType *string `json:"jobType,omitempty"`
	// The number of records processed in this job.
	NumberRecordsProcessed *int64 `json:"numberRecordsProcessed,omitempty"`
	// The number of times that Salesforce attempted to save the results of an operation. Repeated attempts indicate a problem such as a lock contention.
	Retries *int32 `json:"retries,omitempty"`
	// The number of milliseconds taken to process the job.
	TotalProcessingTime *int64 `json:"totalProcessingTime,omitempty"`
}

// NewBulkQueryJobInfoExt instantiates a new BulkQueryJobInfoExt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkQueryJobInfoExt() *BulkQueryJobInfoExt {
	this := BulkQueryJobInfoExt{}
	return &this
}

// NewBulkQueryJobInfoExtWithDefaults instantiates a new BulkQueryJobInfoExt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkQueryJobInfoExtWithDefaults() *BulkQueryJobInfoExt {
	this := BulkQueryJobInfoExt{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BulkQueryJobInfoExt) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQueryJobInfoExt) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BulkQueryJobInfoExt) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BulkQueryJobInfoExt) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *BulkQueryJobInfoExt) GetObject() string {
	if o == nil || isNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQueryJobInfoExt) GetObjectOk() (*string, bool) {
	if o == nil || isNil(o.Object) {
    return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *BulkQueryJobInfoExt) HasObject() bool {
	if o != nil && !isNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *BulkQueryJobInfoExt) SetObject(v string) {
	o.Object = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *BulkQueryJobInfoExt) GetCreatedById() string {
	if o == nil || isNil(o.CreatedById) {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQueryJobInfoExt) GetCreatedByIdOk() (*string, bool) {
	if o == nil || isNil(o.CreatedById) {
    return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *BulkQueryJobInfoExt) HasCreatedById() bool {
	if o != nil && !isNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *BulkQueryJobInfoExt) SetCreatedById(v string) {
	o.CreatedById = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *BulkQueryJobInfoExt) GetCreatedDate() string {
	if o == nil || isNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQueryJobInfoExt) GetCreatedDateOk() (*string, bool) {
	if o == nil || isNil(o.CreatedDate) {
    return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *BulkQueryJobInfoExt) HasCreatedDate() bool {
	if o != nil && !isNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *BulkQueryJobInfoExt) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetSystemModstamp returns the SystemModstamp field value if set, zero value otherwise.
func (o *BulkQueryJobInfoExt) GetSystemModstamp() string {
	if o == nil || isNil(o.SystemModstamp) {
		var ret string
		return ret
	}
	return *o.SystemModstamp
}

// GetSystemModstampOk returns a tuple with the SystemModstamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQueryJobInfoExt) GetSystemModstampOk() (*string, bool) {
	if o == nil || isNil(o.SystemModstamp) {
    return nil, false
	}
	return o.SystemModstamp, true
}

// HasSystemModstamp returns a boolean if a field has been set.
func (o *BulkQueryJobInfoExt) HasSystemModstamp() bool {
	if o != nil && !isNil(o.SystemModstamp) {
		return true
	}

	return false
}

// SetSystemModstamp gets a reference to the given string and assigns it to the SystemModstamp field.
func (o *BulkQueryJobInfoExt) SetSystemModstamp(v string) {
	o.SystemModstamp = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BulkQueryJobInfoExt) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQueryJobInfoExt) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BulkQueryJobInfoExt) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BulkQueryJobInfoExt) SetState(v string) {
	o.State = &v
}

// GetConcurrencyMode returns the ConcurrencyMode field value if set, zero value otherwise.
func (o *BulkQueryJobInfoExt) GetConcurrencyMode() string {
	if o == nil || isNil(o.ConcurrencyMode) {
		var ret string
		return ret
	}
	return *o.ConcurrencyMode
}

// GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQueryJobInfoExt) GetConcurrencyModeOk() (*string, bool) {
	if o == nil || isNil(o.ConcurrencyMode) {
    return nil, false
	}
	return o.ConcurrencyMode, true
}

// HasConcurrencyMode returns a boolean if a field has been set.
func (o *BulkQueryJobInfoExt) HasConcurrencyMode() bool {
	if o != nil && !isNil(o.ConcurrencyMode) {
		return true
	}

	return false
}

// SetConcurrencyMode gets a reference to the given string and assigns it to the ConcurrencyMode field.
func (o *BulkQueryJobInfoExt) SetConcurrencyMode(v string) {
	o.ConcurrencyMode = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *BulkQueryJobInfoExt) GetApiVersion() float32 {
	if o == nil || isNil(o.ApiVersion) {
		var ret float32
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQueryJobInfoExt) GetApiVersionOk() (*float32, bool) {
	if o == nil || isNil(o.ApiVersion) {
    return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *BulkQueryJobInfoExt) HasApiVersion() bool {
	if o != nil && !isNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given float32 and assigns it to the ApiVersion field.
func (o *BulkQueryJobInfoExt) SetApiVersion(v float32) {
	o.ApiVersion = &v
}

// GetJobType returns the JobType field value if set, zero value otherwise.
func (o *BulkQueryJobInfoExt) GetJobType() string {
	if o == nil || isNil(o.JobType) {
		var ret string
		return ret
	}
	return *o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQueryJobInfoExt) GetJobTypeOk() (*string, bool) {
	if o == nil || isNil(o.JobType) {
    return nil, false
	}
	return o.JobType, true
}

// HasJobType returns a boolean if a field has been set.
func (o *BulkQueryJobInfoExt) HasJobType() bool {
	if o != nil && !isNil(o.JobType) {
		return true
	}

	return false
}

// SetJobType gets a reference to the given string and assigns it to the JobType field.
func (o *BulkQueryJobInfoExt) SetJobType(v string) {
	o.JobType = &v
}

// GetNumberRecordsProcessed returns the NumberRecordsProcessed field value if set, zero value otherwise.
func (o *BulkQueryJobInfoExt) GetNumberRecordsProcessed() int64 {
	if o == nil || isNil(o.NumberRecordsProcessed) {
		var ret int64
		return ret
	}
	return *o.NumberRecordsProcessed
}

// GetNumberRecordsProcessedOk returns a tuple with the NumberRecordsProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQueryJobInfoExt) GetNumberRecordsProcessedOk() (*int64, bool) {
	if o == nil || isNil(o.NumberRecordsProcessed) {
    return nil, false
	}
	return o.NumberRecordsProcessed, true
}

// HasNumberRecordsProcessed returns a boolean if a field has been set.
func (o *BulkQueryJobInfoExt) HasNumberRecordsProcessed() bool {
	if o != nil && !isNil(o.NumberRecordsProcessed) {
		return true
	}

	return false
}

// SetNumberRecordsProcessed gets a reference to the given int64 and assigns it to the NumberRecordsProcessed field.
func (o *BulkQueryJobInfoExt) SetNumberRecordsProcessed(v int64) {
	o.NumberRecordsProcessed = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *BulkQueryJobInfoExt) GetRetries() int32 {
	if o == nil || isNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQueryJobInfoExt) GetRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.Retries) {
    return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *BulkQueryJobInfoExt) HasRetries() bool {
	if o != nil && !isNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *BulkQueryJobInfoExt) SetRetries(v int32) {
	o.Retries = &v
}

// GetTotalProcessingTime returns the TotalProcessingTime field value if set, zero value otherwise.
func (o *BulkQueryJobInfoExt) GetTotalProcessingTime() int64 {
	if o == nil || isNil(o.TotalProcessingTime) {
		var ret int64
		return ret
	}
	return *o.TotalProcessingTime
}

// GetTotalProcessingTimeOk returns a tuple with the TotalProcessingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQueryJobInfoExt) GetTotalProcessingTimeOk() (*int64, bool) {
	if o == nil || isNil(o.TotalProcessingTime) {
    return nil, false
	}
	return o.TotalProcessingTime, true
}

// HasTotalProcessingTime returns a boolean if a field has been set.
func (o *BulkQueryJobInfoExt) HasTotalProcessingTime() bool {
	if o != nil && !isNil(o.TotalProcessingTime) {
		return true
	}

	return false
}

// SetTotalProcessingTime gets a reference to the given int64 and assigns it to the TotalProcessingTime field.
func (o *BulkQueryJobInfoExt) SetTotalProcessingTime(v int64) {
	o.TotalProcessingTime = &v
}

func (o BulkQueryJobInfoExt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !isNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if !isNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !isNil(o.SystemModstamp) {
		toSerialize["systemModstamp"] = o.SystemModstamp
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.ConcurrencyMode) {
		toSerialize["concurrencyMode"] = o.ConcurrencyMode
	}
	if !isNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !isNil(o.JobType) {
		toSerialize["jobType"] = o.JobType
	}
	if !isNil(o.NumberRecordsProcessed) {
		toSerialize["numberRecordsProcessed"] = o.NumberRecordsProcessed
	}
	if !isNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	if !isNil(o.TotalProcessingTime) {
		toSerialize["totalProcessingTime"] = o.TotalProcessingTime
	}
	return json.Marshal(toSerialize)
}

type NullableBulkQueryJobInfoExt struct {
	value *BulkQueryJobInfoExt
	isSet bool
}

func (v NullableBulkQueryJobInfoExt) Get() *BulkQueryJobInfoExt {
	return v.value
}

func (v *NullableBulkQueryJobInfoExt) Set(val *BulkQueryJobInfoExt) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkQueryJobInfoExt) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkQueryJobInfoExt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkQueryJobInfoExt(val *BulkQueryJobInfoExt) *NullableBulkQueryJobInfoExt {
	return &NullableBulkQueryJobInfoExt{value: val, isSet: true}
}

func (v NullableBulkQueryJobInfoExt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkQueryJobInfoExt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

