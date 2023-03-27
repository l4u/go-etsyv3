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

// checks if the UserAddressesResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAddressesResultsInner{}

// UserAddressesResultsInner An array of UserAddress resources.
type UserAddressesResultsInner struct {
	// The numeric ID of the user's address.
	UserAddressId *int32 `json:"user_address_id,omitempty"`
	// The user's numeric ID.
	UserId *int32 `json:"user_id,omitempty"`
	// The user's name for this address.
	Name *string `json:"name,omitempty"`
	// The first line of the user's address.
	FirstLine *string `json:"first_line,omitempty"`
	// The second line of the user's address.
	SecondLine NullableString `json:"second_line,omitempty"`
	// The city field of the user's address.
	City *string `json:"city,omitempty"`
	// The state field of the user's address.
	State NullableString `json:"state,omitempty"`
	// The zip code field of the user's address.
	Zip *string `json:"zip,omitempty"`
	// The ISO code of the country in this address.
	IsoCountryCode NullableString `json:"iso_country_code,omitempty"`
	// The name of the user's country.
	CountryName NullableString `json:"country_name,omitempty"`
	// Is this the user's default shipping address.
	IsDefaultShippingAddress *bool `json:"is_default_shipping_address,omitempty"`
}

// NewUserAddressesResultsInner instantiates a new UserAddressesResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAddressesResultsInner() *UserAddressesResultsInner {
	this := UserAddressesResultsInner{}
	return &this
}

// NewUserAddressesResultsInnerWithDefaults instantiates a new UserAddressesResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAddressesResultsInnerWithDefaults() *UserAddressesResultsInner {
	this := UserAddressesResultsInner{}
	return &this
}

// GetUserAddressId returns the UserAddressId field value if set, zero value otherwise.
func (o *UserAddressesResultsInner) GetUserAddressId() int32 {
	if o == nil || IsNil(o.UserAddressId) {
		var ret int32
		return ret
	}
	return *o.UserAddressId
}

// GetUserAddressIdOk returns a tuple with the UserAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAddressesResultsInner) GetUserAddressIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserAddressId) {
		return nil, false
	}
	return o.UserAddressId, true
}

// HasUserAddressId returns a boolean if a field has been set.
func (o *UserAddressesResultsInner) HasUserAddressId() bool {
	if o != nil && !IsNil(o.UserAddressId) {
		return true
	}

	return false
}

// SetUserAddressId gets a reference to the given int32 and assigns it to the UserAddressId field.
func (o *UserAddressesResultsInner) SetUserAddressId(v int32) {
	o.UserAddressId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UserAddressesResultsInner) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAddressesResultsInner) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UserAddressesResultsInner) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *UserAddressesResultsInner) SetUserId(v int32) {
	o.UserId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserAddressesResultsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAddressesResultsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserAddressesResultsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserAddressesResultsInner) SetName(v string) {
	o.Name = &v
}

// GetFirstLine returns the FirstLine field value if set, zero value otherwise.
func (o *UserAddressesResultsInner) GetFirstLine() string {
	if o == nil || IsNil(o.FirstLine) {
		var ret string
		return ret
	}
	return *o.FirstLine
}

// GetFirstLineOk returns a tuple with the FirstLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAddressesResultsInner) GetFirstLineOk() (*string, bool) {
	if o == nil || IsNil(o.FirstLine) {
		return nil, false
	}
	return o.FirstLine, true
}

// HasFirstLine returns a boolean if a field has been set.
func (o *UserAddressesResultsInner) HasFirstLine() bool {
	if o != nil && !IsNil(o.FirstLine) {
		return true
	}

	return false
}

// SetFirstLine gets a reference to the given string and assigns it to the FirstLine field.
func (o *UserAddressesResultsInner) SetFirstLine(v string) {
	o.FirstLine = &v
}

// GetSecondLine returns the SecondLine field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAddressesResultsInner) GetSecondLine() string {
	if o == nil || IsNil(o.SecondLine.Get()) {
		var ret string
		return ret
	}
	return *o.SecondLine.Get()
}

// GetSecondLineOk returns a tuple with the SecondLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAddressesResultsInner) GetSecondLineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecondLine.Get(), o.SecondLine.IsSet()
}

// HasSecondLine returns a boolean if a field has been set.
func (o *UserAddressesResultsInner) HasSecondLine() bool {
	if o != nil && o.SecondLine.IsSet() {
		return true
	}

	return false
}

// SetSecondLine gets a reference to the given NullableString and assigns it to the SecondLine field.
func (o *UserAddressesResultsInner) SetSecondLine(v string) {
	o.SecondLine.Set(&v)
}
// SetSecondLineNil sets the value for SecondLine to be an explicit nil
func (o *UserAddressesResultsInner) SetSecondLineNil() {
	o.SecondLine.Set(nil)
}

