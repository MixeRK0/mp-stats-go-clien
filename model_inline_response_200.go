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

// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	// Номер строки начала получения данных
	StartRow int32 `json:"startRow"`
	// Номер строки конца получения данных
	EndRow int32 `json:"endRow"`
	// 
	FilterModel interface{} `json:"filterModel"`
	SortModel []SortModelItem `json:"sortModel"`
	// Кол-во строк в результирующем запросе без учета пагинации
	Total int32 `json:"total"`
	// Массив данных
	Data []CategoryProduct `json:"data"`
}

// NewInlineResponse200 instantiates a new InlineResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200(startRow int32, endRow int32, filterModel interface{}, sortModel []SortModelItem, total int32, data []CategoryProduct) *InlineResponse200 {
	this := InlineResponse200{}
	this.StartRow = startRow
	this.EndRow = endRow
	this.FilterModel = filterModel
	this.SortModel = sortModel
	this.Total = total
	this.Data = data
	return &this
}

// NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200WithDefaults() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// GetStartRow returns the StartRow field value
func (o *InlineResponse200) GetStartRow() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartRow
}

// GetStartRowOk returns a tuple with the StartRow field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetStartRowOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartRow, true
}

// SetStartRow sets field value
func (o *InlineResponse200) SetStartRow(v int32) {
	o.StartRow = v
}

// GetEndRow returns the EndRow field value
func (o *InlineResponse200) GetEndRow() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EndRow
}

// GetEndRowOk returns a tuple with the EndRow field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetEndRowOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndRow, true
}

// SetEndRow sets field value
func (o *InlineResponse200) SetEndRow(v int32) {
	o.EndRow = v
}

// GetFilterModel returns the FilterModel field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *InlineResponse200) GetFilterModel() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.FilterModel
}

// GetFilterModelOk returns a tuple with the FilterModel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse200) GetFilterModelOk() (*interface{}, bool) {
	if o == nil || o.FilterModel == nil {
		return nil, false
	}
	return &o.FilterModel, true
}

// SetFilterModel sets field value
func (o *InlineResponse200) SetFilterModel(v interface{}) {
	o.FilterModel = v
}

// GetSortModel returns the SortModel field value
func (o *InlineResponse200) GetSortModel() []SortModelItem {
	if o == nil {
		var ret []SortModelItem
		return ret
	}

	return o.SortModel
}

// GetSortModelOk returns a tuple with the SortModel field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetSortModelOk() ([]SortModelItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortModel, true
}

// SetSortModel sets field value
func (o *InlineResponse200) SetSortModel(v []SortModelItem) {
	o.SortModel = v
}

// GetTotal returns the Total field value
func (o *InlineResponse200) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *InlineResponse200) SetTotal(v int32) {
	o.Total = v
}

// GetData returns the Data field value
func (o *InlineResponse200) GetData() []CategoryProduct {
	if o == nil {
		var ret []CategoryProduct
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetDataOk() ([]CategoryProduct, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *InlineResponse200) SetData(v []CategoryProduct) {
	o.Data = v
}

func (o InlineResponse200) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["startRow"] = o.StartRow
	}
	if true {
		toSerialize["endRow"] = o.EndRow
	}
	if o.FilterModel != nil {
		toSerialize["filterModel"] = o.FilterModel
	}
	if true {
		toSerialize["sortModel"] = o.SortModel
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200 struct {
	value *InlineResponse200
	isSet bool
}

func (v NullableInlineResponse200) Get() *InlineResponse200 {
	return v.value
}

func (v *NullableInlineResponse200) Set(val *InlineResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200(val *InlineResponse200) *NullableInlineResponse200 {
	return &NullableInlineResponse200{value: val, isSet: true}
}

func (v NullableInlineResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


