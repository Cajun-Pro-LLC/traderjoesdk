# traderjoesdk\UserAPI

All URIs are relative to *https://api.traderjoexyz.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserFarmPosition**](UserAPI.md#GetUserFarmPosition) | **Get** /v1/user/{chain}/{user_address}/farms/{vault_id} | Get User Farm Position
[**GetUserFarmPositions**](UserAPI.md#GetUserFarmPositions) | **Get** /v1/user/{chain}/{user_address}/farms | Get User Farm Positions
[**GetUserPoolIds**](UserAPI.md#GetUserPoolIds) | **Get** /v1/user/pool-ids/{user_address}/{chain} | Get User Pool Ids
[**UserCurrentBinIds**](UserAPI.md#UserCurrentBinIds) | **Get** /v1/user/bin-ids/{user_address}/{chain}/{pool_address} | User Current Bin Ids
[**UserEarnedFeesPerBin**](UserAPI.md#UserEarnedFeesPerBin) | **Get** /v1/user/fees-earned/{chain}/{user_address}/{pool_address} | User Earned Fees Per Bin
[**UserHistoricalPosition**](UserAPI.md#UserHistoricalPosition) | **Get** /v1/user/{chain}/history/{user_address}/{pool_address} | User Historical Position



## GetUserFarmPosition

> GetUserFarmPosition GetUserFarmPosition(ctx, chain, userAddress, vaultId).Execute()

Get User Farm Position

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
	userAddress := "userAddress_example" // string | 
	vaultId := "vaultId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserFarmPosition(context.Background(), chain, userAddress, vaultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserFarmPosition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserFarmPosition`: GetUserFarmPosition
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserFarmPosition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**Chain**](.md) |  | 
**userAddress** | **string** |  | 
**vaultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFarmPositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetUserFarmPosition**](GetUserFarmPosition.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserFarmPositions

> []GetUserFarmPosition GetUserFarmPositions(ctx, chain, userAddress).Execute()

Get User Farm Positions

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
	userAddress := "userAddress_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserFarmPositions(context.Background(), chain, userAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserFarmPositions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserFarmPositions`: []GetUserFarmPosition
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserFarmPositions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**Chain**](.md) |  | 
**userAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFarmPositionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetUserFarmPosition**](GetUserFarmPosition.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPoolIds

> []UserPoolPosition GetUserPoolIds(ctx, userAddress, chain).PageSize(pageSize).PageNum(pageNum).Execute()

Get User Pool Ids

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
	userAddress := "userAddress_example" // string | 
	chain := openapiclient.Chain("avalanche") // Chain | 
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	pageNum := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserPoolIds(context.Background(), userAddress, chain).PageSize(pageSize).PageNum(pageNum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserPoolIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPoolIds`: []UserPoolPosition
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserPoolIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userAddress** | **string** |  | 
**chain** | [**Chain**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPoolIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** |  | [default to 20]
 **pageNum** | **int32** |  | [default to 1]

### Return type

[**[]UserPoolPosition**](UserPoolPosition.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCurrentBinIds

> []int32 UserCurrentBinIds(ctx, userAddress, chain, poolAddress).Execute()

User Current Bin Ids

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
	userAddress := "userAddress_example" // string | 
	chain := openapiclient.Chain("avalanche") // Chain | 
	poolAddress := "poolAddress_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserCurrentBinIds(context.Background(), userAddress, chain, poolAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserCurrentBinIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserCurrentBinIds`: []int32
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserCurrentBinIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userAddress** | **string** |  | 
**chain** | [**Chain**](.md) |  | 
**poolAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserCurrentBinIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**[]int32**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserEarnedFeesPerBin

> []UserFeesEarnedPerBin UserEarnedFeesPerBin(ctx, chain, userAddress, poolAddress).Execute()

User Earned Fees Per Bin

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
	chain := "chain_example" // ChainParam | 
	userAddress := "userAddress_example" // string | 
	poolAddress := "poolAddress_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserEarnedFeesPerBin(context.Background(), chain, userAddress, poolAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserEarnedFeesPerBin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserEarnedFeesPerBin`: []UserFeesEarnedPerBin
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserEarnedFeesPerBin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **ChainParam** |  | 
**userAddress** | **string** |  | 
**poolAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserEarnedFeesPerBinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]UserFeesEarnedPerBin**](UserFeesEarnedPerBin.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserHistoricalPosition

> []UserBinHistory UserHistoricalPosition(ctx, chain, userAddress, poolAddress).StartTime(startTime).PageSize(pageSize).PageNum(pageNum).EndTime(endTime).Execute()

User Historical Position

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
	userAddress := "userAddress_example" // string | 
	poolAddress := "poolAddress_example" // string | 
	startTime := int32(56) // int32 | 
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	pageNum := int32(56) // int32 |  (optional) (default to 1)
	endTime := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserHistoricalPosition(context.Background(), chain, userAddress, poolAddress).StartTime(startTime).PageSize(pageSize).PageNum(pageNum).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserHistoricalPosition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserHistoricalPosition`: []UserBinHistory
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserHistoricalPosition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**Chain**](.md) |  | 
**userAddress** | **string** |  | 
**poolAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserHistoricalPositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **startTime** | **int32** |  | 
 **pageSize** | **int32** |  | [default to 20]
 **pageNum** | **int32** |  | [default to 1]
 **endTime** | **int32** |  | 

### Return type

[**[]UserBinHistory**](UserBinHistory.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

