# ChildRelationships

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

### NewChildRelationships

`func NewChildRelationships() *ChildRelationships`

NewChildRelationships instantiates a new ChildRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildRelationshipsWithDefaults

`func NewChildRelationshipsWithDefaults() *ChildRelationships`

NewChildRelationshipsWithDefaults instantiates a new ChildRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCascadeDelete

`func (o *ChildRelationships) GetCascadeDelete() bool`

GetCascadeDelete returns the CascadeDelete field if non-nil, zero value otherwise.

### GetCascadeDeleteOk

`func (o *ChildRelationships) GetCascadeDeleteOk() (*bool, bool)`

GetCascadeDeleteOk returns a tuple with the CascadeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadeDelete

`func (o *ChildRelationships) SetCascadeDelete(v bool)`

SetCascadeDelete sets CascadeDelete field to given value.

### HasCascadeDelete

`func (o *ChildRelationships) HasCascadeDelete() bool`

HasCascadeDelete returns a boolean if a field has been set.

### GetChildSObject

`func (o *ChildRelationships) GetChildSObject() string`

GetChildSObject returns the ChildSObject field if non-nil, zero value otherwise.

### GetChildSObjectOk

`func (o *ChildRelationships) GetChildSObjectOk() (*string, bool)`

GetChildSObjectOk returns a tuple with the ChildSObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildSObject

`func (o *ChildRelationships) SetChildSObject(v string)`

SetChildSObject sets ChildSObject field to given value.

### HasChildSObject

`func (o *ChildRelationships) HasChildSObject() bool`

HasChildSObject returns a boolean if a field has been set.

### GetDeprecatedAndHidden

`func (o *ChildRelationships) GetDeprecatedAndHidden() bool`

GetDeprecatedAndHidden returns the DeprecatedAndHidden field if non-nil, zero value otherwise.

### GetDeprecatedAndHiddenOk

`func (o *ChildRelationships) GetDeprecatedAndHiddenOk() (*bool, bool)`

GetDeprecatedAndHiddenOk returns a tuple with the DeprecatedAndHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedAndHidden

`func (o *ChildRelationships) SetDeprecatedAndHidden(v bool)`

SetDeprecatedAndHidden sets DeprecatedAndHidden field to given value.

### HasDeprecatedAndHidden

`func (o *ChildRelationships) HasDeprecatedAndHidden() bool`

HasDeprecatedAndHidden returns a boolean if a field has been set.

### GetField

`func (o *ChildRelationships) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ChildRelationships) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ChildRelationships) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *ChildRelationships) HasField() bool`

HasField returns a boolean if a field has been set.

### GetJunctionIdListNames

`func (o *ChildRelationships) GetJunctionIdListNames() []string`

GetJunctionIdListNames returns the JunctionIdListNames field if non-nil, zero value otherwise.

### GetJunctionIdListNamesOk

`func (o *ChildRelationships) GetJunctionIdListNamesOk() (*[]string, bool)`

GetJunctionIdListNamesOk returns a tuple with the JunctionIdListNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJunctionIdListNames

`func (o *ChildRelationships) SetJunctionIdListNames(v []string)`

SetJunctionIdListNames sets JunctionIdListNames field to given value.

### HasJunctionIdListNames

`func (o *ChildRelationships) HasJunctionIdListNames() bool`

HasJunctionIdListNames returns a boolean if a field has been set.

### GetJunctionReferenceTo

`func (o *ChildRelationships) GetJunctionReferenceTo() []string`

GetJunctionReferenceTo returns the JunctionReferenceTo field if non-nil, zero value otherwise.

### GetJunctionReferenceToOk

`func (o *ChildRelationships) GetJunctionReferenceToOk() (*[]string, bool)`

GetJunctionReferenceToOk returns a tuple with the JunctionReferenceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJunctionReferenceTo

`func (o *ChildRelationships) SetJunctionReferenceTo(v []string)`

SetJunctionReferenceTo sets JunctionReferenceTo field to given value.

### HasJunctionReferenceTo

`func (o *ChildRelationships) HasJunctionReferenceTo() bool`

HasJunctionReferenceTo returns a boolean if a field has been set.

### GetRelationshipName

`func (o *ChildRelationships) GetRelationshipName() string`

GetRelationshipName returns the RelationshipName field if non-nil, zero value otherwise.

### GetRelationshipNameOk

`func (o *ChildRelationships) GetRelationshipNameOk() (*string, bool)`

GetRelationshipNameOk returns a tuple with the RelationshipName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipName

`func (o *ChildRelationships) SetRelationshipName(v string)`

SetRelationshipName sets RelationshipName field to given value.

### HasRelationshipName

`func (o *ChildRelationships) HasRelationshipName() bool`

HasRelationshipName returns a boolean if a field has been set.

### GetRestrictedDelete

`func (o *ChildRelationships) GetRestrictedDelete() bool`

GetRestrictedDelete returns the RestrictedDelete field if non-nil, zero value otherwise.

### GetRestrictedDeleteOk

`func (o *ChildRelationships) GetRestrictedDeleteOk() (*bool, bool)`

GetRestrictedDeleteOk returns a tuple with the RestrictedDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDelete

`func (o *ChildRelationships) SetRestrictedDelete(v bool)`

SetRestrictedDelete sets RestrictedDelete field to given value.

### HasRestrictedDelete

`func (o *ChildRelationships) HasRestrictedDelete() bool`

HasRestrictedDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


