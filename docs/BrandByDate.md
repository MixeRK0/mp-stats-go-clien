# BrandByDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | Дата | 
**Items** | **int32** | Число позиций | 
**Comments** | **float32** | Среднее число комментариев на товар | 
**Rating** | **float32** | Средний рейтинг | 
**Balance** | **int32** | Число товаров, числящихся в наличии | 
**BalancePrice** | **int32** | Произведение остатка на цену продажи | 
**AvgPrice** | **float32** | Средняя цена товара у бренда (все товары, среднее арифметическое) | 
**AvgSalePrice** | **float32** | Средняя цена состоявшейся продажи (деление выручки на число продаж) | 
**Sales** | **int32** | Число продаж единиц товара | 
**Revenue** | **float32** | Выручка за сутки | 

## Methods

### NewBrandByDate

`func NewBrandByDate(data string, items int32, comments float32, rating float32, balance int32, balancePrice int32, avgPrice float32, avgSalePrice float32, sales int32, revenue float32, ) *BrandByDate`

NewBrandByDate instantiates a new BrandByDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandByDateWithDefaults

`func NewBrandByDateWithDefaults() *BrandByDate`

NewBrandByDateWithDefaults instantiates a new BrandByDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BrandByDate) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BrandByDate) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BrandByDate) SetData(v string)`

SetData sets Data field to given value.


### GetItems

`func (o *BrandByDate) GetItems() int32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BrandByDate) GetItemsOk() (*int32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BrandByDate) SetItems(v int32)`

SetItems sets Items field to given value.


### GetComments

`func (o *BrandByDate) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BrandByDate) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BrandByDate) SetComments(v float32)`

SetComments sets Comments field to given value.


### GetRating

`func (o *BrandByDate) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *BrandByDate) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *BrandByDate) SetRating(v float32)`

SetRating sets Rating field to given value.


### GetBalance

`func (o *BrandByDate) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BrandByDate) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BrandByDate) SetBalance(v int32)`

SetBalance sets Balance field to given value.


### GetBalancePrice

`func (o *BrandByDate) GetBalancePrice() int32`

GetBalancePrice returns the BalancePrice field if non-nil, zero value otherwise.

### GetBalancePriceOk

`func (o *BrandByDate) GetBalancePriceOk() (*int32, bool)`

GetBalancePriceOk returns a tuple with the BalancePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePrice

`func (o *BrandByDate) SetBalancePrice(v int32)`

SetBalancePrice sets BalancePrice field to given value.


### GetAvgPrice

`func (o *BrandByDate) GetAvgPrice() float32`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *BrandByDate) GetAvgPriceOk() (*float32, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *BrandByDate) SetAvgPrice(v float32)`

SetAvgPrice sets AvgPrice field to given value.


### GetAvgSalePrice

`func (o *BrandByDate) GetAvgSalePrice() float32`

GetAvgSalePrice returns the AvgSalePrice field if non-nil, zero value otherwise.

### GetAvgSalePriceOk

`func (o *BrandByDate) GetAvgSalePriceOk() (*float32, bool)`

GetAvgSalePriceOk returns a tuple with the AvgSalePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSalePrice

`func (o *BrandByDate) SetAvgSalePrice(v float32)`

SetAvgSalePrice sets AvgSalePrice field to given value.


### GetSales

`func (o *BrandByDate) GetSales() int32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *BrandByDate) GetSalesOk() (*int32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *BrandByDate) SetSales(v int32)`

SetSales sets Sales field to given value.


### GetRevenue

`func (o *BrandByDate) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *BrandByDate) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *BrandByDate) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


