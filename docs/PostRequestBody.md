# PostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartRow** | **int32** | Номер строки начала получения данных | 
**EndRow** | **int32** | Номер строки конца получения данных | 
**FilterModel** | **interface{}** |  | 
**SortModel** | [**[]SortModelItem**](SortModelItem.md) |  | 

## Methods

### NewPostRequestBody

`func NewPostRequestBody(startRow int32, endRow int32, filterModel interface{}, sortModel []SortModelItem, ) *PostRequestBody`

NewPostRequestBody instantiates a new PostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRequestBodyWithDefaults

`func NewPostRequestBodyWithDefaults() *PostRequestBody`

NewPostRequestBodyWithDefaults instantiates a new PostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartRow

`func (o *PostRequestBody) GetStartRow() int32`

GetStartRow returns the StartRow field if non-nil, zero value otherwise.

### GetStartRowOk

`func (o *PostRequestBody) GetStartRowOk() (*int32, bool)`

GetStartRowOk returns a tuple with the StartRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartRow

`func (o *PostRequestBody) SetStartRow(v int32)`

SetStartRow sets StartRow field to given value.


### GetEndRow

`func (o *PostRequestBody) GetEndRow() int32`

GetEndRow returns the EndRow field if non-nil, zero value otherwise.

### GetEndRowOk

`func (o *PostRequestBody) GetEndRowOk() (*int32, bool)`

GetEndRowOk returns a tuple with the EndRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndRow

`func (o *PostRequestBody) SetEndRow(v int32)`

SetEndRow sets EndRow field to given value.


### GetFilterModel

`func (o *PostRequestBody) GetFilterModel() interface{}`

GetFilterModel returns the FilterModel field if non-nil, zero value otherwise.

### GetFilterModelOk

`func (o *PostRequestBody) GetFilterModelOk() (*interface{}, bool)`

GetFilterModelOk returns a tuple with the FilterModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterModel

`func (o *PostRequestBody) SetFilterModel(v interface{})`

SetFilterModel sets FilterModel field to given value.


### SetFilterModelNil

`func (o *PostRequestBody) SetFilterModelNil(b bool)`

 SetFilterModelNil sets the value for FilterModel to be an explicit nil

### UnsetFilterModel
`func (o *PostRequestBody) UnsetFilterModel()`

UnsetFilterModel ensures that no value is present for FilterModel, not even an explicit nil
### GetSortModel

`func (o *PostRequestBody) GetSortModel() []SortModelItem`

GetSortModel returns the SortModel field if non-nil, zero value otherwise.

### GetSortModelOk

`func (o *PostRequestBody) GetSortModelOk() (*[]SortModelItem, bool)`

GetSortModelOk returns a tuple with the SortModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortModel

`func (o *PostRequestBody) SetSortModel(v []SortModelItem)`

SetSortModel sets SortModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


