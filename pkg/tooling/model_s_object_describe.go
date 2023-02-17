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

// SObjectDescribe struct for SObjectDescribe
type SObjectDescribe struct {
	Activateable *bool `json:"activateable,omitempty"`
	AssociateEntityType *string `json:"associateEntityType,omitempty"`
	AssociateParentEntity *string `json:"associateParentEntity,omitempty"`
	Createable *bool `json:"createable,omitempty"`
	Custom *bool `json:"custom,omitempty"`
	CustomSetting *bool `json:"customSetting,omitempty"`
	DeepCloneable *bool `json:"deepCloneable,omitempty"`
	Deletable *bool `json:"deletable,omitempty"`
	DeprecatedAndHidden *bool `json:"deprecatedAndHidden,omitempty"`
	FeedEnabled *bool `json:"feedEnabled,omitempty"`
	HasSubtypes *bool `json:"hasSubtypes,omitempty"`
	IsInterface *bool `json:"isInterface,omitempty"`
	IsSubtype *bool `json:"isSubtype,omitempty"`
	KeyPrefix *string `json:"keyPrefix,omitempty"`
	Label *string `json:"label,omitempty"`
	LabelPlural *string `json:"labelPlural,omitempty"`
	Layoutable *bool `json:"layoutable,omitempty"`
	Mergeable *bool `json:"mergeable,omitempty"`
	MruEnabled *bool `json:"mruEnabled,omitempty"`
	Name *string `json:"name,omitempty"`
	Queryable *bool `json:"queryable,omitempty"`
	Replicateable *bool `json:"replicateable,omitempty"`
	Retrieveable *bool `json:"retrieveable,omitempty"`
	Searchable *bool `json:"searchable,omitempty"`
	Triggerable *bool `json:"triggerable,omitempty"`
	Undeletable *bool `json:"undeletable,omitempty"`
	Updateable *bool `json:"updateable,omitempty"`
	Urls *SObjectDescribeURL `json:"urls,omitempty"`
}

// NewSObjectDescribe instantiates a new SObjectDescribe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSObjectDescribe() *SObjectDescribe {
	this := SObjectDescribe{}
	return &this
}

// NewSObjectDescribeWithDefaults instantiates a new SObjectDescribe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSObjectDescribeWithDefaults() *SObjectDescribe {
	this := SObjectDescribe{}
	return &this
}

// GetActivateable returns the Activateable field value if set, zero value otherwise.
func (o *SObjectDescribe) GetActivateable() bool {
	if o == nil || isNil(o.Activateable) {
		var ret bool
		return ret
	}
	return *o.Activateable
}

// GetActivateableOk returns a tuple with the Activateable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetActivateableOk() (*bool, bool) {
	if o == nil || isNil(o.Activateable) {
    return nil, false
	}
	return o.Activateable, true
}

// HasActivateable returns a boolean if a field has been set.
func (o *SObjectDescribe) HasActivateable() bool {
	if o != nil && !isNil(o.Activateable) {
		return true
	}

	return false
}

// SetActivateable gets a reference to the given bool and assigns it to the Activateable field.
func (o *SObjectDescribe) SetActivateable(v bool) {
	o.Activateable = &v
}

// GetAssociateEntityType returns the AssociateEntityType field value if set, zero value otherwise.
func (o *SObjectDescribe) GetAssociateEntityType() string {
	if o == nil || isNil(o.AssociateEntityType) {
		var ret string
		return ret
	}
	return *o.AssociateEntityType
}

// GetAssociateEntityTypeOk returns a tuple with the AssociateEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetAssociateEntityTypeOk() (*string, bool) {
	if o == nil || isNil(o.AssociateEntityType) {
    return nil, false
	}
	return o.AssociateEntityType, true
}

// HasAssociateEntityType returns a boolean if a field has been set.
func (o *SObjectDescribe) HasAssociateEntityType() bool {
	if o != nil && !isNil(o.AssociateEntityType) {
		return true
	}

	return false
}

// SetAssociateEntityType gets a reference to the given string and assigns it to the AssociateEntityType field.
func (o *SObjectDescribe) SetAssociateEntityType(v string) {
	o.AssociateEntityType = &v
}

