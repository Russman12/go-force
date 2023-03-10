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

// QueryOperation The type of query.
type QueryOperation string

// List of QueryOperation
const (
	QUERYOPERATION_QUERY     QueryOperation = "query"
	QUERYOPERATION_QUERY_ALL QueryOperation = "queryAll"
)

// All allowed values of QueryOperation enum
var AllowedQueryOperationEnumValues = []QueryOperation{
	"query",
	"queryAll",
}

func (v *QueryOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryOperation(value)
	for _, existing := range AllowedQueryOperationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryOperation", value)
}

// NewQueryOperationFromValue returns a pointer to a valid QueryOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryOperationFromValue(v string) (*QueryOperation, error) {
	ev := QueryOperation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryOperation: valid values are %v", v, AllowedQueryOperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryOperation) IsValid() bool {
	for _, existing := range AllowedQueryOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryOperation value
func (v QueryOperation) Ptr() *QueryOperation {
	return &v
}

type NullableQueryOperation struct {
	value *QueryOperation
	isSet bool
}

func (v NullableQueryOperation) Get() *QueryOperation {
	return v.value
}

func (v *NullableQueryOperation) Set(val *QueryOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryOperation(val *QueryOperation) *NullableQueryOperation {
	return &NullableQueryOperation{value: val, isSet: true}
}

func (v NullableQueryOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
