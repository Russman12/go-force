/*
Salesforce Platform Auth APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 56.0
Contact: russell-laboe@outlook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auth

import (
	"encoding/json"
)

// AuthResponse struct for AuthResponse
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	InstanceUrl string `json:"instance_url"`
	Id          string `json:"id"`
	TokenType   string `json:"token_type"`
	IssuedAt    string `json:"issued_at"`
	Signature   string `json:"signature"`
}

// NewAuthResponse instantiates a new AuthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthResponse(accessToken string, instanceUrl string, id string, tokenType string, issuedAt string, signature string) *AuthResponse {
	this := AuthResponse{}
	this.AccessToken = accessToken
	this.InstanceUrl = instanceUrl
	this.Id = id
	this.TokenType = tokenType
	this.IssuedAt = issuedAt
	this.Signature = signature
	return &this
}

// NewAuthResponseWithDefaults instantiates a new AuthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthResponseWithDefaults() *AuthResponse {
	this := AuthResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *AuthResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *AuthResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *AuthResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetInstanceUrl returns the InstanceUrl field value
func (o *AuthResponse) GetInstanceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceUrl
}

// GetInstanceUrlOk returns a tuple with the InstanceUrl field value
// and a boolean to check if the value has been set.
func (o *AuthResponse) GetInstanceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceUrl, true
}

// SetInstanceUrl sets field value
func (o *AuthResponse) SetInstanceUrl(v string) {
	o.InstanceUrl = v
}

// GetId returns the Id field value
func (o *AuthResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuthResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuthResponse) SetId(v string) {
	o.Id = v
}

// GetTokenType returns the TokenType field value
func (o *AuthResponse) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *AuthResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *AuthResponse) SetTokenType(v string) {
	o.TokenType = v
}

// GetIssuedAt returns the IssuedAt field value
func (o *AuthResponse) GetIssuedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value
// and a boolean to check if the value has been set.
func (o *AuthResponse) GetIssuedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedAt, true
}

// SetIssuedAt sets field value
func (o *AuthResponse) SetIssuedAt(v string) {
	o.IssuedAt = v
}

// GetSignature returns the Signature field value
func (o *AuthResponse) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *AuthResponse) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *AuthResponse) SetSignature(v string) {
	o.Signature = v
}

func (o AuthResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["instance_url"] = o.InstanceUrl
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["token_type"] = o.TokenType
	}
	if true {
		toSerialize["issued_at"] = o.IssuedAt
	}
	if true {
		toSerialize["signature"] = o.Signature
	}
	return json.Marshal(toSerialize)
}

type NullableAuthResponse struct {
	value *AuthResponse
	isSet bool
}

func (v NullableAuthResponse) Get() *AuthResponse {
	return v.value
}

func (v *NullableAuthResponse) Set(val *AuthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthResponse(val *AuthResponse) *NullableAuthResponse {
	return &NullableAuthResponse{value: val, isSet: true}
}

func (v NullableAuthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}