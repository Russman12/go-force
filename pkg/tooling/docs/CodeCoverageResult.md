# CodeCoverageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DmlInfo** | Pointer to [**CodeLocation**](CodeLocation.md) |  | [optional] 
**Id** | **string** |  | 
**LocationsNotCovered** | Pointer to [**[]CodeLocation**](CodeLocation.md) |  | [optional] 
**MethodInfo** | Pointer to [**CodeLocation**](CodeLocation.md) |  | [optional] 
**Name** | **string** |  | 
**Namespace** | **string** |  | 
**NumLocations** | **int32** |  | 
**SoqlInfo** | Pointer to [**CodeLocation**](CodeLocation.md) |  | [optional] 
**SoslInfo** | Pointer to [**CodeLocation**](CodeLocation.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewCodeCoverageResult

`func NewCodeCoverageResult(id string, name string, namespace string, numLocations int32, type_ string, ) *CodeCoverageResult`

NewCodeCoverageResult instantiates a new CodeCoverageResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeCoverageResultWithDefaults

`func NewCodeCoverageResultWithDefaults() *CodeCoverageResult`

NewCodeCoverageResultWithDefaults instantiates a new CodeCoverageResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDmlInfo

`func (o *CodeCoverageResult) GetDmlInfo() CodeLocation`

GetDmlInfo returns the DmlInfo field if non-nil, zero value otherwise.

### GetDmlInfoOk

`func (o *CodeCoverageResult) GetDmlInfoOk() (*CodeLocation, bool)`

GetDmlInfoOk returns a tuple with the DmlInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmlInfo

`func (o *CodeCoverageResult) SetDmlInfo(v CodeLocation)`

SetDmlInfo sets DmlInfo field to given value.

### HasDmlInfo

`func (o *CodeCoverageResult) HasDmlInfo() bool`

HasDmlInfo returns a boolean if a field has been set.

### GetId

`func (o *CodeCoverageResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodeCoverageResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodeCoverageResult) SetId(v string)`

SetId sets Id field to given value.


### GetLocationsNotCovered

`func (o *CodeCoverageResult) GetLocationsNotCovered() []CodeLocation`

GetLocationsNotCovered returns the LocationsNotCovered field if non-nil, zero value otherwise.

### GetLocationsNotCoveredOk

`func (o *CodeCoverageResult) GetLocationsNotCoveredOk() (*[]CodeLocation, bool)`

GetLocationsNotCoveredOk returns a tuple with the LocationsNotCovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationsNotCovered

`func (o *CodeCoverageResult) SetLocationsNotCovered(v []CodeLocation)`

SetLocationsNotCovered sets LocationsNotCovered field to given value.

### HasLocationsNotCovered

`func (o *CodeCoverageResult) HasLocationsNotCovered() bool`

HasLocationsNotCovered returns a boolean if a field has been set.

### GetMethodInfo

`func (o *CodeCoverageResult) GetMethodInfo() CodeLocation`

GetMethodInfo returns the MethodInfo field if non-nil, zero value otherwise.

### GetMethodInfoOk

`func (o *CodeCoverageResult) GetMethodInfoOk() (*CodeLocation, bool)`

GetMethodInfoOk returns a tuple with the MethodInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodInfo

`func (o *CodeCoverageResult) SetMethodInfo(v CodeLocation)`

SetMethodInfo sets MethodInfo field to given value.

### HasMethodInfo

`func (o *CodeCoverageResult) HasMethodInfo() bool`

HasMethodInfo returns a boolean if a field has been set.

### GetName

`func (o *CodeCoverageResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodeCoverageResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodeCoverageResult) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *CodeCoverageResult) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CodeCoverageResult) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CodeCoverageResult) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetNumLocations

`func (o *CodeCoverageResult) GetNumLocations() int32`

GetNumLocations returns the NumLocations field if non-nil, zero value otherwise.

### GetNumLocationsOk

`func (o *CodeCoverageResult) GetNumLocationsOk() (*int32, bool)`

GetNumLocationsOk returns a tuple with the NumLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLocations

`func (o *CodeCoverageResult) SetNumLocations(v int32)`

SetNumLocations sets NumLocations field to given value.


### GetSoqlInfo

`func (o *CodeCoverageResult) GetSoqlInfo() CodeLocation`

GetSoqlInfo returns the SoqlInfo field if non-nil, zero value otherwise.

### GetSoqlInfoOk

`func (o *CodeCoverageResult) GetSoqlInfoOk() (*CodeLocation, bool)`

GetSoqlInfoOk returns a tuple with the SoqlInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoqlInfo

`func (o *CodeCoverageResult) SetSoqlInfo(v CodeLocation)`

SetSoqlInfo sets SoqlInfo field to given value.

### HasSoqlInfo

`func (o *CodeCoverageResult) HasSoqlInfo() bool`

HasSoqlInfo returns a boolean if a field has been set.

### GetSoslInfo

`func (o *CodeCoverageResult) GetSoslInfo() CodeLocation`

GetSoslInfo returns the SoslInfo field if non-nil, zero value otherwise.

### GetSoslInfoOk

`func (o *CodeCoverageResult) GetSoslInfoOk() (*CodeLocation, bool)`

GetSoslInfoOk returns a tuple with the SoslInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoslInfo

`func (o *CodeCoverageResult) SetSoslInfo(v CodeLocation)`

SetSoslInfo sets SoslInfo field to given value.

### HasSoslInfo

`func (o *CodeCoverageResult) HasSoslInfo() bool`

HasSoslInfo returns a boolean if a field has been set.

### GetType

`func (o *CodeCoverageResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CodeCoverageResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CodeCoverageResult) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


