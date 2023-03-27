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

// checks if the PaymentAdjustmentItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentAdjustmentItem{}

// PaymentAdjustmentItem A payemnt adjustment line item for a payment adjustment.
type PaymentAdjustmentItem struct {
	// The numeric ID for a payment adjustment.
	PaymentAdjustmentId *int32 `json:"payment_adjustment_id,omitempty"`
	// Unique ID for the adjustment line item.
	PaymentAdjustmentItemId *int32 `json:"payment_adjustment_item_id,omitempty"`
	// String indicating the type of adjustment for this line item.
	AdjustmentType NullableString `json:"adjustment_type,omitempty"`
	// Integer value for the amount of the adjustment in original currency.
	Amount *int32 `json:"amount,omitempty"`
	// Integer value for the amount of the adjustment in currency for the shop.
	ShopAmount *int32 `json:"shop_amount,omitempty"`
	// The unique numeric ID for a transaction.
	TransactionId NullableInt32 `json:"transaction_id,omitempty"`
	// Unique ID for the bill payment adjustment.
	BillPaymentId NullableInt32 `json:"bill_payment_id,omitempty"`
	// The transaction\\'s creation date and time, in epoch seconds.
	CreatedTimestamp *int32 `json:"created_timestamp,omitempty"`
	// The update date and time the payment adjustment in epoch seconds.
	UpdatedTimestamp *int32 `json:"updated_timestamp,omitempty"`
}

// NewPaymentAdjustmentItem instantiates a new PaymentAdjustmentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentAdjustmentItem() *PaymentAdjustmentItem {
	this := PaymentAdjustmentItem{}
	var amount int32 = 0
	this.Amount = &amount
	var shopAmount int32 = 0
	this.ShopAmount = &shopAmount
	return &this
}

// NewPaymentAdjustmentItemWithDefaults instantiates a new PaymentAdjustmentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentAdjustmentItemWithDefaults() *PaymentAdjustmentItem {
	this := PaymentAdjustmentItem{}
	var amount int32 = 0
	this.Amount = &amount
	var shopAmount int32 = 0
	this.ShopAmount = &shopAmount
	return &this
}

// GetPaymentAdjustmentId returns the PaymentAdjustmentId field value if set, zero value otherwise.
func (o *PaymentAdjustmentItem) GetPaymentAdjustmentId() int32 {
	if o == nil || IsNil(o.PaymentAdjustmentId) {
		var ret int32
		return ret
	}
	return *o.PaymentAdjustmentId
}

// GetPaymentAdjustmentIdOk returns a tuple with the PaymentAdjustmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAdjustmentItem) GetPaymentAdjustmentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PaymentAdjustmentId) {
		return nil, false
	}
	return o.PaymentAdjustmentId, true
}

// HasPaymentAdjustmentId returns a boolean if a field has been set.
func (o *PaymentAdjustmentItem) HasPaymentAdjustmentId() bool {
	if o != nil && !IsNil(o.PaymentAdjustmentId) {
		return true
	}

	return false
}

// SetPaymentAdjustmentId gets a reference to the given int32 and assigns it to the PaymentAdjustmentId field.
func (o *PaymentAdjustmentItem) SetPaymentAdjustmentId(v int32) {
	o.PaymentAdjustmentId = &v
}

// GetPaymentAdjustmentItemId returns the PaymentAdjustmentItemId field value if set, zero value otherwise.
func (o *PaymentAdjustmentItem) GetPaymentAdjustmentItemId() int32 {
	if o == nil || IsNil(o.PaymentAdjustmentItemId) {
		var ret int32
		return ret
	}
	return *o.PaymentAdjustmentItemId
}

// GetPaymentAdjustmentItemIdOk returns a tuple with the PaymentAdjustmentItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAdjustmentItem) GetPaymentAdjustmentItemIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PaymentAdjustmentItemId) {
		return nil, false
	}
	return o.PaymentAdjustmentItemId, true
}

// HasPaymentAdjustmentItemId returns a boolean if a field has been set.
func (o *PaymentAdjustmentItem) HasPaymentAdjustmentItemId() bool {
	if o != nil && !IsNil(o.PaymentAdjustmentItemId) {
		return true
	}

	return false
}

