# SearchCategoriesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Brand** | Pointer to **string** |  | [optional] 
**Seller** | Pointer to **string** |  | [optional] 
**SupplierId** | Pointer to **int32** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CategoryPosition** | Pointer to **int32** |  | [optional] 
**PosData** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Balance** | Pointer to **int32** |  | [optional] 
**BalanceFbs** | Pointer to **int32** |  | [optional] 
**Comments** | Pointer to **int32** |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**FinalPrice** | Pointer to **int32** |  | [optional] 
**FinalPriceMax** | Pointer to **int32** |  | [optional] 
**FinalPriceMin** | Pointer to **int32** |  | [optional] 
**FinalPriceAverage** | Pointer to **int32** |  | [optional] 
**BasicSale** | Pointer to **int32** |  | [optional] 
**BasicPrice** | Pointer to **int32** |  | [optional] 
**PromoSale** | Pointer to **int32** |  | [optional] 
**ClientSale** | Pointer to **int32** |  | [optional] 
**ClientPrice** | Pointer to **int32** |  | [optional] 
**StartPrice** | Pointer to **int32** |  | [optional] 
**Sales** | Pointer to **int32** |  | [optional] 
**SalesPerDayAverage** | Pointer to **float32** |  | [optional] 
**Revenue** | Pointer to **int32** |  | [optional] 
**RevenuePotential** | Pointer to **int32** |  | [optional] 
**LostProfit** | Pointer to **int32** |  | [optional] 
**LostProfitPercent** | Pointer to **float32** |  | [optional] 
**DaysInStock** | Pointer to **int32** |  | [optional] 
**DaysWithSales** | Pointer to **int32** |  | [optional] 
**AverageIfInStock** | Pointer to **float32** |  | [optional] 
**Thumb** | Pointer to **string** |  | [optional] 
**ThumbMiddle** | Pointer to **string** |  | [optional] 
**IsFbs** | Pointer to **int32** |  | [optional] 
**SubjectId** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**SkuFirstDate** | Pointer to **string** |  | [optional] 
**Firstcommentdate** | Pointer to **string** |  | [optional] 
**Picscount** | Pointer to **int32** |  | [optional] 
**Has3d** | Pointer to **int32** |  | [optional] 
**Hasvideo** | Pointer to **int32** |  | [optional] 
**Commentsvaluation** | Pointer to **float32** |  | [optional] 
**Cardratingval** | Pointer to **int32** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**CategoriesLastCount** | Pointer to **string** |  | [optional] 
**Graph** | Pointer to **[]int32** |  | [optional] 
**CategoryGraph** | Pointer to **[]int32** |  | [optional] 
**StocksGraph** | Pointer to **[]int32** |  | [optional] 
**PriceGraph** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewSearchCategoriesItem

`func NewSearchCategoriesItem() *SearchCategoriesItem`

NewSearchCategoriesItem instantiates a new SearchCategoriesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchCategoriesItemWithDefaults

`func NewSearchCategoriesItemWithDefaults() *SearchCategoriesItem`

NewSearchCategoriesItemWithDefaults instantiates a new SearchCategoriesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchCategoriesItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchCategoriesItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchCategoriesItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SearchCategoriesItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SearchCategoriesItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchCategoriesItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchCategoriesItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchCategoriesItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBrand

`func (o *SearchCategoriesItem) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *SearchCategoriesItem) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *SearchCategoriesItem) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *SearchCategoriesItem) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetSeller

`func (o *SearchCategoriesItem) GetSeller() string`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *SearchCategoriesItem) GetSellerOk() (*string, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *SearchCategoriesItem) SetSeller(v string)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *SearchCategoriesItem) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetSupplierId

`func (o *SearchCategoriesItem) GetSupplierId() int32`

GetSupplierId returns the SupplierId field if non-nil, zero value otherwise.

### GetSupplierIdOk

`func (o *SearchCategoriesItem) GetSupplierIdOk() (*int32, bool)`

GetSupplierIdOk returns a tuple with the SupplierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierId

`func (o *SearchCategoriesItem) SetSupplierId(v int32)`

SetSupplierId sets SupplierId field to given value.

### HasSupplierId

`func (o *SearchCategoriesItem) HasSupplierId() bool`

HasSupplierId returns a boolean if a field has been set.

### GetCategory

`func (o *SearchCategoriesItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SearchCategoriesItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SearchCategoriesItem) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SearchCategoriesItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCategoryPosition

