/*
mpstats

MPStats API

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CategoryProduct struct for CategoryProduct
type CategoryProduct struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Brand string `json:"brand"`
	Seller string `json:"seller"`
	SupplierId int32 `json:"supplier_id"`
	Category string `json:"category"`
	CategoryPosition int32 `json:"category_position"`
	PosData interface{} `json:"pos_data"`
	Color string `json:"color"`
	Balance int32 `json:"balance"`
	BalanceFbs int32 `json:"balance_fbs"`
	Comments int32 `json:"comments"`
	Rating int32 `json:"rating"`
	FinalPrice int32 `json:"final_price"`
	FinalPriceMax int32 `json:"final_price_max"`
	FinalPriceMin int32 `json:"final_price_min"`
	FinalPriceAverage float32 `json:"final_price_average"`
	BasicSale int32 `json:"basic_sale"`
	BasicPrice int32 `json:"basic_price"`
	PromoSale int32 `json:"promo_sale"`
	ClientSale int32 `json:"client_sale"`
	ClientPrice int32 `json:"client_price"`
	StartPrice int32 `json:"start_price"`
	Sales int32 `json:"sales"`
	SalesPerDayAverage float32 `json:"sales_per_day_average"`
	Revenue int32 `json:"revenue"`
	RevenuePotential int32 `json:"revenue_potential"`
	LostProfit int32 `json:"lost_profit"`
	LostProfitPercent float32 `json:"lost_profit_percent"`
	DaysInStock int32 `json:"days_in_stock"`
	DaysWithSales int32 `json:"days_with_sales"`
	AverageIfInStock float32 `json:"average_if_in_stock"`
	Thumb string `json:"thumb"`
	ThumbMiddle string `json:"thumb_middle"`
	IsFbs int32 `json:"is_fbs"`
	SubjectId int32 `json:"subject_id"`
	Url string `json:"url"`
	Country string `json:"country"`
	Gender string `json:"gender"`
	SkuFirstDate string `json:"sku_first_date"`
	Firstcommentdate string `json:"firstcommentdate"`
	Picscount int32 `json:"picscount"`
	Has3d int32 `json:"has3d"`
	Hasvideo int32 `json:"hasvideo"`
	Commentsvaluation float32 `json:"commentsvaluation"`
	Cardratingval int32 `json:"cardratingval"`
	Position int32 `json:"position"`
	CategoriesLastCount int32 `json:"categories_last_count"`
	Graph []int32 `json:"graph"`
	CategoryGraph []int32 `json:"category_graph"`
	StocksGraph []int32 `json:"stocks_graph"`
	PriceGraph []int32 `json:"price_graph"`
	// Процен выкупа
	Purchase float32 `json:"purchase"`
	// Процент выкупа после возвратов
	PurchaseAfterReturn float32 `json:"purchase_after_return"`
}

// NewCategoryProduct instantiates a new CategoryProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryProduct(id int32, name string, brand string, seller string, supplierId int32, category string, categoryPosition int32, posData interface{}, color string, balance int32, balanceFbs int32, comments int32, rating int32, finalPrice int32, finalPriceMax int32, finalPriceMin int32, finalPriceAverage float32, basicSale int32, basicPrice int32, promoSale int32, clientSale int32, clientPrice int32, startPrice int32, sales int32, salesPerDayAverage float32, revenue int32, revenuePotential int32, lostProfit int32, lostProfitPercent float32, daysInStock int32, daysWithSales int32, averageIfInStock float32, thumb string, thumbMiddle string, isFbs int32, subjectId int32, url string, country string, gender string, skuFirstDate string, firstcommentdate string, picscount int32, has3d int32, hasvideo int32, commentsvaluation float32, cardratingval int32, position int32, categoriesLastCount int32, graph []int32, categoryGraph []int32, stocksGraph []int32, priceGraph []int32, purchase float32, purchaseAfterReturn float32) *CategoryProduct {
	this := CategoryProduct{}
	this.Id = id
	this.Name = name
	this.Brand = brand
	this.Seller = seller
	this.SupplierId = supplierId
	this.Category = category
	this.CategoryPosition = categoryPosition
	this.PosData = posData
	this.Color = color
	this.Balance = balance
	this.BalanceFbs = balanceFbs
	this.Comments = comments
	this.Rating = rating
	this.FinalPrice = finalPrice
	this.FinalPriceMax = finalPriceMax
	this.FinalPriceMin = finalPriceMin
	this.FinalPriceAverage = finalPriceAverage
	this.BasicSale = basicSale
	this.BasicPrice = basicPrice
	this.PromoSale = promoSale
	this.ClientSale = clientSale
	this.ClientPrice = clientPrice
	this.StartPrice = startPrice
	this.Sales = sales
	this.SalesPerDayAverage = salesPerDayAverage
	this.Revenue = revenue
	this.RevenuePotential = revenuePotential
	this.LostProfit = lostProfit
	this.LostProfitPercent = lostProfitPercent
	this.DaysInStock = daysInStock
	this.DaysWithSales = daysWithSales
	this.AverageIfInStock = averageIfInStock
	this.Thumb = thumb
	this.ThumbMiddle = thumbMiddle
	this.IsFbs = isFbs
	this.SubjectId = subjectId
	this.Url = url
	this.Country = country
	this.Gender = gender
	this.SkuFirstDate = skuFirstDate
	this.Firstcommentdate = firstcommentdate
	this.Picscount = picscount
	this.Has3d = has3d
	this.Hasvideo = hasvideo
	this.Commentsvaluation = commentsvaluation
	this.Cardratingval = cardratingval
	this.Position = position
	this.CategoriesLastCount = categoriesLastCount
	this.Graph = graph
	this.CategoryGraph = categoryGraph
	this.StocksGraph = stocksGraph
	this.PriceGraph = priceGraph
	this.Purchase = purchase
	this.PurchaseAfterReturn = purchaseAfterReturn
	return &this
}

// NewCategoryProductWithDefaults instantiates a new CategoryProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryProductWithDefaults() *CategoryProduct {
	this := CategoryProduct{}
	return &this
}

// GetId returns the Id field value
func (o *CategoryProduct) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CategoryProduct) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CategoryProduct) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CategoryProduct) SetName(v string) {
	o.Name = v
}

// GetBrand returns the Brand field value
func (o *CategoryProduct) GetBrand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Brand
}

// GetBrandOk returns a tuple with the Brand field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Brand, true
}

// SetBrand sets field value
func (o *CategoryProduct) SetBrand(v string) {
	o.Brand = v
}

// GetSeller returns the Seller field value
func (o *CategoryProduct) GetSeller() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Seller
}

// GetSellerOk returns a tuple with the Seller field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetSellerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seller, true
}

// SetSeller sets field value
func (o *CategoryProduct) SetSeller(v string) {
	o.Seller = v
}

// GetSupplierId returns the SupplierId field value
func (o *CategoryProduct) GetSupplierId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SupplierId
}

// GetSupplierIdOk returns a tuple with the SupplierId field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetSupplierIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupplierId, true
}

// SetSupplierId sets field value
func (o *CategoryProduct) SetSupplierId(v int32) {
	o.SupplierId = v
}

// GetCategory returns the Category field value
func (o *CategoryProduct) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *CategoryProduct) SetCategory(v string) {
	o.Category = v
}

// GetCategoryPosition returns the CategoryPosition field value
func (o *CategoryProduct) GetCategoryPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CategoryPosition
}

// GetCategoryPositionOk returns a tuple with the CategoryPosition field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetCategoryPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryPosition, true
}

// SetCategoryPosition sets field value
func (o *CategoryProduct) SetCategoryPosition(v int32) {
	o.CategoryPosition = v
}

// GetPosData returns the PosData field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CategoryProduct) GetPosData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PosData
}

// GetPosDataOk returns a tuple with the PosData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CategoryProduct) GetPosDataOk() (*interface{}, bool) {
	if o == nil || o.PosData == nil {
		return nil, false
	}
	return &o.PosData, true
}

// SetPosData sets field value
func (o *CategoryProduct) SetPosData(v interface{}) {
	o.PosData = v
}

// GetColor returns the Color field value
func (o *CategoryProduct) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *CategoryProduct) SetColor(v string) {
	o.Color = v
}

// GetBalance returns the Balance field value
func (o *CategoryProduct) GetBalance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetBalanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *CategoryProduct) SetBalance(v int32) {
	o.Balance = v
}

// GetBalanceFbs returns the BalanceFbs field value
func (o *CategoryProduct) GetBalanceFbs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BalanceFbs
}

// GetBalanceFbsOk returns a tuple with the BalanceFbs field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetBalanceFbsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceFbs, true
}

// SetBalanceFbs sets field value
func (o *CategoryProduct) SetBalanceFbs(v int32) {
	o.BalanceFbs = v
}

// GetComments returns the Comments field value
func (o *CategoryProduct) GetComments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetCommentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comments, true
}

// SetComments sets field value
func (o *CategoryProduct) SetComments(v int32) {
	o.Comments = v
}

// GetRating returns the Rating field value
func (o *CategoryProduct) GetRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *CategoryProduct) SetRating(v int32) {
	o.Rating = v
}

// GetFinalPrice returns the FinalPrice field value
func (o *CategoryProduct) GetFinalPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FinalPrice
}

// GetFinalPriceOk returns a tuple with the FinalPrice field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetFinalPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalPrice, true
}

// SetFinalPrice sets field value
func (o *CategoryProduct) SetFinalPrice(v int32) {
	o.FinalPrice = v
}

// GetFinalPriceMax returns the FinalPriceMax field value
func (o *CategoryProduct) GetFinalPriceMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FinalPriceMax
}

// GetFinalPriceMaxOk returns a tuple with the FinalPriceMax field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetFinalPriceMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalPriceMax, true
}

// SetFinalPriceMax sets field value
func (o *CategoryProduct) SetFinalPriceMax(v int32) {
	o.FinalPriceMax = v
}

// GetFinalPriceMin returns the FinalPriceMin field value
func (o *CategoryProduct) GetFinalPriceMin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FinalPriceMin
}

// GetFinalPriceMinOk returns a tuple with the FinalPriceMin field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetFinalPriceMinOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalPriceMin, true
}

// SetFinalPriceMin sets field value
func (o *CategoryProduct) SetFinalPriceMin(v int32) {
	o.FinalPriceMin = v
}

// GetFinalPriceAverage returns the FinalPriceAverage field value
func (o *CategoryProduct) GetFinalPriceAverage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.FinalPriceAverage
}

// GetFinalPriceAverageOk returns a tuple with the FinalPriceAverage field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetFinalPriceAverageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalPriceAverage, true
}

// SetFinalPriceAverage sets field value
func (o *CategoryProduct) SetFinalPriceAverage(v float32) {
	o.FinalPriceAverage = v
}

// GetBasicSale returns the BasicSale field value
func (o *CategoryProduct) GetBasicSale() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BasicSale
}

// GetBasicSaleOk returns a tuple with the BasicSale field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetBasicSaleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasicSale, true
}

// SetBasicSale sets field value
func (o *CategoryProduct) SetBasicSale(v int32) {
	o.BasicSale = v
}

// GetBasicPrice returns the BasicPrice field value
func (o *CategoryProduct) GetBasicPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BasicPrice
}

// GetBasicPriceOk returns a tuple with the BasicPrice field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetBasicPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasicPrice, true
}

// SetBasicPrice sets field value
func (o *CategoryProduct) SetBasicPrice(v int32) {
	o.BasicPrice = v
}

// GetPromoSale returns the PromoSale field value
func (o *CategoryProduct) GetPromoSale() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PromoSale
}

// GetPromoSaleOk returns a tuple with the PromoSale field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetPromoSaleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromoSale, true
}

// SetPromoSale sets field value
func (o *CategoryProduct) SetPromoSale(v int32) {
	o.PromoSale = v
}

// GetClientSale returns the ClientSale field value
func (o *CategoryProduct) GetClientSale() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClientSale
}

// GetClientSaleOk returns a tuple with the ClientSale field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetClientSaleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSale, true
}

// SetClientSale sets field value
func (o *CategoryProduct) SetClientSale(v int32) {
	o.ClientSale = v
}

// GetClientPrice returns the ClientPrice field value
func (o *CategoryProduct) GetClientPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClientPrice
}

// GetClientPriceOk returns a tuple with the ClientPrice field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetClientPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientPrice, true
}

// SetClientPrice sets field value
func (o *CategoryProduct) SetClientPrice(v int32) {
	o.ClientPrice = v
}

// GetStartPrice returns the StartPrice field value
func (o *CategoryProduct) GetStartPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartPrice
}

// GetStartPriceOk returns a tuple with the StartPrice field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetStartPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartPrice, true
}

// SetStartPrice sets field value
func (o *CategoryProduct) SetStartPrice(v int32) {
	o.StartPrice = v
}

// GetSales returns the Sales field value
func (o *CategoryProduct) GetSales() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sales
}

// GetSalesOk returns a tuple with the Sales field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetSalesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sales, true
}

// SetSales sets field value
func (o *CategoryProduct) SetSales(v int32) {
	o.Sales = v
}

// GetSalesPerDayAverage returns the SalesPerDayAverage field value
func (o *CategoryProduct) GetSalesPerDayAverage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SalesPerDayAverage
}

// GetSalesPerDayAverageOk returns a tuple with the SalesPerDayAverage field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetSalesPerDayAverageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SalesPerDayAverage, true
}

// SetSalesPerDayAverage sets field value
func (o *CategoryProduct) SetSalesPerDayAverage(v float32) {
	o.SalesPerDayAverage = v
}

// GetRevenue returns the Revenue field value
func (o *CategoryProduct) GetRevenue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Revenue
}

// GetRevenueOk returns a tuple with the Revenue field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetRevenueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revenue, true
}

// SetRevenue sets field value
func (o *CategoryProduct) SetRevenue(v int32) {
	o.Revenue = v
}

// GetRevenuePotential returns the RevenuePotential field value
func (o *CategoryProduct) GetRevenuePotential() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RevenuePotential
}

// GetRevenuePotentialOk returns a tuple with the RevenuePotential field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetRevenuePotentialOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevenuePotential, true
}

// SetRevenuePotential sets field value
func (o *CategoryProduct) SetRevenuePotential(v int32) {
	o.RevenuePotential = v
}

// GetLostProfit returns the LostProfit field value
func (o *CategoryProduct) GetLostProfit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LostProfit
}

// GetLostProfitOk returns a tuple with the LostProfit field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetLostProfitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LostProfit, true
}

// SetLostProfit sets field value
func (o *CategoryProduct) SetLostProfit(v int32) {
	o.LostProfit = v
}

// GetLostProfitPercent returns the LostProfitPercent field value
func (o *CategoryProduct) GetLostProfitPercent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LostProfitPercent
}

// GetLostProfitPercentOk returns a tuple with the LostProfitPercent field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetLostProfitPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LostProfitPercent, true
}

// SetLostProfitPercent sets field value
func (o *CategoryProduct) SetLostProfitPercent(v float32) {
	o.LostProfitPercent = v
}

// GetDaysInStock returns the DaysInStock field value
func (o *CategoryProduct) GetDaysInStock() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DaysInStock
}

// GetDaysInStockOk returns a tuple with the DaysInStock field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetDaysInStockOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysInStock, true
}

// SetDaysInStock sets field value
func (o *CategoryProduct) SetDaysInStock(v int32) {
	o.DaysInStock = v
}

// GetDaysWithSales returns the DaysWithSales field value
func (o *CategoryProduct) GetDaysWithSales() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DaysWithSales
}

// GetDaysWithSalesOk returns a tuple with the DaysWithSales field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetDaysWithSalesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysWithSales, true
}

// SetDaysWithSales sets field value
func (o *CategoryProduct) SetDaysWithSales(v int32) {
	o.DaysWithSales = v
}

// GetAverageIfInStock returns the AverageIfInStock field value
func (o *CategoryProduct) GetAverageIfInStock() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AverageIfInStock
}

// GetAverageIfInStockOk returns a tuple with the AverageIfInStock field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetAverageIfInStockOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AverageIfInStock, true
}

// SetAverageIfInStock sets field value
func (o *CategoryProduct) SetAverageIfInStock(v float32) {
	o.AverageIfInStock = v
}

// GetThumb returns the Thumb field value
func (o *CategoryProduct) GetThumb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Thumb
}

// GetThumbOk returns a tuple with the Thumb field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetThumbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Thumb, true
}

// SetThumb sets field value
func (o *CategoryProduct) SetThumb(v string) {
	o.Thumb = v
}

// GetThumbMiddle returns the ThumbMiddle field value
func (o *CategoryProduct) GetThumbMiddle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbMiddle
}

// GetThumbMiddleOk returns a tuple with the ThumbMiddle field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetThumbMiddleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThumbMiddle, true
}

// SetThumbMiddle sets field value
func (o *CategoryProduct) SetThumbMiddle(v string) {
	o.ThumbMiddle = v
}

// GetIsFbs returns the IsFbs field value
func (o *CategoryProduct) GetIsFbs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IsFbs
}

// GetIsFbsOk returns a tuple with the IsFbs field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetIsFbsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFbs, true
}

// SetIsFbs sets field value
func (o *CategoryProduct) SetIsFbs(v int32) {
	o.IsFbs = v
}

// GetSubjectId returns the SubjectId field value
func (o *CategoryProduct) GetSubjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetSubjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectId, true
}

// SetSubjectId sets field value
func (o *CategoryProduct) SetSubjectId(v int32) {
	o.SubjectId = v
}

// GetUrl returns the Url field value
func (o *CategoryProduct) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CategoryProduct) SetUrl(v string) {
	o.Url = v
}

// GetCountry returns the Country field value
func (o *CategoryProduct) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *CategoryProduct) SetCountry(v string) {
	o.Country = v
}

// GetGender returns the Gender field value
func (o *CategoryProduct) GetGender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gender
}

// GetGenderOk returns a tuple with the Gender field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetGenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gender, true
}

// SetGender sets field value
func (o *CategoryProduct) SetGender(v string) {
	o.Gender = v
}

// GetSkuFirstDate returns the SkuFirstDate field value
func (o *CategoryProduct) GetSkuFirstDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SkuFirstDate
}

// GetSkuFirstDateOk returns a tuple with the SkuFirstDate field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetSkuFirstDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkuFirstDate, true
}

// SetSkuFirstDate sets field value
func (o *CategoryProduct) SetSkuFirstDate(v string) {
	o.SkuFirstDate = v
}

// GetFirstcommentdate returns the Firstcommentdate field value
func (o *CategoryProduct) GetFirstcommentdate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Firstcommentdate
}

// GetFirstcommentdateOk returns a tuple with the Firstcommentdate field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetFirstcommentdateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Firstcommentdate, true
}

// SetFirstcommentdate sets field value
func (o *CategoryProduct) SetFirstcommentdate(v string) {
	o.Firstcommentdate = v
}

// GetPicscount returns the Picscount field value
func (o *CategoryProduct) GetPicscount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Picscount
}

// GetPicscountOk returns a tuple with the Picscount field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetPicscountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Picscount, true
}

// SetPicscount sets field value
func (o *CategoryProduct) SetPicscount(v int32) {
	o.Picscount = v
}

// GetHas3d returns the Has3d field value
func (o *CategoryProduct) GetHas3d() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Has3d
}

// GetHas3dOk returns a tuple with the Has3d field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetHas3dOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Has3d, true
}

// SetHas3d sets field value
func (o *CategoryProduct) SetHas3d(v int32) {
	o.Has3d = v
}

// GetHasvideo returns the Hasvideo field value
func (o *CategoryProduct) GetHasvideo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Hasvideo
}

// GetHasvideoOk returns a tuple with the Hasvideo field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetHasvideoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hasvideo, true
}

// SetHasvideo sets field value
func (o *CategoryProduct) SetHasvideo(v int32) {
	o.Hasvideo = v
}

// GetCommentsvaluation returns the Commentsvaluation field value
func (o *CategoryProduct) GetCommentsvaluation() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Commentsvaluation
}

// GetCommentsvaluationOk returns a tuple with the Commentsvaluation field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetCommentsvaluationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commentsvaluation, true
}

// SetCommentsvaluation sets field value
func (o *CategoryProduct) SetCommentsvaluation(v float32) {
	o.Commentsvaluation = v
}

// GetCardratingval returns the Cardratingval field value
func (o *CategoryProduct) GetCardratingval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cardratingval
}

// GetCardratingvalOk returns a tuple with the Cardratingval field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetCardratingvalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cardratingval, true
}

// SetCardratingval sets field value
func (o *CategoryProduct) SetCardratingval(v int32) {
	o.Cardratingval = v
}

// GetPosition returns the Position field value
func (o *CategoryProduct) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *CategoryProduct) SetPosition(v int32) {
	o.Position = v
}

// GetCategoriesLastCount returns the CategoriesLastCount field value
func (o *CategoryProduct) GetCategoriesLastCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CategoriesLastCount
}

// GetCategoriesLastCountOk returns a tuple with the CategoriesLastCount field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetCategoriesLastCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoriesLastCount, true
}

// SetCategoriesLastCount sets field value
func (o *CategoryProduct) SetCategoriesLastCount(v int32) {
	o.CategoriesLastCount = v
}

// GetGraph returns the Graph field value
func (o *CategoryProduct) GetGraph() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Graph
}

// GetGraphOk returns a tuple with the Graph field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetGraphOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Graph, true
}

// SetGraph sets field value
func (o *CategoryProduct) SetGraph(v []int32) {
	o.Graph = v
}

// GetCategoryGraph returns the CategoryGraph field value
func (o *CategoryProduct) GetCategoryGraph() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.CategoryGraph
}

// GetCategoryGraphOk returns a tuple with the CategoryGraph field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetCategoryGraphOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CategoryGraph, true
}

// SetCategoryGraph sets field value
func (o *CategoryProduct) SetCategoryGraph(v []int32) {
	o.CategoryGraph = v
}

// GetStocksGraph returns the StocksGraph field value
func (o *CategoryProduct) GetStocksGraph() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.StocksGraph
}

// GetStocksGraphOk returns a tuple with the StocksGraph field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetStocksGraphOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StocksGraph, true
}

// SetStocksGraph sets field value
func (o *CategoryProduct) SetStocksGraph(v []int32) {
	o.StocksGraph = v
}

// GetPriceGraph returns the PriceGraph field value
func (o *CategoryProduct) GetPriceGraph() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.PriceGraph
}

// GetPriceGraphOk returns a tuple with the PriceGraph field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetPriceGraphOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceGraph, true
}

// SetPriceGraph sets field value
func (o *CategoryProduct) SetPriceGraph(v []int32) {
	o.PriceGraph = v
}

// GetPurchase returns the Purchase field value
func (o *CategoryProduct) GetPurchase() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Purchase
}

// GetPurchaseOk returns a tuple with the Purchase field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetPurchaseOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purchase, true
}

// SetPurchase sets field value
func (o *CategoryProduct) SetPurchase(v float32) {
	o.Purchase = v
}

// GetPurchaseAfterReturn returns the PurchaseAfterReturn field value
func (o *CategoryProduct) GetPurchaseAfterReturn() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PurchaseAfterReturn
}

// GetPurchaseAfterReturnOk returns a tuple with the PurchaseAfterReturn field value
// and a boolean to check if the value has been set.
func (o *CategoryProduct) GetPurchaseAfterReturnOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseAfterReturn, true
}

// SetPurchaseAfterReturn sets field value
func (o *CategoryProduct) SetPurchaseAfterReturn(v float32) {
	o.PurchaseAfterReturn = v
}

func (o CategoryProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["brand"] = o.Brand
	}
	if true {
		toSerialize["seller"] = o.Seller
	}
	if true {
		toSerialize["supplier_id"] = o.SupplierId
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["category_position"] = o.CategoryPosition
	}
	if o.PosData != nil {
		toSerialize["pos_data"] = o.PosData
	}
	if true {
		toSerialize["color"] = o.Color
	}
	if true {
		toSerialize["balance"] = o.Balance
	}
	if true {
		toSerialize["balance_fbs"] = o.BalanceFbs
	}
	if true {
		toSerialize["comments"] = o.Comments
	}
	if true {
		toSerialize["rating"] = o.Rating
	}
	if true {
		toSerialize["final_price"] = o.FinalPrice
	}
	if true {
		toSerialize["final_price_max"] = o.FinalPriceMax
	}
	if true {
		toSerialize["final_price_min"] = o.FinalPriceMin
	}
	if true {
		toSerialize["final_price_average"] = o.FinalPriceAverage
	}
	if true {
		toSerialize["basic_sale"] = o.BasicSale
	}
	if true {
		toSerialize["basic_price"] = o.BasicPrice
	}
	if true {
		toSerialize["promo_sale"] = o.PromoSale
	}
	if true {
		toSerialize["client_sale"] = o.ClientSale
	}
	if true {
		toSerialize["client_price"] = o.ClientPrice
	}
	if true {
		toSerialize["start_price"] = o.StartPrice
	}
	if true {
		toSerialize["sales"] = o.Sales
	}
	if true {
		toSerialize["sales_per_day_average"] = o.SalesPerDayAverage
	}
	if true {
		toSerialize["revenue"] = o.Revenue
	}
	if true {
		toSerialize["revenue_potential"] = o.RevenuePotential
	}
	if true {
		toSerialize["lost_profit"] = o.LostProfit
	}
	if true {
		toSerialize["lost_profit_percent"] = o.LostProfitPercent
	}
	if true {
		toSerialize["days_in_stock"] = o.DaysInStock
	}
	if true {
		toSerialize["days_with_sales"] = o.DaysWithSales
	}
	if true {
		toSerialize["average_if_in_stock"] = o.AverageIfInStock
	}
	if true {
		toSerialize["thumb"] = o.Thumb
	}
	if true {
		toSerialize["thumb_middle"] = o.ThumbMiddle
	}
	if true {
		toSerialize["is_fbs"] = o.IsFbs
	}
	if true {
		toSerialize["subject_id"] = o.SubjectId
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["country"] = o.Country
	}
	if true {
		toSerialize["gender"] = o.Gender
	}
	if true {
		toSerialize["sku_first_date"] = o.SkuFirstDate
	}
	if true {
		toSerialize["firstcommentdate"] = o.Firstcommentdate
	}
	if true {
		toSerialize["picscount"] = o.Picscount
	}
	if true {
		toSerialize["has3d"] = o.Has3d
	}
	if true {
		toSerialize["hasvideo"] = o.Hasvideo
	}
	if true {
		toSerialize["commentsvaluation"] = o.Commentsvaluation
	}
	if true {
		toSerialize["cardratingval"] = o.Cardratingval
	}
	if true {
		toSerialize["position"] = o.Position
	}
	if true {
		toSerialize["categories_last_count"] = o.CategoriesLastCount
	}
	if true {
		toSerialize["graph"] = o.Graph
	}
	if true {
		toSerialize["category_graph"] = o.CategoryGraph
	}
	if true {
		toSerialize["stocks_graph"] = o.StocksGraph
	}
	if true {
		toSerialize["price_graph"] = o.PriceGraph
	}
	if true {
		toSerialize["purchase"] = o.Purchase
	}
	if true {
		toSerialize["purchase_after_return"] = o.PurchaseAfterReturn
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryProduct struct {
	value *CategoryProduct
	isSet bool
}

func (v NullableCategoryProduct) Get() *CategoryProduct {
	return v.value
}

func (v *NullableCategoryProduct) Set(val *CategoryProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryProduct(val *CategoryProduct) *NullableCategoryProduct {
	return &NullableCategoryProduct{value: val, isSet: true}
}

func (v NullableCategoryProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


