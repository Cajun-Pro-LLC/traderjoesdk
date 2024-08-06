# GetPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PairAddress** | **string** |  | 
**Chain** | [**Chain**](Chain.md) |  | 
**Name** | **string** |  | 
**Status** | [**PairStatus**](PairStatus.md) |  | 
**Version** | [**PairVersion**](PairVersion.md) |  | 
**TokenX** | [**TokenWrapper**](TokenWrapper.md) |  | 
**TokenY** | [**TokenWrapper**](TokenWrapper.md) |  | 
**ReserveX** | **float32** |  | 
**ReserveY** | **float32** |  | 
**LbBinStep** | **int32** |  | 
**LbBaseFeePct** | **float32** |  | 
**LbMaxFeePct** | **float32** |  | 
**ActiveBinId** | Pointer to **int32** |  | [optional] 
**LiquidityUsd** | **float32** |  | 
**LiquidityNative** | **string** |  | 
**LiquidityDepthMinus** | **float32** |  | 
**LiquidityDepthPlus** | **float32** |  | 
**LiquidityDepthTokenX** | **float32** |  | 
**LiquidityDepthTokenY** | **float32** |  | 
**VolumeUsd** | **float32** |  | 
**VolumeNative** | **string** |  | 
**FeesUsd** | **float32** |  | 
**FeesNative** | **string** |  | 
**ProtocolSharePct** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetPair

`func NewGetPair(pairAddress string, chain Chain, name string, status PairStatus, version PairVersion, tokenX TokenWrapper, tokenY TokenWrapper, reserveX float32, reserveY float32, lbBinStep int32, lbBaseFeePct float32, lbMaxFeePct float32, liquidityUsd float32, liquidityNative string, liquidityDepthMinus float32, liquidityDepthPlus float32, liquidityDepthTokenX float32, liquidityDepthTokenY float32, volumeUsd float32, volumeNative string, feesUsd float32, feesNative string, ) *GetPair`

NewGetPair instantiates a new GetPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPairWithDefaults

`func NewGetPairWithDefaults() *GetPair`

NewGetPairWithDefaults instantiates a new GetPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPairAddress

`func (o *GetPair) GetPairAddress() string`

GetPairAddress returns the PairAddress field if non-nil, zero value otherwise.

### GetPairAddressOk

`func (o *GetPair) GetPairAddressOk() (*string, bool)`

GetPairAddressOk returns a tuple with the PairAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairAddress

`func (o *GetPair) SetPairAddress(v string)`

SetPairAddress sets PairAddress field to given value.


### GetChain

