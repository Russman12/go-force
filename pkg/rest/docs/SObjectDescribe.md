# SObjectDescribe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activateable** | Pointer to **bool** |  | [optional] 
**AssociateEntityType** | Pointer to **NullableString** |  | [optional] 
**AssociateParentEntity** | Pointer to **NullableString** |  | [optional] 
**CompactLayoutable** | Pointer to **bool** |  | [optional] 
**Createable** | Pointer to **bool** |  | [optional] 
**Custom** | Pointer to **bool** |  | [optional] 
**CustomSetting** | Pointer to **bool** |  | [optional] 
**DeepCloneable** | Pointer to **bool** |  | [optional] 
**DefaultImplementation** | Pointer to **NullableString** |  | [optional] 
**Deletable** | Pointer to **bool** |  | [optional] 
**DeprecatedAndHidden** | Pointer to **bool** |  | [optional] 
**ExtendedBy** | Pointer to **NullableString** |  | [optional] 
**ExtendsInterfaces** | Pointer to **NullableString** |  | [optional] 
**FeedEnabled** | Pointer to **bool** |  | [optional] 
**HasSubtypes** | Pointer to **bool** |  | [optional] 
**ImplementedBy** | Pointer to **NullableString** |  | [optional] 
**ImplementsInterfaces** | Pointer to **NullableString** |  | [optional] 
**IsInterface** | Pointer to **bool** |  | [optional] 
**IsSubtype** | Pointer to **bool** |  | [optional] 
**KeyPrefix** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**LabelPlural** | Pointer to **string** |  | [optional] 
**Layoutable** | Pointer to **bool** |  | [optional] 
**Listviewable** | Pointer to **NullableString** |  | [optional] 
**LookupLayoutable** | Pointer to **NullableString** |  | [optional] 
**Mergeable** | Pointer to **bool** |  | [optional] 
**MruEnabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NetworkScopeFieldName** | Pointer to **NullableString** |  | [optional] 
**Queryable** | Pointer to **bool** |  | [optional] 
**Replicateable** | Pointer to **bool** |  | [optional] 
**Retrieveable** | Pointer to **bool** |  | [optional] 
**SearchLayoutable** | Pointer to **bool** |  | [optional] 
**Searchable** | Pointer to **bool** |  | [optional] 
**SobjectDescribeOption** | Pointer to **string** |  | [optional] 
**Triggerable** | Pointer to **bool** |  | [optional] 
**Undeletable** | Pointer to **bool** |  | [optional] 
**Updateable** | Pointer to **bool** |  | [optional] 
**NamedLayoutInfos** | Pointer to **[]interface{}** |  | [optional] 
**ActionOverrides** | Pointer to [**[]ActionOverrides**](ActionOverrides.md) |  | [optional] 
**ChildRelationships** | Pointer to [**[]ChildRelationships**](ChildRelationships.md) |  | [optional] 
**RecordTypeInfos** | Pointer to [**[]RecordTypeInfo**](RecordTypeInfo.md) |  | [optional] 
**SupportedScopes** | Pointer to [**[]SupportedScopes**](SupportedScopes.md) |  | [optional] 
**Urls** | Pointer to **map[string]interface{}** |  | [optional] 
**Fields** | Pointer to [**[]FieldDescribe**](FieldDescribe.md) |  | [optional] 

## Methods

### NewSObjectDescribe

`func NewSObjectDescribe() *SObjectDescribe`

NewSObjectDescribe instantiates a new SObjectDescribe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSObjectDescribeWithDefaults

`func NewSObjectDescribeWithDefaults() *SObjectDescribe`

NewSObjectDescribeWithDefaults instantiates a new SObjectDescribe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivateable

`func (o *SObjectDescribe) GetActivateable() bool`

GetActivateable returns the Activateable field if non-nil, zero value otherwise.

### GetActivateableOk

`func (o *SObjectDescribe) GetActivateableOk() (*bool, bool)`

GetActivateableOk returns a tuple with the Activateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateable

`func (o *SObjectDescribe) SetActivateable(v bool)`

SetActivateable sets Activateable field to given value.

### HasActivateable

`func (o *SObjectDescribe) HasActivateable() bool`

HasActivateable returns a boolean if a field has been set.

### GetAssociateEntityType

`func (o *SObjectDescribe) GetAssociateEntityType() string`

