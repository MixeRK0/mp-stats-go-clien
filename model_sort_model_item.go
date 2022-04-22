/*
mpstats

MPStats API

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SortModelItem 
type SortModelItem struct {
	ColId string `json:"colId"`
	Sort string `json:"sort"`
}

// NewSortModelItem instantiates a new SortModelItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSortModelItem(colId string, sort string) *SortModelItem {
	this := SortModelItem{}
	this.ColId = colId
	this.Sort = sort
	return &this
}

// NewSortModelItemWithDefaults instantiates a new SortModelItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSortModelItemWithDefaults() *SortModelItem {
	this := SortModelItem{}
	return &this
}

// GetColId returns the ColId field value
func (o *SortModelItem) GetColId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ColId
}

// GetColIdOk returns a tuple with the ColId field value
// and a boolean to check if the value has been set.
func (o *SortModelItem) GetColIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColId, true
}

// SetColId sets field value
func (o *SortModelItem) SetColId(v string) {
	o.ColId = v
}

// GetSort returns the Sort field value
func (o *SortModelItem) GetSort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value
// and a boolean to check if the value has been set.
func (o *SortModelItem) GetSortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sort, true
}

// SetSort sets field value
func (o *SortModelItem) SetSort(v string) {
	o.Sort = v
}

func (o SortModelItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["colId"] = o.ColId
	}
	if true {
		toSerialize["sort"] = o.Sort
	}
	return json.Marshal(toSerialize)
}

type NullableSortModelItem struct {
	value *SortModelItem
	isSet bool
}

func (v NullableSortModelItem) Get() *SortModelItem {
	return v.value
}

func (v *NullableSortModelItem) Set(val *SortModelItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSortModelItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSortModelItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortModelItem(val *SortModelItem) *NullableSortModelItem {
	return &NullableSortModelItem{value: val, isSet: true}
}

func (v NullableSortModelItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortModelItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


