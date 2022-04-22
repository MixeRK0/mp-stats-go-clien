# ItemByCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **[]int32** | Остаток | 
**Categories** | **map[string]interface{}** | Категория и позиции в ней. NaN: Позиция не определена | 
**Days** | **[]string** | Дни | 
**FinalPrice** | **[]float32** | Цена со скидкой | 
**Sales** | **[]int32** | Продаж | 

## Methods

### NewItemByCategory

`func NewItemByCategory(balance []int32, categories map[string]interface{}, days []string, finalPrice []float32, sales []int32, ) *ItemByCategory`

NewItemByCategory instantiates a new ItemByCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemByCategoryWithDefaults

`func NewItemByCategoryWithDefaults() *ItemByCategory`

NewItemByCategoryWithDefaults instantiates a new ItemByCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *ItemByCategory) GetBalance() []int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ItemByCategory) GetBalanceOk() (*[]int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ItemByCategory) SetBalance(v []int32)`

SetBalance sets Balance field to given value.


### GetCategories

`func (o *ItemByCategory) GetCategories() map[string]interface{}`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ItemByCategory) GetCategoriesOk() (*map[string]interface{}, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ItemByCategory) SetCategories(v map[string]interface{})`

SetCategories sets Categories field to given value.


### GetDays

`func (o *ItemByCategory) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ItemByCategory) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ItemByCategory) SetDays(v []string)`

SetDays sets Days field to given value.


### GetFinalPrice

`func (o *ItemByCategory) GetFinalPrice() []float32`

GetFinalPrice returns the FinalPrice field if non-nil, zero value otherwise.

### GetFinalPriceOk

`func (o *ItemByCategory) GetFinalPriceOk() (*[]float32, bool)`

GetFinalPriceOk returns a tuple with the FinalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPrice

`func (o *ItemByCategory) SetFinalPrice(v []float32)`

SetFinalPrice sets FinalPrice field to given value.


### GetSales

`func (o *ItemByCategory) GetSales() []int32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *ItemByCategory) GetSalesOk() (*[]int32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *ItemByCategory) SetSales(v []int32)`

SetSales sets Sales field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


