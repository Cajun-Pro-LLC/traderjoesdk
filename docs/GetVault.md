# GetVault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Chain** | [**Chain**](Chain.md) |  | 
**ChainId** | **int32** |  | 
**Name** | **string** |  | 
**Pair** | [**GetVaultPair**](GetVaultPair.md) |  | 
**TokenX** | [**GetVaultToken**](GetVaultToken.md) |  | 
**TokenY** | [**GetVaultToken**](GetVaultToken.md) |  | 
**TokenX7DayPerformance** | **float32** |  | 
**TokenY7DayPerformance** | **float32** |  | 
**Hodl5050Performance** | **float32** |  | 
**TokenX30DayPerformance** | **float32** |  | 
**TokenY30DayPerformance** | **float32** |  | 
**Hodl30Day5050Performance** | **float32** |  | 
**Strategy** | Pointer to [**GetVaultStrategy**](GetVaultStrategy.md) |  | [optional] 
**AptPrice** | Pointer to **float32** |  | [optional] 
**Apt1dPriceChange** | **float32** |  | 
**TvlUsd** | **float32** |  | 
**FeesUsd** | **float32** |  | 
**Apr1d** | **float32** |  | 
**Farm** | Pointer to [**GetFarm**](GetFarm.md) |  | [optional] 

## Methods

### NewGetVault

`func NewGetVault(address string, chain Chain, chainId int32, name string, pair GetVaultPair, tokenX GetVaultToken, tokenY GetVaultToken, tokenX7DayPerformance float32, tokenY7DayPerformance float32, hodl5050Performance float32, tokenX30DayPerformance float32, tokenY30DayPerformance float32, hodl30Day5050Performance float32, apt1dPriceChange float32, tvlUsd float32, feesUsd float32, apr1d float32, ) *GetVault`

NewGetVault instantiates a new GetVault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVaultWithDefaults

`func NewGetVaultWithDefaults() *GetVault`

NewGetVaultWithDefaults instantiates a new GetVault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GetVault) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetVault) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetVault) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChain

`func (o *GetVault) GetChain() Chain`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *GetVault) GetChainOk() (*Chain, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *GetVault) SetChain(v Chain)`

SetChain sets Chain field to given value.


### GetChainId

`func (o *GetVault) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *GetVault) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *GetVault) SetChainId(v int32)`

SetChainId sets ChainId field to given value.


### GetName

`func (o *GetVault) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetVault) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetVault) SetName(v string)`

SetName sets Name field to given value.


### GetPair

`func (o *GetVault) GetPair() GetVaultPair`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *GetVault) GetPairOk() (*GetVaultPair, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *GetVault) SetPair(v GetVaultPair)`

SetPair sets Pair field to given value.


### GetTokenX

`func (o *GetVault) GetTokenX() GetVaultToken`

GetTokenX returns the TokenX field if non-nil, zero value otherwise.

### GetTokenXOk

`func (o *GetVault) GetTokenXOk() (*GetVaultToken, bool)`

GetTokenXOk returns a tuple with the TokenX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenX

`func (o *GetVault) SetTokenX(v GetVaultToken)`

SetTokenX sets TokenX field to given value.


### GetTokenY

`func (o *GetVault) GetTokenY() GetVaultToken`

GetTokenY returns the TokenY field if non-nil, zero value otherwise.

### GetTokenYOk

`func (o *GetVault) GetTokenYOk() (*GetVaultToken, bool)`

GetTokenYOk returns a tuple with the TokenY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenY

`func (o *GetVault) SetTokenY(v GetVaultToken)`

SetTokenY sets TokenY field to given value.


### GetTokenX7DayPerformance

`func (o *GetVault) GetTokenX7DayPerformance() float32`

GetTokenX7DayPerformance returns the TokenX7DayPerformance field if non-nil, zero value otherwise.

### GetTokenX7DayPerformanceOk

`func (o *GetVault) GetTokenX7DayPerformanceOk() (*float32, bool)`

GetTokenX7DayPerformanceOk returns a tuple with the TokenX7DayPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenX7DayPerformance

`func (o *GetVault) SetTokenX7DayPerformance(v float32)`

SetTokenX7DayPerformance sets TokenX7DayPerformance field to given value.


### GetTokenY7DayPerformance

`func (o *GetVault) GetTokenY7DayPerformance() float32`

GetTokenY7DayPerformance returns the TokenY7DayPerformance field if non-nil, zero value otherwise.

### GetTokenY7DayPerformanceOk

`func (o *GetVault) GetTokenY7DayPerformanceOk() (*float32, bool)`

GetTokenY7DayPerformanceOk returns a tuple with the TokenY7DayPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenY7DayPerformance

`func (o *GetVault) SetTokenY7DayPerformance(v float32)`

SetTokenY7DayPerformance sets TokenY7DayPerformance field to given value.


### GetHodl5050Performance

`func (o *GetVault) GetHodl5050Performance() float32`

GetHodl5050Performance returns the Hodl5050Performance field if non-nil, zero value otherwise.

### GetHodl5050PerformanceOk

`func (o *GetVault) GetHodl5050PerformanceOk() (*float32, bool)`

GetHodl5050PerformanceOk returns a tuple with the Hodl5050Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHodl5050Performance

