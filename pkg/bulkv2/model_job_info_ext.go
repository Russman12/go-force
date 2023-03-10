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

// JobInfoExt struct for JobInfoExt
type JobInfoExt struct {
	// The API version that the job was created in.
	ApiVersion *float32 `json:"apiVersion,omitempty"`
	// For future use. How the request was processed. Currently only parallel mode is supported. (When other modes are added, the mode will be chosen automatically by the API and will not be user configurable.)
	ConcurrencyMode *string `json:"concurrencyMode,omitempty"`
	// The URL to use for Upload Job Data requests for this job. Only valid if the job is in Open state.
	ContentUrl *string `json:"contentUrl,omitempty"`
	// The ID of the user who created the job.
	CreatedById *string `json:"createdById,omitempty"`
	// The date and time in the UTC time zone when the job was created.
	CreatedDate *string `json:"createdDate,omitempty"`
	// Unique ID for this job.
	Id      *string  `json:"id,omitempty"`
	JobType *JobType `json:"jobType,omitempty"`
	// Date and time in the UTC time zone when the job finished.
	SystemModstamp *string   `json:"systemModstamp,omitempty"`
	State          *JobState `json:"state,omitempty"`
	// The number of milliseconds taken to process triggers and other processes related to the job data. This doesn't include the time used for processing asynchronous and batch Apex operations. If there are no triggers, the value is 0.
	ApexProcessingTime *int64 `json:"apexProcessingTime,omitempty"`
	// The number of milliseconds taken to actively process the job and includes apexProcessingTime, but doesn't include the time the job waited in the queue to be processed or the time required for serialization and deserialization.
	ApiActiveProcessingTime *int64 `json:"apiActiveProcessingTime,omitempty"`
	// The error message shown for jobs with errors.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// The number of records that were not processed successfully in this job.
	NumberRecordsFailed *int64 `json:"numberRecordsFailed,omitempty"`
	// The number of records already processed.
	NumberRecordsProcessed *int64 `json:"numberRecordsProcessed,omitempty"`
	// The number of times that Salesforce attempted to save the results of an operation. The repeated attempts are due to a problem, such as a lock contention.
	Retries *int32 `json:"retries,omitempty"`
	// The number of milliseconds taken to process the job.
	TotalProcessingTime *int64 `json:"totalProcessingTime,omitempty"`
}

// NewJobInfoExt instantiates a new JobInfoExt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobInfoExt() *JobInfoExt {
	this := JobInfoExt{}
	return &this
}

