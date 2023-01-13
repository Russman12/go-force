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

// JobType The job’s type.
type JobType string

// List of JobType
const (
	JOBTYPE_BIG_OBJECT_INGEST JobType = "BigObjectIngest"
	JOBTYPE_CLASSIC           JobType = "Classic"
	JOBTYPE_V2_INGEST         JobType = "V2Ingest"
)

// All allowed values of JobType enum
var AllowedJobTypeEnumValues = []JobType{
	"BigObjectIngest",
	"Classic",
	"V2Ingest",
}

func (v *JobType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JobType(value)
	for _, existing := range AllowedJobTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JobType", value)
}

// NewJobTypeFromValue returns a pointer to a valid JobType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJobTypeFromValue(v string) (*JobType, error) {
	ev := JobType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JobType: valid values are %v", v, AllowedJobTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JobType) IsValid() bool {
	for _, existing := range AllowedJobTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JobType value
func (v JobType) Ptr() *JobType {
	return &v
}

type NullableJobType struct {
	value *JobType
	isSet bool
}

func (v NullableJobType) Get() *JobType {
	return v.value
}

func (v *NullableJobType) Set(val *JobType) {
	v.value = val
	v.isSet = true
}

func (v NullableJobType) IsSet() bool {
	return v.isSet
}

func (v *NullableJobType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobType(val *JobType) *NullableJobType {
	return &NullableJobType{value: val, isSet: true}
}

func (v NullableJobType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
