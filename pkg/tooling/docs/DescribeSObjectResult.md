# DescribeSObjectResult

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
**Urls** | Pointer to [**SObjectDescribeURL**](SObjectDescribeURL.md) |  | [optional] 
**ActionOverrides** | Pointer to [**[]ActionOverride**](ActionOverride.md) |  | [optional] 
**ChildRelationships** | Pointer to [**[]ChildRelationship**](ChildRelationship.md) |  | [optional] 
**CompactLayoutable** | Pointer to **bool** |  | [optional] 
**DefaultImplementation** | Pointer to **string** |  | [optional] 
**ExtendedBy** | Pointer to **string** |  | [optional] 
**ExtendsInterfaces** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**[]Field**](Field.md) |  | [optional] 
**ImplementedBy** | Pointer to **string** |  | [optional] 
**ImplementsInterfaces** | Pointer to **string** |  | [optional] 
**Listviewable** | Pointer to **string** |  | [optional] 
**LookupLayoutable** | Pointer to **string** |  | [optional] 
**NamedLayoutInfos** | Pointer to [**[]NamedLayoutInfo**](NamedLayoutInfo.md) |  | [optional] 
**NetworkScopeFieldName** | Pointer to **string** |  | [optional] 
**RecordTypeInfos** | Pointer to [**[]RecordTypeInfo**](RecordTypeInfo.md) |  | [optional] 
**SearchLayoutable** | Pointer to **bool** |  | [optional] 
**SobjectDescribeOption** | Pointer to **string** |  | [optional] 
**SupportedScopes** | Pointer to [**[]ScopeInfo**](ScopeInfo.md) |  | [optional] 

## Methods

### NewDescribeSObjectResult

`func NewDescribeSObjectResult() *DescribeSObjectResult`

NewDescribeSObjectResult instantiates a new DescribeSObjectResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeSObjectResultWithDefaults

`func NewDescribeSObjectResultWithDefaults() *DescribeSObjectResult`

NewDescribeSObjectResultWithDefaults instantiates a new DescribeSObjectResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivateable

`func (o *DescribeSObjectResult) GetActivateable() bool`

GetActivateable returns the Activateable field if non-nil, zero value otherwise.

### GetActivateableOk

`func (o *DescribeSObjectResult) GetActivateableOk() (*bool, bool)`

GetActivateableOk returns a tuple with the Activateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateable

`func (o *DescribeSObjectResult) SetActivateable(v bool)`

SetActivateable sets Activateable field to given value.

### HasActivateable

`func (o *DescribeSObjectResult) HasActivateable() bool`

HasActivateable returns a boolean if a field has been set.

### GetAssociateEntityType

`func (o *DescribeSObjectResult) GetAssociateEntityType() string`

GetAssociateEntityType returns the AssociateEntityType field if non-nil, zero value otherwise.

### GetAssociateEntityTypeOk

`func (o *DescribeSObjectResult) GetAssociateEntityTypeOk() (*string, bool)`

GetAssociateEntityTypeOk returns a tuple with the AssociateEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociateEntityType

`func (o *DescribeSObjectResult) SetAssociateEntityType(v string)`

SetAssociateEntityType sets AssociateEntityType field to given value.

### HasAssociateEntityType

`func (o *DescribeSObjectResult) HasAssociateEntityType() bool`

HasAssociateEntityType returns a boolean if a field has been set.

### GetAssociateParentEntity

`func (o *DescribeSObjectResult) GetAssociateParentEntity() string`

GetAssociateParentEntity returns the AssociateParentEntity field if non-nil, zero value otherwise.

### GetAssociateParentEntityOk

`func (o *DescribeSObjectResult) GetAssociateParentEntityOk() (*string, bool)`

GetAssociateParentEntityOk returns a tuple with the AssociateParentEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociateParentEntity

`func (o *DescribeSObjectResult) SetAssociateParentEntity(v string)`

SetAssociateParentEntity sets AssociateParentEntity field to given value.

### HasAssociateParentEntity

`func (o *DescribeSObjectResult) HasAssociateParentEntity() bool`

HasAssociateParentEntity returns a boolean if a field has been set.

### GetCreateable

`func (o *DescribeSObjectResult) GetCreateable() bool`

