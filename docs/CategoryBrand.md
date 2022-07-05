# CategoryBrand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | **string** | Название бренда | 
**Items** | **int32** | Число товаров | 
**ItemsWithSells** | **int32** | Число товаров с продажами | 
**ItemsWithSellsPercent** | **float32** | Число товаров с продажами в процентах | 
**Sellers** | **int32** | Число продавцов | 
**SellersWithSells** | **int32** | Число продавцов с продажами | 
**SellersWithSellsPercent** | **float32** | Число продавцов с продажами в процентах | 
**Sales** | **int32** | Число зафиксированных продаж (единицы) | 
**Revenue** | **float32** | Сумма произведений числа проданных товаров на их стоимость | 
**SalesPerItemsAverage** | **float32** | Среднее количество продаж на товарную позицию | 
**SalesPerItemsWithSellsAverage** | **float32** | Среднее количество продаж на товарную позицию с продажами | 
**RevenuePerItemsAverage** | **float32** | Средняя выручка за товар | 
**RevenuePerItemsWithSellsAverage** | **float32** | Средняя выручка за товар с продажами | 
**Name** | **string** | Название бренда | 
**Balance** | **float32** | Баланс | 
**AvgPrice** | **float32** | Средняя стоимость товара | 
**Rating** | **float32** | Средний рейтинг | 
**Comments** | **float32** | Среднее число комментариев | 
**Position** | **int32** | Позиция в рейтинге | 
**Graph** | **[]interface{}** | Граф | 

## Methods

### NewCategoryBrand

`func NewCategoryBrand(brand string, items int32, itemsWithSells int32, itemsWithSellsPercent float32, sellers int32, sellersWithSells int32, sellersWithSellsPercent float32, sales int32, revenue float32, salesPerItemsAverage float32, salesPerItemsWithSellsAverage float32, revenuePerItemsAverage float32, revenuePerItemsWithSellsAverage float32, name string, balance float32, avgPrice float32, rating float32, comments float32, position int32, graph []interface{}, ) *CategoryBrand`

NewCategoryBrand instantiates a new CategoryBrand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryBrandWithDefaults

`func NewCategoryBrandWithDefaults() *CategoryBrand`

NewCategoryBrandWithDefaults instantiates a new CategoryBrand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *CategoryBrand) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CategoryBrand) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CategoryBrand) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetItems

`func (o *CategoryBrand) GetItems() int32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CategoryBrand) GetItemsOk() (*int32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CategoryBrand) SetItems(v int32)`

SetItems sets Items field to given value.


### GetItemsWithSells

`func (o *CategoryBrand) GetItemsWithSells() int32`

GetItemsWithSells returns the ItemsWithSells field if non-nil, zero value otherwise.

### GetItemsWithSellsOk

`func (o *CategoryBrand) GetItemsWithSellsOk() (*int32, bool)`

GetItemsWithSellsOk returns a tuple with the ItemsWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsWithSells

`func (o *CategoryBrand) SetItemsWithSells(v int32)`

SetItemsWithSells sets ItemsWithSells field to given value.


### GetItemsWithSellsPercent

`func (o *CategoryBrand) GetItemsWithSellsPercent() float32`

GetItemsWithSellsPercent returns the ItemsWithSellsPercent field if non-nil, zero value otherwise.

### GetItemsWithSellsPercentOk

`func (o *CategoryBrand) GetItemsWithSellsPercentOk() (*float32, bool)`

GetItemsWithSellsPercentOk returns a tuple with the ItemsWithSellsPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsWithSellsPercent

`func (o *CategoryBrand) SetItemsWithSellsPercent(v float32)`

SetItemsWithSellsPercent sets ItemsWithSellsPercent field to given value.


### GetSellers

`func (o *CategoryBrand) GetSellers() int32`

GetSellers returns the Sellers field if non-nil, zero value otherwise.

### GetSellersOk

`func (o *CategoryBrand) GetSellersOk() (*int32, bool)`

GetSellersOk returns a tuple with the Sellers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellers

`func (o *CategoryBrand) SetSellers(v int32)`

SetSellers sets Sellers field to given value.


### GetSellersWithSells

`func (o *CategoryBrand) GetSellersWithSells() int32`

GetSellersWithSells returns the SellersWithSells field if non-nil, zero value otherwise.

### GetSellersWithSellsOk

`func (o *CategoryBrand) GetSellersWithSellsOk() (*int32, bool)`

GetSellersWithSellsOk returns a tuple with the SellersWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellersWithSells

`func (o *CategoryBrand) SetSellersWithSells(v int32)`

SetSellersWithSells sets SellersWithSells field to given value.


### GetSellersWithSellsPercent

`func (o *CategoryBrand) GetSellersWithSellsPercent() float32`

GetSellersWithSellsPercent returns the SellersWithSellsPercent field if non-nil, zero value otherwise.