GetAssociateEntityType returns the AssociateEntityType field if non-nil, zero value otherwise.

### GetAssociateEntityTypeOk

`func (o *SObjectDescribe) GetAssociateEntityTypeOk() (*string, bool)`

GetAssociateEntityTypeOk returns a tuple with the AssociateEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociateEntityType

`func (o *SObjectDescribe) SetAssociateEntityType(v string)`

SetAssociateEntityType sets AssociateEntityType field to given value.

### HasAssociateEntityType

`func (o *SObjectDescribe) HasAssociateEntityType() bool`

HasAssociateEntityType returns a boolean if a field has been set.

### SetAssociateEntityTypeNil

`func (o *SObjectDescribe) SetAssociateEntityTypeNil(b bool)`

 SetAssociateEntityTypeNil sets the value for AssociateEntityType to be an explicit nil

### UnsetAssociateEntityType
`func (o *SObjectDescribe) UnsetAssociateEntityType()`

UnsetAssociateEntityType ensures that no value is present for AssociateEntityType, not even an explicit nil
### GetAssociateParentEntity

`func (o *SObjectDescribe) GetAssociateParentEntity() string`

GetAssociateParentEntity returns the AssociateParentEntity field if non-nil, zero value otherwise.

### GetAssociateParentEntityOk

`func (o *SObjectDescribe) GetAssociateParentEntityOk() (*string, bool)`

GetAssociateParentEntityOk returns a tuple with the AssociateParentEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociateParentEntity

`func (o *SObjectDescribe) SetAssociateParentEntity(v string)`

SetAssociateParentEntity sets AssociateParentEntity field to given value.

### HasAssociateParentEntity

`func (o *SObjectDescribe) HasAssociateParentEntity() bool`

HasAssociateParentEntity returns a boolean if a field has been set.

### SetAssociateParentEntityNil

`func (o *SObjectDescribe) SetAssociateParentEntityNil(b bool)`

 SetAssociateParentEntityNil sets the value for AssociateParentEntity to be an explicit nil

### UnsetAssociateParentEntity
`func (o *SObjectDescribe) UnsetAssociateParentEntity()`

UnsetAssociateParentEntity ensures that no value is present for AssociateParentEntity, not even an explicit nil
### GetCompactLayoutable

`func (o *SObjectDescribe) GetCompactLayoutable() bool`

GetCompactLayoutable returns the CompactLayoutable field if non-nil, zero value otherwise.

### GetCompactLayoutableOk

`func (o *SObjectDescribe) GetCompactLayoutableOk() (*bool, bool)`

GetCompactLayoutableOk returns a tuple with the CompactLayoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactLayoutable

`func (o *SObjectDescribe) SetCompactLayoutable(v bool)`

SetCompactLayoutable sets CompactLayoutable field to given value.

### HasCompactLayoutable

`func (o *SObjectDescribe) HasCompactLayoutable() bool`

HasCompactLayoutable returns a boolean if a field has been set.

### GetCreateable

`func (o *SObjectDescribe) GetCreateable() bool`

GetCreateable returns the Createable field if non-nil, zero value otherwise.

### GetCreateableOk

`func (o *SObjectDescribe) GetCreateableOk() (*bool, bool)`

GetCreateableOk returns a tuple with the Createable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateable

`func (o *SObjectDescribe) SetCreateable(v bool)`

SetCreateable sets Createable field to given value.

### HasCreateable

`func (o *SObjectDescribe) HasCreateable() bool`

HasCreateable returns a boolean if a field has been set.

### GetCustom

