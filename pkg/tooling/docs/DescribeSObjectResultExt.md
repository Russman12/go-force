# DescribeSObjectResultExt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewDescribeSObjectResultExt

`func NewDescribeSObjectResultExt() *DescribeSObjectResultExt`

NewDescribeSObjectResultExt instantiates a new DescribeSObjectResultExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeSObjectResultExtWithDefaults

`func NewDescribeSObjectResultExtWithDefaults() *DescribeSObjectResultExt`

NewDescribeSObjectResultExtWithDefaults instantiates a new DescribeSObjectResultExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionOverrides

`func (o *DescribeSObjectResultExt) GetActionOverrides() []ActionOverride`

GetActionOverrides returns the ActionOverrides field if non-nil, zero value otherwise.

### GetActionOverridesOk

`func (o *DescribeSObjectResultExt) GetActionOverridesOk() (*[]ActionOverride, bool)`

GetActionOverridesOk returns a tuple with the ActionOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOverrides

`func (o *DescribeSObjectResultExt) SetActionOverrides(v []ActionOverride)`

SetActionOverrides sets ActionOverrides field to given value.

### HasActionOverrides

`func (o *DescribeSObjectResultExt) HasActionOverrides() bool`

HasActionOverrides returns a boolean if a field has been set.

### GetChildRelationships

`func (o *DescribeSObjectResultExt) GetChildRelationships() []ChildRelationship`

GetChildRelationships returns the ChildRelationships field if non-nil, zero value otherwise.

### GetChildRelationshipsOk

`func (o *DescribeSObjectResultExt) GetChildRelationshipsOk() (*[]ChildRelationship, bool)`

GetChildRelationshipsOk returns a tuple with the ChildRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildRelationships

`func (o *DescribeSObjectResultExt) SetChildRelationships(v []ChildRelationship)`

SetChildRelationships sets ChildRelationships field to given value.

### HasChildRelationships

`func (o *DescribeSObjectResultExt) HasChildRelationships() bool`

HasChildRelationships returns a boolean if a field has been set.

### GetCompactLayoutable

`func (o *DescribeSObjectResultExt) GetCompactLayoutable() bool`

GetCompactLayoutable returns the CompactLayoutable field if non-nil, zero value otherwise.

### GetCompactLayoutableOk

`func (o *DescribeSObjectResultExt) GetCompactLayoutableOk() (*bool, bool)`

GetCompactLayoutableOk returns a tuple with the CompactLayoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactLayoutable

`func (o *DescribeSObjectResultExt) SetCompactLayoutable(v bool)`

SetCompactLayoutable sets CompactLayoutable field to given value.

### HasCompactLayoutable

`func (o *DescribeSObjectResultExt) HasCompactLayoutable() bool`

HasCompactLayoutable returns a boolean if a field has been set.

### GetDefaultImplementation

`func (o *DescribeSObjectResultExt) GetDefaultImplementation() string`

GetDefaultImplementation returns the DefaultImplementation field if non-nil, zero value otherwise.

### GetDefaultImplementationOk

`func (o *DescribeSObjectResultExt) GetDefaultImplementationOk() (*string, bool)`

GetDefaultImplementationOk returns a tuple with the DefaultImplementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImplementation

`func (o *DescribeSObjectResultExt) SetDefaultImplementation(v string)`

SetDefaultImplementation sets DefaultImplementation field to given value.

### HasDefaultImplementation

`func (o *DescribeSObjectResultExt) HasDefaultImplementation() bool`

HasDefaultImplementation returns a boolean if a field has been set.

### GetExtendedBy

`func (o *DescribeSObjectResultExt) GetExtendedBy() string`

GetExtendedBy returns the ExtendedBy field if non-nil, zero value otherwise.

### GetExtendedByOk

`func (o *DescribeSObjectResultExt) GetExtendedByOk() (*string, bool)`

GetExtendedByOk returns a tuple with the ExtendedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedBy

`func (o *DescribeSObjectResultExt) SetExtendedBy(v string)`

SetExtendedBy sets ExtendedBy field to given value.

### HasExtendedBy

`func (o *DescribeSObjectResultExt) HasExtendedBy() bool`

HasExtendedBy returns a boolean if a field has been set.

### GetExtendsInterfaces

`func (o *DescribeSObjectResultExt) GetExtendsInterfaces() string`

