/*
Salesforce Platform Bulk V2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 56.0
Contact: russell-laboe@outlook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bulkv2

import ()

// JobInfo struct for JobInfo
type JobInfo struct {
	Job
	JobInfoExt
}

// NewJobInfo instantiates a new JobInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobInfo() *JobInfo {
	this := JobInfo{}
	return &this
}

// NewJobInfoWithDefaults instantiates a new JobInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobInfoWithDefaults() *JobInfo {
	this := JobInfo{}
	return &this
}
