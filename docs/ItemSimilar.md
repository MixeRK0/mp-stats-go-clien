# ItemSimilar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **float32** | Последний зафиксированный остаток на складе | 
**Brand** | **string** | Название бренда | 
**Comments** | **float32** | Количество комментариев | 
**FinalPrice** | **float32** | Последняя зафиксированная цена | 
**Graph** | **[]int32** | График продаж | 
**Id** | **float32** | SKU | 
**Name** | **string** | Название товара | 
**Rating** | **float32** | Рейтинг товара | 
**Revenue** | **float32** | Выручка за период | 
**Sales** | **float32** | Количество проданных единиц товара за период | 
**Seller** | **string** | Название продавца | 
**Thumb** | **string** | Изображение товара | 
**ThumbMiddle** | **string** | Изображение товара (среднее) | 

## Methods

### NewItemSimilar

`func NewItemSimilar(balance float32, brand string, comments float32, finalPrice float32, graph []int32, id float32, name string, rating float32, revenue float32, sales float32, seller string, thumb string, thumbMiddle string, ) *ItemSimilar`

NewItemSimilar instantiates a new ItemSimilar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemSimilarWithDefaults

`func NewItemSimilarWithDefaults() *ItemSimilar`

NewItemSimilarWithDefaults instantiates a new ItemSimilar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *ItemSimilar) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ItemSimilar) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ItemSimilar) SetBalance(v float32)`

SetBalance sets Balance field to given value.


### GetBrand

`func (o *ItemSimilar) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ItemSimilar) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ItemSimilar) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetComments

`func (o *ItemSimilar) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ItemSimilar) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ItemSimilar) SetComments(v float32)`

SetComments sets Comments field to given value.


### GetFinalPrice

`func (o *ItemSimilar) GetFinalPrice() float32`

GetFinalPrice returns the FinalPrice field if non-nil, zero value otherwise.

### GetFinalPriceOk

`func (o *ItemSimilar) GetFinalPriceOk() (*float32, bool)`

GetFinalPriceOk returns a tuple with the FinalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPrice

`func (o *ItemSimilar) SetFinalPrice(v float32)`

SetFinalPrice sets FinalPrice field to given value.


### GetGraph

`func (o *ItemSimilar) GetGraph() []int32`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *ItemSimilar) GetGraphOk() (*[]int32, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *ItemSimilar) SetGraph(v []int32)`

SetGraph sets Graph field to given value.


### GetId

`func (o *ItemSimilar) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemSimilar) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemSimilar) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *ItemSimilar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemSimilar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemSimilar) SetName(v string)`

SetName sets Name field to given value.


### GetRating

`func (o *ItemSimilar) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ItemSimilar) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ItemSimilar) SetRating(v float32)`

SetRating sets Rating field to given value.


### GetRevenue

`func (o *ItemSimilar) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *ItemSimilar) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *ItemSimilar) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.


### GetSales

`func (o *ItemSimilar) GetSales() float32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *ItemSimilar) GetSalesOk() (*float32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *ItemSimilar) SetSales(v float32)`

SetSales sets Sales field to given value.


### GetSeller

`func (o *ItemSimilar) GetSeller() string`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *ItemSimilar) GetSellerOk() (*string, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *ItemSimilar) SetSeller(v string)`

SetSeller sets Seller field to given value.


### GetThumb

`func (o *ItemSimilar) GetThumb() string`

GetThumb returns the Thumb field if non-nil, zero value otherwise.

### GetThumbOk

`func (o *ItemSimilar) GetThumbOk() (*string, bool)`

GetThumbOk returns a tuple with the Thumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumb

`func (o *ItemSimilar) SetThumb(v string)`

SetThumb sets Thumb field to given value.


### GetThumbMiddle

`func (o *ItemSimilar) GetThumbMiddle() string`

GetThumbMiddle returns the ThumbMiddle field if non-nil, zero value otherwise.

### GetThumbMiddleOk

`func (o *ItemSimilar) GetThumbMiddleOk() (*string, bool)`

GetThumbMiddleOk returns a tuple with the ThumbMiddle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbMiddle

`func (o *ItemSimilar) SetThumbMiddle(v string)`

SetThumbMiddle sets ThumbMiddle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