// GetAssociateParentEntity returns the AssociateParentEntity field value if set, zero value otherwise.
func (o *SObjectDescribe) GetAssociateParentEntity() string {
	if o == nil || isNil(o.AssociateParentEntity) {
		var ret string
		return ret
	}
	return *o.AssociateParentEntity
}

// GetAssociateParentEntityOk returns a tuple with the AssociateParentEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetAssociateParentEntityOk() (*string, bool) {
	if o == nil || isNil(o.AssociateParentEntity) {
    return nil, false
	}
	return o.AssociateParentEntity, true
}

// HasAssociateParentEntity returns a boolean if a field has been set.
func (o *SObjectDescribe) HasAssociateParentEntity() bool {
	if o != nil && !isNil(o.AssociateParentEntity) {
		return true
	}

	return false
}

// SetAssociateParentEntity gets a reference to the given string and assigns it to the AssociateParentEntity field.
func (o *SObjectDescribe) SetAssociateParentEntity(v string) {
	o.AssociateParentEntity = &v
}

// GetCreateable returns the Createable field value if set, zero value otherwise.
func (o *SObjectDescribe) GetCreateable() bool {
	if o == nil || isNil(o.Createable) {
		var ret bool
		return ret
	}
	return *o.Createable
}

// GetCreateableOk returns a tuple with the Createable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetCreateableOk() (*bool, bool) {
	if o == nil || isNil(o.Createable) {
    return nil, false
	}
	return o.Createable, true
}

// HasCreateable returns a boolean if a field has been set.
func (o *SObjectDescribe) HasCreateable() bool {
	if o != nil && !isNil(o.Createable) {
		return true
	}

	return false
}

// SetCreateable gets a reference to the given bool and assigns it to the Createable field.
func (o *SObjectDescribe) SetCreateable(v bool) {
	o.Createable = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *SObjectDescribe) GetCustom() bool {
	if o == nil || isNil(o.Custom) {
		var ret bool
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetCustomOk() (*bool, bool) {
	if o == nil || isNil(o.Custom) {
    return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *SObjectDescribe) HasCustom() bool {
	if o != nil && !isNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given bool and assigns it to the Custom field.
func (o *SObjectDescribe) SetCustom(v bool) {
	o.Custom = &v
}

// GetCustomSetting returns the CustomSetting field value if set, zero value otherwise.
func (o *SObjectDescribe) GetCustomSetting() bool {
	if o == nil || isNil(o.CustomSetting) {
		var ret bool
		return ret
	}
	return *o.CustomSetting
}

// GetCustomSettingOk returns a tuple with the CustomSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetCustomSettingOk() (*bool, bool) {
	if o == nil || isNil(o.CustomSetting) {
    return nil, false
	}
	return o.CustomSetting, true
}

// HasCustomSetting returns a boolean if a field has been set.
func (o *SObjectDescribe) HasCustomSetting() bool {
	if o != nil && !isNil(o.CustomSetting) {
		return true
	}

	return false
}

// SetCustomSetting gets a reference to the given bool and assigns it to the CustomSetting field.
func (o *SObjectDescribe) SetCustomSetting(v bool) {
	o.CustomSetting = &v
}

// GetDeepCloneable returns the DeepCloneable field value if set, zero value otherwise.
func (o *SObjectDescribe) GetDeepCloneable() bool {
	if o == nil || isNil(o.DeepCloneable) {
		var ret bool
		return ret
	}
	return *o.DeepCloneable
}

// GetDeepCloneableOk returns a tuple with the DeepCloneable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetDeepCloneableOk() (*bool, bool) {
	if o == nil || isNil(o.DeepCloneable) {
    return nil, false
	}
	return o.DeepCloneable, true
}

// HasDeepCloneable returns a boolean if a field has been set.
func (o *SObjectDescribe) HasDeepCloneable() bool {
	if o != nil && !isNil(o.DeepCloneable) {
		return true
	}

	return false
}

// SetDeepCloneable gets a reference to the given bool and assigns it to the DeepCloneable field.
func (o *SObjectDescribe) SetDeepCloneable(v bool) {
	o.DeepCloneable = &v
}

// GetDeletable returns the Deletable field value if set, zero value otherwise.
func (o *SObjectDescribe) GetDeletable() bool {
	if o == nil || isNil(o.Deletable) {
		var ret bool
		return ret
	}
	return *o.Deletable
}

// GetDeletableOk returns a tuple with the Deletable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetDeletableOk() (*bool, bool) {
	if o == nil || isNil(o.Deletable) {
    return nil, false
	}
	return o.Deletable, true
}

// HasDeletable returns a boolean if a field has been set.
func (o *SObjectDescribe) HasDeletable() bool {
	if o != nil && !isNil(o.Deletable) {
		return true
	}

	return false
}

// SetDeletable gets a reference to the given bool and assigns it to the Deletable field.
func (o *SObjectDescribe) SetDeletable(v bool) {
	o.Deletable = &v
}

// GetDeprecatedAndHidden returns the DeprecatedAndHidden field value if set, zero value otherwise.
func (o *SObjectDescribe) GetDeprecatedAndHidden() bool {
	if o == nil || isNil(o.DeprecatedAndHidden) {
		var ret bool
		return ret
	}
	return *o.DeprecatedAndHidden
}

// GetDeprecatedAndHiddenOk returns a tuple with the DeprecatedAndHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetDeprecatedAndHiddenOk() (*bool, bool) {
	if o == nil || isNil(o.DeprecatedAndHidden) {
    return nil, false
	}
	return o.DeprecatedAndHidden, true
}

