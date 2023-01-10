# BulkJobDetailsBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentRuleId** | **string** | The ID of the assignment rule. This property is only shown if an assignment rule is specified when the job is created. | 
**ColumnDelimiter** | **string** | The column delimiter used for CSV job data. | 
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
**SystemModstamp** | Pointer to **string** | Date and time in the UTC time zone when the job finished. | [optional] 
**State** | Pointer to **string** | The current state of processing for the job. | [optional] 

## Methods

### NewBulkJobDetailsBase

`func NewBulkJobDetailsBase(assignmentRuleId string, columnDelimiter string, contentType string, externalIdFieldName string, lineEnding string, object string, operation string, apiVersion float32, concurrencyMode string, contentUrl string, createdById string, createdDate string, id string, jobType string, ) *BulkJobDetailsBase`

NewBulkJobDetailsBase instantiates a new BulkJobDetailsBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkJobDetailsBaseWithDefaults

`func NewBulkJobDetailsBaseWithDefaults() *BulkJobDetailsBase`

NewBulkJobDetailsBaseWithDefaults instantiates a new BulkJobDetailsBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentRuleId

`func (o *BulkJobDetailsBase) GetAssignmentRuleId() string`

GetAssignmentRuleId returns the AssignmentRuleId field if non-nil, zero value otherwise.

### GetAssignmentRuleIdOk

`func (o *BulkJobDetailsBase) GetAssignmentRuleIdOk() (*string, bool)`

GetAssignmentRuleIdOk returns a tuple with the AssignmentRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentRuleId

`func (o *BulkJobDetailsBase) SetAssignmentRuleId(v string)`

SetAssignmentRuleId sets AssignmentRuleId field to given value.


### GetColumnDelimiter

`func (o *BulkJobDetailsBase) GetColumnDelimiter() string`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *BulkJobDetailsBase) GetColumnDelimiterOk() (*string, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *BulkJobDetailsBase) SetColumnDelimiter(v string)`

SetColumnDelimiter sets ColumnDelimiter field to given value.


### GetContentType

`func (o *BulkJobDetailsBase) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BulkJobDetailsBase) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BulkJobDetailsBase) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetExternalIdFieldName

`func (o *BulkJobDetailsBase) GetExternalIdFieldName() string`

GetExternalIdFieldName returns the ExternalIdFieldName field if non-nil, zero value otherwise.

### GetExternalIdFieldNameOk

`func (o *BulkJobDetailsBase) GetExternalIdFieldNameOk() (*string, bool)`

GetExternalIdFieldNameOk returns a tuple with the ExternalIdFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdFieldName

`func (o *BulkJobDetailsBase) SetExternalIdFieldName(v string)`

SetExternalIdFieldName sets ExternalIdFieldName field to given value.


### GetLineEnding

`func (o *BulkJobDetailsBase) GetLineEnding() string`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *BulkJobDetailsBase) GetLineEndingOk() (*string, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *BulkJobDetailsBase) SetLineEnding(v string)`

SetLineEnding sets LineEnding field to given value.


### GetObject

`func (o *BulkJobDetailsBase) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BulkJobDetailsBase) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BulkJobDetailsBase) SetObject(v string)`

SetObject sets Object field to given value.


### GetOperation

`func (o *BulkJobDetailsBase) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BulkJobDetailsBase) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BulkJobDetailsBase) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetApiVersion

`func (o *BulkJobDetailsBase) GetApiVersion() float32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BulkJobDetailsBase) GetApiVersionOk() (*float32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BulkJobDetailsBase) SetApiVersion(v float32)`

SetApiVersion sets ApiVersion field to given value.


### GetConcurrencyMode

`func (o *BulkJobDetailsBase) GetConcurrencyMode() string`

GetConcurrencyMode returns the ConcurrencyMode field if non-nil, zero value otherwise.

### GetConcurrencyModeOk

`func (o *BulkJobDetailsBase) GetConcurrencyModeOk() (*string, bool)`

GetConcurrencyModeOk returns a tuple with the ConcurrencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyMode

`func (o *BulkJobDetailsBase) SetConcurrencyMode(v string)`

SetConcurrencyMode sets ConcurrencyMode field to given value.


### GetContentUrl

`func (o *BulkJobDetailsBase) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *BulkJobDetailsBase) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *BulkJobDetailsBase) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.


### GetCreatedById

`func (o *BulkJobDetailsBase) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *BulkJobDetailsBase) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *BulkJobDetailsBase) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetCreatedDate

`func (o *BulkJobDetailsBase) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *BulkJobDetailsBase) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *BulkJobDetailsBase) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetId

`func (o *BulkJobDetailsBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkJobDetailsBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkJobDetailsBase) SetId(v string)`

SetId sets Id field to given value.


### GetJobType

`func (o *BulkJobDetailsBase) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BulkJobDetailsBase) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BulkJobDetailsBase) SetJobType(v string)`

SetJobType sets JobType field to given value.


### GetSystemModstamp

`func (o *BulkJobDetailsBase) GetSystemModstamp() string`

GetSystemModstamp returns the SystemModstamp field if non-nil, zero value otherwise.

### GetSystemModstampOk

`func (o *BulkJobDetailsBase) GetSystemModstampOk() (*string, bool)`

GetSystemModstampOk returns a tuple with the SystemModstamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModstamp

`func (o *BulkJobDetailsBase) SetSystemModstamp(v string)`

SetSystemModstamp sets SystemModstamp field to given value.

### HasSystemModstamp

`func (o *BulkJobDetailsBase) HasSystemModstamp() bool`

HasSystemModstamp returns a boolean if a field has been set.

### GetState

`func (o *BulkJobDetailsBase) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BulkJobDetailsBase) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BulkJobDetailsBase) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BulkJobDetailsBase) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


