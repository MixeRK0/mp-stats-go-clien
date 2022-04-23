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

// SearchCategories struct for SearchCategories
type SearchCategories struct {
	// Кол-во строк в результирующем запросе без учета пагинации
	Total int32 `json:"total"`
	// Ошибка
	Error bool `json:"error"`
	// Номер строки начала получения данных
	StartRow int32 `json:"startRow"`
	// Номер строки конца получения данных
	EndRow int32 `json:"endRow"`
	RowGroupCols []int32 `json:"rowGroupCols"`
	ValueCols []int32 `json:"valueCols"`
	PivotCols []int32 `json:"pivotCols"`
	PivotMode bool `json:"pivotMode"`
	GroupKeys []int32 `json:"groupKeys"`
	// 
	FilterModel map[string]interface{} `json:"filterModel"`
	SortModel []SortModelItem `json:"sortModel"`
	// tpl
	Tpl string `json:"tpl"`
	// Массив данных
	Data []SearchCategoriesItem `json:"data"`
}

// NewSearchCategories instantiates a new SearchCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCategories(total int32, error_ bool, startRow int32, endRow int32, rowGroupCols []int32, valueCols []int32, pivotCols []int32, pivotMode bool, groupKeys []int32, filterModel map[string]interface{}, sortModel []SortModelItem, tpl string, data []SearchCategoriesItem) *SearchCategories {
	this := SearchCategories{}
	this.Total = total
	this.Error = error_
	this.StartRow = startRow
	this.EndRow = endRow
	this.RowGroupCols = rowGroupCols
	this.ValueCols = valueCols
	this.PivotCols = pivotCols
	this.PivotMode = pivotMode
	this.GroupKeys = groupKeys
	this.FilterModel = filterModel
	this.SortModel = sortModel
	this.Tpl = tpl
	this.Data = data
	return &this
}

// NewSearchCategoriesWithDefaults instantiates a new SearchCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCategoriesWithDefaults() *SearchCategories {
	this := SearchCategories{}
	return &this
}

// GetTotal returns the Total field value
func (o *SearchCategories) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *SearchCategories) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *SearchCategories) SetTotal(v int32) {
	o.Total = v
}

// GetError returns the Error field value
func (o *SearchCategories) GetError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *SearchCategories) GetErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *SearchCategories) SetError(v bool) {
	o.Error = v
}

// GetStartRow returns the StartRow field value
func (o *SearchCategories) GetStartRow() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartRow
}

// GetStartRowOk returns a tuple with the StartRow field value
// and a boolean to check if the value has been set.
func (o *SearchCategories) GetStartRowOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartRow, true
}

// SetStartRow sets field value
func (o *SearchCategories) SetStartRow(v int32) {
	o.StartRow = v
}

// GetEndRow returns the EndRow field value
func (o *SearchCategories) GetEndRow() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EndRow
}

// GetEndRowOk returns a tuple with the EndRow field value
// and a boolean to check if the value has been set.
func (o *SearchCategories) GetEndRowOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndRow, true
}

// SetEndRow sets field value
func (o *SearchCategories) SetEndRow(v int32) {
	o.EndRow = v
}

// GetRowGroupCols returns the RowGroupCols field value
func (o *SearchCategories) GetRowGroupCols() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.RowGroupCols
}

// GetRowGroupColsOk returns a tuple with the RowGroupCols field value
// and a boolean to check if the value has been set.
func (o *SearchCategories) GetRowGroupColsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RowGroupCols, true
}

// SetRowGroupCols sets field value
func (o *SearchCategories) SetRowGroupCols(v []int32) {
	o.RowGroupCols = v
}

// GetValueCols returns the ValueCols field value
func (o *SearchCategories) GetValueCols() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ValueCols
}

// GetValueColsOk returns a tuple with the ValueCols field value
// and a boolean to check if the value has been set.
func (o *SearchCategories) GetValueColsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueCols, true
}

// SetValueCols sets field value
func (o *SearchCategories) SetValueCols(v []int32) {
	o.ValueCols = v
}

// GetPivotCols returns the PivotCols field value
func (o *SearchCategories) GetPivotCols() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.PivotCols
}

