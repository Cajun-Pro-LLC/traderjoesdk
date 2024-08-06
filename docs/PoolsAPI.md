# traderjoesdk\PoolsAPI

All URIs are relative to *https://api.traderjoexyz.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoricalLiquidityProviders**](PoolsAPI.md#GetHistoricalLiquidityProviders) | **Get** /v1/pools/historical-liquidity-providers/{chain}/{address} | Get Historical Liquidity Providers
[**GetPool**](PoolsAPI.md#GetPool) | **Get** /v1/pools/{chain}/{address} | Get Pool
[**ListPools**](PoolsAPI.md#ListPools) | **Get** /v1/pools/{chain} | List Pools



## GetHistoricalLiquidityProviders

> []HistoricalLpFeesEarned GetHistoricalLiquidityProviders(ctx, address, chain).StartTime(startTime).PageSize(pageSize).PageNum(pageNum).EndTime(endTime).OrderBy(orderBy).Execute()

Get Historical Liquidity Providers

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cajun-pro-llc/traderjoesdk"
)

func main() {
	address := "address_example" // string | 
	chain := openapiclient.Chain("avalanche") // Chain | 
	startTime := int32(56) // int32 | 
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	pageNum := int32(56) // int32 |  (optional) (default to 1)
	endTime := int32(56) // int32 |  (optional)
	orderBy := openapiclient.LeaderboardSortType{LeaderboardFeesSortType: openapiclient.LeaderboardFeesSortType("fees")} // LeaderboardSortType |  (optional) (default to "fees")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetHistoricalLiquidityProviders(context.Background(), address, chain).StartTime(startTime).PageSize(pageSize).PageNum(pageNum).EndTime(endTime).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetHistoricalLiquidityProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricalLiquidityProviders`: []HistoricalLpFeesEarned
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetHistoricalLiquidityProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 
**chain** | [**Chain**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricalLiquidityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **int32** |  | 
 **pageSize** | **int32** |  | [default to 20]
 **pageNum** | **int32** |  | [default to 1]
 **endTime** | **int32** |  | 
 **orderBy** | [**LeaderboardSortType**](LeaderboardSortType.md) |  | [default to &quot;fees&quot;]

### Return type

[**[]HistoricalLpFeesEarned**](HistoricalLpFeesEarned.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPool

> GetPair GetPool(ctx, chain, address).FilterBy(filterBy).Execute()

Get Pool

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cajun-pro-llc/traderjoesdk"
)

func main() {
	chain := openapiclient.Chain("avalanche") // Chain | 
	address := "address_example" // string | 
	filterBy := openapiclient.FilterType("1h") // FilterType |  (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetPool(context.Background(), chain, address).FilterBy(filterBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPool`: GetPair
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**Chain**](.md) |  | 
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filterBy** | [**FilterType**](FilterType.md) |  | [default to &quot;1d&quot;]

### Return type

[**GetPair**](GetPair.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPools

> []GetPair ListPools(ctx, chain).PageSize(pageSize).PageNum(pageNum).OrderBy(orderBy).FilterBy(filterBy).Status(status).Version(version).ExcludeLowVolumePools(excludeLowVolumePools).Execute()

List Pools

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cajun-pro-llc/traderjoesdk"
)

func main() {
	chain := openapiclient.Chain("avalanche") // Chain | 
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	pageNum := int32(56) // int32 |  (optional) (default to 1)
	orderBy := openapiclient.PoolSortType{PoolSortParam: openapiclient.PoolSortParam("liquidity")} // PoolSortType |  (optional) (default to volume)
	filterBy := openapiclient.FilterType("1h") // FilterType |  (optional) (default to "1d")
	status := openapiclient.QueryStatus{AllParam: openapiclient.AllParam("all")} // QueryStatus |  (optional) (default to all)
	version := openapiclient.PairVersionParam{AllParam: openapiclient.AllParam("all")} // PairVersionParam |  (optional) (default to "all")
	excludeLowVolumePools := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.ListPools(context.Background(), chain).PageSize(pageSize).PageNum(pageNum).OrderBy(orderBy).FilterBy(filterBy).Status(status).Version(version).ExcludeLowVolumePools(excludeLowVolumePools).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.ListPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPools`: []GetPair
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.ListPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**Chain**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | [default to 20]
 **pageNum** | **int32** |  | [default to 1]
 **orderBy** | [**PoolSortType**](PoolSortType.md) |  | [default to volume]
 **filterBy** | [**FilterType**](FilterType.md) |  | [default to &quot;1d&quot;]
 **status** | [**QueryStatus**](QueryStatus.md) |  | [default to all]
 **version** | [**PairVersionParam**](PairVersionParam.md) |  | [default to &quot;all&quot;]
 **excludeLowVolumePools** | **bool** |  | [default to true]

### Return type

[**[]GetPair**](GetPair.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

