# SObjectDescribesSobjectsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activateable** | Pointer to **bool** |  | [optional] 
**AssociateEntityType** | Pointer to **string** |  | [optional] 
**AssociateParentEntity** | Pointer to **string** |  | [optional] 
**CompactLayoutable** | Pointer to **bool** |  | [optional] 
**Createable** | Pointer to **bool** |  | [optional] 
**Custom** | Pointer to **bool** |  | [optional] 
**CustomSetting** | Pointer to **bool** |  | [optional] 
**DeepCloneable** | Pointer to **bool** |  | [optional] 
**DefaultImplementation** | Pointer to **string** |  | [optional] 
**Deletable** | Pointer to **bool** |  | [optional] 
**DeprecatedAndHidden** | Pointer to **bool** |  | [optional] 
**ExtendedBy** | Pointer to **string** |  | [optional] 
**ExtendsInterfaces** | Pointer to **string** |  | [optional] 
**FeedEnabled** | Pointer to **bool** |  | [optional] 
**HasSubtypes** | Pointer to **bool** |  | [optional] 
**ImplementedBy** | Pointer to **string** |  | [optional] 
**ImplementsInterfaces** | Pointer to **string** |  | [optional] 
**IsInterface** | Pointer to **bool** |  | [optional] 
**IsSubtype** | Pointer to **bool** |  | [optional] 
**KeyPrefix** | **string** |  | 
**Label** | **string** |  | 
**LabelPlural** | **string** |  | 
**Layoutable** | Pointer to **bool** |  | [optional] 
**Listviewable** | Pointer to **string** |  | [optional] 
**LookupLayoutable** | Pointer to **string** |  | [optional] 
**Mergeable** | Pointer to **bool** |  | [optional] 
**MruEnabled** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**NetworkScopeFieldName** | Pointer to **string** |  | [optional] 
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
**ActionOverrides** | Pointer to [**[]SObjectDescribeActionOverrides**](SObjectDescribeActionOverrides.md) |  | [optional] 
**ChildRelationships** | Pointer to [**[]SObjectDescribeChildRelationships**](SObjectDescribeChildRelationships.md) |  | [optional] 
**RecordTypeInfos** | Pointer to [**[]SObjectRecordTypeInfo**](SObjectRecordTypeInfo.md) |  | [optional] 
**SupportedScopes** | Pointer to [**[]SObjectDescribeSupportedScopes**](SObjectDescribeSupportedScopes.md) |  | [optional] 
**Urls** | Pointer to **map[string]interface{}** |  | [optional] 
**Fields** | Pointer to [**[]SObjectFieldDescribe**](SObjectFieldDescribe.md) |  | [optional] 

## Methods

### NewSObjectDescribesSobjectsInner

`func NewSObjectDescribesSobjectsInner(keyPrefix string, label string, labelPlural string, name string, ) *SObjectDescribesSobjectsInner`

NewSObjectDescribesSobjectsInner instantiates a new SObjectDescribesSobjectsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSObjectDescribesSobjectsInnerWithDefaults

`func NewSObjectDescribesSobjectsInnerWithDefaults() *SObjectDescribesSobjectsInner`

NewSObjectDescribesSobjectsInnerWithDefaults instantiates a new SObjectDescribesSobjectsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivateable

`func (o *SObjectDescribesSobjectsInner) GetActivateable() bool`

GetActivateable returns the Activateable field if non-nil, zero value otherwise.

### GetActivateableOk

`func (o *SObjectDescribesSobjectsInner) GetActivateableOk() (*bool, bool)`

GetActivateableOk returns a tuple with the Activateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateable

`func (o *SObjectDescribesSobjectsInner) SetActivateable(v bool)`

SetActivateable sets Activateable field to given value.

### HasActivateable

`func (o *SObjectDescribesSobjectsInner) HasActivateable() bool`

HasActivateable returns a boolean if a field has been set.

### GetAssociateEntityType

`func (o *SObjectDescribesSobjectsInner) GetAssociateEntityType() string`

GetAssociateEntityType returns the AssociateEntityType field if non-nil, zero value otherwise.

### GetAssociateEntityTypeOk