`func (o *SearchCategoriesItem) GetCategoryPosition() int32`

GetCategoryPosition returns the CategoryPosition field if non-nil, zero value otherwise.

### GetCategoryPositionOk

`func (o *SearchCategoriesItem) GetCategoryPositionOk() (*int32, bool)`

GetCategoryPositionOk returns a tuple with the CategoryPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryPosition

`func (o *SearchCategoriesItem) SetCategoryPosition(v int32)`

SetCategoryPosition sets CategoryPosition field to given value.

### HasCategoryPosition

`func (o *SearchCategoriesItem) HasCategoryPosition() bool`

HasCategoryPosition returns a boolean if a field has been set.

### GetPosData

`func (o *SearchCategoriesItem) GetPosData() string`

GetPosData returns the PosData field if non-nil, zero value otherwise.

### GetPosDataOk

`func (o *SearchCategoriesItem) GetPosDataOk() (*string, bool)`

GetPosDataOk returns a tuple with the PosData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosData

`func (o *SearchCategoriesItem) SetPosData(v string)`

SetPosData sets PosData field to given value.

### HasPosData

`func (o *SearchCategoriesItem) HasPosData() bool`

HasPosData returns a boolean if a field has been set.

### GetColor

`func (o *SearchCategoriesItem) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *SearchCategoriesItem) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *SearchCategoriesItem) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *SearchCategoriesItem) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetBalance

`func (o *SearchCategoriesItem) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *SearchCategoriesItem) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *SearchCategoriesItem) SetBalance(v int32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *SearchCategoriesItem) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetBalanceFbs

`func (o *SearchCategoriesItem) GetBalanceFbs() int32`

GetBalanceFbs returns the BalanceFbs field if non-nil, zero value otherwise.

### GetBalanceFbsOk

`func (o *SearchCategoriesItem) GetBalanceFbsOk() (*int32, bool)`

GetBalanceFbsOk returns a tuple with the BalanceFbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceFbs

`func (o *SearchCategoriesItem) SetBalanceFbs(v int32)`

SetBalanceFbs sets BalanceFbs field to given value.

### HasBalanceFbs

`func (o *SearchCategoriesItem) HasBalanceFbs() bool`

HasBalanceFbs returns a boolean if a field has been set.

### GetComments

`func (o *SearchCategoriesItem) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *SearchCategoriesItem) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *SearchCategoriesItem) SetComments(v int32)`

SetComments sets Comments field to given value.

### HasComments

`func (o *SearchCategoriesItem) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetRating

`func (o *SearchCategoriesItem) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *SearchCategoriesItem) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *SearchCategoriesItem) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *SearchCategoriesItem) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetFinalPrice

`func (o *SearchCategoriesItem) GetFinalPrice() int32`

GetFinalPrice returns the FinalPrice field if non-nil, zero value otherwise.

### GetFinalPriceOk

`func (o *SearchCategoriesItem) GetFinalPriceOk() (*int32, bool)`

GetFinalPriceOk returns a tuple with the FinalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPrice

`func (o *SearchCategoriesItem) SetFinalPrice(v int32)`

SetFinalPrice sets FinalPrice field to given value.

### HasFinalPrice

`func (o *SearchCategoriesItem) HasFinalPrice() bool`

HasFinalPrice returns a boolean if a field has been set.

### GetFinalPriceMax

`func (o *SearchCategoriesItem) GetFinalPriceMax() int32`

GetFinalPriceMax returns the FinalPriceMax field if non-nil, zero value otherwise.

### GetFinalPriceMaxOk

`func (o *SearchCategoriesItem) GetFinalPriceMaxOk() (*int32, bool)`

GetFinalPriceMaxOk returns a tuple with the FinalPriceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPriceMax

`func (o *SearchCategoriesItem) SetFinalPriceMax(v int32)`

SetFinalPriceMax sets FinalPriceMax field to given value.

### HasFinalPriceMax

`func (o *SearchCategoriesItem) HasFinalPriceMax() bool`

HasFinalPriceMax returns a boolean if a field has been set.

### GetFinalPriceMin

`func (o *SearchCategoriesItem) GetFinalPriceMin() int32`

