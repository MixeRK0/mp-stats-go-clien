# SearchItemsElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Brand** | **string** |  | 
**Seller** | **string** |  | 
**SupplierId** | **int32** |  | 
**Category** | **string** |  | 
**CategoryPosition** | **int32** |  | 
**PosData** | **interface{}** |  | 
**Color** | **string** |  | 
**Balance** | **int32** |  | 
**BalanceFbs** | **int32** |  | 
**Comments** | **int32** |  | 
**Rating** | **int32** |  | 
**FinalPrice** | **int32** |  | 
**FinalPriceMax** | **int32** |  | 
**FinalPriceMin** | **int32** |  | 
**FinalPriceAverage** | **float32** |  | 
**BasicSale** | **int32** |  | 
**BasicPrice** | **int32** |  | 
**PromoSale** | **int32** |  | 
**ClientSale** | **int32** |  | 
**ClientPrice** | **int32** |  | 
**StartPrice** | **int32** |  | 
**Sales** | **int32** |  | 
**SalesPerDayAverage** | **float32** |  | 
**Revenue** | **int32** |  | 
**RevenuePotential** | **int32** |  | 
**LostProfit** | **int32** |  | 
**LostProfitPercent** | **float32** |  | 
**DaysInStock** | **int32** |  | 
**DaysWithSales** | **int32** |  | 
**AverageIfInStock** | **float32** |  | 
**Thumb** | **string** |  | 
**ThumbMiddle** | **string** |  | 
**IsFbs** | **int32** |  | 
**SubjectId** | **int32** |  | 
**Url** | **string** |  | 
**Country** | **string** |  | 
**Gender** | **string** |  | 
**SkuFirstDate** | **string** |  | 
**Firstcommentdate** | **string** |  | 
**Picscount** | **int32** |  | 
**Has3d** | **int32** |  | 
**Hasvideo** | **int32** |  | 
**Commentsvaluation** | **float32** |  | 
**Cardratingval** | **int32** |  | 
**Position** | **int32** |  | 
**CategoriesLastCount** | **interface{}** |  | 
**Graph** | **[]int32** |  | 
**CategoryGraph** | **[]int32** |  | 
**StocksGraph** | **[]int32** |  | 
**PriceGraph** | **[]int32** |  | 
**Purchase** | **float32** | Процен выкупа | 
**PurchaseAfterReturn** | **float32** | Процент выкупа после возвратов | 

## Methods

### NewSearchItemsElement

`func NewSearchItemsElement(id int32, name string, brand string, seller string, supplierId int32, category string, categoryPosition int32, posData interface{}, color string, balance int32, balanceFbs int32, comments int32, rating int32, finalPrice int32, finalPriceMax int32, finalPriceMin int32, finalPriceAverage float32, basicSale int32, basicPrice int32, promoSale int32, clientSale int32, clientPrice int32, startPrice int32, sales int32, salesPerDayAverage float32, revenue int32, revenuePotential int32, lostProfit int32, lostProfitPercent float32, daysInStock int32, daysWithSales int32, averageIfInStock float32, thumb string, thumbMiddle string, isFbs int32, subjectId int32, url string, country string, gender string, skuFirstDate string, firstcommentdate string, picscount int32, has3d int32, hasvideo int32, commentsvaluation float32, cardratingval int32, position int32, categoriesLastCount interface{}, graph []int32, categoryGraph []int32, stocksGraph []int32, priceGraph []int32, purchase float32, purchaseAfterReturn float32, ) *SearchItemsElement`

NewSearchItemsElement instantiates a new SearchItemsElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchItemsElementWithDefaults

`func NewSearchItemsElementWithDefaults() *SearchItemsElement`

NewSearchItemsElementWithDefaults instantiates a new SearchItemsElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchItemsElement) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchItemsElement) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchItemsElement) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *SearchItemsElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchItemsElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchItemsElement) SetName(v string)`

SetName sets Name field to given value.


### GetBrand

`func (o *SearchItemsElement) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *SearchItemsElement) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *SearchItemsElement) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetSeller

`func (o *SearchItemsElement) GetSeller() string`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *SearchItemsElement) GetSellerOk() (*string, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *SearchItemsElement) SetSeller(v string)`

SetSeller sets Seller field to given value.


### GetSupplierId

`func (o *SearchItemsElement) GetSupplierId() int32`

GetSupplierId returns the SupplierId field if non-nil, zero value otherwise.

### GetSupplierIdOk

`func (o *SearchItemsElement) GetSupplierIdOk() (*int32, bool)`

GetSupplierIdOk returns a tuple with the SupplierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierId

`func (o *SearchItemsElement) SetSupplierId(v int32)`

SetSupplierId sets SupplierId field to given value.


### GetCategory

`func (o *SearchItemsElement) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SearchItemsElement) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SearchItemsElement) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetCategoryPosition

`func (o *SearchItemsElement) GetCategoryPosition() int32`

GetCategoryPosition returns the CategoryPosition field if non-nil, zero value otherwise.

### GetCategoryPositionOk

`func (o *SearchItemsElement) GetCategoryPositionOk() (*int32, bool)`

GetCategoryPositionOk returns a tuple with the CategoryPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryPosition

`func (o *SearchItemsElement) SetCategoryPosition(v int32)`

SetCategoryPosition sets CategoryPosition field to given value.


### GetPosData

`func (o *SearchItemsElement) GetPosData() interface{}`

GetPosData returns the PosData field if non-nil, zero value otherwise.

### GetPosDataOk

`func (o *SearchItemsElement) GetPosDataOk() (*interface{}, bool)`

GetPosDataOk returns a tuple with the PosData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosData

`func (o *SearchItemsElement) SetPosData(v interface{})`

SetPosData sets PosData field to given value.


### SetPosDataNil

`func (o *SearchItemsElement) SetPosDataNil(b bool)`

 SetPosDataNil sets the value for PosData to be an explicit nil

### UnsetPosData
`func (o *SearchItemsElement) UnsetPosData()`

UnsetPosData ensures that no value is present for PosData, not even an explicit nil
### GetColor

`func (o *SearchItemsElement) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *SearchItemsElement) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *SearchItemsElement) SetColor(v string)`

SetColor sets Color field to given value.


### GetBalance

`func (o *SearchItemsElement) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *SearchItemsElement) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *SearchItemsElement) SetBalance(v int32)`

SetBalance sets Balance field to given value.


### GetBalanceFbs

`func (o *SearchItemsElement) GetBalanceFbs() int32`

GetBalanceFbs returns the BalanceFbs field if non-nil, zero value otherwise.

### GetBalanceFbsOk

`func (o *SearchItemsElement) GetBalanceFbsOk() (*int32, bool)`

GetBalanceFbsOk returns a tuple with the BalanceFbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceFbs

`func (o *SearchItemsElement) SetBalanceFbs(v int32)`

SetBalanceFbs sets BalanceFbs field to given value.


### GetComments

`func (o *SearchItemsElement) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *SearchItemsElement) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *SearchItemsElement) SetComments(v int32)`

SetComments sets Comments field to given value.


### GetRating

