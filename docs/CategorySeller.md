# CategorySeller

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupplierId** | **int32** | Идентификатор продавца | 
**Items** | **int32** | Число товаров | 
**ItemsWithSells** | **int32** | Число товаров с продажами | 
**ItemsWithSellsPercent** | **float32** | Процент товаров с продажами | 
**Brands** | **int32** | Число брендов | 
**BrandsWithSells** | **int32** | Число брендов с продажами | 
**BrandsWithSellsPercent** | **float32** | Процент брендов с продажами | 
**Sales** | **int32** | Число зафиксированных продаж (единицы) | 
**Revenue** | **float32** | Сумма произведений числа проданных товаров на их стоимость | 
**SalesPerItemsAverage** | **float32** | Среднее число продаж на товар | 
**SalesPerItemsWithSellsAverage** | **float32** | Среднее число продаж на товар с продажами | 
**RevenuePerItemsAverage** | **float32** | Средняя выручка на товар | 
**RevenuePerItemsWithSellsAverage** | **float32** | Средняя выручка на товар с продажами | 
**Name** | **string** | Название продавца | 
**Balance** | **float32** | Число остатков товаров на складах | 
**AvgPrice** | **float32** | Средняя цена | 
**Rating** | **float32** | Средний рейтинг | 
**Comments** | **float32** | Среднее число комментариев | 
**Position** | **int32** | Позиция продавца | 
**Graph** | **[]int32** | Позиция продавца | 

## Methods

### NewCategorySeller

`func NewCategorySeller(supplierId int32, items int32, itemsWithSells int32, itemsWithSellsPercent float32, brands int32, brandsWithSells int32, brandsWithSellsPercent float32, sales int32, revenue float32, salesPerItemsAverage float32, salesPerItemsWithSellsAverage float32, revenuePerItemsAverage float32, revenuePerItemsWithSellsAverage float32, name string, balance float32, avgPrice float32, rating float32, comments float32, position int32, graph []int32, ) *CategorySeller`

NewCategorySeller instantiates a new CategorySeller object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategorySellerWithDefaults

`func NewCategorySellerWithDefaults() *CategorySeller`

NewCategorySellerWithDefaults instantiates a new CategorySeller object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupplierId

`func (o *CategorySeller) GetSupplierId() int32`

GetSupplierId returns the SupplierId field if non-nil, zero value otherwise.

### GetSupplierIdOk

`func (o *CategorySeller) GetSupplierIdOk() (*int32, bool)`

GetSupplierIdOk returns a tuple with the SupplierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierId

`func (o *CategorySeller) SetSupplierId(v int32)`

SetSupplierId sets SupplierId field to given value.


### GetItems

`func (o *CategorySeller) GetItems() int32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CategorySeller) GetItemsOk() (*int32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CategorySeller) SetItems(v int32)`

SetItems sets Items field to given value.


### GetItemsWithSells

`func (o *CategorySeller) GetItemsWithSells() int32`

GetItemsWithSells returns the ItemsWithSells field if non-nil, zero value otherwise.

### GetItemsWithSellsOk

`func (o *CategorySeller) GetItemsWithSellsOk() (*int32, bool)`

GetItemsWithSellsOk returns a tuple with the ItemsWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsWithSells

`func (o *CategorySeller) SetItemsWithSells(v int32)`

SetItemsWithSells sets ItemsWithSells field to given value.


### GetItemsWithSellsPercent

`func (o *CategorySeller) GetItemsWithSellsPercent() float32`

GetItemsWithSellsPercent returns the ItemsWithSellsPercent field if non-nil, zero value otherwise.

### GetItemsWithSellsPercentOk

`func (o *CategorySeller) GetItemsWithSellsPercentOk() (*float32, bool)`

GetItemsWithSellsPercentOk returns a tuple with the ItemsWithSellsPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsWithSellsPercent

`func (o *CategorySeller) SetItemsWithSellsPercent(v float32)`

SetItemsWithSellsPercent sets ItemsWithSellsPercent field to given value.


### GetBrands

