/*
Salesforce Platform Auth APIs

Testing OAuthApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package auth

import (
	"context"
	openapiclient "github.com/russman12/go-force/pkg/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_auth_OAuthApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OAuthApiService OAuthUserPass", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OAuthApi.OAuthUserPass(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
