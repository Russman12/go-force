# DescribeSObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activateable** | Pointer to **bool** |  | [optional] 
**AssociateEntityType** | Pointer to **string** |  | [optional] 
**AssociateParentEntity** | Pointer to **string** |  | [optional] 
**Createable** | Pointer to **bool** |  | [optional] 
**Custom** | Pointer to **bool** |  | [optional] 
**CustomSetting** | Pointer to **bool** |  | [optional] 
**DeepCloneable** | Pointer to **bool** |  | [optional] 
**Deletable** | Pointer to **bool** |  | [optional] 
**DeprecatedAndHidden** | Pointer to **bool** |  | [optional] 
**FeedEnabled** | Pointer to **bool** |  | [optional] 
**HasSubtypes** | Pointer to **bool** |  | [optional] 
**IsInterface** | Pointer to **bool** |  | [optional] 
**IsSubtype** | Pointer to **bool** |  | [optional] 
**KeyPrefix** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**LabelPlural** | Pointer to **string** |  | [optional] 
**Layoutable** | Pointer to **bool** |  | [optional] 
**Mergeable** | Pointer to **bool** |  | [optional] 
**MruEnabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Queryable** | Pointer to **bool** |  | [optional] 
**Replicateable** | Pointer to **bool** |  | [optional] 
**Retrieveable** | Pointer to **bool** |  | [optional] 
**Searchable** | Pointer to **bool** |  | [optional] 
**Triggerable** | Pointer to **bool** |  | [optional] 
**Undeletable** | Pointer to **bool** |  | [optional] 
**Updateable** | Pointer to **bool** |  | [optional] 
**Urls** | Pointer to [**DescribeSObjectURL**](DescribeSObjectURL.md) |  | [optional] 

## Methods

### NewDescribeSObject

`func NewDescribeSObject() *DescribeSObject`

NewDescribeSObject instantiates a new DescribeSObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeSObjectWithDefaults

`func NewDescribeSObjectWithDefaults() *DescribeSObject`

NewDescribeSObjectWithDefaults instantiates a new DescribeSObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivateable

`func (o *DescribeSObject) GetActivateable() bool`

GetActivateable returns the Activateable field if non-nil, zero value otherwise.

### GetActivateableOk

`func (o *DescribeSObject) GetActivateableOk() (*bool, bool)`

GetActivateableOk returns a tuple with the Activateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateable

`func (o *DescribeSObject) SetActivateable(v bool)`

SetActivateable sets Activateable field to given value.

### HasActivateable

`func (o *DescribeSObject) HasActivateable() bool`

HasActivateable returns a boolean if a field has been set.

### GetAssociateEntityType

`func (o *DescribeSObject) GetAssociateEntityType() string`

GetAssociateEntityType returns the AssociateEntityType field if non-nil, zero value otherwise.

### GetAssociateEntityTypeOk

`func (o *DescribeSObject) GetAssociateEntityTypeOk() (*string, bool)`

GetAssociateEntityTypeOk returns a tuple with the AssociateEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociateEntityType

`func (o *DescribeSObject) SetAssociateEntityType(v string)`

SetAssociateEntityType sets AssociateEntityType field to given value.

### HasAssociateEntityType

`func (o *DescribeSObject) HasAssociateEntityType() bool`

HasAssociateEntityType returns a boolean if a field has been set.

### GetAssociateParentEntity

`func (o *DescribeSObject) GetAssociateParentEntity() string`

GetAssociateParentEntity returns the AssociateParentEntity field if non-nil, zero value otherwise.

### GetAssociateParentEntityOk

`func (o *DescribeSObject) GetAssociateParentEntityOk() (*string, bool)`

GetAssociateParentEntityOk returns a tuple with the AssociateParentEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociateParentEntity

`func (o *DescribeSObject) SetAssociateParentEntity(v string)`

SetAssociateParentEntity sets AssociateParentEntity field to given value.

### HasAssociateParentEntity

`func (o *DescribeSObject) HasAssociateParentEntity() bool`

HasAssociateParentEntity returns a boolean if a field has been set.

### GetCreateable

`func (o *DescribeSObject) GetCreateable() bool`

