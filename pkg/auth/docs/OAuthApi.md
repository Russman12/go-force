# \OAuthApi

All URIs are relative to *https://test.salesforce.com/services*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OAuthUserPass**](OAuthApi.md#OAuthUserPass) | **Post** /oauth2/token | Get OAuth2 token



## OAuthUserPass

> AuthResponse OAuthUserPass(ctx).ClientId(clientId).ClientSecret(clientSecret).Username(username).Password(password).GrantType(grantType).Execute()

Get OAuth2 token

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientId := "clientId_example" // string | The consumer key provided in the connected app.
    clientSecret := "clientSecret_example" // string | The consumer secret provided in the connected app.
    username := "username_example" // string | Your Salesforce username.
    password := "password_example" // string | Your Salesforce password.
    grantType := "grantType_example" // string | Sets the value of this parameter to password. (optional) (default to "password")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthApi.OAuthUserPass(context.Background()).ClientId(clientId).ClientSecret(clientSecret).Username(username).Password(password).GrantType(grantType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.OAuthUserPass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OAuthUserPass`: AuthResponse
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.OAuthUserPass`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOAuthUserPassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | The consumer key provided in the connected app. | 
 **clientSecret** | **string** | The consumer secret provided in the connected app. | 
 **username** | **string** | Your Salesforce username. | 
 **password** | **string** | Your Salesforce password. | 
 **grantType** | **string** | Sets the value of this parameter to password. | [default to &quot;password&quot;]

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

