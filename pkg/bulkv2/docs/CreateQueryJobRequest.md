# CreateQueryJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | [**QueryOperation**](QueryOperation.md) |  | 
**ContentType** | Pointer to [**ContentType**](ContentType.md) |  | [optional] 
**ColumnDelimiter** | Pointer to [**ColumnDelimiter**](ColumnDelimiter.md) |  | [optional] 
**LineEnding** | Pointer to [**LineEnding**](LineEnding.md) |  | [optional] 
**Query** | **string** | The query to be performed. | 

## Methods

### NewCreateQueryJobRequest

`func NewCreateQueryJobRequest(operation QueryOperation, query string, ) *CreateQueryJobRequest`

NewCreateQueryJobRequest instantiates a new CreateQueryJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQueryJobRequestWithDefaults

`func NewCreateQueryJobRequestWithDefaults() *CreateQueryJobRequest`

NewCreateQueryJobRequestWithDefaults instantiates a new CreateQueryJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *CreateQueryJobRequest) GetOperation() QueryOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CreateQueryJobRequest) GetOperationOk() (*QueryOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CreateQueryJobRequest) SetOperation(v QueryOperation)`

SetOperation sets Operation field to given value.


### GetContentType

`func (o *CreateQueryJobRequest) GetContentType() ContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CreateQueryJobRequest) GetContentTypeOk() (*ContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CreateQueryJobRequest) SetContentType(v ContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CreateQueryJobRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetColumnDelimiter

`func (o *CreateQueryJobRequest) GetColumnDelimiter() ColumnDelimiter`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *CreateQueryJobRequest) GetColumnDelimiterOk() (*ColumnDelimiter, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *CreateQueryJobRequest) SetColumnDelimiter(v ColumnDelimiter)`

SetColumnDelimiter sets ColumnDelimiter field to given value.

### HasColumnDelimiter

`func (o *CreateQueryJobRequest) HasColumnDelimiter() bool`

HasColumnDelimiter returns a boolean if a field has been set.

### GetLineEnding

`func (o *CreateQueryJobRequest) GetLineEnding() LineEnding`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *CreateQueryJobRequest) GetLineEndingOk() (*LineEnding, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *CreateQueryJobRequest) SetLineEnding(v LineEnding)`

SetLineEnding sets LineEnding field to given value.

### HasLineEnding

`func (o *CreateQueryJobRequest) HasLineEnding() bool`

HasLineEnding returns a boolean if a field has been set.

### GetQuery

`func (o *CreateQueryJobRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CreateQueryJobRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CreateQueryJobRequest) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


