# SearchCategoriesElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Наименование категории | 
**Items** | **int32** | Число товаров | 
**ItemsWithSells** | **int32** | Число товаров с продажами | 
**ItemsWithSellsPercent** | **float32** | Процент товаров с продажами | 
**Sellers** | **int32** | Число продавцов | 
**SellersWithSells** | **int32** | Число продавцов с продажами | 
**SellersWithSellsPercent** | **float32** | Процент продавцов с продажами | 
**SalesPerItemsAverage** | **float32** | Среднее количество продаж на товар | 
**SalesPerItemsWithSellsAverage** | **float32** | Среднее количество продаж на товар с продажами | 
**Sales** | **int32** | Число продаж | 
**Revenue** | **int32** | Выручка | 
**RevenuePerItemsAverage** | **float32** | Сердняя выручка на товар | 
**RevenuePerItemsWithSellsAverage** | **float32** | Сердняя выручка на товар с продажами | 
**AvgPrice** | **float32** | Сердняя цена товара | 
**Comments** | **float32** | Среднее количество комментариев | 
**Rating** | **float32** | Средний рейтинг товаров | 

## Methods

### NewSearchCategoriesElement

`func NewSearchCategoriesElement(name string, items int32, itemsWithSells int32, itemsWithSellsPercent float32, sellers int32, sellersWithSells int32, sellersWithSellsPercent float32, salesPerItemsAverage float32, salesPerItemsWithSellsAverage float32, sales int32, revenue int32, revenuePerItemsAverage float32, revenuePerItemsWithSellsAverage float32, avgPrice float32, comments float32, rating float32, ) *SearchCategoriesElement`

NewSearchCategoriesElement instantiates a new SearchCategoriesElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchCategoriesElementWithDefaults

`func NewSearchCategoriesElementWithDefaults() *SearchCategoriesElement`

