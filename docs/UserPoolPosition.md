# UserPoolPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolAddress** | **string** |  | 
**PairName** | **string** |  | 
**Status** | [**PairStatus**](PairStatus.md) |  | 
**Version** | [**PairVersion**](PairVersion.md) |  | 
**Chain** | **string** |  | 
**LbBinStep** | **float32** |  | 
**LbBaseFeePct** | **float32** |  | 
**LbMaxFeePct** | **float32** |  | 
**BinIds** | **[]int32** |  | 
**TokenX** | [**UserPoolPositionToken**](UserPoolPositionToken.md) |  | 
**TokenY** | [**UserPoolPositionToken**](UserPoolPositionToken.md) |  | 

## Methods

### NewUserPoolPosition

`func NewUserPoolPosition(poolAddress string, pairName string, status PairStatus, version PairVersion, chain string, lbBinStep float32, lbBaseFeePct float32, lbMaxFeePct float32, binIds []int32, tokenX UserPoolPositionToken, tokenY UserPoolPositionToken, ) *UserPoolPosition`

NewUserPoolPosition instantiates a new UserPoolPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPoolPositionWithDefaults

`func NewUserPoolPositionWithDefaults() *UserPoolPosition`

NewUserPoolPositionWithDefaults instantiates a new UserPoolPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolAddress

`func (o *UserPoolPosition) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *UserPoolPosition) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *UserPoolPosition) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetPairName

`func (o *UserPoolPosition) GetPairName() string`

GetPairName returns the PairName field if non-nil, zero value otherwise.

### GetPairNameOk

`func (o *UserPoolPosition) GetPairNameOk() (*string, bool)`

GetPairNameOk returns a tuple with the PairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairName

`func (o *UserPoolPosition) SetPairName(v string)`

SetPairName sets PairName field to given value.


### GetStatus

`func (o *UserPoolPosition) GetStatus() PairStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserPoolPosition) GetStatusOk() (*PairStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserPoolPosition) SetStatus(v PairStatus)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *UserPoolPosition) GetVersion() PairVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UserPoolPosition) GetVersionOk() (*PairVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UserPoolPosition) SetVersion(v PairVersion)`

SetVersion sets Version field to given value.


### GetChain

`func (o *UserPoolPosition) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *UserPoolPosition) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *UserPoolPosition) SetChain(v string)`

SetChain sets Chain field to given value.


### GetLbBinStep

`func (o *UserPoolPosition) GetLbBinStep() float32`

GetLbBinStep returns the LbBinStep field if non-nil, zero value otherwise.

### GetLbBinStepOk

`func (o *UserPoolPosition) GetLbBinStepOk() (*float32, bool)`

GetLbBinStepOk returns a tuple with the LbBinStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbBinStep

`func (o *UserPoolPosition) SetLbBinStep(v float32)`

SetLbBinStep sets LbBinStep field to given value.


### GetLbBaseFeePct

`func (o *UserPoolPosition) GetLbBaseFeePct() float32`

GetLbBaseFeePct returns the LbBaseFeePct field if non-nil, zero value otherwise.

### GetLbBaseFeePctOk

`func (o *UserPoolPosition) GetLbBaseFeePctOk() (*float32, bool)`

GetLbBaseFeePctOk returns a tuple with the LbBaseFeePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbBaseFeePct

`func (o *UserPoolPosition) SetLbBaseFeePct(v float32)`

SetLbBaseFeePct sets LbBaseFeePct field to given value.


### GetLbMaxFeePct

`func (o *UserPoolPosition) GetLbMaxFeePct() float32`

GetLbMaxFeePct returns the LbMaxFeePct field if non-nil, zero value otherwise.

### GetLbMaxFeePctOk

`func (o *UserPoolPosition) GetLbMaxFeePctOk() (*float32, bool)`

GetLbMaxFeePctOk returns a tuple with the LbMaxFeePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbMaxFeePct

`func (o *UserPoolPosition) SetLbMaxFeePct(v float32)`

SetLbMaxFeePct sets LbMaxFeePct field to given value.


### GetBinIds

`func (o *UserPoolPosition) GetBinIds() []int32`

GetBinIds returns the BinIds field if non-nil, zero value otherwise.

### GetBinIdsOk

`func (o *UserPoolPosition) GetBinIdsOk() (*[]int32, bool)`

GetBinIdsOk returns a tuple with the BinIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinIds

`func (o *UserPoolPosition) SetBinIds(v []int32)`

SetBinIds sets BinIds field to given value.


### GetTokenX

`func (o *UserPoolPosition) GetTokenX() UserPoolPositionToken`

GetTokenX returns the TokenX field if non-nil, zero value otherwise.

### GetTokenXOk

`func (o *UserPoolPosition) GetTokenXOk() (*UserPoolPositionToken, bool)`

GetTokenXOk returns a tuple with the TokenX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenX

`func (o *UserPoolPosition) SetTokenX(v UserPoolPositionToken)`

SetTokenX sets TokenX field to given value.


### GetTokenY

`func (o *UserPoolPosition) GetTokenY() UserPoolPositionToken`

GetTokenY returns the TokenY field if non-nil, zero value otherwise.

### GetTokenYOk

`func (o *UserPoolPosition) GetTokenYOk() (*UserPoolPositionToken, bool)`

GetTokenYOk returns a tuple with the TokenY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenY

`func (o *UserPoolPosition) SetTokenY(v UserPoolPositionToken)`

SetTokenY sets TokenY field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


