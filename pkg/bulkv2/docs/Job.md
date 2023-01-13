# Job

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

## Methods

### NewJob

`func NewJob() *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentRuleId

`func (o *Job) GetAssignmentRuleId() string`

GetAssignmentRuleId returns the AssignmentRuleId field if non-nil, zero value otherwise.

### GetAssignmentRuleIdOk

`func (o *Job) GetAssignmentRuleIdOk() (*string, bool)`

GetAssignmentRuleIdOk returns a tuple with the AssignmentRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentRuleId

`func (o *Job) SetAssignmentRuleId(v string)`

SetAssignmentRuleId sets AssignmentRuleId field to given value.

### HasAssignmentRuleId

`func (o *Job) HasAssignmentRuleId() bool`

HasAssignmentRuleId returns a boolean if a field has been set.

### GetColumnDelimiter

`func (o *Job) GetColumnDelimiter() ColumnDelimiter`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *Job) GetColumnDelimiterOk() (*ColumnDelimiter, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *Job) SetColumnDelimiter(v ColumnDelimiter)`

SetColumnDelimiter sets ColumnDelimiter field to given value.

### HasColumnDelimiter

`func (o *Job) HasColumnDelimiter() bool`

HasColumnDelimiter returns a boolean if a field has been set.

### GetContentType

`func (o *Job) GetContentType() ContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Job) GetContentTypeOk() (*ContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Job) SetContentType(v ContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *Job) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetExternalIdFieldName

`func (o *Job) GetExternalIdFieldName() string`

GetExternalIdFieldName returns the ExternalIdFieldName field if non-nil, zero value otherwise.

### GetExternalIdFieldNameOk

`func (o *Job) GetExternalIdFieldNameOk() (*string, bool)`

GetExternalIdFieldNameOk returns a tuple with the ExternalIdFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdFieldName

`func (o *Job) SetExternalIdFieldName(v string)`

SetExternalIdFieldName sets ExternalIdFieldName field to given value.

### HasExternalIdFieldName

`func (o *Job) HasExternalIdFieldName() bool`

HasExternalIdFieldName returns a boolean if a field has been set.

### GetLineEnding

`func (o *Job) GetLineEnding() LineEnding`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *Job) GetLineEndingOk() (*LineEnding, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *Job) SetLineEnding(v LineEnding)`

SetLineEnding sets LineEnding field to given value.

### HasLineEnding

`func (o *Job) HasLineEnding() bool`

HasLineEnding returns a boolean if a field has been set.

### GetObject

`func (o *Job) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Job) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Job) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Job) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOperation

`func (o *Job) GetOperation() JobOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Job) GetOperationOk() (*JobOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Job) SetOperation(v JobOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *Job) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