GetExtendsInterfaces returns the ExtendsInterfaces field if non-nil, zero value otherwise.

### GetExtendsInterfacesOk

`func (o *DescribeSObjectResultExt) GetExtendsInterfacesOk() (*string, bool)`

GetExtendsInterfacesOk returns a tuple with the ExtendsInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendsInterfaces

`func (o *DescribeSObjectResultExt) SetExtendsInterfaces(v string)`

SetExtendsInterfaces sets ExtendsInterfaces field to given value.

### HasExtendsInterfaces

`func (o *DescribeSObjectResultExt) HasExtendsInterfaces() bool`

HasExtendsInterfaces returns a boolean if a field has been set.

### GetFields

`func (o *DescribeSObjectResultExt) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DescribeSObjectResultExt) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DescribeSObjectResultExt) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DescribeSObjectResultExt) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetImplementedBy

`func (o *DescribeSObjectResultExt) GetImplementedBy() string`

GetImplementedBy returns the ImplementedBy field if non-nil, zero value otherwise.

### GetImplementedByOk

`func (o *DescribeSObjectResultExt) GetImplementedByOk() (*string, bool)`

GetImplementedByOk returns a tuple with the ImplementedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementedBy

`func (o *DescribeSObjectResultExt) SetImplementedBy(v string)`

SetImplementedBy sets ImplementedBy field to given value.

### HasImplementedBy

`func (o *DescribeSObjectResultExt) HasImplementedBy() bool`

HasImplementedBy returns a boolean if a field has been set.

### GetImplementsInterfaces

`func (o *DescribeSObjectResultExt) GetImplementsInterfaces() string`

GetImplementsInterfaces returns the ImplementsInterfaces field if non-nil, zero value otherwise.

### GetImplementsInterfacesOk

`func (o *DescribeSObjectResultExt) GetImplementsInterfacesOk() (*string, bool)`

GetImplementsInterfacesOk returns a tuple with the ImplementsInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementsInterfaces

`func (o *DescribeSObjectResultExt) SetImplementsInterfaces(v string)`

SetImplementsInterfaces sets ImplementsInterfaces field to given value.

### HasImplementsInterfaces

`func (o *DescribeSObjectResultExt) HasImplementsInterfaces() bool`

HasImplementsInterfaces returns a boolean if a field has been set.

### GetListviewable

`func (o *DescribeSObjectResultExt) GetListviewable() string`

GetListviewable returns the Listviewable field if non-nil, zero value otherwise.

### GetListviewableOk

`func (o *DescribeSObjectResultExt) GetListviewableOk() (*string, bool)`

GetListviewableOk returns a tuple with the Listviewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListviewable

`func (o *DescribeSObjectResultExt) SetListviewable(v string)`

SetListviewable sets Listviewable field to given value.

### HasListviewable

`func (o *DescribeSObjectResultExt) HasListviewable() bool`

HasListviewable returns a boolean if a field has been set.

### GetLookupLayoutable

`func (o *DescribeSObjectResultExt) GetLookupLayoutable() string`

GetLookupLayoutable returns the LookupLayoutable field if non-nil, zero value otherwise.

### GetLookupLayoutableOk

`func (o *DescribeSObjectResultExt) GetLookupLayoutableOk() (*string, bool)`

GetLookupLayoutableOk returns a tuple with the LookupLayoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupLayoutable

`func (o *DescribeSObjectResultExt) SetLookupLayoutable(v string)`

SetLookupLayoutable sets LookupLayoutable field to given value.

### HasLookupLayoutable

`func (o *DescribeSObjectResultExt) HasLookupLayoutable() bool`

HasLookupLayoutable returns a boolean if a field has been set.

### GetNamedLayoutInfos

`func (o *DescribeSObjectResultExt) GetNamedLayoutInfos() []NamedLayoutInfo`

GetNamedLayoutInfos returns the NamedLayoutInfos field if non-nil, zero value otherwise.

### GetNamedLayoutInfosOk

`func (o *DescribeSObjectResultExt) GetNamedLayoutInfosOk() (*[]NamedLayoutInfo, bool)`

GetNamedLayoutInfosOk returns a tuple with the NamedLayoutInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedLayoutInfos

`func (o *DescribeSObjectResultExt) SetNamedLayoutInfos(v []NamedLayoutInfo)`

SetNamedLayoutInfos sets NamedLayoutInfos field to given value.

### HasNamedLayoutInfos

`func (o *DescribeSObjectResultExt) HasNamedLayoutInfos() bool`

HasNamedLayoutInfos returns a boolean if a field has been set.

### GetNetworkScopeFieldName

`func (o *DescribeSObjectResultExt) GetNetworkScopeFieldName() string`

GetNetworkScopeFieldName returns the NetworkScopeFieldName field if non-nil, zero value otherwise.

### GetNetworkScopeFieldNameOk

`func (o *DescribeSObjectResultExt) GetNetworkScopeFieldNameOk() (*string, bool)`

GetNetworkScopeFieldNameOk returns a tuple with the NetworkScopeFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkScopeFieldName

`func (o *DescribeSObjectResultExt) SetNetworkScopeFieldName(v string)`

SetNetworkScopeFieldName sets NetworkScopeFieldName field to given value.

### HasNetworkScopeFieldName

`func (o *DescribeSObjectResultExt) HasNetworkScopeFieldName() bool`

HasNetworkScopeFieldName returns a boolean if a field has been set.

### GetRecordTypeInfos

`func (o *DescribeSObjectResultExt) GetRecordTypeInfos() []RecordTypeInfo`

GetRecordTypeInfos returns the RecordTypeInfos field if non-nil, zero value otherwise.

### GetRecordTypeInfosOk

`func (o *DescribeSObjectResultExt) GetRecordTypeInfosOk() (*[]RecordTypeInfo, bool)`

GetRecordTypeInfosOk returns a tuple with the RecordTypeInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTypeInfos

`func (o *DescribeSObjectResultExt) SetRecordTypeInfos(v []RecordTypeInfo)`

SetRecordTypeInfos sets RecordTypeInfos field to given value.

### HasRecordTypeInfos

`func (o *DescribeSObjectResultExt) HasRecordTypeInfos() bool`

HasRecordTypeInfos returns a boolean if a field has been set.

### GetSearchLayoutable

`func (o *DescribeSObjectResultExt) GetSearchLayoutable() bool`

GetSearchLayoutable returns the SearchLayoutable field if non-nil, zero value otherwise.

### GetSearchLayoutableOk

`func (o *DescribeSObjectResultExt) GetSearchLayoutableOk() (*bool, bool)`

GetSearchLayoutableOk returns a tuple with the SearchLayoutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchLayoutable

`func (o *DescribeSObjectResultExt) SetSearchLayoutable(v bool)`

SetSearchLayoutable sets SearchLayoutable field to given value.

### HasSearchLayoutable

`func (o *DescribeSObjectResultExt) HasSearchLayoutable() bool`

HasSearchLayoutable returns a boolean if a field has been set.

### GetSobjectDescribeOption

`func (o *DescribeSObjectResultExt) GetSobjectDescribeOption() string`

GetSobjectDescribeOption returns the SobjectDescribeOption field if non-nil, zero value otherwise.

### GetSobjectDescribeOptionOk

`func (o *DescribeSObjectResultExt) GetSobjectDescribeOptionOk() (*string, bool)`

GetSobjectDescribeOptionOk returns a tuple with the SobjectDescribeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSobjectDescribeOption

`func (o *DescribeSObjectResultExt) SetSobjectDescribeOption(v string)`

SetSobjectDescribeOption sets SobjectDescribeOption field to given value.

### HasSobjectDescribeOption

`func (o *DescribeSObjectResultExt) HasSobjectDescribeOption() bool`

HasSobjectDescribeOption returns a boolean if a field has been set.

### GetSupportedScopes

`func (o *DescribeSObjectResultExt) GetSupportedScopes() []ScopeInfo`

GetSupportedScopes returns the SupportedScopes field if non-nil, zero value otherwise.

### GetSupportedScopesOk

`func (o *DescribeSObjectResultExt) GetSupportedScopesOk() (*[]ScopeInfo, bool)`

GetSupportedScopesOk returns a tuple with the SupportedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedScopes

`func (o *DescribeSObjectResultExt) SetSupportedScopes(v []ScopeInfo)`

SetSupportedScopes sets SupportedScopes field to given value.

### HasSupportedScopes

`func (o *DescribeSObjectResultExt) HasSupportedScopes() bool`

HasSupportedScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