// GetPivotColsOk returns a tuple with the PivotCols field value
// and a boolean to check if the value has been set.
func (o *SearchCategories) GetPivotColsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PivotCols, true
}

// SetPivotCols sets field value
func (o *SearchCategories) SetPivotCols(v []int32) {
	o.PivotCols = v
}

// GetPivotMode returns the PivotMode field value
func (o *SearchCategories) GetPivotMode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PivotMode
}

// GetPivotModeOk returns a tuple with the PivotMode field value
// and a boolean to check if the value has been set.
func (o *SearchCategories) GetPivotModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PivotMode, true
}

// SetPivotMode sets field value
func (o *SearchCategories) SetPivotMode(v bool) {
	o.PivotMode = v
}

// GetGroupKeys returns the GroupKeys field value
func (o *SearchCategories) GetGroupKeys() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.GroupKeys
}

// GetGroupKeysOk returns a tuple with the GroupKeys field value
// and a boolean to check if the value has been set.
func (o *SearchCategories) GetGroupKeysOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupKeys, true
}

// SetGroupKeys sets field value
func (o *SearchCategories) SetGroupKeys(v []int32) {
	o.GroupKeys = v
}

// GetFilterModel returns the FilterModel field value
func (o *SearchCategories) GetFilterModel() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.FilterModel
}

// GetFilterModelOk returns a tuple with the FilterModel field value
// and a boolean to check if the value has been set.
func (o *SearchCategories) GetFilterModelOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterModel, true
}

// SetFilterModel sets field value
func (o *SearchCategories) SetFilterModel(v map[string]interface{}) {
	o.FilterModel = v
}

// GetSortModel returns the SortModel field value
func (o *SearchCategories) GetSortModel() []SortModelItem {
	if o == nil {
		var ret []SortModelItem
		return ret
	}

	return o.SortModel
}

// GetSortModelOk returns a tuple with the SortModel field value
// and a boolean to check if the value has been set.
func (o *SearchCategories) GetSortModelOk() ([]SortModelItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortModel, true
}

// SetSortModel sets field value
func (o *SearchCategories) SetSortModel(v []SortModelItem) {
	o.SortModel = v
}

// GetTpl returns the Tpl field value
func (o *SearchCategories) GetTpl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tpl
}

// GetTplOk returns a tuple with the Tpl field value
// and a boolean to check if the value has been set.
func (o *SearchCategories) GetTplOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tpl, true
}

// SetTpl sets field value
func (o *SearchCategories) SetTpl(v string) {
	o.Tpl = v
}

// GetData returns the Data field value
func (o *SearchCategories) GetData() []SearchCategoriesItem {
	if o == nil {
		var ret []SearchCategoriesItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SearchCategories) GetDataOk() ([]SearchCategoriesItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SearchCategories) SetData(v []SearchCategoriesItem) {
	o.Data = v
}

func (o SearchCategories) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["error"] = o.Error
	}
	if true {
		toSerialize["startRow"] = o.StartRow
	}
	if true {
		toSerialize["endRow"] = o.EndRow
	}
	if true {
		toSerialize["rowGroupCols"] = o.RowGroupCols
	}
	if true {
		toSerialize["valueCols"] = o.ValueCols
	}
	if true {
		toSerialize["pivotCols"] = o.PivotCols
	}
	if true {
		toSerialize["pivotMode"] = o.PivotMode
	}
	if true {
		toSerialize["groupKeys"] = o.GroupKeys
	}
	if true {
		toSerialize["filterModel"] = o.FilterModel
	}
	if true {
		toSerialize["sortModel"] = o.SortModel
	}
	if true {
		toSerialize["tpl"] = o.Tpl
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSearchCategories struct {
	value *SearchCategories
	isSet bool
}

func (v NullableSearchCategories) Get() *SearchCategories {
	return v.value
}

func (v *NullableSearchCategories) Set(val *SearchCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCategories(val *SearchCategories) *NullableSearchCategories {
	return &NullableSearchCategories{value: val, isSet: true}
}

func (v NullableSearchCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