GetCreateable returns the Createable field if non-nil, zero value otherwise.

### GetCreateableOk

`func (o *DescribeSObjectResult) GetCreateableOk() (*bool, bool)`

GetCreateableOk returns a tuple with the Createable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateable

`func (o *DescribeSObjectResult) SetCreateable(v bool)`

SetCreateable sets Createable field to given value.

### HasCreateable

`func (o *DescribeSObjectResult) HasCreateable() bool`

HasCreateable returns a boolean if a field has been set.

### GetCustom

`func (o *DescribeSObjectResult) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *DescribeSObjectResult) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *DescribeSObjectResult) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *DescribeSObjectResult) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomSetting

`func (o *DescribeSObjectResult) GetCustomSetting() bool`

GetCustomSetting returns the CustomSetting field if non-nil, zero value otherwise.

### GetCustomSettingOk

`func (o *DescribeSObjectResult) GetCustomSettingOk() (*bool, bool)`

GetCustomSettingOk returns a tuple with the CustomSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSetting

`func (o *DescribeSObjectResult) SetCustomSetting(v bool)`

SetCustomSetting sets CustomSetting field to given value.

### HasCustomSetting

`func (o *DescribeSObjectResult) HasCustomSetting() bool`

HasCustomSetting returns a boolean if a field has been set.

### GetDeepCloneable

`func (o *DescribeSObjectResult) GetDeepCloneable() bool`

GetDeepCloneable returns the DeepCloneable field if non-nil, zero value otherwise.

### GetDeepCloneableOk

`func (o *DescribeSObjectResult) GetDeepCloneableOk() (*bool, bool)`

GetDeepCloneableOk returns a tuple with the DeepCloneable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepCloneable

`func (o *DescribeSObjectResult) SetDeepCloneable(v bool)`

SetDeepCloneable sets DeepCloneable field to given value.

### HasDeepCloneable

`func (o *DescribeSObjectResult) HasDeepCloneable() bool`

HasDeepCloneable returns a boolean if a field has been set.

### GetDeletable

`func (o *DescribeSObjectResult) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *DescribeSObjectResult) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *DescribeSObjectResult) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *DescribeSObjectResult) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDeprecatedAndHidden

`func (o *DescribeSObjectResult) GetDeprecatedAndHidden() bool`

GetDeprecatedAndHidden returns the DeprecatedAndHidden field if non-nil, zero value otherwise.

### GetDeprecatedAndHiddenOk

`func (o *DescribeSObjectResult) GetDeprecatedAndHiddenOk() (*bool, bool)`

GetDeprecatedAndHiddenOk returns a tuple with the DeprecatedAndHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedAndHidden

`func (o *DescribeSObjectResult) SetDeprecatedAndHidden(v bool)`

SetDeprecatedAndHidden sets DeprecatedAndHidden field to given value.

### HasDeprecatedAndHidden

`func (o *DescribeSObjectResult) HasDeprecatedAndHidden() bool`

HasDeprecatedAndHidden returns a boolean if a field has been set.

### GetFeedEnabled

`func (o *DescribeSObjectResult) GetFeedEnabled() bool`

GetFeedEnabled returns the FeedEnabled field if non-nil, zero value otherwise.

### GetFeedEnabledOk

`func (o *DescribeSObjectResult) GetFeedEnabledOk() (*bool, bool)`

GetFeedEnabledOk returns a tuple with the FeedEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedEnabled

`func (o *DescribeSObjectResult) SetFeedEnabled(v bool)`

SetFeedEnabled sets FeedEnabled field to given value.

### HasFeedEnabled

`func (o *DescribeSObjectResult) HasFeedEnabled() bool`

HasFeedEnabled returns a boolean if a field has been set.

### GetHasSubtypes

`func (o *DescribeSObjectResult) GetHasSubtypes() bool`

GetHasSubtypes returns the HasSubtypes field if non-nil, zero value otherwise.

### GetHasSubtypesOk

`func (o *DescribeSObjectResult) GetHasSubtypesOk() (*bool, bool)`

GetHasSubtypesOk returns a tuple with the HasSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSubtypes

`func (o *DescribeSObjectResult) SetHasSubtypes(v bool)`

