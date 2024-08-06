# VaultSharePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | [**Chain**](Chain.md) |  | 
**VaultAddress** | **string** |  | 
**BlockNumber** | **int32** |  | 
**SharePrice** | **string** |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewVaultSharePrice

`func NewVaultSharePrice(chain Chain, vaultAddress string, blockNumber int32, sharePrice string, timestamp time.Time, ) *VaultSharePrice`

NewVaultSharePrice instantiates a new VaultSharePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultSharePriceWithDefaults

`func NewVaultSharePriceWithDefaults() *VaultSharePrice`

NewVaultSharePriceWithDefaults instantiates a new VaultSharePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *VaultSharePrice) GetChain() Chain`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *VaultSharePrice) GetChainOk() (*Chain, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *VaultSharePrice) SetChain(v Chain)`

SetChain sets Chain field to given value.


### GetVaultAddress

`func (o *VaultSharePrice) GetVaultAddress() string`

GetVaultAddress returns the VaultAddress field if non-nil, zero value otherwise.

### GetVaultAddressOk

`func (o *VaultSharePrice) GetVaultAddressOk() (*string, bool)`

GetVaultAddressOk returns a tuple with the VaultAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAddress

`func (o *VaultSharePrice) SetVaultAddress(v string)`

SetVaultAddress sets VaultAddress field to given value.


### GetBlockNumber

`func (o *VaultSharePrice) GetBlockNumber() int32`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *VaultSharePrice) GetBlockNumberOk() (*int32, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *VaultSharePrice) SetBlockNumber(v int32)`

SetBlockNumber sets BlockNumber field to given value.


### GetSharePrice

`func (o *VaultSharePrice) GetSharePrice() string`

GetSharePrice returns the SharePrice field if non-nil, zero value otherwise.

### GetSharePriceOk

`func (o *VaultSharePrice) GetSharePriceOk() (*string, bool)`

GetSharePriceOk returns a tuple with the SharePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePrice

`func (o *VaultSharePrice) SetSharePrice(v string)`

SetSharePrice sets SharePrice field to given value.


### GetTimestamp

`func (o *VaultSharePrice) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VaultSharePrice) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VaultSharePrice) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