NewSearchCategoriesElementWithDefaults instantiates a new SearchCategoriesElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SearchCategoriesElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchCategoriesElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchCategoriesElement) SetName(v string)`

SetName sets Name field to given value.


### GetItems

`func (o *SearchCategoriesElement) GetItems() int32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SearchCategoriesElement) GetItemsOk() (*int32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SearchCategoriesElement) SetItems(v int32)`

SetItems sets Items field to given value.


### GetItemsWithSells

`func (o *SearchCategoriesElement) GetItemsWithSells() int32`

GetItemsWithSells returns the ItemsWithSells field if non-nil, zero value otherwise.

### GetItemsWithSellsOk

`func (o *SearchCategoriesElement) GetItemsWithSellsOk() (*int32, bool)`

GetItemsWithSellsOk returns a tuple with the ItemsWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsWithSells

`func (o *SearchCategoriesElement) SetItemsWithSells(v int32)`

SetItemsWithSells sets ItemsWithSells field to given value.


### GetItemsWithSellsPercent

`func (o *SearchCategoriesElement) GetItemsWithSellsPercent() float32`

GetItemsWithSellsPercent returns the ItemsWithSellsPercent field if non-nil, zero value otherwise.

### GetItemsWithSellsPercentOk

`func (o *SearchCategoriesElement) GetItemsWithSellsPercentOk() (*float32, bool)`

GetItemsWithSellsPercentOk returns a tuple with the ItemsWithSellsPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsWithSellsPercent

`func (o *SearchCategoriesElement) SetItemsWithSellsPercent(v float32)`

SetItemsWithSellsPercent sets ItemsWithSellsPercent field to given value.


### GetSellers

`func (o *SearchCategoriesElement) GetSellers() int32`

GetSellers returns the Sellers field if non-nil, zero value otherwise.

### GetSellersOk

`func (o *SearchCategoriesElement) GetSellersOk() (*int32, bool)`

GetSellersOk returns a tuple with the Sellers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellers

`func (o *SearchCategoriesElement) SetSellers(v int32)`

SetSellers sets Sellers field to given value.


### GetSellersWithSells

`func (o *SearchCategoriesElement) GetSellersWithSells() int32`

GetSellersWithSells returns the SellersWithSells field if non-nil, zero value otherwise.

### GetSellersWithSellsOk

`func (o *SearchCategoriesElement) GetSellersWithSellsOk() (*int32, bool)`

GetSellersWithSellsOk returns a tuple with the SellersWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellersWithSells

`func (o *SearchCategoriesElement) SetSellersWithSells(v int32)`

SetSellersWithSells sets SellersWithSells field to given value.


### GetSellersWithSellsPercent

`func (o *SearchCategoriesElement) GetSellersWithSellsPercent() float32`

GetSellersWithSellsPercent returns the SellersWithSellsPercent field if non-nil, zero value otherwise.

### GetSellersWithSellsPercentOk

`func (o *SearchCategoriesElement) GetSellersWithSellsPercentOk() (*float32, bool)`

GetSellersWithSellsPercentOk returns a tuple with the SellersWithSellsPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellersWithSellsPercent

`func (o *SearchCategoriesElement) SetSellersWithSellsPercent(v float32)`

SetSellersWithSellsPercent sets SellersWithSellsPercent field to given value.


### GetSalesPerItemsAverage

`func (o *SearchCategoriesElement) GetSalesPerItemsAverage() float32`

GetSalesPerItemsAverage returns the SalesPerItemsAverage field if non-nil, zero value otherwise.

### GetSalesPerItemsAverageOk

`func (o *SearchCategoriesElement) GetSalesPerItemsAverageOk() (*float32, bool)`

GetSalesPerItemsAverageOk returns a tuple with the SalesPerItemsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPerItemsAverage

`func (o *SearchCategoriesElement) SetSalesPerItemsAverage(v float32)`

SetSalesPerItemsAverage sets SalesPerItemsAverage field to given value.


### GetSalesPerItemsWithSellsAverage

`func (o *SearchCategoriesElement) GetSalesPerItemsWithSellsAverage() float32`

GetSalesPerItemsWithSellsAverage returns the SalesPerItemsWithSellsAverage field if non-nil, zero value otherwise.

### GetSalesPerItemsWithSellsAverageOk

`func (o *SearchCategoriesElement) GetSalesPerItemsWithSellsAverageOk() (*float32, bool)`

GetSalesPerItemsWithSellsAverageOk returns a tuple with the SalesPerItemsWithSellsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPerItemsWithSellsAverage

`func (o *SearchCategoriesElement) SetSalesPerItemsWithSellsAverage(v float32)`

SetSalesPerItemsWithSellsAverage sets SalesPerItemsWithSellsAverage field to given value.


### GetSales

`func (o *SearchCategoriesElement) GetSales() int32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *SearchCategoriesElement) GetSalesOk() (*int32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *SearchCategoriesElement) SetSales(v int32)`

SetSales sets Sales field to given value.


### GetRevenue

`func (o *SearchCategoriesElement) GetRevenue() int32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *SearchCategoriesElement) GetRevenueOk() (*int32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *SearchCategoriesElement) SetRevenue(v int32)`

SetRevenue sets Revenue field to given value.


### GetRevenuePerItemsAverage

`func (o *SearchCategoriesElement) GetRevenuePerItemsAverage() float32`

GetRevenuePerItemsAverage returns the RevenuePerItemsAverage field if non-nil, zero value otherwise.

### GetRevenuePerItemsAverageOk

`func (o *SearchCategoriesElement) GetRevenuePerItemsAverageOk() (*float32, bool)`

GetRevenuePerItemsAverageOk returns a tuple with the RevenuePerItemsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePerItemsAverage

`func (o *SearchCategoriesElement) SetRevenuePerItemsAverage(v float32)`

SetRevenuePerItemsAverage sets RevenuePerItemsAverage field to given value.


### GetRevenuePerItemsWithSellsAverage

`func (o *SearchCategoriesElement) GetRevenuePerItemsWithSellsAverage() float32`

GetRevenuePerItemsWithSellsAverage returns the RevenuePerItemsWithSellsAverage field if non-nil, zero value otherwise.

### GetRevenuePerItemsWithSellsAverageOk

`func (o *SearchCategoriesElement) GetRevenuePerItemsWithSellsAverageOk() (*float32, bool)`

GetRevenuePerItemsWithSellsAverageOk returns a tuple with the RevenuePerItemsWithSellsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePerItemsWithSellsAverage

`func (o *SearchCategoriesElement) SetRevenuePerItemsWithSellsAverage(v float32)`

SetRevenuePerItemsWithSellsAverage sets RevenuePerItemsWithSellsAverage field to given value.


### GetAvgPrice

`func (o *SearchCategoriesElement) GetAvgPrice() float32`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *SearchCategoriesElement) GetAvgPriceOk() (*float32, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *SearchCategoriesElement) SetAvgPrice(v float32)`

SetAvgPrice sets AvgPrice field to given value.


### GetComments

`func (o *SearchCategoriesElement) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *SearchCategoriesElement) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *SearchCategoriesElement) SetComments(v float32)`

SetComments sets Comments field to given value.


### GetRating

`func (o *SearchCategoriesElement) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *SearchCategoriesElement) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *SearchCategoriesElement) SetRating(v float32)`

SetRating sets Rating field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


