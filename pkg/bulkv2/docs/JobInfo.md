# JobInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentRuleId** | Pointer to **string** | The ID of the assignment rule. This property is only shown if an assignment rule is specified when the job is created. | [optional] 
**ColumnDelimiter** | Pointer to [**ColumnDelimiter**](ColumnDelimiter.md) |  | [optional] 
**ContentType** | Pointer to [**ContentType**](ContentType.md) |  | [optional] 
**ExternalIdFieldName** | Pointer to **string** | The name of the external ID field for an upsert. | [optional] 
**LineEnding** | Pointer to [**LineEnding**](LineEnding.md) |  | [optional] 
**Object** | Pointer to **string** | The object type for the data being processed. | [optional] 
**Operation** | Pointer to [**JobOperation**](JobOperation.md) |  | [optional] 
**ApiVersion** | Pointer to **float32** | The API version that the job was created in. | [optional] 
**ConcurrencyMode** | Pointer to **string** | For future use. How the request was processed. Currently only parallel mode is supported. (When other modes are added, the mode will be chosen automatically by the API and will not be user configurable.) | [optional] 
**ContentUrl** | Pointer to **string** | The URL to use for Upload Job Data requests for this job. Only valid if the job is in Open state. | [optional] 
**CreatedById** | Pointer to **string** | The ID of the user who created the job. | [optional] 
**CreatedDate** | Pointer to **string** | The date and time in the UTC time zone when the job was created. | [optional] 
**Id** | Pointer to **string** | Unique ID for this job. | [optional] 
**JobType** | Pointer to [**JobType**](JobType.md) |  | [optional] 
**SystemModstamp** | Pointer to **string** | Date and time in the UTC time zone when the job finished. | [optional] 
**State** | Pointer to [**JobState**](JobState.md) |  | [optional] 
**ApexProcessingTime** | Pointer to **int64** | The number of milliseconds taken to process triggers and other processes related to the job data. This doesn&#39;t include the time used for processing asynchronous and batch Apex operations. If there are no triggers, the value is 0. | [optional] 
**ApiActiveProcessingTime** | Pointer to **int64** | The number of milliseconds taken to actively process the job and includes apexProcessingTime, but doesn&#39;t include the time the job waited in the queue to be processed or the time required for serialization and deserialization. | [optional] 
**ErrorMessage** | Pointer to **string** | The error message shown for jobs with errors. | [optional] 
**NumberRecordsFailed** | Pointer to **int64** | The number of records that were not processed successfully in this job. | [optional] 
**NumberRecordsProcessed** | Pointer to **int64** | The number of records already processed. | [optional] 
**Retries** | Pointer to **int32** | The number of times that Salesforce attempted to save the results of an operation. The repeated attempts are due to a problem, such as a lock contention. | [optional] 
**TotalProcessingTime** | Pointer to **int64** | The number of milliseconds taken to process the job. | [optional] 

## Methods

### NewJobInfo

`func NewJobInfo() *JobInfo`

NewJobInfo instantiates a new JobInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobInfoWithDefaults

`func NewJobInfoWithDefaults() *JobInfo`

NewJobInfoWithDefaults instantiates a new JobInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentRuleId

`func (o *JobInfo) GetAssignmentRuleId() string`

GetAssignmentRuleId returns the AssignmentRuleId field if non-nil, zero value otherwise.

### GetAssignmentRuleIdOk

`func (o *JobInfo) GetAssignmentRuleIdOk() (*string, bool)`

GetAssignmentRuleIdOk returns a tuple with the AssignmentRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentRuleId

`func (o *JobInfo) SetAssignmentRuleId(v string)`

SetAssignmentRuleId sets AssignmentRuleId field to given value.

### HasAssignmentRuleId

`func (o *JobInfo) HasAssignmentRuleId() bool`

HasAssignmentRuleId returns a boolean if a field has been set.

### GetColumnDelimiter

