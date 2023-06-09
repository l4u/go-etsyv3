# Go API client for openapi

<div class=\"wt-text-body-01\"><p class=\"wt-pt-xs-2 wt-pb-xs-2\">Etsy's Open API provides a simple RESTful interface for various Etsy.com features. The API endpoints are meant to replace Etsy's Open API v2, which is scheduled to end service in 2022.</p><p class=\"wt-pb-xs-2\">All of the endpoints are callable and the majority of the API endpoints are now in a beta phase. This means we do not expect to make any breaking changes before our general release. A handful of endpoints are currently interface stubs (labeled “Feedback Only”) and returns a \"501 Not Implemented\" response code when called.</p><p class=\"wt-pb-xs-2\">If you'd like to report an issue or provide feedback on the API design, <a target=\"_blank\" class=\"wt-text-link wt-p-xs-0\" href=\"https://github.com/etsy/open-api/discussions\">please add an issue in Github</a>.</p></div>&copy; 2021-2023 Etsy, Inc. All Rights Reserved. Use of this code is subject to Etsy's <a class='wt-text-link wt-p-xs-0' target='_blank' href='https://www.etsy.com/legal/api'>API Developer Terms of Use</a>.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/l4u/go-etsyv3"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://openapi.etsy.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BuyerTaxonomyApi* | [**GetBuyerTaxonomyNodes**](docs/BuyerTaxonomyApi.md#getbuyertaxonomynodes) | **Get** /v3/application/buyer-taxonomy/nodes | 
*BuyerTaxonomyApi* | [**GetPropertiesByBuyerTaxonomyId**](docs/BuyerTaxonomyApi.md#getpropertiesbybuyertaxonomyid) | **Get** /v3/application/buyer-taxonomy/nodes/{taxonomy_id}/properties | 
*LedgerEntryApi* | [**GetShopPaymentAccountLedgerEntries**](docs/LedgerEntryApi.md#getshoppaymentaccountledgerentries) | **Get** /v3/application/shops/{shop_id}/payment-account/ledger-entries | 
*LedgerEntryApi* | [**GetShopPaymentAccountLedgerEntry**](docs/LedgerEntryApi.md#getshoppaymentaccountledgerentry) | **Get** /v3/application/shops/{shop_id}/payment-account/ledger-entries/{ledger_entry_id} | 
*OtherApi* | [**Ping**](docs/OtherApi.md#ping) | **Get** /v3/application/openapi-ping | 
*OtherApi* | [**TokenScopes**](docs/OtherApi.md#tokenscopes) | **Post** /v3/application/scopes | 
*PaymentApi* | [**GetPaymentAccountLedgerEntryPayments**](docs/PaymentApi.md#getpaymentaccountledgerentrypayments) | **Get** /v3/application/shops/{shop_id}/payment-account/ledger-entries/payments | 
*PaymentApi* | [**GetPayments**](docs/PaymentApi.md#getpayments) | **Get** /v3/application/shops/{shop_id}/payments | 
*PaymentApi* | [**GetShopPaymentByReceiptId**](docs/PaymentApi.md#getshoppaymentbyreceiptid) | **Get** /v3/application/shops/{shop_id}/receipts/{receipt_id}/payments | 
*ReviewApi* | [**GetReviewsByListing**](docs/ReviewApi.md#getreviewsbylisting) | **Get** /v3/application/listings/{listing_id}/reviews | 
*ReviewApi* | [**GetReviewsByShop**](docs/ReviewApi.md#getreviewsbyshop) | **Get** /v3/application/shops/{shop_id}/reviews | 
*SellerTaxonomyApi* | [**GetPropertiesByTaxonomyId**](docs/SellerTaxonomyApi.md#getpropertiesbytaxonomyid) | **Get** /v3/application/seller-taxonomy/nodes/{taxonomy_id}/properties | 
*SellerTaxonomyApi* | [**GetSellerTaxonomyNodes**](docs/SellerTaxonomyApi.md#getsellertaxonomynodes) | **Get** /v3/application/seller-taxonomy/nodes | 
*ShopApi* | [**FindShops**](docs/ShopApi.md#findshops) | **Get** /v3/application/shops | 
*ShopApi* | [**GetShop**](docs/ShopApi.md#getshop) | **Get** /v3/application/shops/{shop_id} | 
*ShopApi* | [**GetShopByOwnerUserId**](docs/ShopApi.md#getshopbyowneruserid) | **Get** /v3/application/users/{user_id}/shops | 
*ShopApi* | [**UpdateShop**](docs/ShopApi.md#updateshop) | **Put** /v3/application/shops/{shop_id} | 
*ShopListingApi* | [**CreateDraftListing**](docs/ShopListingApi.md#createdraftlisting) | **Post** /v3/application/shops/{shop_id}/listings | 
*ShopListingApi* | [**DeleteListing**](docs/ShopListingApi.md#deletelisting) | **Delete** /v3/application/listings/{listing_id} | 
*ShopListingApi* | [**DeleteListingProperty**](docs/ShopListingApi.md#deletelistingproperty) | **Delete** /v3/application/shops/{shop_id}/listings/{listing_id}/properties/{property_id} | 
*ShopListingApi* | [**FindAllActiveListingsByShop**](docs/ShopListingApi.md#findallactivelistingsbyshop) | **Get** /v3/application/shops/{shop_id}/listings/active | 
*ShopListingApi* | [**FindAllListingsActive**](docs/ShopListingApi.md#findalllistingsactive) | **Get** /v3/application/listings/active | 
*ShopListingApi* | [**GetFeaturedListingsByShop**](docs/ShopListingApi.md#getfeaturedlistingsbyshop) | **Get** /v3/application/shops/{shop_id}/listings/featured | 
*ShopListingApi* | [**GetListing**](docs/ShopListingApi.md#getlisting) | **Get** /v3/application/listings/{listing_id} | 
*ShopListingApi* | [**GetListingProperties**](docs/ShopListingApi.md#getlistingproperties) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/properties | 
*ShopListingApi* | [**GetListingProperty**](docs/ShopListingApi.md#getlistingproperty) | **Get** /v3/application/listings/{listing_id}/properties/{property_id} | 
*ShopListingApi* | [**GetListingsByListingIds**](docs/ShopListingApi.md#getlistingsbylistingids) | **Get** /v3/application/listings/batch | 
*ShopListingApi* | [**GetListingsByShop**](docs/ShopListingApi.md#getlistingsbyshop) | **Get** /v3/application/shops/{shop_id}/listings | 
*ShopListingApi* | [**GetListingsByShopReceipt**](docs/ShopListingApi.md#getlistingsbyshopreceipt) | **Get** /v3/application/shops/{shop_id}/receipts/{receipt_id}/listings | 
*ShopListingApi* | [**GetListingsByShopReturnPolicy**](docs/ShopListingApi.md#getlistingsbyshopreturnpolicy) | **Get** /v3/application/shops/{shop_id}/policies/return/{return_policy_id}/listings | 
*ShopListingApi* | [**GetListingsByShopSectionId**](docs/ShopListingApi.md#getlistingsbyshopsectionid) | **Get** /v3/application/shops/{shop_id}/shop-sections/listings | 
*ShopListingApi* | [**UpdateListing**](docs/ShopListingApi.md#updatelisting) | **Patch** /v3/application/shops/{shop_id}/listings/{listing_id} | 
*ShopListingApi* | [**UpdateListingDeprecated**](docs/ShopListingApi.md#updatelistingdeprecated) | **Put** /v3/application/shops/{shop_id}/listings/{listing_id} | 
*ShopListingApi* | [**UpdateListingProperty**](docs/ShopListingApi.md#updatelistingproperty) | **Put** /v3/application/shops/{shop_id}/listings/{listing_id}/properties/{property_id} | 
*ShopListingFileApi* | [**DeleteListingFile**](docs/ShopListingFileApi.md#deletelistingfile) | **Delete** /v3/application/shops/{shop_id}/listings/{listing_id}/files/{listing_file_id} | 
*ShopListingFileApi* | [**GetAllListingFiles**](docs/ShopListingFileApi.md#getalllistingfiles) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/files | 
*ShopListingFileApi* | [**GetListingFile**](docs/ShopListingFileApi.md#getlistingfile) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/files/{listing_file_id} | 
*ShopListingFileApi* | [**UploadListingFile**](docs/ShopListingFileApi.md#uploadlistingfile) | **Post** /v3/application/shops/{shop_id}/listings/{listing_id}/files | 
*ShopListingImageApi* | [**DeleteListingImage**](docs/ShopListingImageApi.md#deletelistingimage) | **Delete** /v3/application/shops/{shop_id}/listings/{listing_id}/images/{listing_image_id} | 
*ShopListingImageApi* | [**GetListingImage**](docs/ShopListingImageApi.md#getlistingimage) | **Get** /v3/application/listings/{listing_id}/images/{listing_image_id} | 
*ShopListingImageApi* | [**GetListingImageDeprecated**](docs/ShopListingImageApi.md#getlistingimagedeprecated) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/images/{listing_image_id} | 
*ShopListingImageApi* | [**GetListingImages**](docs/ShopListingImageApi.md#getlistingimages) | **Get** /v3/application/listings/{listing_id}/images | 
*ShopListingImageApi* | [**GetListingImagesDeprecated**](docs/ShopListingImageApi.md#getlistingimagesdeprecated) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/images | 
*ShopListingImageApi* | [**UploadListingImage**](docs/ShopListingImageApi.md#uploadlistingimage) | **Post** /v3/application/shops/{shop_id}/listings/{listing_id}/images | 
*ShopListingInventoryApi* | [**GetListingInventory**](docs/ShopListingInventoryApi.md#getlistinginventory) | **Get** /v3/application/listings/{listing_id}/inventory | 
*ShopListingInventoryApi* | [**UpdateListingInventory**](docs/ShopListingInventoryApi.md#updatelistinginventory) | **Put** /v3/application/listings/{listing_id}/inventory | 
*ShopListingOfferingApi* | [**GetListingOffering**](docs/ShopListingOfferingApi.md#getlistingoffering) | **Get** /v3/application/listings/{listing_id}/products/{product_id}/offerings/{product_offering_id} | 
*ShopListingProductApi* | [**GetListingProduct**](docs/ShopListingProductApi.md#getlistingproduct) | **Get** /v3/application/listings/{listing_id}/inventory/products/{product_id} | 
*ShopListingTranslationApi* | [**CreateListingTranslation**](docs/ShopListingTranslationApi.md#createlistingtranslation) | **Post** /v3/application/shops/{shop_id}/listings/{listing_id}/translations/{language} | 
*ShopListingTranslationApi* | [**GetListingTranslation**](docs/ShopListingTranslationApi.md#getlistingtranslation) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/translations/{language} | 
*ShopListingTranslationApi* | [**UpdateListingTranslation**](docs/ShopListingTranslationApi.md#updatelistingtranslation) | **Put** /v3/application/shops/{shop_id}/listings/{listing_id}/translations/{language} | 
*ShopListingVariationImageApi* | [**GetListingVariationImages**](docs/ShopListingVariationImageApi.md#getlistingvariationimages) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/variation-images | 
*ShopListingVariationImageApi* | [**UpdateVariationImages**](docs/ShopListingVariationImageApi.md#updatevariationimages) | **Post** /v3/application/shops/{shop_id}/listings/{listing_id}/variation-images | 
*ShopListingVideoApi* | [**DeleteListingVideo**](docs/ShopListingVideoApi.md#deletelistingvideo) | **Delete** /v3/application/shops/{shop_id}/listings/{listing_id}/videos/{video_id} | 
*ShopListingVideoApi* | [**GetListingVideo**](docs/ShopListingVideoApi.md#getlistingvideo) | **Get** /v3/application/listings/{listing_id}/videos/{video_id} | 
*ShopListingVideoApi* | [**GetListingVideos**](docs/ShopListingVideoApi.md#getlistingvideos) | **Get** /v3/application/listings/{listing_id}/videos | 
*ShopListingVideoApi* | [**UploadListingVideo**](docs/ShopListingVideoApi.md#uploadlistingvideo) | **Post** /v3/application/shops/{shop_id}/listings/{listing_id}/videos | 
*ShopProductionPartnerApi* | [**GetShopProductionPartners**](docs/ShopProductionPartnerApi.md#getshopproductionpartners) | **Get** /v3/application/shops/{shop_id}/production-partners | 
*ShopReceiptApi* | [**CreateReceiptShipment**](docs/ShopReceiptApi.md#createreceiptshipment) | **Post** /v3/application/shops/{shop_id}/receipts/{receipt_id}/tracking | 
*ShopReceiptApi* | [**GetShopReceipt**](docs/ShopReceiptApi.md#getshopreceipt) | **Get** /v3/application/shops/{shop_id}/receipts/{receipt_id} | 
*ShopReceiptApi* | [**GetShopReceipts**](docs/ShopReceiptApi.md#getshopreceipts) | **Get** /v3/application/shops/{shop_id}/receipts | 
*ShopReceiptApi* | [**UpdateShopReceipt**](docs/ShopReceiptApi.md#updateshopreceipt) | **Put** /v3/application/shops/{shop_id}/receipts/{receipt_id} | 
*ShopReceiptTransactionsApi* | [**GetShopReceiptTransaction**](docs/ShopReceiptTransactionsApi.md#getshopreceipttransaction) | **Get** /v3/application/shops/{shop_id}/transactions/{transaction_id} | 
*ShopReceiptTransactionsApi* | [**GetShopReceiptTransactionsByListing**](docs/ShopReceiptTransactionsApi.md#getshopreceipttransactionsbylisting) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/transactions | 
*ShopReceiptTransactionsApi* | [**GetShopReceiptTransactionsByReceipt**](docs/ShopReceiptTransactionsApi.md#getshopreceipttransactionsbyreceipt) | **Get** /v3/application/shops/{shop_id}/receipts/{receipt_id}/transactions | 
*ShopReceiptTransactionsApi* | [**GetShopReceiptTransactionsByShop**](docs/ShopReceiptTransactionsApi.md#getshopreceipttransactionsbyshop) | **Get** /v3/application/shops/{shop_id}/transactions | 
*ShopReturnPolicyApi* | [**ConsolidateShopReturnPolicies**](docs/ShopReturnPolicyApi.md#consolidateshopreturnpolicies) | **Post** /v3/application/shops/{shop_id}/policies/return/consolidate | 
*ShopReturnPolicyApi* | [**CreateShopReturnPolicy**](docs/ShopReturnPolicyApi.md#createshopreturnpolicy) | **Post** /v3/application/shops/{shop_id}/policies/return | 
*ShopReturnPolicyApi* | [**DeleteShopReturnPolicy**](docs/ShopReturnPolicyApi.md#deleteshopreturnpolicy) | **Delete** /v3/application/shops/{shop_id}/policies/return/{return_policy_id} | 
*ShopReturnPolicyApi* | [**GetShopReturnPolicies**](docs/ShopReturnPolicyApi.md#getshopreturnpolicies) | **Get** /v3/application/shops/{shop_id}/policies/return | 
*ShopReturnPolicyApi* | [**GetShopReturnPolicy**](docs/ShopReturnPolicyApi.md#getshopreturnpolicy) | **Get** /v3/application/shops/{shop_id}/policies/return/{return_policy_id} | 
*ShopReturnPolicyApi* | [**UpdateShopReturnPolicy**](docs/ShopReturnPolicyApi.md#updateshopreturnpolicy) | **Put** /v3/application/shops/{shop_id}/policies/return/{return_policy_id} | 
*ShopSectionApi* | [**CreateShopSection**](docs/ShopSectionApi.md#createshopsection) | **Post** /v3/application/shops/{shop_id}/sections | 
*ShopSectionApi* | [**DeleteShopSection**](docs/ShopSectionApi.md#deleteshopsection) | **Delete** /v3/application/shops/{shop_id}/sections/{shop_section_id} | 
*ShopSectionApi* | [**GetShopSection**](docs/ShopSectionApi.md#getshopsection) | **Get** /v3/application/shops/{shop_id}/sections/{shop_section_id} | 
*ShopSectionApi* | [**GetShopSections**](docs/ShopSectionApi.md#getshopsections) | **Get** /v3/application/shops/{shop_id}/sections | 
*ShopSectionApi* | [**UpdateShopSection**](docs/ShopSectionApi.md#updateshopsection) | **Put** /v3/application/shops/{shop_id}/sections/{shop_section_id} | 
*ShopShippingProfileApi* | [**CreateShopShippingProfile**](docs/ShopShippingProfileApi.md#createshopshippingprofile) | **Post** /v3/application/shops/{shop_id}/shipping-profiles | 
*ShopShippingProfileApi* | [**CreateShopShippingProfileDestination**](docs/ShopShippingProfileApi.md#createshopshippingprofiledestination) | **Post** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/destinations | 
*ShopShippingProfileApi* | [**CreateShopShippingProfileUpgrade**](docs/ShopShippingProfileApi.md#createshopshippingprofileupgrade) | **Post** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/upgrades | 
*ShopShippingProfileApi* | [**DeleteShopShippingProfile**](docs/ShopShippingProfileApi.md#deleteshopshippingprofile) | **Delete** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id} | 
*ShopShippingProfileApi* | [**DeleteShopShippingProfileDestination**](docs/ShopShippingProfileApi.md#deleteshopshippingprofiledestination) | **Delete** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/destinations/{shipping_profile_destination_id} | 
*ShopShippingProfileApi* | [**DeleteShopShippingProfileUpgrade**](docs/ShopShippingProfileApi.md#deleteshopshippingprofileupgrade) | **Delete** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/upgrades/{upgrade_id} | 
*ShopShippingProfileApi* | [**GetShippingCarriers**](docs/ShopShippingProfileApi.md#getshippingcarriers) | **Get** /v3/application/shipping-carriers | 
*ShopShippingProfileApi* | [**GetShopShippingProfile**](docs/ShopShippingProfileApi.md#getshopshippingprofile) | **Get** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id} | 
*ShopShippingProfileApi* | [**GetShopShippingProfileDestinationsByShippingProfile**](docs/ShopShippingProfileApi.md#getshopshippingprofiledestinationsbyshippingprofile) | **Get** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/destinations | 
*ShopShippingProfileApi* | [**GetShopShippingProfileUpgrades**](docs/ShopShippingProfileApi.md#getshopshippingprofileupgrades) | **Get** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/upgrades | 
*ShopShippingProfileApi* | [**GetShopShippingProfiles**](docs/ShopShippingProfileApi.md#getshopshippingprofiles) | **Get** /v3/application/shops/{shop_id}/shipping-profiles | 
*ShopShippingProfileApi* | [**UpdateShopShippingProfile**](docs/ShopShippingProfileApi.md#updateshopshippingprofile) | **Put** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id} | 
*ShopShippingProfileApi* | [**UpdateShopShippingProfileDestination**](docs/ShopShippingProfileApi.md#updateshopshippingprofiledestination) | **Put** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/destinations/{shipping_profile_destination_id} | 
*ShopShippingProfileApi* | [**UpdateShopShippingProfileUpgrade**](docs/ShopShippingProfileApi.md#updateshopshippingprofileupgrade) | **Put** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/upgrades/{upgrade_id} | 
*UserApi* | [**GetMe**](docs/UserApi.md#getme) | **Get** /v3/application/users/me | 
*UserApi* | [**GetUser**](docs/UserApi.md#getuser) | **Get** /v3/application/users/{user_id} | 
*UserAddressApi* | [**DeleteUserAddress**](docs/UserAddressApi.md#deleteuseraddress) | **Delete** /v3/application/user/addresses/{user_address_id} | 
*UserAddressApi* | [**GetUserAddress**](docs/UserAddressApi.md#getuseraddress) | **Get** /v3/application/user/addresses/{user_address_id} | 
*UserAddressApi* | [**GetUserAddresses**](docs/UserAddressApi.md#getuseraddresses) | **Get** /v3/application/user/addresses | 


## Documentation For Models

 - [BuyerTaxonomyNode](docs/BuyerTaxonomyNode.md)
 - [BuyerTaxonomyNodeChildrenInner](docs/BuyerTaxonomyNodeChildrenInner.md)
 - [BuyerTaxonomyNodeProperties](docs/BuyerTaxonomyNodeProperties.md)
 - [BuyerTaxonomyNodePropertiesResultsInner](docs/BuyerTaxonomyNodePropertiesResultsInner.md)
 - [BuyerTaxonomyNodeProperty](docs/BuyerTaxonomyNodeProperty.md)
 - [BuyerTaxonomyNodePropertyPossibleValuesInner](docs/BuyerTaxonomyNodePropertyPossibleValuesInner.md)
 - [BuyerTaxonomyNodePropertyScalesInner](docs/BuyerTaxonomyNodePropertyScalesInner.md)
 - [BuyerTaxonomyNodePropertySelectedValuesInner](docs/BuyerTaxonomyNodePropertySelectedValuesInner.md)
 - [BuyerTaxonomyNodes](docs/BuyerTaxonomyNodes.md)
 - [BuyerTaxonomyNodesResultsInner](docs/BuyerTaxonomyNodesResultsInner.md)
 - [BuyerTaxonomyPropertyScale](docs/BuyerTaxonomyPropertyScale.md)
 - [BuyerTaxonomyPropertyValue](docs/BuyerTaxonomyPropertyValue.md)
 - [ErrorSchema](docs/ErrorSchema.md)
 - [ListingImage](docs/ListingImage.md)
 - [ListingImages](docs/ListingImages.md)
 - [ListingImagesResultsInner](docs/ListingImagesResultsInner.md)
 - [ListingInventory](docs/ListingInventory.md)
 - [ListingInventoryProduct](docs/ListingInventoryProduct.md)
 - [ListingInventoryProductOffering](docs/ListingInventoryProductOffering.md)
 - [ListingInventoryProductOfferingPrice](docs/ListingInventoryProductOfferingPrice.md)
 - [ListingInventoryProductOfferingsInner](docs/ListingInventoryProductOfferingsInner.md)
 - [ListingInventoryProductPropertyValuesInner](docs/ListingInventoryProductPropertyValuesInner.md)
 - [ListingInventoryProductsInner](docs/ListingInventoryProductsInner.md)
 - [ListingInventoryWithAssociations](docs/ListingInventoryWithAssociations.md)
 - [ListingInventoryWithAssociationsListing](docs/ListingInventoryWithAssociationsListing.md)
 - [ListingPropertyValue](docs/ListingPropertyValue.md)
 - [ListingPropertyValues](docs/ListingPropertyValues.md)
 - [ListingPropertyValuesResultsInner](docs/ListingPropertyValuesResultsInner.md)
 - [ListingReview](docs/ListingReview.md)
 - [ListingReviews](docs/ListingReviews.md)
 - [ListingReviewsResultsInner](docs/ListingReviewsResultsInner.md)
 - [ListingTranslation](docs/ListingTranslation.md)
 - [ListingVariationImage](docs/ListingVariationImage.md)
 - [ListingVariationImages](docs/ListingVariationImages.md)
 - [ListingVariationImagesResultsInner](docs/ListingVariationImagesResultsInner.md)
 - [ListingVideo](docs/ListingVideo.md)
 - [ListingVideos](docs/ListingVideos.md)
 - [ListingVideosResultsInner](docs/ListingVideosResultsInner.md)
 - [Money](docs/Money.md)
 - [Payment](docs/Payment.md)
 - [PaymentAccountLedgerEntries](docs/PaymentAccountLedgerEntries.md)
 - [PaymentAccountLedgerEntriesResultsInner](docs/PaymentAccountLedgerEntriesResultsInner.md)
 - [PaymentAccountLedgerEntry](docs/PaymentAccountLedgerEntry.md)
 - [PaymentAccountLedgerEntryPaymentAdjustmentsInner](docs/PaymentAccountLedgerEntryPaymentAdjustmentsInner.md)
 - [PaymentAdjustedFees](docs/PaymentAdjustedFees.md)
 - [PaymentAdjustedGross](docs/PaymentAdjustedGross.md)
 - [PaymentAdjustedNet](docs/PaymentAdjustedNet.md)
 - [PaymentAdjustment](docs/PaymentAdjustment.md)
 - [PaymentAdjustmentItem](docs/PaymentAdjustmentItem.md)
 - [PaymentAdjustmentPaymentAdjustmentItemsInner](docs/PaymentAdjustmentPaymentAdjustmentItemsInner.md)
 - [PaymentAmountFees](docs/PaymentAmountFees.md)
 - [PaymentAmountGross](docs/PaymentAmountGross.md)
 - [PaymentAmountNet](docs/PaymentAmountNet.md)
 - [PaymentPostedFees](docs/PaymentPostedFees.md)
 - [PaymentPostedGross](docs/PaymentPostedGross.md)
 - [PaymentPostedNet](docs/PaymentPostedNet.md)
 - [Payments](docs/Payments.md)
 - [PaymentsResultsInner](docs/PaymentsResultsInner.md)
 - [Pong](docs/Pong.md)
 - [Self](docs/Self.md)
 - [SellerTaxonomyNode](docs/SellerTaxonomyNode.md)
 - [SellerTaxonomyNodeChildrenInner](docs/SellerTaxonomyNodeChildrenInner.md)
 - [SellerTaxonomyNodes](docs/SellerTaxonomyNodes.md)
 - [SellerTaxonomyNodesResultsInner](docs/SellerTaxonomyNodesResultsInner.md)
 - [ShippingCarrier](docs/ShippingCarrier.md)
 - [ShippingCarrierDomesticClassesInner](docs/ShippingCarrierDomesticClassesInner.md)
 - [ShippingCarrierInternationalClassesInner](docs/ShippingCarrierInternationalClassesInner.md)
 - [ShippingCarrierMailClass](docs/ShippingCarrierMailClass.md)
 - [ShippingCarriers](docs/ShippingCarriers.md)
 - [ShippingCarriersResultsInner](docs/ShippingCarriersResultsInner.md)
 - [Shop](docs/Shop.md)
 - [ShopListing](docs/ShopListing.md)
 - [ShopListingFile](docs/ShopListingFile.md)
 - [ShopListingFiles](docs/ShopListingFiles.md)
 - [ShopListingFilesResultsInner](docs/ShopListingFilesResultsInner.md)
 - [ShopListingPrice](docs/ShopListingPrice.md)
 - [ShopListingWithAssociations](docs/ShopListingWithAssociations.md)
 - [ShopListingWithAssociationsImagesInner](docs/ShopListingWithAssociationsImagesInner.md)
 - [ShopListingWithAssociationsInventory](docs/ShopListingWithAssociationsInventory.md)
 - [ShopListingWithAssociationsProductionPartnersInner](docs/ShopListingWithAssociationsProductionPartnersInner.md)
 - [ShopListingWithAssociationsShippingProfile](docs/ShopListingWithAssociationsShippingProfile.md)
 - [ShopListingWithAssociationsShop](docs/ShopListingWithAssociationsShop.md)
 - [ShopListingWithAssociationsTranslationsInner](docs/ShopListingWithAssociationsTranslationsInner.md)
 - [ShopListingWithAssociationsUser](docs/ShopListingWithAssociationsUser.md)
 - [ShopListingWithAssociationsVideosInner](docs/ShopListingWithAssociationsVideosInner.md)
 - [ShopListings](docs/ShopListings.md)
 - [ShopListingsResultsInner](docs/ShopListingsResultsInner.md)
 - [ShopListingsWithAssociations](docs/ShopListingsWithAssociations.md)
 - [ShopListingsWithAssociationsResultsInner](docs/ShopListingsWithAssociationsResultsInner.md)
 - [ShopProductionPartner](docs/ShopProductionPartner.md)
 - [ShopProductionPartners](docs/ShopProductionPartners.md)
 - [ShopProductionPartnersResultsInner](docs/ShopProductionPartnersResultsInner.md)
 - [ShopReceipt](docs/ShopReceipt.md)
 - [ShopReceiptDiscountAmt](docs/ShopReceiptDiscountAmt.md)
 - [ShopReceiptGiftWrapPrice](docs/ShopReceiptGiftWrapPrice.md)
 - [ShopReceiptGrandtotal](docs/ShopReceiptGrandtotal.md)
 - [ShopReceiptRefundsInner](docs/ShopReceiptRefundsInner.md)
 - [ShopReceiptShipment](docs/ShopReceiptShipment.md)
 - [ShopReceiptShipmentsInner](docs/ShopReceiptShipmentsInner.md)
 - [ShopReceiptSubtotal](docs/ShopReceiptSubtotal.md)
 - [ShopReceiptTotalPrice](docs/ShopReceiptTotalPrice.md)
 - [ShopReceiptTotalShippingCost](docs/ShopReceiptTotalShippingCost.md)
 - [ShopReceiptTotalTaxCost](docs/ShopReceiptTotalTaxCost.md)
 - [ShopReceiptTotalVatCost](docs/ShopReceiptTotalVatCost.md)
 - [ShopReceiptTransaction](docs/ShopReceiptTransaction.md)
 - [ShopReceiptTransactionPrice](docs/ShopReceiptTransactionPrice.md)
 - [ShopReceiptTransactionShippingCost](docs/ShopReceiptTransactionShippingCost.md)
 - [ShopReceiptTransactionVariationsInner](docs/ShopReceiptTransactionVariationsInner.md)
 - [ShopReceiptTransactions](docs/ShopReceiptTransactions.md)
 - [ShopReceiptTransactionsInner](docs/ShopReceiptTransactionsInner.md)
 - [ShopReceiptTransactionsResultsInner](docs/ShopReceiptTransactionsResultsInner.md)
 - [ShopReceipts](docs/ShopReceipts.md)
 - [ShopReceiptsResultsInner](docs/ShopReceiptsResultsInner.md)
 - [ShopRefund](docs/ShopRefund.md)
 - [ShopRefundAmount](docs/ShopRefundAmount.md)
 - [ShopReturnPolicies](docs/ShopReturnPolicies.md)
 - [ShopReturnPoliciesResultsInner](docs/ShopReturnPoliciesResultsInner.md)
 - [ShopReturnPolicy](docs/ShopReturnPolicy.md)
 - [ShopSection](docs/ShopSection.md)
 - [ShopSections](docs/ShopSections.md)
 - [ShopSectionsResultsInner](docs/ShopSectionsResultsInner.md)
 - [ShopShippingProfile](docs/ShopShippingProfile.md)
 - [ShopShippingProfileDestination](docs/ShopShippingProfileDestination.md)
 - [ShopShippingProfileDestinationPrimaryCost](docs/ShopShippingProfileDestinationPrimaryCost.md)
 - [ShopShippingProfileDestinationSecondaryCost](docs/ShopShippingProfileDestinationSecondaryCost.md)
 - [ShopShippingProfileDestinations](docs/ShopShippingProfileDestinations.md)
 - [ShopShippingProfileDestinationsResultsInner](docs/ShopShippingProfileDestinationsResultsInner.md)
 - [ShopShippingProfileShippingProfileDestinationsInner](docs/ShopShippingProfileShippingProfileDestinationsInner.md)
 - [ShopShippingProfileShippingProfileUpgradesInner](docs/ShopShippingProfileShippingProfileUpgradesInner.md)
 - [ShopShippingProfileUpgrade](docs/ShopShippingProfileUpgrade.md)
 - [ShopShippingProfileUpgradePrice](docs/ShopShippingProfileUpgradePrice.md)
 - [ShopShippingProfileUpgradeSecondaryPrice](docs/ShopShippingProfileUpgradeSecondaryPrice.md)
 - [ShopShippingProfileUpgrades](docs/ShopShippingProfileUpgrades.md)
 - [ShopShippingProfileUpgradesResultsInner](docs/ShopShippingProfileUpgradesResultsInner.md)
 - [ShopShippingProfiles](docs/ShopShippingProfiles.md)
 - [ShopShippingProfilesResultsInner](docs/ShopShippingProfilesResultsInner.md)
 - [Shops](docs/Shops.md)
 - [ShopsResultsInner](docs/ShopsResultsInner.md)
 - [TaxonomyNodeProperties](docs/TaxonomyNodeProperties.md)
 - [TaxonomyNodePropertiesResultsInner](docs/TaxonomyNodePropertiesResultsInner.md)
 - [TaxonomyNodeProperty](docs/TaxonomyNodeProperty.md)
 - [TaxonomyNodePropertyPossibleValuesInner](docs/TaxonomyNodePropertyPossibleValuesInner.md)
 - [TaxonomyNodePropertyScalesInner](docs/TaxonomyNodePropertyScalesInner.md)
 - [TaxonomyNodePropertySelectedValuesInner](docs/TaxonomyNodePropertySelectedValuesInner.md)
 - [TaxonomyPropertyScale](docs/TaxonomyPropertyScale.md)
 - [TaxonomyPropertyValue](docs/TaxonomyPropertyValue.md)
 - [TransactionReview](docs/TransactionReview.md)
 - [TransactionReviews](docs/TransactionReviews.md)
 - [TransactionReviewsResultsInner](docs/TransactionReviewsResultsInner.md)
 - [TransactionVariations](docs/TransactionVariations.md)
 - [UpdateListingInventoryRequest](docs/UpdateListingInventoryRequest.md)
 - [UpdateListingInventoryRequestProductsInner](docs/UpdateListingInventoryRequestProductsInner.md)
 - [UpdateListingInventoryRequestProductsInnerOfferingsInner](docs/UpdateListingInventoryRequestProductsInnerOfferingsInner.md)
 - [UpdateListingInventoryRequestProductsInnerPropertyValuesInner](docs/UpdateListingInventoryRequestProductsInnerPropertyValuesInner.md)
 - [UpdateVariationImagesRequest](docs/UpdateVariationImagesRequest.md)
 - [UpdateVariationImagesRequestVariationImagesInner](docs/UpdateVariationImagesRequestVariationImagesInner.md)
 - [User](docs/User.md)
 - [UserAddress](docs/UserAddress.md)
 - [UserAddresses](docs/UserAddresses.md)
 - [UserAddressesResultsInner](docs/UserAddressesResultsInner.md)


## Documentation For Authorization



### api_key

- **Type**: API key
- **API key parameter name**: x-api-key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: x-api-key and passed in as the auth context for each request.


### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://www.etsy.com/oauth/connect
- **Scopes**: 
 - **address_r**: see billing and shipping addresses
 - **address_w**: update billing and shipping addresses
 - **billing_r**: see all billing statement data
 - **cart_r**: read shopping carts
 - **cart_w**: add/remove from shopping carts
 - **email_r**: Read a member's email address
 - **favorites_r**: see private favorites
 - **favorites_w**: add/remove favorites
 - **feedback_r**: see purchase info in feedback
 - **listings_d**: delete listings
 - **listings_r**: see all listings (including expired etc)
 - **listings_w**: create/edit listings
 - **profile_r**: see all profile data
 - **profile_w**: update user profile, avatar, etc
 - **recommend_r**: see recommended listings
 - **recommend_w**: accept/reject recommended listings
 - **shops_r**: see private shop info
 - **shops_w**: update shop
 - **transactions_r**: see all checkout/payment data
 - **transactions_w**: update receipts

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

developers@etsy.com

