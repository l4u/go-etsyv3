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

// checks if the ShopSectionsResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShopSectionsResultsInner{}

// ShopSectionsResultsInner The list of requested resources.
type ShopSectionsResultsInner struct {
	// The numeric ID of a section in a specific Etsy shop.
	ShopSectionId *int32 `json:"shop_section_id,omitempty"`
	// The title string for a shop section.
	Title *string `json:"title,omitempty"`
	// The positive non-zero numeric position of this section in the section display order for a shop, with rank 1 sections appearing first.
	Rank *int32 `json:"rank,omitempty"`
	// The numeric ID of the [user](/documentation/reference#tag/User) who owns this shop section.
	UserId *int32 `json:"user_id,omitempty"`
	// The number of active listings in one section of a specific Etsy shop.
	ActiveListingCount *int32 `json:"active_listing_count,omitempty"`
}

// NewShopSectionsResultsInner instantiates a new ShopSectionsResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShopSectionsResultsInner() *ShopSectionsResultsInner {
	this := ShopSectionsResultsInner{}
	return &this
}

// NewShopSectionsResultsInnerWithDefaults instantiates a new ShopSectionsResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShopSectionsResultsInnerWithDefaults() *ShopSectionsResultsInner {
	this := ShopSectionsResultsInner{}
	return &this
}

// GetShopSectionId returns the ShopSectionId field value if set, zero value otherwise.
func (o *ShopSectionsResultsInner) GetShopSectionId() int32 {
	if o == nil || IsNil(o.ShopSectionId) {
		var ret int32
		return ret
	}
	return *o.ShopSectionId
}

// GetShopSectionIdOk returns a tuple with the ShopSectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopSectionsResultsInner) GetShopSectionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ShopSectionId) {
		return nil, false
	}
	return o.ShopSectionId, true
}

// HasShopSectionId returns a boolean if a field has been set.
func (o *ShopSectionsResultsInner) HasShopSectionId() bool {
	if o != nil && !IsNil(o.ShopSectionId) {
		return true
	}

	return false
}

// SetShopSectionId gets a reference to the given int32 and assigns it to the ShopSectionId field.
func (o *ShopSectionsResultsInner) SetShopSectionId(v int32) {
	o.ShopSectionId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ShopSectionsResultsInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopSectionsResultsInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ShopSectionsResultsInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ShopSectionsResultsInner) SetTitle(v string) {
	o.Title = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *ShopSectionsResultsInner) GetRank() int32 {
	if o == nil || IsNil(o.Rank) {
		var ret int32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopSectionsResultsInner) GetRankOk() (*int32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *ShopSectionsResultsInner) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given int32 and assigns it to the Rank field.
func (o *ShopSectionsResultsInner) SetRank(v int32) {
	o.Rank = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ShopSectionsResultsInner) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopSectionsResultsInner) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ShopSectionsResultsInner) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *ShopSectionsResultsInner) SetUserId(v int32) {
	o.UserId = &v
}

// GetActiveListingCount returns the ActiveListingCount field value if set, zero value otherwise.
func (o *ShopSectionsResultsInner) GetActiveListingCount() int32 {
	if o == nil || IsNil(o.ActiveListingCount) {
		var ret int32
		return ret
	}
	return *o.ActiveListingCount
}

// GetActiveListingCountOk returns a tuple with the ActiveListingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopSectionsResultsInner) GetActiveListingCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveListingCount) {
		return nil, false
	}
	return o.ActiveListingCount, true
}

// HasActiveListingCount returns a boolean if a field has been set.
func (o *ShopSectionsResultsInner) HasActiveListingCount() bool {
	if o != nil && !IsNil(o.ActiveListingCount) {
		return true
	}

	return false
}

// SetActiveListingCount gets a reference to the given int32 and assigns it to the ActiveListingCount field.
func (o *ShopSectionsResultsInner) SetActiveListingCount(v int32) {
	o.ActiveListingCount = &v
}

func (o ShopSectionsResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShopSectionsResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShopSectionId) {
		toSerialize["shop_section_id"] = o.ShopSectionId
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.ActiveListingCount) {
		toSerialize["active_listing_count"] = o.ActiveListingCount
	}
	return toSerialize, nil
}

type NullableShopSectionsResultsInner struct {
	value *ShopSectionsResultsInner
	isSet bool
}

func (v NullableShopSectionsResultsInner) Get() *ShopSectionsResultsInner {
	return v.value
}

func (v *NullableShopSectionsResultsInner) Set(val *ShopSectionsResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableShopSectionsResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableShopSectionsResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShopSectionsResultsInner(val *ShopSectionsResultsInner) *NullableShopSectionsResultsInner {
	return &NullableShopSectionsResultsInner{value: val, isSet: true}
}

func (v NullableShopSectionsResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShopSectionsResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