// HasDeprecatedAndHidden returns a boolean if a field has been set.
func (o *SObjectDescribe) HasDeprecatedAndHidden() bool {
	if o != nil && !isNil(o.DeprecatedAndHidden) {
		return true
	}

	return false
}

// SetDeprecatedAndHidden gets a reference to the given bool and assigns it to the DeprecatedAndHidden field.
func (o *SObjectDescribe) SetDeprecatedAndHidden(v bool) {
	o.DeprecatedAndHidden = &v
}

// GetFeedEnabled returns the FeedEnabled field value if set, zero value otherwise.
func (o *SObjectDescribe) GetFeedEnabled() bool {
	if o == nil || isNil(o.FeedEnabled) {
		var ret bool
		return ret
	}
	return *o.FeedEnabled
}

// GetFeedEnabledOk returns a tuple with the FeedEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetFeedEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.FeedEnabled) {
    return nil, false
	}
	return o.FeedEnabled, true
}

// HasFeedEnabled returns a boolean if a field has been set.
func (o *SObjectDescribe) HasFeedEnabled() bool {
	if o != nil && !isNil(o.FeedEnabled) {
		return true
	}

	return false
}

// SetFeedEnabled gets a reference to the given bool and assigns it to the FeedEnabled field.
func (o *SObjectDescribe) SetFeedEnabled(v bool) {
	o.FeedEnabled = &v
}

// GetHasSubtypes returns the HasSubtypes field value if set, zero value otherwise.
func (o *SObjectDescribe) GetHasSubtypes() bool {
	if o == nil || isNil(o.HasSubtypes) {
		var ret bool
		return ret
	}
	return *o.HasSubtypes
}

// GetHasSubtypesOk returns a tuple with the HasSubtypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetHasSubtypesOk() (*bool, bool) {
	if o == nil || isNil(o.HasSubtypes) {
    return nil, false
	}
	return o.HasSubtypes, true
}

// HasHasSubtypes returns a boolean if a field has been set.
func (o *SObjectDescribe) HasHasSubtypes() bool {
	if o != nil && !isNil(o.HasSubtypes) {
		return true
	}

	return false
}

// SetHasSubtypes gets a reference to the given bool and assigns it to the HasSubtypes field.
func (o *SObjectDescribe) SetHasSubtypes(v bool) {
	o.HasSubtypes = &v
}

// GetIsInterface returns the IsInterface field value if set, zero value otherwise.
func (o *SObjectDescribe) GetIsInterface() bool {
	if o == nil || isNil(o.IsInterface) {
		var ret bool
		return ret
	}
	return *o.IsInterface
}