`func (o *SObjectDescribe) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *SObjectDescribe) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *SObjectDescribe) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *SObjectDescribe) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomSetting

`func (o *SObjectDescribe) GetCustomSetting() bool`

GetCustomSetting returns the CustomSetting field if non-nil, zero value otherwise.

### GetCustomSettingOk

`func (o *SObjectDescribe) GetCustomSettingOk() (*bool, bool)`

GetCustomSettingOk returns a tuple with the CustomSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSetting

`func (o *SObjectDescribe) SetCustomSetting(v bool)`

SetCustomSetting sets CustomSetting field to given value.

### HasCustomSetting

`func (o *SObjectDescribe) HasCustomSetting() bool`

HasCustomSetting returns a boolean if a field has been set.

### GetDeepCloneable

`func (o *SObjectDescribe) GetDeepCloneable() bool`

GetDeepCloneable returns the DeepCloneable field if non-nil, zero value otherwise.

### GetDeepCloneableOk

`func (o *SObjectDescribe) GetDeepCloneableOk() (*bool, bool)`

GetDeepCloneableOk returns a tuple with the DeepCloneable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepCloneable

`func (o *SObjectDescribe) SetDeepCloneable(v bool)`

SetDeepCloneable sets DeepCloneable field to given value.

### HasDeepCloneable

`func (o *SObjectDescribe) HasDeepCloneable() bool`

HasDeepCloneable returns a boolean if a field has been set.

### GetDefaultImplementation

`func (o *SObjectDescribe) GetDefaultImplementation() string`

GetDefaultImplementation returns the DefaultImplementation field if non-nil, zero value otherwise.

### GetDefaultImplementationOk

`func (o *SObjectDescribe) GetDefaultImplementationOk() (*string, bool)`

GetDefaultImplementationOk returns a tuple with the DefaultImplementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImplementation

`func (o *SObjectDescribe) SetDefaultImplementation(v string)`

SetDefaultImplementation sets DefaultImplementation field to given value.

### HasDefaultImplementation

`func (o *SObjectDescribe) HasDefaultImplementation() bool`

HasDefaultImplementation returns a boolean if a field has been set.

### SetDefaultImplementationNil

`func (o *SObjectDescribe) SetDefaultImplementationNil(b bool)`

 SetDefaultImplementationNil sets the value for DefaultImplementation to be an explicit nil

### UnsetDefaultImplementation
`func (o *SObjectDescribe) UnsetDefaultImplementation()`

UnsetDefaultImplementation ensures that no value is present for DefaultImplementation, not even an explicit nil
### GetDeletable

`func (o *SObjectDescribe) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *SObjectDescribe) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *SObjectDescribe) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *SObjectDescribe) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDeprecatedAndHidden

`func (o *SObjectDescribe) GetDeprecatedAndHidden() bool`

GetDeprecatedAndHidden returns the DeprecatedAndHidden field if non-nil, zero value otherwise.

### GetDeprecatedAndHiddenOk

`func (o *SObjectDescribe) GetDeprecatedAndHiddenOk() (*bool, bool)`

GetDeprecatedAndHiddenOk returns a tuple with the DeprecatedAndHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedAndHidden

`func (o *SObjectDescribe) SetDeprecatedAndHidden(v bool)`

SetDeprecatedAndHidden sets DeprecatedAndHidden field to given value.

### HasDeprecatedAndHidden

`func (o *SObjectDescribe) HasDeprecatedAndHidden() bool`

HasDeprecatedAndHidden returns a boolean if a field has been set.

### GetExtendedBy

`func (o *SObjectDescribe) GetExtendedBy() string`

GetExtendedBy returns the ExtendedBy field if non-nil, zero value otherwise.

### GetExtendedByOk

`func (o *SObjectDescribe) GetExtendedByOk() (*string, bool)`

GetExtendedByOk returns a tuple with the ExtendedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedBy

`func (o *SObjectDescribe) SetExtendedBy(v string)`

SetExtendedBy sets ExtendedBy field to given value.

### HasExtendedBy

`func (o *SObjectDescribe) HasExtendedBy() bool`

HasExtendedBy returns a boolean if a field has been set.

### SetExtendedByNil

`func (o *SObjectDescribe) SetExtendedByNil(b bool)`

 SetExtendedByNil sets the value for ExtendedBy to be an explicit nil

### UnsetExtendedBy
`func (o *SObjectDescribe) UnsetExtendedBy()`

UnsetExtendedBy ensures that no value is present for ExtendedBy, not even an explicit nil
### GetExtendsInterfaces

`func (o *SObjectDescribe) GetExtendsInterfaces() string`

GetExtendsInterfaces returns the ExtendsInterfaces field if non-nil, zero value otherwise.

### GetExtendsInterfacesOk

`func (o *SObjectDescribe) GetExtendsInterfacesOk() (*string, bool)`

GetExtendsInterfacesOk returns a tuple with the ExtendsInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendsInterfaces

`func (o *SObjectDescribe) SetExtendsInterfaces(v string)`

SetExtendsInterfaces sets ExtendsInterfaces field to given value.

