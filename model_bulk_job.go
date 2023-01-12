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

// BulkJob struct for BulkJob
type BulkJob struct {
	// The ID of the assignment rule. This property is only shown if an assignment rule is specified when the job is created.
	AssignmentRuleId *string `json:"assignmentRuleId,omitempty"`
	// The column delimiter used for CSV job data.
	ColumnDelimiter *string `json:"columnDelimiter,omitempty"`
	// The format of the data being processed. Only CSV is supported.
	ContentType *string `json:"contentType,omitempty"`
	// The name of the external ID field for an upsert.
	ExternalIdFieldName *string `json:"externalIdFieldName,omitempty"`
	// The line ending used for CSV job data.
	LineEnding *string `json:"lineEnding,omitempty"`
	// The object type for the data being processed.
	Object string `json:"object"`
	// The processing operation for the job.
	Operation string `json:"operation"`
}

// NewBulkJob instantiates a new BulkJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkJob(object string, operation string) *BulkJob {
	this := BulkJob{}
	this.Object = object
	this.Operation = operation
	return &this
}

// NewBulkJobWithDefaults instantiates a new BulkJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkJobWithDefaults() *BulkJob {
	this := BulkJob{}
	return &this
}

// GetAssignmentRuleId returns the AssignmentRuleId field value if set, zero value otherwise.
func (o *BulkJob) GetAssignmentRuleId() string {
	if o == nil || isNil(o.AssignmentRuleId) {
		var ret string
		return ret
	}
	return *o.AssignmentRuleId
}

// GetAssignmentRuleIdOk returns a tuple with the AssignmentRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkJob) GetAssignmentRuleIdOk() (*string, bool) {
	if o == nil || isNil(o.AssignmentRuleId) {
		return nil, false
	}
	return o.AssignmentRuleId, true
}

// HasAssignmentRuleId returns a boolean if a field has been set.
func (o *BulkJob) HasAssignmentRuleId() bool {
	if o != nil && !isNil(o.AssignmentRuleId) {
		return true
	}

	return false
}

// SetAssignmentRuleId gets a reference to the given string and assigns it to the AssignmentRuleId field.
func (o *BulkJob) SetAssignmentRuleId(v string) {
	o.AssignmentRuleId = &v
}

// GetColumnDelimiter returns the ColumnDelimiter field value if set, zero value otherwise.
func (o *BulkJob) GetColumnDelimiter() string {
	if o == nil || isNil(o.ColumnDelimiter) {
		var ret string
		return ret
	}
	return *o.ColumnDelimiter
}

// GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkJob) GetColumnDelimiterOk() (*string, bool) {
	if o == nil || isNil(o.ColumnDelimiter) {
		return nil, false
	}
	return o.ColumnDelimiter, true
}

// HasColumnDelimiter returns a boolean if a field has been set.
func (o *BulkJob) HasColumnDelimiter() bool {
	if o != nil && !isNil(o.ColumnDelimiter) {
		return true
	}

	return false
}

// SetColumnDelimiter gets a reference to the given string and assigns it to the ColumnDelimiter field.
func (o *BulkJob) SetColumnDelimiter(v string) {
	o.ColumnDelimiter = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *BulkJob) GetContentType() string {
	if o == nil || isNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkJob) GetContentTypeOk() (*string, bool) {
	if o == nil || isNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *BulkJob) HasContentType() bool {
	if o != nil && !isNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *BulkJob) SetContentType(v string) {
	o.ContentType = &v
}

// GetExternalIdFieldName returns the ExternalIdFieldName field value if set, zero value otherwise.
func (o *BulkJob) GetExternalIdFieldName() string {
	if o == nil || isNil(o.ExternalIdFieldName) {
		var ret string
		return ret
	}
	return *o.ExternalIdFieldName
}

// GetExternalIdFieldNameOk returns a tuple with the ExternalIdFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkJob) GetExternalIdFieldNameOk() (*string, bool) {
	if o == nil || isNil(o.ExternalIdFieldName) {
		return nil, false
	}
	return o.ExternalIdFieldName, true
}

// HasExternalIdFieldName returns a boolean if a field has been set.
func (o *BulkJob) HasExternalIdFieldName() bool {
	if o != nil && !isNil(o.ExternalIdFieldName) {
		return true
	}

	return false
}

// SetExternalIdFieldName gets a reference to the given string and assigns it to the ExternalIdFieldName field.
func (o *BulkJob) SetExternalIdFieldName(v string) {
	o.ExternalIdFieldName = &v
}

// GetLineEnding returns the LineEnding field value if set, zero value otherwise.
func (o *BulkJob) GetLineEnding() string {
	if o == nil || isNil(o.LineEnding) {
		var ret string
		return ret
	}
	return *o.LineEnding
}

// GetLineEndingOk returns a tuple with the LineEnding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkJob) GetLineEndingOk() (*string, bool) {
	if o == nil || isNil(o.LineEnding) {
		return nil, false
	}
	return o.LineEnding, true
}

// HasLineEnding returns a boolean if a field has been set.
func (o *BulkJob) HasLineEnding() bool {
	if o != nil && !isNil(o.LineEnding) {
		return true
	}

	return false
}

// SetLineEnding gets a reference to the given string and assigns it to the LineEnding field.
func (o *BulkJob) SetLineEnding(v string) {
	o.LineEnding = &v
}

// GetObject returns the Object field value
func (o *BulkJob) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *BulkJob) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *BulkJob) SetObject(v string) {
	o.Object = v
}

// GetOperation returns the Operation field value
func (o *BulkJob) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *BulkJob) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *BulkJob) SetOperation(v string) {
	o.Operation = v
}

func (o BulkJob) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AssignmentRuleId) {
		toSerialize["assignmentRuleId"] = o.AssignmentRuleId
	}
	if !isNil(o.ColumnDelimiter) {
		toSerialize["columnDelimiter"] = o.ColumnDelimiter
	}
	if !isNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	if !isNil(o.ExternalIdFieldName) {
		toSerialize["externalIdFieldName"] = o.ExternalIdFieldName
	}
	if !isNil(o.LineEnding) {
		toSerialize["lineEnding"] = o.LineEnding
	}
	if true {
		toSerialize["object"] = o.Object
	}
	if true {
		toSerialize["operation"] = o.Operation
	}
	return json.Marshal(toSerialize)
}

type NullableBulkJob struct {
	value *BulkJob
	isSet bool
}

func (v NullableBulkJob) Get() *BulkJob {
	return v.value
}

func (v *NullableBulkJob) Set(val *BulkJob) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkJob) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkJob(val *BulkJob) *NullableBulkJob {
	return &NullableBulkJob{value: val, isSet: true}
}

func (v NullableBulkJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
