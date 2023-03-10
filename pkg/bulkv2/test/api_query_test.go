/*
Salesforce Platform Bulk V2 API

Testing QueryApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package bulkv2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/russman12/go-force/mocks"
	openapiclient "github.com/russman12/go-force/pkg/bulkv2"
)

func Test_bulkv2_QueryApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration, mocks.NewTokenSource(t))

	t.Run("Test QueryApiService AbortQueryJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.QueryApi.AbortQueryJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QueryApiService CreateQueryJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.QueryApi.CreateQueryJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QueryApiService DeleteQueryJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		httpRes, err := apiClient.QueryApi.DeleteQueryJob(context.Background(), jobId).Execute()

		require.Nil(t, err)

		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QueryApiService GetQueryJobInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.QueryApi.GetQueryJobInfo(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QueryApiService GetQueryJobResults", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.QueryApi.GetQueryJobResults(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QueryApiService GetQueryJobs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.QueryApi.GetQueryJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