### HasExtendsInterfaces

`func (o *SObjectDescribe) HasExtendsInterfaces() bool`

HasExtendsInterfaces returns a boolean if a field has been set.

### SetExtendsInterfacesNil

`func (o *SObjectDescribe) SetExtendsInterfacesNil(b bool)`

 SetExtendsInterfacesNil sets the value for ExtendsInterfaces to be an explicit nil

### UnsetExtendsInterfaces
`func (o *SObjectDescribe) UnsetExtendsInterfaces()`

UnsetExtendsInterfaces ensures that no value is present for ExtendsInterfaces, not even an explicit nil
### GetFeedEnabled

`func (o *SObjectDescribe) GetFeedEnabled() bool`

GetFeedEnabled returns the FeedEnabled field if non-nil, zero value otherwise.

### GetFeedEnabledOk

`func (o *SObjectDescribe) GetFeedEnabledOk() (*bool, bool)`

GetFeedEnabledOk returns a tuple with the FeedEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedEnabled

`func (o *SObjectDescribe) SetFeedEnabled(v bool)`

SetFeedEnabled sets FeedEnabled field to given value.

### HasFeedEnabled

`func (o *SObjectDescribe) HasFeedEnabled() bool`

HasFeedEnabled returns a boolean if a field has been set.

### GetHasSubtypes

`func (o *SObjectDescribe) GetHasSubtypes() bool`

GetHasSubtypes returns the HasSubtypes field if non-nil, zero value otherwise.

### GetHasSubtypesOk

`func (o *SObjectDescribe) GetHasSubtypesOk() (*bool, bool)`

GetHasSubtypesOk returns a tuple with the HasSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSubtypes

`func (o *SObjectDescribe) SetHasSubtypes(v bool)`

SetHasSubtypes sets HasSubtypes field to given value.

### HasHasSubtypes

`func (o *SObjectDescribe) HasHasSubtypes() bool`

HasHasSubtypes returns a boolean if a field has been set.

### GetImplementedBy

`func (o *SObjectDescribe) GetImplementedBy() string`

GetImplementedBy returns the ImplementedBy field if non-nil, zero value otherwise.

### GetImplementedByOk

`func (o *SObjectDescribe) GetImplementedByOk() (*string, bool)`

GetImplementedByOk returns a tuple with the ImplementedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementedBy

`func (o *SObjectDescribe) SetImplementedBy(v string)`

SetImplementedBy sets ImplementedBy field to given value.

### HasImplementedBy

`func (o *SObjectDescribe) HasImplementedBy() bool`

HasImplementedBy returns a boolean if a field has been set.

### SetImplementedByNil

`func (o *SObjectDescribe) SetImplementedByNil(b bool)`

 SetImplementedByNil sets the value for ImplementedBy to be an explicit nil

### UnsetImplementedBy
`func (o *SObjectDescribe) UnsetImplementedBy()`

UnsetImplementedBy ensures that no value is present for ImplementedBy, not even an explicit nil
### GetImplementsInterfaces

`func (o *SObjectDescribe) GetImplementsInterfaces() string`

GetImplementsInterfaces returns the ImplementsInterfaces field if non-nil, zero value otherwise.

### GetImplementsInterfacesOk

`func (o *SObjectDescribe) GetImplementsInterfacesOk() (*string, bool)`

GetImplementsInterfacesOk returns a tuple with the ImplementsInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementsInterfaces

`func (o *SObjectDescribe) SetImplementsInterfaces(v string)`

SetImplementsInterfaces sets ImplementsInterfaces field to given value.

### HasImplementsInterfaces

`func (o *SObjectDescribe) HasImplementsInterfaces() bool`

HasImplementsInterfaces returns a boolean if a field has been set.

### SetImplementsInterfacesNil

`func (o *SObjectDescribe) SetImplementsInterfacesNil(b bool)`

 SetImplementsInterfacesNil sets the value for ImplementsInterfaces to be an explicit nil

### UnsetImplementsInterfaces
`func (o *SObjectDescribe) UnsetImplementsInterfaces()`

UnsetImplementsInterfaces ensures that no value is present for ImplementsInterfaces, not even an explicit nil
### GetIsInterface

`func (o *SObjectDescribe) GetIsInterface() bool`

GetIsInterface returns the IsInterface field if non-nil, zero value otherwise.

