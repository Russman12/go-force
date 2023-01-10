/*
Salesforce Platform APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 56.0
Contact: russell-laboe@outlook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"

)

// SObjectRecordTypeInfo struct for SObjectRecordTypeInfo
type SObjectRecordTypeInfo struct {
	Active *bool `json:"active,omitempty"`
	Available *bool `json:"available,omitempty"`
	DefaultRecordTypeMapping *bool `json:"defaultRecordTypeMapping,omitempty"`
	DeveloperName *string `json:"developerName,omitempty"`
	Master *bool `json:"master,omitempty"`
	Name *string `json:"name,omitempty"`
	RecordTypeId *string `json:"recordTypeId,omitempty"`
	Urls *SObjectRecordTypeInfoUrls `json:"urls,omitempty"`
}

// NewSObjectRecordTypeInfo instantiates a new SObjectRecordTypeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSObjectRecordTypeInfo() *SObjectRecordTypeInfo {
	this := SObjectRecordTypeInfo{}
	return &this
}

// NewSObjectRecordTypeInfoWithDefaults instantiates a new SObjectRecordTypeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSObjectRecordTypeInfoWithDefaults() *SObjectRecordTypeInfo {
	this := SObjectRecordTypeInfo{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SObjectRecordTypeInfo) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectRecordTypeInfo) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SObjectRecordTypeInfo) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *SObjectRecordTypeInfo) SetActive(v bool) {
	o.Active = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *SObjectRecordTypeInfo) GetAvailable() bool {
	if o == nil || isNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectRecordTypeInfo) GetAvailableOk() (*bool, bool) {
	if o == nil || isNil(o.Available) {
    return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *SObjectRecordTypeInfo) HasAvailable() bool {
	if o != nil && !isNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *SObjectRecordTypeInfo) SetAvailable(v bool) {
	o.Available = &v
}

// GetDefaultRecordTypeMapping returns the DefaultRecordTypeMapping field value if set, zero value otherwise.
func (o *SObjectRecordTypeInfo) GetDefaultRecordTypeMapping() bool {
	if o == nil || isNil(o.DefaultRecordTypeMapping) {
		var ret bool
		return ret
	}
	return *o.DefaultRecordTypeMapping
}

// GetDefaultRecordTypeMappingOk returns a tuple with the DefaultRecordTypeMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectRecordTypeInfo) GetDefaultRecordTypeMappingOk() (*bool, bool) {
	if o == nil || isNil(o.DefaultRecordTypeMapping) {
    return nil, false
	}
	return o.DefaultRecordTypeMapping, true
}

// HasDefaultRecordTypeMapping returns a boolean if a field has been set.
func (o *SObjectRecordTypeInfo) HasDefaultRecordTypeMapping() bool {
	if o != nil && !isNil(o.DefaultRecordTypeMapping) {
		return true
	}

	return false
}

// SetDefaultRecordTypeMapping gets a reference to the given bool and assigns it to the DefaultRecordTypeMapping field.
func (o *SObjectRecordTypeInfo) SetDefaultRecordTypeMapping(v bool) {
	o.DefaultRecordTypeMapping = &v
}

// GetDeveloperName returns the DeveloperName field value if set, zero value otherwise.
func (o *SObjectRecordTypeInfo) GetDeveloperName() string {
	if o == nil || isNil(o.DeveloperName) {
		var ret string
		return ret
	}
	return *o.DeveloperName
}

// GetDeveloperNameOk returns a tuple with the DeveloperName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectRecordTypeInfo) GetDeveloperNameOk() (*string, bool) {
	if o == nil || isNil(o.DeveloperName) {
    return nil, false
	}
	return o.DeveloperName, true
}

// HasDeveloperName returns a boolean if a field has been set.
func (o *SObjectRecordTypeInfo) HasDeveloperName() bool {
	if o != nil && !isNil(o.DeveloperName) {
		return true
	}

	return false
}

// SetDeveloperName gets a reference to the given string and assigns it to the DeveloperName field.
func (o *SObjectRecordTypeInfo) SetDeveloperName(v string) {
	o.DeveloperName = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *SObjectRecordTypeInfo) GetMaster() bool {
	if o == nil || isNil(o.Master) {
		var ret bool
		return ret
	}
	return *o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectRecordTypeInfo) GetMasterOk() (*bool, bool) {
	if o == nil || isNil(o.Master) {
    return nil, false
	}
	return o.Master, true
}

// HasMaster returns a boolean if a field has been set.
func (o *SObjectRecordTypeInfo) HasMaster() bool {
	if o != nil && !isNil(o.Master) {
		return true
	}

	return false
}

// SetMaster gets a reference to the given bool and assigns it to the Master field.
func (o *SObjectRecordTypeInfo) SetMaster(v bool) {
	o.Master = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SObjectRecordTypeInfo) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectRecordTypeInfo) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SObjectRecordTypeInfo) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SObjectRecordTypeInfo) SetName(v string) {
	o.Name = &v
}

// GetRecordTypeId returns the RecordTypeId field value if set, zero value otherwise.
func (o *SObjectRecordTypeInfo) GetRecordTypeId() string {
	if o == nil || isNil(o.RecordTypeId) {
		var ret string
		return ret
	}
	return *o.RecordTypeId
}

// GetRecordTypeIdOk returns a tuple with the RecordTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectRecordTypeInfo) GetRecordTypeIdOk() (*string, bool) {
	if o == nil || isNil(o.RecordTypeId) {
    return nil, false
	}
	return o.RecordTypeId, true
}

// HasRecordTypeId returns a boolean if a field has been set.
func (o *SObjectRecordTypeInfo) HasRecordTypeId() bool {
	if o != nil && !isNil(o.RecordTypeId) {
		return true
	}

	return false
}

// SetRecordTypeId gets a reference to the given string and assigns it to the RecordTypeId field.
func (o *SObjectRecordTypeInfo) SetRecordTypeId(v string) {
	o.RecordTypeId = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *SObjectRecordTypeInfo) GetUrls() SObjectRecordTypeInfoUrls {
	if o == nil || isNil(o.Urls) {
		var ret SObjectRecordTypeInfoUrls
		return ret
	}
	return *o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectRecordTypeInfo) GetUrlsOk() (*SObjectRecordTypeInfoUrls, bool) {
	if o == nil || isNil(o.Urls) {
    return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *SObjectRecordTypeInfo) HasUrls() bool {
	if o != nil && !isNil(o.Urls) {
		return true
	}

	return false
}

// SetUrls gets a reference to the given SObjectRecordTypeInfoUrls and assigns it to the Urls field.
func (o *SObjectRecordTypeInfo) SetUrls(v SObjectRecordTypeInfoUrls) {
	o.Urls = &v
}

func (o SObjectRecordTypeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !isNil(o.DefaultRecordTypeMapping) {
		toSerialize["defaultRecordTypeMapping"] = o.DefaultRecordTypeMapping
	}
	if !isNil(o.DeveloperName) {
		toSerialize["developerName"] = o.DeveloperName
	}
	if !isNil(o.Master) {
		toSerialize["master"] = o.Master
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.RecordTypeId) {
		toSerialize["recordTypeId"] = o.RecordTypeId
	}
	if !isNil(o.Urls) {
		toSerialize["urls"] = o.Urls
	}
	return json.Marshal(toSerialize)
}

type NullableSObjectRecordTypeInfo struct {
	value *SObjectRecordTypeInfo
	isSet bool
}

func (v NullableSObjectRecordTypeInfo) Get() *SObjectRecordTypeInfo {
	return v.value
}

func (v *NullableSObjectRecordTypeInfo) Set(val *SObjectRecordTypeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSObjectRecordTypeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSObjectRecordTypeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSObjectRecordTypeInfo(val *SObjectRecordTypeInfo) *NullableSObjectRecordTypeInfo {
	return &NullableSObjectRecordTypeInfo{value: val, isSet: true}
}

func (v NullableSObjectRecordTypeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSObjectRecordTypeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


