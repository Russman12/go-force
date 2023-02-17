# RecordTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Indicates whether this record type is available (true) or not (false). Availability is used to display a list of available record types to the user when they are creating a new record. | [optional] 
**DefaultRecordTypeMapping** | Pointer to **bool** | Indicates whether this is the default record type mapping (true) or not (false). | [optional] 
**DeveloperName** | Pointer to **string** | Developer name of this record type. Available in API versions 43.0 and later. | [optional] 
**Master** | Pointer to **bool** | Indicates whether this is the main record type (true) or not (false). The main record type is the default record type that’s used when a record has no custom record type associated with it. | [optional] 
**Name** | Pointer to **string** | Name of this record type. | [optional] 
**RecordTypeId** | Pointer to **string** | ID of this record type. | [optional] 

## Methods

### NewRecordTypeInfo

`func NewRecordTypeInfo() *RecordTypeInfo`

NewRecordTypeInfo instantiates a new RecordTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordTypeInfoWithDefaults

`func NewRecordTypeInfoWithDefaults() *RecordTypeInfo`

NewRecordTypeInfoWithDefaults instantiates a new RecordTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *RecordTypeInfo) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *RecordTypeInfo) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *RecordTypeInfo) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *RecordTypeInfo) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetDefaultRecordTypeMapping

`func (o *RecordTypeInfo) GetDefaultRecordTypeMapping() bool`

GetDefaultRecordTypeMapping returns the DefaultRecordTypeMapping field if non-nil, zero value otherwise.

### GetDefaultRecordTypeMappingOk

`func (o *RecordTypeInfo) GetDefaultRecordTypeMappingOk() (*bool, bool)`

GetDefaultRecordTypeMappingOk returns a tuple with the DefaultRecordTypeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRecordTypeMapping

`func (o *RecordTypeInfo) SetDefaultRecordTypeMapping(v bool)`

SetDefaultRecordTypeMapping sets DefaultRecordTypeMapping field to given value.

### HasDefaultRecordTypeMapping

`func (o *RecordTypeInfo) HasDefaultRecordTypeMapping() bool`

HasDefaultRecordTypeMapping returns a boolean if a field has been set.

### GetDeveloperName

`func (o *RecordTypeInfo) GetDeveloperName() string`

GetDeveloperName returns the DeveloperName field if non-nil, zero value otherwise.

### GetDeveloperNameOk

`func (o *RecordTypeInfo) GetDeveloperNameOk() (*string, bool)`

GetDeveloperNameOk returns a tuple with the DeveloperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperName

`func (o *RecordTypeInfo) SetDeveloperName(v string)`

SetDeveloperName sets DeveloperName field to given value.

### HasDeveloperName

`func (o *RecordTypeInfo) HasDeveloperName() bool`

HasDeveloperName returns a boolean if a field has been set.

### GetMaster

`func (o *RecordTypeInfo) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *RecordTypeInfo) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *RecordTypeInfo) SetMaster(v bool)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *RecordTypeInfo) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetName

`func (o *RecordTypeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordTypeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordTypeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordTypeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecordTypeId

`func (o *RecordTypeInfo) GetRecordTypeId() string`

GetRecordTypeId returns the RecordTypeId field if non-nil, zero value otherwise.

### GetRecordTypeIdOk

`func (o *RecordTypeInfo) GetRecordTypeIdOk() (*string, bool)`

GetRecordTypeIdOk returns a tuple with the RecordTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTypeId

`func (o *RecordTypeInfo) SetRecordTypeId(v string)`

SetRecordTypeId sets RecordTypeId field to given value.

### HasRecordTypeId

`func (o *RecordTypeInfo) HasRecordTypeId() bool`

HasRecordTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