`func (o *SObjectDescribesSobjectsInner) GetAssociateEntityTypeOk() (*string, bool)`

GetAssociateEntityTypeOk returns a tuple with the AssociateEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociateEntityType

`func (o *SObjectDescribesSobjectsInner) SetAssociateEntityType(v string)`

SetAssociateEntityType sets AssociateEntityType field to given value.

### HasAssociateEntityType

`func (o *SObjectDescribesSobjectsInner) HasAssociateEntityType() bool`

HasAssociateEntityType returns a boolean if a field has been set.

### GetAssociateParentEntity

`func (o *SObjectDescribesSobjectsInner) GetAssociateParentEntity() string`

GetAssociateParentEntity returns the AssociateParentEntity field if non-nil, zero value otherwise.

### GetAssociateParentEntityOk

`func (o *SObjectDescribesSobjectsInner) GetAssociateParentEntityOk() (*string, bool)`

GetAssociateParentEntityOk returns a tuple with the AssociateParentEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociateParentEntity

`func (o *SObjectDescribesSobjectsInner) SetAssociateParentEntity(v string)`

SetAssociateParentEntity sets AssociateParentEntity field to given value.

### HasAssociateParentEntity

`func (o *SObjectDescribesSobjectsInner) HasAssociateParentEntity() bool`

HasAssociateParentEntity returns a boolean if a field has been set.

### GetCompactLayoutable

`func (o *SObjectDescribesSobjectsInner) GetCompactLayoutable() bool`

GetCompactLayoutable returns the CompactLayoutable field if non-nil, zero value otherwise.

### GetCompactLayoutableOk

`func (o *SObjectDescribesSobjectsInner) GetCompactLayoutableOk() (*bool, bool)`

GetCompactLayoutableOk returns a tuple with the CompactLayoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactLayoutable

`func (o *SObjectDescribesSobjectsInner) SetCompactLayoutable(v bool)`

SetCompactLayoutable sets CompactLayoutable field to given value.

### HasCompactLayoutable

`func (o *SObjectDescribesSobjectsInner) HasCompactLayoutable() bool`

HasCompactLayoutable returns a boolean if a field has been set.

### GetCreateable

`func (o *SObjectDescribesSobjectsInner) GetCreateable() bool`

GetCreateable returns the Createable field if non-nil, zero value otherwise.

### GetCreateableOk

`func (o *SObjectDescribesSobjectsInner) GetCreateableOk() (*bool, bool)`

GetCreateableOk returns a tuple with the Createable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateable

`func (o *SObjectDescribesSobjectsInner) SetCreateable(v bool)`

SetCreateable sets Createable field to given value.

### HasCreateable

`func (o *SObjectDescribesSobjectsInner) HasCreateable() bool`

HasCreateable returns a boolean if a field has been set.

### GetCustom

