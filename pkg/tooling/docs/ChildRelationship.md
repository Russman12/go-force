# ChildRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CascadeDelete** | Pointer to **bool** |  | [optional] 
**ChildSObject** | Pointer to **string** |  | [optional] 
**DeprecatedAndHidden** | Pointer to **bool** |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**JunctionIdListNames** | Pointer to **[]string** |  | [optional] 
**JunctionReferenceTo** | Pointer to **[]string** |  | [optional] 
**RelationshipName** | Pointer to **string** |  | [optional] 
**RestrictedDelete** | Pointer to **bool** |  | [optional] 

## Methods

### NewChildRelationship

`func NewChildRelationship() *ChildRelationship`

NewChildRelationship instantiates a new ChildRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildRelationshipWithDefaults

`func NewChildRelationshipWithDefaults() *ChildRelationship`

NewChildRelationshipWithDefaults instantiates a new ChildRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCascadeDelete

`func (o *ChildRelationship) GetCascadeDelete() bool`

GetCascadeDelete returns the CascadeDelete field if non-nil, zero value otherwise.

### GetCascadeDeleteOk

`func (o *ChildRelationship) GetCascadeDeleteOk() (*bool, bool)`

GetCascadeDeleteOk returns a tuple with the CascadeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadeDelete

`func (o *ChildRelationship) SetCascadeDelete(v bool)`

SetCascadeDelete sets CascadeDelete field to given value.

### HasCascadeDelete

`func (o *ChildRelationship) HasCascadeDelete() bool`

HasCascadeDelete returns a boolean if a field has been set.

### GetChildSObject

`func (o *ChildRelationship) GetChildSObject() string`

GetChildSObject returns the ChildSObject field if non-nil, zero value otherwise.

### GetChildSObjectOk

`func (o *ChildRelationship) GetChildSObjectOk() (*string, bool)`

GetChildSObjectOk returns a tuple with the ChildSObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildSObject

`func (o *ChildRelationship) SetChildSObject(v string)`

SetChildSObject sets ChildSObject field to given value.

### HasChildSObject

`func (o *ChildRelationship) HasChildSObject() bool`

HasChildSObject returns a boolean if a field has been set.

### GetDeprecatedAndHidden

`func (o *ChildRelationship) GetDeprecatedAndHidden() bool`

GetDeprecatedAndHidden returns the DeprecatedAndHidden field if non-nil, zero value otherwise.

### GetDeprecatedAndHiddenOk

`func (o *ChildRelationship) GetDeprecatedAndHiddenOk() (*bool, bool)`

GetDeprecatedAndHiddenOk returns a tuple with the DeprecatedAndHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedAndHidden

`func (o *ChildRelationship) SetDeprecatedAndHidden(v bool)`

SetDeprecatedAndHidden sets DeprecatedAndHidden field to given value.

### HasDeprecatedAndHidden

`func (o *ChildRelationship) HasDeprecatedAndHidden() bool`

HasDeprecatedAndHidden returns a boolean if a field has been set.

### GetField

`func (o *ChildRelationship) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ChildRelationship) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ChildRelationship) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *ChildRelationship) HasField() bool`

HasField returns a boolean if a field has been set.

### GetJunctionIdListNames

`func (o *ChildRelationship) GetJunctionIdListNames() []string`

GetJunctionIdListNames returns the JunctionIdListNames field if non-nil, zero value otherwise.

### GetJunctionIdListNamesOk

`func (o *ChildRelationship) GetJunctionIdListNamesOk() (*[]string, bool)`

GetJunctionIdListNamesOk returns a tuple with the JunctionIdListNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJunctionIdListNames

`func (o *ChildRelationship) SetJunctionIdListNames(v []string)`

SetJunctionIdListNames sets JunctionIdListNames field to given value.

### HasJunctionIdListNames

`func (o *ChildRelationship) HasJunctionIdListNames() bool`

HasJunctionIdListNames returns a boolean if a field has been set.

### GetJunctionReferenceTo

`func (o *ChildRelationship) GetJunctionReferenceTo() []string`

GetJunctionReferenceTo returns the JunctionReferenceTo field if non-nil, zero value otherwise.

### GetJunctionReferenceToOk

`func (o *ChildRelationship) GetJunctionReferenceToOk() (*[]string, bool)`

GetJunctionReferenceToOk returns a tuple with the JunctionReferenceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJunctionReferenceTo

`func (o *ChildRelationship) SetJunctionReferenceTo(v []string)`

SetJunctionReferenceTo sets JunctionReferenceTo field to given value.

### HasJunctionReferenceTo

`func (o *ChildRelationship) HasJunctionReferenceTo() bool`

HasJunctionReferenceTo returns a boolean if a field has been set.

### GetRelationshipName

`func (o *ChildRelationship) GetRelationshipName() string`

GetRelationshipName returns the RelationshipName field if non-nil, zero value otherwise.

### GetRelationshipNameOk

`func (o *ChildRelationship) GetRelationshipNameOk() (*string, bool)`

GetRelationshipNameOk returns a tuple with the RelationshipName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipName

`func (o *ChildRelationship) SetRelationshipName(v string)`

SetRelationshipName sets RelationshipName field to given value.

### HasRelationshipName

`func (o *ChildRelationship) HasRelationshipName() bool`

HasRelationshipName returns a boolean if a field has been set.

### GetRestrictedDelete

`func (o *ChildRelationship) GetRestrictedDelete() bool`

GetRestrictedDelete returns the RestrictedDelete field if non-nil, zero value otherwise.

### GetRestrictedDeleteOk

`func (o *ChildRelationship) GetRestrictedDeleteOk() (*bool, bool)`

GetRestrictedDeleteOk returns a tuple with the RestrictedDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDelete

`func (o *ChildRelationship) SetRestrictedDelete(v bool)`

SetRestrictedDelete sets RestrictedDelete field to given value.

### HasRestrictedDelete

`func (o *ChildRelationship) HasRestrictedDelete() bool`

HasRestrictedDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


