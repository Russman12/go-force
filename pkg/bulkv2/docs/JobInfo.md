# JobInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentRuleId** | **string** | The ID of the assignment rule. This property is only shown if an assignment rule is specified when the job is created. | 
**ColumnDelimiter** | **NullableString** | The column delimiter used for CSV job data. | 
**ContentType** | **string** | The format of the data being processed. Only CSV is supported. | 
**ExternalIdFieldName** | **string** | The name of the external ID field for an upsert. | 
**LineEnding** | **string** | The line ending used for CSV job data. | 
**Object** | **string** | The object type for the data being processed. | 
**Operation** | **string** | The processing operation for the job. | 
**ApiVersion** | **float32** | The API version that the job was created in. | 
**ConcurrencyMode** | **string** | For future use. How the request was processed. Currently only parallel mode is supported. (When other modes are added, the mode will be chosen automatically by the API and will not be user configurable.) | 
**ContentUrl** | **string** | The URL to use for Upload Job Data requests for this job. Only valid if the job is in Open state. | 
**CreatedById** | **string** | The ID of the user who created the job. | 
**CreatedDate** | **string** | The date and time in the UTC time zone when the job was created. | 
**Id** | **string** | Unique ID for this job. | 
**JobType** | **string** | The job’s type. | 
**SystemModstamp** | **string** | Date and time in the UTC time zone when the job finished. | 
**State** | **string** | The current state of processing for the job. | 
**ApexProcessingTime** | Pointer to **int64** | The number of milliseconds taken to process triggers and other processes related to the job data. This doesn&#39;t include the time used for processing asynchronous and batch Apex operations. If there are no triggers, the value is 0. | [optional] 
**ApiActiveProcessingTime** | Pointer to **int64** | The number of milliseconds taken to actively process the job and includes apexProcessingTime, but doesn&#39;t include the time the job waited in the queue to be processed or the time required for serialization and deserialization. | [optional] 
**ErrorMessage** | Pointer to **string** | The error message shown for jobs with errors. | [optional] 
**NumberRecordsFailed** | Pointer to **int64** | The number of records that were not processed successfully in this job. | [optional] 
**NumberRecordsProcessed** | Pointer to **int64** | The number of records already processed. | [optional] 
**Retries** | Pointer to **int32** | The number of times that Salesforce attempted to save the results of an operation. The repeated attempts are due to a problem, such as a lock contention. | [optional] 
**TotalProcessingTime** | Pointer to **int64** | The number of milliseconds taken to process the job. | [optional] 

## Methods

### NewJobInfo

`func NewJobInfo(assignmentRuleId string, columnDelimiter NullableString, contentType string, externalIdFieldName string, lineEnding string, object string, operation string, apiVersion float32, concurrencyMode string, contentUrl string, createdById string, createdDate string, id string, jobType string, systemModstamp string, state string, ) *JobInfo`

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


### GetColumnDelimiter

`func (o *JobInfo) GetColumnDelimiter() string`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *JobInfo) GetColumnDelimiterOk() (*string, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *JobInfo) SetColumnDelimiter(v string)`

SetColumnDelimiter sets ColumnDelimiter field to given value.


### SetColumnDelimiterNil

`func (o *JobInfo) SetColumnDelimiterNil(b bool)`

 SetColumnDelimiterNil sets the value for ColumnDelimiter to be an explicit nil

### UnsetColumnDelimiter
`func (o *JobInfo) UnsetColumnDelimiter()`

UnsetColumnDelimiter ensures that no value is present for ColumnDelimiter, not even an explicit nil
### GetContentType

`func (o *JobInfo) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *JobInfo) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *JobInfo) SetContentType(v string)`

SetContentType sets ContentType field to given value.


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


### GetLineEnding

`func (o *JobInfo) GetLineEnding() string`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *JobInfo) GetLineEndingOk() (*string, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *JobInfo) SetLineEnding(v string)`

SetLineEnding sets LineEnding field to given value.


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


### GetOperation

`func (o *JobInfo) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *JobInfo) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *JobInfo) SetOperation(v string)`

SetOperation sets Operation field to given value.


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


### GetJobType

`func (o *JobInfo) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *JobInfo) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *JobInfo) SetJobType(v string)`

SetJobType sets JobType field to given value.


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


### GetState

`func (o *JobInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JobInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JobInfo) SetState(v string)`

SetState sets State field to given value.


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


