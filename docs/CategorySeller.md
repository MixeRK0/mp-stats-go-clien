# CategorySeller

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название продавца | 
**Items** | **int32** | Число товаров | 
**Sales** | **int32** | Число зафиксированных продаж (единицы) | 
**Revenue** | **float32** | Сумма произведений числа проданных товаров на их стоимость | 
**Comments** | **float32** | Среднее число комментариев | 
**Rating** | **float32** | Средний рейтинг | 

## Methods

### NewCategorySeller

`func NewCategorySeller(name string, items int32, sales int32, revenue float32, comments float32, rating float32, ) *CategorySeller`

NewCategorySeller instantiates a new CategorySeller object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategorySellerWithDefaults

`func NewCategorySellerWithDefaults() *CategorySeller`

NewCategorySellerWithDefaults instantiates a new CategorySeller object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CategorySeller) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategorySeller) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategorySeller) SetName(v string)`

SetName sets Name field to given value.


### GetItems

`func (o *CategorySeller) GetItems() int32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CategorySeller) GetItemsOk() (*int32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CategorySeller) SetItems(v int32)`

SetItems sets Items field to given value.


### GetSales

`func (o *CategorySeller) GetSales() int32`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *CategorySeller) GetSalesOk() (*int32, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *CategorySeller) SetSales(v int32)`

SetSales sets Sales field to given value.


### GetRevenue

`func (o *CategorySeller) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CategorySeller) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CategorySeller) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.


### GetComments

`func (o *CategorySeller) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CategorySeller) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CategorySeller) SetComments(v float32)`

SetComments sets Comments field to given value.


### GetRating

`func (o *CategorySeller) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *CategorySeller) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *CategorySeller) SetRating(v float32)`

SetRating sets Rating field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


