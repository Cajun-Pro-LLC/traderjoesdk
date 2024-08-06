# traderjoesdk\VaultsAPI

All URIs are relative to *https://api.traderjoexyz.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVault**](VaultsAPI.md#GetVault) | **Get** /v1/vaults/{chain}/{vault_address} | Get Vault
[**GetVaultRecentActivity**](VaultsAPI.md#GetVaultRecentActivity) | **Get** /v1/vaults/{chain}/{vault_address}/recent-activity | Get Vault Recent Activity
[**GetVaultSharePrice**](VaultsAPI.md#GetVaultSharePrice) | **Get** /v1/vaults/{chain}/{vault_address}/share-price | Get Vault Share Price
[**GetVaultTvlHistory**](VaultsAPI.md#GetVaultTvlHistory) | **Get** /v1/vaults/{chain}/{vault_address}/tvl-history | Get Vault Tvl History
[**GetVaultWithdrawalsByUser**](VaultsAPI.md#GetVaultWithdrawalsByUser) | **Get** /v1/vaults/{chain}/withdrawals/{user_address} | Get Vault Withdrawals By User
[**GetVaultWithdrawalsByUserAndVault**](VaultsAPI.md#GetVaultWithdrawalsByUserAndVault) | **Get** /v1/vaults/{chain}/{vault_address}/withdrawals/{user_address} | Get Vault Withdrawals By User And Vault
[**ListVaultsByChain**](VaultsAPI.md#ListVaultsByChain) | **Get** /v1/vaults/{chain} | List Vaults By Chain
[**ListVaultsV1VaultsGet**](VaultsAPI.md#ListVaultsV1VaultsGet) | **Get** /v1/vaults | List Vaults



## GetVault

> GetVault GetVault(ctx, chain, vaultAddress).Execute()

Get Vault

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
	vaultAddress := "vaultAddress_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.GetVault(context.Background(), chain, vaultAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.GetVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVault`: GetVault
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.GetVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **ChainParam** |  | 
**vaultAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVault**](GetVault.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultRecentActivity

> []GetVaultActivity GetVaultRecentActivity(ctx, chain, vaultAddress).PageSize(pageSize).PageNum(pageNum).Execute()

Get Vault Recent Activity

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
	vaultAddress := "vaultAddress_example" // string | 
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	pageNum := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.GetVaultRecentActivity(context.Background(), chain, vaultAddress).PageSize(pageSize).PageNum(pageNum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.GetVaultRecentActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVaultRecentActivity`: []GetVaultActivity
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.GetVaultRecentActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **ChainParam** |  | 
**vaultAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultRecentActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** |  | [default to 20]
 **pageNum** | **int32** |  | [default to 1]

### Return type

[**[]GetVaultActivity**](GetVaultActivity.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultSharePrice

> []VaultSharePrice GetVaultSharePrice(ctx, chain, vaultAddress).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Execute()

Get Vault Share Price

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
	vaultAddress := "vaultAddress_example" // string | 
	fromTimestamp := int32(56) // int32 | 
	toTimestamp := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.GetVaultSharePrice(context.Background(), chain, vaultAddress).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.GetVaultSharePrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVaultSharePrice`: []VaultSharePrice
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.GetVaultSharePrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**Chain**](.md) |  | 
**vaultAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultSharePriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromTimestamp** | **int32** |  | 
 **toTimestamp** | **int32** |  | 

### Return type

[**[]VaultSharePrice**](VaultSharePrice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultTvlHistory

> []GetVaultTVL GetVaultTvlHistory(ctx, chain, vaultAddress).StartTime(startTime).EndTime(endTime).Execute()

Get Vault Tvl History

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
	vaultAddress := "vaultAddress_example" // string | 
	startTime := int32(56) // int32 | 
	endTime := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.GetVaultTvlHistory(context.Background(), chain, vaultAddress).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.GetVaultTvlHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVaultTvlHistory`: []GetVaultTVL
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.GetVaultTvlHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **ChainParam** |  | 
**vaultAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultTvlHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **int32** |  | 
 **endTime** | **int32** |  | 

### Return type

[**[]GetVaultTVL**](GetVaultTVL.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultWithdrawalsByUser

> []GetVaultWithdrawal GetVaultWithdrawalsByUser(ctx, chain, userAddress).PageSize(pageSize).PageNum(pageNum).Execute()

Get Vault Withdrawals By User

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
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	pageNum := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.GetVaultWithdrawalsByUser(context.Background(), chain, userAddress).PageSize(pageSize).PageNum(pageNum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.GetVaultWithdrawalsByUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVaultWithdrawalsByUser`: []GetVaultWithdrawal
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.GetVaultWithdrawalsByUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **ChainParam** |  | 
**userAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultWithdrawalsByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** |  | [default to 20]
 **pageNum** | **int32** |  | [default to 1]

### Return type

[**[]GetVaultWithdrawal**](GetVaultWithdrawal.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultWithdrawalsByUserAndVault

> []GetVaultWithdrawal GetVaultWithdrawalsByUserAndVault(ctx, chain, vaultAddress, userAddress).PageSize(pageSize).PageNum(pageNum).Execute()

Get Vault Withdrawals By User And Vault

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
	vaultAddress := "vaultAddress_example" // string | 
	userAddress := "userAddress_example" // string | 
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	pageNum := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.GetVaultWithdrawalsByUserAndVault(context.Background(), chain, vaultAddress, userAddress).PageSize(pageSize).PageNum(pageNum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.GetVaultWithdrawalsByUserAndVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVaultWithdrawalsByUserAndVault`: []GetVaultWithdrawal
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.GetVaultWithdrawalsByUserAndVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **ChainParam** |  | 
**vaultAddress** | **string** |  | 
**userAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultWithdrawalsByUserAndVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageSize** | **int32** |  | [default to 20]
 **pageNum** | **int32** |  | [default to 1]

### Return type

[**[]GetVaultWithdrawal**](GetVaultWithdrawal.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVaultsByChain

> []GetVault ListVaultsByChain(ctx, chain).PageSize(pageSize).PageNum(pageNum).UserAddress(userAddress).Execute()

List Vaults By Chain

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
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	pageNum := int32(56) // int32 |  (optional) (default to 1)
	userAddress := "userAddress_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.ListVaultsByChain(context.Background(), chain).PageSize(pageSize).PageNum(pageNum).UserAddress(userAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.ListVaultsByChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVaultsByChain`: []GetVault
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.ListVaultsByChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **ChainParam** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultsByChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | [default to 20]
 **pageNum** | **int32** |  | [default to 1]
 **userAddress** | **string** |  | 

### Return type

[**[]GetVault**](GetVault.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVaultsV1VaultsGet

> []GetVault ListVaultsV1VaultsGet(ctx).PageSize(pageSize).PageNum(pageNum).Execute()

List Vaults

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
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	pageNum := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.ListVaultsV1VaultsGet(context.Background()).PageSize(pageSize).PageNum(pageNum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.ListVaultsV1VaultsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVaultsV1VaultsGet`: []GetVault
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.ListVaultsV1VaultsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVaultsV1VaultsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** |  | [default to 20]
 **pageNum** | **int32** |  | [default to 1]

### Return type

[**[]GetVault**](GetVault.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