`func (o *SearchItemsElement) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *SearchItemsElement) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *SearchItemsElement) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetFinalPrice

`func (o *SearchItemsElement) GetFinalPrice() int32`

GetFinalPrice returns the FinalPrice field if non-nil, zero value otherwise.

### GetFinalPriceOk

`func (o *SearchItemsElement) GetFinalPriceOk() (*int32, bool)`

GetFinalPriceOk returns a tuple with the FinalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPrice

`func (o *SearchItemsElement) SetFinalPrice(v int32)`

SetFinalPrice sets FinalPrice field to given value.


### GetFinalPriceMax

`func (o *SearchItemsElement) GetFinalPriceMax() int32`

GetFinalPriceMax returns the FinalPriceMax field if non-nil, zero value otherwise.

### GetFinalPriceMaxOk

`func (o *SearchItemsElement) GetFinalPriceMaxOk() (*int32, bool)`

GetFinalPriceMaxOk returns a tuple with the FinalPriceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPriceMax

`func (o *SearchItemsElement) SetFinalPriceMax(v int32)`

SetFinalPriceMax sets FinalPriceMax field to given value.


### GetFinalPriceMin

`func (o *SearchItemsElement) GetFinalPriceMin() int32`

GetFinalPriceMin returns the FinalPriceMin field if non-nil, zero value otherwise.

### GetFinalPriceMinOk

`func (o *SearchItemsElement) GetFinalPriceMinOk() (*int32, bool)`

GetFinalPriceMinOk returns a tuple with the FinalPriceMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPriceMin

`func (o *SearchItemsElement) SetFinalPriceMin(v int32)`

SetFinalPriceMin sets FinalPriceMin field to given value.


### GetFinalPriceAverage

`func (o *SearchItemsElement) GetFinalPriceAverage() float32`

GetFinalPriceAverage returns the FinalPriceAverage field if non-nil, zero value otherwise.

### GetFinalPriceAverageOk

`func (o *SearchItemsElement) GetFinalPriceAverageOk() (*float32, bool)`

GetFinalPriceAverageOk returns a tuple with the FinalPriceAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPriceAverage

`func (o *SearchItemsElement) SetFinalPriceAverage(v float32)`

SetFinalPriceAverage sets FinalPriceAverage field to given value.


### GetBasicSale

`func (o *SearchItemsElement) GetBasicSale() int32`

GetBasicSale returns the BasicSale field if non-nil, zero value otherwise.

### GetBasicSaleOk

`func (o *SearchItemsElement) GetBasicSaleOk() (*int32, bool)`

GetBasicSaleOk returns a tuple with the BasicSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSale

`func (o *SearchItemsElement) SetBasicSale(v int32)`

SetBasicSale sets BasicSale field to given value.


### GetBasicPrice

`func (o *SearchItemsElement) GetBasicPrice() int32`

GetBasicPrice returns the BasicPrice field if non-nil, zero value otherwise.

### GetBasicPriceOk

`func (o *SearchItemsElement) GetBasicPriceOk() (*int32, bool)`

GetBasicPriceOk returns a tuple with the BasicPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicPrice

`func (o *SearchItemsElement) SetBasicPrice(v int32)`

SetBasicPrice sets BasicPrice field to given value.


### GetPromoSale

`func (o *SearchItemsElement) GetPromoSale() int32`

GetPromoSale returns the PromoSale field if non-nil, zero value otherwise.

### GetPromoSaleOk

`func (o *SearchItemsElement) GetPromoSaleOk() (*int32, bool)`

GetPromoSaleOk returns a tuple with the PromoSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoSale

`func (o *SearchItemsElement) SetPromoSale(v int32)`

SetPromoSale sets PromoSale field to given value.


### GetClientSale

`func (o *SearchItemsElement) GetClientSale() int32`

GetClientSale returns the ClientSale field if non-nil, zero value otherwise.

### GetClientSaleOk

`func (o *SearchItemsElement) GetClientSaleOk() (*int32, bool)`

GetClientSaleOk returns a tuple with the ClientSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSale

`func (o *SearchItemsElement) SetClientSale(v int32)`

SetClientSale sets ClientSale field to given value.


### GetClientPrice

`func (o *SearchItemsElement) GetClientPrice() int32`

GetClientPrice returns the ClientPrice field if non-nil, zero value otherwise.

### GetClientPriceOk

`func (o *SearchItemsElement) GetClientPriceOk() (*int32, bool)`

GetClientPriceOk returns a tuple with the ClientPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrice

`func (o *SearchItemsElement) SetClientPrice(v int32)`

SetClientPrice sets ClientPrice field to given value.


### GetStartPrice

`func (o *SearchItemsElement) GetStartPrice() int32`

GetStartPrice returns the StartPrice field if non-nil, zero value otherwise.

### GetStartPriceOk

`func (o *SearchItemsElement) GetStartPriceOk() (*int32, bool)`

GetStartPriceOk returns a tuple with the StartPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPrice

`func (o *SearchItemsElement) SetStartPrice(v int32)`

SetStartPrice sets StartPrice field to given value.


### GetSales

`func (o *SearchItemsElement) GetSales() int32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *SearchItemsElement) GetSalesOk() (*int32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *SearchItemsElement) SetSales(v int32)`

SetSales sets Sales field to given value.


### GetSalesPerDayAverage

`func (o *SearchItemsElement) GetSalesPerDayAverage() float32`

GetSalesPerDayAverage returns the SalesPerDayAverage field if non-nil, zero value otherwise.

### GetSalesPerDayAverageOk

`func (o *SearchItemsElement) GetSalesPerDayAverageOk() (*float32, bool)`

GetSalesPerDayAverageOk returns a tuple with the SalesPerDayAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPerDayAverage

`func (o *SearchItemsElement) SetSalesPerDayAverage(v float32)`

SetSalesPerDayAverage sets SalesPerDayAverage field to given value.


### GetRevenue

`func (o *SearchItemsElement) GetRevenue() int32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *SearchItemsElement) GetRevenueOk() (*int32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *SearchItemsElement) SetRevenue(v int32)`

SetRevenue sets Revenue field to given value.


### GetRevenuePotential

`func (o *SearchItemsElement) GetRevenuePotential() int32`

GetRevenuePotential returns the RevenuePotential field if non-nil, zero value otherwise.

### GetRevenuePotentialOk

`func (o *SearchItemsElement) GetRevenuePotentialOk() (*int32, bool)`

GetRevenuePotentialOk returns a tuple with the RevenuePotential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePotential

`func (o *SearchItemsElement) SetRevenuePotential(v int32)`

SetRevenuePotential sets RevenuePotential field to given value.


### GetLostProfit

`func (o *SearchItemsElement) GetLostProfit() int32`

GetLostProfit returns the LostProfit field if non-nil, zero value otherwise.

### GetLostProfitOk

`func (o *SearchItemsElement) GetLostProfitOk() (*int32, bool)`

GetLostProfitOk returns a tuple with the LostProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostProfit

`func (o *SearchItemsElement) SetLostProfit(v int32)`

SetLostProfit sets LostProfit field to given value.


### GetLostProfitPercent

`func (o *SearchItemsElement) GetLostProfitPercent() float32`

GetLostProfitPercent returns the LostProfitPercent field if non-nil, zero value otherwise.

### GetLostProfitPercentOk

`func (o *SearchItemsElement) GetLostProfitPercentOk() (*float32, bool)`

GetLostProfitPercentOk returns a tuple with the LostProfitPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostProfitPercent

`func (o *SearchItemsElement) SetLostProfitPercent(v float32)`

SetLostProfitPercent sets LostProfitPercent field to given value.


### GetDaysInStock

`func (o *SearchItemsElement) GetDaysInStock() int32`

GetDaysInStock returns the DaysInStock field if non-nil, zero value otherwise.

### GetDaysInStockOk

`func (o *SearchItemsElement) GetDaysInStockOk() (*int32, bool)`

GetDaysInStockOk returns a tuple with the DaysInStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInStock

`func (o *SearchItemsElement) SetDaysInStock(v int32)`

SetDaysInStock sets DaysInStock field to given value.


### GetDaysWithSales

`func (o *SearchItemsElement) GetDaysWithSales() int32`

GetDaysWithSales returns the DaysWithSales field if non-nil, zero value otherwise.

### GetDaysWithSalesOk

`func (o *SearchItemsElement) GetDaysWithSalesOk() (*int32, bool)`

GetDaysWithSalesOk returns a tuple with the DaysWithSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysWithSales

`func (o *SearchItemsElement) SetDaysWithSales(v int32)`

SetDaysWithSales sets DaysWithSales field to given value.


### GetAverageIfInStock

`func (o *SearchItemsElement) GetAverageIfInStock() float32`

GetAverageIfInStock returns the AverageIfInStock field if non-nil, zero value otherwise.

### GetAverageIfInStockOk

`func (o *SearchItemsElement) GetAverageIfInStockOk() (*float32, bool)`

GetAverageIfInStockOk returns a tuple with the AverageIfInStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageIfInStock

`func (o *SearchItemsElement) SetAverageIfInStock(v float32)`

SetAverageIfInStock sets AverageIfInStock field to given value.


### GetThumb

`func (o *SearchItemsElement) GetThumb() string`

GetThumb returns the Thumb field if non-nil, zero value otherwise.

### GetThumbOk

`func (o *SearchItemsElement) GetThumbOk() (*string, bool)`

GetThumbOk returns a tuple with the Thumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumb

`func (o *SearchItemsElement) SetThumb(v string)`

SetThumb sets Thumb field to given value.


### GetThumbMiddle

`func (o *SearchItemsElement) GetThumbMiddle() string`

GetThumbMiddle returns the ThumbMiddle field if non-nil, zero value otherwise.

### GetThumbMiddleOk

`func (o *SearchItemsElement) GetThumbMiddleOk() (*string, bool)`

GetThumbMiddleOk returns a tuple with the ThumbMiddle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbMiddle

`func (o *SearchItemsElement) SetThumbMiddle(v string)`

SetThumbMiddle sets ThumbMiddle field to given value.


### GetIsFbs

`func (o *SearchItemsElement) GetIsFbs() int32`

GetIsFbs returns the IsFbs field if non-nil, zero value otherwise.

### GetIsFbsOk

`func (o *SearchItemsElement) GetIsFbsOk() (*int32, bool)`

GetIsFbsOk returns a tuple with the IsFbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFbs

`func (o *SearchItemsElement) SetIsFbs(v int32)`

SetIsFbs sets IsFbs field to given value.


### GetSubjectId

`func (o *SearchItemsElement) GetSubjectId() int32`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *SearchItemsElement) GetSubjectIdOk() (*int32, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *SearchItemsElement) SetSubjectId(v int32)`

