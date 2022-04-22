# ItemSale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoData** | **int32** | 0: OK, 1: Продаж и остатков не зафиксировано | 
**Data** | **string** | Дата | 
**Balance** | **int32** | Остаток | 
**Sales** | **int32** | Продажи | 
**Rating** | **float32** | Рейтинг | 
**Price** | **float32** | Цена | 
**FinalPrice** | **float32** | Со скидкой | 
**IsNew** | **int32** | Новый товар (?) | 
**Comments** | **float32** | Комментариев | 

## Methods

### NewItemSale

`func NewItemSale(noData int32, data string, balance int32, sales int32, rating float32, price float32, finalPrice float32, isNew int32, comments float32, ) *ItemSale`

NewItemSale instantiates a new ItemSale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemSaleWithDefaults

`func NewItemSaleWithDefaults() *ItemSale`

NewItemSaleWithDefaults instantiates a new ItemSale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoData

`func (o *ItemSale) GetNoData() int32`

GetNoData returns the NoData field if non-nil, zero value otherwise.

### GetNoDataOk

`func (o *ItemSale) GetNoDataOk() (*int32, bool)`

GetNoDataOk returns a tuple with the NoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoData

`func (o *ItemSale) SetNoData(v int32)`

SetNoData sets NoData field to given value.


### GetData

`func (o *ItemSale) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ItemSale) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ItemSale) SetData(v string)`

SetData sets Data field to given value.


### GetBalance

`func (o *ItemSale) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ItemSale) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ItemSale) SetBalance(v int32)`

SetBalance sets Balance field to given value.


### GetSales

`func (o *ItemSale) GetSales() int32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *ItemSale) GetSalesOk() (*int32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *ItemSale) SetSales(v int32)`

SetSales sets Sales field to given value.


### GetRating

`func (o *ItemSale) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ItemSale) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ItemSale) SetRating(v float32)`

SetRating sets Rating field to given value.


### GetPrice

`func (o *ItemSale) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ItemSale) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ItemSale) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetFinalPrice

`func (o *ItemSale) GetFinalPrice() float32`

GetFinalPrice returns the FinalPrice field if non-nil, zero value otherwise.

### GetFinalPriceOk

`func (o *ItemSale) GetFinalPriceOk() (*float32, bool)`

GetFinalPriceOk returns a tuple with the FinalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPrice

`func (o *ItemSale) SetFinalPrice(v float32)`

SetFinalPrice sets FinalPrice field to given value.


### GetIsNew

`func (o *ItemSale) GetIsNew() int32`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *ItemSale) GetIsNewOk() (*int32, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *ItemSale) SetIsNew(v int32)`

SetIsNew sets IsNew field to given value.


### GetComments

`func (o *ItemSale) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ItemSale) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ItemSale) SetComments(v float32)`

SetComments sets Comments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


