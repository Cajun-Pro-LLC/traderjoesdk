# traderjoesdk\DexAnalyticsAPI

All URIs are relative to *https://api.traderjoexyz.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DailyDexAnalytics**](DexAnalyticsAPI.md#DailyDexAnalytics) | **Get** /v1/dex/analytics/{chain} | Daily Dex Analytics



## DailyDexAnalytics

> []DexAnalytics DailyDexAnalytics(ctx, chain).StartTime(startTime).EndTime(endTime).Execute()

Daily Dex Analytics

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
	startTime := int32(56) // int32 | 
	endTime := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DexAnalyticsAPI.DailyDexAnalytics(context.Background(), chain).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DexAnalyticsAPI.DailyDexAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DailyDexAnalytics`: []DexAnalytics
	fmt.Fprintf(os.Stdout, "Response from `DexAnalyticsAPI.DailyDexAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **ChainParam** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDailyDexAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int32** |  | 
 **endTime** | **int32** |  | 

### Return type

[**[]DexAnalytics**](DexAnalytics.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

