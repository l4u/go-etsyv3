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

// checks if the ShopReceiptRefundsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShopReceiptRefundsInner{}

// ShopReceiptRefundsInner Refunds for a given receipt.
type ShopReceiptRefundsInner struct {
	Amount *ShopRefundAmount `json:"amount,omitempty"`
	// The date & time of the refund, in epoch seconds.
	CreatedTimestamp *int32 `json:"created_timestamp,omitempty"`
	// The reason string given for the refund.
	Reason NullableString `json:"reason,omitempty"`
	// The note string created by the refund issuer.
	NoteFromIssuer NullableString `json:"note_from_issuer,omitempty"`
	// The status indication string for the refund.
	Status NullableString `json:"status,omitempty"`
}

// NewShopReceiptRefundsInner instantiates a new ShopReceiptRefundsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShopReceiptRefundsInner() *ShopReceiptRefundsInner {
	this := ShopReceiptRefundsInner{}
	return &this
}

// NewShopReceiptRefundsInnerWithDefaults instantiates a new ShopReceiptRefundsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShopReceiptRefundsInnerWithDefaults() *ShopReceiptRefundsInner {
	this := ShopReceiptRefundsInner{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ShopReceiptRefundsInner) GetAmount() ShopRefundAmount {
	if o == nil || IsNil(o.Amount) {
		var ret ShopRefundAmount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopReceiptRefundsInner) GetAmountOk() (*ShopRefundAmount, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ShopReceiptRefundsInner) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given ShopRefundAmount and assigns it to the Amount field.
func (o *ShopReceiptRefundsInner) SetAmount(v ShopRefundAmount) {
	o.Amount = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value if set, zero value otherwise.
func (o *ShopReceiptRefundsInner) GetCreatedTimestamp() int32 {
	if o == nil || IsNil(o.CreatedTimestamp) {
		var ret int32
		return ret
	}
	return *o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopReceiptRefundsInner) GetCreatedTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedTimestamp) {
		return nil, false
	}
	return o.CreatedTimestamp, true
}

// HasCreatedTimestamp returns a boolean if a field has been set.
func (o *ShopReceiptRefundsInner) HasCreatedTimestamp() bool {
	if o != nil && !IsNil(o.CreatedTimestamp) {
		return true
	}

	return false
}

// SetCreatedTimestamp gets a reference to the given int32 and assigns it to the CreatedTimestamp field.
func (o *ShopReceiptRefundsInner) SetCreatedTimestamp(v int32) {
	o.CreatedTimestamp = &v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShopReceiptRefundsInner) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShopReceiptRefundsInner) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *ShopReceiptRefundsInner) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *ShopReceiptRefundsInner) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *ShopReceiptRefundsInner) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *ShopReceiptRefundsInner) UnsetReason() {
	o.Reason.Unset()
}

// GetNoteFromIssuer returns the NoteFromIssuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShopReceiptRefundsInner) GetNoteFromIssuer() string {
	if o == nil || IsNil(o.NoteFromIssuer.Get()) {
		var ret string
		return ret
	}
	return *o.NoteFromIssuer.Get()
}

// GetNoteFromIssuerOk returns a tuple with the NoteFromIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShopReceiptRefundsInner) GetNoteFromIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NoteFromIssuer.Get(), o.NoteFromIssuer.IsSet()
}

// HasNoteFromIssuer returns a boolean if a field has been set.
func (o *ShopReceiptRefundsInner) HasNoteFromIssuer() bool {
	if o != nil && o.NoteFromIssuer.IsSet() {
		return true
	}

	return false
}

// SetNoteFromIssuer gets a reference to the given NullableString and assigns it to the NoteFromIssuer field.
func (o *ShopReceiptRefundsInner) SetNoteFromIssuer(v string) {
	o.NoteFromIssuer.Set(&v)
}
// SetNoteFromIssuerNil sets the value for NoteFromIssuer to be an explicit nil
func (o *ShopReceiptRefundsInner) SetNoteFromIssuerNil() {
	o.NoteFromIssuer.Set(nil)
}

// UnsetNoteFromIssuer ensures that no value is present for NoteFromIssuer, not even an explicit nil
func (o *ShopReceiptRefundsInner) UnsetNoteFromIssuer() {
	o.NoteFromIssuer.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShopReceiptRefundsInner) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShopReceiptRefundsInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ShopReceiptRefundsInner) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *ShopReceiptRefundsInner) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ShopReceiptRefundsInner) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ShopReceiptRefundsInner) UnsetStatus() {
	o.Status.Unset()
}

func (o ShopReceiptRefundsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShopReceiptRefundsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CreatedTimestamp) {
		toSerialize["created_timestamp"] = o.CreatedTimestamp
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.NoteFromIssuer.IsSet() {
		toSerialize["note_from_issuer"] = o.NoteFromIssuer.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	return toSerialize, nil
}

type NullableShopReceiptRefundsInner struct {
	value *ShopReceiptRefundsInner
	isSet bool
}

func (v NullableShopReceiptRefundsInner) Get() *ShopReceiptRefundsInner {
	return v.value
}

func (v *NullableShopReceiptRefundsInner) Set(val *ShopReceiptRefundsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableShopReceiptRefundsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableShopReceiptRefundsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShopReceiptRefundsInner(val *ShopReceiptRefundsInner) *NullableShopReceiptRefundsInner {
	return &NullableShopReceiptRefundsInner{value: val, isSet: true}
}

func (v NullableShopReceiptRefundsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShopReceiptRefundsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