SetSubjectId sets SubjectId field to given value.


### GetUrl

`func (o *SearchItemsElement) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SearchItemsElement) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SearchItemsElement) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCountry

`func (o *SearchItemsElement) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SearchItemsElement) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SearchItemsElement) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetGender

`func (o *SearchItemsElement) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *SearchItemsElement) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *SearchItemsElement) SetGender(v string)`

SetGender sets Gender field to given value.


### GetSkuFirstDate

`func (o *SearchItemsElement) GetSkuFirstDate() string`

GetSkuFirstDate returns the SkuFirstDate field if non-nil, zero value otherwise.

### GetSkuFirstDateOk

`func (o *SearchItemsElement) GetSkuFirstDateOk() (*string, bool)`

GetSkuFirstDateOk returns a tuple with the SkuFirstDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuFirstDate

`func (o *SearchItemsElement) SetSkuFirstDate(v string)`

SetSkuFirstDate sets SkuFirstDate field to given value.


### GetFirstcommentdate

`func (o *SearchItemsElement) GetFirstcommentdate() string`

GetFirstcommentdate returns the Firstcommentdate field if non-nil, zero value otherwise.

### GetFirstcommentdateOk

`func (o *SearchItemsElement) GetFirstcommentdateOk() (*string, bool)`

GetFirstcommentdateOk returns a tuple with the Firstcommentdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstcommentdate

