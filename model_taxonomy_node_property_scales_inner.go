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

// checks if the TaxonomyNodePropertyScalesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxonomyNodePropertyScalesInner{}

// TaxonomyNodePropertyScalesInner A list of available scales.
type TaxonomyNodePropertyScalesInner struct {
	// The unique numeric ID of a scale.
	ScaleId *int32 `json:"scale_id,omitempty"`
	// The name string for a scale.
	DisplayName *string `json:"display_name,omitempty"`
	// The description string for a scale.
	Description *string `json:"description,omitempty"`
}

// NewTaxonomyNodePropertyScalesInner instantiates a new TaxonomyNodePropertyScalesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxonomyNodePropertyScalesInner() *TaxonomyNodePropertyScalesInner {
	this := TaxonomyNodePropertyScalesInner{}
	return &this
}

// NewTaxonomyNodePropertyScalesInnerWithDefaults instantiates a new TaxonomyNodePropertyScalesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxonomyNodePropertyScalesInnerWithDefaults() *TaxonomyNodePropertyScalesInner {
	this := TaxonomyNodePropertyScalesInner{}
	return &this
}

// GetScaleId returns the ScaleId field value if set, zero value otherwise.
func (o *TaxonomyNodePropertyScalesInner) GetScaleId() int32 {
	if o == nil || IsNil(o.ScaleId) {
		var ret int32
		return ret
	}
	return *o.ScaleId
}

// GetScaleIdOk returns a tuple with the ScaleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxonomyNodePropertyScalesInner) GetScaleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ScaleId) {
		return nil, false
	}
	return o.ScaleId, true
}

// HasScaleId returns a boolean if a field has been set.
func (o *TaxonomyNodePropertyScalesInner) HasScaleId() bool {
	if o != nil && !IsNil(o.ScaleId) {
		return true
	}

	return false
}

// SetScaleId gets a reference to the given int32 and assigns it to the ScaleId field.
func (o *TaxonomyNodePropertyScalesInner) SetScaleId(v int32) {
	o.ScaleId = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *TaxonomyNodePropertyScalesInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxonomyNodePropertyScalesInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TaxonomyNodePropertyScalesInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *TaxonomyNodePropertyScalesInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TaxonomyNodePropertyScalesInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxonomyNodePropertyScalesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TaxonomyNodePropertyScalesInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TaxonomyNodePropertyScalesInner) SetDescription(v string) {
	o.Description = &v
}

func (o TaxonomyNodePropertyScalesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxonomyNodePropertyScalesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScaleId) {
		toSerialize["scale_id"] = o.ScaleId
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableTaxonomyNodePropertyScalesInner struct {
	value *TaxonomyNodePropertyScalesInner
	isSet bool
}

func (v NullableTaxonomyNodePropertyScalesInner) Get() *TaxonomyNodePropertyScalesInner {
	return v.value
}

func (v *NullableTaxonomyNodePropertyScalesInner) Set(val *TaxonomyNodePropertyScalesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxonomyNodePropertyScalesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxonomyNodePropertyScalesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxonomyNodePropertyScalesInner(val *TaxonomyNodePropertyScalesInner) *NullableTaxonomyNodePropertyScalesInner {
	return &NullableTaxonomyNodePropertyScalesInner{value: val, isSet: true}
}

func (v NullableTaxonomyNodePropertyScalesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxonomyNodePropertyScalesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


