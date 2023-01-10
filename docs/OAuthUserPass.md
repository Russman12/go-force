# OAuthUserPass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | Pointer to **string** | Sets the value of this parameter to password. | [optional] [default to "password"]
**ClientId** | **string** | The consumer key provided in the connected app. | 
**ClientSecret** | **string** | The consumer secret provided in the connected app. | 
**Username** | **string** | Your Salesforce username. | 
**Password** | **string** | Your Salesforce password. | 

## Methods

### NewOAuthUserPass

`func NewOAuthUserPass(clientId string, clientSecret string, username string, password string, ) *OAuthUserPass`

NewOAuthUserPass instantiates a new OAuthUserPass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthUserPassWithDefaults

`func NewOAuthUserPassWithDefaults() *OAuthUserPass`

NewOAuthUserPassWithDefaults instantiates a new OAuthUserPass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *OAuthUserPass) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *OAuthUserPass) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *OAuthUserPass) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *OAuthUserPass) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetClientId

`func (o *OAuthUserPass) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthUserPass) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthUserPass) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OAuthUserPass) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthUserPass) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthUserPass) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetUsername

`func (o *OAuthUserPass) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OAuthUserPass) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OAuthUserPass) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *OAuthUserPass) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OAuthUserPass) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OAuthUserPass) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


