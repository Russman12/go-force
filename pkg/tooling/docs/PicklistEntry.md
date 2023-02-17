# PicklistEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether this item must be displayed (true) or not (false) in the drop-down list for the picklist field in the user interface. | [optional] 
**ValidFor** | Pointer to **string** | A set of bits where each bit indicates a controlling value for which this &#x60;PicklistEntry&#x60; is valid. See About Dependent Picklists. | [optional] 
**DefaultValue** | Pointer to **bool** | Indicates whether this item is the default item (true) in the picklist or not (false). Only one item in a picklist can be designated as the default. | [optional] 
**Label** | Pointer to **string** | Display name of this item in the picklist. | [optional] 
**Value** | Pointer to **string** | Value of this item in the picklist. | [optional] 

## Methods

### NewPicklistEntry

`func NewPicklistEntry() *PicklistEntry`

NewPicklistEntry instantiates a new PicklistEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPicklistEntryWithDefaults

`func NewPicklistEntryWithDefaults() *PicklistEntry`

NewPicklistEntryWithDefaults instantiates a new PicklistEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *PicklistEntry) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PicklistEntry) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PicklistEntry) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PicklistEntry) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetValidFor

`func (o *PicklistEntry) GetValidFor() string`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *PicklistEntry) GetValidForOk() (*string, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *PicklistEntry) SetValidFor(v string)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *PicklistEntry) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.

### GetDefaultValue

`func (o *PicklistEntry) GetDefaultValue() bool`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *PicklistEntry) GetDefaultValueOk() (*bool, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *PicklistEntry) SetDefaultValue(v bool)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *PicklistEntry) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetLabel

`func (o *PicklistEntry) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PicklistEntry) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PicklistEntry) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PicklistEntry) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetValue

`func (o *PicklistEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PicklistEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PicklistEntry) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PicklistEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