// GetIsInterfaceOk returns a tuple with the IsInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetIsInterfaceOk() (*bool, bool) {
	if o == nil || isNil(o.IsInterface) {
    return nil, false
	}
	return o.IsInterface, true
}

// HasIsInterface returns a boolean if a field has been set.
func (o *SObjectDescribe) HasIsInterface() bool {
	if o != nil && !isNil(o.IsInterface) {
		return true
	}

	return false
}

// SetIsInterface gets a reference to the given bool and assigns it to the IsInterface field.
func (o *SObjectDescribe) SetIsInterface(v bool) {
	o.IsInterface = &v
}

// GetIsSubtype returns the IsSubtype field value if set, zero value otherwise.
func (o *SObjectDescribe) GetIsSubtype() bool {
	if o == nil || isNil(o.IsSubtype) {
		var ret bool
		return ret
	}
	return *o.IsSubtype
}

// GetIsSubtypeOk returns a tuple with the IsSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetIsSubtypeOk() (*bool, bool) {
	if o == nil || isNil(o.IsSubtype) {
    return nil, false
	}
	return o.IsSubtype, true
}

// HasIsSubtype returns a boolean if a field has been set.
func (o *SObjectDescribe) HasIsSubtype() bool {
	if o != nil && !isNil(o.IsSubtype) {
		return true
	}

	return false
}

// SetIsSubtype gets a reference to the given bool and assigns it to the IsSubtype field.
func (o *SObjectDescribe) SetIsSubtype(v bool) {
	o.IsSubtype = &v
}

// GetKeyPrefix returns the KeyPrefix field value if set, zero value otherwise.
func (o *SObjectDescribe) GetKeyPrefix() string {
	if o == nil || isNil(o.KeyPrefix) {
		var ret string
		return ret
	}
	return *o.KeyPrefix
}

// GetKeyPrefixOk returns a tuple with the KeyPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetKeyPrefixOk() (*string, bool) {
	if o == nil || isNil(o.KeyPrefix) {
    return nil, false
	}
	return o.KeyPrefix, true
}

// HasKeyPrefix returns a boolean if a field has been set.
func (o *SObjectDescribe) HasKeyPrefix() bool {
	if o != nil && !isNil(o.KeyPrefix) {
		return true
	}

	return false
}

// SetKeyPrefix gets a reference to the given string and assigns it to the KeyPrefix field.
func (o *SObjectDescribe) SetKeyPrefix(v string) {
	o.KeyPrefix = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *SObjectDescribe) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *SObjectDescribe) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *SObjectDescribe) SetLabel(v string) {
	o.Label = &v
}

// GetLabelPlural returns the LabelPlural field value if set, zero value otherwise.
func (o *SObjectDescribe) GetLabelPlural() string {
	if o == nil || isNil(o.LabelPlural) {
		var ret string
		return ret
	}
	return *o.LabelPlural
}

// GetLabelPluralOk returns a tuple with the LabelPlural field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetLabelPluralOk() (*string, bool) {
	if o == nil || isNil(o.LabelPlural) {
    return nil, false
	}
	return o.LabelPlural, true
}

// HasLabelPlural returns a boolean if a field has been set.
func (o *SObjectDescribe) HasLabelPlural() bool {
	if o != nil && !isNil(o.LabelPlural) {
		return true
	}

	return false
}

// SetLabelPlural gets a reference to the given string and assigns it to the LabelPlural field.
func (o *SObjectDescribe) SetLabelPlural(v string) {
	o.LabelPlural = &v
}

// GetLayoutable returns the Layoutable field value if set, zero value otherwise.
func (o *SObjectDescribe) GetLayoutable() bool {
	if o == nil || isNil(o.Layoutable) {
		var ret bool
		return ret
	}
	return *o.Layoutable
}

// GetLayoutableOk returns a tuple with the Layoutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetLayoutableOk() (*bool, bool) {
	if o == nil || isNil(o.Layoutable) {
    return nil, false
	}
	return o.Layoutable, true
}