`func (o *SObjectDescribesSobjectsInner) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *SObjectDescribesSobjectsInner) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *SObjectDescribesSobjectsInner) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *SObjectDescribesSobjectsInner) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomSetting

`func (o *SObjectDescribesSobjectsInner) GetCustomSetting() bool`

GetCustomSetting returns the CustomSetting field if non-nil, zero value otherwise.

### GetCustomSettingOk

`func (o *SObjectDescribesSobjectsInner) GetCustomSettingOk() (*bool, bool)`

GetCustomSettingOk returns a tuple with the CustomSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSetting

`func (o *SObjectDescribesSobjectsInner) SetCustomSetting(v bool)`

SetCustomSetting sets CustomSetting field to given value.

### HasCustomSetting

`func (o *SObjectDescribesSobjectsInner) HasCustomSetting() bool`

HasCustomSetting returns a boolean if a field has been set.

### GetDeepCloneable

`func (o *SObjectDescribesSobjectsInner) GetDeepCloneable() bool`

GetDeepCloneable returns the DeepCloneable field if non-nil, zero value otherwise.

### GetDeepCloneableOk

`func (o *SObjectDescribesSobjectsInner) GetDeepCloneableOk() (*bool, bool)`

GetDeepCloneableOk returns a tuple with the DeepCloneable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepCloneable

`func (o *SObjectDescribesSobjectsInner) SetDeepCloneable(v bool)`

SetDeepCloneable sets DeepCloneable field to given value.

### HasDeepCloneable

`func (o *SObjectDescribesSobjectsInner) HasDeepCloneable() bool`

HasDeepCloneable returns a boolean if a field has been set.

### GetDefaultImplementation

`func (o *SObjectDescribesSobjectsInner) GetDefaultImplementation() string`

GetDefaultImplementation returns the DefaultImplementation field if non-nil, zero value otherwise.

### GetDefaultImplementationOk

`func (o *SObjectDescribesSobjectsInner) GetDefaultImplementationOk() (*string, bool)`

GetDefaultImplementationOk returns a tuple with the DefaultImplementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImplementation

`func (o *SObjectDescribesSobjectsInner) SetDefaultImplementation(v string)`

SetDefaultImplementation sets DefaultImplementation field to given value.

### HasDefaultImplementation

`func (o *SObjectDescribesSobjectsInner) HasDefaultImplementation() bool`

HasDefaultImplementation returns a boolean if a field has been set.

### GetDeletable

`func (o *SObjectDescribesSobjectsInner) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *SObjectDescribesSobjectsInner) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *SObjectDescribesSobjectsInner) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *SObjectDescribesSobjectsInner) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDeprecatedAndHidden

`func (o *SObjectDescribesSobjectsInner) GetDeprecatedAndHidden() bool`

GetDeprecatedAndHidden returns the DeprecatedAndHidden field if non-nil, zero value otherwise.

### GetDeprecatedAndHiddenOk

`func (o *SObjectDescribesSobjectsInner) GetDeprecatedAndHiddenOk() (*bool, bool)`

GetDeprecatedAndHiddenOk returns a tuple with the DeprecatedAndHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedAndHidden

`func (o *SObjectDescribesSobjectsInner) SetDeprecatedAndHidden(v bool)`

SetDeprecatedAndHidden sets DeprecatedAndHidden field to given value.

### HasDeprecatedAndHidden

`func (o *SObjectDescribesSobjectsInner) HasDeprecatedAndHidden() bool`

HasDeprecatedAndHidden returns a boolean if a field has been set.

### GetExtendedBy

`func (o *SObjectDescribesSobjectsInner) GetExtendedBy() string`

GetExtendedBy returns the ExtendedBy field if non-nil, zero value otherwise.

### GetExtendedByOk

`func (o *SObjectDescribesSobjectsInner) GetExtendedByOk() (*string, bool)`

GetExtendedByOk returns a tuple with the ExtendedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedBy

`func (o *SObjectDescribesSobjectsInner) SetExtendedBy(v string)`

SetExtendedBy sets ExtendedBy field to given value.

### HasExtendedBy

`func (o *SObjectDescribesSobjectsInner) HasExtendedBy() bool`

HasExtendedBy returns a boolean if a field has been set.

### GetExtendsInterfaces

`func (o *SObjectDescribesSobjectsInner) GetExtendsInterfaces() string`

GetExtendsInterfaces returns the ExtendsInterfaces field if non-nil, zero value otherwise.

### GetExtendsInterfacesOk

`func (o *SObjectDescribesSobjectsInner) GetExtendsInterfacesOk() (*string, bool)`

GetExtendsInterfacesOk returns a tuple with the ExtendsInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendsInterfaces

`func (o *SObjectDescribesSobjectsInner) SetExtendsInterfaces(v string)`

SetExtendsInterfaces sets ExtendsInterfaces field to given value.

### HasExtendsInterfaces

`func (o *SObjectDescribesSobjectsInner) HasExtendsInterfaces() bool`

HasExtendsInterfaces returns a boolean if a field has been set.

### GetFeedEnabled

`func (o *SObjectDescribesSobjectsInner) GetFeedEnabled() bool`

GetFeedEnabled returns the FeedEnabled field if non-nil, zero value otherwise.

### GetFeedEnabledOk

`func (o *SObjectDescribesSobjectsInner) GetFeedEnabledOk() (*bool, bool)`