GetFinalPriceMin returns the FinalPriceMin field if non-nil, zero value otherwise.

### GetFinalPriceMinOk

`func (o *SearchCategoriesItem) GetFinalPriceMinOk() (*int32, bool)`

GetFinalPriceMinOk returns a tuple with the FinalPriceMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPriceMin

`func (o *SearchCategoriesItem) SetFinalPriceMin(v int32)`

SetFinalPriceMin sets FinalPriceMin field to given value.

### HasFinalPriceMin

`func (o *SearchCategoriesItem) HasFinalPriceMin() bool`

HasFinalPriceMin returns a boolean if a field has been set.

### GetFinalPriceAverage

`func (o *SearchCategoriesItem) GetFinalPriceAverage() int32`

GetFinalPriceAverage returns the FinalPriceAverage field if non-nil, zero value otherwise.

### GetFinalPriceAverageOk

`func (o *SearchCategoriesItem) GetFinalPriceAverageOk() (*int32, bool)`

GetFinalPriceAverageOk returns a tuple with the FinalPriceAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPriceAverage

`func (o *SearchCategoriesItem) SetFinalPriceAverage(v int32)`

SetFinalPriceAverage sets FinalPriceAverage field to given value.

### HasFinalPriceAverage

`func (o *SearchCategoriesItem) HasFinalPriceAverage() bool`

HasFinalPriceAverage returns a boolean if a field has been set.

### GetBasicSale

`func (o *SearchCategoriesItem) GetBasicSale() int32`

GetBasicSale returns the BasicSale field if non-nil, zero value otherwise.

### GetBasicSaleOk

`func (o *SearchCategoriesItem) GetBasicSaleOk() (*int32, bool)`

GetBasicSaleOk returns a tuple with the BasicSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSale

`func (o *SearchCategoriesItem) SetBasicSale(v int32)`

SetBasicSale sets BasicSale field to given value.

### HasBasicSale

`func (o *SearchCategoriesItem) HasBasicSale() bool`

HasBasicSale returns a boolean if a field has been set.

### GetBasicPrice

`func (o *SearchCategoriesItem) GetBasicPrice() int32`

GetBasicPrice returns the BasicPrice field if non-nil, zero value otherwise.

### GetBasicPriceOk

`func (o *SearchCategoriesItem) GetBasicPriceOk() (*int32, bool)`

GetBasicPriceOk returns a tuple with the BasicPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicPrice

`func (o *SearchCategoriesItem) SetBasicPrice(v int32)`

SetBasicPrice sets BasicPrice field to given value.

### HasBasicPrice

`func (o *SearchCategoriesItem) HasBasicPrice() bool`

HasBasicPrice returns a boolean if a field has been set.

### GetPromoSale

`func (o *SearchCategoriesItem) GetPromoSale() int32`

GetPromoSale returns the PromoSale field if non-nil, zero value otherwise.

### GetPromoSaleOk

`func (o *SearchCategoriesItem) GetPromoSaleOk() (*int32, bool)`

GetPromoSaleOk returns a tuple with the PromoSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoSale

`func (o *SearchCategoriesItem) SetPromoSale(v int32)`

SetPromoSale sets PromoSale field to given value.

### HasPromoSale

`func (o *SearchCategoriesItem) HasPromoSale() bool`

HasPromoSale returns a boolean if a field has been set.

### GetClientSale

`func (o *SearchCategoriesItem) GetClientSale() int32`

GetClientSale returns the ClientSale field if non-nil, zero value otherwise.

### GetClientSaleOk

`func (o *SearchCategoriesItem) GetClientSaleOk() (*int32, bool)`

GetClientSaleOk returns a tuple with the ClientSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSale

`func (o *SearchCategoriesItem) SetClientSale(v int32)`

SetClientSale sets ClientSale field to given value.

### HasClientSale

`func (o *SearchCategoriesItem) HasClientSale() bool`

HasClientSale returns a boolean if a field has been set.

### GetClientPrice

`func (o *SearchCategoriesItem) GetClientPrice() int32`

GetClientPrice returns the ClientPrice field if non-nil, zero value otherwise.

### GetClientPriceOk

`func (o *SearchCategoriesItem) GetClientPriceOk() (*int32, bool)`

