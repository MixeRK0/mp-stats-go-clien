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

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
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
	Data []SearchItemsElement `json:"data"`
}

// NewInlineResponse2001 instantiates a new InlineResponse2001 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001(startRow int32, endRow int32, filterModel interface{}, sortModel []SortModelItem, total int32, data []SearchItemsElement) *InlineResponse2001 {
	this := InlineResponse2001{}
	this.StartRow = startRow
	this.EndRow = endRow
	this.FilterModel = filterModel
	this.SortModel = sortModel
	this.Total = total
	this.Data = data
	return &this
}

// NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001WithDefaults() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// GetStartRow returns the StartRow field value
func (o *InlineResponse2001) GetStartRow() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartRow
}

// GetStartRowOk returns a tuple with the StartRow field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetStartRowOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartRow, true
}

// SetStartRow sets field value
func (o *InlineResponse2001) SetStartRow(v int32) {
	o.StartRow = v
}

// GetEndRow returns the EndRow field value
func (o *InlineResponse2001) GetEndRow() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EndRow
}

// GetEndRowOk returns a tuple with the EndRow field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetEndRowOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndRow, true
}

// SetEndRow sets field value
func (o *InlineResponse2001) SetEndRow(v int32) {
	o.EndRow = v
}

// GetFilterModel returns the FilterModel field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *InlineResponse2001) GetFilterModel() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.FilterModel
}

// GetFilterModelOk returns a tuple with the FilterModel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse2001) GetFilterModelOk() (*interface{}, bool) {
	if o == nil || o.FilterModel == nil {
		return nil, false
	}
	return &o.FilterModel, true
}

// SetFilterModel sets field value
func (o *InlineResponse2001) SetFilterModel(v interface{}) {
	o.FilterModel = v
}

// GetSortModel returns the SortModel field value
func (o *InlineResponse2001) GetSortModel() []SortModelItem {
	if o == nil {
		var ret []SortModelItem
		return ret
	}

	return o.SortModel
}

// GetSortModelOk returns a tuple with the SortModel field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetSortModelOk() ([]SortModelItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortModel, true
}

// SetSortModel sets field value
func (o *InlineResponse2001) SetSortModel(v []SortModelItem) {
	o.SortModel = v
}

// GetTotal returns the Total field value
func (o *InlineResponse2001) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *InlineResponse2001) SetTotal(v int32) {
	o.Total = v
}

// GetData returns the Data field value
func (o *InlineResponse2001) GetData() []SearchItemsElement {
	if o == nil {
		var ret []SearchItemsElement
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetDataOk() ([]SearchItemsElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *InlineResponse2001) SetData(v []SearchItemsElement) {
	o.Data = v
}

func (o InlineResponse2001) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse2001 struct {
	value *InlineResponse2001
	isSet bool
}

func (v NullableInlineResponse2001) Get() *InlineResponse2001 {
	return v.value
}

func (v *NullableInlineResponse2001) Set(val *InlineResponse2001) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001(val *InlineResponse2001) *NullableInlineResponse2001 {
	return &NullableInlineResponse2001{value: val, isSet: true}
}

func (v NullableInlineResponse2001) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


