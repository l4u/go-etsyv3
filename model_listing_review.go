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

// checks if the ListingReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListingReview{}

// ListingReview A listing review record left by a User.
type ListingReview struct {
	// The shop's numeric ID.
	ShopId *int32 `json:"shop_id,omitempty"`
	// The ID of the ShopListing that the TransactionReview belongs to.
	ListingId *int32 `json:"listing_id,omitempty"`
	// Rating value on scale from 1 to 5
	Rating *int32 `json:"rating,omitempty"`
	// A message left by the author, explaining the feedback, if provided.
	Review NullableString `json:"review,omitempty"`
	// The language of the TransactionReview
	Language *string `json:"language,omitempty"`
	// The url to a photo provided with the feedback, dimensions fullxfull. Note: This field may be absent, depending on the buyer's privacy settings.
	ImageUrlFullxfull NullableString `json:"image_url_fullxfull,omitempty"`
	// The date and time the TransactionReview was created in epoch seconds.
	CreateTimestamp *int32 `json:"create_timestamp,omitempty"`
	// The date and time the TransactionReview was created in epoch seconds.
	CreatedTimestamp *int32 `json:"created_timestamp,omitempty"`
	// The date and time the TransactionReview was updated in epoch seconds.
	UpdateTimestamp *int32 `json:"update_timestamp,omitempty"`
	// The date and time the TransactionReview was updated in epoch seconds.
	UpdatedTimestamp *int32 `json:"updated_timestamp,omitempty"`
}

// NewListingReview instantiates a new ListingReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingReview() *ListingReview {
	this := ListingReview{}
	return &this
}

// NewListingReviewWithDefaults instantiates a new ListingReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingReviewWithDefaults() *ListingReview {
	this := ListingReview{}
	return &this
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *ListingReview) GetShopId() int32 {
	if o == nil || IsNil(o.ShopId) {
		var ret int32
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingReview) GetShopIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *ListingReview) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int32 and assigns it to the ShopId field.
func (o *ListingReview) SetShopId(v int32) {
	o.ShopId = &v
}

// GetListingId returns the ListingId field value if set, zero value otherwise.
func (o *ListingReview) GetListingId() int32 {
	if o == nil || IsNil(o.ListingId) {
		var ret int32
		return ret
	}
	return *o.ListingId
}

// GetListingIdOk returns a tuple with the ListingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingReview) GetListingIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ListingId) {
		return nil, false
	}
	return o.ListingId, true
}

// HasListingId returns a boolean if a field has been set.
func (o *ListingReview) HasListingId() bool {
	if o != nil && !IsNil(o.ListingId) {
		return true
	}

	return false
}

// SetListingId gets a reference to the given int32 and assigns it to the ListingId field.
func (o *ListingReview) SetListingId(v int32) {
	o.ListingId = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *ListingReview) GetRating() int32 {
	if o == nil || IsNil(o.Rating) {
		var ret int32
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingReview) GetRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *ListingReview) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given int32 and assigns it to the Rating field.
func (o *ListingReview) SetRating(v int32) {
	o.Rating = &v
}

// GetReview returns the Review field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingReview) GetReview() string {
	if o == nil || IsNil(o.Review.Get()) {
		var ret string
		return ret
	}
	return *o.Review.Get()
}

// GetReviewOk returns a tuple with the Review field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingReview) GetReviewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Review.Get(), o.Review.IsSet()
}

// HasReview returns a boolean if a field has been set.
func (o *ListingReview) HasReview() bool {
	if o != nil && o.Review.IsSet() {
		return true
	}

	return false
}

// SetReview gets a reference to the given NullableString and assigns it to the Review field.
func (o *ListingReview) SetReview(v string) {
	o.Review.Set(&v)
}
// SetReviewNil sets the value for Review to be an explicit nil
func (o *ListingReview) SetReviewNil() {
	o.Review.Set(nil)
}

// UnsetReview ensures that no value is present for Review, not even an explicit nil
func (o *ListingReview) UnsetReview() {
	o.Review.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ListingReview) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingReview) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ListingReview) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ListingReview) SetLanguage(v string) {
	o.Language = &v
}

// GetImageUrlFullxfull returns the ImageUrlFullxfull field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingReview) GetImageUrlFullxfull() string {
	if o == nil || IsNil(o.ImageUrlFullxfull.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrlFullxfull.Get()
}

// GetImageUrlFullxfullOk returns a tuple with the ImageUrlFullxfull field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingReview) GetImageUrlFullxfullOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrlFullxfull.Get(), o.ImageUrlFullxfull.IsSet()
}

// HasImageUrlFullxfull returns a boolean if a field has been set.
func (o *ListingReview) HasImageUrlFullxfull() bool {
	if o != nil && o.ImageUrlFullxfull.IsSet() {
		return true
	}

	return false
}

// SetImageUrlFullxfull gets a reference to the given NullableString and assigns it to the ImageUrlFullxfull field.
func (o *ListingReview) SetImageUrlFullxfull(v string) {
	o.ImageUrlFullxfull.Set(&v)
}
// SetImageUrlFullxfullNil sets the value for ImageUrlFullxfull to be an explicit nil
func (o *ListingReview) SetImageUrlFullxfullNil() {
	o.ImageUrlFullxfull.Set(nil)
}

