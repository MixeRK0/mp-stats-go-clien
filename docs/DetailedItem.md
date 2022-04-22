# DetailedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Идентификатор товарной позиции | 
**Name** | **string** | Название товара | 
**Brand** | **string** | Название бренда | 
**Seller** | **string** | Название продавца | 
**Category** | **string** | Категория товара | 
**Color** | **string** | Цвет товара | 
**Balance** | **float32** | Последний зафиксированный остаток на складе | 
**Rating** | **float32** | Рейтинг товара | 
**Comments** | **float32** | Количество комментариев | 
**FinalPrice** | **float32** | Последняя зафиксированная цена | 
**FinalPriceMax** | **float32** | Максимальная цена за период | 
**FinalPriceMin** | **float32** | Минимальная цена за период | 
**FinalPriceAverage** | **float32** | Средняя цена за период (выручка / число продаж) | 
**BasicSale** | **float32** | Размер скидки | 
**BasicPrice** | **float32** | Цена после применения скидки | 
**PromoSale** | **float32** | Размер скидки по промокоду | 
**ClientSale** | **float32** | Размер Скидки Постоянного Покупателя | 
**ClientPrice** | **float32** | Итоговая цена для посетителя, с учетом СПП | 
**StartPrice** | **float32** | Цена без учета скидок | 
**Sales** | **float32** | Количество проданных единиц товара за период | 
**Revenue** | **float32** | Выручка за период | 
**RevenuePotential** | **float32** | Потенциал выручки (выручка / число дней в наличии) * дней в отчете | 
**LostProfit** | **float32** | Упущенная выручка (Потенциал - Выручка) | 
**DaysInStock** | **float32** | Количество дней, когда товар был в наличии на конец дня | 
**DaysWithSales** | **float32** | Количество дней, когда были продажи | 
**AverageIfInStock** | **float32** | Среднее число продаж, при наличии товара на конец дня | 
**Thumb** | **string** | Изображение товара | 
**Graph** | **[]float32** | График продаж | 

## Methods

### NewDetailedItem

`func NewDetailedItem(id float32, name string, brand string, seller string, category string, color string, balance float32, rating float32, comments float32, finalPrice float32, finalPriceMax float32, finalPriceMin float32, finalPriceAverage float32, basicSale float32, basicPrice float32, promoSale float32, clientSale float32, clientPrice float32, startPrice float32, sales float32, revenue float32, revenuePotential float32, lostProfit float32, daysInStock float32, daysWithSales float32, averageIfInStock float32, thumb string, graph []float32, ) *DetailedItem`

NewDetailedItem instantiates a new DetailedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedItemWithDefaults

`func NewDetailedItemWithDefaults() *DetailedItem`

NewDetailedItemWithDefaults instantiates a new DetailedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DetailedItem) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailedItem) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailedItem) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *DetailedItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailedItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailedItem) SetName(v string)`

SetName sets Name field to given value.


### GetBrand

`func (o *DetailedItem) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *DetailedItem) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *DetailedItem) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetSeller

`func (o *DetailedItem) GetSeller() string`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *DetailedItem) GetSellerOk() (*string, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *DetailedItem) SetSeller(v string)`

SetSeller sets Seller field to given value.


### GetCategory

`func (o *DetailedItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DetailedItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DetailedItem) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetColor

`func (o *DetailedItem) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DetailedItem) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DetailedItem) SetColor(v string)`

SetColor sets Color field to given value.


### GetBalance

`func (o *DetailedItem) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *DetailedItem) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *DetailedItem) SetBalance(v float32)`

SetBalance sets Balance field to given value.


### GetRating

`func (o *DetailedItem) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *DetailedItem) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *DetailedItem) SetRating(v float32)`

SetRating sets Rating field to given value.


### GetComments

`func (o *DetailedItem) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DetailedItem) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DetailedItem) SetComments(v float32)`

SetComments sets Comments field to given value.


### GetFinalPrice

`func (o *DetailedItem) GetFinalPrice() float32`

GetFinalPrice returns the FinalPrice field if non-nil, zero value otherwise.

