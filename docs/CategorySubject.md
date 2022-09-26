# CategorySubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sales** | **int64** | Количество продаж | 
**Revenue** | **int64** | Выручка | 
**Items** | **int64** | Количество товаров | 
**ItemsWithSells** | **int64** | Количество товаров с продажами | 
**ItemsWithSellsPercent** | **float32** | Процент товаров с продажами | 
**Sellers** | **int64** | Число продавцов | 
**SellersWithSells** | **int64** | Число продавцов с продажами | 
**Brands** | **int64** | Число брендов | 
**BrandsWithSells** | **int64** | Число брендов с продажами | 
**SellersWithSellsPercent** | **float32** | Процент продавцов с продажами | 
**SalesPerItemsAverage** | **float32** | Среднее количество продаж | 
**SalesPerItemsWithSellsAverage** | **float32** | Среднее количество продаж у товаров с продажами | 
**RevenuePerItemsAverage** | **float32** | Средняя выручка за товар | 
**RevenuePerItemsWithSellsAverage** | **float32** | Средняя выручка за товар с продажами | 
**AvgPrice** | **float32** | Средняя цена | 
**Rating** | **float32** | Рейтинг | 
**Comments** | **float32** | Количество комментариев | 
**Balance** | **int64** | Остаток | 
**LiveItems** | **int64** | Товаров с движением | 
**RevenuePotential** | **int64** | Потенциал | 
**LostProfit** | **int64** | Упущенная выручка | 
**LostProfitPercent** | **float32** | Процент упущенной выручки | 
**Graph** | **[]int32** | Граф | 
**Id** | **int32** | Идентификатор | 
**Name** | **string** | Наименование | 
**CommisionFbo** | **float32** | Комиссия FBO | 
**CommisionFbs** | **float32** | Комиссия FBS | 
**BasicLogistics** | **float32** | Базовая логистика | 
**StoragePrice** | **float32** | Цена хранения | 
**AcceptancePrice** | **float32** | Цена приемки | 
**DeliveryByVolume** | **string** |  | 
**Purchase** | **float32** | Средний процент выкупа | 
**PurchaseAfterReturn** | **float32** | Средний процент выкупа с учетом возврат | 

## Methods

### NewCategorySubject

`func NewCategorySubject(sales int64, revenue int64, items int64, itemsWithSells int64, itemsWithSellsPercent float32, sellers int64, sellersWithSells int64, brands int64, brandsWithSells int64, sellersWithSellsPercent float32, salesPerItemsAverage float32, salesPerItemsWithSellsAverage float32, revenuePerItemsAverage float32, revenuePerItemsWithSellsAverage float32, avgPrice float32, rating float32, comments float32, balance int64, liveItems int64, revenuePotential int64, lostProfit int64, lostProfitPercent float32, graph []int32, id int32, name string, commisionFbo float32, commisionFbs float32, basicLogistics float32, storagePrice float32, acceptancePrice float32, deliveryByVolume string, purchase float32, purchaseAfterReturn float32, ) *CategorySubject`

NewCategorySubject instantiates a new CategorySubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategorySubjectWithDefaults

`func NewCategorySubjectWithDefaults() *CategorySubject`

NewCategorySubjectWithDefaults instantiates a new CategorySubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSales

`func (o *CategorySubject) GetSales() int64`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *CategorySubject) GetSalesOk() (*int64, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *CategorySubject) SetSales(v int64)`

SetSales sets Sales field to given value.


### GetRevenue

`func (o *CategorySubject) GetRevenue() int64`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CategorySubject) GetRevenueOk() (*int64, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CategorySubject) SetRevenue(v int64)`

SetRevenue sets Revenue field to given value.


### GetItems

