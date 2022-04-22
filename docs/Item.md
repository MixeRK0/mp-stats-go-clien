# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | [**ItemItem**](ItemItem.md) |  | 
**Photos** | [**[]ItemPhotos**](ItemPhotos.md) |  | 

## Methods

### NewItem

`func NewItem(item ItemItem, photos []ItemPhotos, ) *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *Item) GetItem() ItemItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *Item) GetItemOk() (*ItemItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *Item) SetItem(v ItemItem)`

SetItem sets Item field to given value.


### GetPhotos

`func (o *Item) GetPhotos() []ItemPhotos`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *Item) GetPhotosOk() (*[]ItemPhotos, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *Item) SetPhotos(v []ItemPhotos)`

SetPhotos sets Photos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