`func (o *CategorySeller) GetBrands() int32`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *CategorySeller) GetBrandsOk() (*int32, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *CategorySeller) SetBrands(v int32)`

SetBrands sets Brands field to given value.


### GetBrandsWithSells

`func (o *CategorySeller) GetBrandsWithSells() int32`

GetBrandsWithSells returns the BrandsWithSells field if non-nil, zero value otherwise.

### GetBrandsWithSellsOk

`func (o *CategorySeller) GetBrandsWithSellsOk() (*int32, bool)`

GetBrandsWithSellsOk returns a tuple with the BrandsWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandsWithSells

`func (o *CategorySeller) SetBrandsWithSells(v int32)`

SetBrandsWithSells sets BrandsWithSells field to given value.


### GetBrandsWithSellsPercent

`func (o *CategorySeller) GetBrandsWithSellsPercent() float32`

GetBrandsWithSellsPercent returns the BrandsWithSellsPercent field if non-nil, zero value otherwise.

### GetBrandsWithSellsPercentOk

`func (o *CategorySeller) GetBrandsWithSellsPercentOk() (*float32, bool)`

GetBrandsWithSellsPercentOk returns a tuple with the BrandsWithSellsPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandsWithSellsPercent

`func (o *CategorySeller) SetBrandsWithSellsPercent(v float32)`

SetBrandsWithSellsPercent sets BrandsWithSellsPercent field to given value.


### GetSales

`func (o *CategorySeller) GetSales() int32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *CategorySeller) GetSalesOk() (*int32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *CategorySeller) SetSales(v int32)`

SetSales sets Sales field to given value.


### GetRevenue

`func (o *CategorySeller) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CategorySeller) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CategorySeller) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.


### GetSalesPerItemsAverage

`func (o *CategorySeller) GetSalesPerItemsAverage() float32`

GetSalesPerItemsAverage returns the SalesPerItemsAverage field if non-nil, zero value otherwise.

### GetSalesPerItemsAverageOk

`func (o *CategorySeller) GetSalesPerItemsAverageOk() (*float32, bool)`

GetSalesPerItemsAverageOk returns a tuple with the SalesPerItemsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPerItemsAverage

`func (o *CategorySeller) SetSalesPerItemsAverage(v float32)`

SetSalesPerItemsAverage sets SalesPerItemsAverage field to given value.


### GetSalesPerItemsWithSellsAverage

`func (o *CategorySeller) GetSalesPerItemsWithSellsAverage() float32`

GetSalesPerItemsWithSellsAverage returns the SalesPerItemsWithSellsAverage field if non-nil, zero value otherwise.

### GetSalesPerItemsWithSellsAverageOk

`func (o *CategorySeller) GetSalesPerItemsWithSellsAverageOk() (*float32, bool)`

GetSalesPerItemsWithSellsAverageOk returns a tuple with the SalesPerItemsWithSellsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPerItemsWithSellsAverage

`func (o *CategorySeller) SetSalesPerItemsWithSellsAverage(v float32)`

SetSalesPerItemsWithSellsAverage sets SalesPerItemsWithSellsAverage field to given value.


### GetRevenuePerItemsAverage

`func (o *CategorySeller) GetRevenuePerItemsAverage() float32`

GetRevenuePerItemsAverage returns the RevenuePerItemsAverage field if non-nil, zero value otherwise.

### GetRevenuePerItemsAverageOk

`func (o *CategorySeller) GetRevenuePerItemsAverageOk() (*float32, bool)`

GetRevenuePerItemsAverageOk returns a tuple with the RevenuePerItemsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePerItemsAverage

`func (o *CategorySeller) SetRevenuePerItemsAverage(v float32)`

SetRevenuePerItemsAverage sets RevenuePerItemsAverage field to given value.


### GetRevenuePerItemsWithSellsAverage

`func (o *CategorySeller) GetRevenuePerItemsWithSellsAverage() float32`

GetRevenuePerItemsWithSellsAverage returns the RevenuePerItemsWithSellsAverage field if non-nil, zero value otherwise.

### GetRevenuePerItemsWithSellsAverageOk

`func (o *CategorySeller) GetRevenuePerItemsWithSellsAverageOk() (*float32, bool)`

GetRevenuePerItemsWithSellsAverageOk returns a tuple with the RevenuePerItemsWithSellsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePerItemsWithSellsAverage

`func (o *CategorySeller) SetRevenuePerItemsWithSellsAverage(v float32)`

SetRevenuePerItemsWithSellsAverage sets RevenuePerItemsWithSellsAverage field to given value.


### GetName

`func (o *CategorySeller) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategorySeller) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategorySeller) SetName(v string)`

SetName sets Name field to given value.


### GetBalance

`func (o *CategorySeller) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CategorySeller) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CategorySeller) SetBalance(v float32)`

SetBalance sets Balance field to given value.


### GetAvgPrice

`func (o *CategorySeller) GetAvgPrice() float32`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *CategorySeller) GetAvgPriceOk() (*float32, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *CategorySeller) SetAvgPrice(v float32)`

SetAvgPrice sets AvgPrice field to given value.


### GetRating

`func (o *CategorySeller) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *CategorySeller) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *CategorySeller) SetRating(v float32)`

SetRating sets Rating field to given value.


### GetComments

`func (o *CategorySeller) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CategorySeller) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CategorySeller) SetComments(v float32)`

SetComments sets Comments field to given value.


### GetPosition

`func (o *CategorySeller) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CategorySeller) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CategorySeller) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetGraph

`func (o *CategorySeller) GetGraph() []int32`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *CategorySeller) GetGraphOk() (*[]int32, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *CategorySeller) SetGraph(v []int32)`

SetGraph sets Graph field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