`func (o *CategorySubject) GetItems() int64`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CategorySubject) GetItemsOk() (*int64, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CategorySubject) SetItems(v int64)`

SetItems sets Items field to given value.


### GetItemsWithSells

`func (o *CategorySubject) GetItemsWithSells() int64`

GetItemsWithSells returns the ItemsWithSells field if non-nil, zero value otherwise.

### GetItemsWithSellsOk

`func (o *CategorySubject) GetItemsWithSellsOk() (*int64, bool)`

GetItemsWithSellsOk returns a tuple with the ItemsWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsWithSells

`func (o *CategorySubject) SetItemsWithSells(v int64)`

SetItemsWithSells sets ItemsWithSells field to given value.


### GetItemsWithSellsPercent

`func (o *CategorySubject) GetItemsWithSellsPercent() float32`

GetItemsWithSellsPercent returns the ItemsWithSellsPercent field if non-nil, zero value otherwise.

### GetItemsWithSellsPercentOk

`func (o *CategorySubject) GetItemsWithSellsPercentOk() (*float32, bool)`

GetItemsWithSellsPercentOk returns a tuple with the ItemsWithSellsPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsWithSellsPercent

`func (o *CategorySubject) SetItemsWithSellsPercent(v float32)`

SetItemsWithSellsPercent sets ItemsWithSellsPercent field to given value.


### GetSellers

`func (o *CategorySubject) GetSellers() int64`

GetSellers returns the Sellers field if non-nil, zero value otherwise.

### GetSellersOk

`func (o *CategorySubject) GetSellersOk() (*int64, bool)`

GetSellersOk returns a tuple with the Sellers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellers

`func (o *CategorySubject) SetSellers(v int64)`

SetSellers sets Sellers field to given value.


### GetSellersWithSells

`func (o *CategorySubject) GetSellersWithSells() int64`

GetSellersWithSells returns the SellersWithSells field if non-nil, zero value otherwise.

### GetSellersWithSellsOk

`func (o *CategorySubject) GetSellersWithSellsOk() (*int64, bool)`

GetSellersWithSellsOk returns a tuple with the SellersWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellersWithSells

`func (o *CategorySubject) SetSellersWithSells(v int64)`

SetSellersWithSells sets SellersWithSells field to given value.


### GetBrands

`func (o *CategorySubject) GetBrands() int64`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *CategorySubject) GetBrandsOk() (*int64, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *CategorySubject) SetBrands(v int64)`

SetBrands sets Brands field to given value.


### GetBrandsWithSells

`func (o *CategorySubject) GetBrandsWithSells() int64`

GetBrandsWithSells returns the BrandsWithSells field if non-nil, zero value otherwise.

### GetBrandsWithSellsOk

`func (o *CategorySubject) GetBrandsWithSellsOk() (*int64, bool)`

GetBrandsWithSellsOk returns a tuple with the BrandsWithSells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandsWithSells

`func (o *CategorySubject) SetBrandsWithSells(v int64)`

SetBrandsWithSells sets BrandsWithSells field to given value.


### GetSellersWithSellsPercent

`func (o *CategorySubject) GetSellersWithSellsPercent() float32`

GetSellersWithSellsPercent returns the SellersWithSellsPercent field if non-nil, zero value otherwise.

### GetSellersWithSellsPercentOk

`func (o *CategorySubject) GetSellersWithSellsPercentOk() (*float32, bool)`

GetSellersWithSellsPercentOk returns a tuple with the SellersWithSellsPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellersWithSellsPercent

`func (o *CategorySubject) SetSellersWithSellsPercent(v float32)`

SetSellersWithSellsPercent sets SellersWithSellsPercent field to given value.


### GetSalesPerItemsAverage

`func (o *CategorySubject) GetSalesPerItemsAverage() float32`

GetSalesPerItemsAverage returns the SalesPerItemsAverage field if non-nil, zero value otherwise.

### GetSalesPerItemsAverageOk

`func (o *CategorySubject) GetSalesPerItemsAverageOk() (*float32, bool)`

GetSalesPerItemsAverageOk returns a tuple with the SalesPerItemsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPerItemsAverage

`func (o *CategorySubject) SetSalesPerItemsAverage(v float32)`

SetSalesPerItemsAverage sets SalesPerItemsAverage field to given value.


### GetSalesPerItemsWithSellsAverage

`func (o *CategorySubject) GetSalesPerItemsWithSellsAverage() float32`

GetSalesPerItemsWithSellsAverage returns the SalesPerItemsWithSellsAverage field if non-nil, zero value otherwise.

### GetSalesPerItemsWithSellsAverageOk

`func (o *CategorySubject) GetSalesPerItemsWithSellsAverageOk() (*float32, bool)`

GetSalesPerItemsWithSellsAverageOk returns a tuple with the SalesPerItemsWithSellsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPerItemsWithSellsAverage

`func (o *CategorySubject) SetSalesPerItemsWithSellsAverage(v float32)`

