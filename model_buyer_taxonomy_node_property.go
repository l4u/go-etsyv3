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

// checks if the BuyerTaxonomyNodeProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyerTaxonomyNodeProperty{}

// BuyerTaxonomyNodeProperty A product property definition.
type BuyerTaxonomyNodeProperty struct {
	// The unique numeric ID of this product property.
	PropertyId *int32 `json:"property_id,omitempty"`
	// The name string for this taxonomy node.
	Name *string `json:"name,omitempty"`
	// The human-readable product property name string.
	DisplayName *string `json:"display_name,omitempty"`
	// A list of available scales.
	Scales []BuyerTaxonomyNodePropertyScalesInner `json:"scales,omitempty"`
	// When true, listings assigned eligible taxonomy IDs require this property.
	IsRequired *bool `json:"is_required,omitempty"`
	// When true, you can use this property in listing attributes.
	SupportsAttributes *bool `json:"supports_attributes,omitempty"`
	// When true, you can use this property in listing variations.
	SupportsVariations *bool `json:"supports_variations,omitempty"`
	// When true, you can assign multiple property values to this property
	IsMultivalued *bool `json:"is_multivalued,omitempty"`
	// When true, you can assign multiple property values to this property
	MaxValuesAllowed NullableInt32 `json:"max_values_allowed,omitempty"`
	// A list of supported property value strings for this property.
	PossibleValues []BuyerTaxonomyNodePropertyPossibleValuesInner `json:"possible_values,omitempty"`
	// A list of property value strings automatically and always selected for the given property.
	SelectedValues []BuyerTaxonomyNodePropertySelectedValuesInner `json:"selected_values,omitempty"`
}

// NewBuyerTaxonomyNodeProperty instantiates a new BuyerTaxonomyNodeProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyerTaxonomyNodeProperty() *BuyerTaxonomyNodeProperty {
	this := BuyerTaxonomyNodeProperty{}
	return &this
}

// NewBuyerTaxonomyNodePropertyWithDefaults instantiates a new BuyerTaxonomyNodeProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyerTaxonomyNodePropertyWithDefaults() *BuyerTaxonomyNodeProperty {
	this := BuyerTaxonomyNodeProperty{}
	return &this
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeProperty) GetPropertyId() int32 {
	if o == nil || IsNil(o.PropertyId) {
		var ret int32
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeProperty) GetPropertyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PropertyId) {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeProperty) HasPropertyId() bool {
	if o != nil && !IsNil(o.PropertyId) {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given int32 and assigns it to the PropertyId field.
func (o *BuyerTaxonomyNodeProperty) SetPropertyId(v int32) {
	o.PropertyId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeProperty) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeProperty) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeProperty) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BuyerTaxonomyNodeProperty) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeProperty) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeProperty) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeProperty) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *BuyerTaxonomyNodeProperty) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetScales returns the Scales field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeProperty) GetScales() []BuyerTaxonomyNodePropertyScalesInner {
	if o == nil || IsNil(o.Scales) {
		var ret []BuyerTaxonomyNodePropertyScalesInner
		return ret
	}
	return o.Scales
}

// GetScalesOk returns a tuple with the Scales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeProperty) GetScalesOk() ([]BuyerTaxonomyNodePropertyScalesInner, bool) {
	if o == nil || IsNil(o.Scales) {
		return nil, false
	}
	return o.Scales, true
}

// HasScales returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeProperty) HasScales() bool {
	if o != nil && !IsNil(o.Scales) {
		return true
	}

	return false
}

// SetScales gets a reference to the given []BuyerTaxonomyNodePropertyScalesInner and assigns it to the Scales field.
func (o *BuyerTaxonomyNodeProperty) SetScales(v []BuyerTaxonomyNodePropertyScalesInner) {
	o.Scales = v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeProperty) GetIsRequired() bool {
	if o == nil || IsNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeProperty) GetIsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRequired) {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeProperty) HasIsRequired() bool {
	if o != nil && !IsNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *BuyerTaxonomyNodeProperty) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetSupportsAttributes returns the SupportsAttributes field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeProperty) GetSupportsAttributes() bool {
	if o == nil || IsNil(o.SupportsAttributes) {
		var ret bool
		return ret
	}
	return *o.SupportsAttributes
}

