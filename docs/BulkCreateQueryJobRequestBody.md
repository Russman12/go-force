# BulkCreateQueryJobRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** |  | 
**Query** | **string** |  | 
**ContentType** | Pointer to **string** |  | [optional] 
**ColumnDelimiter** | Pointer to **string** |  | [optional] 
**LineEnding** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkCreateQueryJobRequestBody

`func NewBulkCreateQueryJobRequestBody(operation string, query string, ) *BulkCreateQueryJobRequestBody`

NewBulkCreateQueryJobRequestBody instantiates a new BulkCreateQueryJobRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCreateQueryJobRequestBodyWithDefaults

`func NewBulkCreateQueryJobRequestBodyWithDefaults() *BulkCreateQueryJobRequestBody`

NewBulkCreateQueryJobRequestBodyWithDefaults instantiates a new BulkCreateQueryJobRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *BulkCreateQueryJobRequestBody) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BulkCreateQueryJobRequestBody) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BulkCreateQueryJobRequestBody) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetQuery

`func (o *BulkCreateQueryJobRequestBody) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *BulkCreateQueryJobRequestBody) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *BulkCreateQueryJobRequestBody) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetContentType

`func (o *BulkCreateQueryJobRequestBody) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BulkCreateQueryJobRequestBody) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BulkCreateQueryJobRequestBody) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *BulkCreateQueryJobRequestBody) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetColumnDelimiter

`func (o *BulkCreateQueryJobRequestBody) GetColumnDelimiter() string`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *BulkCreateQueryJobRequestBody) GetColumnDelimiterOk() (*string, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *BulkCreateQueryJobRequestBody) SetColumnDelimiter(v string)`

SetColumnDelimiter sets ColumnDelimiter field to given value.

### HasColumnDelimiter

`func (o *BulkCreateQueryJobRequestBody) HasColumnDelimiter() bool`

HasColumnDelimiter returns a boolean if a field has been set.

### GetLineEnding

`func (o *BulkCreateQueryJobRequestBody) GetLineEnding() string`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *BulkCreateQueryJobRequestBody) GetLineEndingOk() (*string, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *BulkCreateQueryJobRequestBody) SetLineEnding(v string)`

SetLineEnding sets LineEnding field to given value.

### HasLineEnding

`func (o *BulkCreateQueryJobRequestBody) HasLineEnding() bool`

HasLineEnding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