### GetSellersWithSellsPercentOk

`func (o *CategoryBrand) GetSellersWithSellsPercentOk() (*float32, bool)`

GetSellersWithSellsPercentOk returns a tuple with the SellersWithSellsPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellersWithSellsPercent

`func (o *CategoryBrand) SetSellersWithSellsPercent(v float32)`

SetSellersWithSellsPercent sets SellersWithSellsPercent field to given value.


### GetSales

`func (o *CategoryBrand) GetSales() int32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *CategoryBrand) GetSalesOk() (*int32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *CategoryBrand) SetSales(v int32)`

SetSales sets Sales field to given value.


### GetRevenue

`func (o *CategoryBrand) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CategoryBrand) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CategoryBrand) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.


### GetSalesPerItemsAverage

`func (o *CategoryBrand) GetSalesPerItemsAverage() float32`

GetSalesPerItemsAverage returns the SalesPerItemsAverage field if non-nil, zero value otherwise.

### GetSalesPerItemsAverageOk

`func (o *CategoryBrand) GetSalesPerItemsAverageOk() (*float32, bool)`

GetSalesPerItemsAverageOk returns a tuple with the SalesPerItemsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPerItemsAverage

`func (o *CategoryBrand) SetSalesPerItemsAverage(v float32)`

SetSalesPerItemsAverage sets SalesPerItemsAverage field to given value.


### GetSalesPerItemsWithSellsAverage

`func (o *CategoryBrand) GetSalesPerItemsWithSellsAverage() float32`

GetSalesPerItemsWithSellsAverage returns the SalesPerItemsWithSellsAverage field if non-nil, zero value otherwise.

### GetSalesPerItemsWithSellsAverageOk

`func (o *CategoryBrand) GetSalesPerItemsWithSellsAverageOk() (*float32, bool)`

GetSalesPerItemsWithSellsAverageOk returns a tuple with the SalesPerItemsWithSellsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPerItemsWithSellsAverage

`func (o *CategoryBrand) SetSalesPerItemsWithSellsAverage(v float32)`

SetSalesPerItemsWithSellsAverage sets SalesPerItemsWithSellsAverage field to given value.


### GetRevenuePerItemsAverage

`func (o *CategoryBrand) GetRevenuePerItemsAverage() float32`

GetRevenuePerItemsAverage returns the RevenuePerItemsAverage field if non-nil, zero value otherwise.

### GetRevenuePerItemsAverageOk

`func (o *CategoryBrand) GetRevenuePerItemsAverageOk() (*float32, bool)`

GetRevenuePerItemsAverageOk returns a tuple with the RevenuePerItemsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePerItemsAverage

`func (o *CategoryBrand) SetRevenuePerItemsAverage(v float32)`

SetRevenuePerItemsAverage sets RevenuePerItemsAverage field to given value.


### GetRevenuePerItemsWithSellsAverage

`func (o *CategoryBrand) GetRevenuePerItemsWithSellsAverage() float32`

GetRevenuePerItemsWithSellsAverage returns the RevenuePerItemsWithSellsAverage field if non-nil, zero value otherwise.

### GetRevenuePerItemsWithSellsAverageOk

`func (o *CategoryBrand) GetRevenuePerItemsWithSellsAverageOk() (*float32, bool)`

GetRevenuePerItemsWithSellsAverageOk returns a tuple with the RevenuePerItemsWithSellsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePerItemsWithSellsAverage

`func (o *CategoryBrand) SetRevenuePerItemsWithSellsAverage(v float32)`

SetRevenuePerItemsWithSellsAverage sets RevenuePerItemsWithSellsAverage field to given value.


### GetName

`func (o *CategoryBrand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryBrand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryBrand) SetName(v string)`

SetName sets Name field to given value.


### GetBalance

`func (o *CategoryBrand) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CategoryBrand) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CategoryBrand) SetBalance(v float32)`

SetBalance sets Balance field to given value.


### GetAvgPrice

`func (o *CategoryBrand) GetAvgPrice() float32`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *CategoryBrand) GetAvgPriceOk() (*float32, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *CategoryBrand) SetAvgPrice(v float32)`

SetAvgPrice sets AvgPrice field to given value.


### GetRating

`func (o *CategoryBrand) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *CategoryBrand) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *CategoryBrand) SetRating(v float32)`

SetRating sets Rating field to given value.


### GetComments

`func (o *CategoryBrand) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CategoryBrand) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CategoryBrand) SetComments(v float32)`

SetComments sets Comments field to given value.


### GetPosition

`func (o *CategoryBrand) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CategoryBrand) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CategoryBrand) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetGraph

`func (o *CategoryBrand) GetGraph() []interface{}`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *CategoryBrand) GetGraphOk() (*[]interface{}, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *CategoryBrand) SetGraph(v []interface{})`

SetGraph sets Graph field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


