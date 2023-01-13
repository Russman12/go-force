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

// JobInfo struct for JobInfo
type JobInfo struct {
	// The ID of the assignment rule. This property is only shown if an assignment rule is specified when the job is created.
	AssignmentRuleId string `json:"assignmentRuleId"`
	// The column delimiter used for CSV job data.
	ColumnDelimiter string `json:"columnDelimiter"`
	// The format of the data being processed. Only CSV is supported.
	ContentType string `json:"contentType"`
	// The name of the external ID field for an upsert.
	ExternalIdFieldName string `json:"externalIdFieldName"`
	// The line ending used for CSV job data.
	LineEnding string `json:"lineEnding"`
	// The object type for the data being processed.
	Object string `json:"object"`
	// The processing operation for the job.
	Operation string `json:"operation"`
	// The API version that the job was created in.
	ApiVersion float32 `json:"apiVersion"`
	// For future use. How the request was processed. Currently only parallel mode is supported. (When other modes are added, the mode will be chosen automatically by the API and will not be user configurable.)
	ConcurrencyMode string `json:"concurrencyMode"`
	// The URL to use for Upload Job Data requests for this job. Only valid if the job is in Open state.
	ContentUrl string `json:"contentUrl"`
	// The ID of the user who created the job.
	CreatedById string `json:"createdById"`
	// The date and time in the UTC time zone when the job was created.
	CreatedDate string `json:"createdDate"`
	// Unique ID for this job.
	Id string `json:"id"`
	// The job’s type.
	JobType string `json:"jobType"`
	// Date and time in the UTC time zone when the job finished.
	SystemModstamp string `json:"systemModstamp"`
	// The current state of processing for the job.
	State string `json:"state"`
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

// NewJobInfo instantiates a new JobInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobInfo(assignmentRuleId string, columnDelimiter string, contentType string, externalIdFieldName string, lineEnding string, object string, operation string, apiVersion float32, concurrencyMode string, contentUrl string, createdById string, createdDate string, id string, jobType string, systemModstamp string, state string) *JobInfo {
	this := JobInfo{}
	this.AssignmentRuleId = assignmentRuleId
	this.ColumnDelimiter = columnDelimiter
	this.ContentType = contentType
	this.ExternalIdFieldName = externalIdFieldName
	this.LineEnding = lineEnding
	this.Object = object
	this.Operation = operation
	this.ApiVersion = apiVersion
	this.ConcurrencyMode = concurrencyMode
	this.ContentUrl = contentUrl
	this.CreatedById = createdById
	this.CreatedDate = createdDate
	this.Id = id
	this.JobType = jobType
	this.SystemModstamp = systemModstamp
	this.State = state
	return &this
}

// NewJobInfoWithDefaults instantiates a new JobInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobInfoWithDefaults() *JobInfo {
	this := JobInfo{}
	return &this
}

// GetAssignmentRuleId returns the AssignmentRuleId field value
func (o *JobInfo) GetAssignmentRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignmentRuleId
}

// GetAssignmentRuleIdOk returns a tuple with the AssignmentRuleId field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetAssignmentRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignmentRuleId, true
}

// SetAssignmentRuleId sets field value
func (o *JobInfo) SetAssignmentRuleId(v string) {
	o.AssignmentRuleId = v
}

// GetColumnDelimiter returns the ColumnDelimiter field value
func (o *JobInfo) GetColumnDelimiter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ColumnDelimiter
}

// GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetColumnDelimiterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnDelimiter, true
}

// SetColumnDelimiter sets field value
func (o *JobInfo) SetColumnDelimiter(v string) {
	o.ColumnDelimiter = v
}

// GetContentType returns the ContentType field value
func (o *JobInfo) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *JobInfo) SetContentType(v string) {
	o.ContentType = v
}

// GetExternalIdFieldName returns the ExternalIdFieldName field value
func (o *JobInfo) GetExternalIdFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalIdFieldName
}

// GetExternalIdFieldNameOk returns a tuple with the ExternalIdFieldName field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetExternalIdFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalIdFieldName, true
}

// SetExternalIdFieldName sets field value
func (o *JobInfo) SetExternalIdFieldName(v string) {
	o.ExternalIdFieldName = v
}

// GetLineEnding returns the LineEnding field value
func (o *JobInfo) GetLineEnding() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LineEnding
}

// GetLineEndingOk returns a tuple with the LineEnding field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetLineEndingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LineEnding, true
}

// SetLineEnding sets field value
func (o *JobInfo) SetLineEnding(v string) {
	o.LineEnding = v
}

