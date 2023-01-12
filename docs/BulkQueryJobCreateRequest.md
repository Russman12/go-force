# BulkQueryJobCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | The type of query. | 
**ContentType** | Pointer to **string** | The format that is used for the results. Currently the only supported value is CSV. | [optional] 
**ColumnDelimiter** | Pointer to **string** | The column delimiter used for CSV job data. The default value is COMMA. | [optional] 
**LineEnding** | Pointer to **string** | The line ending used for CSV job data, marking the end of a data row. The default is LF. | [optional] 
**Query** | **string** | The query to be performed. | 

## Methods

### NewBulkQueryJobCreateRequest

`func NewBulkQueryJobCreateRequest(operation string, query string, ) *BulkQueryJobCreateRequest`

NewBulkQueryJobCreateRequest instantiates a new BulkQueryJobCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkQueryJobCreateRequestWithDefaults

`func NewBulkQueryJobCreateRequestWithDefaults() *BulkQueryJobCreateRequest`

NewBulkQueryJobCreateRequestWithDefaults instantiates a new BulkQueryJobCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *BulkQueryJobCreateRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BulkQueryJobCreateRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BulkQueryJobCreateRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetContentType

`func (o *BulkQueryJobCreateRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BulkQueryJobCreateRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BulkQueryJobCreateRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *BulkQueryJobCreateRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetColumnDelimiter

`func (o *BulkQueryJobCreateRequest) GetColumnDelimiter() string`

GetColumnDelimiter returns the ColumnDelimiter field if non-nil, zero value otherwise.

### GetColumnDelimiterOk

`func (o *BulkQueryJobCreateRequest) GetColumnDelimiterOk() (*string, bool)`

GetColumnDelimiterOk returns a tuple with the ColumnDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDelimiter

`func (o *BulkQueryJobCreateRequest) SetColumnDelimiter(v string)`

SetColumnDelimiter sets ColumnDelimiter field to given value.

### HasColumnDelimiter

`func (o *BulkQueryJobCreateRequest) HasColumnDelimiter() bool`

HasColumnDelimiter returns a boolean if a field has been set.

### GetLineEnding

`func (o *BulkQueryJobCreateRequest) GetLineEnding() string`

GetLineEnding returns the LineEnding field if non-nil, zero value otherwise.

### GetLineEndingOk

`func (o *BulkQueryJobCreateRequest) GetLineEndingOk() (*string, bool)`

GetLineEndingOk returns a tuple with the LineEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnding

`func (o *BulkQueryJobCreateRequest) SetLineEnding(v string)`

SetLineEnding sets LineEnding field to given value.

### HasLineEnding

`func (o *BulkQueryJobCreateRequest) HasLineEnding() bool`

HasLineEnding returns a boolean if a field has been set.

### GetQuery

`func (o *BulkQueryJobCreateRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *BulkQueryJobCreateRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *BulkQueryJobCreateRequest) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

