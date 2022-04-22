# CategoryTrend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Week** | **string** | Дата начала недели | 
**Sales** | **float32** | Число продаж | 
**Revenue** | **float32** | Суммарная выручка в рубрике | 
**ProductRevenue** | **float32** | Сумма выручки категории, разделенная на число товаров | 
**Items** | **float32** | Число товаров в рубрике на указанной неделе | 
**Brands** | **float32** | Число брендов в рубрике на указанной неделе | 
**Sellers** | **float32** |  Число продавцов в рубрике на указанной неделе | 

## Methods

### NewCategoryTrend

`func NewCategoryTrend(week string, sales float32, revenue float32, productRevenue float32, items float32, brands float32, sellers float32, ) *CategoryTrend`

NewCategoryTrend instantiates a new CategoryTrend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryTrendWithDefaults

`func NewCategoryTrendWithDefaults() *CategoryTrend`

NewCategoryTrendWithDefaults instantiates a new CategoryTrend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeek

`func (o *CategoryTrend) GetWeek() string`

GetWeek returns the Week field if non-nil, zero value otherwise.

### GetWeekOk

`func (o *CategoryTrend) GetWeekOk() (*string, bool)`

GetWeekOk returns a tuple with the Week field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeek

`func (o *CategoryTrend) SetWeek(v string)`

SetWeek sets Week field to given value.


### GetSales

`func (o *CategoryTrend) GetSales() float32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *CategoryTrend) GetSalesOk() (*float32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *CategoryTrend) SetSales(v float32)`

SetSales sets Sales field to given value.


### GetRevenue

`func (o *CategoryTrend) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CategoryTrend) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CategoryTrend) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.


### GetProductRevenue

`func (o *CategoryTrend) GetProductRevenue() float32`

GetProductRevenue returns the ProductRevenue field if non-nil, zero value otherwise.

### GetProductRevenueOk

`func (o *CategoryTrend) GetProductRevenueOk() (*float32, bool)`

GetProductRevenueOk returns a tuple with the ProductRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductRevenue

`func (o *CategoryTrend) SetProductRevenue(v float32)`

SetProductRevenue sets ProductRevenue field to given value.


### GetItems

`func (o *CategoryTrend) GetItems() float32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CategoryTrend) GetItemsOk() (*float32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CategoryTrend) SetItems(v float32)`

SetItems sets Items field to given value.


### GetBrands

`func (o *CategoryTrend) GetBrands() float32`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *CategoryTrend) GetBrandsOk() (*float32, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *CategoryTrend) SetBrands(v float32)`

SetBrands sets Brands field to given value.


### GetSellers

`func (o *CategoryTrend) GetSellers() float32`

GetSellers returns the Sellers field if non-nil, zero value otherwise.

### GetSellersOk

`func (o *CategoryTrend) GetSellersOk() (*float32, bool)`

GetSellersOk returns a tuple with the Sellers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellers

`func (o *CategoryTrend) SetSellers(v float32)`

SetSellers sets Sellers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