`func (o *GetVault) SetHodl5050Performance(v float32)`

SetHodl5050Performance sets Hodl5050Performance field to given value.


### GetTokenX30DayPerformance

`func (o *GetVault) GetTokenX30DayPerformance() float32`

GetTokenX30DayPerformance returns the TokenX30DayPerformance field if non-nil, zero value otherwise.

### GetTokenX30DayPerformanceOk

`func (o *GetVault) GetTokenX30DayPerformanceOk() (*float32, bool)`

GetTokenX30DayPerformanceOk returns a tuple with the TokenX30DayPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenX30DayPerformance

`func (o *GetVault) SetTokenX30DayPerformance(v float32)`

SetTokenX30DayPerformance sets TokenX30DayPerformance field to given value.


### GetTokenY30DayPerformance

`func (o *GetVault) GetTokenY30DayPerformance() float32`

GetTokenY30DayPerformance returns the TokenY30DayPerformance field if non-nil, zero value otherwise.

### GetTokenY30DayPerformanceOk

`func (o *GetVault) GetTokenY30DayPerformanceOk() (*float32, bool)`

GetTokenY30DayPerformanceOk returns a tuple with the TokenY30DayPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenY30DayPerformance

`func (o *GetVault) SetTokenY30DayPerformance(v float32)`

SetTokenY30DayPerformance sets TokenY30DayPerformance field to given value.


### GetHodl30Day5050Performance

`func (o *GetVault) GetHodl30Day5050Performance() float32`

GetHodl30Day5050Performance returns the Hodl30Day5050Performance field if non-nil, zero value otherwise.

### GetHodl30Day5050PerformanceOk

`func (o *GetVault) GetHodl30Day5050PerformanceOk() (*float32, bool)`

GetHodl30Day5050PerformanceOk returns a tuple with the Hodl30Day5050Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHodl30Day5050Performance

`func (o *GetVault) SetHodl30Day5050Performance(v float32)`

SetHodl30Day5050Performance sets Hodl30Day5050Performance field to given value.


### GetStrategy

`func (o *GetVault) GetStrategy() GetVaultStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *GetVault) GetStrategyOk() (*GetVaultStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *GetVault) SetStrategy(v GetVaultStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *GetVault) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetAptPrice

`func (o *GetVault) GetAptPrice() float32`

GetAptPrice returns the AptPrice field if non-nil, zero value otherwise.

### GetAptPriceOk

`func (o *GetVault) GetAptPriceOk() (*float32, bool)`

GetAptPriceOk returns a tuple with the AptPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAptPrice

`func (o *GetVault) SetAptPrice(v float32)`

SetAptPrice sets AptPrice field to given value.

### HasAptPrice

`func (o *GetVault) HasAptPrice() bool`

HasAptPrice returns a boolean if a field has been set.

### GetApt1dPriceChange

`func (o *GetVault) GetApt1dPriceChange() float32`

GetApt1dPriceChange returns the Apt1dPriceChange field if non-nil, zero value otherwise.

### GetApt1dPriceChangeOk

`func (o *GetVault) GetApt1dPriceChangeOk() (*float32, bool)`

GetApt1dPriceChangeOk returns a tuple with the Apt1dPriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApt1dPriceChange

`func (o *GetVault) SetApt1dPriceChange(v float32)`

SetApt1dPriceChange sets Apt1dPriceChange field to given value.


### GetTvlUsd

`func (o *GetVault) GetTvlUsd() float32`

GetTvlUsd returns the TvlUsd field if non-nil, zero value otherwise.

### GetTvlUsdOk

`func (o *GetVault) GetTvlUsdOk() (*float32, bool)`

GetTvlUsdOk returns a tuple with the TvlUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvlUsd

`func (o *GetVault) SetTvlUsd(v float32)`

SetTvlUsd sets TvlUsd field to given value.


### GetFeesUsd

`func (o *GetVault) GetFeesUsd() float32`

GetFeesUsd returns the FeesUsd field if non-nil, zero value otherwise.

### GetFeesUsdOk

`func (o *GetVault) GetFeesUsdOk() (*float32, bool)`

GetFeesUsdOk returns a tuple with the FeesUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesUsd

`func (o *GetVault) SetFeesUsd(v float32)`

SetFeesUsd sets FeesUsd field to given value.


### GetApr1d

`func (o *GetVault) GetApr1d() float32`

GetApr1d returns the Apr1d field if non-nil, zero value otherwise.

### GetApr1dOk

`func (o *GetVault) GetApr1dOk() (*float32, bool)`

GetApr1dOk returns a tuple with the Apr1d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApr1d

`func (o *GetVault) SetApr1d(v float32)`

SetApr1d sets Apr1d field to given value.


### GetFarm

`func (o *GetVault) GetFarm() GetFarm`

GetFarm returns the Farm field if non-nil, zero value otherwise.

### GetFarmOk

`func (o *GetVault) GetFarmOk() (*GetFarm, bool)`

GetFarmOk returns a tuple with the Farm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarm

`func (o *GetVault) SetFarm(v GetFarm)`

SetFarm sets Farm field to given value.

### HasFarm

`func (o *GetVault) HasFarm() bool`

HasFarm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


