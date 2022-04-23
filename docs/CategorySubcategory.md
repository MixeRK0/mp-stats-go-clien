# CategorySubcategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Вложенная рубрика | 
**Items** | **int32** | Число товаров | 
**ItemsWithSells** | **int32** | Число товаров с продажами | 
**Brands** | **int32** | Число брендов | 
**BrandsWithSells** | **int32** | Число брендов с продажами | 
**Sellers** | **int32** | Число продавцов | 
**SellersWithSells** | **int32** | Число продавцов с продажами | 
**Sales** | **int32** | Число зафиксированных продаж (единицы) | 
**Revenue** | **float32** | Сумма произведений числа проданных товаров на их стоимость | 
**AvgPrice** | **float32** | Средняя стоимость товара | 
**Comments** | **float32** | Среднее число комментариев | 
**Rating** | **float32** | Средний рейтинг | 
**ItemsWithSellsPercent** | **float32** | Число товаров с продажами в процентах | 
**BrandsWithSellsPercent** | **float32** | Число брендов с продажами в процентах | 
**SellersWithSellsPercent** | **float32** | Число продавцов с продажами в процентах | 
**SalesPerItemsAverage** | **float32** | Среднее количество продаж на товар | 
**SalesPerItemsWithSellsAverage** | **float32** | Среднее количество продаж на товар с продажами | 
**RevenuePerItemsAverage** | **float32** | Среднее выручка на товар | 
**RevenuePerItemsWithSellsAverage** | **float32** | Среднее выручка на товар с продажами | 

## Methods

### NewCategorySubcategory

`func NewCategorySubcategory(name string, items int32, itemsWithSells int32, brands int32, brandsWithSells int32, sellers int32, sellersWithSells int32, sales int32, revenue float32, avgPrice float32, comments float32, rating float32, itemsWithSellsPercent float32, brandsWithSellsPercent float32, sellersWithSellsPercent float32, salesPerItemsAverage float32, salesPerItemsWithSellsAverage float32, revenuePerItemsAverage float32, revenuePerItemsWithSellsAverage float32, ) *CategorySubcategory`

NewCategorySubcategory instantiates a new CategorySubcategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategorySubcategoryWithDefaults

`func NewCategorySubcategoryWithDefaults() *CategorySubcategory`

NewCategorySubcategoryWithDefaults instantiates a new CategorySubcategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CategorySubcategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategorySubcategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategorySubcategory) SetName(v string)`

SetName sets Name field to given value.


### GetItems

`func (o *CategorySubcategory) GetItems() int32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CategorySubcategory) GetItemsOk() (*int32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CategorySubcategory) SetItems(v int32)`

SetItems sets Items field to given value.


### GetItemsWithSells

`func (o *CategorySubcategory) GetItemsWithSells() int32`

GetItemsWithSells returns the ItemsWithSells field if non-nil, zero value otherwise.

### GetItemsWithSellsOk

`func (o *CategorySubcategory) GetItemsWithSellsOk() (*int32, bool)`

GetItemsWithSellsOk returns a tuple with the ItemsWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsWithSells

`func (o *CategorySubcategory) SetItemsWithSells(v int32)`

SetItemsWithSells sets ItemsWithSells field to given value.


### GetBrands

`func (o *CategorySubcategory) GetBrands() int32`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *CategorySubcategory) GetBrandsOk() (*int32, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *CategorySubcategory) SetBrands(v int32)`

SetBrands sets Brands field to given value.


### GetBrandsWithSells

`func (o *CategorySubcategory) GetBrandsWithSells() int32`

GetBrandsWithSells returns the BrandsWithSells field if non-nil, zero value otherwise.

### GetBrandsWithSellsOk

`func (o *CategorySubcategory) GetBrandsWithSellsOk() (*int32, bool)`

GetBrandsWithSellsOk returns a tuple with the BrandsWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandsWithSells

`func (o *CategorySubcategory) SetBrandsWithSells(v int32)`

SetBrandsWithSells sets BrandsWithSells field to given value.


### GetSellers

`func (o *CategorySubcategory) GetSellers() int32`

GetSellers returns the Sellers field if non-nil, zero value otherwise.

### GetSellersOk

`func (o *CategorySubcategory) GetSellersOk() (*int32, bool)`

GetSellersOk returns a tuple with the Sellers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellers

`func (o *CategorySubcategory) SetSellers(v int32)`

SetSellers sets Sellers field to given value.


### GetSellersWithSells

`func (o *CategorySubcategory) GetSellersWithSells() int32`

GetSellersWithSells returns the SellersWithSells field if non-nil, zero value otherwise.

### GetSellersWithSellsOk

`func (o *CategorySubcategory) GetSellersWithSellsOk() (*int32, bool)`

GetSellersWithSellsOk returns a tuple with the SellersWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellersWithSells

`func (o *CategorySubcategory) SetSellersWithSells(v int32)`

SetSellersWithSells sets SellersWithSells field to given value.


### GetSales

`func (o *CategorySubcategory) GetSales() int32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *CategorySubcategory) GetSalesOk() (*int32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *CategorySubcategory) SetSales(v int32)`

SetSales sets Sales field to given value.


### GetRevenue

`func (o *CategorySubcategory) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CategorySubcategory) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CategorySubcategory) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.


### GetAvgPrice

`func (o *CategorySubcategory) GetAvgPrice() float32`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *CategorySubcategory) GetAvgPriceOk() (*float32, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *CategorySubcategory) SetAvgPrice(v float32)`