`func (o *GetPair) GetChain() Chain`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *GetPair) GetChainOk() (*Chain, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *GetPair) SetChain(v Chain)`

SetChain sets Chain field to given value.


### GetName

`func (o *GetPair) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetPair) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetPair) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *GetPair) GetStatus() PairStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPair) GetStatusOk() (*PairStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPair) SetStatus(v PairStatus)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *GetPair) GetVersion() PairVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetPair) GetVersionOk() (*PairVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetPair) SetVersion(v PairVersion)`

SetVersion sets Version field to given value.


### GetTokenX

`func (o *GetPair) GetTokenX() TokenWrapper`

GetTokenX returns the TokenX field if non-nil, zero value otherwise.

### GetTokenXOk

`func (o *GetPair) GetTokenXOk() (*TokenWrapper, bool)`

GetTokenXOk returns a tuple with the TokenX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenX

`func (o *GetPair) SetTokenX(v TokenWrapper)`

SetTokenX sets TokenX field to given value.


### GetTokenY

`func (o *GetPair) GetTokenY() TokenWrapper`

GetTokenY returns the TokenY field if non-nil, zero value otherwise.

### GetTokenYOk

`func (o *GetPair) GetTokenYOk() (*TokenWrapper, bool)`

GetTokenYOk returns a tuple with the TokenY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenY

`func (o *GetPair) SetTokenY(v TokenWrapper)`

SetTokenY sets TokenY field to given value.


### GetReserveX

`func (o *GetPair) GetReserveX() float32`

GetReserveX returns the ReserveX field if non-nil, zero value otherwise.

### GetReserveXOk

`func (o *GetPair) GetReserveXOk() (*float32, bool)`

GetReserveXOk returns a tuple with the ReserveX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveX

`func (o *GetPair) SetReserveX(v float32)`

SetReserveX sets ReserveX field to given value.


### GetReserveY

`func (o *GetPair) GetReserveY() float32`

GetReserveY returns the ReserveY field if non-nil, zero value otherwise.

### GetReserveYOk

`func (o *GetPair) GetReserveYOk() (*float32, bool)`

GetReserveYOk returns a tuple with the ReserveY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveY

`func (o *GetPair) SetReserveY(v float32)`

SetReserveY sets ReserveY field to given value.


### GetLbBinStep

`func (o *GetPair) GetLbBinStep() int32`

GetLbBinStep returns the LbBinStep field if non-nil, zero value otherwise.

### GetLbBinStepOk

`func (o *GetPair) GetLbBinStepOk() (*int32, bool)`

GetLbBinStepOk returns a tuple with the LbBinStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbBinStep

`func (o *GetPair) SetLbBinStep(v int32)`

SetLbBinStep sets LbBinStep field to given value.


### GetLbBaseFeePct

`func (o *GetPair) GetLbBaseFeePct() float32`

GetLbBaseFeePct returns the LbBaseFeePct field if non-nil, zero value otherwise.

### GetLbBaseFeePctOk

`func (o *GetPair) GetLbBaseFeePctOk() (*float32, bool)`

GetLbBaseFeePctOk returns a tuple with the LbBaseFeePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbBaseFeePct

`func (o *GetPair) SetLbBaseFeePct(v float32)`

SetLbBaseFeePct sets LbBaseFeePct field to given value.


### GetLbMaxFeePct

`func (o *GetPair) GetLbMaxFeePct() float32`

GetLbMaxFeePct returns the LbMaxFeePct field if non-nil, zero value otherwise.

### GetLbMaxFeePctOk

`func (o *GetPair) GetLbMaxFeePctOk() (*float32, bool)`

GetLbMaxFeePctOk returns a tuple with the LbMaxFeePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbMaxFeePct

`func (o *GetPair) SetLbMaxFeePct(v float32)`

SetLbMaxFeePct sets LbMaxFeePct field to given value.


### GetActiveBinId

`func (o *GetPair) GetActiveBinId() int32`

GetActiveBinId returns the ActiveBinId field if non-nil, zero value otherwise.

### GetActiveBinIdOk

`func (o *GetPair) GetActiveBinIdOk() (*int32, bool)`

GetActiveBinIdOk returns a tuple with the ActiveBinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBinId

`func (o *GetPair) SetActiveBinId(v int32)`

SetActiveBinId sets ActiveBinId field to given value.

### HasActiveBinId

`func (o *GetPair) HasActiveBinId() bool`

HasActiveBinId returns a boolean if a field has been set.

### GetLiquidityUsd

`func (o *GetPair) GetLiquidityUsd() float32`

GetLiquidityUsd returns the LiquidityUsd field if non-nil, zero value otherwise.

### GetLiquidityUsdOk

`func (o *GetPair) GetLiquidityUsdOk() (*float32, bool)`

GetLiquidityUsdOk returns a tuple with the LiquidityUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidityUsd

`func (o *GetPair) SetLiquidityUsd(v float32)`

SetLiquidityUsd sets LiquidityUsd field to given value.


### GetLiquidityNative

`func (o *GetPair) GetLiquidityNative() string`

GetLiquidityNative returns the LiquidityNative field if non-nil, zero value otherwise.

### GetLiquidityNativeOk

`func (o *GetPair) GetLiquidityNativeOk() (*string, bool)`

GetLiquidityNativeOk returns a tuple with the LiquidityNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidityNative

`func (o *GetPair) SetLiquidityNative(v string)`

SetLiquidityNative sets LiquidityNative field to given value.


### GetLiquidityDepthMinus

`func (o *GetPair) GetLiquidityDepthMinus() float32`

GetLiquidityDepthMinus returns the LiquidityDepthMinus field if non-nil, zero value otherwise.

### GetLiquidityDepthMinusOk

`func (o *GetPair) GetLiquidityDepthMinusOk() (*float32, bool)`

GetLiquidityDepthMinusOk returns a tuple with the LiquidityDepthMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidityDepthMinus

`func (o *GetPair) SetLiquidityDepthMinus(v float32)`

SetLiquidityDepthMinus sets LiquidityDepthMinus field to given value.


### GetLiquidityDepthPlus

`func (o *GetPair) GetLiquidityDepthPlus() float32`

GetLiquidityDepthPlus returns the LiquidityDepthPlus field if non-nil, zero value otherwise.

### GetLiquidityDepthPlusOk

`func (o *GetPair) GetLiquidityDepthPlusOk() (*float32, bool)`

GetLiquidityDepthPlusOk returns a tuple with the LiquidityDepthPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidityDepthPlus

`func (o *GetPair) SetLiquidityDepthPlus(v float32)`

SetLiquidityDepthPlus sets LiquidityDepthPlus field to given value.


### GetLiquidityDepthTokenX

`func (o *GetPair) GetLiquidityDepthTokenX() float32`

GetLiquidityDepthTokenX returns the LiquidityDepthTokenX field if non-nil, zero value otherwise.

### GetLiquidityDepthTokenXOk

`func (o *GetPair) GetLiquidityDepthTokenXOk() (*float32, bool)`

GetLiquidityDepthTokenXOk returns a tuple with the LiquidityDepthTokenX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidityDepthTokenX

`func (o *GetPair) SetLiquidityDepthTokenX(v float32)`

SetLiquidityDepthTokenX sets LiquidityDepthTokenX field to given value.


### GetLiquidityDepthTokenY

`func (o *GetPair) GetLiquidityDepthTokenY() float32`

GetLiquidityDepthTokenY returns the LiquidityDepthTokenY field if non-nil, zero value otherwise.

### GetLiquidityDepthTokenYOk

`func (o *GetPair) GetLiquidityDepthTokenYOk() (*float32, bool)`

GetLiquidityDepthTokenYOk returns a tuple with the LiquidityDepthTokenY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidityDepthTokenY

`func (o *GetPair) SetLiquidityDepthTokenY(v float32)`

SetLiquidityDepthTokenY sets LiquidityDepthTokenY field to given value.


### GetVolumeUsd

`func (o *GetPair) GetVolumeUsd() float32`

GetVolumeUsd returns the VolumeUsd field if non-nil, zero value otherwise.

### GetVolumeUsdOk

`func (o *GetPair) GetVolumeUsdOk() (*float32, bool)`

GetVolumeUsdOk returns a tuple with the VolumeUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUsd

`func (o *GetPair) SetVolumeUsd(v float32)`

SetVolumeUsd sets VolumeUsd field to given value.


### GetVolumeNative

`func (o *GetPair) GetVolumeNative() string`

GetVolumeNative returns the VolumeNative field if non-nil, zero value otherwise.

### GetVolumeNativeOk

`func (o *GetPair) GetVolumeNativeOk() (*string, bool)`

GetVolumeNativeOk returns a tuple with the VolumeNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeNative

`func (o *GetPair) SetVolumeNative(v string)`

SetVolumeNative sets VolumeNative field to given value.


### GetFeesUsd

`func (o *GetPair) GetFeesUsd() float32`

GetFeesUsd returns the FeesUsd field if non-nil, zero value otherwise.

### GetFeesUsdOk

`func (o *GetPair) GetFeesUsdOk() (*float32, bool)`

GetFeesUsdOk returns a tuple with the FeesUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesUsd

`func (o *GetPair) SetFeesUsd(v float32)`

SetFeesUsd sets FeesUsd field to given value.


### GetFeesNative

`func (o *GetPair) GetFeesNative() string`

GetFeesNative returns the FeesNative field if non-nil, zero value otherwise.

### GetFeesNativeOk

`func (o *GetPair) GetFeesNativeOk() (*string, bool)`

GetFeesNativeOk returns a tuple with the FeesNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesNative

`func (o *GetPair) SetFeesNative(v string)`

SetFeesNative sets FeesNative field to given value.


### GetProtocolSharePct

`func (o *GetPair) GetProtocolSharePct() float32`

GetProtocolSharePct returns the ProtocolSharePct field if non-nil, zero value otherwise.

### GetProtocolSharePctOk

`func (o *GetPair) GetProtocolSharePctOk() (*float32, bool)`

GetProtocolSharePctOk returns a tuple with the ProtocolSharePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolSharePct

`func (o *GetPair) SetProtocolSharePct(v float32)`

SetProtocolSharePct sets ProtocolSharePct field to given value.

### HasProtocolSharePct

`func (o *GetPair) HasProtocolSharePct() bool`

HasProtocolSharePct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


