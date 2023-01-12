# Jobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Done** | **bool** | Indicates whether there are more jobs to get. If false, use the nextRecordsUrl value to retrieve the next group of jobs. | 
**Records** | [**[]JobInfo**](JobInfo.md) | Contains information for each retrieved job. | 
**NextRecordsUrl** | Pointer to **string** | A URL that contains a query locator used to get the next set of results in a subsequent request if done isn’t true. | [optional] 

## Methods

### NewJobs

`func NewJobs(done bool, records []JobInfo, ) *Jobs`

NewJobs instantiates a new Jobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobsWithDefaults

`func NewJobsWithDefaults() *Jobs`

NewJobsWithDefaults instantiates a new Jobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDone

`func (o *Jobs) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *Jobs) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *Jobs) SetDone(v bool)`

SetDone sets Done field to given value.


### GetRecords

`func (o *Jobs) GetRecords() []JobInfo`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *Jobs) GetRecordsOk() (*[]JobInfo, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *Jobs) SetRecords(v []JobInfo)`

SetRecords sets Records field to given value.


### GetNextRecordsUrl

`func (o *Jobs) GetNextRecordsUrl() string`

GetNextRecordsUrl returns the NextRecordsUrl field if non-nil, zero value otherwise.

### GetNextRecordsUrlOk

`func (o *Jobs) GetNextRecordsUrlOk() (*string, bool)`

GetNextRecordsUrlOk returns a tuple with the NextRecordsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRecordsUrl

`func (o *Jobs) SetNextRecordsUrl(v string)`

SetNextRecordsUrl sets NextRecordsUrl field to given value.

### HasNextRecordsUrl

`func (o *Jobs) HasNextRecordsUrl() bool`

HasNextRecordsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