GetClientPriceOk returns a tuple with the ClientPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrice

`func (o *SearchCategoriesItem) SetClientPrice(v int32)`

SetClientPrice sets ClientPrice field to given value.

### HasClientPrice

`func (o *SearchCategoriesItem) HasClientPrice() bool`

HasClientPrice returns a boolean if a field has been set.

### GetStartPrice

`func (o *SearchCategoriesItem) GetStartPrice() int32`

GetStartPrice returns the StartPrice field if non-nil, zero value otherwise.

### GetStartPriceOk

`func (o *SearchCategoriesItem) GetStartPriceOk() (*int32, bool)`

GetStartPriceOk returns a tuple with the StartPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPrice

`func (o *SearchCategoriesItem) SetStartPrice(v int32)`

SetStartPrice sets StartPrice field to given value.

### HasStartPrice

`func (o *SearchCategoriesItem) HasStartPrice() bool`

HasStartPrice returns a boolean if a field has been set.

### GetSales

`func (o *SearchCategoriesItem) GetSales() int32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *SearchCategoriesItem) GetSalesOk() (*int32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *SearchCategoriesItem) SetSales(v int32)`

SetSales sets Sales field to given value.

### HasSales

`func (o *SearchCategoriesItem) HasSales() bool`

HasSales returns a boolean if a field has been set.

### GetSalesPerDayAverage

`func (o *SearchCategoriesItem) GetSalesPerDayAverage() float32`

GetSalesPerDayAverage returns the SalesPerDayAverage field if non-nil, zero value otherwise.

### GetSalesPerDayAverageOk

`func (o *SearchCategoriesItem) GetSalesPerDayAverageOk() (*float32, bool)`

GetSalesPerDayAverageOk returns a tuple with the SalesPerDayAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPerDayAverage

`func (o *SearchCategoriesItem) SetSalesPerDayAverage(v float32)`

SetSalesPerDayAverage sets SalesPerDayAverage field to given value.

### HasSalesPerDayAverage

`func (o *SearchCategoriesItem) HasSalesPerDayAverage() bool`

HasSalesPerDayAverage returns a boolean if a field has been set.

### GetRevenue