// GetSupportsAttributesOk returns a tuple with the SupportsAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeProperty) GetSupportsAttributesOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsAttributes) {
		return nil, false
	}
	return o.SupportsAttributes, true
}

// HasSupportsAttributes returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeProperty) HasSupportsAttributes() bool {
	if o != nil && !IsNil(o.SupportsAttributes) {
		return true
	}

	return false
}

// SetSupportsAttributes gets a reference to the given bool and assigns it to the SupportsAttributes field.
func (o *BuyerTaxonomyNodeProperty) SetSupportsAttributes(v bool) {
	o.SupportsAttributes = &v
}

// GetSupportsVariations returns the SupportsVariations field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeProperty) GetSupportsVariations() bool {
	if o == nil || IsNil(o.SupportsVariations) {
		var ret bool
		return ret
	}
	return *o.SupportsVariations
}

// GetSupportsVariationsOk returns a tuple with the SupportsVariations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeProperty) GetSupportsVariationsOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsVariations) {
		return nil, false
	}
	return o.SupportsVariations, true
}

// HasSupportsVariations returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeProperty) HasSupportsVariations() bool {
	if o != nil && !IsNil(o.SupportsVariations) {
		return true
	}

	return false
}

// SetSupportsVariations gets a reference to the given bool and assigns it to the SupportsVariations field.
func (o *BuyerTaxonomyNodeProperty) SetSupportsVariations(v bool) {
	o.SupportsVariations = &v
}

// GetIsMultivalued returns the IsMultivalued field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeProperty) GetIsMultivalued() bool {
	if o == nil || IsNil(o.IsMultivalued) {
		var ret bool
		return ret
	}
	return *o.IsMultivalued
}

// GetIsMultivaluedOk returns a tuple with the IsMultivalued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeProperty) GetIsMultivaluedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMultivalued) {
		return nil, false
	}
	return o.IsMultivalued, true
}

// HasIsMultivalued returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeProperty) HasIsMultivalued() bool {
	if o != nil && !IsNil(o.IsMultivalued) {
		return true
	}

	return false
}

// SetIsMultivalued gets a reference to the given bool and assigns it to the IsMultivalued field.
func (o *BuyerTaxonomyNodeProperty) SetIsMultivalued(v bool) {
	o.IsMultivalued = &v
}

// GetMaxValuesAllowed returns the MaxValuesAllowed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BuyerTaxonomyNodeProperty) GetMaxValuesAllowed() int32 {
	if o == nil || IsNil(o.MaxValuesAllowed.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxValuesAllowed.Get()
}

// GetMaxValuesAllowedOk returns a tuple with the MaxValuesAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuyerTaxonomyNodeProperty) GetMaxValuesAllowedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxValuesAllowed.Get(), o.MaxValuesAllowed.IsSet()
}

// HasMaxValuesAllowed returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeProperty) HasMaxValuesAllowed() bool {
	if o != nil && o.MaxValuesAllowed.IsSet() {
		return true
	}

	return false
}

// SetMaxValuesAllowed gets a reference to the given NullableInt32 and assigns it to the MaxValuesAllowed field.
func (o *BuyerTaxonomyNodeProperty) SetMaxValuesAllowed(v int32) {
	o.MaxValuesAllowed.Set(&v)
}
// SetMaxValuesAllowedNil sets the value for MaxValuesAllowed to be an explicit nil
func (o *BuyerTaxonomyNodeProperty) SetMaxValuesAllowedNil() {
	o.MaxValuesAllowed.Set(nil)
}

// UnsetMaxValuesAllowed ensures that no value is present for MaxValuesAllowed, not even an explicit nil
func (o *BuyerTaxonomyNodeProperty) UnsetMaxValuesAllowed() {
	o.MaxValuesAllowed.Unset()
}

