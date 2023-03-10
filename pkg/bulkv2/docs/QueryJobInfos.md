# QueryJobInfos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Done** | Pointer to **bool** | This is true if this is the last (or only) set of results. It is false if there are more records to fetch. | [optional] 
**Records** | Pointer to [**[]QueryJobInfo**](QueryJobInfo.md) | An array of record objects. | [optional] 
**NextRecordsUrl** | Pointer to **string** | The URI to get the next set of records (if there are any). This method returns up to 1,000 result rows per request. If there are more than 1,000 records, use the nextRecordsUrl to get the next set of records. This parameter is null if there are no more records to fetch.  | [optional] 

## Methods

### NewQueryJobInfos

`func NewQueryJobInfos() *QueryJobInfos`

NewQueryJobInfos instantiates a new QueryJobInfos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryJobInfosWithDefaults

`func NewQueryJobInfosWithDefaults() *QueryJobInfos`

NewQueryJobInfosWithDefaults instantiates a new QueryJobInfos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDone

`func (o *QueryJobInfos) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *QueryJobInfos) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *QueryJobInfos) SetDone(v bool)`

SetDone sets Done field to given value.

### HasDone

`func (o *QueryJobInfos) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetRecords

`func (o *QueryJobInfos) GetRecords() []QueryJobInfo`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *QueryJobInfos) GetRecordsOk() (*[]QueryJobInfo, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *QueryJobInfos) SetRecords(v []QueryJobInfo)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *QueryJobInfos) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetNextRecordsUrl

`func (o *QueryJobInfos) GetNextRecordsUrl() string`

GetNextRecordsUrl returns the NextRecordsUrl field if non-nil, zero value otherwise.

### GetNextRecordsUrlOk

`func (o *QueryJobInfos) GetNextRecordsUrlOk() (*string, bool)`

GetNextRecordsUrlOk returns a tuple with the NextRecordsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRecordsUrl

`func (o *QueryJobInfos) SetNextRecordsUrl(v string)`

SetNextRecordsUrl sets NextRecordsUrl field to given value.

### HasNextRecordsUrl

`func (o *QueryJobInfos) HasNextRecordsUrl() bool`

HasNextRecordsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


