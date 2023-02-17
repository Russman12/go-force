/*
Salesforce Platform REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 56.0
Contact: russell-laboe@outlook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tooling

import (
	"encoding/json"
)

// CodeCoverageResult struct for CodeCoverageResult
type CodeCoverageResult struct {
	DmlInfo *CodeLocation `json:"dmlInfo,omitempty"`
	Id string `json:"id"`
	LocationsNotCovered []CodeLocation `json:"locationsNotCovered,omitempty"`
	MethodInfo *CodeLocation `json:"methodInfo,omitempty"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	NumLocations int32 `json:"numLocations"`
	SoqlInfo *CodeLocation `json:"soqlInfo,omitempty"`
	SoslInfo *CodeLocation `json:"soslInfo,omitempty"`
	Type string `json:"type"`
}

// NewCodeCoverageResult instantiates a new CodeCoverageResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeCoverageResult(id string, name string, namespace string, numLocations int32, type_ string) *CodeCoverageResult {
	this := CodeCoverageResult{}
	this.Id = id
	this.Name = name
	this.Namespace = namespace
	this.NumLocations = numLocations
	this.Type = type_
	return &this
}

// NewCodeCoverageResultWithDefaults instantiates a new CodeCoverageResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeCoverageResultWithDefaults() *CodeCoverageResult {
	this := CodeCoverageResult{}
	return &this
}

// GetDmlInfo returns the DmlInfo field value if set, zero value otherwise.
func (o *CodeCoverageResult) GetDmlInfo() CodeLocation {
	if o == nil || isNil(o.DmlInfo) {
		var ret CodeLocation
		return ret
	}
	return *o.DmlInfo
}

// GetDmlInfoOk returns a tuple with the DmlInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeCoverageResult) GetDmlInfoOk() (*CodeLocation, bool) {
	if o == nil || isNil(o.DmlInfo) {
    return nil, false
	}
	return o.DmlInfo, true
}

// HasDmlInfo returns a boolean if a field has been set.
func (o *CodeCoverageResult) HasDmlInfo() bool {
	if o != nil && !isNil(o.DmlInfo) {
		return true
	}

	return false
}

// SetDmlInfo gets a reference to the given CodeLocation and assigns it to the DmlInfo field.
func (o *CodeCoverageResult) SetDmlInfo(v CodeLocation) {
	o.DmlInfo = &v
}

// GetId returns the Id field value
func (o *CodeCoverageResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CodeCoverageResult) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CodeCoverageResult) SetId(v string) {
	o.Id = v
}

// GetLocationsNotCovered returns the LocationsNotCovered field value if set, zero value otherwise.
func (o *CodeCoverageResult) GetLocationsNotCovered() []CodeLocation {
	if o == nil || isNil(o.LocationsNotCovered) {
		var ret []CodeLocation
		return ret
	}
	return o.LocationsNotCovered
}

// GetLocationsNotCoveredOk returns a tuple with the LocationsNotCovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeCoverageResult) GetLocationsNotCoveredOk() ([]CodeLocation, bool) {
	if o == nil || isNil(o.LocationsNotCovered) {
    return nil, false
	}
	return o.LocationsNotCovered, true
}

// HasLocationsNotCovered returns a boolean if a field has been set.
func (o *CodeCoverageResult) HasLocationsNotCovered() bool {
	if o != nil && !isNil(o.LocationsNotCovered) {
		return true
	}

	return false
}

// SetLocationsNotCovered gets a reference to the given []CodeLocation and assigns it to the LocationsNotCovered field.
func (o *CodeCoverageResult) SetLocationsNotCovered(v []CodeLocation) {
	o.LocationsNotCovered = v
}

// GetMethodInfo returns the MethodInfo field value if set, zero value otherwise.
func (o *CodeCoverageResult) GetMethodInfo() CodeLocation {
	if o == nil || isNil(o.MethodInfo) {
		var ret CodeLocation
		return ret
	}
	return *o.MethodInfo
}

// GetMethodInfoOk returns a tuple with the MethodInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeCoverageResult) GetMethodInfoOk() (*CodeLocation, bool) {
	if o == nil || isNil(o.MethodInfo) {
    return nil, false
	}
	return o.MethodInfo, true
}

// HasMethodInfo returns a boolean if a field has been set.
func (o *CodeCoverageResult) HasMethodInfo() bool {
	if o != nil && !isNil(o.MethodInfo) {
		return true
	}

	return false
}

// SetMethodInfo gets a reference to the given CodeLocation and assigns it to the MethodInfo field.
func (o *CodeCoverageResult) SetMethodInfo(v CodeLocation) {
	o.MethodInfo = &v
}

// GetName returns the Name field value
func (o *CodeCoverageResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CodeCoverageResult) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CodeCoverageResult) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *CodeCoverageResult) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *CodeCoverageResult) GetNamespaceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *CodeCoverageResult) SetNamespace(v string) {
	o.Namespace = v
}

// GetNumLocations returns the NumLocations field value
func (o *CodeCoverageResult) GetNumLocations() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumLocations
}

// GetNumLocationsOk returns a tuple with the NumLocations field value
// and a boolean to check if the value has been set.
func (o *CodeCoverageResult) GetNumLocationsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NumLocations, true
}

// SetNumLocations sets field value
func (o *CodeCoverageResult) SetNumLocations(v int32) {
	o.NumLocations = v
}

// GetSoqlInfo returns the SoqlInfo field value if set, zero value otherwise.
func (o *CodeCoverageResult) GetSoqlInfo() CodeLocation {
	if o == nil || isNil(o.SoqlInfo) {
		var ret CodeLocation
		return ret
	}
	return *o.SoqlInfo
}

// GetSoqlInfoOk returns a tuple with the SoqlInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeCoverageResult) GetSoqlInfoOk() (*CodeLocation, bool) {
	if o == nil || isNil(o.SoqlInfo) {
    return nil, false
	}
	return o.SoqlInfo, true
}

// HasSoqlInfo returns a boolean if a field has been set.
func (o *CodeCoverageResult) HasSoqlInfo() bool {
	if o != nil && !isNil(o.SoqlInfo) {
		return true
	}

	return false
}

// SetSoqlInfo gets a reference to the given CodeLocation and assigns it to the SoqlInfo field.
func (o *CodeCoverageResult) SetSoqlInfo(v CodeLocation) {
	o.SoqlInfo = &v
}

// GetSoslInfo returns the SoslInfo field value if set, zero value otherwise.
func (o *CodeCoverageResult) GetSoslInfo() CodeLocation {
	if o == nil || isNil(o.SoslInfo) {
		var ret CodeLocation
		return ret
	}
	return *o.SoslInfo
}

// GetSoslInfoOk returns a tuple with the SoslInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeCoverageResult) GetSoslInfoOk() (*CodeLocation, bool) {
	if o == nil || isNil(o.SoslInfo) {
    return nil, false
	}
	return o.SoslInfo, true
}

// HasSoslInfo returns a boolean if a field has been set.
func (o *CodeCoverageResult) HasSoslInfo() bool {
	if o != nil && !isNil(o.SoslInfo) {
		return true
	}

	return false
}

// SetSoslInfo gets a reference to the given CodeLocation and assigns it to the SoslInfo field.
func (o *CodeCoverageResult) SetSoslInfo(v CodeLocation) {
	o.SoslInfo = &v
}

// GetType returns the Type field value
func (o *CodeCoverageResult) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CodeCoverageResult) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CodeCoverageResult) SetType(v string) {
	o.Type = v
}

func (o CodeCoverageResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DmlInfo) {
		toSerialize["dmlInfo"] = o.DmlInfo
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.LocationsNotCovered) {
		toSerialize["locationsNotCovered"] = o.LocationsNotCovered
	}
	if !isNil(o.MethodInfo) {
		toSerialize["methodInfo"] = o.MethodInfo
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["namespace"] = o.Namespace
	}
	if true {
		toSerialize["numLocations"] = o.NumLocations
	}
	if !isNil(o.SoqlInfo) {
		toSerialize["soqlInfo"] = o.SoqlInfo
	}
	if !isNil(o.SoslInfo) {
		toSerialize["soslInfo"] = o.SoslInfo
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCodeCoverageResult struct {
	value *CodeCoverageResult
	isSet bool
}

func (v NullableCodeCoverageResult) Get() *CodeCoverageResult {
	return v.value
}

func (v *NullableCodeCoverageResult) Set(val *CodeCoverageResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeCoverageResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeCoverageResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeCoverageResult(val *CodeCoverageResult) *NullableCodeCoverageResult {
	return &NullableCodeCoverageResult{value: val, isSet: true}
}

func (v NullableCodeCoverageResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeCoverageResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

