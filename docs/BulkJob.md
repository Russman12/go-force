# BulkJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentRuleId** | Pointer to **string** | The ID of the assignment rule. This property is only shown if an assignment rule is specified when the job is created. | [optional] 
**ColumnDelimiter** | Pointer to **string** | The column delimiter used for CSV job data. | [optional] 
**ContentType** | Pointer to **string** | The format of the data being processed. Only CSV is supported. | [optional] 
**ExternalIdFieldName** | Pointer to **string** | The name of the external ID field for an upsert. | [optional] 
**LineEnding** | Pointer to **string** | The line ending used for CSV job data. | [optional] 
**Object** | **string** | The object type for the data being processed. | 
**Operation** | **string** | The processing operation for the job. | 

## Methods

### NewBulkJob

`func NewBulkJob(object string, operation string, ) *BulkJob`

NewBulkJob instantiates a new BulkJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkJobWithDefaults

`func NewBulkJobWithDefaults() *BulkJob`

NewBulkJobWithDefaults instantiates a new BulkJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentRuleId

`func (o *BulkJob) GetAssignmentRuleId() string`

GetAssignmentRuleId returns the AssignmentRuleId field if non-nil, zero value otherwise.

### GetAssignmentRuleIdOk

`func (o *BulkJob) GetAssignmentRuleIdOk() (*string, bool)`

GetAssignmentRuleIdOk returns a tuple with the AssignmentRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentRuleId

`func (o *BulkJob) SetAssignmentRuleId(v string)`

SetAssignmentRuleId sets AssignmentRuleId field to given value.

### HasAssignmentRuleId

`func (o *BulkJob) HasAssignmentRuleId() bool`

HasAssignmentRuleId returns a boolean if a field has been set.

### GetColumnDelimiter

`func (o *BulkJob) GetColumnDelimiter() string`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *BulkJob) GetColumnDelimiterOk() (*string, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *BulkJob) SetColumnDelimiter(v string)`

SetColumnDelimiter sets ColumnDelimiter field to given value.

### HasColumnDelimiter

`func (o *BulkJob) HasColumnDelimiter() bool`

HasColumnDelimiter returns a boolean if a field has been set.

### GetContentType

`func (o *BulkJob) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BulkJob) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BulkJob) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *BulkJob) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetExternalIdFieldName

`func (o *BulkJob) GetExternalIdFieldName() string`

GetExternalIdFieldName returns the ExternalIdFieldName field if non-nil, zero value otherwise.

### GetExternalIdFieldNameOk

`func (o *BulkJob) GetExternalIdFieldNameOk() (*string, bool)`

GetExternalIdFieldNameOk returns a tuple with the ExternalIdFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdFieldName

`func (o *BulkJob) SetExternalIdFieldName(v string)`

SetExternalIdFieldName sets ExternalIdFieldName field to given value.

### HasExternalIdFieldName

`func (o *BulkJob) HasExternalIdFieldName() bool`

HasExternalIdFieldName returns a boolean if a field has been set.

### GetLineEnding

`func (o *BulkJob) GetLineEnding() string`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *BulkJob) GetLineEndingOk() (*string, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *BulkJob) SetLineEnding(v string)`

SetLineEnding sets LineEnding field to given value.

### HasLineEnding

`func (o *BulkJob) HasLineEnding() bool`

HasLineEnding returns a boolean if a field has been set.

### GetObject

`func (o *BulkJob) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BulkJob) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BulkJob) SetObject(v string)`

SetObject sets Object field to given value.


### GetOperation

`func (o *BulkJob) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BulkJob) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BulkJob) SetOperation(v string)`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