### GetIsInterfaceOk

`func (o *SObjectDescribe) GetIsInterfaceOk() (*bool, bool)`

GetIsInterfaceOk returns a tuple with the IsInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInterface

`func (o *SObjectDescribe) SetIsInterface(v bool)`

SetIsInterface sets IsInterface field to given value.

### HasIsInterface

`func (o *SObjectDescribe) HasIsInterface() bool`

HasIsInterface returns a boolean if a field has been set.

### GetIsSubtype

`func (o *SObjectDescribe) GetIsSubtype() bool`

GetIsSubtype returns the IsSubtype field if non-nil, zero value otherwise.

### GetIsSubtypeOk

`func (o *SObjectDescribe) GetIsSubtypeOk() (*bool, bool)`

GetIsSubtypeOk returns a tuple with the IsSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubtype

`func (o *SObjectDescribe) SetIsSubtype(v bool)`

SetIsSubtype sets IsSubtype field to given value.

### HasIsSubtype

`func (o *SObjectDescribe) HasIsSubtype() bool`

HasIsSubtype returns a boolean if a field has been set.

### GetKeyPrefix

`func (o *SObjectDescribe) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *SObjectDescribe) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *SObjectDescribe) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.

### HasKeyPrefix

`func (o *SObjectDescribe) HasKeyPrefix() bool`

HasKeyPrefix returns a boolean if a field has been set.

### GetLabel

`func (o *SObjectDescribe) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SObjectDescribe) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SObjectDescribe) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SObjectDescribe) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLabelPlural

`func (o *SObjectDescribe) GetLabelPlural() string`

GetLabelPlural returns the LabelPlural field if non-nil, zero value otherwise.

### GetLabelPluralOk

`func (o *SObjectDescribe) GetLabelPluralOk() (*string, bool)`

GetLabelPluralOk returns a tuple with the LabelPlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPlural

`func (o *SObjectDescribe) SetLabelPlural(v string)`

SetLabelPlural sets LabelPlural field to given value.

### HasLabelPlural

`func (o *SObjectDescribe) HasLabelPlural() bool`

HasLabelPlural returns a boolean if a field has been set.

### GetLayoutable

`func (o *SObjectDescribe) GetLayoutable() bool`

GetLayoutable returns the Layoutable field if non-nil, zero value otherwise.

### GetLayoutableOk

`func (o *SObjectDescribe) GetLayoutableOk() (*bool, bool)`

GetLayoutableOk returns a tuple with the Layoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutable

`func (o *SObjectDescribe) SetLayoutable(v bool)`

SetLayoutable sets Layoutable field to given value.

### HasLayoutable

`func (o *SObjectDescribe) HasLayoutable() bool`

HasLayoutable returns a boolean if a field has been set.

### GetListviewable

`func (o *SObjectDescribe) GetListviewable() string`

GetListviewable returns the Listviewable field if non-nil, zero value otherwise.

### GetListviewableOk

`func (o *SObjectDescribe) GetListviewableOk() (*string, bool)`

GetListviewableOk returns a tuple with the Listviewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListviewable

`func (o *SObjectDescribe) SetListviewable(v string)`

SetListviewable sets Listviewable field to given value.

### HasListviewable

`func (o *SObjectDescribe) HasListviewable() bool`

HasListviewable returns a boolean if a field has been set.

### SetListviewableNil

`func (o *SObjectDescribe) SetListviewableNil(b bool)`

 SetListviewableNil sets the value for Listviewable to be an explicit nil

### UnsetListviewable
`func (o *SObjectDescribe) UnsetListviewable()`

UnsetListviewable ensures that no value is present for Listviewable, not even an explicit nil
### GetLookupLayoutable

`func (o *SObjectDescribe) GetLookupLayoutable() string`

GetLookupLayoutable returns the LookupLayoutable field if non-nil, zero value otherwise.

### GetLookupLayoutableOk

`func (o *SObjectDescribe) GetLookupLayoutableOk() (*string, bool)`

GetLookupLayoutableOk returns a tuple with the LookupLayoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupLayoutable

`func (o *SObjectDescribe) SetLookupLayoutable(v string)`

SetLookupLayoutable sets LookupLayoutable field to given value.

### HasLookupLayoutable

`func (o *SObjectDescribe) HasLookupLayoutable() bool`