SetHasSubtypes sets HasSubtypes field to given value.

### HasHasSubtypes

`func (o *DescribeSObjectResult) HasHasSubtypes() bool`

HasHasSubtypes returns a boolean if a field has been set.

### GetIsInterface

`func (o *DescribeSObjectResult) GetIsInterface() bool`

GetIsInterface returns the IsInterface field if non-nil, zero value otherwise.

### GetIsInterfaceOk

`func (o *DescribeSObjectResult) GetIsInterfaceOk() (*bool, bool)`

GetIsInterfaceOk returns a tuple with the IsInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInterface

`func (o *DescribeSObjectResult) SetIsInterface(v bool)`

SetIsInterface sets IsInterface field to given value.

### HasIsInterface

`func (o *DescribeSObjectResult) HasIsInterface() bool`

HasIsInterface returns a boolean if a field has been set.

### GetIsSubtype

`func (o *DescribeSObjectResult) GetIsSubtype() bool`

GetIsSubtype returns the IsSubtype field if non-nil, zero value otherwise.

### GetIsSubtypeOk

`func (o *DescribeSObjectResult) GetIsSubtypeOk() (*bool, bool)`

GetIsSubtypeOk returns a tuple with the IsSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubtype

`func (o *DescribeSObjectResult) SetIsSubtype(v bool)`

SetIsSubtype sets IsSubtype field to given value.

### HasIsSubtype

`func (o *DescribeSObjectResult) HasIsSubtype() bool`

HasIsSubtype returns a boolean if a field has been set.

### GetKeyPrefix

`func (o *DescribeSObjectResult) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *DescribeSObjectResult) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *DescribeSObjectResult) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.

### HasKeyPrefix

`func (o *DescribeSObjectResult) HasKeyPrefix() bool`

HasKeyPrefix returns a boolean if a field has been set.

### GetLabel

`func (o *DescribeSObjectResult) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DescribeSObjectResult) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DescribeSObjectResult) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DescribeSObjectResult) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLabelPlural

`func (o *DescribeSObjectResult) GetLabelPlural() string`

GetLabelPlural returns the LabelPlural field if non-nil, zero value otherwise.

### GetLabelPluralOk

`func (o *DescribeSObjectResult) GetLabelPluralOk() (*string, bool)`

GetLabelPluralOk returns a tuple with the LabelPlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPlural

`func (o *DescribeSObjectResult) SetLabelPlural(v string)`

SetLabelPlural sets LabelPlural field to given value.

### HasLabelPlural

`func (o *DescribeSObjectResult) HasLabelPlural() bool`

HasLabelPlural returns a boolean if a field has been set.

### GetLayoutable

`func (o *DescribeSObjectResult) GetLayoutable() bool`

GetLayoutable returns the Layoutable field if non-nil, zero value otherwise.

### GetLayoutableOk

`func (o *DescribeSObjectResult) GetLayoutableOk() (*bool, bool)`

GetLayoutableOk returns a tuple with the Layoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutable

`func (o *DescribeSObjectResult) SetLayoutable(v bool)`

SetLayoutable sets Layoutable field to given value.

### HasLayoutable

`func (o *DescribeSObjectResult) HasLayoutable() bool`

HasLayoutable returns a boolean if a field has been set.

### GetMergeable

`func (o *DescribeSObjectResult) GetMergeable() bool`

GetMergeable returns the Mergeable field if non-nil, zero value otherwise.

### GetMergeableOk

`func (o *DescribeSObjectResult) GetMergeableOk() (*bool, bool)`

GetMergeableOk returns a tuple with the Mergeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeable

`func (o *DescribeSObjectResult) SetMergeable(v bool)`

SetMergeable sets Mergeable field to given value.

### HasMergeable

`func (o *DescribeSObjectResult) HasMergeable() bool`

HasMergeable returns a boolean if a field has been set.

### GetMruEnabled

`func (o *DescribeSObjectResult) GetMruEnabled() bool`

GetMruEnabled returns the MruEnabled field if non-nil, zero value otherwise.

### GetMruEnabledOk

`func (o *DescribeSObjectResult) GetMruEnabledOk() (*bool, bool)`

GetMruEnabledOk returns a tuple with the MruEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMruEnabled

