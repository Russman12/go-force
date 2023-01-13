# CreateJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentRuleId** | Pointer to **string** | The ID of the assignment rule. This property is only shown if an assignment rule is specified when the job is created. | [optional] 
**ColumnDelimiter** | Pointer to [**ColumnDelimiter**](ColumnDelimiter.md) |  | [optional] 
**ContentType** | Pointer to [**ContentType**](ContentType.md) |  | [optional] 
**ExternalIdFieldName** | Pointer to **string** | The name of the external ID field for an upsert. | [optional] 
**LineEnding** | Pointer to [**LineEnding**](LineEnding.md) |  | [optional] 
**Object** | **string** | The object type for the data being processed. | 
**Operation** | [**JobOperation**](JobOperation.md) |  | 

## Methods

### NewCreateJobRequest

`func NewCreateJobRequest(object string, operation JobOperation, ) *CreateJobRequest`

NewCreateJobRequest instantiates a new CreateJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateJobRequestWithDefaults

`func NewCreateJobRequestWithDefaults() *CreateJobRequest`

NewCreateJobRequestWithDefaults instantiates a new CreateJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentRuleId

`func (o *CreateJobRequest) GetAssignmentRuleId() string`

GetAssignmentRuleId returns the AssignmentRuleId field if non-nil, zero value otherwise.

### GetAssignmentRuleIdOk

`func (o *CreateJobRequest) GetAssignmentRuleIdOk() (*string, bool)`

GetAssignmentRuleIdOk returns a tuple with the AssignmentRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentRuleId

`func (o *CreateJobRequest) SetAssignmentRuleId(v string)`

SetAssignmentRuleId sets AssignmentRuleId field to given value.

### HasAssignmentRuleId

`func (o *CreateJobRequest) HasAssignmentRuleId() bool`

HasAssignmentRuleId returns a boolean if a field has been set.

### GetColumnDelimiter

`func (o *CreateJobRequest) GetColumnDelimiter() ColumnDelimiter`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *CreateJobRequest) GetColumnDelimiterOk() (*ColumnDelimiter, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *CreateJobRequest) SetColumnDelimiter(v ColumnDelimiter)`

SetColumnDelimiter sets ColumnDelimiter field to given value.

### HasColumnDelimiter

`func (o *CreateJobRequest) HasColumnDelimiter() bool`

HasColumnDelimiter returns a boolean if a field has been set.

### GetContentType

`func (o *CreateJobRequest) GetContentType() ContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CreateJobRequest) GetContentTypeOk() (*ContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CreateJobRequest) SetContentType(v ContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CreateJobRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetExternalIdFieldName

`func (o *CreateJobRequest) GetExternalIdFieldName() string`

GetExternalIdFieldName returns the ExternalIdFieldName field if non-nil, zero value otherwise.

### GetExternalIdFieldNameOk

`func (o *CreateJobRequest) GetExternalIdFieldNameOk() (*string, bool)`

GetExternalIdFieldNameOk returns a tuple with the ExternalIdFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdFieldName

`func (o *CreateJobRequest) SetExternalIdFieldName(v string)`

SetExternalIdFieldName sets ExternalIdFieldName field to given value.

### HasExternalIdFieldName

`func (o *CreateJobRequest) HasExternalIdFieldName() bool`

HasExternalIdFieldName returns a boolean if a field has been set.

### GetLineEnding

`func (o *CreateJobRequest) GetLineEnding() LineEnding`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *CreateJobRequest) GetLineEndingOk() (*LineEnding, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *CreateJobRequest) SetLineEnding(v LineEnding)`

SetLineEnding sets LineEnding field to given value.

### HasLineEnding

`func (o *CreateJobRequest) HasLineEnding() bool`

HasLineEnding returns a boolean if a field has been set.

### GetObject

`func (o *CreateJobRequest) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreateJobRequest) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreateJobRequest) SetObject(v string)`

SetObject sets Object field to given value.


### GetOperation

`func (o *CreateJobRequest) GetOperation() JobOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CreateJobRequest) GetOperationOk() (*JobOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CreateJobRequest) SetOperation(v JobOperation)`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