HasLookupLayoutable returns a boolean if a field has been set.

### SetLookupLayoutableNil

`func (o *SObjectDescribe) SetLookupLayoutableNil(b bool)`

 SetLookupLayoutableNil sets the value for LookupLayoutable to be an explicit nil

### UnsetLookupLayoutable
`func (o *SObjectDescribe) UnsetLookupLayoutable()`

UnsetLookupLayoutable ensures that no value is present for LookupLayoutable, not even an explicit nil
### GetMergeable

`func (o *SObjectDescribe) GetMergeable() bool`

GetMergeable returns the Mergeable field if non-nil, zero value otherwise.

### GetMergeableOk

`func (o *SObjectDescribe) GetMergeableOk() (*bool, bool)`

GetMergeableOk returns a tuple with the Mergeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeable

`func (o *SObjectDescribe) SetMergeable(v bool)`

SetMergeable sets Mergeable field to given value.

### HasMergeable

`func (o *SObjectDescribe) HasMergeable() bool`

HasMergeable returns a boolean if a field has been set.

### GetMruEnabled

`func (o *SObjectDescribe) GetMruEnabled() bool`

GetMruEnabled returns the MruEnabled field if non-nil, zero value otherwise.

### GetMruEnabledOk

`func (o *SObjectDescribe) GetMruEnabledOk() (*bool, bool)`

GetMruEnabledOk returns a tuple with the MruEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMruEnabled

`func (o *SObjectDescribe) SetMruEnabled(v bool)`

SetMruEnabled sets MruEnabled field to given value.

### HasMruEnabled

`func (o *SObjectDescribe) HasMruEnabled() bool`

HasMruEnabled returns a boolean if a field has been set.

### GetName

`func (o *SObjectDescribe) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SObjectDescribe) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SObjectDescribe) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SObjectDescribe) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkScopeFieldName

`func (o *SObjectDescribe) GetNetworkScopeFieldName() string`

GetNetworkScopeFieldName returns the NetworkScopeFieldName field if non-nil, zero value otherwise.

### GetNetworkScopeFieldNameOk

`func (o *SObjectDescribe) GetNetworkScopeFieldNameOk() (*string, bool)`

GetNetworkScopeFieldNameOk returns a tuple with the NetworkScopeFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkScopeFieldName

`func (o *SObjectDescribe) SetNetworkScopeFieldName(v string)`

SetNetworkScopeFieldName sets NetworkScopeFieldName field to given value.

### HasNetworkScopeFieldName

`func (o *SObjectDescribe) HasNetworkScopeFieldName() bool`

HasNetworkScopeFieldName returns a boolean if a field has been set.

### SetNetworkScopeFieldNameNil

`func (o *SObjectDescribe) SetNetworkScopeFieldNameNil(b bool)`

 SetNetworkScopeFieldNameNil sets the value for NetworkScopeFieldName to be an explicit nil

### UnsetNetworkScopeFieldName
`func (o *SObjectDescribe) UnsetNetworkScopeFieldName()`

UnsetNetworkScopeFieldName ensures that no value is present for NetworkScopeFieldName, not even an explicit nil
### GetQueryable

`func (o *SObjectDescribe) GetQueryable() bool`

GetQueryable returns the Queryable field if non-nil, zero value otherwise.

### GetQueryableOk

`func (o *SObjectDescribe) GetQueryableOk() (*bool, bool)`

GetQueryableOk returns a tuple with the Queryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryable

`func (o *SObjectDescribe) SetQueryable(v bool)`

SetQueryable sets Queryable field to given value.

### HasQueryable

`func (o *SObjectDescribe) HasQueryable() bool`

HasQueryable returns a boolean if a field has been set.

### GetReplicateable

`func (o *SObjectDescribe) GetReplicateable() bool`

GetReplicateable returns the Replicateable field if non-nil, zero value otherwise.

### GetReplicateableOk

`func (o *SObjectDescribe) GetReplicateableOk() (*bool, bool)`

GetReplicateableOk returns a tuple with the Replicateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateable

`func (o *SObjectDescribe) SetReplicateable(v bool)`

SetReplicateable sets Replicateable field to given value.

### HasReplicateable

`func (o *SObjectDescribe) HasReplicateable() bool`

HasReplicateable returns a boolean if a field has been set.

### GetRetrieveable

`func (o *SObjectDescribe) GetRetrieveable() bool`

GetRetrieveable returns the Retrieveable field if non-nil, zero value otherwise.

### GetRetrieveableOk

`func (o *SObjectDescribe) GetRetrieveableOk() (*bool, bool)`

GetRetrieveableOk returns a tuple with the Retrieveable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveable

`func (o *SObjectDescribe) SetRetrieveable(v bool)`

SetRetrieveable sets Retrieveable field to given value.

### HasRetrieveable

`func (o *SObjectDescribe) HasRetrieveable() bool`

HasRetrieveable returns a boolean if a field has been set.

### GetSearchLayoutable

`func (o *SObjectDescribe) GetSearchLayoutable() bool`

GetSearchLayoutable returns the SearchLayoutable field if non-nil, zero value otherwise.

### GetSearchLayoutableOk

`func (o *SObjectDescribe) GetSearchLayoutableOk() (*bool, bool)`

GetSearchLayoutableOk returns a tuple with the SearchLayoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchLayoutable

`func (o *SObjectDescribe) SetSearchLayoutable(v bool)`

SetSearchLayoutable sets SearchLayoutable field to given value.

### HasSearchLayoutable

`func (o *SObjectDescribe) HasSearchLayoutable() bool`

HasSearchLayoutable returns a boolean if a field has been set.

### GetSearchable

`func (o *SObjectDescribe) GetSearchable() bool`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *SObjectDescribe) GetSearchableOk() (*bool, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *SObjectDescribe) SetSearchable(v bool)`

