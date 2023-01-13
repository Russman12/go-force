# QueryJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to [**QueryOperation**](QueryOperation.md) |  | [optional] 
**ContentType** | Pointer to [**ContentType**](ContentType.md) |  | [optional] 
**ColumnDelimiter** | Pointer to [**ColumnDelimiter**](ColumnDelimiter.md) |  | [optional] 
**LineEnding** | Pointer to [**LineEnding**](LineEnding.md) |  | [optional] 

## Methods

### NewQueryJob

`func NewQueryJob() *QueryJob`

NewQueryJob instantiates a new QueryJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryJobWithDefaults

`func NewQueryJobWithDefaults() *QueryJob`

NewQueryJobWithDefaults instantiates a new QueryJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *QueryJob) GetOperation() QueryOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *QueryJob) GetOperationOk() (*QueryOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *QueryJob) SetOperation(v QueryOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *QueryJob) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetContentType

`func (o *QueryJob) GetContentType() ContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *QueryJob) GetContentTypeOk() (*ContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *QueryJob) SetContentType(v ContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *QueryJob) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetColumnDelimiter

`func (o *QueryJob) GetColumnDelimiter() ColumnDelimiter`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *QueryJob) GetColumnDelimiterOk() (*ColumnDelimiter, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *QueryJob) SetColumnDelimiter(v ColumnDelimiter)`

SetColumnDelimiter sets ColumnDelimiter field to given value.

### HasColumnDelimiter

`func (o *QueryJob) HasColumnDelimiter() bool`

HasColumnDelimiter returns a boolean if a field has been set.

### GetLineEnding

`func (o *QueryJob) GetLineEnding() LineEnding`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *QueryJob) GetLineEndingOk() (*LineEnding, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *QueryJob) SetLineEnding(v LineEnding)`

SetLineEnding sets LineEnding field to given value.

### HasLineEnding

`func (o *QueryJob) HasLineEnding() bool`

HasLineEnding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


