/*
Etsy Open API v3

Testing ShopProductionPartnerApiService

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

func Test_openapi_ShopProductionPartnerApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ShopProductionPartnerApiService GetShopProductionPartners", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shopId int32

		resp, httpRes, err := apiClient.ShopProductionPartnerApi.GetShopProductionPartners(context.Background(), shopId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