`func (o *SearchItemsElement) SetFirstcommentdate(v string)`

SetFirstcommentdate sets Firstcommentdate field to given value.


### GetPicscount

`func (o *SearchItemsElement) GetPicscount() int32`

GetPicscount returns the Picscount field if non-nil, zero value otherwise.

### GetPicscountOk

`func (o *SearchItemsElement) GetPicscountOk() (*int32, bool)`

GetPicscountOk returns a tuple with the Picscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicscount

`func (o *SearchItemsElement) SetPicscount(v int32)`

SetPicscount sets Picscount field to given value.


### GetHas3d

`func (o *SearchItemsElement) GetHas3d() int32`

GetHas3d returns the Has3d field if non-nil, zero value otherwise.

### GetHas3dOk

`func (o *SearchItemsElement) GetHas3dOk() (*int32, bool)`

GetHas3dOk returns a tuple with the Has3d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHas3d

`func (o *SearchItemsElement) SetHas3d(v int32)`

SetHas3d sets Has3d field to given value.


### GetHasvideo

`func (o *SearchItemsElement) GetHasvideo() int32`

GetHasvideo returns the Hasvideo field if non-nil, zero value otherwise.

### GetHasvideoOk

`func (o *SearchItemsElement) GetHasvideoOk() (*int32, bool)`

GetHasvideoOk returns a tuple with the Hasvideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasvideo

`func (o *SearchItemsElement) SetHasvideo(v int32)`

SetHasvideo sets Hasvideo field to given value.


### GetCommentsvaluation

`func (o *SearchItemsElement) GetCommentsvaluation() float32`

GetCommentsvaluation returns the Commentsvaluation field if non-nil, zero value otherwise.

### GetCommentsvaluationOk

`func (o *SearchItemsElement) GetCommentsvaluationOk() (*float32, bool)`

GetCommentsvaluationOk returns a tuple with the Commentsvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsvaluation

`func (o *SearchItemsElement) SetCommentsvaluation(v float32)`

SetCommentsvaluation sets Commentsvaluation field to given value.