SetSearchable sets Searchable field to given value.

### HasSearchable

`func (o *SObjectDescribe) HasSearchable() bool`

HasSearchable returns a boolean if a field has been set.

### GetSobjectDescribeOption

`func (o *SObjectDescribe) GetSobjectDescribeOption() string`

GetSobjectDescribeOption returns the SobjectDescribeOption field if non-nil, zero value otherwise.

### GetSobjectDescribeOptionOk

`func (o *SObjectDescribe) GetSobjectDescribeOptionOk() (*string, bool)`

GetSobjectDescribeOptionOk returns a tuple with the SobjectDescribeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSobjectDescribeOption

`func (o *SObjectDescribe) SetSobjectDescribeOption(v string)`

SetSobjectDescribeOption sets SobjectDescribeOption field to given value.

### HasSobjectDescribeOption

`func (o *SObjectDescribe) HasSobjectDescribeOption() bool`

HasSobjectDescribeOption returns a boolean if a field has been set.

### GetTriggerable

`func (o *SObjectDescribe) GetTriggerable() bool`

GetTriggerable returns the Triggerable field if non-nil, zero value otherwise.

### GetTriggerableOk

`func (o *SObjectDescribe) GetTriggerableOk() (*bool, bool)`

GetTriggerableOk returns a tuple with the Triggerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerable

`func (o *SObjectDescribe) SetTriggerable(v bool)`

SetTriggerable sets Triggerable field to given value.

### HasTriggerable

`func (o *SObjectDescribe) HasTriggerable() bool`

HasTriggerable returns a boolean if a field has been set.

### GetUndeletable

`func (o *SObjectDescribe) GetUndeletable() bool`

GetUndeletable returns the Undeletable field if non-nil, zero value otherwise.

### GetUndeletableOk

`func (o *SObjectDescribe) GetUndeletableOk() (*bool, bool)`

GetUndeletableOk returns a tuple with the Undeletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndeletable

`func (o *SObjectDescribe) SetUndeletable(v bool)`

SetUndeletable sets Undeletable field to given value.

### HasUndeletable

`func (o *SObjectDescribe) HasUndeletable() bool`

HasUndeletable returns a boolean if a field has been set.

### GetUpdateable

`func (o *SObjectDescribe) GetUpdateable() bool`

GetUpdateable returns the Updateable field if non-nil, zero value otherwise.

### GetUpdateableOk

`func (o *SObjectDescribe) GetUpdateableOk() (*bool, bool)`

GetUpdateableOk returns a tuple with the Updateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateable

`func (o *SObjectDescribe) SetUpdateable(v bool)`

SetUpdateable sets Updateable field to given value.

### HasUpdateable

`func (o *SObjectDescribe) HasUpdateable() bool`