// HasLayoutable returns a boolean if a field has been set.
func (o *SObjectDescribe) HasLayoutable() bool {
	if o != nil && !isNil(o.Layoutable) {
		return true
	}

	return false
}

// SetLayoutable gets a reference to the given bool and assigns it to the Layoutable field.
func (o *SObjectDescribe) SetLayoutable(v bool) {
	o.Layoutable = &v
}

// GetMergeable returns the Mergeable field value if set, zero value otherwise.
func (o *SObjectDescribe) GetMergeable() bool {
	if o == nil || isNil(o.Mergeable) {
		var ret bool
		return ret
	}
	return *o.Mergeable
}

// GetMergeableOk returns a tuple with the Mergeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetMergeableOk() (*bool, bool) {
	if o == nil || isNil(o.Mergeable) {
    return nil, false
	}
	return o.Mergeable, true
}

// HasMergeable returns a boolean if a field has been set.
func (o *SObjectDescribe) HasMergeable() bool {
	if o != nil && !isNil(o.Mergeable) {
		return true
	}

	return false
}

// SetMergeable gets a reference to the given bool and assigns it to the Mergeable field.
func (o *SObjectDescribe) SetMergeable(v bool) {
	o.Mergeable = &v
}

// GetMruEnabled returns the MruEnabled field value if set, zero value otherwise.
func (o *SObjectDescribe) GetMruEnabled() bool {
	if o == nil || isNil(o.MruEnabled) {
		var ret bool
		return ret
	}
	return *o.MruEnabled
}

// GetMruEnabledOk returns a tuple with the MruEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetMruEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.MruEnabled) {
    return nil, false
	}
	return o.MruEnabled, true
}

// HasMruEnabled returns a boolean if a field has been set.
func (o *SObjectDescribe) HasMruEnabled() bool {
	if o != nil && !isNil(o.MruEnabled) {
		return true
	}

	return false
}

// SetMruEnabled gets a reference to the given bool and assigns it to the MruEnabled field.
func (o *SObjectDescribe) SetMruEnabled(v bool) {
	o.MruEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SObjectDescribe) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SObjectDescribe) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SObjectDescribe) SetName(v string) {
	o.Name = &v
}

// GetQueryable returns the Queryable field value if set, zero value otherwise.
func (o *SObjectDescribe) GetQueryable() bool {
	if o == nil || isNil(o.Queryable) {
		var ret bool
		return ret
	}
	return *o.Queryable
}

// GetQueryableOk returns a tuple with the Queryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetQueryableOk() (*bool, bool) {
	if o == nil || isNil(o.Queryable) {
    return nil, false
	}
	return o.Queryable, true
}

// HasQueryable returns a boolean if a field has been set.
func (o *SObjectDescribe) HasQueryable() bool {
	if o != nil && !isNil(o.Queryable) {
		return true
	}

	return false
}

// SetQueryable gets a reference to the given bool and assigns it to the Queryable field.
func (o *SObjectDescribe) SetQueryable(v bool) {
	o.Queryable = &v
}

// GetReplicateable returns the Replicateable field value if set, zero value otherwise.
func (o *SObjectDescribe) GetReplicateable() bool {
	if o == nil || isNil(o.Replicateable) {
		var ret bool
		return ret
	}
	return *o.Replicateable
}

// GetReplicateableOk returns a tuple with the Replicateable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetReplicateableOk() (*bool, bool) {
	if o == nil || isNil(o.Replicateable) {
    return nil, false
	}
	return o.Replicateable, true
}

// HasReplicateable returns a boolean if a field has been set.
func (o *SObjectDescribe) HasReplicateable() bool {
	if o != nil && !isNil(o.Replicateable) {
		return true
	}

	return false
}

// SetReplicateable gets a reference to the given bool and assigns it to the Replicateable field.
func (o *SObjectDescribe) SetReplicateable(v bool) {
	o.Replicateable = &v
}

// GetRetrieveable returns the Retrieveable field value if set, zero value otherwise.
func (o *SObjectDescribe) GetRetrieveable() bool {
	if o == nil || isNil(o.Retrieveable) {
		var ret bool
		return ret
	}
	return *o.Retrieveable
}