### GetCardratingval

`func (o *SearchItemsElement) GetCardratingval() int32`

GetCardratingval returns the Cardratingval field if non-nil, zero value otherwise.

### GetCardratingvalOk

`func (o *SearchItemsElement) GetCardratingvalOk() (*int32, bool)`

GetCardratingvalOk returns a tuple with the Cardratingval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardratingval

`func (o *SearchItemsElement) SetCardratingval(v int32)`

SetCardratingval sets Cardratingval field to given value.


### GetPosition

`func (o *SearchItemsElement) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SearchItemsElement) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SearchItemsElement) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetCategoriesLastCount

`func (o *SearchItemsElement) GetCategoriesLastCount() interface{}`

GetCategoriesLastCount returns the CategoriesLastCount field if non-nil, zero value otherwise.

### GetCategoriesLastCountOk

`func (o *SearchItemsElement) GetCategoriesLastCountOk() (*interface{}, bool)`

GetCategoriesLastCountOk returns a tuple with the CategoriesLastCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesLastCount

`func (o *SearchItemsElement) SetCategoriesLastCount(v interface{})`

SetCategoriesLastCount sets CategoriesLastCount field to given value.


### SetCategoriesLastCountNil

`func (o *SearchItemsElement) SetCategoriesLastCountNil(b bool)`

 SetCategoriesLastCountNil sets the value for CategoriesLastCount to be an explicit nil

### UnsetCategoriesLastCount
`func (o *SearchItemsElement) UnsetCategoriesLastCount()`

UnsetCategoriesLastCount ensures that no value is present for CategoriesLastCount, not even an explicit nil
### GetGraph

`func (o *SearchItemsElement) GetGraph() []int32`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *SearchItemsElement) GetGraphOk() (*[]int32, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *SearchItemsElement) SetGraph(v []int32)`

SetGraph sets Graph field to given value.


### GetCategoryGraph

`func (o *SearchItemsElement) GetCategoryGraph() []int32`

GetCategoryGraph returns the CategoryGraph field if non-nil, zero value otherwise.

### GetCategoryGraphOk

`func (o *SearchItemsElement) GetCategoryGraphOk() (*[]int32, bool)`

GetCategoryGraphOk returns a tuple with the CategoryGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGraph

`func (o *SearchItemsElement) SetCategoryGraph(v []int32)`

SetCategoryGraph sets CategoryGraph field to given value.


### GetStocksGraph

`func (o *SearchItemsElement) GetStocksGraph() []int32`

GetStocksGraph returns the StocksGraph field if non-nil, zero value otherwise.

### GetStocksGraphOk

`func (o *SearchItemsElement) GetStocksGraphOk() (*[]int32, bool)`

GetStocksGraphOk returns a tuple with the StocksGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStocksGraph

`func (o *SearchItemsElement) SetStocksGraph(v []int32)`

SetStocksGraph sets StocksGraph field to given value.


### GetPriceGraph

`func (o *SearchItemsElement) GetPriceGraph() []int32`

GetPriceGraph returns the PriceGraph field if non-nil, zero value otherwise.

### GetPriceGraphOk

`func (o *SearchItemsElement) GetPriceGraphOk() (*[]int32, bool)`

GetPriceGraphOk returns a tuple with the PriceGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGraph

`func (o *SearchItemsElement) SetPriceGraph(v []int32)`

SetPriceGraph sets PriceGraph field to given value.


### GetPurchase

`func (o *SearchItemsElement) GetPurchase() float32`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *SearchItemsElement) GetPurchaseOk() (*float32, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *SearchItemsElement) SetPurchase(v float32)`

SetPurchase sets Purchase field to given value.


### GetPurchaseAfterReturn

`func (o *SearchItemsElement) GetPurchaseAfterReturn() float32`

GetPurchaseAfterReturn returns the PurchaseAfterReturn field if non-nil, zero value otherwise.

### GetPurchaseAfterReturnOk

`func (o *SearchItemsElement) GetPurchaseAfterReturnOk() (*float32, bool)`

GetPurchaseAfterReturnOk returns a tuple with the PurchaseAfterReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseAfterReturn

`func (o *SearchItemsElement) SetPurchaseAfterReturn(v float32)`

SetPurchaseAfterReturn sets PurchaseAfterReturn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