GetCreateable returns the Createable field if non-nil, zero value otherwise.

### GetCreateableOk

`func (o *DescribeSObject) GetCreateableOk() (*bool, bool)`

GetCreateableOk returns a tuple with the Createable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateable

`func (o *DescribeSObject) SetCreateable(v bool)`

SetCreateable sets Createable field to given value.

### HasCreateable

`func (o *DescribeSObject) HasCreateable() bool`

HasCreateable returns a boolean if a field has been set.

### GetCustom

`func (o *DescribeSObject) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *DescribeSObject) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *DescribeSObject) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *DescribeSObject) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomSetting

`func (o *DescribeSObject) GetCustomSetting() bool`

GetCustomSetting returns the CustomSetting field if non-nil, zero value otherwise.

### GetCustomSettingOk

`func (o *DescribeSObject) GetCustomSettingOk() (*bool, bool)`

GetCustomSettingOk returns a tuple with the CustomSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSetting

`func (o *DescribeSObject) SetCustomSetting(v bool)`

SetCustomSetting sets CustomSetting field to given value.

### HasCustomSetting

`func (o *DescribeSObject) HasCustomSetting() bool`

HasCustomSetting returns a boolean if a field has been set.

### GetDeepCloneable

`func (o *DescribeSObject) GetDeepCloneable() bool`

GetDeepCloneable returns the DeepCloneable field if non-nil, zero value otherwise.

### GetDeepCloneableOk

`func (o *DescribeSObject) GetDeepCloneableOk() (*bool, bool)`

GetDeepCloneableOk returns a tuple with the DeepCloneable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepCloneable

`func (o *DescribeSObject) SetDeepCloneable(v bool)`

SetDeepCloneable sets DeepCloneable field to given value.

### HasDeepCloneable

`func (o *DescribeSObject) HasDeepCloneable() bool`

HasDeepCloneable returns a boolean if a field has been set.

### GetDeletable

`func (o *DescribeSObject) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *DescribeSObject) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *DescribeSObject) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *DescribeSObject) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDeprecatedAndHidden

`func (o *DescribeSObject) GetDeprecatedAndHidden() bool`

GetDeprecatedAndHidden returns the DeprecatedAndHidden field if non-nil, zero value otherwise.

### GetDeprecatedAndHiddenOk

`func (o *DescribeSObject) GetDeprecatedAndHiddenOk() (*bool, bool)`

GetDeprecatedAndHiddenOk returns a tuple with the DeprecatedAndHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedAndHidden

`func (o *DescribeSObject) SetDeprecatedAndHidden(v bool)`

SetDeprecatedAndHidden sets DeprecatedAndHidden field to given value.

### HasDeprecatedAndHidden

`func (o *DescribeSObject) HasDeprecatedAndHidden() bool`

HasDeprecatedAndHidden returns a boolean if a field has been set.

### GetFeedEnabled

`func (o *DescribeSObject) GetFeedEnabled() bool`

GetFeedEnabled returns the FeedEnabled field if non-nil, zero value otherwise.

### GetFeedEnabledOk

`func (o *DescribeSObject) GetFeedEnabledOk() (*bool, bool)`

GetFeedEnabledOk returns a tuple with the FeedEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedEnabled

`func (o *DescribeSObject) SetFeedEnabled(v bool)`

SetFeedEnabled sets FeedEnabled field to given value.

### HasFeedEnabled

`func (o *DescribeSObject) HasFeedEnabled() bool`

HasFeedEnabled returns a boolean if a field has been set.

### GetHasSubtypes

`func (o *DescribeSObject) GetHasSubtypes() bool`

GetHasSubtypes returns the HasSubtypes field if non-nil, zero value otherwise.

### GetHasSubtypesOk

`func (o *DescribeSObject) GetHasSubtypesOk() (*bool, bool)`

GetHasSubtypesOk returns a tuple with the HasSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSubtypes

`func (o *DescribeSObject) SetHasSubtypes(v bool)`

SetHasSubtypes sets HasSubtypes field to given value.

### HasHasSubtypes

`func (o *DescribeSObject) HasHasSubtypes() bool`

HasHasSubtypes returns a boolean if a field has been set.