// GetObject returns the Object field value
func (o *JobInfo) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *JobInfo) SetObject(v string) {
	o.Object = v
}

// GetOperation returns the Operation field value
func (o *JobInfo) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *JobInfo) SetOperation(v string) {
	o.Operation = v
}

// GetApiVersion returns the ApiVersion field value
func (o *JobInfo) GetApiVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetApiVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *JobInfo) SetApiVersion(v float32) {
	o.ApiVersion = v
}

// GetConcurrencyMode returns the ConcurrencyMode field value
func (o *JobInfo) GetConcurrencyMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConcurrencyMode
}

// GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetConcurrencyModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConcurrencyMode, true
}

// SetConcurrencyMode sets field value
func (o *JobInfo) SetConcurrencyMode(v string) {
	o.ConcurrencyMode = v
}

// GetContentUrl returns the ContentUrl field value
func (o *JobInfo) GetContentUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentUrl
}

// GetContentUrlOk returns a tuple with the ContentUrl field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetContentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentUrl, true
}

// SetContentUrl sets field value
func (o *JobInfo) SetContentUrl(v string) {
	o.ContentUrl = v
}

// GetCreatedById returns the CreatedById field value
func (o *JobInfo) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *JobInfo) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *JobInfo) GetCreatedDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetCreatedDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *JobInfo) SetCreatedDate(v string) {
	o.CreatedDate = v
}

// GetId returns the Id field value
func (o *JobInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobInfo) SetId(v string) {
	o.Id = v
}

// GetJobType returns the JobType field value
func (o *JobInfo) GetJobType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetJobTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobType, true
}

// SetJobType sets field value
func (o *JobInfo) SetJobType(v string) {
	o.JobType = v
}

// GetSystemModstamp returns the SystemModstamp field value
func (o *JobInfo) GetSystemModstamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemModstamp
}

// GetSystemModstampOk returns a tuple with the SystemModstamp field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetSystemModstampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemModstamp, true
}

// SetSystemModstamp sets field value
func (o *JobInfo) SetSystemModstamp(v string) {
	o.SystemModstamp = v
}

// GetState returns the State field value
func (o *JobInfo) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *JobInfo) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *JobInfo) SetState(v string) {
	o.State = v
}

// GetApexProcessingTime returns the ApexProcessingTime field value if set, zero value otherwise.
func (o *JobInfo) GetApexProcessingTime() int64 {
	if o == nil || isNil(o.ApexProcessingTime) {
		var ret int64
		return ret
	}
	return *o.ApexProcessingTime
}

// GetApexProcessingTimeOk returns a tuple with the ApexProcessingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfo) GetApexProcessingTimeOk() (*int64, bool) {
	if o == nil || isNil(o.ApexProcessingTime) {
		return nil, false
	}
	return o.ApexProcessingTime, true
}

// HasApexProcessingTime returns a boolean if a field has been set.
func (o *JobInfo) HasApexProcessingTime() bool {
	if o != nil && !isNil(o.ApexProcessingTime) {
		return true
	}

	return false
}

// SetApexProcessingTime gets a reference to the given int64 and assigns it to the ApexProcessingTime field.
func (o *JobInfo) SetApexProcessingTime(v int64) {
	o.ApexProcessingTime = &v
}

// GetApiActiveProcessingTime returns the ApiActiveProcessingTime field value if set, zero value otherwise.
func (o *JobInfo) GetApiActiveProcessingTime() int64 {
	if o == nil || isNil(o.ApiActiveProcessingTime) {
		var ret int64
		return ret
	}
	return *o.ApiActiveProcessingTime
}

// GetApiActiveProcessingTimeOk returns a tuple with the ApiActiveProcessingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfo) GetApiActiveProcessingTimeOk() (*int64, bool) {
	if o == nil || isNil(o.ApiActiveProcessingTime) {
		return nil, false
	}
	return o.ApiActiveProcessingTime, true
}

// HasApiActiveProcessingTime returns a boolean if a field has been set.
func (o *JobInfo) HasApiActiveProcessingTime() bool {
	if o != nil && !isNil(o.ApiActiveProcessingTime) {
		return true
	}

	return false
}

// SetApiActiveProcessingTime gets a reference to the given int64 and assigns it to the ApiActiveProcessingTime field.
func (o *JobInfo) SetApiActiveProcessingTime(v int64) {
	o.ApiActiveProcessingTime = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *JobInfo) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfo) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *JobInfo) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *JobInfo) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetNumberRecordsFailed returns the NumberRecordsFailed field value if set, zero value otherwise.
func (o *JobInfo) GetNumberRecordsFailed() int64 {
	if o == nil || isNil(o.NumberRecordsFailed) {
		var ret int64
		return ret
	}
	return *o.NumberRecordsFailed
}

