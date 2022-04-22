# SellerPriceSegmentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | **string** | Диапазон цен | 
**Items** | **float32** | Число товаров в диапазоне | 
**Brands** | **float32** | Число брендов в диапазоне | 
**BrandsWithSells** | **float32** | Число брендов в диапазоне с продажами больше 0 | 
**Sellers** | **float32** | Число продавцов в диапазоне | 
**SellersWithSells** | **float32** | Число продавцов в диапазоне продажами больше 0 | 
**Revenue** | **float32** | Суммарная выручка в диапазоне | 
**Sales** | **float32** | Число продаж | 
**ProductRevenue** | **float32** | Сумма выручки диапазона, разделенная на число товаров в нем | 
**MinRangePrice** | **float32** | Цена в диапазоне от | 
**MaxRangePrice** | **float32** | Цена в диапазоне до | 

## Methods

### NewSellerPriceSegmentation

`func NewSellerPriceSegmentation(range_ string, items float32, brands float32, brandsWithSells float32, sellers float32, sellersWithSells float32, revenue float32, sales float32, productRevenue float32, minRangePrice float32, maxRangePrice float32, ) *SellerPriceSegmentation`

NewSellerPriceSegmentation instantiates a new SellerPriceSegmentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellerPriceSegmentationWithDefaults

`func NewSellerPriceSegmentationWithDefaults() *SellerPriceSegmentation`

NewSellerPriceSegmentationWithDefaults instantiates a new SellerPriceSegmentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *SellerPriceSegmentation) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *SellerPriceSegmentation) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *SellerPriceSegmentation) SetRange(v string)`

SetRange sets Range field to given value.


### GetItems

`func (o *SellerPriceSegmentation) GetItems() float32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SellerPriceSegmentation) GetItemsOk() (*float32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SellerPriceSegmentation) SetItems(v float32)`

SetItems sets Items field to given value.


### GetBrands

`func (o *SellerPriceSegmentation) GetBrands() float32`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *SellerPriceSegmentation) GetBrandsOk() (*float32, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *SellerPriceSegmentation) SetBrands(v float32)`

SetBrands sets Brands field to given value.


### GetBrandsWithSells

`func (o *SellerPriceSegmentation) GetBrandsWithSells() float32`

GetBrandsWithSells returns the BrandsWithSells field if non-nil, zero value otherwise.

### GetBrandsWithSellsOk

`func (o *SellerPriceSegmentation) GetBrandsWithSellsOk() (*float32, bool)`

GetBrandsWithSellsOk returns a tuple with the BrandsWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandsWithSells

`func (o *SellerPriceSegmentation) SetBrandsWithSells(v float32)`

SetBrandsWithSells sets BrandsWithSells field to given value.


### GetSellers

`func (o *SellerPriceSegmentation) GetSellers() float32`

GetSellers returns the Sellers field if non-nil, zero value otherwise.

### GetSellersOk

`func (o *SellerPriceSegmentation) GetSellersOk() (*float32, bool)`

GetSellersOk returns a tuple with the Sellers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellers

`func (o *SellerPriceSegmentation) SetSellers(v float32)`

SetSellers sets Sellers field to given value.


### GetSellersWithSells

`func (o *SellerPriceSegmentation) GetSellersWithSells() float32`

GetSellersWithSells returns the SellersWithSells field if non-nil, zero value otherwise.

### GetSellersWithSellsOk

`func (o *SellerPriceSegmentation) GetSellersWithSellsOk() (*float32, bool)`

GetSellersWithSellsOk returns a tuple with the SellersWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellersWithSells

`func (o *SellerPriceSegmentation) SetSellersWithSells(v float32)`

SetSellersWithSells sets SellersWithSells field to given value.


### GetRevenue

`func (o *SellerPriceSegmentation) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *SellerPriceSegmentation) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *SellerPriceSegmentation) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.


### GetSales

`func (o *SellerPriceSegmentation) GetSales() float32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *SellerPriceSegmentation) GetSalesOk() (*float32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *SellerPriceSegmentation) SetSales(v float32)`

SetSales sets Sales field to given value.


### GetProductRevenue

`func (o *SellerPriceSegmentation) GetProductRevenue() float32`

GetProductRevenue returns the ProductRevenue field if non-nil, zero value otherwise.

### GetProductRevenueOk

`func (o *SellerPriceSegmentation) GetProductRevenueOk() (*float32, bool)`

GetProductRevenueOk returns a tuple with the ProductRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductRevenue

`func (o *SellerPriceSegmentation) SetProductRevenue(v float32)`

SetProductRevenue sets ProductRevenue field to given value.


### GetMinRangePrice

`func (o *SellerPriceSegmentation) GetMinRangePrice() float32`

GetMinRangePrice returns the MinRangePrice field if non-nil, zero value otherwise.

### GetMinRangePriceOk

`func (o *SellerPriceSegmentation) GetMinRangePriceOk() (*float32, bool)`

GetMinRangePriceOk returns a tuple with the MinRangePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRangePrice

`func (o *SellerPriceSegmentation) SetMinRangePrice(v float32)`

SetMinRangePrice sets MinRangePrice field to given value.


### GetMaxRangePrice

`func (o *SellerPriceSegmentation) GetMaxRangePrice() float32`

GetMaxRangePrice returns the MaxRangePrice field if non-nil, zero value otherwise.

### GetMaxRangePriceOk

`func (o *SellerPriceSegmentation) GetMaxRangePriceOk() (*float32, bool)`

GetMaxRangePriceOk returns a tuple with the MaxRangePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRangePrice

`func (o *SellerPriceSegmentation) SetMaxRangePrice(v float32)`

SetMaxRangePrice sets MaxRangePrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


