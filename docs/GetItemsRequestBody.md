# GetItemsRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartRow** | **int32** |  | 
**EndRow** | **int32** |  | 
**FilterModel** | **interface{}** |  | 
**GroupKeys** | **[]string** |  | 
**PivotCols** | **[]string** |  | 
**PivotMode** | **bool** |  | 
**RowGroupCols** | **[]string** |  | 
**SortModel** | [**[]SortModelItem**](SortModelItem.md) |  | 
**ValueCols** | **[]string** |  | 

## Methods

### NewGetItemsRequestBody

`func NewGetItemsRequestBody(startRow int32, endRow int32, filterModel interface{}, groupKeys []string, pivotCols []string, pivotMode bool, rowGroupCols []string, sortModel []SortModelItem, valueCols []string, ) *GetItemsRequestBody`

NewGetItemsRequestBody instantiates a new GetItemsRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetItemsRequestBodyWithDefaults

`func NewGetItemsRequestBodyWithDefaults() *GetItemsRequestBody`

NewGetItemsRequestBodyWithDefaults instantiates a new GetItemsRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartRow

`func (o *GetItemsRequestBody) GetStartRow() int32`

GetStartRow returns the StartRow field if non-nil, zero value otherwise.

### GetStartRowOk

`func (o *GetItemsRequestBody) GetStartRowOk() (*int32, bool)`

GetStartRowOk returns a tuple with the StartRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartRow

`func (o *GetItemsRequestBody) SetStartRow(v int32)`

SetStartRow sets StartRow field to given value.


### GetEndRow

`func (o *GetItemsRequestBody) GetEndRow() int32`

GetEndRow returns the EndRow field if non-nil, zero value otherwise.

### GetEndRowOk

`func (o *GetItemsRequestBody) GetEndRowOk() (*int32, bool)`

GetEndRowOk returns a tuple with the EndRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndRow

`func (o *GetItemsRequestBody) SetEndRow(v int32)`

SetEndRow sets EndRow field to given value.


### GetFilterModel

`func (o *GetItemsRequestBody) GetFilterModel() interface{}`

GetFilterModel returns the FilterModel field if non-nil, zero value otherwise.

### GetFilterModelOk

`func (o *GetItemsRequestBody) GetFilterModelOk() (*interface{}, bool)`

GetFilterModelOk returns a tuple with the FilterModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterModel

`func (o *GetItemsRequestBody) SetFilterModel(v interface{})`

SetFilterModel sets FilterModel field to given value.


### SetFilterModelNil

`func (o *GetItemsRequestBody) SetFilterModelNil(b bool)`

 SetFilterModelNil sets the value for FilterModel to be an explicit nil

### UnsetFilterModel
`func (o *GetItemsRequestBody) UnsetFilterModel()`

UnsetFilterModel ensures that no value is present for FilterModel, not even an explicit nil
### GetGroupKeys

`func (o *GetItemsRequestBody) GetGroupKeys() []string`

GetGroupKeys returns the GroupKeys field if non-nil, zero value otherwise.

### GetGroupKeysOk

`func (o *GetItemsRequestBody) GetGroupKeysOk() (*[]string, bool)`

GetGroupKeysOk returns a tuple with the GroupKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKeys

`func (o *GetItemsRequestBody) SetGroupKeys(v []string)`

SetGroupKeys sets GroupKeys field to given value.


### GetPivotCols

`func (o *GetItemsRequestBody) GetPivotCols() []string`

GetPivotCols returns the PivotCols field if non-nil, zero value otherwise.

### GetPivotColsOk

`func (o *GetItemsRequestBody) GetPivotColsOk() (*[]string, bool)`

GetPivotColsOk returns a tuple with the PivotCols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPivotCols

`func (o *GetItemsRequestBody) SetPivotCols(v []string)`

SetPivotCols sets PivotCols field to given value.


### GetPivotMode

`func (o *GetItemsRequestBody) GetPivotMode() bool`

GetPivotMode returns the PivotMode field if non-nil, zero value otherwise.

### GetPivotModeOk

`func (o *GetItemsRequestBody) GetPivotModeOk() (*bool, bool)`

GetPivotModeOk returns a tuple with the PivotMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPivotMode

`func (o *GetItemsRequestBody) SetPivotMode(v bool)`

SetPivotMode sets PivotMode field to given value.


### GetRowGroupCols

`func (o *GetItemsRequestBody) GetRowGroupCols() []string`

GetRowGroupCols returns the RowGroupCols field if non-nil, zero value otherwise.

### GetRowGroupColsOk

`func (o *GetItemsRequestBody) GetRowGroupColsOk() (*[]string, bool)`

GetRowGroupColsOk returns a tuple with the RowGroupCols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowGroupCols

`func (o *GetItemsRequestBody) SetRowGroupCols(v []string)`

SetRowGroupCols sets RowGroupCols field to given value.


### GetSortModel

`func (o *GetItemsRequestBody) GetSortModel() []SortModelItem`

GetSortModel returns the SortModel field if non-nil, zero value otherwise.

### GetSortModelOk

`func (o *GetItemsRequestBody) GetSortModelOk() (*[]SortModelItem, bool)`

GetSortModelOk returns a tuple with the SortModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortModel

`func (o *GetItemsRequestBody) SetSortModel(v []SortModelItem)`

SetSortModel sets SortModel field to given value.


### GetValueCols

`func (o *GetItemsRequestBody) GetValueCols() []string`

GetValueCols returns the ValueCols field if non-nil, zero value otherwise.

### GetValueColsOk

`func (o *GetItemsRequestBody) GetValueColsOk() (*[]string, bool)`

GetValueColsOk returns a tuple with the ValueCols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCols

`func (o *GetItemsRequestBody) SetValueCols(v []string)`

SetValueCols sets ValueCols field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


