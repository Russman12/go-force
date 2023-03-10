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

// ChildRelationship struct for ChildRelationship
type ChildRelationship struct {
	CascadeDelete *bool `json:"cascadeDelete,omitempty"`
	ChildSObject *string `json:"childSObject,omitempty"`
	DeprecatedAndHidden *bool `json:"deprecatedAndHidden,omitempty"`
	Field *string `json:"field,omitempty"`
	JunctionIdListNames []string `json:"junctionIdListNames,omitempty"`
	JunctionReferenceTo []string `json:"junctionReferenceTo,omitempty"`
	RelationshipName *string `json:"relationshipName,omitempty"`
	RestrictedDelete *bool `json:"restrictedDelete,omitempty"`
}

// NewChildRelationship instantiates a new ChildRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChildRelationship() *ChildRelationship {
	this := ChildRelationship{}
	return &this
}

// NewChildRelationshipWithDefaults instantiates a new ChildRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChildRelationshipWithDefaults() *ChildRelationship {
	this := ChildRelationship{}
	return &this
}

// GetCascadeDelete returns the CascadeDelete field value if set, zero value otherwise.
func (o *ChildRelationship) GetCascadeDelete() bool {
	if o == nil || isNil(o.CascadeDelete) {
		var ret bool
		return ret
	}
	return *o.CascadeDelete
}

// GetCascadeDeleteOk returns a tuple with the CascadeDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildRelationship) GetCascadeDeleteOk() (*bool, bool) {
	if o == nil || isNil(o.CascadeDelete) {
    return nil, false
	}
	return o.CascadeDelete, true
}

// HasCascadeDelete returns a boolean if a field has been set.
func (o *ChildRelationship) HasCascadeDelete() bool {
	if o != nil && !isNil(o.CascadeDelete) {
		return true
	}

	return false
}

// SetCascadeDelete gets a reference to the given bool and assigns it to the CascadeDelete field.
func (o *ChildRelationship) SetCascadeDelete(v bool) {
	o.CascadeDelete = &v
}

// GetChildSObject returns the ChildSObject field value if set, zero value otherwise.
func (o *ChildRelationship) GetChildSObject() string {
	if o == nil || isNil(o.ChildSObject) {
		var ret string
		return ret
	}
	return *o.ChildSObject
}

// GetChildSObjectOk returns a tuple with the ChildSObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildRelationship) GetChildSObjectOk() (*string, bool) {
	if o == nil || isNil(o.ChildSObject) {
    return nil, false
	}
	return o.ChildSObject, true
}

// HasChildSObject returns a boolean if a field has been set.
func (o *ChildRelationship) HasChildSObject() bool {
	if o != nil && !isNil(o.ChildSObject) {
		return true
	}

	return false
}

// SetChildSObject gets a reference to the given string and assigns it to the ChildSObject field.
func (o *ChildRelationship) SetChildSObject(v string) {
	o.ChildSObject = &v
}

// GetDeprecatedAndHidden returns the DeprecatedAndHidden field value if set, zero value otherwise.
func (o *ChildRelationship) GetDeprecatedAndHidden() bool {
	if o == nil || isNil(o.DeprecatedAndHidden) {
		var ret bool
		return ret
	}
	return *o.DeprecatedAndHidden
}

// GetDeprecatedAndHiddenOk returns a tuple with the DeprecatedAndHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildRelationship) GetDeprecatedAndHiddenOk() (*bool, bool) {
	if o == nil || isNil(o.DeprecatedAndHidden) {
    return nil, false
	}
	return o.DeprecatedAndHidden, true
}

// HasDeprecatedAndHidden returns a boolean if a field has been set.
func (o *ChildRelationship) HasDeprecatedAndHidden() bool {
	if o != nil && !isNil(o.DeprecatedAndHidden) {
		return true
	}

	return false
}

// SetDeprecatedAndHidden gets a reference to the given bool and assigns it to the DeprecatedAndHidden field.
func (o *ChildRelationship) SetDeprecatedAndHidden(v bool) {
	o.DeprecatedAndHidden = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ChildRelationship) GetField() string {
	if o == nil || isNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildRelationship) GetFieldOk() (*string, bool) {
	if o == nil || isNil(o.Field) {
    return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ChildRelationship) HasField() bool {
	if o != nil && !isNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *ChildRelationship) SetField(v string) {
	o.Field = &v
}

// GetJunctionIdListNames returns the JunctionIdListNames field value if set, zero value otherwise.
func (o *ChildRelationship) GetJunctionIdListNames() []string {
	if o == nil || isNil(o.JunctionIdListNames) {
		var ret []string
		return ret
	}
	return o.JunctionIdListNames
}

// GetJunctionIdListNamesOk returns a tuple with the JunctionIdListNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildRelationship) GetJunctionIdListNamesOk() ([]string, bool) {
	if o == nil || isNil(o.JunctionIdListNames) {
    return nil, false
	}
	return o.JunctionIdListNames, true
}

