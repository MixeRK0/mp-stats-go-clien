# \WbApi

All URIs are relative to *https://mpstats.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBrandByDate**](WbApi.md#GetBrandByDate) | **Get** /wb/get/brand/by_date | Данные по дням по бренду
[**GetBrandCategories**](WbApi.md#GetBrandCategories) | **Get** /wb/get/brand/categories | Категории бренда
[**GetBrandDetailedItems**](WbApi.md#GetBrandDetailedItems) | **Post** /wb/get/brand | Товары бренда
[**GetBrandPriceSegmentation**](WbApi.md#GetBrandPriceSegmentation) | **Get** /wb/get/brand/price_segmentation | Ценовая сегментация бренда
[**GetBrandSellers**](WbApi.md#GetBrandSellers) | **Get** /wb/get/brand/sellers | Продавцы бренда
[**GetCategories**](WbApi.md#GetCategories) | **Get** /wb/get/categories | Текущие категории
[**GetCategoryBrands**](WbApi.md#GetCategoryBrands) | **Get** /wb/get/category/brands | Бренды категории
[**GetCategoryByDate**](WbApi.md#GetCategoryByDate) | **Get** /wb/get/category/by_date | Данные по дням по категории
[**GetCategoryItems**](WbApi.md#GetCategoryItems) | **Post** /wb/get/category | Товары категории
[**GetCategoryPriceSegmentation**](WbApi.md#GetCategoryPriceSegmentation) | **Get** /wb/get/category/price_segmentation | Ценовая сегментация категории
[**GetCategorySellers**](WbApi.md#GetCategorySellers) | **Get** /wb/get/category/sellers | Продавцы категории
[**GetCategorySubcategories**](WbApi.md#GetCategorySubcategories) | **Get** /wb/get/category/subcategories | Подкатегории категории
[**GetCategorySubjects**](WbApi.md#GetCategorySubjects) | **Get** /wb/get/category/items | Предметы категории
[**GetCategoryTrends**](WbApi.md#GetCategoryTrends) | **Get** /wb/get/category/trends | Тренды категории
[**GetItemBalanceByDay**](WbApi.md#GetItemBalanceByDay) | **Get** /wb/get/item/{sku}/balance_by_day | Продажи и остатки товарной позиции за сутки
[**GetItemByCategory**](WbApi.md#GetItemByCategory) | **Get** /wb/get/item/{sku}/by_category | История товарной позиции по категориям
[**GetItemBySku**](WbApi.md#GetItemBySku) | **Get** /wb/get/item/{sku} | Одна товарная позиция по SKU
[**GetItemOrdersByRegion**](WbApi.md#GetItemOrdersByRegion) | **Get** /wb/get/item/{sku}/orders_by_region | История заказов и остатков товарной позиции по складам
[**GetItemOrdersBySize**](WbApi.md#GetItemOrdersBySize) | **Get** /wb/get/item/{sku}/orders_by_size | История заказов и остатков товарной позиции по размерам
[**GetItemSales**](WbApi.md#GetItemSales) | **Get** /wb/get/item/{sku}/sales | Продажи и остатки товарной позиции
[**GetItemSimilar**](WbApi.md#GetItemSimilar) | **Get** /wb/get/item/{sku}/similar | Похожие товарные позиции
[**GetItemsBatch**](WbApi.md#GetItemsBatch) | **Post** /wb/get/items/batch | Товарные позиции по SKU
[**GetSellerBrands**](WbApi.md#GetSellerBrands) | **Get** /wb/get/seller/brands | Бренды продавца
[**GetSellerByDate**](WbApi.md#GetSellerByDate) | **Get** /wb/get/seller/by_date | Данные по дням по продавцу
[**GetSellerCategories**](WbApi.md#GetSellerCategories) | **Get** /wb/get/seller/categories | Категории продавца
[**GetSellerDetailedItems**](WbApi.md#GetSellerDetailedItems) | **Post** /wb/get/seller | Товары продавца
[**GetSellerPriceSegmentation**](WbApi.md#GetSellerPriceSegmentation) | **Get** /wb/get/seller/price_segmentation | Ценовая сегментация продавца
[**GetSimilarBrands**](WbApi.md#GetSimilarBrands) | **Get** /wb/get/similar/brands | Бренды похожего товара
[**GetSimilarByDate**](WbApi.md#GetSimilarByDate) | **Get** /wb/get/similar/by_date | По дням
[**GetSimilarCategories**](WbApi.md#GetSimilarCategories) | **Get** /wb/get/similar/categories | Категории похожего товара
[**GetSimilarDetailedItems**](WbApi.md#GetSimilarDetailedItems) | **Post** /wb/get/similar | Товары по похожему товару
[**GetSimilarSellers**](WbApi.md#GetSimilarSellers) | **Get** /wb/get/similar/sellers | Продавцы похожего товара
[**GetSubjectItems**](WbApi.md#GetSubjectItems) | **Post** /wb/get/subject | Получить список товаров по предмету
[**PostWbGetSearchCategories**](WbApi.md#PostWbGetSearchCategories) | **Post** /wb/get/search/categories | GetSearchCategories
[**PostWbGetSearchItems**](WbApi.md#PostWbGetSearchItems) | **Post** /wb/get/search | GetSearchItems



## GetBrandByDate

> []BrandByDate GetBrandByDate(ctx).Path(path).D1(d1).D2(d2).Execute()

Данные по дням по бренду



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "Mango" // string | Бренд
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetBrandByDate(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetBrandByDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandByDate`: []BrandByDate
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetBrandByDate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandByDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Бренд | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]BrandByDate**](BrandByDate.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandCategories

> []BrandCategory GetBrandCategories(ctx).Path(path).D1(d1).D2(d2).Execute()

Категории бренда



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "Mango" // string | Бренд
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetBrandCategories(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetBrandCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandCategories`: []BrandCategory
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetBrandCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Бренд | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]BrandCategory**](BrandCategory.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandDetailedItems

> InlineResponse200 GetBrandDetailedItems(ctx).Path(path).D1(d1).D2(d2).GetItemsRequestBody(getItemsRequestBody).Execute()

Товары бренда



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "Mango" // string | Бренд
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)
    getItemsRequestBody := *openapiclient.NewGetItemsRequestBody(int32(123), int32(123), interface{}(123), []string{"GroupKeys_example"}, []string{"PivotCols_example"}, false, []string{"RowGroupCols_example"}, []openapiclient.SortModelItem{*openapiclient.NewSortModelItem("ColId_example", "Sort_example")}, []string{"ValueCols_example"}) // GetItemsRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetBrandDetailedItems(context.Background()).Path(path).D1(d1).D2(d2).GetItemsRequestBody(getItemsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetBrandDetailedItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandDetailedItems`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetBrandDetailedItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandDetailedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Бренд | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 
 **getItemsRequestBody** | [**GetItemsRequestBody**](GetItemsRequestBody.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandPriceSegmentation

> []BrandPriceSegmentation GetBrandPriceSegmentation(ctx).Path(path).D1(d1).D2(d2).Execute()

Ценовая сегментация бренда



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "Mango" // string | Бренд
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetBrandPriceSegmentation(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetBrandPriceSegmentation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandPriceSegmentation`: []BrandPriceSegmentation
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetBrandPriceSegmentation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandPriceSegmentationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Бренд | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]BrandPriceSegmentation**](BrandPriceSegmentation.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandSellers

> []BrandSeller GetBrandSellers(ctx).Path(path).D1(d1).D2(d2).Execute()

Продавцы бренда



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "Mango" // string | Бренд
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetBrandSellers(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetBrandSellers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandSellers`: []BrandSeller
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetBrandSellers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandSellersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Бренд | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]BrandSeller**](BrandSeller.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategories

> []CategoryPath GetCategories(ctx).Execute()

Текущие категории



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetCategories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategories`: []CategoryPath
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoriesRequest struct via the builder pattern


### Return type

[**[]CategoryPath**](category-path.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryBrands

> []CategoryBrand GetCategoryBrands(ctx).Path(path).D1(d1).D2(d2).Execute()

Бренды категории



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "Женщинам/Одежда" // string | Категория
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetCategoryBrands(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetCategoryBrands``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategoryBrands`: []CategoryBrand
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetCategoryBrands`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryBrandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Категория | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]CategoryBrand**](CategoryBrand.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryByDate

> []CategoryByDate GetCategoryByDate(ctx).Path(path).D1(d1).D2(d2).Execute()

Данные по дням по категории



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "Женщинам/Одежда" // string | Категория
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetCategoryByDate(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetCategoryByDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategoryByDate`: []CategoryByDate
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetCategoryByDate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryByDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Категория | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]CategoryByDate**](CategoryByDate.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryItems

> InlineResponse200 GetCategoryItems(ctx).Path(path).D1(d1).D2(d2).GetItemsRequestBody(getItemsRequestBody).Execute()

Товары категории



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "Женщинам/Одежда" // string | Категория
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)
    getItemsRequestBody := *openapiclient.NewGetItemsRequestBody(int32(123), int32(123), interface{}(123), []string{"GroupKeys_example"}, []string{"PivotCols_example"}, false, []string{"RowGroupCols_example"}, []openapiclient.SortModelItem{*openapiclient.NewSortModelItem("ColId_example", "Sort_example")}, []string{"ValueCols_example"}) // GetItemsRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetCategoryItems(context.Background()).Path(path).D1(d1).D2(d2).GetItemsRequestBody(getItemsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetCategoryItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategoryItems`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetCategoryItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Категория | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 
 **getItemsRequestBody** | [**GetItemsRequestBody**](GetItemsRequestBody.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryPriceSegmentation

> []CategoryPriceSegmentation GetCategoryPriceSegmentation(ctx).Path(path).D1(d1).D2(d2).Execute()

Ценовая сегментация категории



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "Женщинам/Одежда" // string | Категория
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetCategoryPriceSegmentation(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetCategoryPriceSegmentation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategoryPriceSegmentation`: []CategoryPriceSegmentation
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetCategoryPriceSegmentation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryPriceSegmentationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Категория | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]CategoryPriceSegmentation**](CategoryPriceSegmentation.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategorySellers

> []CategorySeller GetCategorySellers(ctx).Path(path).D1(d1).D2(d2).Execute()

Продавцы категории



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "Женщинам/Одежда" // string | Категория
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetCategorySellers(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetCategorySellers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategorySellers`: []CategorySeller
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetCategorySellers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategorySellersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Категория | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]CategorySeller**](CategorySeller.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategorySubcategories

> []CategorySubcategory GetCategorySubcategories(ctx).Path(path).D1(d1).D2(d2).Execute()

Подкатегории категории



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "Женщинам/Одежда" // string | Категория
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetCategorySubcategories(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetCategorySubcategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategorySubcategories`: []CategorySubcategory
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetCategorySubcategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategorySubcategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Категория | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]CategorySubcategory**](CategorySubcategory.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategorySubjects

> []CategorySubject GetCategorySubjects(ctx).Path(path).D1(d1).D2(d2).Execute()

Предметы категории



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "Женщинам/Одежда" // string | Категория
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetCategorySubjects(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetCategorySubjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategorySubjects`: []CategorySubject
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetCategorySubjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategorySubjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Категория | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]CategorySubject**](CategorySubject.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryTrends

> []CategoryTrend GetCategoryTrends(ctx).Path(path).D1(d1).D2(d2).Execute()

Тренды категории



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "Женщинам/Одежда" // string | Категория
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetCategoryTrends(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetCategoryTrends``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategoryTrends`: []CategoryTrend
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetCategoryTrends`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryTrendsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Категория | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]CategoryTrend**](CategoryTrend.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemBalanceByDay

> []ItemBalanceByDay GetItemBalanceByDay(ctx, sku).D(d).Execute()

Продажи и остатки товарной позиции за сутки



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    sku := int32(56) // int32 | SKU
    d := time.Now() // string | Дата (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetItemBalanceByDay(context.Background(), sku).D(d).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetItemBalanceByDay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemBalanceByDay`: []ItemBalanceByDay
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetItemBalanceByDay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **int32** | SKU | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemBalanceByDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **d** | **string** | Дата | 

### Return type

[**[]ItemBalanceByDay**](ItemBalanceByDay.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemByCategory

> []ItemByCategory GetItemByCategory(ctx, sku).D1(d1).D2(d2).Execute()

История товарной позиции по категориям



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    sku := int32(56) // int32 | SKU
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetItemByCategory(context.Background(), sku).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetItemByCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemByCategory`: []ItemByCategory
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetItemByCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **int32** | SKU | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemByCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]ItemByCategory**](ItemByCategory.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemBySku

> Item GetItemBySku(ctx, sku).Execute()

Одна товарная позиция по SKU



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sku := int32(56) // int32 | SKU

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetItemBySku(context.Background(), sku).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetItemBySku``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemBySku`: Item
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetItemBySku`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **int32** | SKU | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemBySkuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Item**](Item.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemOrdersByRegion

> []map[string]interface{} GetItemOrdersByRegion(ctx, sku).D1(d1).D2(d2).Execute()

История заказов и остатков товарной позиции по складам



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    sku := int32(56) // int32 | SKU
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetItemOrdersByRegion(context.Background(), sku).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetItemOrdersByRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemOrdersByRegion`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetItemOrdersByRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **int32** | SKU | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemOrdersByRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

**[]map[string]interface{}**

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemOrdersBySize

> map[string]interface{} GetItemOrdersBySize(ctx, sku).D1(d1).D2(d2).Execute()

История заказов и остатков товарной позиции по размерам



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    sku := int32(56) // int32 | SKU
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetItemOrdersBySize(context.Background(), sku).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetItemOrdersBySize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemOrdersBySize`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetItemOrdersBySize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **int32** | SKU | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemOrdersBySizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

**map[string]interface{}**

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemSales

> []ItemSale GetItemSales(ctx, sku).D1(d1).D2(d2).Execute()

Продажи и остатки товарной позиции



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    sku := int32(56) // int32 | SKU
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetItemSales(context.Background(), sku).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetItemSales``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemSales`: []ItemSale
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetItemSales`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **int32** | SKU | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemSalesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]ItemSale**](ItemSale.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemSimilar

> ItemSimilar GetItemSimilar(ctx, sku).D1(d1).D2(d2).Execute()

Похожие товарные позиции



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    sku := int32(56) // int32 | SKU
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetItemSimilar(context.Background(), sku).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetItemSimilar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemSimilar`: ItemSimilar
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetItemSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **int32** | SKU | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**ItemSimilar**](ItemSimilar.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsBatch

> []Item GetItemsBatch(ctx).GetItemsBatchRequestBody(getItemsBatchRequestBody).Execute()

Товарные позиции по SKU



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getItemsBatchRequestBody := *openapiclient.NewGetItemsBatchRequestBody([]int32{int32(123)}) // GetItemsBatchRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetItemsBatch(context.Background()).GetItemsBatchRequestBody(getItemsBatchRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetItemsBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemsBatch`: []Item
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetItemsBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getItemsBatchRequestBody** | [**GetItemsBatchRequestBody**](GetItemsBatchRequestBody.md) |  | 

### Return type

[**[]Item**](Item.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSellerBrands

> []SellerBrand GetSellerBrands(ctx).Path(path).D1(d1).D2(d2).Execute()

Бренды продавца



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "ВАЙЛДБЕРРИЗ ООО" // string | Продавец
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetSellerBrands(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetSellerBrands``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSellerBrands`: []SellerBrand
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetSellerBrands`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSellerBrandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Продавец | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]SellerBrand**](SellerBrand.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSellerByDate

> []SellerByDate GetSellerByDate(ctx).Path(path).D1(d1).D2(d2).Execute()

Данные по дням по продавцу



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "ВАЙЛДБЕРРИЗ ООО" // string | Продавец
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetSellerByDate(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetSellerByDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSellerByDate`: []SellerByDate
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetSellerByDate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSellerByDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Продавец | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]SellerByDate**](SellerByDate.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSellerCategories

> []SellerCategory GetSellerCategories(ctx).Path(path).D1(d1).D2(d2).Execute()

Категории продавца



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "ВАЙЛДБЕРРИЗ ООО" // string | Продавец
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetSellerCategories(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetSellerCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSellerCategories`: []SellerCategory
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetSellerCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSellerCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Продавец | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]SellerCategory**](SellerCategory.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSellerDetailedItems

> InlineResponse200 GetSellerDetailedItems(ctx).Path(path).D1(d1).D2(d2).GetItemsRequestBody(getItemsRequestBody).Execute()

Товары продавца



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "ВАЙЛДБЕРРИЗ ООО" // string | Продавец
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)
    getItemsRequestBody := *openapiclient.NewGetItemsRequestBody(int32(123), int32(123), interface{}(123), []string{"GroupKeys_example"}, []string{"PivotCols_example"}, false, []string{"RowGroupCols_example"}, []openapiclient.SortModelItem{*openapiclient.NewSortModelItem("ColId_example", "Sort_example")}, []string{"ValueCols_example"}) // GetItemsRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetSellerDetailedItems(context.Background()).Path(path).D1(d1).D2(d2).GetItemsRequestBody(getItemsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetSellerDetailedItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSellerDetailedItems`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetSellerDetailedItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSellerDetailedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Продавец | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 
 **getItemsRequestBody** | [**GetItemsRequestBody**](GetItemsRequestBody.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSellerPriceSegmentation

> []SellerPriceSegmentation GetSellerPriceSegmentation(ctx).Path(path).D1(d1).D2(d2).Execute()

Ценовая сегментация продавца



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "ВАЙЛДБЕРРИЗ ООО" // string | Продавец
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetSellerPriceSegmentation(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetSellerPriceSegmentation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSellerPriceSegmentation`: []SellerPriceSegmentation
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetSellerPriceSegmentation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSellerPriceSegmentationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Продавец | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]SellerPriceSegmentation**](SellerPriceSegmentation.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimilarBrands

> []SimilarBrand GetSimilarBrands(ctx).Path(path).D1(d1).D2(d2).Execute()

Бренды похожего товара



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "6929090" // string | SKU
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetSimilarBrands(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetSimilarBrands``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSimilarBrands`: []SimilarBrand
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetSimilarBrands`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarBrandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | SKU | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]SimilarBrand**](SimilarBrand.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimilarByDate

> []SimilarByDate GetSimilarByDate(ctx).Path(path).D1(d1).D2(d2).Execute()

По дням



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "6929090" // string | SKU
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetSimilarByDate(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetSimilarByDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSimilarByDate`: []SimilarByDate
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetSimilarByDate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarByDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | SKU | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]SimilarByDate**](SimilarByDate.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimilarCategories

> []SimilarCategory GetSimilarCategories(ctx).Path(path).D1(d1).D2(d2).Execute()

Категории похожего товара



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "6929090" // string | SKU
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetSimilarCategories(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetSimilarCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSimilarCategories`: []SimilarCategory
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetSimilarCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | SKU | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]SimilarCategory**](SimilarCategory.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimilarDetailedItems

> InlineResponse200 GetSimilarDetailedItems(ctx).Path(path).D1(d1).D2(d2).GetItemsRequestBody(getItemsRequestBody).Execute()

Товары по похожему товару



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "6929090" // string | SKU
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)
    getItemsRequestBody := *openapiclient.NewGetItemsRequestBody(int32(123), int32(123), interface{}(123), []string{"GroupKeys_example"}, []string{"PivotCols_example"}, false, []string{"RowGroupCols_example"}, []openapiclient.SortModelItem{*openapiclient.NewSortModelItem("ColId_example", "Sort_example")}, []string{"ValueCols_example"}) // GetItemsRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetSimilarDetailedItems(context.Background()).Path(path).D1(d1).D2(d2).GetItemsRequestBody(getItemsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetSimilarDetailedItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSimilarDetailedItems`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetSimilarDetailedItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarDetailedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | SKU | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 
 **getItemsRequestBody** | [**GetItemsRequestBody**](GetItemsRequestBody.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimilarSellers

> []SimilarSeller GetSimilarSellers(ctx).Path(path).D1(d1).D2(d2).Execute()

Продавцы похожего товара



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "6929090" // string | SKU
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetSimilarSellers(context.Background()).Path(path).D1(d1).D2(d2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetSimilarSellers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSimilarSellers`: []SimilarSeller
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetSimilarSellers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarSellersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | SKU | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 

### Return type

[**[]SimilarSeller**](SimilarSeller.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubjectItems

> InlineResponse200 GetSubjectItems(ctx).Path(path).D1(d1).D2(d2).GetItemsRequestBody(getItemsRequestBody).Execute()

Получить список товаров по предмету



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "Женщинам/Одежда" // string | Категория
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)
    getItemsRequestBody := *openapiclient.NewGetItemsRequestBody(int32(123), int32(123), interface{}(123), []string{"GroupKeys_example"}, []string{"PivotCols_example"}, false, []string{"RowGroupCols_example"}, []openapiclient.SortModelItem{*openapiclient.NewSortModelItem("ColId_example", "Sort_example")}, []string{"ValueCols_example"}) // GetItemsRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.GetSubjectItems(context.Background()).Path(path).D1(d1).D2(d2).GetItemsRequestBody(getItemsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.GetSubjectItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubjectItems`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `WbApi.GetSubjectItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubjectItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Категория | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 
 **getItemsRequestBody** | [**GetItemsRequestBody**](GetItemsRequestBody.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWbGetSearchCategories

> []SearchCategoriesElement PostWbGetSearchCategories(ctx).Path(path).D1(d1).D2(d2).TplsRequestBody(tplsRequestBody).Execute()

GetSearchCategories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Поисковой запрос
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)
    tplsRequestBody := *openapiclient.NewTplsRequestBody("Tpls_example") // TplsRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.PostWbGetSearchCategories(context.Background()).Path(path).D1(d1).D2(d2).TplsRequestBody(tplsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.PostWbGetSearchCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWbGetSearchCategories`: []SearchCategoriesElement
    fmt.Fprintf(os.Stdout, "Response from `WbApi.PostWbGetSearchCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWbGetSearchCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Поисковой запрос | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 
 **tplsRequestBody** | [**TplsRequestBody**](TplsRequestBody.md) |  | 

### Return type

[**[]SearchCategoriesElement**](SearchCategoriesElement.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWbGetSearchItems

> SearchItems PostWbGetSearchItems(ctx).Path(path).D1(d1).D2(d2).GetSearchItemsRequestBody(getSearchItemsRequestBody).Execute()

GetSearchItems



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    path := "path_example" // string | Поисковой запрос
    d1 := time.Now() // string | Дата начала периода (optional)
    d2 := time.Now() // string | Дата окончания периода (optional)
    getSearchItemsRequestBody := *openapiclient.NewGetSearchItemsRequestBody(int32(123), int32(123), interface{}(123), []string{"GroupKeys_example"}, []string{"PivotCols_example"}, false, []string{"RowGroupCols_example"}, []openapiclient.SortModelItem{*openapiclient.NewSortModelItem("ColId_example", "Sort_example")}, "Tpl_example", []string{"ValueCols_example"}) // GetSearchItemsRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WbApi.PostWbGetSearchItems(context.Background()).Path(path).D1(d1).D2(d2).GetSearchItemsRequestBody(getSearchItemsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WbApi.PostWbGetSearchItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWbGetSearchItems`: SearchItems
    fmt.Fprintf(os.Stdout, "Response from `WbApi.PostWbGetSearchItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWbGetSearchItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Поисковой запрос | 
 **d1** | **string** | Дата начала периода | 
 **d2** | **string** | Дата окончания периода | 
 **getSearchItemsRequestBody** | [**GetSearchItemsRequestBody**](GetSearchItemsRequestBody.md) |  | 

### Return type

[**SearchItems**](SearchItems.md)

### Authorization

[Header-token](../README.md#Header-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