`func (o *JobInfo) GetColumnDelimiter() ColumnDelimiter`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *JobInfo) GetColumnDelimiterOk() (*ColumnDelimiter, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *JobInfo) SetColumnDelimiter(v ColumnDelimiter)`

SetColumnDelimiter sets ColumnDelimiter field to given value.

### HasColumnDelimiter

`func (o *JobInfo) HasColumnDelimiter() bool`

HasColumnDelimiter returns a boolean if a field has been set.

### GetContentType

`func (o *JobInfo) GetContentType() ContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *JobInfo) GetContentTypeOk() (*ContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *JobInfo) SetContentType(v ContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *JobInfo) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetExternalIdFieldName

`func (o *JobInfo) GetExternalIdFieldName() string`

GetExternalIdFieldName returns the ExternalIdFieldName field if non-nil, zero value otherwise.

### GetExternalIdFieldNameOk

`func (o *JobInfo) GetExternalIdFieldNameOk() (*string, bool)`

GetExternalIdFieldNameOk returns a tuple with the ExternalIdFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdFieldName

`func (o *JobInfo) SetExternalIdFieldName(v string)`

SetExternalIdFieldName sets ExternalIdFieldName field to given value.

### HasExternalIdFieldName

`func (o *JobInfo) HasExternalIdFieldName() bool`

HasExternalIdFieldName returns a boolean if a field has been set.

### GetLineEnding

`func (o *JobInfo) GetLineEnding() LineEnding`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *JobInfo) GetLineEndingOk() (*LineEnding, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *JobInfo) SetLineEnding(v LineEnding)`

SetLineEnding sets LineEnding field to given value.

### HasLineEnding

`func (o *JobInfo) HasLineEnding() bool`

HasLineEnding returns a boolean if a field has been set.

### GetObject

`func (o *JobInfo) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *JobInfo) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *JobInfo) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *JobInfo) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOperation

`func (o *JobInfo) GetOperation() JobOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *JobInfo) GetOperationOk() (*JobOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *JobInfo) SetOperation(v JobOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *JobInfo) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetApiVersion

`func (o *JobInfo) GetApiVersion() float32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *JobInfo) GetApiVersionOk() (*float32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *JobInfo) SetApiVersion(v float32)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *JobInfo) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetConcurrencyMode

`func (o *JobInfo) GetConcurrencyMode() string`

GetConcurrencyMode returns the ConcurrencyMode field if non-nil, zero value otherwise.

### GetConcurrencyModeOk

`func (o *JobInfo) GetConcurrencyModeOk() (*string, bool)`

GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyMode

`func (o *JobInfo) SetConcurrencyMode(v string)`

SetConcurrencyMode sets ConcurrencyMode field to given value.

### HasConcurrencyMode

`func (o *JobInfo) HasConcurrencyMode() bool`

HasConcurrencyMode returns a boolean if a field has been set.

### GetContentUrl

`func (o *JobInfo) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *JobInfo) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *JobInfo) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.

### HasContentUrl

`func (o *JobInfo) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### GetCreatedById

`func (o *JobInfo) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *JobInfo) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *JobInfo) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *JobInfo) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedDate

`func (o *JobInfo) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *JobInfo) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *JobInfo) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *JobInfo) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetId

`func (o *JobInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobType

`func (o *JobInfo) GetJobType() JobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *JobInfo) GetJobTypeOk() (*JobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *JobInfo) SetJobType(v JobType)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *JobInfo) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetSystemModstamp

`func (o *JobInfo) GetSystemModstamp() string`

GetSystemModstamp returns the SystemModstamp field if non-nil, zero value otherwise.

### GetSystemModstampOk

`func (o *JobInfo) GetSystemModstampOk() (*string, bool)`

GetSystemModstampOk returns a tuple with the SystemModstamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModstamp

`func (o *JobInfo) SetSystemModstamp(v string)`

SetSystemModstamp sets SystemModstamp field to given value.

### HasSystemModstamp

`func (o *JobInfo) HasSystemModstamp() bool`

HasSystemModstamp returns a boolean if a field has been set.

### GetState

`func (o *JobInfo) GetState() JobState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JobInfo) GetStateOk() (*JobState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JobInfo) SetState(v JobState)`