// SetPaymentAdjustmentItemId gets a reference to the given int32 and assigns it to the PaymentAdjustmentItemId field.
func (o *PaymentAdjustmentItem) SetPaymentAdjustmentItemId(v int32) {
	o.PaymentAdjustmentItemId = &v
}

// GetAdjustmentType returns the AdjustmentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentAdjustmentItem) GetAdjustmentType() string {
	if o == nil || IsNil(o.AdjustmentType.Get()) {
		var ret string
		return ret
	}
	return *o.AdjustmentType.Get()
}

// GetAdjustmentTypeOk returns a tuple with the AdjustmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAdjustmentItem) GetAdjustmentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdjustmentType.Get(), o.AdjustmentType.IsSet()
}

// HasAdjustmentType returns a boolean if a field has been set.
func (o *PaymentAdjustmentItem) HasAdjustmentType() bool {
	if o != nil && o.AdjustmentType.IsSet() {
		return true
	}

	return false
}

// SetAdjustmentType gets a reference to the given NullableString and assigns it to the AdjustmentType field.
func (o *PaymentAdjustmentItem) SetAdjustmentType(v string) {
	o.AdjustmentType.Set(&v)
}
// SetAdjustmentTypeNil sets the value for AdjustmentType to be an explicit nil
func (o *PaymentAdjustmentItem) SetAdjustmentTypeNil() {
	o.AdjustmentType.Set(nil)
}

// UnsetAdjustmentType ensures that no value is present for AdjustmentType, not even an explicit nil
func (o *PaymentAdjustmentItem) UnsetAdjustmentType() {
	o.AdjustmentType.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaymentAdjustmentItem) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAdjustmentItem) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentAdjustmentItem) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *PaymentAdjustmentItem) SetAmount(v int32) {
	o.Amount = &v
}

// GetShopAmount returns the ShopAmount field value if set, zero value otherwise.
func (o *PaymentAdjustmentItem) GetShopAmount() int32 {
	if o == nil || IsNil(o.ShopAmount) {
		var ret int32
		return ret
	}
	return *o.ShopAmount
}

// GetShopAmountOk returns a tuple with the ShopAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAdjustmentItem) GetShopAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.ShopAmount) {
		return nil, false
	}
	return o.ShopAmount, true
}

// HasShopAmount returns a boolean if a field has been set.
func (o *PaymentAdjustmentItem) HasShopAmount() bool {
	if o != nil && !IsNil(o.ShopAmount) {
		return true
	}

	return false
}

// SetShopAmount gets a reference to the given int32 and assigns it to the ShopAmount field.
func (o *PaymentAdjustmentItem) SetShopAmount(v int32) {
	o.ShopAmount = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentAdjustmentItem) GetTransactionId() int32 {
	if o == nil || IsNil(o.TransactionId.Get()) {
		var ret int32
		return ret
	}
	return *o.TransactionId.Get()
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAdjustmentItem) GetTransactionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionId.Get(), o.TransactionId.IsSet()
}

// HasTransactionId returns a boolean if a field has been set.
func (o *PaymentAdjustmentItem) HasTransactionId() bool {
	if o != nil && o.TransactionId.IsSet() {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given NullableInt32 and assigns it to the TransactionId field.
func (o *PaymentAdjustmentItem) SetTransactionId(v int32) {
	o.TransactionId.Set(&v)
}
// SetTransactionIdNil sets the value for TransactionId to be an explicit nil
func (o *PaymentAdjustmentItem) SetTransactionIdNil() {
	o.TransactionId.Set(nil)
}

// UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
func (o *PaymentAdjustmentItem) UnsetTransactionId() {
	o.TransactionId.Unset()
}

// GetBillPaymentId returns the BillPaymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentAdjustmentItem) GetBillPaymentId() int32 {
	if o == nil || IsNil(o.BillPaymentId.Get()) {
		var ret int32
		return ret
	}
	return *o.BillPaymentId.Get()
}

// GetBillPaymentIdOk returns a tuple with the BillPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAdjustmentItem) GetBillPaymentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillPaymentId.Get(), o.BillPaymentId.IsSet()
}

// HasBillPaymentId returns a boolean if a field has been set.
func (o *PaymentAdjustmentItem) HasBillPaymentId() bool {
	if o != nil && o.BillPaymentId.IsSet() {
		return true
	}

	return false
}

