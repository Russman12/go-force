# SObjectResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectDescribe** | Pointer to [**SObjectDescribe**](SObjectDescribe.md) |  | [optional] 
**RecentItems** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewSObjectResult

`func NewSObjectResult() *SObjectResult`

NewSObjectResult instantiates a new SObjectResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSObjectResultWithDefaults

`func NewSObjectResultWithDefaults() *SObjectResult`

NewSObjectResultWithDefaults instantiates a new SObjectResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectDescribe

`func (o *SObjectResult) GetObjectDescribe() SObjectDescribe`

GetObjectDescribe returns the ObjectDescribe field if non-nil, zero value otherwise.

### GetObjectDescribeOk

`func (o *SObjectResult) GetObjectDescribeOk() (*SObjectDescribe, bool)`

GetObjectDescribeOk returns a tuple with the ObjectDescribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDescribe

`func (o *SObjectResult) SetObjectDescribe(v SObjectDescribe)`

SetObjectDescribe sets ObjectDescribe field to given value.

### HasObjectDescribe

`func (o *SObjectResult) HasObjectDescribe() bool`

HasObjectDescribe returns a boolean if a field has been set.

### GetRecentItems

`func (o *SObjectResult) GetRecentItems() []map[string]interface{}`

GetRecentItems returns the RecentItems field if non-nil, zero value otherwise.

### GetRecentItemsOk

`func (o *SObjectResult) GetRecentItemsOk() (*[]map[string]interface{}, bool)`

GetRecentItemsOk returns a tuple with the RecentItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentItems

`func (o *SObjectResult) SetRecentItems(v []map[string]interface{})`

SetRecentItems sets RecentItems field to given value.

### HasRecentItems

`func (o *SObjectResult) HasRecentItems() bool`

HasRecentItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