SetState sets State field to given value.

### HasState

`func (o *JobInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApexProcessingTime

`func (o *JobInfo) GetApexProcessingTime() int64`

GetApexProcessingTime returns the ApexProcessingTime field if non-nil, zero value otherwise.

### GetApexProcessingTimeOk

`func (o *JobInfo) GetApexProcessingTimeOk() (*int64, bool)`

GetApexProcessingTimeOk returns a tuple with the ApexProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApexProcessingTime

`func (o *JobInfo) SetApexProcessingTime(v int64)`

SetApexProcessingTime sets ApexProcessingTime field to given value.

### HasApexProcessingTime

`func (o *JobInfo) HasApexProcessingTime() bool`

HasApexProcessingTime returns a boolean if a field has been set.

### GetApiActiveProcessingTime

`func (o *JobInfo) GetApiActiveProcessingTime() int64`

GetApiActiveProcessingTime returns the ApiActiveProcessingTime field if non-nil, zero value otherwise.

### GetApiActiveProcessingTimeOk

`func (o *JobInfo) GetApiActiveProcessingTimeOk() (*int64, bool)`

GetApiActiveProcessingTimeOk returns a tuple with the ApiActiveProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiActiveProcessingTime

`func (o *JobInfo) SetApiActiveProcessingTime(v int64)`

SetApiActiveProcessingTime sets ApiActiveProcessingTime field to given value.

### HasApiActiveProcessingTime

`func (o *JobInfo) HasApiActiveProcessingTime() bool`

HasApiActiveProcessingTime returns a boolean if a field has been set.

### GetErrorMessage

`func (o *JobInfo) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *JobInfo) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *JobInfo) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *JobInfo) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetNumberRecordsFailed

`func (o *JobInfo) GetNumberRecordsFailed() int64`

GetNumberRecordsFailed returns the NumberRecordsFailed field if non-nil, zero value otherwise.

### GetNumberRecordsFailedOk

`func (o *JobInfo) GetNumberRecordsFailedOk() (*int64, bool)`

GetNumberRecordsFailedOk returns a tuple with the NumberRecordsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberRecordsFailed

`func (o *JobInfo) SetNumberRecordsFailed(v int64)`

SetNumberRecordsFailed sets NumberRecordsFailed field to given value.

### HasNumberRecordsFailed

`func (o *JobInfo) HasNumberRecordsFailed() bool`

HasNumberRecordsFailed returns a boolean if a field has been set.

### GetNumberRecordsProcessed

`func (o *JobInfo) GetNumberRecordsProcessed() int64`

GetNumberRecordsProcessed returns the NumberRecordsProcessed field if non-nil, zero value otherwise.

### GetNumberRecordsProcessedOk

`func (o *JobInfo) GetNumberRecordsProcessedOk() (*int64, bool)`

GetNumberRecordsProcessedOk returns a tuple with the NumberRecordsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberRecordsProcessed

`func (o *JobInfo) SetNumberRecordsProcessed(v int64)`

SetNumberRecordsProcessed sets NumberRecordsProcessed field to given value.

### HasNumberRecordsProcessed

`func (o *JobInfo) HasNumberRecordsProcessed() bool`

HasNumberRecordsProcessed returns a boolean if a field has been set.

### GetRetries

`func (o *JobInfo) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *JobInfo) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *JobInfo) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *JobInfo) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetTotalProcessingTime

`func (o *JobInfo) GetTotalProcessingTime() int64`

GetTotalProcessingTime returns the TotalProcessingTime field if non-nil, zero value otherwise.

### GetTotalProcessingTimeOk

`func (o *JobInfo) GetTotalProcessingTimeOk() (*int64, bool)`

GetTotalProcessingTimeOk returns a tuple with the TotalProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProcessingTime

`func (o *JobInfo) SetTotalProcessingTime(v int64)`

SetTotalProcessingTime sets TotalProcessingTime field to given value.

### HasTotalProcessingTime

`func (o *JobInfo) HasTotalProcessingTime() bool`

HasTotalProcessingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