### GetFinalPriceOk

`func (o *DetailedItem) GetFinalPriceOk() (*float32, bool)`

GetFinalPriceOk returns a tuple with the FinalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPrice

`func (o *DetailedItem) SetFinalPrice(v float32)`

SetFinalPrice sets FinalPrice field to given value.


### GetFinalPriceMax

`func (o *DetailedItem) GetFinalPriceMax() float32`

GetFinalPriceMax returns the FinalPriceMax field if non-nil, zero value otherwise.

### GetFinalPriceMaxOk

`func (o *DetailedItem) GetFinalPriceMaxOk() (*float32, bool)`

GetFinalPriceMaxOk returns a tuple with the FinalPriceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPriceMax

`func (o *DetailedItem) SetFinalPriceMax(v float32)`

SetFinalPriceMax sets FinalPriceMax field to given value.


### GetFinalPriceMin

`func (o *DetailedItem) GetFinalPriceMin() float32`

GetFinalPriceMin returns the FinalPriceMin field if non-nil, zero value otherwise.

### GetFinalPriceMinOk

`func (o *DetailedItem) GetFinalPriceMinOk() (*float32, bool)`

GetFinalPriceMinOk returns a tuple with the FinalPriceMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPriceMin

`func (o *DetailedItem) SetFinalPriceMin(v float32)`

SetFinalPriceMin sets FinalPriceMin field to given value.


### GetFinalPriceAverage

`func (o *DetailedItem) GetFinalPriceAverage() float32`

GetFinalPriceAverage returns the FinalPriceAverage field if non-nil, zero value otherwise.

### GetFinalPriceAverageOk

`func (o *DetailedItem) GetFinalPriceAverageOk() (*float32, bool)`

GetFinalPriceAverageOk returns a tuple with the FinalPriceAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPriceAverage

`func (o *DetailedItem) SetFinalPriceAverage(v float32)`

SetFinalPriceAverage sets FinalPriceAverage field to given value.


### GetBasicSale

`func (o *DetailedItem) GetBasicSale() float32`

GetBasicSale returns the BasicSale field if non-nil, zero value otherwise.

### GetBasicSaleOk

`func (o *DetailedItem) GetBasicSaleOk() (*float32, bool)`

GetBasicSaleOk returns a tuple with the BasicSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSale

`func (o *DetailedItem) SetBasicSale(v float32)`

SetBasicSale sets BasicSale field to given value.


### GetBasicPrice

`func (o *DetailedItem) GetBasicPrice() float32`

GetBasicPrice returns the BasicPrice field if non-nil, zero value otherwise.

### GetBasicPriceOk

`func (o *DetailedItem) GetBasicPriceOk() (*float32, bool)`

GetBasicPriceOk returns a tuple with the BasicPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicPrice

`func (o *DetailedItem) SetBasicPrice(v float32)`

SetBasicPrice sets BasicPrice field to given value.


### GetPromoSale

`func (o *DetailedItem) GetPromoSale() float32`

GetPromoSale returns the PromoSale field if non-nil, zero value otherwise.

### GetPromoSaleOk

`func (o *DetailedItem) GetPromoSaleOk() (*float32, bool)`

GetPromoSaleOk returns a tuple with the PromoSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoSale

`func (o *DetailedItem) SetPromoSale(v float32)`

SetPromoSale sets PromoSale field to given value.


### GetClientSale

`func (o *DetailedItem) GetClientSale() float32`

GetClientSale returns the ClientSale field if non-nil, zero value otherwise.

### GetClientSaleOk

`func (o *DetailedItem) GetClientSaleOk() (*float32, bool)`

GetClientSaleOk returns a tuple with the ClientSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSale

`func (o *DetailedItem) SetClientSale(v float32)`

SetClientSale sets ClientSale field to given value.


### GetClientPrice

`func (o *DetailedItem) GetClientPrice() float32`

GetClientPrice returns the ClientPrice field if non-nil, zero value otherwise.

### GetClientPriceOk

`func (o *DetailedItem) GetClientPriceOk() (*float32, bool)`