// GetRetrieveableOk returns a tuple with the Retrieveable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetRetrieveableOk() (*bool, bool) {
	if o == nil || isNil(o.Retrieveable) {
    return nil, false
	}
	return o.Retrieveable, true
}

// HasRetrieveable returns a boolean if a field has been set.
func (o *SObjectDescribe) HasRetrieveable() bool {
	if o != nil && !isNil(o.Retrieveable) {
		return true
	}

	return false
}

// SetRetrieveable gets a reference to the given bool and assigns it to the Retrieveable field.
func (o *SObjectDescribe) SetRetrieveable(v bool) {
	o.Retrieveable = &v
}

// GetSearchable returns the Searchable field value if set, zero value otherwise.
func (o *SObjectDescribe) GetSearchable() bool {
	if o == nil || isNil(o.Searchable) {
		var ret bool
		return ret
	}
	return *o.Searchable
}

// GetSearchableOk returns a tuple with the Searchable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetSearchableOk() (*bool, bool) {
	if o == nil || isNil(o.Searchable) {
    return nil, false
	}
	return o.Searchable, true
}

// HasSearchable returns a boolean if a field has been set.
func (o *SObjectDescribe) HasSearchable() bool {
	if o != nil && !isNil(o.Searchable) {
		return true
	}

	return false
}

// SetSearchable gets a reference to the given bool and assigns it to the Searchable field.
func (o *SObjectDescribe) SetSearchable(v bool) {
	o.Searchable = &v
}

// GetTriggerable returns the Triggerable field value if set, zero value otherwise.
func (o *SObjectDescribe) GetTriggerable() bool {
	if o == nil || isNil(o.Triggerable) {
		var ret bool
		return ret
	}
	return *o.Triggerable
}

// GetTriggerableOk returns a tuple with the Triggerable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetTriggerableOk() (*bool, bool) {
	if o == nil || isNil(o.Triggerable) {
    return nil, false
	}
	return o.Triggerable, true
}

// HasTriggerable returns a boolean if a field has been set.
func (o *SObjectDescribe) HasTriggerable() bool {
	if o != nil && !isNil(o.Triggerable) {
		return true
	}

	return false
}

// SetTriggerable gets a reference to the given bool and assigns it to the Triggerable field.
func (o *SObjectDescribe) SetTriggerable(v bool) {
	o.Triggerable = &v
}

// GetUndeletable returns the Undeletable field value if set, zero value otherwise.
func (o *SObjectDescribe) GetUndeletable() bool {
	if o == nil || isNil(o.Undeletable) {
		var ret bool
		return ret
	}
	return *o.Undeletable
}

// GetUndeletableOk returns a tuple with the Undeletable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetUndeletableOk() (*bool, bool) {
	if o == nil || isNil(o.Undeletable) {
    return nil, false
	}
	return o.Undeletable, true
}

// HasUndeletable returns a boolean if a field has been set.
func (o *SObjectDescribe) HasUndeletable() bool {
	if o != nil && !isNil(o.Undeletable) {
		return true
	}

	return false
}

// SetUndeletable gets a reference to the given bool and assigns it to the Undeletable field.
func (o *SObjectDescribe) SetUndeletable(v bool) {
	o.Undeletable = &v
}

// GetUpdateable returns the Updateable field value if set, zero value otherwise.
func (o *SObjectDescribe) GetUpdateable() bool {
	if o == nil || isNil(o.Updateable) {
		var ret bool
		return ret
	}
	return *o.Updateable
}

// GetUpdateableOk returns a tuple with the Updateable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetUpdateableOk() (*bool, bool) {
	if o == nil || isNil(o.Updateable) {
    return nil, false
	}
	return o.Updateable, true
}

// HasUpdateable returns a boolean if a field has been set.
func (o *SObjectDescribe) HasUpdateable() bool {
	if o != nil && !isNil(o.Updateable) {
		return true
	}

	return false
}

