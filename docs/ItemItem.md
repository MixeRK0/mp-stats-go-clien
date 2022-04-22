# ItemItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор товара (SKU) | 
**Name** | **string** | Название | 
**FullName** | **string** | Полное название | 
**Link** | **string** | Ссылка на площадке | 
**Brand** | **string** | Бренд | 
**Seller** | **string** | Продавец | 
**Rating** | **float32** | Рейтинг | 
**Comments** | **float32** | Комментариев | 
**Price** | **float32** | Цена | 
**FinalPrice** | **float32** | Цена с учетом скидки | 
**Discount** | **float32** | Скидка | 
**Updated** | **string** | Обновлено | 
**IsNew** | **float32** | Новинка | 

## Methods

### NewItemItem

`func NewItemItem(id string, name string, fullName string, link string, brand string, seller string, rating float32, comments float32, price float32, finalPrice float32, discount float32, updated string, isNew float32, ) *ItemItem`

NewItemItem instantiates a new ItemItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemItemWithDefaults

`func NewItemItemWithDefaults() *ItemItem`

NewItemItemWithDefaults instantiates a new ItemItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ItemItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemItem) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *ItemItem) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ItemItem) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ItemItem) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetLink

`func (o *ItemItem) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ItemItem) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ItemItem) SetLink(v string)`

SetLink sets Link field to given value.


### GetBrand

`func (o *ItemItem) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ItemItem) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ItemItem) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetSeller

`func (o *ItemItem) GetSeller() string`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *ItemItem) GetSellerOk() (*string, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *ItemItem) SetSeller(v string)`

SetSeller sets Seller field to given value.


### GetRating

`func (o *ItemItem) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ItemItem) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ItemItem) SetRating(v float32)`

SetRating sets Rating field to given value.


### GetComments

`func (o *ItemItem) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ItemItem) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ItemItem) SetComments(v float32)`

SetComments sets Comments field to given value.


### GetPrice

`func (o *ItemItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ItemItem) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ItemItem) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetFinalPrice

`func (o *ItemItem) GetFinalPrice() float32`

GetFinalPrice returns the FinalPrice field if non-nil, zero value otherwise.

### GetFinalPriceOk

`func (o *ItemItem) GetFinalPriceOk() (*float32, bool)`

GetFinalPriceOk returns a tuple with the FinalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPrice

`func (o *ItemItem) SetFinalPrice(v float32)`

SetFinalPrice sets FinalPrice field to given value.


### GetDiscount

`func (o *ItemItem) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *ItemItem) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *ItemItem) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.


### GetUpdated

`func (o *ItemItem) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ItemItem) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ItemItem) SetUpdated(v string)`

SetUpdated sets Updated field to given value.


### GetIsNew

`func (o *ItemItem) GetIsNew() float32`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *ItemItem) GetIsNewOk() (*float32, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *ItemItem) SetIsNew(v float32)`

SetIsNew sets IsNew field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


