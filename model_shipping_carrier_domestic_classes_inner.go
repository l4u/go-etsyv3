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

// checks if the ShippingCarrierDomesticClassesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingCarrierDomesticClassesInner{}

// ShippingCarrierDomesticClassesInner Set of domestic mail classes of this shipping carrier.
type ShippingCarrierDomesticClassesInner struct {
	// The unique identifier of this mail class.
	MailClassKey *string `json:"mail_class_key,omitempty"`
	// The name of this mail class.
	Name *string `json:"name,omitempty"`
}

// NewShippingCarrierDomesticClassesInner instantiates a new ShippingCarrierDomesticClassesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingCarrierDomesticClassesInner() *ShippingCarrierDomesticClassesInner {
	this := ShippingCarrierDomesticClassesInner{}
	return &this
}

// NewShippingCarrierDomesticClassesInnerWithDefaults instantiates a new ShippingCarrierDomesticClassesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingCarrierDomesticClassesInnerWithDefaults() *ShippingCarrierDomesticClassesInner {
	this := ShippingCarrierDomesticClassesInner{}
	return &this
}

// GetMailClassKey returns the MailClassKey field value if set, zero value otherwise.
func (o *ShippingCarrierDomesticClassesInner) GetMailClassKey() string {
	if o == nil || IsNil(o.MailClassKey) {
		var ret string
		return ret
	}
	return *o.MailClassKey
}

// GetMailClassKeyOk returns a tuple with the MailClassKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingCarrierDomesticClassesInner) GetMailClassKeyOk() (*string, bool) {
	if o == nil || IsNil(o.MailClassKey) {
		return nil, false
	}
	return o.MailClassKey, true
}

// HasMailClassKey returns a boolean if a field has been set.
func (o *ShippingCarrierDomesticClassesInner) HasMailClassKey() bool {
	if o != nil && !IsNil(o.MailClassKey) {
		return true
	}

	return false
}

// SetMailClassKey gets a reference to the given string and assigns it to the MailClassKey field.
func (o *ShippingCarrierDomesticClassesInner) SetMailClassKey(v string) {
	o.MailClassKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShippingCarrierDomesticClassesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingCarrierDomesticClassesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShippingCarrierDomesticClassesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShippingCarrierDomesticClassesInner) SetName(v string) {
	o.Name = &v
}

func (o ShippingCarrierDomesticClassesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingCarrierDomesticClassesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MailClassKey) {
		toSerialize["mail_class_key"] = o.MailClassKey
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableShippingCarrierDomesticClassesInner struct {
	value *ShippingCarrierDomesticClassesInner
	isSet bool
}

func (v NullableShippingCarrierDomesticClassesInner) Get() *ShippingCarrierDomesticClassesInner {
	return v.value
}

func (v *NullableShippingCarrierDomesticClassesInner) Set(val *ShippingCarrierDomesticClassesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingCarrierDomesticClassesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingCarrierDomesticClassesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingCarrierDomesticClassesInner(val *ShippingCarrierDomesticClassesInner) *NullableShippingCarrierDomesticClassesInner {
	return &NullableShippingCarrierDomesticClassesInner{value: val, isSet: true}
}

func (v NullableShippingCarrierDomesticClassesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingCarrierDomesticClassesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


