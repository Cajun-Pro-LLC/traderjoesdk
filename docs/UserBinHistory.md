# UserBinHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolAddress** | **string** |  | 
**PairName** | **string** |  | 
**BinId** | **int32** |  | 
**LbBinStep** | **float32** |  | 
**LbBaseFeePct** | **float32** |  | 
**LbMaxFeePct** | **float32** |  | 
**TokenX** | [**UserBinPositionTokenWrapper**](UserBinPositionTokenWrapper.md) |  | 
**TokenY** | [**UserBinPositionTokenWrapper**](UserBinPositionTokenWrapper.md) |  | 
**IsDeposit** | **bool** |  | 
**Timestamp** | **time.Time** |  | 
**BlockNumber** | **int32** |  | 

## Methods

### NewUserBinHistory

`func NewUserBinHistory(poolAddress string, pairName string, binId int32, lbBinStep float32, lbBaseFeePct float32, lbMaxFeePct float32, tokenX UserBinPositionTokenWrapper, tokenY UserBinPositionTokenWrapper, isDeposit bool, timestamp time.Time, blockNumber int32, ) *UserBinHistory`

NewUserBinHistory instantiates a new UserBinHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBinHistoryWithDefaults

`func NewUserBinHistoryWithDefaults() *UserBinHistory`

NewUserBinHistoryWithDefaults instantiates a new UserBinHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolAddress

`func (o *UserBinHistory) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *UserBinHistory) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *UserBinHistory) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetPairName

`func (o *UserBinHistory) GetPairName() string`

GetPairName returns the PairName field if non-nil, zero value otherwise.

### GetPairNameOk

`func (o *UserBinHistory) GetPairNameOk() (*string, bool)`

GetPairNameOk returns a tuple with the PairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairName

`func (o *UserBinHistory) SetPairName(v string)`

SetPairName sets PairName field to given value.


### GetBinId

`func (o *UserBinHistory) GetBinId() int32`

GetBinId returns the BinId field if non-nil, zero value otherwise.

### GetBinIdOk

`func (o *UserBinHistory) GetBinIdOk() (*int32, bool)`

GetBinIdOk returns a tuple with the BinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinId

`func (o *UserBinHistory) SetBinId(v int32)`

SetBinId sets BinId field to given value.


### GetLbBinStep

`func (o *UserBinHistory) GetLbBinStep() float32`

GetLbBinStep returns the LbBinStep field if non-nil, zero value otherwise.

### GetLbBinStepOk

`func (o *UserBinHistory) GetLbBinStepOk() (*float32, bool)`

GetLbBinStepOk returns a tuple with the LbBinStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbBinStep

`func (o *UserBinHistory) SetLbBinStep(v float32)`

SetLbBinStep sets LbBinStep field to given value.


### GetLbBaseFeePct

`func (o *UserBinHistory) GetLbBaseFeePct() float32`

GetLbBaseFeePct returns the LbBaseFeePct field if non-nil, zero value otherwise.

### GetLbBaseFeePctOk

`func (o *UserBinHistory) GetLbBaseFeePctOk() (*float32, bool)`

GetLbBaseFeePctOk returns a tuple with the LbBaseFeePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbBaseFeePct

`func (o *UserBinHistory) SetLbBaseFeePct(v float32)`

SetLbBaseFeePct sets LbBaseFeePct field to given value.


### GetLbMaxFeePct

`func (o *UserBinHistory) GetLbMaxFeePct() float32`

GetLbMaxFeePct returns the LbMaxFeePct field if non-nil, zero value otherwise.

### GetLbMaxFeePctOk

`func (o *UserBinHistory) GetLbMaxFeePctOk() (*float32, bool)`

GetLbMaxFeePctOk returns a tuple with the LbMaxFeePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbMaxFeePct

`func (o *UserBinHistory) SetLbMaxFeePct(v float32)`

SetLbMaxFeePct sets LbMaxFeePct field to given value.


### GetTokenX

`func (o *UserBinHistory) GetTokenX() UserBinPositionTokenWrapper`

GetTokenX returns the TokenX field if non-nil, zero value otherwise.

### GetTokenXOk

`func (o *UserBinHistory) GetTokenXOk() (*UserBinPositionTokenWrapper, bool)`

GetTokenXOk returns a tuple with the TokenX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenX

`func (o *UserBinHistory) SetTokenX(v UserBinPositionTokenWrapper)`

SetTokenX sets TokenX field to given value.


### GetTokenY

`func (o *UserBinHistory) GetTokenY() UserBinPositionTokenWrapper`

GetTokenY returns the TokenY field if non-nil, zero value otherwise.

### GetTokenYOk

`func (o *UserBinHistory) GetTokenYOk() (*UserBinPositionTokenWrapper, bool)`

GetTokenYOk returns a tuple with the TokenY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenY

`func (o *UserBinHistory) SetTokenY(v UserBinPositionTokenWrapper)`

SetTokenY sets TokenY field to given value.


### GetIsDeposit

`func (o *UserBinHistory) GetIsDeposit() bool`

GetIsDeposit returns the IsDeposit field if non-nil, zero value otherwise.

### GetIsDepositOk

`func (o *UserBinHistory) GetIsDepositOk() (*bool, bool)`

GetIsDepositOk returns a tuple with the IsDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeposit

`func (o *UserBinHistory) SetIsDeposit(v bool)`

SetIsDeposit sets IsDeposit field to given value.


### GetTimestamp

`func (o *UserBinHistory) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserBinHistory) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserBinHistory) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetBlockNumber

`func (o *UserBinHistory) GetBlockNumber() int32`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *UserBinHistory) GetBlockNumberOk() (*int32, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *UserBinHistory) SetBlockNumber(v int32)`

SetBlockNumber sets BlockNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