SetSalesPerItemsWithSellsAverage sets SalesPerItemsWithSellsAverage field to given value.


### GetRevenuePerItemsAverage

`func (o *CategorySubject) GetRevenuePerItemsAverage() float32`

GetRevenuePerItemsAverage returns the RevenuePerItemsAverage field if non-nil, zero value otherwise.

### GetRevenuePerItemsAverageOk

`func (o *CategorySubject) GetRevenuePerItemsAverageOk() (*float32, bool)`

GetRevenuePerItemsAverageOk returns a tuple with the RevenuePerItemsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePerItemsAverage

`func (o *CategorySubject) SetRevenuePerItemsAverage(v float32)`

SetRevenuePerItemsAverage sets RevenuePerItemsAverage field to given value.


### GetRevenuePerItemsWithSellsAverage

`func (o *CategorySubject) GetRevenuePerItemsWithSellsAverage() float32`

GetRevenuePerItemsWithSellsAverage returns the RevenuePerItemsWithSellsAverage field if non-nil, zero value otherwise.

### GetRevenuePerItemsWithSellsAverageOk

`func (o *CategorySubject) GetRevenuePerItemsWithSellsAverageOk() (*float32, bool)`

GetRevenuePerItemsWithSellsAverageOk returns a tuple with the RevenuePerItemsWithSellsAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePerItemsWithSellsAverage

`func (o *CategorySubject) SetRevenuePerItemsWithSellsAverage(v float32)`

SetRevenuePerItemsWithSellsAverage sets RevenuePerItemsWithSellsAverage field to given value.


### GetAvgPrice

`func (o *CategorySubject) GetAvgPrice() float32`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *CategorySubject) GetAvgPriceOk() (*float32, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *CategorySubject) SetAvgPrice(v float32)`

SetAvgPrice sets AvgPrice field to given value.


### GetRating

`func (o *CategorySubject) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *CategorySubject) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *CategorySubject) SetRating(v float32)`

SetRating sets Rating field to given value.


### GetComments

`func (o *CategorySubject) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CategorySubject) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CategorySubject) SetComments(v float32)`

SetComments sets Comments field to given value.


### GetBalance

`func (o *CategorySubject) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CategorySubject) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CategorySubject) SetBalance(v int64)`

SetBalance sets Balance field to given value.


### GetLiveItems

`func (o *CategorySubject) GetLiveItems() int64`

GetLiveItems returns the LiveItems field if non-nil, zero value otherwise.

### GetLiveItemsOk

`func (o *CategorySubject) GetLiveItemsOk() (*int64, bool)`

GetLiveItemsOk returns a tuple with the LiveItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveItems

`func (o *CategorySubject) SetLiveItems(v int64)`

SetLiveItems sets LiveItems field to given value.


### GetRevenuePotential

`func (o *CategorySubject) GetRevenuePotential() int64`

GetRevenuePotential returns the RevenuePotential field if non-nil, zero value otherwise.

### GetRevenuePotentialOk

`func (o *CategorySubject) GetRevenuePotentialOk() (*int64, bool)`

GetRevenuePotentialOk returns a tuple with the RevenuePotential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePotential

`func (o *CategorySubject) SetRevenuePotential(v int64)`

SetRevenuePotential sets RevenuePotential field to given value.


### GetLostProfit

`func (o *CategorySubject) GetLostProfit() int64`

GetLostProfit returns the LostProfit field if non-nil, zero value otherwise.

### GetLostProfitOk

`func (o *CategorySubject) GetLostProfitOk() (*int64, bool)`

GetLostProfitOk returns a tuple with the LostProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostProfit

`func (o *CategorySubject) SetLostProfit(v int64)`

SetLostProfit sets LostProfit field to given value.


### GetLostProfitPercent

`func (o *CategorySubject) GetLostProfitPercent() float32`

GetLostProfitPercent returns the LostProfitPercent field if non-nil, zero value otherwise.

### GetLostProfitPercentOk

`func (o *CategorySubject) GetLostProfitPercentOk() (*float32, bool)`

GetLostProfitPercentOk returns a tuple with the LostProfitPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostProfitPercent

`func (o *CategorySubject) SetLostProfitPercent(v float32)`

SetLostProfitPercent sets LostProfitPercent field to given value.


### GetGraph

`func (o *CategorySubject) GetGraph() []int32`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *CategorySubject) GetGraphOk() (*[]int32, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *CategorySubject) SetGraph(v []int32)`