GetFeedEnabledOk returns a tuple with the FeedEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedEnabled

`func (o *SObjectDescribesSobjectsInner) SetFeedEnabled(v bool)`

SetFeedEnabled sets FeedEnabled field to given value.

### HasFeedEnabled

`func (o *SObjectDescribesSobjectsInner) HasFeedEnabled() bool`

HasFeedEnabled returns a boolean if a field has been set.

### GetHasSubtypes

`func (o *SObjectDescribesSobjectsInner) GetHasSubtypes() bool`

GetHasSubtypes returns the HasSubtypes field if non-nil, zero value otherwise.

### GetHasSubtypesOk

`func (o *SObjectDescribesSobjectsInner) GetHasSubtypesOk() (*bool, bool)`

GetHasSubtypesOk returns a tuple with the HasSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSubtypes

`func (o *SObjectDescribesSobjectsInner) SetHasSubtypes(v bool)`

SetHasSubtypes sets HasSubtypes field to given value.

### HasHasSubtypes

`func (o *SObjectDescribesSobjectsInner) HasHasSubtypes() bool`

HasHasSubtypes returns a boolean if a field has been set.

### GetImplementedBy

`func (o *SObjectDescribesSobjectsInner) GetImplementedBy() string`

GetImplementedBy returns the ImplementedBy field if non-nil, zero value otherwise.

### GetImplementedByOk

`func (o *SObjectDescribesSobjectsInner) GetImplementedByOk() (*string, bool)`

GetImplementedByOk returns a tuple with the ImplementedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementedBy

`func (o *SObjectDescribesSobjectsInner) SetImplementedBy(v string)`

SetImplementedBy sets ImplementedBy field to given value.

### HasImplementedBy

`func (o *SObjectDescribesSobjectsInner) HasImplementedBy() bool`

HasImplementedBy returns a boolean if a field has been set.

### GetImplementsInterfaces

`func (o *SObjectDescribesSobjectsInner) GetImplementsInterfaces() string`

GetImplementsInterfaces returns the ImplementsInterfaces field if non-nil, zero value otherwise.

### GetImplementsInterfacesOk

`func (o *SObjectDescribesSobjectsInner) GetImplementsInterfacesOk() (*string, bool)`

GetImplementsInterfacesOk returns a tuple with the ImplementsInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementsInterfaces

`func (o *SObjectDescribesSobjectsInner) SetImplementsInterfaces(v string)`

SetImplementsInterfaces sets ImplementsInterfaces field to given value.

### HasImplementsInterfaces

`func (o *SObjectDescribesSobjectsInner) HasImplementsInterfaces() bool`

HasImplementsInterfaces returns a boolean if a field has been set.

### GetIsInterface

`func (o *SObjectDescribesSobjectsInner) GetIsInterface() bool`

GetIsInterface returns the IsInterface field if non-nil, zero value otherwise.

### GetIsInterfaceOk

`func (o *SObjectDescribesSobjectsInner) GetIsInterfaceOk() (*bool, bool)`

GetIsInterfaceOk returns a tuple with the IsInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInterface

`func (o *SObjectDescribesSobjectsInner) SetIsInterface(v bool)`

SetIsInterface sets IsInterface field to given value.

### HasIsInterface

`func (o *SObjectDescribesSobjectsInner) HasIsInterface() bool`

HasIsInterface returns a boolean if a field has been set.

### GetIsSubtype

`func (o *SObjectDescribesSobjectsInner) GetIsSubtype() bool`

GetIsSubtype returns the IsSubtype field if non-nil, zero value otherwise.

### GetIsSubtypeOk

`func (o *SObjectDescribesSobjectsInner) GetIsSubtypeOk() (*bool, bool)`

GetIsSubtypeOk returns a tuple with the IsSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubtype

`func (o *SObjectDescribesSobjectsInner) SetIsSubtype(v bool)`

SetIsSubtype sets IsSubtype field to given value.

### HasIsSubtype

`func (o *SObjectDescribesSobjectsInner) HasIsSubtype() bool`

HasIsSubtype returns a boolean if a field has been set.

### GetKeyPrefix

