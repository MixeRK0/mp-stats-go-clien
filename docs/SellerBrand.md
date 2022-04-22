# SellerBrand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название бренда | 
**Items** | **int32** | Число товаров | 
**Sales** | **int32** | Число зафиксированных продаж (единицы) | 
**Revenue** | **float32** | Сумма произведений числа проданных товаров на их стоимость | 
**Comments** | **float32** | Среднее число комментариев | 
**Rating** | **float32** | Средний рейтинг | 

## Methods

### NewSellerBrand

`func NewSellerBrand(name string, items int32, sales int32, revenue float32, comments float32, rating float32, ) *SellerBrand`

NewSellerBrand instantiates a new SellerBrand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellerBrandWithDefaults

`func NewSellerBrandWithDefaults() *SellerBrand`

NewSellerBrandWithDefaults instantiates a new SellerBrand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SellerBrand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SellerBrand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SellerBrand) SetName(v string)`

SetName sets Name field to given value.


### GetItems

`func (o *SellerBrand) GetItems() int32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SellerBrand) GetItemsOk() (*int32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SellerBrand) SetItems(v int32)`

SetItems sets Items field to given value.


### GetSales

`func (o *SellerBrand) GetSales() int32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *SellerBrand) GetSalesOk() (*int32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *SellerBrand) SetSales(v int32)`

SetSales sets Sales field to given value.


### GetRevenue

`func (o *SellerBrand) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *SellerBrand) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *SellerBrand) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.


### GetComments

`func (o *SellerBrand) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *SellerBrand) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *SellerBrand) SetComments(v float32)`

SetComments sets Comments field to given value.


### GetRating

`func (o *SellerBrand) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *SellerBrand) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *SellerBrand) SetRating(v float32)`

SetRating sets Rating field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