// NewJobInfoExtWithDefaults instantiates a new JobInfoExt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobInfoExtWithDefaults() *JobInfoExt {
	this := JobInfoExt{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *JobInfoExt) GetApiVersion() float32 {
	if o == nil || isNil(o.ApiVersion) {
		var ret float32
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetApiVersionOk() (*float32, bool) {
	if o == nil || isNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *JobInfoExt) HasApiVersion() bool {
	if o != nil && !isNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given float32 and assigns it to the ApiVersion field.
func (o *JobInfoExt) SetApiVersion(v float32) {
	o.ApiVersion = &v
}

// GetConcurrencyMode returns the ConcurrencyMode field value if set, zero value otherwise.
func (o *JobInfoExt) GetConcurrencyMode() string {
	if o == nil || isNil(o.ConcurrencyMode) {
		var ret string
		return ret
	}
	return *o.ConcurrencyMode
}

// GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetConcurrencyModeOk() (*string, bool) {
	if o == nil || isNil(o.ConcurrencyMode) {
		return nil, false
	}
	return o.ConcurrencyMode, true
}

// HasConcurrencyMode returns a boolean if a field has been set.
func (o *JobInfoExt) HasConcurrencyMode() bool {
	if o != nil && !isNil(o.ConcurrencyMode) {
		return true
	}

	return false
}

// SetConcurrencyMode gets a reference to the given string and assigns it to the ConcurrencyMode field.
func (o *JobInfoExt) SetConcurrencyMode(v string) {
	o.ConcurrencyMode = &v
}

// GetContentUrl returns the ContentUrl field value if set, zero value otherwise.
func (o *JobInfoExt) GetContentUrl() string {
	if o == nil || isNil(o.ContentUrl) {
		var ret string
		return ret
	}
	return *o.ContentUrl
}

// GetContentUrlOk returns a tuple with the ContentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetContentUrlOk() (*string, bool) {
	if o == nil || isNil(o.ContentUrl) {
		return nil, false
	}
	return o.ContentUrl, true
}

// HasContentUrl returns a boolean if a field has been set.
func (o *JobInfoExt) HasContentUrl() bool {
	if o != nil && !isNil(o.ContentUrl) {
		return true
	}

	return false
}

// SetContentUrl gets a reference to the given string and assigns it to the ContentUrl field.
func (o *JobInfoExt) SetContentUrl(v string) {
	o.ContentUrl = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *JobInfoExt) GetCreatedById() string {
	if o == nil || isNil(o.CreatedById) {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetCreatedByIdOk() (*string, bool) {
	if o == nil || isNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *JobInfoExt) HasCreatedById() bool {
	if o != nil && !isNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *JobInfoExt) SetCreatedById(v string) {
	o.CreatedById = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *JobInfoExt) GetCreatedDate() string {
	if o == nil || isNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetCreatedDateOk() (*string, bool) {
	if o == nil || isNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *JobInfoExt) HasCreatedDate() bool {
	if o != nil && !isNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *JobInfoExt) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JobInfoExt) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JobInfoExt) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JobInfoExt) SetId(v string) {
	o.Id = &v
}

// GetJobType returns the JobType field value if set, zero value otherwise.
func (o *JobInfoExt) GetJobType() JobType {
	if o == nil || isNil(o.JobType) {
		var ret JobType
		return ret
	}
	return *o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetJobTypeOk() (*JobType, bool) {
	if o == nil || isNil(o.JobType) {
		return nil, false
	}
	return o.JobType, true
}

// HasJobType returns a boolean if a field has been set.
func (o *JobInfoExt) HasJobType() bool {
	if o != nil && !isNil(o.JobType) {
		return true
	}

	return false
}

// SetJobType gets a reference to the given JobType and assigns it to the JobType field.
func (o *JobInfoExt) SetJobType(v JobType) {
	o.JobType = &v
}

// GetSystemModstamp returns the SystemModstamp field value if set, zero value otherwise.
func (o *JobInfoExt) GetSystemModstamp() string {
	if o == nil || isNil(o.SystemModstamp) {
		var ret string
		return ret
	}
	return *o.SystemModstamp
}

// GetSystemModstampOk returns a tuple with the SystemModstamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetSystemModstampOk() (*string, bool) {
	if o == nil || isNil(o.SystemModstamp) {
		return nil, false
	}
	return o.SystemModstamp, true
}

// HasSystemModstamp returns a boolean if a field has been set.
func (o *JobInfoExt) HasSystemModstamp() bool {
	if o != nil && !isNil(o.SystemModstamp) {
		return true
	}

	return false
}

// SetSystemModstamp gets a reference to the given string and assigns it to the SystemModstamp field.
func (o *JobInfoExt) SetSystemModstamp(v string) {
	o.SystemModstamp = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *JobInfoExt) GetState() JobState {
	if o == nil || isNil(o.State) {
		var ret JobState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetStateOk() (*JobState, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *JobInfoExt) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given JobState and assigns it to the State field.
func (o *JobInfoExt) SetState(v JobState) {
	o.State = &v
}

// GetApexProcessingTime returns the ApexProcessingTime field value if set, zero value otherwise.
func (o *JobInfoExt) GetApexProcessingTime() int64 {
	if o == nil || isNil(o.ApexProcessingTime) {
		var ret int64
		return ret
	}
	return *o.ApexProcessingTime
}

// GetApexProcessingTimeOk returns a tuple with the ApexProcessingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetApexProcessingTimeOk() (*int64, bool) {
	if o == nil || isNil(o.ApexProcessingTime) {
		return nil, false
	}
	return o.ApexProcessingTime, true
}

// HasApexProcessingTime returns a boolean if a field has been set.
func (o *JobInfoExt) HasApexProcessingTime() bool {
	if o != nil && !isNil(o.ApexProcessingTime) {
		return true
	}

	return false
}

// SetApexProcessingTime gets a reference to the given int64 and assigns it to the ApexProcessingTime field.
func (o *JobInfoExt) SetApexProcessingTime(v int64) {
	o.ApexProcessingTime = &v
}

// GetApiActiveProcessingTime returns the ApiActiveProcessingTime field value if set, zero value otherwise.
func (o *JobInfoExt) GetApiActiveProcessingTime() int64 {
	if o == nil || isNil(o.ApiActiveProcessingTime) {
		var ret int64
		return ret
	}
	return *o.ApiActiveProcessingTime
}

// GetApiActiveProcessingTimeOk returns a tuple with the ApiActiveProcessingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetApiActiveProcessingTimeOk() (*int64, bool) {
	if o == nil || isNil(o.ApiActiveProcessingTime) {
		return nil, false
	}
	return o.ApiActiveProcessingTime, true
}

// HasApiActiveProcessingTime returns a boolean if a field has been set.
func (o *JobInfoExt) HasApiActiveProcessingTime() bool {
	if o != nil && !isNil(o.ApiActiveProcessingTime) {
		return true
	}

	return false
}

// SetApiActiveProcessingTime gets a reference to the given int64 and assigns it to the ApiActiveProcessingTime field.
func (o *JobInfoExt) SetApiActiveProcessingTime(v int64) {
	o.ApiActiveProcessingTime = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *JobInfoExt) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *JobInfoExt) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *JobInfoExt) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetNumberRecordsFailed returns the NumberRecordsFailed field value if set, zero value otherwise.
