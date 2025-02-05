/*
Trader Joe Dex API

<p>Discover DeFi with Trader Joe, a leading decentralized exchange. Trade a wide variety of tokens, earn rewards, and engage in secure, peer-to-peer transactions. Trader Joe makes DeFi easy and accessible.

API version: 1.0.0
Contact: public-api@traderjoexyz.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package traderjoesdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type PoolsAPI interface {

	/*
		GetHistoricalLiquidityProviders Get Historical Liquidity Providers

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param address
		@param chain
		@return ApiGetHistoricalLiquidityProvidersRequest
	*/
	GetHistoricalLiquidityProviders(ctx context.Context, address string, chain Chain) ApiGetHistoricalLiquidityProvidersRequest

	// GetHistoricalLiquidityProvidersExecute executes the request
	//  @return []HistoricalLpFeesEarned
	GetHistoricalLiquidityProvidersExecute(r ApiGetHistoricalLiquidityProvidersRequest) ([]HistoricalLpFeesEarned, *http.Response, error)

	/*
		GetPool Get Pool

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain
		@param address
		@return ApiGetPoolRequest
	*/
	GetPool(ctx context.Context, chain Chain, address string) ApiGetPoolRequest

	// GetPoolExecute executes the request
	//  @return GetPair
	GetPoolExecute(r ApiGetPoolRequest) (*GetPair, *http.Response, error)

	/*
		ListPools List Pools

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain
		@return ApiListPoolsRequest
	*/
	ListPools(ctx context.Context, chain Chain) ApiListPoolsRequest

	// ListPoolsExecute executes the request
	//  @return []GetPair
	ListPoolsExecute(r ApiListPoolsRequest) ([]GetPair, *http.Response, error)
}

// PoolsAPIService PoolsAPI service
type PoolsAPIService service

type ApiGetHistoricalLiquidityProvidersRequest struct {
	ctx        context.Context
	ApiService PoolsAPI
	address    string
	chain      Chain
	startTime  *int32
	pageSize   *int32
	pageNum    *int32
	endTime    *int32
	orderBy    *LeaderboardSortType
}

func (r ApiGetHistoricalLiquidityProvidersRequest) StartTime(startTime int32) ApiGetHistoricalLiquidityProvidersRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetHistoricalLiquidityProvidersRequest) PageSize(pageSize int32) ApiGetHistoricalLiquidityProvidersRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetHistoricalLiquidityProvidersRequest) PageNum(pageNum int32) ApiGetHistoricalLiquidityProvidersRequest {
	r.pageNum = &pageNum
	return r
}

func (r ApiGetHistoricalLiquidityProvidersRequest) EndTime(endTime int32) ApiGetHistoricalLiquidityProvidersRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetHistoricalLiquidityProvidersRequest) OrderBy(orderBy LeaderboardSortType) ApiGetHistoricalLiquidityProvidersRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiGetHistoricalLiquidityProvidersRequest) Execute() ([]HistoricalLpFeesEarned, *http.Response, error) {
	return r.ApiService.GetHistoricalLiquidityProvidersExecute(r)
}

/*
GetHistoricalLiquidityProviders Get Historical Liquidity Providers

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param address
	@param chain
	@return ApiGetHistoricalLiquidityProvidersRequest
*/
func (a *PoolsAPIService) GetHistoricalLiquidityProviders(ctx context.Context, address string, chain Chain) ApiGetHistoricalLiquidityProvidersRequest {
	return ApiGetHistoricalLiquidityProvidersRequest{
		ApiService: a,
		ctx:        ctx,
		address:    address,
		chain:      chain,
	}
}

