# CategoryByDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** |  Дата | 
**Brands** | **int32** | Кол-во брендов | 
**Sellers** | **int32** | Кол-во продавцов | 
**AvgPrice** | **float32** | Средняя цена товара в категории (все товары, среднее арифметическое) | 
**AvgSalePrice** | **float32** | Средняя цена состоявшейся продажи (деление выручки на число продаж) | 
**Items** | **int32** | Число товаров | 
**Sales** | **int32** | Число зафиксированных продаж (единицы) | 
**Revenue** | **float32** | Сумма произведений числа проданных товаров на их стоимость | 
**Comments** | **float32** | Среднее число комментариев | 
**Rating** | **float32** | Средний рейтинг | 

## Methods

### NewCategoryByDate

`func NewCategoryByDate(data string, brands int32, sellers int32, avgPrice float32, avgSalePrice float32, items int32, sales int32, revenue float32, comments float32, rating float32, ) *CategoryByDate`

NewCategoryByDate instantiates a new CategoryByDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryByDateWithDefaults

`func NewCategoryByDateWithDefaults() *CategoryByDate`

NewCategoryByDateWithDefaults instantiates a new CategoryByDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CategoryByDate) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CategoryByDate) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CategoryByDate) SetData(v string)`

SetData sets Data field to given value.


### GetBrands

`func (o *CategoryByDate) GetBrands() int32`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *CategoryByDate) GetBrandsOk() (*int32, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *CategoryByDate) SetBrands(v int32)`

SetBrands sets Brands field to given value.


### GetSellers

`func (o *CategoryByDate) GetSellers() int32`

GetSellers returns the Sellers field if non-nil, zero value otherwise.

### GetSellersOk

`func (o *CategoryByDate) GetSellersOk() (*int32, bool)`

GetSellersOk returns a tuple with the Sellers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellers

`func (o *CategoryByDate) SetSellers(v int32)`

SetSellers sets Sellers field to given value.


### GetAvgPrice

`func (o *CategoryByDate) GetAvgPrice() float32`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *CategoryByDate) GetAvgPriceOk() (*float32, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *CategoryByDate) SetAvgPrice(v float32)`

SetAvgPrice sets AvgPrice field to given value.


### GetAvgSalePrice

`func (o *CategoryByDate) GetAvgSalePrice() float32`

GetAvgSalePrice returns the AvgSalePrice field if non-nil, zero value otherwise.

### GetAvgSalePriceOk

`func (o *CategoryByDate) GetAvgSalePriceOk() (*float32, bool)`

GetAvgSalePriceOk returns a tuple with the AvgSalePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSalePrice

`func (o *CategoryByDate) SetAvgSalePrice(v float32)`

SetAvgSalePrice sets AvgSalePrice field to given value.


### GetItems

`func (o *CategoryByDate) GetItems() int32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CategoryByDate) GetItemsOk() (*int32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CategoryByDate) SetItems(v int32)`

SetItems sets Items field to given value.


### GetSales

`func (o *CategoryByDate) GetSales() int32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *CategoryByDate) GetSalesOk() (*int32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *CategoryByDate) SetSales(v int32)`

SetSales sets Sales field to given value.


### GetRevenue

`func (o *CategoryByDate) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CategoryByDate) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CategoryByDate) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.


### GetComments

`func (o *CategoryByDate) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CategoryByDate) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CategoryByDate) SetComments(v float32)`

SetComments sets Comments field to given value.


### GetRating

`func (o *CategoryByDate) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *CategoryByDate) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *CategoryByDate) SetRating(v float32)`

SetRating sets Rating field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


