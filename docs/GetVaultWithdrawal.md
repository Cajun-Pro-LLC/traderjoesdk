# GetVaultWithdrawal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | [**Chain**](Chain.md) |  | 
**VaultAddress** | **string** |  | 
**UserAddress** | **string** |  | 
**Round** | **int32** |  | 
**Shares** | **int32** |  | 

## Methods

### NewGetVaultWithdrawal

`func NewGetVaultWithdrawal(chain Chain, vaultAddress string, userAddress string, round int32, shares int32, ) *GetVaultWithdrawal`

NewGetVaultWithdrawal instantiates a new GetVaultWithdrawal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVaultWithdrawalWithDefaults

`func NewGetVaultWithdrawalWithDefaults() *GetVaultWithdrawal`

NewGetVaultWithdrawalWithDefaults instantiates a new GetVaultWithdrawal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *GetVaultWithdrawal) GetChain() Chain`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *GetVaultWithdrawal) GetChainOk() (*Chain, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *GetVaultWithdrawal) SetChain(v Chain)`

SetChain sets Chain field to given value.


### GetVaultAddress

`func (o *GetVaultWithdrawal) GetVaultAddress() string`

GetVaultAddress returns the VaultAddress field if non-nil, zero value otherwise.

### GetVaultAddressOk

`func (o *GetVaultWithdrawal) GetVaultAddressOk() (*string, bool)`

GetVaultAddressOk returns a tuple with the VaultAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAddress

`func (o *GetVaultWithdrawal) SetVaultAddress(v string)`

SetVaultAddress sets VaultAddress field to given value.


### GetUserAddress

`func (o *GetVaultWithdrawal) GetUserAddress() string`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *GetVaultWithdrawal) GetUserAddressOk() (*string, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *GetVaultWithdrawal) SetUserAddress(v string)`

SetUserAddress sets UserAddress field to given value.


### GetRound

`func (o *GetVaultWithdrawal) GetRound() int32`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *GetVaultWithdrawal) GetRoundOk() (*int32, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *GetVaultWithdrawal) SetRound(v int32)`

SetRound sets Round field to given value.


### GetShares

`func (o *GetVaultWithdrawal) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *GetVaultWithdrawal) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *GetVaultWithdrawal) SetShares(v int32)`

SetShares sets Shares field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