`func (o *SearchCategoriesItem) GetRevenue() int32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *SearchCategoriesItem) GetRevenueOk() (*int32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *SearchCategoriesItem) SetRevenue(v int32)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *SearchCategoriesItem) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetRevenuePotential

`func (o *SearchCategoriesItem) GetRevenuePotential() int32`

GetRevenuePotential returns the RevenuePotential field if non-nil, zero value otherwise.

### GetRevenuePotentialOk

`func (o *SearchCategoriesItem) GetRevenuePotentialOk() (*int32, bool)`

GetRevenuePotentialOk returns a tuple with the RevenuePotential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePotential

`func (o *SearchCategoriesItem) SetRevenuePotential(v int32)`

SetRevenuePotential sets RevenuePotential field to given value.

### HasRevenuePotential

`func (o *SearchCategoriesItem) HasRevenuePotential() bool`

HasRevenuePotential returns a boolean if a field has been set.

### GetLostProfit

`func (o *SearchCategoriesItem) GetLostProfit() int32`

GetLostProfit returns the LostProfit field if non-nil, zero value otherwise.

### GetLostProfitOk

`func (o *SearchCategoriesItem) GetLostProfitOk() (*int32, bool)`

GetLostProfitOk returns a tuple with the LostProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostProfit

`func (o *SearchCategoriesItem) SetLostProfit(v int32)`

SetLostProfit sets LostProfit field to given value.

### HasLostProfit

`func (o *SearchCategoriesItem) HasLostProfit() bool`

HasLostProfit returns a boolean if a field has been set.

### GetLostProfitPercent

`func (o *SearchCategoriesItem) GetLostProfitPercent() float32`

GetLostProfitPercent returns the LostProfitPercent field if non-nil, zero value otherwise.

### GetLostProfitPercentOk

`func (o *SearchCategoriesItem) GetLostProfitPercentOk() (*float32, bool)`

GetLostProfitPercentOk returns a tuple with the LostProfitPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostProfitPercent

`func (o *SearchCategoriesItem) SetLostProfitPercent(v float32)`

SetLostProfitPercent sets LostProfitPercent field to given value.

### HasLostProfitPercent

`func (o *SearchCategoriesItem) HasLostProfitPercent() bool`

HasLostProfitPercent returns a boolean if a field has been set.

### GetDaysInStock

`func (o *SearchCategoriesItem) GetDaysInStock() int32`

GetDaysInStock returns the DaysInStock field if non-nil, zero value otherwise.

### GetDaysInStockOk

`func (o *SearchCategoriesItem) GetDaysInStockOk() (*int32, bool)`

GetDaysInStockOk returns a tuple with the DaysInStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInStock

`func (o *SearchCategoriesItem) SetDaysInStock(v int32)`

SetDaysInStock sets DaysInStock field to given value.

### HasDaysInStock

`func (o *SearchCategoriesItem) HasDaysInStock() bool`

HasDaysInStock returns a boolean if a field has been set.

### GetDaysWithSales

`func (o *SearchCategoriesItem) GetDaysWithSales() int32`

GetDaysWithSales returns the DaysWithSales field if non-nil, zero value otherwise.

### GetDaysWithSalesOk

`func (o *SearchCategoriesItem) GetDaysWithSalesOk() (*int32, bool)`

GetDaysWithSalesOk returns a tuple with the DaysWithSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysWithSales

`func (o *SearchCategoriesItem) SetDaysWithSales(v int32)`

SetDaysWithSales sets DaysWithSales field to given value.

### HasDaysWithSales

`func (o *SearchCategoriesItem) HasDaysWithSales() bool`

HasDaysWithSales returns a boolean if a field has been set.

### GetAverageIfInStock

`func (o *SearchCategoriesItem) GetAverageIfInStock() float32`

GetAverageIfInStock returns the AverageIfInStock field if non-nil, zero value otherwise.

### GetAverageIfInStockOk

`func (o *SearchCategoriesItem) GetAverageIfInStockOk() (*float32, bool)`

GetAverageIfInStockOk returns a tuple with the AverageIfInStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageIfInStock

`func (o *SearchCategoriesItem) SetAverageIfInStock(v float32)`

SetAverageIfInStock sets AverageIfInStock field to given value.

### HasAverageIfInStock

`func (o *SearchCategoriesItem) HasAverageIfInStock() bool`

HasAverageIfInStock returns a boolean if a field has been set.

### GetThumb

`func (o *SearchCategoriesItem) GetThumb() string`

GetThumb returns the Thumb field if non-nil, zero value otherwise.

### GetThumbOk

`func (o *SearchCategoriesItem) GetThumbOk() (*string, bool)`

GetThumbOk returns a tuple with the Thumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumb

`func (o *SearchCategoriesItem) SetThumb(v string)`

SetThumb sets Thumb field to given value.

### HasThumb

`func (o *SearchCategoriesItem) HasThumb() bool`

HasThumb returns a boolean if a field has been set.

### GetThumbMiddle

`func (o *SearchCategoriesItem) GetThumbMiddle() string`

GetThumbMiddle returns the ThumbMiddle field if non-nil, zero value otherwise.

### GetThumbMiddleOk

`func (o *SearchCategoriesItem) GetThumbMiddleOk() (*string, bool)`

GetThumbMiddleOk returns a tuple with the ThumbMiddle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbMiddle

`func (o *SearchCategoriesItem) SetThumbMiddle(v string)`

SetThumbMiddle sets ThumbMiddle field to given value.

### HasThumbMiddle

`func (o *SearchCategoriesItem) HasThumbMiddle() bool`

HasThumbMiddle returns a boolean if a field has been set.

### GetIsFbs

`func (o *SearchCategoriesItem) GetIsFbs() int32`

GetIsFbs returns the IsFbs field if non-nil, zero value otherwise.

### GetIsFbsOk

`func (o *SearchCategoriesItem) GetIsFbsOk() (*int32, bool)`

GetIsFbsOk returns a tuple with the IsFbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFbs

`func (o *SearchCategoriesItem) SetIsFbs(v int32)`

SetIsFbs sets IsFbs field to given value.

### HasIsFbs

`func (o *SearchCategoriesItem) HasIsFbs() bool`

HasIsFbs returns a boolean if a field has been set.

### GetSubjectId

`func (o *SearchCategoriesItem) GetSubjectId() int32`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *SearchCategoriesItem) GetSubjectIdOk() (*int32, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *SearchCategoriesItem) SetSubjectId(v int32)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *SearchCategoriesItem) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetUrl

`func (o *SearchCategoriesItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SearchCategoriesItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SearchCategoriesItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SearchCategoriesItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCountry

`func (o *SearchCategoriesItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SearchCategoriesItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SearchCategoriesItem) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SearchCategoriesItem) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetGender

`func (o *SearchCategoriesItem) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *SearchCategoriesItem) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *SearchCategoriesItem) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *SearchCategoriesItem) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetSkuFirstDate

`func (o *SearchCategoriesItem) GetSkuFirstDate() string`

GetSkuFirstDate returns the SkuFirstDate field if non-nil, zero value otherwise.

### GetSkuFirstDateOk

`func (o *SearchCategoriesItem) GetSkuFirstDateOk() (*string, bool)`

GetSkuFirstDateOk returns a tuple with the SkuFirstDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuFirstDate

`func (o *SearchCategoriesItem) SetSkuFirstDate(v string)`

SetSkuFirstDate sets SkuFirstDate field to given value.

### HasSkuFirstDate

`func (o *SearchCategoriesItem) HasSkuFirstDate() bool`

HasSkuFirstDate returns a boolean if a field has been set.

### GetFirstcommentdate

`func (o *SearchCategoriesItem) GetFirstcommentdate() string`

GetFirstcommentdate returns the Firstcommentdate field if non-nil, zero value otherwise.

### GetFirstcommentdateOk

`func (o *SearchCategoriesItem) GetFirstcommentdateOk() (*string, bool)`

GetFirstcommentdateOk returns a tuple with the Firstcommentdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstcommentdate

`func (o *SearchCategoriesItem) SetFirstcommentdate(v string)`

SetFirstcommentdate sets Firstcommentdate field to given value.

### HasFirstcommentdate

`func (o *SearchCategoriesItem) HasFirstcommentdate() bool`

HasFirstcommentdate returns a boolean if a field has been set.

### GetPicscount

`func (o *SearchCategoriesItem) GetPicscount() int32`

GetPicscount returns the Picscount field if non-nil, zero value otherwise.

### GetPicscountOk

`func (o *SearchCategoriesItem) GetPicscountOk() (*int32, bool)`

GetPicscountOk returns a tuple with the Picscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicscount

`func (o *SearchCategoriesItem) SetPicscount(v int32)`

SetPicscount sets Picscount field to given value.

### HasPicscount

`func (o *SearchCategoriesItem) HasPicscount() bool`

HasPicscount returns a boolean if a field has been set.

### GetHas3d

`func (o *SearchCategoriesItem) GetHas3d() int32`

GetHas3d returns the Has3d field if non-nil, zero value otherwise.

### GetHas3dOk

`func (o *SearchCategoriesItem) GetHas3dOk() (*int32, bool)`

GetHas3dOk returns a tuple with the Has3d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHas3d

`func (o *SearchCategoriesItem) SetHas3d(v int32)`

SetHas3d sets Has3d field to given value.

### HasHas3d

`func (o *SearchCategoriesItem) HasHas3d() bool`

HasHas3d returns a boolean if a field has been set.

### GetHasvideo

`func (o *SearchCategoriesItem) GetHasvideo() int32`

GetHasvideo returns the Hasvideo field if non-nil, zero value otherwise.

### GetHasvideoOk

`func (o *SearchCategoriesItem) GetHasvideoOk() (*int32, bool)`

GetHasvideoOk returns a tuple with the Hasvideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasvideo

`func (o *SearchCategoriesItem) SetHasvideo(v int32)`

SetHasvideo sets Hasvideo field to given value.

### HasHasvideo

`func (o *SearchCategoriesItem) HasHasvideo() bool`

HasHasvideo returns a boolean if a field has been set.

### GetCommentsvaluation

`func (o *SearchCategoriesItem) GetCommentsvaluation() float32`

GetCommentsvaluation returns the Commentsvaluation field if non-nil, zero value otherwise.

### GetCommentsvaluationOk

`func (o *SearchCategoriesItem) GetCommentsvaluationOk() (*float32, bool)`

GetCommentsvaluationOk returns a tuple with the Commentsvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsvaluation

`func (o *SearchCategoriesItem) SetCommentsvaluation(v float32)`

SetCommentsvaluation sets Commentsvaluation field to given value.

### HasCommentsvaluation

`func (o *SearchCategoriesItem) HasCommentsvaluation() bool`

HasCommentsvaluation returns a boolean if a field has been set.

### GetCardratingval

`func (o *SearchCategoriesItem) GetCardratingval() int32`

GetCardratingval returns the Cardratingval field if non-nil, zero value otherwise.

### GetCardratingvalOk

`func (o *SearchCategoriesItem) GetCardratingvalOk() (*int32, bool)`

GetCardratingvalOk returns a tuple with the Cardratingval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardratingval

`func (o *SearchCategoriesItem) SetCardratingval(v int32)`

SetCardratingval sets Cardratingval field to given value.

### HasCardratingval

`func (o *SearchCategoriesItem) HasCardratingval() bool`

HasCardratingval returns a boolean if a field has been set.

### GetPosition

`func (o *SearchCategoriesItem) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SearchCategoriesItem) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SearchCategoriesItem) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *SearchCategoriesItem) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetCategoriesLastCount

`func (o *SearchCategoriesItem) GetCategoriesLastCount() string`

GetCategoriesLastCount returns the CategoriesLastCount field if non-nil, zero value otherwise.

### GetCategoriesLastCountOk

`func (o *SearchCategoriesItem) GetCategoriesLastCountOk() (*string, bool)`

GetCategoriesLastCountOk returns a tuple with the CategoriesLastCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesLastCount

`func (o *SearchCategoriesItem) SetCategoriesLastCount(v string)`

SetCategoriesLastCount sets CategoriesLastCount field to given value.

### HasCategoriesLastCount

`func (o *SearchCategoriesItem) HasCategoriesLastCount() bool`

HasCategoriesLastCount returns a boolean if a field has been set.

### GetGraph

`func (o *SearchCategoriesItem) GetGraph() []int32`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *SearchCategoriesItem) GetGraphOk() (*[]int32, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *SearchCategoriesItem) SetGraph(v []int32)`

SetGraph sets Graph field to given value.

### HasGraph

`func (o *SearchCategoriesItem) HasGraph() bool`

HasGraph returns a boolean if a field has been set.

### GetCategoryGraph

`func (o *SearchCategoriesItem) GetCategoryGraph() []int32`

GetCategoryGraph returns the CategoryGraph field if non-nil, zero value otherwise.

### GetCategoryGraphOk

`func (o *SearchCategoriesItem) GetCategoryGraphOk() (*[]int32, bool)`

GetCategoryGraphOk returns a tuple with the CategoryGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGraph

`func (o *SearchCategoriesItem) SetCategoryGraph(v []int32)`

SetCategoryGraph sets CategoryGraph field to given value.

### HasCategoryGraph

`func (o *SearchCategoriesItem) HasCategoryGraph() bool`

HasCategoryGraph returns a boolean if a field has been set.

### GetStocksGraph

`func (o *SearchCategoriesItem) GetStocksGraph() []int32`

GetStocksGraph returns the StocksGraph field if non-nil, zero value otherwise.

### GetStocksGraphOk

`func (o *SearchCategoriesItem) GetStocksGraphOk() (*[]int32, bool)`

GetStocksGraphOk returns a tuple with the StocksGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStocksGraph

`func (o *SearchCategoriesItem) SetStocksGraph(v []int32)`

SetStocksGraph sets StocksGraph field to given value.

### HasStocksGraph

`func (o *SearchCategoriesItem) HasStocksGraph() bool`

HasStocksGraph returns a boolean if a field has been set.

### GetPriceGraph

`func (o *SearchCategoriesItem) GetPriceGraph() []int32`

GetPriceGraph returns the PriceGraph field if non-nil, zero value otherwise.

### GetPriceGraphOk

`func (o *SearchCategoriesItem) GetPriceGraphOk() (*[]int32, bool)`

GetPriceGraphOk returns a tuple with the PriceGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGraph

`func (o *SearchCategoriesItem) SetPriceGraph(v []int32)`

SetPriceGraph sets PriceGraph field to given value.

### HasPriceGraph

`func (o *SearchCategoriesItem) HasPriceGraph() bool`

HasPriceGraph returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


