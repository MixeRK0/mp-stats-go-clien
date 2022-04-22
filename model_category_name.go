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

// CategoryName struct for CategoryName
type CategoryName struct {
	// Название категории
	Name string `json:"name"`
}

// NewCategoryName instantiates a new CategoryName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryName(name string) *CategoryName {
	this := CategoryName{}
	this.Name = name
	return &this
}

// NewCategoryNameWithDefaults instantiates a new CategoryName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryNameWithDefaults() *CategoryName {
	this := CategoryName{}
	return &this
}

// GetName returns the Name field value
func (o *CategoryName) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CategoryName) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CategoryName) SetName(v string) {
	o.Name = v
}

func (o CategoryName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryName struct {
	value *CategoryName
	isSet bool
}

func (v NullableCategoryName) Get() *CategoryName {
	return v.value
}

func (v *NullableCategoryName) Set(val *CategoryName) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryName) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryName(val *CategoryName) *NullableCategoryName {
	return &NullableCategoryName{value: val, isSet: true}
}

func (v NullableCategoryName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