// SetUpdateable gets a reference to the given bool and assigns it to the Updateable field.
func (o *SObjectDescribe) SetUpdateable(v bool) {
	o.Updateable = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *SObjectDescribe) GetUrls() SObjectDescribeURL {
	if o == nil || isNil(o.Urls) {
		var ret SObjectDescribeURL
		return ret
	}
	return *o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SObjectDescribe) GetUrlsOk() (*SObjectDescribeURL, bool) {
	if o == nil || isNil(o.Urls) {
    return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *SObjectDescribe) HasUrls() bool {
	if o != nil && !isNil(o.Urls) {
		return true
	}

	return false
}

// SetUrls gets a reference to the given SObjectDescribeURL and assigns it to the Urls field.
func (o *SObjectDescribe) SetUrls(v SObjectDescribeURL) {
	o.Urls = &v
}

func (o SObjectDescribe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Activateable) {
		toSerialize["activateable"] = o.Activateable
	}
	if !isNil(o.AssociateEntityType) {
		toSerialize["associateEntityType"] = o.AssociateEntityType
	}
	if !isNil(o.AssociateParentEntity) {
		toSerialize["associateParentEntity"] = o.AssociateParentEntity
	}
	if !isNil(o.Createable) {
		toSerialize["createable"] = o.Createable
	}
	if !isNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if !isNil(o.CustomSetting) {
		toSerialize["customSetting"] = o.CustomSetting
	}
	if !isNil(o.DeepCloneable) {
		toSerialize["deepCloneable"] = o.DeepCloneable
	}
	if !isNil(o.Deletable) {
		toSerialize["deletable"] = o.Deletable
	}
	if !isNil(o.DeprecatedAndHidden) {
		toSerialize["deprecatedAndHidden"] = o.DeprecatedAndHidden
	}
	if !isNil(o.FeedEnabled) {
		toSerialize["feedEnabled"] = o.FeedEnabled
	}
	if !isNil(o.HasSubtypes) {
		toSerialize["hasSubtypes"] = o.HasSubtypes
	}
	if !isNil(o.IsInterface) {
		toSerialize["isInterface"] = o.IsInterface
	}
	if !isNil(o.IsSubtype) {
		toSerialize["isSubtype"] = o.IsSubtype
	}
	if !isNil(o.KeyPrefix) {
		toSerialize["keyPrefix"] = o.KeyPrefix
	}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.LabelPlural) {
		toSerialize["labelPlural"] = o.LabelPlural
	}
	if !isNil(o.Layoutable) {
		toSerialize["layoutable"] = o.Layoutable
	}
	if !isNil(o.Mergeable) {
		toSerialize["mergeable"] = o.Mergeable
	}
	if !isNil(o.MruEnabled) {
		toSerialize["mruEnabled"] = o.MruEnabled
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Queryable) {
		toSerialize["queryable"] = o.Queryable
	}
	if !isNil(o.Replicateable) {
		toSerialize["replicateable"] = o.Replicateable
	}
	if !isNil(o.Retrieveable) {
		toSerialize["retrieveable"] = o.Retrieveable
	}
	if !isNil(o.Searchable) {
		toSerialize["searchable"] = o.Searchable
	}
	if !isNil(o.Triggerable) {
		toSerialize["triggerable"] = o.Triggerable
	}
	if !isNil(o.Undeletable) {
		toSerialize["undeletable"] = o.Undeletable
	}
	if !isNil(o.Updateable) {
		toSerialize["updateable"] = o.Updateable
	}
	if !isNil(o.Urls) {
		toSerialize["urls"] = o.Urls
	}
	return json.Marshal(toSerialize)
}

type NullableSObjectDescribe struct {
	value *SObjectDescribe
	isSet bool
}

func (v NullableSObjectDescribe) Get() *SObjectDescribe {
	return v.value
}

func (v *NullableSObjectDescribe) Set(val *SObjectDescribe) {
	v.value = val
	v.isSet = true
}

func (v NullableSObjectDescribe) IsSet() bool {
	return v.isSet
}

func (v *NullableSObjectDescribe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSObjectDescribe(val *SObjectDescribe) *NullableSObjectDescribe {
	return &NullableSObjectDescribe{value: val, isSet: true}
}

func (v NullableSObjectDescribe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSObjectDescribe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