`func (o *DescribeSObjectResult) SetMruEnabled(v bool)`

SetMruEnabled sets MruEnabled field to given value.

### HasMruEnabled

`func (o *DescribeSObjectResult) HasMruEnabled() bool`

HasMruEnabled returns a boolean if a field has been set.

### GetName

`func (o *DescribeSObjectResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeSObjectResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeSObjectResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DescribeSObjectResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQueryable

`func (o *DescribeSObjectResult) GetQueryable() bool`

GetQueryable returns the Queryable field if non-nil, zero value otherwise.

### GetQueryableOk

`func (o *DescribeSObjectResult) GetQueryableOk() (*bool, bool)`

GetQueryableOk returns a tuple with the Queryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryable

`func (o *DescribeSObjectResult) SetQueryable(v bool)`

SetQueryable sets Queryable field to given value.

### HasQueryable

`func (o *DescribeSObjectResult) HasQueryable() bool`

HasQueryable returns a boolean if a field has been set.

### GetReplicateable

`func (o *DescribeSObjectResult) GetReplicateable() bool`

GetReplicateable returns the Replicateable field if non-nil, zero value otherwise.

### GetReplicateableOk

`func (o *DescribeSObjectResult) GetReplicateableOk() (*bool, bool)`

GetReplicateableOk returns a tuple with the Replicateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateable

`func (o *DescribeSObjectResult) SetReplicateable(v bool)`

SetReplicateable sets Replicateable field to given value.

### HasReplicateable

`func (o *DescribeSObjectResult) HasReplicateable() bool`

HasReplicateable returns a boolean if a field has been set.

### GetRetrieveable

`func (o *DescribeSObjectResult) GetRetrieveable() bool`

GetRetrieveable returns the Retrieveable field if non-nil, zero value otherwise.

### GetRetrieveableOk

`func (o *DescribeSObjectResult) GetRetrieveableOk() (*bool, bool)`

GetRetrieveableOk returns a tuple with the Retrieveable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveable

`func (o *DescribeSObjectResult) SetRetrieveable(v bool)`

SetRetrieveable sets Retrieveable field to given value.

### HasRetrieveable

`func (o *DescribeSObjectResult) HasRetrieveable() bool`

HasRetrieveable returns a boolean if a field has been set.

### GetSearchable

`func (o *DescribeSObjectResult) GetSearchable() bool`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *DescribeSObjectResult) GetSearchableOk() (*bool, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *DescribeSObjectResult) SetSearchable(v bool)`

SetSearchable sets Searchable field to given value.

### HasSearchable

`func (o *DescribeSObjectResult) HasSearchable() bool`

HasSearchable returns a boolean if a field has been set.

### GetTriggerable

`func (o *DescribeSObjectResult) GetTriggerable() bool`

GetTriggerable returns the Triggerable field if non-nil, zero value otherwise.

### GetTriggerableOk

`func (o *DescribeSObjectResult) GetTriggerableOk() (*bool, bool)`

GetTriggerableOk returns a tuple with the Triggerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerable

`func (o *DescribeSObjectResult) SetTriggerable(v bool)`

SetTriggerable sets Triggerable field to given value.

### HasTriggerable

`func (o *DescribeSObjectResult) HasTriggerable() bool`

HasTriggerable returns a boolean if a field has been set.

### GetUndeletable

`func (o *DescribeSObjectResult) GetUndeletable() bool`

GetUndeletable returns the Undeletable field if non-nil, zero value otherwise.

### GetUndeletableOk

`func (o *DescribeSObjectResult) GetUndeletableOk() (*bool, bool)`

GetUndeletableOk returns a tuple with the Undeletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndeletable

`func (o *DescribeSObjectResult) SetUndeletable(v bool)`

SetUndeletable sets Undeletable field to given value.

### HasUndeletable

`func (o *DescribeSObjectResult) HasUndeletable() bool`

HasUndeletable returns a boolean if a field has been set.

### GetUpdateable

`func (o *DescribeSObjectResult) GetUpdateable() bool`

GetUpdateable returns the Updateable field if non-nil, zero value otherwise.

### GetUpdateableOk

`func (o *DescribeSObjectResult) GetUpdateableOk() (*bool, bool)`

