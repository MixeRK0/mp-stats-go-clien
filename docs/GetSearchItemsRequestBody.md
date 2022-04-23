# GetSearchItemsRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartRow** | **int32** |  | 
**EndRow** | **int32** |  | 
**FilterModel** | Pointer to **interface{}** |  | [optional] 
**GroupKeys** | Pointer to **[]string** |  | [optional] 
**PivotCols** | Pointer to **[]string** |  | [optional] 
**PivotMode** | **bool** |  | 
**RowGroupCols** | Pointer to **[]string** |  | [optional] 
**SortModel** | Pointer to [**[]SortModelItem**](SortModelItem.md) |  | [optional] 
**Tpl** | **string** |  | 
**ValueCols** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetSearchItemsRequestBody

`func NewGetSearchItemsRequestBody(startRow int32, endRow int32, pivotMode bool, tpl string, ) *GetSearchItemsRequestBody`

NewGetSearchItemsRequestBody instantiates a new GetSearchItemsRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSearchItemsRequestBodyWithDefaults

`func NewGetSearchItemsRequestBodyWithDefaults() *GetSearchItemsRequestBody`

NewGetSearchItemsRequestBodyWithDefaults instantiates a new GetSearchItemsRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartRow

`func (o *GetSearchItemsRequestBody) GetStartRow() int32`

GetStartRow returns the StartRow field if non-nil, zero value otherwise.

### GetStartRowOk

`func (o *GetSearchItemsRequestBody) GetStartRowOk() (*int32, bool)`

GetStartRowOk returns a tuple with the StartRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartRow

`func (o *GetSearchItemsRequestBody) SetStartRow(v int32)`

SetStartRow sets StartRow field to given value.


### GetEndRow

`func (o *GetSearchItemsRequestBody) GetEndRow() int32`

GetEndRow returns the EndRow field if non-nil, zero value otherwise.

### GetEndRowOk

`func (o *GetSearchItemsRequestBody) GetEndRowOk() (*int32, bool)`

GetEndRowOk returns a tuple with the EndRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndRow

`func (o *GetSearchItemsRequestBody) SetEndRow(v int32)`

SetEndRow sets EndRow field to given value.


### GetFilterModel

`func (o *GetSearchItemsRequestBody) GetFilterModel() interface{}`

GetFilterModel returns the FilterModel field if non-nil, zero value otherwise.

### GetFilterModelOk

`func (o *GetSearchItemsRequestBody) GetFilterModelOk() (*interface{}, bool)`

GetFilterModelOk returns a tuple with the FilterModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterModel

`func (o *GetSearchItemsRequestBody) SetFilterModel(v interface{})`

SetFilterModel sets FilterModel field to given value.

### HasFilterModel

`func (o *GetSearchItemsRequestBody) HasFilterModel() bool`

HasFilterModel returns a boolean if a field has been set.

### SetFilterModelNil

`func (o *GetSearchItemsRequestBody) SetFilterModelNil(b bool)`

 SetFilterModelNil sets the value for FilterModel to be an explicit nil

### UnsetFilterModel
`func (o *GetSearchItemsRequestBody) UnsetFilterModel()`

UnsetFilterModel ensures that no value is present for FilterModel, not even an explicit nil
### GetGroupKeys

`func (o *GetSearchItemsRequestBody) GetGroupKeys() []string`

GetGroupKeys returns the GroupKeys field if non-nil, zero value otherwise.

### GetGroupKeysOk

`func (o *GetSearchItemsRequestBody) GetGroupKeysOk() (*[]string, bool)`

GetGroupKeysOk returns a tuple with the GroupKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKeys

`func (o *GetSearchItemsRequestBody) SetGroupKeys(v []string)`

SetGroupKeys sets GroupKeys field to given value.

### HasGroupKeys

`func (o *GetSearchItemsRequestBody) HasGroupKeys() bool`

HasGroupKeys returns a boolean if a field has been set.

### GetPivotCols

`func (o *GetSearchItemsRequestBody) GetPivotCols() []string`

GetPivotCols returns the PivotCols field if non-nil, zero value otherwise.

### GetPivotColsOk

`func (o *GetSearchItemsRequestBody) GetPivotColsOk() (*[]string, bool)`

GetPivotColsOk returns a tuple with the PivotCols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPivotCols

`func (o *GetSearchItemsRequestBody) SetPivotCols(v []string)`

SetPivotCols sets PivotCols field to given value.

### HasPivotCols

`func (o *GetSearchItemsRequestBody) HasPivotCols() bool`

HasPivotCols returns a boolean if a field has been set.

### GetPivotMode

`func (o *GetSearchItemsRequestBody) GetPivotMode() bool`

GetPivotMode returns the PivotMode field if non-nil, zero value otherwise.

### GetPivotModeOk

`func (o *GetSearchItemsRequestBody) GetPivotModeOk() (*bool, bool)`

GetPivotModeOk returns a tuple with the PivotMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPivotMode

`func (o *GetSearchItemsRequestBody) SetPivotMode(v bool)`

SetPivotMode sets PivotMode field to given value.


### GetRowGroupCols

`func (o *GetSearchItemsRequestBody) GetRowGroupCols() []string`

GetRowGroupCols returns the RowGroupCols field if non-nil, zero value otherwise.

### GetRowGroupColsOk

`func (o *GetSearchItemsRequestBody) GetRowGroupColsOk() (*[]string, bool)`

GetRowGroupColsOk returns a tuple with the RowGroupCols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowGroupCols

`func (o *GetSearchItemsRequestBody) SetRowGroupCols(v []string)`

SetRowGroupCols sets RowGroupCols field to given value.

### HasRowGroupCols

`func (o *GetSearchItemsRequestBody) HasRowGroupCols() bool`

HasRowGroupCols returns a boolean if a field has been set.

### GetSortModel

`func (o *GetSearchItemsRequestBody) GetSortModel() []SortModelItem`

GetSortModel returns the SortModel field if non-nil, zero value otherwise.

### GetSortModelOk

`func (o *GetSearchItemsRequestBody) GetSortModelOk() (*[]SortModelItem, bool)`

GetSortModelOk returns a tuple with the SortModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortModel

`func (o *GetSearchItemsRequestBody) SetSortModel(v []SortModelItem)`

SetSortModel sets SortModel field to given value.

### HasSortModel

`func (o *GetSearchItemsRequestBody) HasSortModel() bool`

HasSortModel returns a boolean if a field has been set.

### GetTpl

`func (o *GetSearchItemsRequestBody) GetTpl() string`

GetTpl returns the Tpl field if non-nil, zero value otherwise.

### GetTplOk

`func (o *GetSearchItemsRequestBody) GetTplOk() (*string, bool)`

GetTplOk returns a tuple with the Tpl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpl

`func (o *GetSearchItemsRequestBody) SetTpl(v string)`

SetTpl sets Tpl field to given value.


### GetValueCols

`func (o *GetSearchItemsRequestBody) GetValueCols() []string`

GetValueCols returns the ValueCols field if non-nil, zero value otherwise.

### GetValueColsOk

`func (o *GetSearchItemsRequestBody) GetValueColsOk() (*[]string, bool)`

GetValueColsOk returns a tuple with the ValueCols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCols

`func (o *GetSearchItemsRequestBody) SetValueCols(v []string)`

SetValueCols sets ValueCols field to given value.

### HasValueCols

`func (o *GetSearchItemsRequestBody) HasValueCols() bool`

HasValueCols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