// UnsetSecondLine ensures that no value is present for SecondLine, not even an explicit nil
func (o *UserAddressesResultsInner) UnsetSecondLine() {
	o.SecondLine.Unset()
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *UserAddressesResultsInner) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAddressesResultsInner) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *UserAddressesResultsInner) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *UserAddressesResultsInner) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAddressesResultsInner) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAddressesResultsInner) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *UserAddressesResultsInner) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *UserAddressesResultsInner) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *UserAddressesResultsInner) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *UserAddressesResultsInner) UnsetState() {
	o.State.Unset()
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *UserAddressesResultsInner) GetZip() string {
	if o == nil || IsNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAddressesResultsInner) GetZipOk() (*string, bool) {
	if o == nil || IsNil(o.Zip) {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *UserAddressesResultsInner) HasZip() bool {
	if o != nil && !IsNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *UserAddressesResultsInner) SetZip(v string) {
	o.Zip = &v
}

// GetIsoCountryCode returns the IsoCountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAddressesResultsInner) GetIsoCountryCode() string {
	if o == nil || IsNil(o.IsoCountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.IsoCountryCode.Get()
}

// GetIsoCountryCodeOk returns a tuple with the IsoCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAddressesResultsInner) GetIsoCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsoCountryCode.Get(), o.IsoCountryCode.IsSet()
}

// HasIsoCountryCode returns a boolean if a field has been set.
func (o *UserAddressesResultsInner) HasIsoCountryCode() bool {
	if o != nil && o.IsoCountryCode.IsSet() {
		return true
	}

	return false
}

// SetIsoCountryCode gets a reference to the given NullableString and assigns it to the IsoCountryCode field.
func (o *UserAddressesResultsInner) SetIsoCountryCode(v string) {
	o.IsoCountryCode.Set(&v)
}
// SetIsoCountryCodeNil sets the value for IsoCountryCode to be an explicit nil
func (o *UserAddressesResultsInner) SetIsoCountryCodeNil() {
	o.IsoCountryCode.Set(nil)
}

// UnsetIsoCountryCode ensures that no value is present for IsoCountryCode, not even an explicit nil
func (o *UserAddressesResultsInner) UnsetIsoCountryCode() {
	o.IsoCountryCode.Unset()
}

// GetCountryName returns the CountryName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAddressesResultsInner) GetCountryName() string {
	if o == nil || IsNil(o.CountryName.Get()) {
		var ret string
		return ret
	}
	return *o.CountryName.Get()
}

// GetCountryNameOk returns a tuple with the CountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAddressesResultsInner) GetCountryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryName.Get(), o.CountryName.IsSet()
}

// HasCountryName returns a boolean if a field has been set.
func (o *UserAddressesResultsInner) HasCountryName() bool {
	if o != nil && o.CountryName.IsSet() {
		return true
	}

	return false
}

// SetCountryName gets a reference to the given NullableString and assigns it to the CountryName field.
func (o *UserAddressesResultsInner) SetCountryName(v string) {
	o.CountryName.Set(&v)
}
// SetCountryNameNil sets the value for CountryName to be an explicit nil
func (o *UserAddressesResultsInner) SetCountryNameNil() {
	o.CountryName.Set(nil)
}

// UnsetCountryName ensures that no value is present for CountryName, not even an explicit nil
func (o *UserAddressesResultsInner) UnsetCountryName() {
	o.CountryName.Unset()
}

// GetIsDefaultShippingAddress returns the IsDefaultShippingAddress field value if set, zero value otherwise.
func (o *UserAddressesResultsInner) GetIsDefaultShippingAddress() bool {
	if o == nil || IsNil(o.IsDefaultShippingAddress) {
		var ret bool
		return ret
	}
	return *o.IsDefaultShippingAddress
}

// GetIsDefaultShippingAddressOk returns a tuple with the IsDefaultShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAddressesResultsInner) GetIsDefaultShippingAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultShippingAddress) {
		return nil, false
	}
	return o.IsDefaultShippingAddress, true
}

// HasIsDefaultShippingAddress returns a boolean if a field has been set.
func (o *UserAddressesResultsInner) HasIsDefaultShippingAddress() bool {
	if o != nil && !IsNil(o.IsDefaultShippingAddress) {
		return true
	}

	return false
}

// SetIsDefaultShippingAddress gets a reference to the given bool and assigns it to the IsDefaultShippingAddress field.
func (o *UserAddressesResultsInner) SetIsDefaultShippingAddress(v bool) {
	o.IsDefaultShippingAddress = &v
}

func (o UserAddressesResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAddressesResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserAddressId) {
		toSerialize["user_address_id"] = o.UserAddressId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FirstLine) {
		toSerialize["first_line"] = o.FirstLine
	}
	if o.SecondLine.IsSet() {
		toSerialize["second_line"] = o.SecondLine.Get()
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if !IsNil(o.Zip) {
		toSerialize["zip"] = o.Zip
	}
	if o.IsoCountryCode.IsSet() {
		toSerialize["iso_country_code"] = o.IsoCountryCode.Get()
	}
	if o.CountryName.IsSet() {
		toSerialize["country_name"] = o.CountryName.Get()
	}
	if !IsNil(o.IsDefaultShippingAddress) {
		toSerialize["is_default_shipping_address"] = o.IsDefaultShippingAddress
	}
	return toSerialize, nil
}

type NullableUserAddressesResultsInner struct {
	value *UserAddressesResultsInner
	isSet bool
}

func (v NullableUserAddressesResultsInner) Get() *UserAddressesResultsInner {
	return v.value
}

func (v *NullableUserAddressesResultsInner) Set(val *UserAddressesResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAddressesResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAddressesResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAddressesResultsInner(val *UserAddressesResultsInner) *NullableUserAddressesResultsInner {
	return &NullableUserAddressesResultsInner{value: val, isSet: true}
}

func (v NullableUserAddressesResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAddressesResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

