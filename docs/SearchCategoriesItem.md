# SearchCategoriesItem

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
**PosData** | **string** |  | 
**Color** | **string** |  | 
**Balance** | **int32** |  | 
**BalanceFbs** | **int32** |  | 
**Comments** | **int32** |  | 
**Rating** | **int32** |  | 
**FinalPrice** | **int32** |  | 
**FinalPriceMax** | **int32** |  | 
**FinalPriceMin** | **int32** |  | 
**FinalPriceAverage** | **int32** |  | 
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
**CategoriesLastCount** | **string** |  | 
**Graph** | **[]int32** |  | 
**CategoryGraph** | **[]int32** |  | 
**StocksGraph** | **[]int32** |  | 
**PriceGraph** | **[]int32** |  | 

## Methods

### NewSearchCategoriesItem

`func NewSearchCategoriesItem(id int32, name string, brand string, seller string, supplierId int32, category string, categoryPosition int32, posData string, color string, balance int32, balanceFbs int32, comments int32, rating int32, finalPrice int32, finalPriceMax int32, finalPriceMin int32, finalPriceAverage int32, basicSale int32, basicPrice int32, promoSale int32, clientSale int32, clientPrice int32, startPrice int32, sales int32, salesPerDayAverage float32, revenue int32, revenuePotential int32, lostProfit int32, lostProfitPercent float32, daysInStock int32, daysWithSales int32, averageIfInStock float32, thumb string, thumbMiddle string, isFbs int32, subjectId int32, url string, country string, gender string, skuFirstDate string, firstcommentdate string, picscount int32, has3d int32, hasvideo int32, commentsvaluation float32, cardratingval int32, position int32, categoriesLastCount string, graph []int32, categoryGraph []int32, stocksGraph []int32, priceGraph []int32, ) *SearchCategoriesItem`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


