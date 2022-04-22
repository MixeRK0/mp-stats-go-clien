# CategoryPriceSegmentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | **string** | Диапазон цен | 
**Items** | **float32** | Число товаров в диапазоне | 
**ItemsWithSells** | Pointer to **float32** | Число товаров с продажами в диапазоне | [optional] 
**Brands** | **float32** | Число брендов в диапазоне | 
**BrandsWithSells** | **float32** | Число брендов в диапазоне с продажами больше 0 | 
**Sellers** | **float32** | Число продавцов в диапазоне | 
**SellersWithSells** | **float32** | Число продавцов в диапазоне продажами больше 0 | 
**Revenue** | **float32** | Суммарная выручка в диапазоне | 
**Sales** | **float32** | Число продаж | 
**ProductRevenue** | **float32** | Сумма выручки диапазона, разделенная на число товаров в нем | 
**MinRangePrice** | **float32** | Цена в диапазоне от | 
**MaxRangePrice** | **float32** | Цена в диапазоне до | 
**LostProfit** | Pointer to **float32** | Упущенная выручка | [optional] 
**LostProfitPercent** | Pointer to **float32** | Процент упущенной выручка | [optional] 

## Methods

### NewCategoryPriceSegmentation

`func NewCategoryPriceSegmentation(range_ string, items float32, brands float32, brandsWithSells float32, sellers float32, sellersWithSells float32, revenue float32, sales float32, productRevenue float32, minRangePrice float32, maxRangePrice float32, ) *CategoryPriceSegmentation`

NewCategoryPriceSegmentation instantiates a new CategoryPriceSegmentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryPriceSegmentationWithDefaults

`func NewCategoryPriceSegmentationWithDefaults() *CategoryPriceSegmentation`

NewCategoryPriceSegmentationWithDefaults instantiates a new CategoryPriceSegmentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *CategoryPriceSegmentation) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CategoryPriceSegmentation) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CategoryPriceSegmentation) SetRange(v string)`

SetRange sets Range field to given value.


### GetItems

`func (o *CategoryPriceSegmentation) GetItems() float32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CategoryPriceSegmentation) GetItemsOk() (*float32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CategoryPriceSegmentation) SetItems(v float32)`

SetItems sets Items field to given value.


### GetItemsWithSells

`func (o *CategoryPriceSegmentation) GetItemsWithSells() float32`

GetItemsWithSells returns the ItemsWithSells field if non-nil, zero value otherwise.

### GetItemsWithSellsOk

`func (o *CategoryPriceSegmentation) GetItemsWithSellsOk() (*float32, bool)`

GetItemsWithSellsOk returns a tuple with the ItemsWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsWithSells

`func (o *CategoryPriceSegmentation) SetItemsWithSells(v float32)`

SetItemsWithSells sets ItemsWithSells field to given value.

### HasItemsWithSells

`func (o *CategoryPriceSegmentation) HasItemsWithSells() bool`

HasItemsWithSells returns a boolean if a field has been set.

### GetBrands

`func (o *CategoryPriceSegmentation) GetBrands() float32`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *CategoryPriceSegmentation) GetBrandsOk() (*float32, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *CategoryPriceSegmentation) SetBrands(v float32)`

SetBrands sets Brands field to given value.


### GetBrandsWithSells

`func (o *CategoryPriceSegmentation) GetBrandsWithSells() float32`

GetBrandsWithSells returns the BrandsWithSells field if non-nil, zero value otherwise.

### GetBrandsWithSellsOk

`func (o *CategoryPriceSegmentation) GetBrandsWithSellsOk() (*float32, bool)`

GetBrandsWithSellsOk returns a tuple with the BrandsWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandsWithSells

`func (o *CategoryPriceSegmentation) SetBrandsWithSells(v float32)`

SetBrandsWithSells sets BrandsWithSells field to given value.


### GetSellers

`func (o *CategoryPriceSegmentation) GetSellers() float32`

GetSellers returns the Sellers field if non-nil, zero value otherwise.

### GetSellersOk

`func (o *CategoryPriceSegmentation) GetSellersOk() (*float32, bool)`

GetSellersOk returns a tuple with the Sellers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellers

`func (o *CategoryPriceSegmentation) SetSellers(v float32)`

SetSellers sets Sellers field to given value.


### GetSellersWithSells

`func (o *CategoryPriceSegmentation) GetSellersWithSells() float32`

GetSellersWithSells returns the SellersWithSells field if non-nil, zero value otherwise.

### GetSellersWithSellsOk

`func (o *CategoryPriceSegmentation) GetSellersWithSellsOk() (*float32, bool)`

GetSellersWithSellsOk returns a tuple with the SellersWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellersWithSells

`func (o *CategoryPriceSegmentation) SetSellersWithSells(v float32)`

SetSellersWithSells sets SellersWithSells field to given value.


### GetRevenue

`func (o *CategoryPriceSegmentation) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CategoryPriceSegmentation) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CategoryPriceSegmentation) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.


### GetSales

`func (o *CategoryPriceSegmentation) GetSales() float32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *CategoryPriceSegmentation) GetSalesOk() (*float32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *CategoryPriceSegmentation) SetSales(v float32)`

SetSales sets Sales field to given value.


### GetProductRevenue

`func (o *CategoryPriceSegmentation) GetProductRevenue() float32`

GetProductRevenue returns the ProductRevenue field if non-nil, zero value otherwise.

### GetProductRevenueOk

`func (o *CategoryPriceSegmentation) GetProductRevenueOk() (*float32, bool)`

GetProductRevenueOk returns a tuple with the ProductRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductRevenue

`func (o *CategoryPriceSegmentation) SetProductRevenue(v float32)`

SetProductRevenue sets ProductRevenue field to given value.


### GetMinRangePrice

`func (o *CategoryPriceSegmentation) GetMinRangePrice() float32`

GetMinRangePrice returns the MinRangePrice field if non-nil, zero value otherwise.

### GetMinRangePriceOk

`func (o *CategoryPriceSegmentation) GetMinRangePriceOk() (*float32, bool)`

GetMinRangePriceOk returns a tuple with the MinRangePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRangePrice

`func (o *CategoryPriceSegmentation) SetMinRangePrice(v float32)`

SetMinRangePrice sets MinRangePrice field to given value.


### GetMaxRangePrice

`func (o *CategoryPriceSegmentation) GetMaxRangePrice() float32`

GetMaxRangePrice returns the MaxRangePrice field if non-nil, zero value otherwise.

### GetMaxRangePriceOk

`func (o *CategoryPriceSegmentation) GetMaxRangePriceOk() (*float32, bool)`

GetMaxRangePriceOk returns a tuple with the MaxRangePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRangePrice

`func (o *CategoryPriceSegmentation) SetMaxRangePrice(v float32)`

SetMaxRangePrice sets MaxRangePrice field to given value.


### GetLostProfit

`func (o *CategoryPriceSegmentation) GetLostProfit() float32`

GetLostProfit returns the LostProfit field if non-nil, zero value otherwise.

### GetLostProfitOk

`func (o *CategoryPriceSegmentation) GetLostProfitOk() (*float32, bool)`

GetLostProfitOk returns a tuple with the LostProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostProfit

`func (o *CategoryPriceSegmentation) SetLostProfit(v float32)`

SetLostProfit sets LostProfit field to given value.

### HasLostProfit

`func (o *CategoryPriceSegmentation) HasLostProfit() bool`

HasLostProfit returns a boolean if a field has been set.

### GetLostProfitPercent

`func (o *CategoryPriceSegmentation) GetLostProfitPercent() float32`

GetLostProfitPercent returns the LostProfitPercent field if non-nil, zero value otherwise.

### GetLostProfitPercentOk

`func (o *CategoryPriceSegmentation) GetLostProfitPercentOk() (*float32, bool)`

GetLostProfitPercentOk returns a tuple with the LostProfitPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostProfitPercent

`func (o *CategoryPriceSegmentation) SetLostProfitPercent(v float32)`

SetLostProfitPercent sets LostProfitPercent field to given value.

### HasLostProfitPercent

`func (o *CategoryPriceSegmentation) HasLostProfitPercent() bool`

HasLostProfitPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


