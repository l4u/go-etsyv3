/*
Etsy Open API v3

Testing ShopReturnPolicyApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ShopReturnPolicyApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ShopReturnPolicyApiService ConsolidateShopReturnPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shopId int32

		resp, httpRes, err := apiClient.ShopReturnPolicyApi.ConsolidateShopReturnPolicies(context.Background(), shopId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopReturnPolicyApiService CreateShopReturnPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shopId int32

		resp, httpRes, err := apiClient.ShopReturnPolicyApi.CreateShopReturnPolicy(context.Background(), shopId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopReturnPolicyApiService DeleteShopReturnPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shopId int32
		var returnPolicyId int32

		httpRes, err := apiClient.ShopReturnPolicyApi.DeleteShopReturnPolicy(context.Background(), shopId, returnPolicyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopReturnPolicyApiService GetShopReturnPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shopId int32

		resp, httpRes, err := apiClient.ShopReturnPolicyApi.GetShopReturnPolicies(context.Background(), shopId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopReturnPolicyApiService GetShopReturnPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shopId int32
		var returnPolicyId int32

		resp, httpRes, err := apiClient.ShopReturnPolicyApi.GetShopReturnPolicy(context.Background(), shopId, returnPolicyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopReturnPolicyApiService UpdateShopReturnPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shopId int32
		var returnPolicyId int32

		resp, httpRes, err := apiClient.ShopReturnPolicyApi.UpdateShopReturnPolicy(context.Background(), shopId, returnPolicyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}