### GetIsInterface

`func (o *DescribeSObject) GetIsInterface() bool`

GetIsInterface returns the IsInterface field if non-nil, zero value otherwise.

### GetIsInterfaceOk

`func (o *DescribeSObject) GetIsInterfaceOk() (*bool, bool)`

GetIsInterfaceOk returns a tuple with the IsInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInterface

`func (o *DescribeSObject) SetIsInterface(v bool)`

SetIsInterface sets IsInterface field to given value.

### HasIsInterface

`func (o *DescribeSObject) HasIsInterface() bool`

HasIsInterface returns a boolean if a field has been set.

### GetIsSubtype

`func (o *DescribeSObject) GetIsSubtype() bool`

GetIsSubtype returns the IsSubtype field if non-nil, zero value otherwise.

### GetIsSubtypeOk

`func (o *DescribeSObject) GetIsSubtypeOk() (*bool, bool)`

GetIsSubtypeOk returns a tuple with the IsSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubtype

`func (o *DescribeSObject) SetIsSubtype(v bool)`

SetIsSubtype sets IsSubtype field to given value.

### HasIsSubtype

`func (o *DescribeSObject) HasIsSubtype() bool`

HasIsSubtype returns a boolean if a field has been set.

### GetKeyPrefix

`func (o *DescribeSObject) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *DescribeSObject) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *DescribeSObject) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.

### HasKeyPrefix

`func (o *DescribeSObject) HasKeyPrefix() bool`

HasKeyPrefix returns a boolean if a field has been set.

### GetLabel

`func (o *DescribeSObject) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DescribeSObject) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DescribeSObject) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DescribeSObject) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLabelPlural

`func (o *DescribeSObject) GetLabelPlural() string`

GetLabelPlural returns the LabelPlural field if non-nil, zero value otherwise.

### GetLabelPluralOk

`func (o *DescribeSObject) GetLabelPluralOk() (*string, bool)`

GetLabelPluralOk returns a tuple with the LabelPlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPlural

`func (o *DescribeSObject) SetLabelPlural(v string)`

SetLabelPlural sets LabelPlural field to given value.

### HasLabelPlural

`func (o *DescribeSObject) HasLabelPlural() bool`

HasLabelPlural returns a boolean if a field has been set.

### GetLayoutable

`func (o *DescribeSObject) GetLayoutable() bool`

GetLayoutable returns the Layoutable field if non-nil, zero value otherwise.

### GetLayoutableOk

`func (o *DescribeSObject) GetLayoutableOk() (*bool, bool)`

GetLayoutableOk returns a tuple with the Layoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutable

`func (o *DescribeSObject) SetLayoutable(v bool)`

SetLayoutable sets Layoutable field to given value.

### HasLayoutable

`func (o *DescribeSObject) HasLayoutable() bool`

HasLayoutable returns a boolean if a field has been set.

### GetMergeable

`func (o *DescribeSObject) GetMergeable() bool`

GetMergeable returns the Mergeable field if non-nil, zero value otherwise.

### GetMergeableOk

`func (o *DescribeSObject) GetMergeableOk() (*bool, bool)`

GetMergeableOk returns a tuple with the Mergeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeable

`func (o *DescribeSObject) SetMergeable(v bool)`

SetMergeable sets Mergeable field to given value.

### HasMergeable

`func (o *DescribeSObject) HasMergeable() bool`

HasMergeable returns a boolean if a field has been set.

### GetMruEnabled

`func (o *DescribeSObject) GetMruEnabled() bool`

GetMruEnabled returns the MruEnabled field if non-nil, zero value otherwise.

### GetMruEnabledOk

`func (o *DescribeSObject) GetMruEnabledOk() (*bool, bool)`

GetMruEnabledOk returns a tuple with the MruEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMruEnabled

`func (o *DescribeSObject) SetMruEnabled(v bool)`

SetMruEnabled sets MruEnabled field to given value.

### HasMruEnabled

`func (o *DescribeSObject) HasMruEnabled() bool`

HasMruEnabled returns a boolean if a field has been set.

### GetName