SetGraph sets Graph field to given value.


### GetId

`func (o *CategorySubject) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CategorySubject) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CategorySubject) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CategorySubject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategorySubject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategorySubject) SetName(v string)`

SetName sets Name field to given value.


### GetCommisionFbo

`func (o *CategorySubject) GetCommisionFbo() float32`

GetCommisionFbo returns the CommisionFbo field if non-nil, zero value otherwise.

### GetCommisionFboOk

`func (o *CategorySubject) GetCommisionFboOk() (*float32, bool)`

GetCommisionFboOk returns a tuple with the CommisionFbo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommisionFbo

`func (o *CategorySubject) SetCommisionFbo(v float32)`

SetCommisionFbo sets CommisionFbo field to given value.


### GetCommisionFbs

`func (o *CategorySubject) GetCommisionFbs() float32`

GetCommisionFbs returns the CommisionFbs field if non-nil, zero value otherwise.

### GetCommisionFbsOk

`func (o *CategorySubject) GetCommisionFbsOk() (*float32, bool)`

GetCommisionFbsOk returns a tuple with the CommisionFbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommisionFbs

`func (o *CategorySubject) SetCommisionFbs(v float32)`

SetCommisionFbs sets CommisionFbs field to given value.


### GetBasicLogistics

`func (o *CategorySubject) GetBasicLogistics() float32`

GetBasicLogistics returns the BasicLogistics field if non-nil, zero value otherwise.

### GetBasicLogisticsOk

`func (o *CategorySubject) GetBasicLogisticsOk() (*float32, bool)`

GetBasicLogisticsOk returns a tuple with the BasicLogistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicLogistics

`func (o *CategorySubject) SetBasicLogistics(v float32)`

SetBasicLogistics sets BasicLogistics field to given value.


### GetStoragePrice

`func (o *CategorySubject) GetStoragePrice() float32`

GetStoragePrice returns the StoragePrice field if non-nil, zero value otherwise.

### GetStoragePriceOk

`func (o *CategorySubject) GetStoragePriceOk() (*float32, bool)`

GetStoragePriceOk returns a tuple with the StoragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePrice

`func (o *CategorySubject) SetStoragePrice(v float32)`

SetStoragePrice sets StoragePrice field to given value.


### GetAcceptancePrice

`func (o *CategorySubject) GetAcceptancePrice() float32`

GetAcceptancePrice returns the AcceptancePrice field if non-nil, zero value otherwise.

### GetAcceptancePriceOk

`func (o *CategorySubject) GetAcceptancePriceOk() (*float32, bool)`

GetAcceptancePriceOk returns a tuple with the AcceptancePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptancePrice

`func (o *CategorySubject) SetAcceptancePrice(v float32)`

SetAcceptancePrice sets AcceptancePrice field to given value.


### GetDeliveryByVolume

`func (o *CategorySubject) GetDeliveryByVolume() string`

GetDeliveryByVolume returns the DeliveryByVolume field if non-nil, zero value otherwise.

### GetDeliveryByVolumeOk

`func (o *CategorySubject) GetDeliveryByVolumeOk() (*string, bool)`

GetDeliveryByVolumeOk returns a tuple with the DeliveryByVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryByVolume

`func (o *CategorySubject) SetDeliveryByVolume(v string)`

SetDeliveryByVolume sets DeliveryByVolume field to given value.


### GetPurchase

`func (o *CategorySubject) GetPurchase() float32`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *CategorySubject) GetPurchaseOk() (*float32, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *CategorySubject) SetPurchase(v float32)`

SetPurchase sets Purchase field to given value.


### GetPurchaseAfterReturn

`func (o *CategorySubject) GetPurchaseAfterReturn() float32`

GetPurchaseAfterReturn returns the PurchaseAfterReturn field if non-nil, zero value otherwise.

### GetPurchaseAfterReturnOk

`func (o *CategorySubject) GetPurchaseAfterReturnOk() (*float32, bool)`

GetPurchaseAfterReturnOk returns a tuple with the PurchaseAfterReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseAfterReturn

`func (o *CategorySubject) SetPurchaseAfterReturn(v float32)`

SetPurchaseAfterReturn sets PurchaseAfterReturn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


