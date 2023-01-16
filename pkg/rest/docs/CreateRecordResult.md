# CreateRecordResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Errors** | Pointer to **[]interface{}** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateRecordResult

`func NewCreateRecordResult() *CreateRecordResult`

NewCreateRecordResult instantiates a new CreateRecordResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRecordResultWithDefaults

`func NewCreateRecordResultWithDefaults() *CreateRecordResult`

NewCreateRecordResultWithDefaults instantiates a new CreateRecordResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateRecordResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateRecordResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateRecordResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateRecordResult) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreateRecordResult) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreateRecordResult) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetErrors

`func (o *CreateRecordResult) GetErrors() []interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateRecordResult) GetErrorsOk() (*[]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateRecordResult) SetErrors(v []interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateRecordResult) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccess

`func (o *CreateRecordResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateRecordResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateRecordResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateRecordResult) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