HasUpdateable returns a boolean if a field has been set.

### GetNamedLayoutInfos

`func (o *SObjectDescribe) GetNamedLayoutInfos() []interface{}`

GetNamedLayoutInfos returns the NamedLayoutInfos field if non-nil, zero value otherwise.

### GetNamedLayoutInfosOk

`func (o *SObjectDescribe) GetNamedLayoutInfosOk() (*[]interface{}, bool)`

GetNamedLayoutInfosOk returns a tuple with the NamedLayoutInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedLayoutInfos

`func (o *SObjectDescribe) SetNamedLayoutInfos(v []interface{})`

SetNamedLayoutInfos sets NamedLayoutInfos field to given value.

### HasNamedLayoutInfos

`func (o *SObjectDescribe) HasNamedLayoutInfos() bool`

HasNamedLayoutInfos returns a boolean if a field has been set.

### GetActionOverrides

`func (o *SObjectDescribe) GetActionOverrides() []ActionOverrides`

GetActionOverrides returns the ActionOverrides field if non-nil, zero value otherwise.

### GetActionOverridesOk

`func (o *SObjectDescribe) GetActionOverridesOk() (*[]ActionOverrides, bool)`

GetActionOverridesOk returns a tuple with the ActionOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOverrides

`func (o *SObjectDescribe) SetActionOverrides(v []ActionOverrides)`

SetActionOverrides sets ActionOverrides field to given value.

### HasActionOverrides

`func (o *SObjectDescribe) HasActionOverrides() bool`

HasActionOverrides returns a boolean if a field has been set.

### GetChildRelationships

`func (o *SObjectDescribe) GetChildRelationships() []ChildRelationships`

GetChildRelationships returns the ChildRelationships field if non-nil, zero value otherwise.

### GetChildRelationshipsOk

`func (o *SObjectDescribe) GetChildRelationshipsOk() (*[]ChildRelationships, bool)`

GetChildRelationshipsOk returns a tuple with the ChildRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildRelationships

`func (o *SObjectDescribe) SetChildRelationships(v []ChildRelationships)`

SetChildRelationships sets ChildRelationships field to given value.

### HasChildRelationships

`func (o *SObjectDescribe) HasChildRelationships() bool`

HasChildRelationships returns a boolean if a field has been set.

### GetRecordTypeInfos

`func (o *SObjectDescribe) GetRecordTypeInfos() []RecordTypeInfo`

GetRecordTypeInfos returns the RecordTypeInfos field if non-nil, zero value otherwise.

### GetRecordTypeInfosOk

`func (o *SObjectDescribe) GetRecordTypeInfosOk() (*[]RecordTypeInfo, bool)`

GetRecordTypeInfosOk returns a tuple with the RecordTypeInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTypeInfos

`func (o *SObjectDescribe) SetRecordTypeInfos(v []RecordTypeInfo)`

SetRecordTypeInfos sets RecordTypeInfos field to given value.

### HasRecordTypeInfos

`func (o *SObjectDescribe) HasRecordTypeInfos() bool`

HasRecordTypeInfos returns a boolean if a field has been set.

### GetSupportedScopes

`func (o *SObjectDescribe) GetSupportedScopes() []SupportedScopes`

GetSupportedScopes returns the SupportedScopes field if non-nil, zero value otherwise.

### GetSupportedScopesOk

`func (o *SObjectDescribe) GetSupportedScopesOk() (*[]SupportedScopes, bool)`

GetSupportedScopesOk returns a tuple with the SupportedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedScopes

`func (o *SObjectDescribe) SetSupportedScopes(v []SupportedScopes)`

SetSupportedScopes sets SupportedScopes field to given value.

### HasSupportedScopes

`func (o *SObjectDescribe) HasSupportedScopes() bool`

HasSupportedScopes returns a boolean if a field has been set.

### GetUrls

`func (o *SObjectDescribe) GetUrls() map[string]interface{}`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *SObjectDescribe) GetUrlsOk() (*map[string]interface{}, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *SObjectDescribe) SetUrls(v map[string]interface{})`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *SObjectDescribe) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetFields

`func (o *SObjectDescribe) GetFields() []FieldDescribe`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SObjectDescribe) GetFieldsOk() (*[]FieldDescribe, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SObjectDescribe) SetFields(v []FieldDescribe)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SObjectDescribe) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