GetClientPriceOk returns a tuple with the ClientPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrice

`func (o *DetailedItem) SetClientPrice(v float32)`

SetClientPrice sets ClientPrice field to given value.


### GetStartPrice

`func (o *DetailedItem) GetStartPrice() float32`

GetStartPrice returns the StartPrice field if non-nil, zero value otherwise.

### GetStartPriceOk

`func (o *DetailedItem) GetStartPriceOk() (*float32, bool)`

GetStartPriceOk returns a tuple with the StartPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPrice

`func (o *DetailedItem) SetStartPrice(v float32)`

SetStartPrice sets StartPrice field to given value.


### GetSales

`func (o *DetailedItem) GetSales() float32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *DetailedItem) GetSalesOk() (*float32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *DetailedItem) SetSales(v float32)`

SetSales sets Sales field to given value.


### GetRevenue

`func (o *DetailedItem) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *DetailedItem) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *DetailedItem) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.


### GetRevenuePotential

`func (o *DetailedItem) GetRevenuePotential() float32`

GetRevenuePotential returns the RevenuePotential field if non-nil, zero value otherwise.

### GetRevenuePotentialOk

`func (o *DetailedItem) GetRevenuePotentialOk() (*float32, bool)`

GetRevenuePotentialOk returns a tuple with the RevenuePotential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePotential

`func (o *DetailedItem) SetRevenuePotential(v float32)`

SetRevenuePotential sets RevenuePotential field to given value.


### GetLostProfit

`func (o *DetailedItem) GetLostProfit() float32`

GetLostProfit returns the LostProfit field if non-nil, zero value otherwise.

### GetLostProfitOk

`func (o *DetailedItem) GetLostProfitOk() (*float32, bool)`

GetLostProfitOk returns a tuple with the LostProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostProfit

`func (o *DetailedItem) SetLostProfit(v float32)`

SetLostProfit sets LostProfit field to given value.


### GetDaysInStock

`func (o *DetailedItem) GetDaysInStock() float32`

GetDaysInStock returns the DaysInStock field if non-nil, zero value otherwise.

### GetDaysInStockOk

`func (o *DetailedItem) GetDaysInStockOk() (*float32, bool)`

GetDaysInStockOk returns a tuple with the DaysInStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInStock

`func (o *DetailedItem) SetDaysInStock(v float32)`

SetDaysInStock sets DaysInStock field to given value.


### GetDaysWithSales

`func (o *DetailedItem) GetDaysWithSales() float32`

GetDaysWithSales returns the DaysWithSales field if non-nil, zero value otherwise.

### GetDaysWithSalesOk

`func (o *DetailedItem) GetDaysWithSalesOk() (*float32, bool)`

GetDaysWithSalesOk returns a tuple with the DaysWithSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysWithSales

`func (o *DetailedItem) SetDaysWithSales(v float32)`

SetDaysWithSales sets DaysWithSales field to given value.


### GetAverageIfInStock

`func (o *DetailedItem) GetAverageIfInStock() float32`

GetAverageIfInStock returns the AverageIfInStock field if non-nil, zero value otherwise.

### GetAverageIfInStockOk

`func (o *DetailedItem) GetAverageIfInStockOk() (*float32, bool)`

GetAverageIfInStockOk returns a tuple with the AverageIfInStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageIfInStock

`func (o *DetailedItem) SetAverageIfInStock(v float32)`

SetAverageIfInStock sets AverageIfInStock field to given value.


### GetThumb

`func (o *DetailedItem) GetThumb() string`

GetThumb returns the Thumb field if non-nil, zero value otherwise.

### GetThumbOk

`func (o *DetailedItem) GetThumbOk() (*string, bool)`

GetThumbOk returns a tuple with the Thumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumb

`func (o *DetailedItem) SetThumb(v string)`

SetThumb sets Thumb field to given value.


### GetGraph

`func (o *DetailedItem) GetGraph() []float32`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *DetailedItem) GetGraphOk() (*[]float32, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *DetailedItem) SetGraph(v []float32)`

SetGraph sets Graph field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


