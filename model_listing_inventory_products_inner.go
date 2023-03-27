/*
Etsy Open API v3

<div class=\"wt-text-body-01\"><p class=\"wt-pt-xs-2 wt-pb-xs-2\">Etsy's Open API provides a simple RESTful interface for various Etsy.com features. The API endpoints are meant to replace Etsy's Open API v2, which is scheduled to end service in 2022.</p><p class=\"wt-pb-xs-2\">All of the endpoints are callable and the majority of the API endpoints are now in a beta phase. This means we do not expect to make any breaking changes before our general release. A handful of endpoints are currently interface stubs (labeled “Feedback Only”) and returns a \"501 Not Implemented\" response code when called.</p><p class=\"wt-pb-xs-2\">If you'd like to report an issue or provide feedback on the API design, <a target=\"_blank\" class=\"wt-text-link wt-p-xs-0\" href=\"https://github.com/etsy/open-api/discussions\">please add an issue in Github</a>.</p></div>&copy; 2021-2023 Etsy, Inc. All Rights Reserved. Use of this code is subject to Etsy's <a class='wt-text-link wt-p-xs-0' target='_blank' href='https://www.etsy.com/legal/api'>API Developer Terms of Use</a>.

API version: 3.0.0
Contact: developers@etsy.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ListingInventoryProductsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListingInventoryProductsInner{}

// ListingInventoryProductsInner A JSON array of products available in a listing, even if only one product. All field names in the JSON blobs are lowercase.
type ListingInventoryProductsInner struct {
	// The numeric ID for a specific [product](/documentation/reference#tag/ShopListing-Product) purchased from a listing.
	ProductId *int32 `json:"product_id,omitempty"`
	// The SKU string for the product
	Sku *string `json:"sku,omitempty"`
	// When true, someone deleted this product.
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// A list of product offering entries for this product.
	Offerings []ListingInventoryProductOfferingsInner `json:"offerings,omitempty"`
	// A list of property value entries for this product. Note: parenthesis characters (`(` and `)`) are not allowed.
	PropertyValues []ListingInventoryProductPropertyValuesInner `json:"property_values,omitempty"`
}

// NewListingInventoryProductsInner instantiates a new ListingInventoryProductsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingInventoryProductsInner() *ListingInventoryProductsInner {
	this := ListingInventoryProductsInner{}
	return &this
}

// NewListingInventoryProductsInnerWithDefaults instantiates a new ListingInventoryProductsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingInventoryProductsInnerWithDefaults() *ListingInventoryProductsInner {
	this := ListingInventoryProductsInner{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *ListingInventoryProductsInner) GetProductId() int32 {
	if o == nil || IsNil(o.ProductId) {
		var ret int32
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingInventoryProductsInner) GetProductIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *ListingInventoryProductsInner) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int32 and assigns it to the ProductId field.
func (o *ListingInventoryProductsInner) SetProductId(v int32) {
	o.ProductId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ListingInventoryProductsInner) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingInventoryProductsInner) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ListingInventoryProductsInner) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ListingInventoryProductsInner) SetSku(v string) {
	o.Sku = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *ListingInventoryProductsInner) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingInventoryProductsInner) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *ListingInventoryProductsInner) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *ListingInventoryProductsInner) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetOfferings returns the Offerings field value if set, zero value otherwise.
func (o *ListingInventoryProductsInner) GetOfferings() []ListingInventoryProductOfferingsInner {
	if o == nil || IsNil(o.Offerings) {
		var ret []ListingInventoryProductOfferingsInner
		return ret
	}
	return o.Offerings
}

// GetOfferingsOk returns a tuple with the Offerings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingInventoryProductsInner) GetOfferingsOk() ([]ListingInventoryProductOfferingsInner, bool) {
	if o == nil || IsNil(o.Offerings) {
		return nil, false
	}
	return o.Offerings, true
}

// HasOfferings returns a boolean if a field has been set.
func (o *ListingInventoryProductsInner) HasOfferings() bool {
	if o != nil && !IsNil(o.Offerings) {
		return true
	}

	return false
}

// SetOfferings gets a reference to the given []ListingInventoryProductOfferingsInner and assigns it to the Offerings field.
func (o *ListingInventoryProductsInner) SetOfferings(v []ListingInventoryProductOfferingsInner) {
	o.Offerings = v
}

// GetPropertyValues returns the PropertyValues field value if set, zero value otherwise.
func (o *ListingInventoryProductsInner) GetPropertyValues() []ListingInventoryProductPropertyValuesInner {
	if o == nil || IsNil(o.PropertyValues) {
		var ret []ListingInventoryProductPropertyValuesInner
		return ret
	}
	return o.PropertyValues
}

// GetPropertyValuesOk returns a tuple with the PropertyValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingInventoryProductsInner) GetPropertyValuesOk() ([]ListingInventoryProductPropertyValuesInner, bool) {
	if o == nil || IsNil(o.PropertyValues) {
		return nil, false
	}
	return o.PropertyValues, true
}

// HasPropertyValues returns a boolean if a field has been set.
func (o *ListingInventoryProductsInner) HasPropertyValues() bool {
	if o != nil && !IsNil(o.PropertyValues) {
		return true
	}

	return false
}

// SetPropertyValues gets a reference to the given []ListingInventoryProductPropertyValuesInner and assigns it to the PropertyValues field.
func (o *ListingInventoryProductsInner) SetPropertyValues(v []ListingInventoryProductPropertyValuesInner) {
	o.PropertyValues = v
}

func (o ListingInventoryProductsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListingInventoryProductsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.Offerings) {
		toSerialize["offerings"] = o.Offerings
	}
	if !IsNil(o.PropertyValues) {
		toSerialize["property_values"] = o.PropertyValues
	}
	return toSerialize, nil
}

type NullableListingInventoryProductsInner struct {
	value *ListingInventoryProductsInner
	isSet bool
}

func (v NullableListingInventoryProductsInner) Get() *ListingInventoryProductsInner {
	return v.value
}

func (v *NullableListingInventoryProductsInner) Set(val *ListingInventoryProductsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListingInventoryProductsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListingInventoryProductsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingInventoryProductsInner(val *ListingInventoryProductsInner) *NullableListingInventoryProductsInner {
	return &NullableListingInventoryProductsInner{value: val, isSet: true}
}

func (v NullableListingInventoryProductsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListingInventoryProductsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