func (o *JobInfoExt) GetNumberRecordsFailed() int64 {
	if o == nil || isNil(o.NumberRecordsFailed) {
		var ret int64
		return ret
	}
	return *o.NumberRecordsFailed
}

// GetNumberRecordsFailedOk returns a tuple with the NumberRecordsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetNumberRecordsFailedOk() (*int64, bool) {
	if o == nil || isNil(o.NumberRecordsFailed) {
		return nil, false
	}
	return o.NumberRecordsFailed, true
}

// HasNumberRecordsFailed returns a boolean if a field has been set.
func (o *JobInfoExt) HasNumberRecordsFailed() bool {
	if o != nil && !isNil(o.NumberRecordsFailed) {
		return true
	}

	return false
}

// SetNumberRecordsFailed gets a reference to the given int64 and assigns it to the NumberRecordsFailed field.
func (o *JobInfoExt) SetNumberRecordsFailed(v int64) {
	o.NumberRecordsFailed = &v
}

// GetNumberRecordsProcessed returns the NumberRecordsProcessed field value if set, zero value otherwise.
func (o *JobInfoExt) GetNumberRecordsProcessed() int64 {
	if o == nil || isNil(o.NumberRecordsProcessed) {
		var ret int64
		return ret
	}
	return *o.NumberRecordsProcessed
}

// GetNumberRecordsProcessedOk returns a tuple with the NumberRecordsProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetNumberRecordsProcessedOk() (*int64, bool) {
	if o == nil || isNil(o.NumberRecordsProcessed) {
		return nil, false
	}
	return o.NumberRecordsProcessed, true
}

// HasNumberRecordsProcessed returns a boolean if a field has been set.
func (o *JobInfoExt) HasNumberRecordsProcessed() bool {
	if o != nil && !isNil(o.NumberRecordsProcessed) {
		return true
	}

	return false
}

// SetNumberRecordsProcessed gets a reference to the given int64 and assigns it to the NumberRecordsProcessed field.
func (o *JobInfoExt) SetNumberRecordsProcessed(v int64) {
	o.NumberRecordsProcessed = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *JobInfoExt) GetRetries() int32 {
	if o == nil || isNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.Retries) {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *JobInfoExt) HasRetries() bool {
	if o != nil && !isNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *JobInfoExt) SetRetries(v int32) {
	o.Retries = &v
}

// GetTotalProcessingTime returns the TotalProcessingTime field value if set, zero value otherwise.
func (o *JobInfoExt) GetTotalProcessingTime() int64 {
	if o == nil || isNil(o.TotalProcessingTime) {
		var ret int64
		return ret
	}
	return *o.TotalProcessingTime
}

// GetTotalProcessingTimeOk returns a tuple with the TotalProcessingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfoExt) GetTotalProcessingTimeOk() (*int64, bool) {
	if o == nil || isNil(o.TotalProcessingTime) {
		return nil, false
	}
	return o.TotalProcessingTime, true
}

// HasTotalProcessingTime returns a boolean if a field has been set.
func (o *JobInfoExt) HasTotalProcessingTime() bool {
	if o != nil && !isNil(o.TotalProcessingTime) {
		return true
	}

	return false
}

// SetTotalProcessingTime gets a reference to the given int64 and assigns it to the TotalProcessingTime field.
func (o *JobInfoExt) SetTotalProcessingTime(v int64) {
	o.TotalProcessingTime = &v
}

func (o JobInfoExt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !isNil(o.ConcurrencyMode) {
		toSerialize["concurrencyMode"] = o.ConcurrencyMode
	}
	if !isNil(o.ContentUrl) {
		toSerialize["contentUrl"] = o.ContentUrl
	}
	if !isNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if !isNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.JobType) {
		toSerialize["jobType"] = o.JobType
	}
	if !isNil(o.SystemModstamp) {
		toSerialize["systemModstamp"] = o.SystemModstamp
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.ApexProcessingTime) {
		toSerialize["apexProcessingTime"] = o.ApexProcessingTime
	}
	if !isNil(o.ApiActiveProcessingTime) {
		toSerialize["apiActiveProcessingTime"] = o.ApiActiveProcessingTime
	}
	if !isNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !isNil(o.NumberRecordsFailed) {
		toSerialize["numberRecordsFailed"] = o.NumberRecordsFailed
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

type NullableJobInfoExt struct {
	value *JobInfoExt
	isSet bool
}

func (v NullableJobInfoExt) Get() *JobInfoExt {
	return v.value
}

func (v *NullableJobInfoExt) Set(val *JobInfoExt) {
	v.value = val
	v.isSet = true
}

func (v NullableJobInfoExt) IsSet() bool {
	return v.isSet
}

func (v *NullableJobInfoExt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobInfoExt(val *JobInfoExt) *NullableJobInfoExt {
	return &NullableJobInfoExt{value: val, isSet: true}
}

func (v NullableJobInfoExt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobInfoExt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
