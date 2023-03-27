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

// checks if the BuyerTaxonomyNodeChildrenInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyerTaxonomyNodeChildrenInner{}

// BuyerTaxonomyNodeChildrenInner An array of taxonomy nodes for all the direct children of this taxonomy node in the seller taxonomy tree.
type BuyerTaxonomyNodeChildrenInner struct {
	// The unique numeric ID of an Etsy taxonomy node, which is a metadata category for listings organized into the seller taxonomy hierarchy tree. For example, the \\\"shoes\\\" taxonomy node (ID: 1429, level: 1) is higher in the hierarchy than \\\"girls' shoes\\\" (ID: 1440, level: 2). The taxonomy nodes assigned to a listing support access to specific standardized product scales and properties. For example, listings assigned the taxonomy nodes \\\"shoes\\\" or \\\"girls' shoes\\\" support access to the \\\"EU\\\" shoe size scale with its associated property names and IDs for EU shoe sizes, such as property `value_id`:\\\"1394\\\", and `name`:\\\"38\\\".
	Id *int32 `json:"id,omitempty"`
	// The integer depth of this taxonomy node in the seller taxonomy tree, with roots at level 0.
	Level *int32 `json:"level,omitempty"`
	// The name string for this taxonomy node.
	Name *string `json:"name,omitempty"`
	// The numeric taxonomy ID of the parent of this node.
	ParentId NullableInt32 `json:"parent_id,omitempty"`
	// An array of taxonomy nodes for all the direct children of this taxonomy node in the seller taxonomy tree.
	Children []BuyerTaxonomyNodeChildrenInner `json:"children,omitempty"`
	// An array of `taxonomy_id`s including this node and all of its direct parents in the seller taxonomy tree up to a root node. They are listed in order from root to leaf.
	FullPathTaxonomyIds []int32 `json:"full_path_taxonomy_ids,omitempty"`
}

// NewBuyerTaxonomyNodeChildrenInner instantiates a new BuyerTaxonomyNodeChildrenInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyerTaxonomyNodeChildrenInner() *BuyerTaxonomyNodeChildrenInner {
	this := BuyerTaxonomyNodeChildrenInner{}
	return &this
}

// NewBuyerTaxonomyNodeChildrenInnerWithDefaults instantiates a new BuyerTaxonomyNodeChildrenInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyerTaxonomyNodeChildrenInnerWithDefaults() *BuyerTaxonomyNodeChildrenInner {
	this := BuyerTaxonomyNodeChildrenInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeChildrenInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeChildrenInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeChildrenInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BuyerTaxonomyNodeChildrenInner) SetId(v int32) {
	o.Id = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeChildrenInner) GetLevel() int32 {
	if o == nil || IsNil(o.Level) {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeChildrenInner) GetLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeChildrenInner) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *BuyerTaxonomyNodeChildrenInner) SetLevel(v int32) {
	o.Level = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeChildrenInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeChildrenInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeChildrenInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BuyerTaxonomyNodeChildrenInner) SetName(v string) {
	o.Name = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BuyerTaxonomyNodeChildrenInner) GetParentId() int32 {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret int32
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuyerTaxonomyNodeChildrenInner) GetParentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeChildrenInner) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableInt32 and assigns it to the ParentId field.
func (o *BuyerTaxonomyNodeChildrenInner) SetParentId(v int32) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *BuyerTaxonomyNodeChildrenInner) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *BuyerTaxonomyNodeChildrenInner) UnsetParentId() {
	o.ParentId.Unset()
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeChildrenInner) GetChildren() []BuyerTaxonomyNodeChildrenInner {
	if o == nil || IsNil(o.Children) {
		var ret []BuyerTaxonomyNodeChildrenInner
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeChildrenInner) GetChildrenOk() ([]BuyerTaxonomyNodeChildrenInner, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeChildrenInner) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []BuyerTaxonomyNodeChildrenInner and assigns it to the Children field.
func (o *BuyerTaxonomyNodeChildrenInner) SetChildren(v []BuyerTaxonomyNodeChildrenInner) {
	o.Children = v
}

// GetFullPathTaxonomyIds returns the FullPathTaxonomyIds field value if set, zero value otherwise.
func (o *BuyerTaxonomyNodeChildrenInner) GetFullPathTaxonomyIds() []int32 {
	if o == nil || IsNil(o.FullPathTaxonomyIds) {
		var ret []int32
		return ret
	}
	return o.FullPathTaxonomyIds
}

// GetFullPathTaxonomyIdsOk returns a tuple with the FullPathTaxonomyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxonomyNodeChildrenInner) GetFullPathTaxonomyIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.FullPathTaxonomyIds) {
		return nil, false
	}
	return o.FullPathTaxonomyIds, true
}

// HasFullPathTaxonomyIds returns a boolean if a field has been set.
func (o *BuyerTaxonomyNodeChildrenInner) HasFullPathTaxonomyIds() bool {
	if o != nil && !IsNil(o.FullPathTaxonomyIds) {
		return true
	}

	return false
}

// SetFullPathTaxonomyIds gets a reference to the given []int32 and assigns it to the FullPathTaxonomyIds field.
func (o *BuyerTaxonomyNodeChildrenInner) SetFullPathTaxonomyIds(v []int32) {
	o.FullPathTaxonomyIds = v
}

func (o BuyerTaxonomyNodeChildrenInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyerTaxonomyNodeChildrenInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.ParentId.IsSet() {
		toSerialize["parent_id"] = o.ParentId.Get()
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !IsNil(o.FullPathTaxonomyIds) {
		toSerialize["full_path_taxonomy_ids"] = o.FullPathTaxonomyIds
	}
	return toSerialize, nil
}

type NullableBuyerTaxonomyNodeChildrenInner struct {
	value *BuyerTaxonomyNodeChildrenInner
	isSet bool
}

func (v NullableBuyerTaxonomyNodeChildrenInner) Get() *BuyerTaxonomyNodeChildrenInner {
	return v.value
}

func (v *NullableBuyerTaxonomyNodeChildrenInner) Set(val *BuyerTaxonomyNodeChildrenInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyerTaxonomyNodeChildrenInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyerTaxonomyNodeChildrenInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyerTaxonomyNodeChildrenInner(val *BuyerTaxonomyNodeChildrenInner) *NullableBuyerTaxonomyNodeChildrenInner {
	return &NullableBuyerTaxonomyNodeChildrenInner{value: val, isSet: true}
}

func (v NullableBuyerTaxonomyNodeChildrenInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyerTaxonomyNodeChildrenInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


