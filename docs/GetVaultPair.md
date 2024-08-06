# GetVaultPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Chain** | [**Chain**](Chain.md) |  | 
**Version** | [**PairVersion**](PairVersion.md) |  | 
**BinStep** | **int32** |  | 
**BaseFeePct** | **float32** |  | 

## Methods

### NewGetVaultPair

`func NewGetVaultPair(address string, chain Chain, version PairVersion, binStep int32, baseFeePct float32, ) *GetVaultPair`

NewGetVaultPair instantiates a new GetVaultPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVaultPairWithDefaults

`func NewGetVaultPairWithDefaults() *GetVaultPair`

NewGetVaultPairWithDefaults instantiates a new GetVaultPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GetVaultPair) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetVaultPair) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetVaultPair) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChain

`func (o *GetVaultPair) GetChain() Chain`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *GetVaultPair) GetChainOk() (*Chain, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *GetVaultPair) SetChain(v Chain)`

SetChain sets Chain field to given value.


### GetVersion

`func (o *GetVaultPair) GetVersion() PairVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetVaultPair) GetVersionOk() (*PairVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetVaultPair) SetVersion(v PairVersion)`

SetVersion sets Version field to given value.


### GetBinStep

`func (o *GetVaultPair) GetBinStep() int32`

GetBinStep returns the BinStep field if non-nil, zero value otherwise.

### GetBinStepOk

`func (o *GetVaultPair) GetBinStepOk() (*int32, bool)`

GetBinStepOk returns a tuple with the BinStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinStep

`func (o *GetVaultPair) SetBinStep(v int32)`

SetBinStep sets BinStep field to given value.


### GetBaseFeePct

`func (o *GetVaultPair) GetBaseFeePct() float32`

GetBaseFeePct returns the BaseFeePct field if non-nil, zero value otherwise.

### GetBaseFeePctOk

`func (o *GetVaultPair) GetBaseFeePctOk() (*float32, bool)`

GetBaseFeePctOk returns a tuple with the BaseFeePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFeePct

`func (o *GetVaultPair) SetBaseFeePct(v float32)`

SetBaseFeePct sets BaseFeePct field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