// GetNumberRecordsFailedOk returns a tuple with the NumberRecordsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfo) GetNumberRecordsFailedOk() (*int64, bool) {
	if o == nil || isNil(o.NumberRecordsFailed) {
		return nil, false
	}
	return o.NumberRecordsFailed, true
}

// HasNumberRecordsFailed returns a boolean if a field has been set.
func (o *JobInfo) HasNumberRecordsFailed() bool {
	if o != nil && !isNil(o.NumberRecordsFailed) {
		return true
	}

	return false
}

// SetNumberRecordsFailed gets a reference to the given int64 and assigns it to the NumberRecordsFailed field.
func (o *JobInfo) SetNumberRecordsFailed(v int64) {
	o.NumberRecordsFailed = &v
}

// GetNumberRecordsProcessed returns the NumberRecordsProcessed field value if set, zero value otherwise.
func (o *JobInfo) GetNumberRecordsProcessed() int64 {
	if o == nil || isNil(o.NumberRecordsProcessed) {
		var ret int64
		return ret
	}
	return *o.NumberRecordsProcessed
}

// GetNumberRecordsProcessedOk returns a tuple with the NumberRecordsProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfo) GetNumberRecordsProcessedOk() (*int64, bool) {
	if o == nil || isNil(o.NumberRecordsProcessed) {
		return nil, false
	}
	return o.NumberRecordsProcessed, true
}

// HasNumberRecordsProcessed returns a boolean if a field has been set.
func (o *JobInfo) HasNumberRecordsProcessed() bool {
	if o != nil && !isNil(o.NumberRecordsProcessed) {
		return true
	}

	return false
}

// SetNumberRecordsProcessed gets a reference to the given int64 and assigns it to the NumberRecordsProcessed field.
func (o *JobInfo) SetNumberRecordsProcessed(v int64) {
	o.NumberRecordsProcessed = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *JobInfo) GetRetries() int32 {
	if o == nil || isNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfo) GetRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.Retries) {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *JobInfo) HasRetries() bool {
	if o != nil && !isNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *JobInfo) SetRetries(v int32) {
	o.Retries = &v
}

// GetTotalProcessingTime returns the TotalProcessingTime field value if set, zero value otherwise.
func (o *JobInfo) GetTotalProcessingTime() int64 {
	if o == nil || isNil(o.TotalProcessingTime) {
		var ret int64
		return ret
	}
	return *o.TotalProcessingTime
}

// GetTotalProcessingTimeOk returns a tuple with the TotalProcessingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInfo) GetTotalProcessingTimeOk() (*int64, bool) {
	if o == nil || isNil(o.TotalProcessingTime) {
		return nil, false
	}
	return o.TotalProcessingTime, true
}

// HasTotalProcessingTime returns a boolean if a field has been set.
func (o *JobInfo) HasTotalProcessingTime() bool {
	if o != nil && !isNil(o.TotalProcessingTime) {
		return true
	}

	return false
}

// SetTotalProcessingTime gets a reference to the given int64 and assigns it to the TotalProcessingTime field.
func (o *JobInfo) SetTotalProcessingTime(v int64) {
	o.TotalProcessingTime = &v
}

func (o JobInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assignmentRuleId"] = o.AssignmentRuleId
	}
	if true {
		toSerialize["columnDelimiter"] = o.ColumnDelimiter
	}
	if true {
		toSerialize["contentType"] = o.ContentType
	}
	if true {
		toSerialize["externalIdFieldName"] = o.ExternalIdFieldName
	}
	if true {
		toSerialize["lineEnding"] = o.LineEnding
	}
	if true {
		toSerialize["object"] = o.Object
	}
	if true {
		toSerialize["operation"] = o.Operation
	}
	if true {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if true {
		toSerialize["concurrencyMode"] = o.ConcurrencyMode
	}
	if true {
		toSerialize["contentUrl"] = o.ContentUrl
	}
	if true {
		toSerialize["createdById"] = o.CreatedById
	}
	if true {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["jobType"] = o.JobType
	}
	if true {
		toSerialize["systemModstamp"] = o.SystemModstamp
	}
	if true {
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

type NullableJobInfo struct {
	value *JobInfo
	isSet bool
}

func (v NullableJobInfo) Get() *JobInfo {
	return v.value
}

func (v *NullableJobInfo) Set(val *JobInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableJobInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableJobInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobInfo(val *JobInfo) *NullableJobInfo {
	return &NullableJobInfo{value: val, isSet: true}
}

func (v NullableJobInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
