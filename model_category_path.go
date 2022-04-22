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

// CategoryPath struct for CategoryPath
type CategoryPath struct {
	// Путь по категориям каталога Wildberries
	Path string `json:"path"`
}

// NewCategoryPath instantiates a new CategoryPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryPath(path string) *CategoryPath {
	this := CategoryPath{}
	this.Path = path
	return &this
}

// NewCategoryPathWithDefaults instantiates a new CategoryPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryPathWithDefaults() *CategoryPath {
	this := CategoryPath{}
	return &this
}

// GetPath returns the Path field value
func (o *CategoryPath) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *CategoryPath) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *CategoryPath) SetPath(v string) {
	o.Path = v
}

func (o CategoryPath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryPath struct {
	value *CategoryPath
	isSet bool
}

func (v NullableCategoryPath) Get() *CategoryPath {
	return v.value
}

func (v *NullableCategoryPath) Set(val *CategoryPath) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryPath) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryPath(val *CategoryPath) *NullableCategoryPath {
	return &NullableCategoryPath{value: val, isSet: true}
}

func (v NullableCategoryPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