// GetPossibleValues returns the PossibleValues field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeProperty) GetPossibleValues() []BuyerTaxonomyNodePropertyPossibleValuesInner {
	if o == nil || IsNil(o.PossibleValues) {
		var ret []BuyerTaxonomyNodePropertyPossibleValuesInner
		return ret
	}
	return o.PossibleValues
}

// GetPossibleValuesOk returns a tuple with the PossibleValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeProperty) GetPossibleValuesOk() ([]BuyerTaxonomyNodePropertyPossibleValuesInner, bool) {
	if o == nil || IsNil(o.PossibleValues) {
		return nil, false
	}
	return o.PossibleValues, true
}

// HasPossibleValues returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeProperty) HasPossibleValues() bool {
	if o != nil && !IsNil(o.PossibleValues) {
		return true
	}

	return false
}

// SetPossibleValues gets a reference to the given []BuyerTaxonomyNodePropertyPossibleValuesInner and assigns it to the PossibleValues field.
func (o *BuyerTaxonomyNodeProperty) SetPossibleValues(v []BuyerTaxonomyNodePropertyPossibleValuesInner) {
	o.PossibleValues = v
}

// GetSelectedValues returns the SelectedValues field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeProperty) GetSelectedValues() []BuyerTaxonomyNodePropertySelectedValuesInner {
	if o == nil || IsNil(o.SelectedValues) {
		var ret []BuyerTaxonomyNodePropertySelectedValuesInner
		return ret
	}
	return o.SelectedValues
}

// GetSelectedValuesOk returns a tuple with the SelectedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeProperty) GetSelectedValuesOk() ([]BuyerTaxonomyNodePropertySelectedValuesInner, bool) {
	if o == nil || IsNil(o.SelectedValues) {
		return nil, false
	}
	return o.SelectedValues, true
}

// HasSelectedValues returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeProperty) HasSelectedValues() bool {
	if o != nil && !IsNil(o.SelectedValues) {
		return true
	}

	return false
}

// SetSelectedValues gets a reference to the given []BuyerTaxonomyNodePropertySelectedValuesInner and assigns it to the SelectedValues field.
func (o *BuyerTaxonomyNodeProperty) SetSelectedValues(v []BuyerTaxonomyNodePropertySelectedValuesInner) {
	o.SelectedValues = v
}

func (o BuyerTaxonomyNodeProperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyerTaxonomyNodeProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyId) {
		toSerialize["property_id"] = o.PropertyId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Scales) {
		toSerialize["scales"] = o.Scales
	}
	if !IsNil(o.IsRequired) {
		toSerialize["is_required"] = o.IsRequired
	}
	if !IsNil(o.SupportsAttributes) {
		toSerialize["supports_attributes"] = o.SupportsAttributes
	}
	if !IsNil(o.SupportsVariations) {
		toSerialize["supports_variations"] = o.SupportsVariations
	}
	if !IsNil(o.IsMultivalued) {
		toSerialize["is_multivalued"] = o.IsMultivalued
	}
	if o.MaxValuesAllowed.IsSet() {
		toSerialize["max_values_allowed"] = o.MaxValuesAllowed.Get()
	}
	if !IsNil(o.PossibleValues) {
		toSerialize["possible_values"] = o.PossibleValues
	}
	if !IsNil(o.SelectedValues) {
		toSerialize["selected_values"] = o.SelectedValues
	}
	return toSerialize, nil
}

type NullableBuyerTaxonomyNodeProperty struct {
	value *BuyerTaxonomyNodeProperty
	isSet bool
}

func (v NullableBuyerTaxonomyNodeProperty) Get() *BuyerTaxonomyNodeProperty {
	return v.value
}

func (v *NullableBuyerTaxonomyNodeProperty) Set(val *BuyerTaxonomyNodeProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyerTaxonomyNodeProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyerTaxonomyNodeProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyerTaxonomyNodeProperty(val *BuyerTaxonomyNodeProperty) *NullableBuyerTaxonomyNodeProperty {
	return &NullableBuyerTaxonomyNodeProperty{value: val, isSet: true}
}

func (v NullableBuyerTaxonomyNodeProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyerTaxonomyNodeProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

