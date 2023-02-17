/*
Salesforce Platform REST API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tooling

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"

    openapiclient "github.com/russman12/go-force/pkg/tooling"
    "github.com/russman12/go-force/mocks"
)

func Test_tooling_DefaultApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration, mocks.NewTokenSource(t))

    t.Run("Test DefaultApiService CreateRecord", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sObjectName string

        resp, httpRes, err := apiClient.DefaultApi.CreateRecord(context.Background(), sObjectName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService DeleteRecord", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sObjectName string
        var id string

        httpRes, err := apiClient.DefaultApi.DeleteRecord(context.Background(), sObjectName, id).Execute()

        require.Nil(t, err)
        
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService DescribeSObject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sObjectName string

        resp, httpRes, err := apiClient.DefaultApi.DescribeSObject(context.Background(), sObjectName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ExecuteAnonymous", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.ExecuteAnonymous(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetCompletions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.GetCompletions(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetSObject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sObjectName string

        resp, httpRes, err := apiClient.DefaultApi.GetSObject(context.Background(), sObjectName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetSObjects", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.GetSObjects(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService Query", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.Query(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService RetrieveRecord", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sObjectName string
        var id string

        resp, httpRes, err := apiClient.DefaultApi.RetrieveRecord(context.Background(), sObjectName, id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService RunTestsAsync", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.RunTestsAsync(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService RunTestsSync", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.RunTestsSync(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService Search", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.Search(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService UpdateRecord", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sObjectName string
        var id string

        httpRes, err := apiClient.DefaultApi.UpdateRecord(context.Background(), sObjectName, id).Execute()

        require.Nil(t, err)
        
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}