`func (o *SObjectDescribesSobjectsInner) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *SObjectDescribesSobjectsInner) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *SObjectDescribesSobjectsInner) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.


### GetLabel

`func (o *SObjectDescribesSobjectsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SObjectDescribesSobjectsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SObjectDescribesSobjectsInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLabelPlural

`func (o *SObjectDescribesSobjectsInner) GetLabelPlural() string`

GetLabelPlural returns the LabelPlural field if non-nil, zero value otherwise.

### GetLabelPluralOk

`func (o *SObjectDescribesSobjectsInner) GetLabelPluralOk() (*string, bool)`

GetLabelPluralOk returns a tuple with the LabelPlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPlural

`func (o *SObjectDescribesSobjectsInner) SetLabelPlural(v string)`

SetLabelPlural sets LabelPlural field to given value.


### GetLayoutable

`func (o *SObjectDescribesSobjectsInner) GetLayoutable() bool`

GetLayoutable returns the Layoutable field if non-nil, zero value otherwise.

### GetLayoutableOk

`func (o *SObjectDescribesSobjectsInner) GetLayoutableOk() (*bool, bool)`

GetLayoutableOk returns a tuple with the Layoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutable

`func (o *SObjectDescribesSobjectsInner) SetLayoutable(v bool)`

SetLayoutable sets Layoutable field to given value.

### HasLayoutable

`func (o *SObjectDescribesSobjectsInner) HasLayoutable() bool`

HasLayoutable returns a boolean if a field has been set.

### GetListviewable

`func (o *SObjectDescribesSobjectsInner) GetListviewable() string`

GetListviewable returns the Listviewable field if non-nil, zero value otherwise.

### GetListviewableOk

`func (o *SObjectDescribesSobjectsInner) GetListviewableOk() (*string, bool)`

GetListviewableOk returns a tuple with the Listviewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListviewable

`func (o *SObjectDescribesSobjectsInner) SetListviewable(v string)`

SetListviewable sets Listviewable field to given value.

### HasListviewable

`func (o *SObjectDescribesSobjectsInner) HasListviewable() bool`

HasListviewable returns a boolean if a field has been set.

### GetLookupLayoutable

`func (o *SObjectDescribesSobjectsInner) GetLookupLayoutable() string`

GetLookupLayoutable returns the LookupLayoutable field if non-nil, zero value otherwise.

### GetLookupLayoutableOk

`func (o *SObjectDescribesSobjectsInner) GetLookupLayoutableOk() (*string, bool)`

GetLookupLayoutableOk returns a tuple with the LookupLayoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupLayoutable

`func (o *SObjectDescribesSobjectsInner) SetLookupLayoutable(v string)`

SetLookupLayoutable sets LookupLayoutable field to given value.

### HasLookupLayoutable

`func (o *SObjectDescribesSobjectsInner) HasLookupLayoutable() bool`

HasLookupLayoutable returns a boolean if a field has been set.

### GetMergeable

`func (o *SObjectDescribesSobjectsInner) GetMergeable() bool`

GetMergeable returns the Mergeable field if non-nil, zero value otherwise.

### GetMergeableOk

`func (o *SObjectDescribesSobjectsInner) GetMergeableOk() (*bool, bool)`

GetMergeableOk returns a tuple with the Mergeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeable

`func (o *SObjectDescribesSobjectsInner) SetMergeable(v bool)`

SetMergeable sets Mergeable field to given value.

### HasMergeable

`func (o *SObjectDescribesSobjectsInner) HasMergeable() bool`

HasMergeable returns a boolean if a field has been set.

### GetMruEnabled

`func (o *SObjectDescribesSobjectsInner) GetMruEnabled() bool`

GetMruEnabled returns the MruEnabled field if non-nil, zero value otherwise.

### GetMruEnabledOk

`func (o *SObjectDescribesSobjectsInner) GetMruEnabledOk() (*bool, bool)`

GetMruEnabledOk returns a tuple with the MruEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMruEnabled

`func (o *SObjectDescribesSobjectsInner) SetMruEnabled(v bool)`

SetMruEnabled sets MruEnabled field to given value.

### HasMruEnabled

`func (o *SObjectDescribesSobjectsInner) HasMruEnabled() bool`

HasMruEnabled returns a boolean if a field has been set.

### GetName

`func (o *SObjectDescribesSobjectsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SObjectDescribesSobjectsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SObjectDescribesSobjectsInner) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkScopeFieldName

`func (o *SObjectDescribesSobjectsInner) GetNetworkScopeFieldName() string`

GetNetworkScopeFieldName returns the NetworkScopeFieldName field if non-nil, zero value otherwise.

### GetNetworkScopeFieldNameOk

`func (o *SObjectDescribesSobjectsInner) GetNetworkScopeFieldNameOk() (*string, bool)`

GetNetworkScopeFieldNameOk returns a tuple with the NetworkScopeFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkScopeFieldName

`func (o *SObjectDescribesSobjectsInner) SetNetworkScopeFieldName(v string)`

SetNetworkScopeFieldName sets NetworkScopeFieldName field to given value.

### HasNetworkScopeFieldName

`func (o *SObjectDescribesSobjectsInner) HasNetworkScopeFieldName() bool`

HasNetworkScopeFieldName returns a boolean if a field has been set.

### GetQueryable

`func (o *SObjectDescribesSobjectsInner) GetQueryable() bool`

GetQueryable returns the Queryable field if non-nil, zero value otherwise.

### GetQueryableOk

`func (o *SObjectDescribesSobjectsInner) GetQueryableOk() (*bool, bool)`

GetQueryableOk returns a tuple with the Queryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryable

`func (o *SObjectDescribesSobjectsInner) SetQueryable(v bool)`

SetQueryable sets Queryable field to given value.

### HasQueryable

`func (o *SObjectDescribesSobjectsInner) HasQueryable() bool`

HasQueryable returns a boolean if a field has been set.

### GetReplicateable

`func (o *SObjectDescribesSobjectsInner) GetReplicateable() bool`

GetReplicateable returns the Replicateable field if non-nil, zero value otherwise.

### GetReplicateableOk

`func (o *SObjectDescribesSobjectsInner) GetReplicateableOk() (*bool, bool)`

GetReplicateableOk returns a tuple with the Replicateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateable

`func (o *SObjectDescribesSobjectsInner) SetReplicateable(v bool)`

SetReplicateable sets Replicateable field to given value.

### HasReplicateable

`func (o *SObjectDescribesSobjectsInner) HasReplicateable() bool`

HasReplicateable returns a boolean if a field has been set.

### GetRetrieveable

`func (o *SObjectDescribesSobjectsInner) GetRetrieveable() bool`

GetRetrieveable returns the Retrieveable field if non-nil, zero value otherwise.

### GetRetrieveableOk

`func (o *SObjectDescribesSobjectsInner) GetRetrieveableOk() (*bool, bool)`

GetRetrieveableOk returns a tuple with the Retrieveable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveable

`func (o *SObjectDescribesSobjectsInner) SetRetrieveable(v bool)`

SetRetrieveable sets Retrieveable field to given value.

### HasRetrieveable

`func (o *SObjectDescribesSobjectsInner) HasRetrieveable() bool`

HasRetrieveable returns a boolean if a field has been set.

### GetSearchLayoutable

`func (o *SObjectDescribesSobjectsInner) GetSearchLayoutable() bool`

GetSearchLayoutable returns the SearchLayoutable field if non-nil, zero value otherwise.

### GetSearchLayoutableOk

`func (o *SObjectDescribesSobjectsInner) GetSearchLayoutableOk() (*bool, bool)`

GetSearchLayoutableOk returns a tuple with the SearchLayoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchLayoutable

`func (o *SObjectDescribesSobjectsInner) SetSearchLayoutable(v bool)`

SetSearchLayoutable sets SearchLayoutable field to given value.

### HasSearchLayoutable

`func (o *SObjectDescribesSobjectsInner) HasSearchLayoutable() bool`

HasSearchLayoutable returns a boolean if a field has been set.

### GetSearchable

`func (o *SObjectDescribesSobjectsInner) GetSearchable() bool`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *SObjectDescribesSobjectsInner) GetSearchableOk() (*bool, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *SObjectDescribesSobjectsInner) SetSearchable(v bool)`

SetSearchable sets Searchable field to given value.

### HasSearchable

`func (o *SObjectDescribesSobjectsInner) HasSearchable() bool`

HasSearchable returns a boolean if a field has been set.

### GetSobjectDescribeOption

`func (o *SObjectDescribesSobjectsInner) GetSobjectDescribeOption() string`

GetSobjectDescribeOption returns the SobjectDescribeOption field if non-nil, zero value otherwise.

### GetSobjectDescribeOptionOk

`func (o *SObjectDescribesSobjectsInner) GetSobjectDescribeOptionOk() (*string, bool)`

GetSobjectDescribeOptionOk returns a tuple with the SobjectDescribeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSobjectDescribeOption

`func (o *SObjectDescribesSobjectsInner) SetSobjectDescribeOption(v string)`

SetSobjectDescribeOption sets SobjectDescribeOption field to given value.

### HasSobjectDescribeOption

`func (o *SObjectDescribesSobjectsInner) HasSobjectDescribeOption() bool`

HasSobjectDescribeOption returns a boolean if a field has been set.

### GetTriggerable

`func (o *SObjectDescribesSobjectsInner) GetTriggerable() bool`

GetTriggerable returns the Triggerable field if non-nil, zero value otherwise.

### GetTriggerableOk

`func (o *SObjectDescribesSobjectsInner) GetTriggerableOk() (*bool, bool)`

GetTriggerableOk returns a tuple with the Triggerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerable

`func (o *SObjectDescribesSobjectsInner) SetTriggerable(v bool)`

SetTriggerable sets Triggerable field to given value.

### HasTriggerable

`func (o *SObjectDescribesSobjectsInner) HasTriggerable() bool`

HasTriggerable returns a boolean if a field has been set.

### GetUndeletable

`func (o *SObjectDescribesSobjectsInner) GetUndeletable() bool`

GetUndeletable returns the Undeletable field if non-nil, zero value otherwise.

### GetUndeletableOk

`func (o *SObjectDescribesSobjectsInner) GetUndeletableOk() (*bool, bool)`

GetUndeletableOk returns a tuple with the Undeletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndeletable

`func (o *SObjectDescribesSobjectsInner) SetUndeletable(v bool)`

SetUndeletable sets Undeletable field to given value.

### HasUndeletable

`func (o *SObjectDescribesSobjectsInner) HasUndeletable() bool`

HasUndeletable returns a boolean if a field has been set.

### GetUpdateable

`func (o *SObjectDescribesSobjectsInner) GetUpdateable() bool`

GetUpdateable returns the Updateable field if non-nil, zero value otherwise.

### GetUpdateableOk

`func (o *SObjectDescribesSobjectsInner) GetUpdateableOk() (*bool, bool)`

GetUpdateableOk returns a tuple with the Updateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateable

`func (o *SObjectDescribesSobjectsInner) SetUpdateable(v bool)`

SetUpdateable sets Updateable field to given value.

### HasUpdateable

`func (o *SObjectDescribesSobjectsInner) HasUpdateable() bool`

HasUpdateable returns a boolean if a field has been set.

### GetNamedLayoutInfos

`func (o *SObjectDescribesSobjectsInner) GetNamedLayoutInfos() []interface{}`

GetNamedLayoutInfos returns the NamedLayoutInfos field if non-nil, zero value otherwise.

### GetNamedLayoutInfosOk

`func (o *SObjectDescribesSobjectsInner) GetNamedLayoutInfosOk() (*[]interface{}, bool)`

GetNamedLayoutInfosOk returns a tuple with the NamedLayoutInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedLayoutInfos

`func (o *SObjectDescribesSobjectsInner) SetNamedLayoutInfos(v []interface{})`

SetNamedLayoutInfos sets NamedLayoutInfos field to given value.

### HasNamedLayoutInfos

`func (o *SObjectDescribesSobjectsInner) HasNamedLayoutInfos() bool`

HasNamedLayoutInfos returns a boolean if a field has been set.

### GetActionOverrides

`func (o *SObjectDescribesSobjectsInner) GetActionOverrides() []SObjectDescribeActionOverrides`

GetActionOverrides returns the ActionOverrides field if non-nil, zero value otherwise.

### GetActionOverridesOk

`func (o *SObjectDescribesSobjectsInner) GetActionOverridesOk() (*[]SObjectDescribeActionOverrides, bool)`

GetActionOverridesOk returns a tuple with the ActionOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOverrides

`func (o *SObjectDescribesSobjectsInner) SetActionOverrides(v []SObjectDescribeActionOverrides)`

SetActionOverrides sets ActionOverrides field to given value.

### HasActionOverrides

`func (o *SObjectDescribesSobjectsInner) HasActionOverrides() bool`

HasActionOverrides returns a boolean if a field has been set.

### GetChildRelationships

`func (o *SObjectDescribesSobjectsInner) GetChildRelationships() []SObjectDescribeChildRelationships`

GetChildRelationships returns the ChildRelationships field if non-nil, zero value otherwise.

### GetChildRelationshipsOk

`func (o *SObjectDescribesSobjectsInner) GetChildRelationshipsOk() (*[]SObjectDescribeChildRelationships, bool)`

GetChildRelationshipsOk returns a tuple with the ChildRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildRelationships

`func (o *SObjectDescribesSobjectsInner) SetChildRelationships(v []SObjectDescribeChildRelationships)`

SetChildRelationships sets ChildRelationships field to given value.

### HasChildRelationships

`func (o *SObjectDescribesSobjectsInner) HasChildRelationships() bool`

HasChildRelationships returns a boolean if a field has been set.

### GetRecordTypeInfos

`func (o *SObjectDescribesSobjectsInner) GetRecordTypeInfos() []SObjectRecordTypeInfo`

GetRecordTypeInfos returns the RecordTypeInfos field if non-nil, zero value otherwise.

### GetRecordTypeInfosOk

`func (o *SObjectDescribesSobjectsInner) GetRecordTypeInfosOk() (*[]SObjectRecordTypeInfo, bool)`

GetRecordTypeInfosOk returns a tuple with the RecordTypeInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTypeInfos

`func (o *SObjectDescribesSobjectsInner) SetRecordTypeInfos(v []SObjectRecordTypeInfo)`

SetRecordTypeInfos sets RecordTypeInfos field to given value.

### HasRecordTypeInfos

`func (o *SObjectDescribesSobjectsInner) HasRecordTypeInfos() bool`

HasRecordTypeInfos returns a boolean if a field has been set.

### GetSupportedScopes

`func (o *SObjectDescribesSobjectsInner) GetSupportedScopes() []SObjectDescribeSupportedScopes`

GetSupportedScopes returns the SupportedScopes field if non-nil, zero value otherwise.

### GetSupportedScopesOk

`func (o *SObjectDescribesSobjectsInner) GetSupportedScopesOk() (*[]SObjectDescribeSupportedScopes, bool)`

GetSupportedScopesOk returns a tuple with the SupportedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedScopes

`func (o *SObjectDescribesSobjectsInner) SetSupportedScopes(v []SObjectDescribeSupportedScopes)`

SetSupportedScopes sets SupportedScopes field to given value.

### HasSupportedScopes

`func (o *SObjectDescribesSobjectsInner) HasSupportedScopes() bool`

HasSupportedScopes returns a boolean if a field has been set.

### GetUrls

`func (o *SObjectDescribesSobjectsInner) GetUrls() map[string]interface{}`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *SObjectDescribesSobjectsInner) GetUrlsOk() (*map[string]interface{}, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *SObjectDescribesSobjectsInner) SetUrls(v map[string]interface{})`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *SObjectDescribesSobjectsInner) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetFields

`func (o *SObjectDescribesSobjectsInner) GetFields() []SObjectFieldDescribe`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SObjectDescribesSobjectsInner) GetFieldsOk() (*[]SObjectFieldDescribe, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SObjectDescribesSobjectsInner) SetFields(v []SObjectFieldDescribe)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SObjectDescribesSobjectsInner) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