GetUpdateableOk returns a tuple with the Updateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateable

`func (o *DescribeSObjectResult) SetUpdateable(v bool)`

SetUpdateable sets Updateable field to given value.

### HasUpdateable

`func (o *DescribeSObjectResult) HasUpdateable() bool`

HasUpdateable returns a boolean if a field has been set.

### GetUrls

`func (o *DescribeSObjectResult) GetUrls() SObjectDescribeURL`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *DescribeSObjectResult) GetUrlsOk() (*SObjectDescribeURL, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *DescribeSObjectResult) SetUrls(v SObjectDescribeURL)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *DescribeSObjectResult) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetActionOverrides

`func (o *DescribeSObjectResult) GetActionOverrides() []ActionOverride`

GetActionOverrides returns the ActionOverrides field if non-nil, zero value otherwise.

### GetActionOverridesOk

`func (o *DescribeSObjectResult) GetActionOverridesOk() (*[]ActionOverride, bool)`

GetActionOverridesOk returns a tuple with the ActionOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOverrides

`func (o *DescribeSObjectResult) SetActionOverrides(v []ActionOverride)`

SetActionOverrides sets ActionOverrides field to given value.

### HasActionOverrides

`func (o *DescribeSObjectResult) HasActionOverrides() bool`

HasActionOverrides returns a boolean if a field has been set.

### GetChildRelationships

`func (o *DescribeSObjectResult) GetChildRelationships() []ChildRelationship`

GetChildRelationships returns the ChildRelationships field if non-nil, zero value otherwise.

### GetChildRelationshipsOk

`func (o *DescribeSObjectResult) GetChildRelationshipsOk() (*[]ChildRelationship, bool)`

GetChildRelationshipsOk returns a tuple with the ChildRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildRelationships

`func (o *DescribeSObjectResult) SetChildRelationships(v []ChildRelationship)`

SetChildRelationships sets ChildRelationships field to given value.

### HasChildRelationships

`func (o *DescribeSObjectResult) HasChildRelationships() bool`

HasChildRelationships returns a boolean if a field has been set.

### GetCompactLayoutable

`func (o *DescribeSObjectResult) GetCompactLayoutable() bool`

GetCompactLayoutable returns the CompactLayoutable field if non-nil, zero value otherwise.

### GetCompactLayoutableOk

`func (o *DescribeSObjectResult) GetCompactLayoutableOk() (*bool, bool)`

GetCompactLayoutableOk returns a tuple with the CompactLayoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactLayoutable

`func (o *DescribeSObjectResult) SetCompactLayoutable(v bool)`

SetCompactLayoutable sets CompactLayoutable field to given value.

### HasCompactLayoutable

`func (o *DescribeSObjectResult) HasCompactLayoutable() bool`

HasCompactLayoutable returns a boolean if a field has been set.

### GetDefaultImplementation

`func (o *DescribeSObjectResult) GetDefaultImplementation() string`

GetDefaultImplementation returns the DefaultImplementation field if non-nil, zero value otherwise.

### GetDefaultImplementationOk

`func (o *DescribeSObjectResult) GetDefaultImplementationOk() (*string, bool)`

GetDefaultImplementationOk returns a tuple with the DefaultImplementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImplementation

`func (o *DescribeSObjectResult) SetDefaultImplementation(v string)`

SetDefaultImplementation sets DefaultImplementation field to given value.

### HasDefaultImplementation

`func (o *DescribeSObjectResult) HasDefaultImplementation() bool`

HasDefaultImplementation returns a boolean if a field has been set.

### GetExtendedBy

`func (o *DescribeSObjectResult) GetExtendedBy() string`

GetExtendedBy returns the ExtendedBy field if non-nil, zero value otherwise.

### GetExtendedByOk

`func (o *DescribeSObjectResult) GetExtendedByOk() (*string, bool)`

GetExtendedByOk returns a tuple with the ExtendedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedBy

`func (o *DescribeSObjectResult) SetExtendedBy(v string)`

SetExtendedBy sets ExtendedBy field to given value.

### HasExtendedBy

`func (o *DescribeSObjectResult) HasExtendedBy() bool`

HasExtendedBy returns a boolean if a field has been set.

### GetExtendsInterfaces

