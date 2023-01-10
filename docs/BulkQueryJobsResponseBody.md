# BulkQueryJobsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Done** | Pointer to **bool** |  | [optional] 
**Records** | Pointer to [**[]BulkQueryJobsResponseBodyRecords**](BulkQueryJobsResponseBodyRecords.md) |  | [optional] 
**NextRecordsUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkQueryJobsResponseBody

`func NewBulkQueryJobsResponseBody() *BulkQueryJobsResponseBody`

NewBulkQueryJobsResponseBody instantiates a new BulkQueryJobsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkQueryJobsResponseBodyWithDefaults

`func NewBulkQueryJobsResponseBodyWithDefaults() *BulkQueryJobsResponseBody`

NewBulkQueryJobsResponseBodyWithDefaults instantiates a new BulkQueryJobsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDone

`func (o *BulkQueryJobsResponseBody) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *BulkQueryJobsResponseBody) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *BulkQueryJobsResponseBody) SetDone(v bool)`

SetDone sets Done field to given value.

### HasDone

`func (o *BulkQueryJobsResponseBody) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetRecords

`func (o *BulkQueryJobsResponseBody) GetRecords() []BulkQueryJobsResponseBodyRecords`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *BulkQueryJobsResponseBody) GetRecordsOk() (*[]BulkQueryJobsResponseBodyRecords, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *BulkQueryJobsResponseBody) SetRecords(v []BulkQueryJobsResponseBodyRecords)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *BulkQueryJobsResponseBody) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetNextRecordsUrl

`func (o *BulkQueryJobsResponseBody) GetNextRecordsUrl() string`

GetNextRecordsUrl returns the NextRecordsUrl field if non-nil, zero value otherwise.

### GetNextRecordsUrlOk

`func (o *BulkQueryJobsResponseBody) GetNextRecordsUrlOk() (*string, bool)`

GetNextRecordsUrlOk returns a tuple with the NextRecordsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRecordsUrl

`func (o *BulkQueryJobsResponseBody) SetNextRecordsUrl(v string)`

SetNextRecordsUrl sets NextRecordsUrl field to given value.

### HasNextRecordsUrl

`func (o *BulkQueryJobsResponseBody) HasNextRecordsUrl() bool`

HasNextRecordsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