`func (o *DescribeSObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeSObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeSObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DescribeSObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQueryable

`func (o *DescribeSObject) GetQueryable() bool`

GetQueryable returns the Queryable field if non-nil, zero value otherwise.

### GetQueryableOk

`func (o *DescribeSObject) GetQueryableOk() (*bool, bool)`

GetQueryableOk returns a tuple with the Queryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryable

`func (o *DescribeSObject) SetQueryable(v bool)`

SetQueryable sets Queryable field to given value.

### HasQueryable

`func (o *DescribeSObject) HasQueryable() bool`

HasQueryable returns a boolean if a field has been set.

### GetReplicateable

`func (o *DescribeSObject) GetReplicateable() bool`

GetReplicateable returns the Replicateable field if non-nil, zero value otherwise.

### GetReplicateableOk

`func (o *DescribeSObject) GetReplicateableOk() (*bool, bool)`

GetReplicateableOk returns a tuple with the Replicateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateable

`func (o *DescribeSObject) SetReplicateable(v bool)`

SetReplicateable sets Replicateable field to given value.

### HasReplicateable

`func (o *DescribeSObject) HasReplicateable() bool`

HasReplicateable returns a boolean if a field has been set.

### GetRetrieveable

`func (o *DescribeSObject) GetRetrieveable() bool`

GetRetrieveable returns the Retrieveable field if non-nil, zero value otherwise.

### GetRetrieveableOk

`func (o *DescribeSObject) GetRetrieveableOk() (*bool, bool)`

GetRetrieveableOk returns a tuple with the Retrieveable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveable

`func (o *DescribeSObject) SetRetrieveable(v bool)`

SetRetrieveable sets Retrieveable field to given value.

### HasRetrieveable

`func (o *DescribeSObject) HasRetrieveable() bool`

HasRetrieveable returns a boolean if a field has been set.

### GetSearchable

`func (o *DescribeSObject) GetSearchable() bool`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *DescribeSObject) GetSearchableOk() (*bool, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *DescribeSObject) SetSearchable(v bool)`

SetSearchable sets Searchable field to given value.

### HasSearchable

`func (o *DescribeSObject) HasSearchable() bool`

HasSearchable returns a boolean if a field has been set.

### GetTriggerable

`func (o *DescribeSObject) GetTriggerable() bool`

GetTriggerable returns the Triggerable field if non-nil, zero value otherwise.

### GetTriggerableOk

`func (o *DescribeSObject) GetTriggerableOk() (*bool, bool)`

GetTriggerableOk returns a tuple with the Triggerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerable

`func (o *DescribeSObject) SetTriggerable(v bool)`

SetTriggerable sets Triggerable field to given value.

### HasTriggerable

`func (o *DescribeSObject) HasTriggerable() bool`

HasTriggerable returns a boolean if a field has been set.

### GetUndeletable

`func (o *DescribeSObject) GetUndeletable() bool`

GetUndeletable returns the Undeletable field if non-nil, zero value otherwise.

### GetUndeletableOk

`func (o *DescribeSObject) GetUndeletableOk() (*bool, bool)`

GetUndeletableOk returns a tuple with the Undeletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndeletable

`func (o *DescribeSObject) SetUndeletable(v bool)`

SetUndeletable sets Undeletable field to given value.

### HasUndeletable

`func (o *DescribeSObject) HasUndeletable() bool`

HasUndeletable returns a boolean if a field has been set.

### GetUpdateable

`func (o *DescribeSObject) GetUpdateable() bool`

GetUpdateable returns the Updateable field if non-nil, zero value otherwise.

### GetUpdateableOk

`func (o *DescribeSObject) GetUpdateableOk() (*bool, bool)`

GetUpdateableOk returns a tuple with the Updateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateable

`func (o *DescribeSObject) SetUpdateable(v bool)`

SetUpdateable sets Updateable field to given value.

### HasUpdateable

`func (o *DescribeSObject) HasUpdateable() bool`

HasUpdateable returns a boolean if a field has been set.

### GetUrls

`func (o *DescribeSObject) GetUrls() DescribeSObjectURL`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *DescribeSObject) GetUrlsOk() (*DescribeSObjectURL, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *DescribeSObject) SetUrls(v DescribeSObjectURL)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *DescribeSObject) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


