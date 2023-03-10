/*
Salesforce Platform Bulk V2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 56.0
Contact: russell-laboe@outlook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bulkv2

import ()

// CreateJobRequest struct for CreateJobRequest
type CreateJobRequest struct {
	Job
}

// NewCreateJobRequest instantiates a new CreateJobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateJobRequest(object string, operation JobOperation) *CreateJobRequest {
	this := CreateJobRequest{}
	this.SetObject(object)
	this.SetOperation(operation)
	return &this
}

// NewCreateJobRequestWithDefaults instantiates a new CreateJobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateJobRequestWithDefaults() *CreateJobRequest {
	this := CreateJobRequest{}
	return &this
}
