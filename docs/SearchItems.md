# SearchItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Кол-во строк в результирующем запросе без учета пагинации | 
**Error** | **bool** | Ошибка | 
**StartRow** | **int32** | Номер строки начала получения данных | 
**EndRow** | **int32** | Номер строки конца получения данных | 
**RowGroupCols** | **[]int32** |  | 
**ValueCols** | **[]int32** |  | 
**PivotCols** | **[]int32** |  | 
**PivotMode** | **bool** |  | 
**GroupKeys** | **[]int32** |  | 
**FilterModel** | **interface{}** |  | 
**SortModel** | [**[]SortModelItem**](SortModelItem.md) |  | 
**Tpl** | **string** | tpl | 
**Data** | [**[]SearchItemsElement**](SearchItemsElement.md) | Массив данных | 

## Methods

### NewSearchItems

`func NewSearchItems(total int32, error_ bool, startRow int32, endRow int32, rowGroupCols []int32, valueCols []int32, pivotCols []int32, pivotMode bool, groupKeys []int32, filterModel interface{}, sortModel []SortModelItem, tpl string, data []SearchItemsElement, ) *SearchItems`

NewSearchItems instantiates a new SearchItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchItemsWithDefaults

`func NewSearchItemsWithDefaults() *SearchItems`

NewSearchItemsWithDefaults instantiates a new SearchItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *SearchItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchItems) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetError

`func (o *SearchItems) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchItems) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchItems) SetError(v bool)`

SetError sets Error field to given value.


### GetStartRow

`func (o *SearchItems) GetStartRow() int32`

GetStartRow returns the StartRow field if non-nil, zero value otherwise.

### GetStartRowOk

`func (o *SearchItems) GetStartRowOk() (*int32, bool)`

GetStartRowOk returns a tuple with the StartRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartRow

`func (o *SearchItems) SetStartRow(v int32)`

SetStartRow sets StartRow field to given value.


### GetEndRow

`func (o *SearchItems) GetEndRow() int32`

GetEndRow returns the EndRow field if non-nil, zero value otherwise.

### GetEndRowOk

`func (o *SearchItems) GetEndRowOk() (*int32, bool)`

GetEndRowOk returns a tuple with the EndRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndRow

`func (o *SearchItems) SetEndRow(v int32)`

SetEndRow sets EndRow field to given value.


### GetRowGroupCols

`func (o *SearchItems) GetRowGroupCols() []int32`

GetRowGroupCols returns the RowGroupCols field if non-nil, zero value otherwise.

### GetRowGroupColsOk

`func (o *SearchItems) GetRowGroupColsOk() (*[]int32, bool)`

GetRowGroupColsOk returns a tuple with the RowGroupCols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowGroupCols

`func (o *SearchItems) SetRowGroupCols(v []int32)`

SetRowGroupCols sets RowGroupCols field to given value.


### GetValueCols

`func (o *SearchItems) GetValueCols() []int32`

GetValueCols returns the ValueCols field if non-nil, zero value otherwise.

### GetValueColsOk

`func (o *SearchItems) GetValueColsOk() (*[]int32, bool)`

GetValueColsOk returns a tuple with the ValueCols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCols

`func (o *SearchItems) SetValueCols(v []int32)`

SetValueCols sets ValueCols field to given value.


### GetPivotCols

`func (o *SearchItems) GetPivotCols() []int32`

GetPivotCols returns the PivotCols field if non-nil, zero value otherwise.

### GetPivotColsOk

`func (o *SearchItems) GetPivotColsOk() (*[]int32, bool)`

GetPivotColsOk returns a tuple with the PivotCols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPivotCols

`func (o *SearchItems) SetPivotCols(v []int32)`

SetPivotCols sets PivotCols field to given value.


### GetPivotMode

`func (o *SearchItems) GetPivotMode() bool`

GetPivotMode returns the PivotMode field if non-nil, zero value otherwise.

### GetPivotModeOk

`func (o *SearchItems) GetPivotModeOk() (*bool, bool)`

GetPivotModeOk returns a tuple with the PivotMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPivotMode

`func (o *SearchItems) SetPivotMode(v bool)`

SetPivotMode sets PivotMode field to given value.


### GetGroupKeys

`func (o *SearchItems) GetGroupKeys() []int32`

GetGroupKeys returns the GroupKeys field if non-nil, zero value otherwise.

### GetGroupKeysOk

`func (o *SearchItems) GetGroupKeysOk() (*[]int32, bool)`

GetGroupKeysOk returns a tuple with the GroupKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKeys

`func (o *SearchItems) SetGroupKeys(v []int32)`

SetGroupKeys sets GroupKeys field to given value.


### GetFilterModel

`func (o *SearchItems) GetFilterModel() interface{}`

GetFilterModel returns the FilterModel field if non-nil, zero value otherwise.

### GetFilterModelOk

`func (o *SearchItems) GetFilterModelOk() (*interface{}, bool)`

GetFilterModelOk returns a tuple with the FilterModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterModel

`func (o *SearchItems) SetFilterModel(v interface{})`

SetFilterModel sets FilterModel field to given value.


### SetFilterModelNil

`func (o *SearchItems) SetFilterModelNil(b bool)`

 SetFilterModelNil sets the value for FilterModel to be an explicit nil

### UnsetFilterModel
`func (o *SearchItems) UnsetFilterModel()`

UnsetFilterModel ensures that no value is present for FilterModel, not even an explicit nil
### GetSortModel

`func (o *SearchItems) GetSortModel() []SortModelItem`

GetSortModel returns the SortModel field if non-nil, zero value otherwise.

### GetSortModelOk

`func (o *SearchItems) GetSortModelOk() (*[]SortModelItem, bool)`

GetSortModelOk returns a tuple with the SortModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortModel

`func (o *SearchItems) SetSortModel(v []SortModelItem)`

SetSortModel sets SortModel field to given value.


### GetTpl

`func (o *SearchItems) GetTpl() string`

GetTpl returns the Tpl field if non-nil, zero value otherwise.

### GetTplOk

`func (o *SearchItems) GetTplOk() (*string, bool)`

GetTplOk returns a tuple with the Tpl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpl

`func (o *SearchItems) SetTpl(v string)`

SetTpl sets Tpl field to given value.


### GetData

`func (o *SearchItems) GetData() []SearchItemsElement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SearchItems) GetDataOk() (*[]SearchItemsElement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SearchItems) SetData(v []SearchItemsElement)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