`func (o *DescribeSObjectResult) GetExtendsInterfaces() string`

GetExtendsInterfaces returns the ExtendsInterfaces field if non-nil, zero value otherwise.

### GetExtendsInterfacesOk

`func (o *DescribeSObjectResult) GetExtendsInterfacesOk() (*string, bool)`

GetExtendsInterfacesOk returns a tuple with the ExtendsInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendsInterfaces

`func (o *DescribeSObjectResult) SetExtendsInterfaces(v string)`

SetExtendsInterfaces sets ExtendsInterfaces field to given value.

### HasExtendsInterfaces

`func (o *DescribeSObjectResult) HasExtendsInterfaces() bool`

HasExtendsInterfaces returns a boolean if a field has been set.

### GetFields

`func (o *DescribeSObjectResult) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DescribeSObjectResult) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DescribeSObjectResult) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DescribeSObjectResult) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetImplementedBy

`func (o *DescribeSObjectResult) GetImplementedBy() string`

GetImplementedBy returns the ImplementedBy field if non-nil, zero value otherwise.

### GetImplementedByOk

`func (o *DescribeSObjectResult) GetImplementedByOk() (*string, bool)`

GetImplementedByOk returns a tuple with the ImplementedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementedBy

`func (o *DescribeSObjectResult) SetImplementedBy(v string)`

SetImplementedBy sets ImplementedBy field to given value.

### HasImplementedBy

`func (o *DescribeSObjectResult) HasImplementedBy() bool`

HasImplementedBy returns a boolean if a field has been set.

### GetImplementsInterfaces

`func (o *DescribeSObjectResult) GetImplementsInterfaces() string`

GetImplementsInterfaces returns the ImplementsInterfaces field if non-nil, zero value otherwise.

### GetImplementsInterfacesOk

`func (o *DescribeSObjectResult) GetImplementsInterfacesOk() (*string, bool)`

GetImplementsInterfacesOk returns a tuple with the ImplementsInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementsInterfaces

`func (o *DescribeSObjectResult) SetImplementsInterfaces(v string)`

SetImplementsInterfaces sets ImplementsInterfaces field to given value.

### HasImplementsInterfaces

`func (o *DescribeSObjectResult) HasImplementsInterfaces() bool`

HasImplementsInterfaces returns a boolean if a field has been set.

### GetListviewable

`func (o *DescribeSObjectResult) GetListviewable() string`

GetListviewable returns the Listviewable field if non-nil, zero value otherwise.

### GetListviewableOk

`func (o *DescribeSObjectResult) GetListviewableOk() (*string, bool)`

GetListviewableOk returns a tuple with the Listviewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListviewable

`func (o *DescribeSObjectResult) SetListviewable(v string)`

SetListviewable sets Listviewable field to given value.

### HasListviewable

`func (o *DescribeSObjectResult) HasListviewable() bool`

HasListviewable returns a boolean if a field has been set.

### GetLookupLayoutable

`func (o *DescribeSObjectResult) GetLookupLayoutable() string`

GetLookupLayoutable returns the LookupLayoutable field if non-nil, zero value otherwise.

### GetLookupLayoutableOk

`func (o *DescribeSObjectResult) GetLookupLayoutableOk() (*string, bool)`

GetLookupLayoutableOk returns a tuple with the LookupLayoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupLayoutable

`func (o *DescribeSObjectResult) SetLookupLayoutable(v string)`

SetLookupLayoutable sets LookupLayoutable field to given value.

### HasLookupLayoutable

`func (o *DescribeSObjectResult) HasLookupLayoutable() bool`

HasLookupLayoutable returns a boolean if a field has been set.

### GetNamedLayoutInfos

`func (o *DescribeSObjectResult) GetNamedLayoutInfos() []NamedLayoutInfo`

GetNamedLayoutInfos returns the NamedLayoutInfos field if non-nil, zero value otherwise.

### GetNamedLayoutInfosOk

`func (o *DescribeSObjectResult) GetNamedLayoutInfosOk() (*[]NamedLayoutInfo, bool)`

GetNamedLayoutInfosOk returns a tuple with the NamedLayoutInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedLayoutInfos

`func (o *DescribeSObjectResult) SetNamedLayoutInfos(v []NamedLayoutInfo)`

SetNamedLayoutInfos sets NamedLayoutInfos field to given value.

### HasNamedLayoutInfos

`func (o *DescribeSObjectResult) HasNamedLayoutInfos() bool`

HasNamedLayoutInfos returns a boolean if a field has been set.

### GetNetworkScopeFieldName

`func (o *DescribeSObjectResult) GetNetworkScopeFieldName() string`

GetNetworkScopeFieldName returns the NetworkScopeFieldName field if non-nil, zero value otherwise.

### GetNetworkScopeFieldNameOk

`func (o *DescribeSObjectResult) GetNetworkScopeFieldNameOk() (*string, bool)`

GetNetworkScopeFieldNameOk returns a tuple with the NetworkScopeFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkScopeFieldName

`func (o *DescribeSObjectResult) SetNetworkScopeFieldName(v string)`

SetNetworkScopeFieldName sets NetworkScopeFieldName field to given value.

### HasNetworkScopeFieldName

`func (o *DescribeSObjectResult) HasNetworkScopeFieldName() bool`

HasNetworkScopeFieldName returns a boolean if a field has been set.

### GetRecordTypeInfos

`func (o *DescribeSObjectResult) GetRecordTypeInfos() []RecordTypeInfo`

GetRecordTypeInfos returns the RecordTypeInfos field if non-nil, zero value otherwise.

### GetRecordTypeInfosOk

`func (o *DescribeSObjectResult) GetRecordTypeInfosOk() (*[]RecordTypeInfo, bool)`

GetRecordTypeInfosOk returns a tuple with the RecordTypeInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTypeInfos

`func (o *DescribeSObjectResult) SetRecordTypeInfos(v []RecordTypeInfo)`

SetRecordTypeInfos sets RecordTypeInfos field to given value.

### HasRecordTypeInfos

`func (o *DescribeSObjectResult) HasRecordTypeInfos() bool`

HasRecordTypeInfos returns a boolean if a field has been set.

### GetSearchLayoutable

`func (o *DescribeSObjectResult) GetSearchLayoutable() bool`

GetSearchLayoutable returns the SearchLayoutable field if non-nil, zero value otherwise.

### GetSearchLayoutableOk

`func (o *DescribeSObjectResult) GetSearchLayoutableOk() (*bool, bool)`

GetSearchLayoutableOk returns a tuple with the SearchLayoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchLayoutable

`func (o *DescribeSObjectResult) SetSearchLayoutable(v bool)`

SetSearchLayoutable sets SearchLayoutable field to given value.

### HasSearchLayoutable

`func (o *DescribeSObjectResult) HasSearchLayoutable() bool`

HasSearchLayoutable returns a boolean if a field has been set.

### GetSobjectDescribeOption

`func (o *DescribeSObjectResult) GetSobjectDescribeOption() string`

GetSobjectDescribeOption returns the SobjectDescribeOption field if non-nil, zero value otherwise.

### GetSobjectDescribeOptionOk

`func (o *DescribeSObjectResult) GetSobjectDescribeOptionOk() (*string, bool)`

GetSobjectDescribeOptionOk returns a tuple with the SobjectDescribeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSobjectDescribeOption

`func (o *DescribeSObjectResult) SetSobjectDescribeOption(v string)`

SetSobjectDescribeOption sets SobjectDescribeOption field to given value.

### HasSobjectDescribeOption

`func (o *DescribeSObjectResult) HasSobjectDescribeOption() bool`

HasSobjectDescribeOption returns a boolean if a field has been set.

### GetSupportedScopes

`func (o *DescribeSObjectResult) GetSupportedScopes() []ScopeInfo`

GetSupportedScopes returns the SupportedScopes field if non-nil, zero value otherwise.

### GetSupportedScopesOk

`func (o *DescribeSObjectResult) GetSupportedScopesOk() (*[]ScopeInfo, bool)`

GetSupportedScopesOk returns a tuple with the SupportedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedScopes

`func (o *DescribeSObjectResult) SetSupportedScopes(v []ScopeInfo)`

SetSupportedScopes sets SupportedScopes field to given value.

### HasSupportedScopes

`func (o *DescribeSObjectResult) HasSupportedScopes() bool`

HasSupportedScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