// SetBillPaymentId gets a reference to the given NullableInt32 and assigns it to the BillPaymentId field.
func (o *PaymentAdjustmentItem) SetBillPaymentId(v int32) {
	o.BillPaymentId.Set(&v)
}
// SetBillPaymentIdNil sets the value for BillPaymentId to be an explicit nil
func (o *PaymentAdjustmentItem) SetBillPaymentIdNil() {
	o.BillPaymentId.Set(nil)
}

// UnsetBillPaymentId ensures that no value is present for BillPaymentId, not even an explicit nil
func (o *PaymentAdjustmentItem) UnsetBillPaymentId() {
	o.BillPaymentId.Unset()
}

// GetCreatedTimestamp returns the CreatedTimestamp field value if set, zero value otherwise.
func (o *PaymentAdjustmentItem) GetCreatedTimestamp() int32 {
	if o == nil || IsNil(o.CreatedTimestamp) {
		var ret int32
		return ret
	}
	return *o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAdjustmentItem) GetCreatedTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedTimestamp) {
		return nil, false
	}
	return o.CreatedTimestamp, true
}

// HasCreatedTimestamp returns a boolean if a field has been set.
func (o *PaymentAdjustmentItem) HasCreatedTimestamp() bool {
	if o != nil && !IsNil(o.CreatedTimestamp) {
		return true
	}

	return false
}

// SetCreatedTimestamp gets a reference to the given int32 and assigns it to the CreatedTimestamp field.
func (o *PaymentAdjustmentItem) SetCreatedTimestamp(v int32) {
	o.CreatedTimestamp = &v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value if set, zero value otherwise.
func (o *PaymentAdjustmentItem) GetUpdatedTimestamp() int32 {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		var ret int32
		return ret
	}
	return *o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAdjustmentItem) GetUpdatedTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		return nil, false
	}
	return o.UpdatedTimestamp, true
}

// HasUpdatedTimestamp returns a boolean if a field has been set.
func (o *PaymentAdjustmentItem) HasUpdatedTimestamp() bool {
	if o != nil && !IsNil(o.UpdatedTimestamp) {
		return true
	}

	return false
}

// SetUpdatedTimestamp gets a reference to the given int32 and assigns it to the UpdatedTimestamp field.
func (o *PaymentAdjustmentItem) SetUpdatedTimestamp(v int32) {
	o.UpdatedTimestamp = &v
}

func (o PaymentAdjustmentItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentAdjustmentItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentAdjustmentId) {
		toSerialize["payment_adjustment_id"] = o.PaymentAdjustmentId
	}
	if !IsNil(o.PaymentAdjustmentItemId) {
		toSerialize["payment_adjustment_item_id"] = o.PaymentAdjustmentItemId
	}
	if o.AdjustmentType.IsSet() {
		toSerialize["adjustment_type"] = o.AdjustmentType.Get()
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.ShopAmount) {
		toSerialize["shop_amount"] = o.ShopAmount
	}
	if o.TransactionId.IsSet() {
		toSerialize["transaction_id"] = o.TransactionId.Get()
	}
	if o.BillPaymentId.IsSet() {
		toSerialize["bill_payment_id"] = o.BillPaymentId.Get()
	}
	if !IsNil(o.CreatedTimestamp) {
		toSerialize["created_timestamp"] = o.CreatedTimestamp
	}
	if !IsNil(o.UpdatedTimestamp) {
		toSerialize["updated_timestamp"] = o.UpdatedTimestamp
	}
	return toSerialize, nil
}

type NullablePaymentAdjustmentItem struct {
	value *PaymentAdjustmentItem
	isSet bool
}

func (v NullablePaymentAdjustmentItem) Get() *PaymentAdjustmentItem {
	return v.value
}

func (v *NullablePaymentAdjustmentItem) Set(val *PaymentAdjustmentItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentAdjustmentItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentAdjustmentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentAdjustmentItem(val *PaymentAdjustmentItem) *NullablePaymentAdjustmentItem {
	return &NullablePaymentAdjustmentItem{value: val, isSet: true}
}

func (v NullablePaymentAdjustmentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentAdjustmentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


