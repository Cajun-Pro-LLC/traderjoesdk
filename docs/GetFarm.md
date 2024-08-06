# GetFarm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FarmId** | **string** |  | 
**VaultId** | **string** |  | 
**Liquidity** | **string** |  | 
**LiquidityRaw** | **int32** |  | 
**LiquidityUsd** | **float32** |  | 
**AptDecimals** | **int32** |  | 
**Apr1d** | **float32** |  | 
**RewardsPerSec** | **float32** |  | 
**Reward** | [**GetReward**](GetReward.md) |  | 
**Rewarder** | Pointer to [**GetRewarder**](GetRewarder.md) |  | [optional] 

## Methods

### NewGetFarm

`func NewGetFarm(farmId string, vaultId string, liquidity string, liquidityRaw int32, liquidityUsd float32, aptDecimals int32, apr1d float32, rewardsPerSec float32, reward GetReward, ) *GetFarm`

NewGetFarm instantiates a new GetFarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFarmWithDefaults

`func NewGetFarmWithDefaults() *GetFarm`

NewGetFarmWithDefaults instantiates a new GetFarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFarmId

`func (o *GetFarm) GetFarmId() string`

GetFarmId returns the FarmId field if non-nil, zero value otherwise.

### GetFarmIdOk

`func (o *GetFarm) GetFarmIdOk() (*string, bool)`

GetFarmIdOk returns a tuple with the FarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarmId

`func (o *GetFarm) SetFarmId(v string)`

SetFarmId sets FarmId field to given value.


### GetVaultId

`func (o *GetFarm) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *GetFarm) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *GetFarm) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.


### GetLiquidity

`func (o *GetFarm) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *GetFarm) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *GetFarm) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.


### GetLiquidityRaw

`func (o *GetFarm) GetLiquidityRaw() int32`

GetLiquidityRaw returns the LiquidityRaw field if non-nil, zero value otherwise.

### GetLiquidityRawOk

`func (o *GetFarm) GetLiquidityRawOk() (*int32, bool)`

GetLiquidityRawOk returns a tuple with the LiquidityRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidityRaw

`func (o *GetFarm) SetLiquidityRaw(v int32)`

SetLiquidityRaw sets LiquidityRaw field to given value.


### GetLiquidityUsd

`func (o *GetFarm) GetLiquidityUsd() float32`

GetLiquidityUsd returns the LiquidityUsd field if non-nil, zero value otherwise.

### GetLiquidityUsdOk

`func (o *GetFarm) GetLiquidityUsdOk() (*float32, bool)`

GetLiquidityUsdOk returns a tuple with the LiquidityUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidityUsd

`func (o *GetFarm) SetLiquidityUsd(v float32)`

SetLiquidityUsd sets LiquidityUsd field to given value.


### GetAptDecimals

`func (o *GetFarm) GetAptDecimals() int32`

GetAptDecimals returns the AptDecimals field if non-nil, zero value otherwise.

### GetAptDecimalsOk

`func (o *GetFarm) GetAptDecimalsOk() (*int32, bool)`

GetAptDecimalsOk returns a tuple with the AptDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAptDecimals

`func (o *GetFarm) SetAptDecimals(v int32)`

SetAptDecimals sets AptDecimals field to given value.


### GetApr1d

`func (o *GetFarm) GetApr1d() float32`

GetApr1d returns the Apr1d field if non-nil, zero value otherwise.

### GetApr1dOk

`func (o *GetFarm) GetApr1dOk() (*float32, bool)`

GetApr1dOk returns a tuple with the Apr1d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApr1d

`func (o *GetFarm) SetApr1d(v float32)`

SetApr1d sets Apr1d field to given value.


### GetRewardsPerSec

`func (o *GetFarm) GetRewardsPerSec() float32`

GetRewardsPerSec returns the RewardsPerSec field if non-nil, zero value otherwise.

### GetRewardsPerSecOk

`func (o *GetFarm) GetRewardsPerSecOk() (*float32, bool)`

GetRewardsPerSecOk returns a tuple with the RewardsPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsPerSec

`func (o *GetFarm) SetRewardsPerSec(v float32)`

SetRewardsPerSec sets RewardsPerSec field to given value.


### GetReward

`func (o *GetFarm) GetReward() GetReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *GetFarm) GetRewardOk() (*GetReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *GetFarm) SetReward(v GetReward)`

SetReward sets Reward field to given value.


### GetRewarder

`func (o *GetFarm) GetRewarder() GetRewarder`

GetRewarder returns the Rewarder field if non-nil, zero value otherwise.

### GetRewarderOk

`func (o *GetFarm) GetRewarderOk() (*GetRewarder, bool)`

GetRewarderOk returns a tuple with the Rewarder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewarder

`func (o *GetFarm) SetRewarder(v GetRewarder)`

SetRewarder sets Rewarder field to given value.

### HasRewarder

`func (o *GetFarm) HasRewarder() bool`

HasRewarder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


