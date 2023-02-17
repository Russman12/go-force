# ScopeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | UI label for this scope. | [optional] 
**Name** | Pointer to **string** | Name of this scope. | [optional] 

## Methods

### NewScopeInfo

`func NewScopeInfo() *ScopeInfo`

NewScopeInfo instantiates a new ScopeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeInfoWithDefaults

`func NewScopeInfoWithDefaults() *ScopeInfo`

NewScopeInfoWithDefaults instantiates a new ScopeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ScopeInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ScopeInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ScopeInfo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ScopeInfo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *ScopeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScopeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScopeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScopeInfo) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