// Execute executes the request
//
//	@return []HistoricalLpFeesEarned
func (a *PoolsAPIService) GetHistoricalLiquidityProvidersExecute(r ApiGetHistoricalLiquidityProvidersRequest) ([]HistoricalLpFeesEarned, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []HistoricalLpFeesEarned
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolsAPIService.GetHistoricalLiquidityProviders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/pools/historical-liquidity-providers/{chain}/{address}"
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", url.PathEscape(parameterValueToString(r.address, "address")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.startTime == nil {
		return localVarReturnValue, nil, reportError("startTime is required and must be specified")
	}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	} else {
		var defaultValue int32 = 20
		r.pageSize = &defaultValue
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int32 = 1
		r.pageNum = &defaultValue
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "")
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "")
	} else {
		var defaultValue LeaderboardSortType = "fees"
		r.orderBy = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-traderjoe-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPoolRequest struct {
	ctx        context.Context
	ApiService PoolsAPI
	chain      Chain
	address    string
	filterBy   *FilterType
}

func (r ApiGetPoolRequest) FilterBy(filterBy FilterType) ApiGetPoolRequest {
	r.filterBy = &filterBy
	return r
}

func (r ApiGetPoolRequest) Execute() (*GetPair, *http.Response, error) {
	return r.ApiService.GetPoolExecute(r)
}

/*
GetPool Get Pool

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain
	@param address
	@return ApiGetPoolRequest
*/
func (a *PoolsAPIService) GetPool(ctx context.Context, chain Chain, address string) ApiGetPoolRequest {
	return ApiGetPoolRequest{
		ApiService: a,
		ctx:        ctx,
		chain:      chain,
		address:    address,
	}
}

// Execute executes the request
//
//	@return GetPair
func (a *PoolsAPIService) GetPoolExecute(r ApiGetPoolRequest) (*GetPair, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetPair
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolsAPIService.GetPool")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/pools/{chain}/{address}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", url.PathEscape(parameterValueToString(r.address, "address")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterBy", r.filterBy, "")
	} else {
		var defaultValue FilterType = "1d"
		r.filterBy = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-traderjoe-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListPoolsRequest struct {
	ctx                   context.Context
	ApiService            PoolsAPI
	chain                 Chain
	pageSize              *int32
	pageNum               *int32
	orderBy               *PoolSortType
	filterBy              *FilterType
	status                *QueryStatus
	version               *PairVersionParam
	excludeLowVolumePools *bool
}

func (r ApiListPoolsRequest) PageSize(pageSize int32) ApiListPoolsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiListPoolsRequest) PageNum(pageNum int32) ApiListPoolsRequest {
	r.pageNum = &pageNum
	return r
}

func (r ApiListPoolsRequest) OrderBy(orderBy PoolSortType) ApiListPoolsRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiListPoolsRequest) FilterBy(filterBy FilterType) ApiListPoolsRequest {
	r.filterBy = &filterBy
	return r
}

func (r ApiListPoolsRequest) Status(status QueryStatus) ApiListPoolsRequest {
	r.status = &status
	return r
}

func (r ApiListPoolsRequest) Version(version PairVersionParam) ApiListPoolsRequest {
	r.version = &version
	return r
}

func (r ApiListPoolsRequest) ExcludeLowVolumePools(excludeLowVolumePools bool) ApiListPoolsRequest {
	r.excludeLowVolumePools = &excludeLowVolumePools
	return r
}

func (r ApiListPoolsRequest) Execute() ([]GetPair, *http.Response, error) {
	return r.ApiService.ListPoolsExecute(r)
}

/*
ListPools List Pools

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain
	@return ApiListPoolsRequest
*/
func (a *PoolsAPIService) ListPools(ctx context.Context, chain Chain) ApiListPoolsRequest {
	return ApiListPoolsRequest{
		ApiService: a,
		ctx:        ctx,
		chain:      chain,
	}
}

// Execute executes the request
//
//	@return []GetPair
func (a *PoolsAPIService) ListPoolsExecute(r ApiListPoolsRequest) ([]GetPair, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []GetPair
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolsAPIService.ListPools")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/pools/{chain}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	} else {
		var defaultValue int32 = 20
		r.pageSize = &defaultValue
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int32 = 1
		r.pageNum = &defaultValue
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "")
	} else {
		var defaultValue PoolSortType = "volume"
		r.orderBy = &defaultValue
	}
	if r.filterBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterBy", r.filterBy, "")
	} else {
		var defaultValue FilterType = "1d"
		r.filterBy = &defaultValue
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	} else {
		var defaultValue QueryStatus = "all"
		r.status = &defaultValue
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	} else {
		var defaultValue PairVersionParam = "all"
		r.version = &defaultValue
	}
	if r.excludeLowVolumePools != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeLowVolumePools", r.excludeLowVolumePools, "")
	} else {
		var defaultValue bool = true
		r.excludeLowVolumePools = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-traderjoe-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