// UnsetImageUrlFullxfull ensures that no value is present for ImageUrlFullxfull, not even an explicit nil
func (o *ListingReview) UnsetImageUrlFullxfull() {
	o.ImageUrlFullxfull.Unset()
}

// GetCreateTimestamp returns the CreateTimestamp field value if set, zero value otherwise.
func (o *ListingReview) GetCreateTimestamp() int32 {
	if o == nil || IsNil(o.CreateTimestamp) {
		var ret int32
		return ret
	}
	return *o.CreateTimestamp
}

// GetCreateTimestampOk returns a tuple with the CreateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingReview) GetCreateTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.CreateTimestamp) {
		return nil, false
	}
	return o.CreateTimestamp, true
}

// HasCreateTimestamp returns a boolean if a field has been set.
func (o *ListingReview) HasCreateTimestamp() bool {
	if o != nil && !IsNil(o.CreateTimestamp) {
		return true
	}

	return false
}

// SetCreateTimestamp gets a reference to the given int32 and assigns it to the CreateTimestamp field.
func (o *ListingReview) SetCreateTimestamp(v int32) {
	o.CreateTimestamp = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value if set, zero value otherwise.
func (o *ListingReview) GetCreatedTimestamp() int32 {
	if o == nil || IsNil(o.CreatedTimestamp) {
		var ret int32
		return ret
	}
	return *o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingReview) GetCreatedTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedTimestamp) {
		return nil, false
	}
	return o.CreatedTimestamp, true
}

// HasCreatedTimestamp returns a boolean if a field has been set.
func (o *ListingReview) HasCreatedTimestamp() bool {
	if o != nil && !IsNil(o.CreatedTimestamp) {
		return true
	}

	return false
}

// SetCreatedTimestamp gets a reference to the given int32 and assigns it to the CreatedTimestamp field.
func (o *ListingReview) SetCreatedTimestamp(v int32) {
	o.CreatedTimestamp = &v
}

// GetUpdateTimestamp returns the UpdateTimestamp field value if set, zero value otherwise.
func (o *ListingReview) GetUpdateTimestamp() int32 {
	if o == nil || IsNil(o.UpdateTimestamp) {
		var ret int32
		return ret
	}
	return *o.UpdateTimestamp
}

// GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingReview) GetUpdateTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdateTimestamp) {
		return nil, false
	}
	return o.UpdateTimestamp, true
}

// HasUpdateTimestamp returns a boolean if a field has been set.
func (o *ListingReview) HasUpdateTimestamp() bool {
	if o != nil && !IsNil(o.UpdateTimestamp) {
		return true
	}

	return false
}

// SetUpdateTimestamp gets a reference to the given int32 and assigns it to the UpdateTimestamp field.
func (o *ListingReview) SetUpdateTimestamp(v int32) {
	o.UpdateTimestamp = &v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value if set, zero value otherwise.
func (o *ListingReview) GetUpdatedTimestamp() int32 {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		var ret int32
		return ret
	}
	return *o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingReview) GetUpdatedTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		return nil, false
	}
	return o.UpdatedTimestamp, true
}

// HasUpdatedTimestamp returns a boolean if a field has been set.
func (o *ListingReview) HasUpdatedTimestamp() bool {
	if o != nil && !IsNil(o.UpdatedTimestamp) {
		return true
	}

	return false
}

// SetUpdatedTimestamp gets a reference to the given int32 and assigns it to the UpdatedTimestamp field.
func (o *ListingReview) SetUpdatedTimestamp(v int32) {
	o.UpdatedTimestamp = &v
}

func (o ListingReview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListingReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.ListingId) {
		toSerialize["listing_id"] = o.ListingId
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if o.Review.IsSet() {
		toSerialize["review"] = o.Review.Get()
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if o.ImageUrlFullxfull.IsSet() {
		toSerialize["image_url_fullxfull"] = o.ImageUrlFullxfull.Get()
	}
	if !IsNil(o.CreateTimestamp) {
		toSerialize["create_timestamp"] = o.CreateTimestamp
	}
	if !IsNil(o.CreatedTimestamp) {
		toSerialize["created_timestamp"] = o.CreatedTimestamp
	}
	if !IsNil(o.UpdateTimestamp) {
		toSerialize["update_timestamp"] = o.UpdateTimestamp
	}
	if !IsNil(o.UpdatedTimestamp) {
		toSerialize["updated_timestamp"] = o.UpdatedTimestamp
	}
	return toSerialize, nil
}

type NullableListingReview struct {
	value *ListingReview
	isSet bool
}

func (v NullableListingReview) Get() *ListingReview {
	return v.value
}

func (v *NullableListingReview) Set(val *ListingReview) {
	v.value = val
	v.isSet = true
}

func (v NullableListingReview) IsSet() bool {
	return v.isSet
}

func (v *NullableListingReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingReview(val *ListingReview) *NullableListingReview {
	return &NullableListingReview{value: val, isSet: true}
}

func (v NullableListingReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListingReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