// HasJunctionIdListNames returns a boolean if a field has been set.
func (o *ChildRelationship) HasJunctionIdListNames() bool {
	if o != nil && !isNil(o.JunctionIdListNames) {
		return true
	}

	return false
}

// SetJunctionIdListNames gets a reference to the given []string and assigns it to the JunctionIdListNames field.
func (o *ChildRelationship) SetJunctionIdListNames(v []string) {
	o.JunctionIdListNames = v
}

// GetJunctionReferenceTo returns the JunctionReferenceTo field value if set, zero value otherwise.
func (o *ChildRelationship) GetJunctionReferenceTo() []string {
	if o == nil || isNil(o.JunctionReferenceTo) {
		var ret []string
		return ret
	}
	return o.JunctionReferenceTo
}

// GetJunctionReferenceToOk returns a tuple with the JunctionReferenceTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildRelationship) GetJunctionReferenceToOk() ([]string, bool) {
	if o == nil || isNil(o.JunctionReferenceTo) {
    return nil, false
	}
	return o.JunctionReferenceTo, true
}

// HasJunctionReferenceTo returns a boolean if a field has been set.
func (o *ChildRelationship) HasJunctionReferenceTo() bool {
	if o != nil && !isNil(o.JunctionReferenceTo) {
		return true
	}

	return false
}

// SetJunctionReferenceTo gets a reference to the given []string and assigns it to the JunctionReferenceTo field.
func (o *ChildRelationship) SetJunctionReferenceTo(v []string) {
	o.JunctionReferenceTo = v
}

// GetRelationshipName returns the RelationshipName field value if set, zero value otherwise.
func (o *ChildRelationship) GetRelationshipName() string {
	if o == nil || isNil(o.RelationshipName) {
		var ret string
		return ret
	}
	return *o.RelationshipName
}

// GetRelationshipNameOk returns a tuple with the RelationshipName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildRelationship) GetRelationshipNameOk() (*string, bool) {
	if o == nil || isNil(o.RelationshipName) {
    return nil, false
	}
	return o.RelationshipName, true
}

// HasRelationshipName returns a boolean if a field has been set.
func (o *ChildRelationship) HasRelationshipName() bool {
	if o != nil && !isNil(o.RelationshipName) {
		return true
	}

	return false
}

// SetRelationshipName gets a reference to the given string and assigns it to the RelationshipName field.
func (o *ChildRelationship) SetRelationshipName(v string) {
	o.RelationshipName = &v
}

// GetRestrictedDelete returns the RestrictedDelete field value if set, zero value otherwise.
func (o *ChildRelationship) GetRestrictedDelete() bool {
	if o == nil || isNil(o.RestrictedDelete) {
		var ret bool
		return ret
	}
	return *o.RestrictedDelete
}

// GetRestrictedDeleteOk returns a tuple with the RestrictedDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildRelationship) GetRestrictedDeleteOk() (*bool, bool) {
	if o == nil || isNil(o.RestrictedDelete) {
    return nil, false
	}
	return o.RestrictedDelete, true
}

// HasRestrictedDelete returns a boolean if a field has been set.
func (o *ChildRelationship) HasRestrictedDelete() bool {
	if o != nil && !isNil(o.RestrictedDelete) {
		return true
	}

	return false
}

// SetRestrictedDelete gets a reference to the given bool and assigns it to the RestrictedDelete field.
func (o *ChildRelationship) SetRestrictedDelete(v bool) {
	o.RestrictedDelete = &v
}

func (o ChildRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CascadeDelete) {
		toSerialize["cascadeDelete"] = o.CascadeDelete
	}
	if !isNil(o.ChildSObject) {
		toSerialize["childSObject"] = o.ChildSObject
	}
	if !isNil(o.DeprecatedAndHidden) {
		toSerialize["deprecatedAndHidden"] = o.DeprecatedAndHidden
	}
	if !isNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !isNil(o.JunctionIdListNames) {
		toSerialize["junctionIdListNames"] = o.JunctionIdListNames
	}
	if !isNil(o.JunctionReferenceTo) {
		toSerialize["junctionReferenceTo"] = o.JunctionReferenceTo
	}
	if !isNil(o.RelationshipName) {
		toSerialize["relationshipName"] = o.RelationshipName
	}
	if !isNil(o.RestrictedDelete) {
		toSerialize["restrictedDelete"] = o.RestrictedDelete
	}
	return json.Marshal(toSerialize)
}

type NullableChildRelationship struct {
	value *ChildRelationship
	isSet bool
}

func (v NullableChildRelationship) Get() *ChildRelationship {
	return v.value
}

func (v *NullableChildRelationship) Set(val *ChildRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableChildRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableChildRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChildRelationship(val *ChildRelationship) *NullableChildRelationship {
	return &NullableChildRelationship{value: val, isSet: true}
}

func (v NullableChildRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChildRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


