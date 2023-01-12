# BulkQueryJobsInfos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Done** | **bool** | This is true if this is the last (or only) set of results. It is false if there are more records to fetch. | 
**Records** | [**[]BulkQueryJobInfo**](BulkQueryJobInfo.md) | An array of record objects. | 
**NextRecordsUrl** | Pointer to **string** | The URI to get the next set of records (if there are any). This method returns up to 1,000 result rows per request. If there are more than 1,000 records, use the nextRecordsUrl to get the next set of records. This parameter is null if there are no more records to fetch.  | [optional] 

## Methods

### NewBulkQueryJobsInfos

`func NewBulkQueryJobsInfos(done bool, records []BulkQueryJobInfo, ) *BulkQueryJobsInfos`

NewBulkQueryJobsInfos instantiates a new BulkQueryJobsInfos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkQueryJobsInfosWithDefaults

`func NewBulkQueryJobsInfosWithDefaults() *BulkQueryJobsInfos`

NewBulkQueryJobsInfosWithDefaults instantiates a new BulkQueryJobsInfos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDone

`func (o *BulkQueryJobsInfos) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *BulkQueryJobsInfos) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *BulkQueryJobsInfos) SetDone(v bool)`

SetDone sets Done field to given value.


### GetRecords

`func (o *BulkQueryJobsInfos) GetRecords() []BulkQueryJobInfo`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *BulkQueryJobsInfos) GetRecordsOk() (*[]BulkQueryJobInfo, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *BulkQueryJobsInfos) SetRecords(v []BulkQueryJobInfo)`

SetRecords sets Records field to given value.


### GetNextRecordsUrl

`func (o *BulkQueryJobsInfos) GetNextRecordsUrl() string`

GetNextRecordsUrl returns the NextRecordsUrl field if non-nil, zero value otherwise.

### GetNextRecordsUrlOk

`func (o *BulkQueryJobsInfos) GetNextRecordsUrlOk() (*string, bool)`

GetNextRecordsUrlOk returns a tuple with the NextRecordsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRecordsUrl

`func (o *BulkQueryJobsInfos) SetNextRecordsUrl(v string)`

SetNextRecordsUrl sets NextRecordsUrl field to given value.

### HasNextRecordsUrl

`func (o *BulkQueryJobsInfos) HasNextRecordsUrl() bool`

HasNextRecordsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