SetAvgPrice sets AvgPrice field to given value.


### GetComments

`func (o *CategorySubcategory) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CategorySubcategory) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CategorySubcategory) SetComments(v float32)`

SetComments sets Comments field to given value.


### GetRating

`func (o *CategorySubcategory) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *CategorySubcategory) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *CategorySubcategory) SetRating(v float32)`

SetRating sets Rating field to given value.


### GetItemsWithSellsPercent

`func (o *CategorySubcategory) GetItemsWithSellsPercent() float32`

GetItemsWithSellsPercent returns the ItemsWithSellsPercent field if non-nil, zero value otherwise.

### GetItemsWithSellsPercentOk

`func (o *CategorySubcategory) GetItemsWithSellsPercentOk() (*float32, bool)`

GetItemsWithSellsPercentOk returns a tuple with the ItemsWithSellsPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsWithSellsPercent

`func (o *CategorySubcategory) SetItemsWithSellsPercent(v float32)`

SetItemsWithSellsPercent sets ItemsWithSellsPercent field to given value.


### GetBrandsWithSellsPercent

`func (o *CategorySubcategory) GetBrandsWithSellsPercent() float32`

GetBrandsWithSellsPercent returns the BrandsWithSellsPercent field if non-nil, zero value otherwise.

### GetBrandsWithSellsPercentOk

`func (o *CategorySubcategory) GetBrandsWithSellsPercentOk() (*float32, bool)`

GetBrandsWithSellsPercentOk returns a tuple with the BrandsWithSellsPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandsWithSellsPercent

`func (o *CategorySubcategory) SetBrandsWithSellsPercent(v float32)`

SetBrandsWithSellsPercent sets BrandsWithSellsPercent field to given value.


### GetSellersWithSellsPercent

`func (o *CategorySubcategory) GetSellersWithSellsPercent() float32`

GetSellersWithSellsPercent returns the SellersWithSellsPercent field if non-nil, zero value otherwise.

### GetSellersWithSellsPercentOk

`func (o *CategorySubcategory) GetSellersWithSellsPercentOk() (*float32, bool)`

GetSellersWithSellsPercentOk returns a tuple with the SellersWithSellsPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellersWithSellsPercent

`func (o *CategorySubcategory) SetSellersWithSellsPercent(v float32)`

SetSellersWithSellsPercent sets SellersWithSellsPercent field to given value.


### GetSalesPerItemsAverage

`func (o *CategorySubcategory) GetSalesPerItemsAverage() float32`

GetSalesPerItemsAverage returns the SalesPerItemsAverage field if non-nil, zero value otherwise.

### GetSalesPerItemsAverageOk

`func (o *CategorySubcategory) GetSalesPerItemsAverageOk() (*float32, bool)`

GetSalesPerItemsAverageOk returns a tuple with the SalesPerItemsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPerItemsAverage

`func (o *CategorySubcategory) SetSalesPerItemsAverage(v float32)`

SetSalesPerItemsAverage sets SalesPerItemsAverage field to given value.


### GetSalesPerItemsWithSellsAverage

`func (o *CategorySubcategory) GetSalesPerItemsWithSellsAverage() float32`

GetSalesPerItemsWithSellsAverage returns the SalesPerItemsWithSellsAverage field if non-nil, zero value otherwise.

### GetSalesPerItemsWithSellsAverageOk

`func (o *CategorySubcategory) GetSalesPerItemsWithSellsAverageOk() (*float32, bool)`

GetSalesPerItemsWithSellsAverageOk returns a tuple with the SalesPerItemsWithSellsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPerItemsWithSellsAverage

`func (o *CategorySubcategory) SetSalesPerItemsWithSellsAverage(v float32)`

SetSalesPerItemsWithSellsAverage sets SalesPerItemsWithSellsAverage field to given value.


### GetRevenuePerItemsAverage

`func (o *CategorySubcategory) GetRevenuePerItemsAverage() float32`

GetRevenuePerItemsAverage returns the RevenuePerItemsAverage field if non-nil, zero value otherwise.

### GetRevenuePerItemsAverageOk

`func (o *CategorySubcategory) GetRevenuePerItemsAverageOk() (*float32, bool)`

GetRevenuePerItemsAverageOk returns a tuple with the RevenuePerItemsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePerItemsAverage

`func (o *CategorySubcategory) SetRevenuePerItemsAverage(v float32)`

SetRevenuePerItemsAverage sets RevenuePerItemsAverage field to given value.


### GetRevenuePerItemsWithSellsAverage

`func (o *CategorySubcategory) GetRevenuePerItemsWithSellsAverage() float32`

GetRevenuePerItemsWithSellsAverage returns the RevenuePerItemsWithSellsAverage field if non-nil, zero value otherwise.

### GetRevenuePerItemsWithSellsAverageOk

`func (o *CategorySubcategory) GetRevenuePerItemsWithSellsAverageOk() (*float32, bool)`

GetRevenuePerItemsWithSellsAverageOk returns a tuple with the RevenuePerItemsWithSellsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePerItemsWithSellsAverage

`func (o *CategorySubcategory) SetRevenuePerItemsWithSellsAverage(v float32)`

SetRevenuePerItemsWithSellsAverage sets RevenuePerItemsWithSellsAverage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


