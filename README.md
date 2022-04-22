# Go API client for openapi

MPStats API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://mpstats.io/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*WbApi* | [**GetBrandByDate**](docs/WbApi.md#getbrandbydate) | **Get** /wb/get/brand/by_date | Данные по дням по бренду
*WbApi* | [**GetBrandCategories**](docs/WbApi.md#getbrandcategories) | **Get** /wb/get/brand/categories | Категории бренда
*WbApi* | [**GetBrandDetailedItems**](docs/WbApi.md#getbranddetaileditems) | **Post** /wb/get/brand | Товары бренда
*WbApi* | [**GetBrandPriceSegmentation**](docs/WbApi.md#getbrandpricesegmentation) | **Get** /wb/get/brand/price_segmentation | Ценовая сегментация бренда
*WbApi* | [**GetBrandSellers**](docs/WbApi.md#getbrandsellers) | **Get** /wb/get/brand/sellers | Продавцы бренда
*WbApi* | [**GetCategories**](docs/WbApi.md#getcategories) | **Get** /wb/get/categories | Текущие категории
*WbApi* | [**GetCategoryBrands**](docs/WbApi.md#getcategorybrands) | **Get** /wb/get/category/brands | Бренды категории
*WbApi* | [**GetCategoryByDate**](docs/WbApi.md#getcategorybydate) | **Get** /wb/get/category/by_date | Данные по дням по категории
*WbApi* | [**GetCategoryDetailedItems**](docs/WbApi.md#getcategorydetaileditems) | **Post** /wb/get/category | Товары категории
*WbApi* | [**GetCategoryPriceSegmentation**](docs/WbApi.md#getcategorypricesegmentation) | **Get** /wb/get/category/price_segmentation | Ценовая сегментация категории
*WbApi* | [**GetCategorySellers**](docs/WbApi.md#getcategorysellers) | **Get** /wb/get/category/sellers | Продавцы категории
*WbApi* | [**GetCategorySubcategories**](docs/WbApi.md#getcategorysubcategories) | **Get** /wb/get/category/subcategories | Подкатегории категории
*WbApi* | [**GetCategoryTrends**](docs/WbApi.md#getcategorytrends) | **Get** /wb/get/category/trends | Тренды категории
*WbApi* | [**GetItemBalanceByDay**](docs/WbApi.md#getitembalancebyday) | **Get** /wb/get/item/{sku}/balance_by_day | Продажи и остатки товарной позиции за сутки
*WbApi* | [**GetItemByCategory**](docs/WbApi.md#getitembycategory) | **Get** /wb/get/item/{sku}/by_category | История товарной позиции по категориям
*WbApi* | [**GetItemBySku**](docs/WbApi.md#getitembysku) | **Get** /wb/get/item/{sku} | Одна товарная позиция по SKU
*WbApi* | [**GetItemOrdersByRegion**](docs/WbApi.md#getitemordersbyregion) | **Get** /wb/get/item/{sku}/orders_by_region | История заказов и остатков товарной позиции по складам
*WbApi* | [**GetItemOrdersBySize**](docs/WbApi.md#getitemordersbysize) | **Get** /wb/get/item/{sku}/orders_by_size | История заказов и остатков товарной позиции по размерам
*WbApi* | [**GetItemSales**](docs/WbApi.md#getitemsales) | **Get** /wb/get/item/{sku}/sales | Продажи и остатки товарной позиции
*WbApi* | [**GetItemSimilar**](docs/WbApi.md#getitemsimilar) | **Get** /wb/get/item/{sku}/similar | Похожие товарные позиции
*WbApi* | [**GetItemsBatch**](docs/WbApi.md#getitemsbatch) | **Post** /wb/get/items/batch | Товарные позиции по SKU
*WbApi* | [**GetSellerBrands**](docs/WbApi.md#getsellerbrands) | **Get** /wb/get/seller/brands | Бренды продавца
*WbApi* | [**GetSellerByDate**](docs/WbApi.md#getsellerbydate) | **Get** /wb/get/seller/by_date | Данные по дням по продавцу
*WbApi* | [**GetSellerCategories**](docs/WbApi.md#getsellercategories) | **Get** /wb/get/seller/categories | Категории продавца
*WbApi* | [**GetSellerDetailedItems**](docs/WbApi.md#getsellerdetaileditems) | **Post** /wb/get/seller | Товары продавца
*WbApi* | [**GetSellerPriceSegmentation**](docs/WbApi.md#getsellerpricesegmentation) | **Get** /wb/get/seller/price_segmentation | Ценовая сегментация продавца
*WbApi* | [**GetSimilarBrands**](docs/WbApi.md#getsimilarbrands) | **Get** /wb/get/similar/brands | Бренды похожего товара
*WbApi* | [**GetSimilarByDate**](docs/WbApi.md#getsimilarbydate) | **Get** /wb/get/similar/by_date | По дням
*WbApi* | [**GetSimilarCategories**](docs/WbApi.md#getsimilarcategories) | **Get** /wb/get/similar/categories | Категории похожего товара
*WbApi* | [**GetSimilarDetailedItems**](docs/WbApi.md#getsimilardetaileditems) | **Post** /wb/get/similar | Товары по похожему товару
*WbApi* | [**GetSimilarSellers**](docs/WbApi.md#getsimilarsellers) | **Get** /wb/get/similar/sellers | Продавцы похожего товара
*WbApi* | [**PostWbGetSearchCategories**](docs/WbApi.md#postwbgetsearchcategories) | **Post** /wb/get/search/categories | 


## Documentation For Models

 - [BrandByDate](docs/BrandByDate.md)
 - [BrandCategory](docs/BrandCategory.md)
 - [BrandPriceSegmentation](docs/BrandPriceSegmentation.md)
 - [BrandSeller](docs/BrandSeller.md)
 - [CategoryBrand](docs/CategoryBrand.md)
 - [CategoryByDate](docs/CategoryByDate.md)
 - [CategoryName](docs/CategoryName.md)
 - [CategoryPath](docs/CategoryPath.md)
 - [CategoryPriceSegmentation](docs/CategoryPriceSegmentation.md)
 - [CategorySeller](docs/CategorySeller.md)
 - [CategorySubcategory](docs/CategorySubcategory.md)
 - [CategoryTrend](docs/CategoryTrend.md)
 - [CategoryUrl](docs/CategoryUrl.md)
 - [DetailedItem](docs/DetailedItem.md)
 - [FilterModelItem](docs/FilterModelItem.md)
 - [GetItemsBatchRequestBody](docs/GetItemsBatchRequestBody.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [Item](docs/Item.md)
 - [ItemBalanceByDay](docs/ItemBalanceByDay.md)
 - [ItemByCategory](docs/ItemByCategory.md)
 - [ItemItem](docs/ItemItem.md)
 - [ItemPhotos](docs/ItemPhotos.md)
 - [ItemSale](docs/ItemSale.md)
 - [ItemSimilar](docs/ItemSimilar.md)
 - [PostRequestBody](docs/PostRequestBody.md)
 - [SearchCategoriesItem](docs/SearchCategoriesItem.md)
 - [SearchCategory](docs/SearchCategory.md)
 - [SellerBrand](docs/SellerBrand.md)
 - [SellerByDate](docs/SellerByDate.md)
 - [SellerCategory](docs/SellerCategory.md)
 - [SellerPriceSegmentation](docs/SellerPriceSegmentation.md)
 - [SimilarBrand](docs/SimilarBrand.md)
 - [SimilarByDate](docs/SimilarByDate.md)
 - [SimilarCategory](docs/SimilarCategory.md)
 - [SimilarSeller](docs/SimilarSeller.md)
 - [SortModelItem](docs/SortModelItem.md)


## Documentation For Authorization



### Header-token

- **Type**: API key
- **API key parameter name**: X-Mpstats-TOKEN
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Mpstats-TOKEN and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



