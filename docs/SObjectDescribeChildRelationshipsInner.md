# SObjectDescribeChildRelationshipsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CascadeDelete** | Pointer to **bool** |  | [optional] 
**ChildSObject** | Pointer to **string** |  | [optional] 
**DeprecatedAndHidden** | Pointer to **bool** |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**JunctionIdListNames** | Pointer to **[]map[string]interface{}** |  | [optional] 
**JunctionReferenceTo** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RelationshipName** | Pointer to **string** |  | [optional] 
**RestrictedDelete** | Pointer to **bool** |  | [optional] 

## Methods

### NewSObjectDescribeChildRelationshipsInner

`func NewSObjectDescribeChildRelationshipsInner() *SObjectDescribeChildRelationshipsInner`

NewSObjectDescribeChildRelationshipsInner instantiates a new SObjectDescribeChildRelationshipsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSObjectDescribeChildRelationshipsInnerWithDefaults

`func NewSObjectDescribeChildRelationshipsInnerWithDefaults() *SObjectDescribeChildRelationshipsInner`

NewSObjectDescribeChildRelationshipsInnerWithDefaults instantiates a new SObjectDescribeChildRelationshipsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCascadeDelete

`func (o *SObjectDescribeChildRelationshipsInner) GetCascadeDelete() bool`

GetCascadeDelete returns the CascadeDelete field if non-nil, zero value otherwise.

### GetCascadeDeleteOk

`func (o *SObjectDescribeChildRelationshipsInner) GetCascadeDeleteOk() (*bool, bool)`

GetCascadeDeleteOk returns a tuple with the CascadeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadeDelete

`func (o *SObjectDescribeChildRelationshipsInner) SetCascadeDelete(v bool)`

SetCascadeDelete sets CascadeDelete field to given value.

### HasCascadeDelete

`func (o *SObjectDescribeChildRelationshipsInner) HasCascadeDelete() bool`

HasCascadeDelete returns a boolean if a field has been set.

### GetChildSObject

`func (o *SObjectDescribeChildRelationshipsInner) GetChildSObject() string`

GetChildSObject returns the ChildSObject field if non-nil, zero value otherwise.

### GetChildSObjectOk

`func (o *SObjectDescribeChildRelationshipsInner) GetChildSObjectOk() (*string, bool)`

GetChildSObjectOk returns a tuple with the ChildSObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildSObject

`func (o *SObjectDescribeChildRelationshipsInner) SetChildSObject(v string)`

SetChildSObject sets ChildSObject field to given value.

### HasChildSObject

`func (o *SObjectDescribeChildRelationshipsInner) HasChildSObject() bool`

HasChildSObject returns a boolean if a field has been set.

### GetDeprecatedAndHidden

`func (o *SObjectDescribeChildRelationshipsInner) GetDeprecatedAndHidden() bool`

GetDeprecatedAndHidden returns the DeprecatedAndHidden field if non-nil, zero value otherwise.

### GetDeprecatedAndHiddenOk

`func (o *SObjectDescribeChildRelationshipsInner) GetDeprecatedAndHiddenOk() (*bool, bool)`

GetDeprecatedAndHiddenOk returns a tuple with the DeprecatedAndHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedAndHidden

`func (o *SObjectDescribeChildRelationshipsInner) SetDeprecatedAndHidden(v bool)`

SetDeprecatedAndHidden sets DeprecatedAndHidden field to given value.

### HasDeprecatedAndHidden

`func (o *SObjectDescribeChildRelationshipsInner) HasDeprecatedAndHidden() bool`

HasDeprecatedAndHidden returns a boolean if a field has been set.

### GetField

`func (o *SObjectDescribeChildRelationshipsInner) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SObjectDescribeChildRelationshipsInner) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SObjectDescribeChildRelationshipsInner) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *SObjectDescribeChildRelationshipsInner) HasField() bool`

HasField returns a boolean if a field has been set.

### GetJunctionIdListNames

`func (o *SObjectDescribeChildRelationshipsInner) GetJunctionIdListNames() []map[string]interface{}`

GetJunctionIdListNames returns the JunctionIdListNames field if non-nil, zero value otherwise.

### GetJunctionIdListNamesOk

`func (o *SObjectDescribeChildRelationshipsInner) GetJunctionIdListNamesOk() (*[]map[string]interface{}, bool)`

GetJunctionIdListNamesOk returns a tuple with the JunctionIdListNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJunctionIdListNames

`func (o *SObjectDescribeChildRelationshipsInner) SetJunctionIdListNames(v []map[string]interface{})`

SetJunctionIdListNames sets JunctionIdListNames field to given value.

### HasJunctionIdListNames

`func (o *SObjectDescribeChildRelationshipsInner) HasJunctionIdListNames() bool`

HasJunctionIdListNames returns a boolean if a field has been set.

### GetJunctionReferenceTo

`func (o *SObjectDescribeChildRelationshipsInner) GetJunctionReferenceTo() []map[string]interface{}`

GetJunctionReferenceTo returns the JunctionReferenceTo field if non-nil, zero value otherwise.

### GetJunctionReferenceToOk

`func (o *SObjectDescribeChildRelationshipsInner) GetJunctionReferenceToOk() (*[]map[string]interface{}, bool)`

GetJunctionReferenceToOk returns a tuple with the JunctionReferenceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJunctionReferenceTo

`func (o *SObjectDescribeChildRelationshipsInner) SetJunctionReferenceTo(v []map[string]interface{})`

SetJunctionReferenceTo sets JunctionReferenceTo field to given value.

### HasJunctionReferenceTo

`func (o *SObjectDescribeChildRelationshipsInner) HasJunctionReferenceTo() bool`

HasJunctionReferenceTo returns a boolean if a field has been set.

### GetRelationshipName

`func (o *SObjectDescribeChildRelationshipsInner) GetRelationshipName() string`

GetRelationshipName returns the RelationshipName field if non-nil, zero value otherwise.

### GetRelationshipNameOk

`func (o *SObjectDescribeChildRelationshipsInner) GetRelationshipNameOk() (*string, bool)`

GetRelationshipNameOk returns a tuple with the RelationshipName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipName

`func (o *SObjectDescribeChildRelationshipsInner) SetRelationshipName(v string)`

SetRelationshipName sets RelationshipName field to given value.

### HasRelationshipName

`func (o *SObjectDescribeChildRelationshipsInner) HasRelationshipName() bool`

HasRelationshipName returns a boolean if a field has been set.

### GetRestrictedDelete

`func (o *SObjectDescribeChildRelationshipsInner) GetRestrictedDelete() bool`

GetRestrictedDelete returns the RestrictedDelete field if non-nil, zero value otherwise.

### GetRestrictedDeleteOk

`func (o *SObjectDescribeChildRelationshipsInner) GetRestrictedDeleteOk() (*bool, bool)`

GetRestrictedDeleteOk returns a tuple with the RestrictedDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDelete

`func (o *SObjectDescribeChildRelationshipsInner) SetRestrictedDelete(v bool)`

SetRestrictedDelete sets RestrictedDelete field to given value.

### HasRestrictedDelete

`func (o *SObjectDescribeChildRelationshipsInner) HasRestrictedDelete() bool`

HasRestrictedDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


