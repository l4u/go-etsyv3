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

// checks if the ShopReceiptGrandtotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShopReceiptGrandtotal{}

// ShopReceiptGrandtotal A number equal to the total_price minus the coupon discount plus tax and shipping costs.
type ShopReceiptGrandtotal struct {
	// The amount of represented by this data.
	Amount *int32 `json:"amount,omitempty"`
	// The divisor to render the amount.
	Divisor *int32 `json:"divisor,omitempty"`
	// The ISO currency code for this data.
	CurrencyCode *string `json:"currency_code,omitempty"`
}

// NewShopReceiptGrandtotal instantiates a new ShopReceiptGrandtotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShopReceiptGrandtotal() *ShopReceiptGrandtotal {
	this := ShopReceiptGrandtotal{}
	return &this
}

// NewShopReceiptGrandtotalWithDefaults instantiates a new ShopReceiptGrandtotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShopReceiptGrandtotalWithDefaults() *ShopReceiptGrandtotal {
	this := ShopReceiptGrandtotal{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ShopReceiptGrandtotal) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopReceiptGrandtotal) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ShopReceiptGrandtotal) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *ShopReceiptGrandtotal) SetAmount(v int32) {
	o.Amount = &v
}

// GetDivisor returns the Divisor field value if set, zero value otherwise.
func (o *ShopReceiptGrandtotal) GetDivisor() int32 {
	if o == nil || IsNil(o.Divisor) {
		var ret int32
		return ret
	}
	return *o.Divisor
}

// GetDivisorOk returns a tuple with the Divisor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopReceiptGrandtotal) GetDivisorOk() (*int32, bool) {
	if o == nil || IsNil(o.Divisor) {
		return nil, false
	}
	return o.Divisor, true
}

// HasDivisor returns a boolean if a field has been set.
func (o *ShopReceiptGrandtotal) HasDivisor() bool {
	if o != nil && !IsNil(o.Divisor) {
		return true
	}

	return false
}

// SetDivisor gets a reference to the given int32 and assigns it to the Divisor field.
func (o *ShopReceiptGrandtotal) SetDivisor(v int32) {
	o.Divisor = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *ShopReceiptGrandtotal) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopReceiptGrandtotal) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ShopReceiptGrandtotal) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *ShopReceiptGrandtotal) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

func (o ShopReceiptGrandtotal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShopReceiptGrandtotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Divisor) {
		toSerialize["divisor"] = o.Divisor
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	return toSerialize, nil
}

type NullableShopReceiptGrandtotal struct {
	value *ShopReceiptGrandtotal
	isSet bool
}

func (v NullableShopReceiptGrandtotal) Get() *ShopReceiptGrandtotal {
	return v.value
}

func (v *NullableShopReceiptGrandtotal) Set(val *ShopReceiptGrandtotal) {
	v.value = val
	v.isSet = true
}

func (v NullableShopReceiptGrandtotal) IsSet() bool {
	return v.isSet
}

func (v *NullableShopReceiptGrandtotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShopReceiptGrandtotal(val *ShopReceiptGrandtotal) *NullableShopReceiptGrandtotal {
	return &NullableShopReceiptGrandtotal{value: val, isSet: true}
}

func (v NullableShopReceiptGrandtotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShopReceiptGrandtotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


