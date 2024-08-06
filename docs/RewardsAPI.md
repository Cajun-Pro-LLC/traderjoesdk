# traderjoesdk\RewardsAPI

All URIs are relative to *https://api.traderjoexyz.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProof**](RewardsAPI.md#GetProof) | **Get** /v1/rewards/{chain}/{user_address} | Get Proof
[**GetUserClaimableRewards**](RewardsAPI.md#GetUserClaimableRewards) | **Get** /v1/rewards/claimable/{chain}/{user_address} | Get User Claimable Rewards
[**GetUserProofs**](RewardsAPI.md#GetUserProofs) | **Post** /v1/rewards/batch-proof/{chain}/{user_address} | Get User Proofs
[**GetUserRewardHistory**](RewardsAPI.md#GetUserRewardHistory) | **Get** /v1/rewards/history/{chain}/{user_address} | Get User Reward History



## GetProof

> []string GetProof(ctx, chain, userAddress).Market(market).Epoch(epoch).Token(token).Execute()

Get Proof

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
	market := "market_example" // string | 
	epoch := int32(56) // int32 | 
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RewardsAPI.GetProof(context.Background(), chain, userAddress).Market(market).Epoch(epoch).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardsAPI.GetProof``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProof`: []string
	fmt.Fprintf(os.Stdout, "Response from `RewardsAPI.GetProof`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**Chain**](.md) |  | 
**userAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProofRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **market** | **string** |  | 
 **epoch** | **int32** |  | 
 **token** | **string** |  | 

### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserClaimableRewards

> []ClaimableReward GetUserClaimableRewards(ctx, chain, userAddress).Market(market).Execute()

Get User Claimable Rewards

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
	market := "market_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RewardsAPI.GetUserClaimableRewards(context.Background(), chain, userAddress).Market(market).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardsAPI.GetUserClaimableRewards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserClaimableRewards`: []ClaimableReward
	fmt.Fprintf(os.Stdout, "Response from `RewardsAPI.GetUserClaimableRewards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**Chain**](.md) |  | 
**userAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserClaimableRewardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **market** | **string** |  | 

### Return type

[**[]ClaimableReward**](ClaimableReward.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserProofs

> [][]string GetUserProofs(ctx, chain, userAddress).GetProofBatch(getProofBatch).Execute()

Get User Proofs

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
	getProofBatch := *openapiclient.NewGetProofBatch([]openapiclient.GetProof{*openapiclient.NewGetProof("Market_example", int32(123), "Token_example")}) // GetProofBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RewardsAPI.GetUserProofs(context.Background(), chain, userAddress).GetProofBatch(getProofBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardsAPI.GetUserProofs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserProofs`: [][]string
	fmt.Fprintf(os.Stdout, "Response from `RewardsAPI.GetUserProofs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**Chain**](.md) |  | 
**userAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserProofsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **getProofBatch** | [**GetProofBatch**](GetProofBatch.md) |  | 

### Return type

[**[][]string**](array.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserRewardHistory

> []GetUserRewardHistoryResponse GetUserRewardHistory(ctx, chain, userAddress).Market(market).Execute()

Get User Reward History

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
	market := "market_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RewardsAPI.GetUserRewardHistory(context.Background(), chain, userAddress).Market(market).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardsAPI.GetUserRewardHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserRewardHistory`: []GetUserRewardHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RewardsAPI.GetUserRewardHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**Chain**](.md) |  | 
**userAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRewardHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **market** | **string** |  | 

### Return type

[**[]GetUserRewardHistoryResponse**](GetUserRewardHistoryResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

