# GetVaultActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** |  | 
**Timestamp** | **int32** |  | 
**TransactionHash** | **string** |  | 
**Deposits** | [**[]VaultBinActivity**](VaultBinActivity.md) |  | 
**Withdrawals** | [**[]VaultBinActivity**](VaultBinActivity.md) |  | 

## Methods

### NewGetVaultActivity

`func NewGetVaultActivity(date time.Time, timestamp int32, transactionHash string, deposits []VaultBinActivity, withdrawals []VaultBinActivity, ) *GetVaultActivity`

NewGetVaultActivity instantiates a new GetVaultActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVaultActivityWithDefaults

`func NewGetVaultActivityWithDefaults() *GetVaultActivity`

NewGetVaultActivityWithDefaults instantiates a new GetVaultActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetVaultActivity) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetVaultActivity) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetVaultActivity) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetTimestamp

`func (o *GetVaultActivity) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetVaultActivity) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetVaultActivity) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *GetVaultActivity) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *GetVaultActivity) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *GetVaultActivity) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetDeposits

`func (o *GetVaultActivity) GetDeposits() []VaultBinActivity`

GetDeposits returns the Deposits field if non-nil, zero value otherwise.

### GetDepositsOk

`func (o *GetVaultActivity) GetDepositsOk() (*[]VaultBinActivity, bool)`

GetDepositsOk returns a tuple with the Deposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposits

`func (o *GetVaultActivity) SetDeposits(v []VaultBinActivity)`

SetDeposits sets Deposits field to given value.


### GetWithdrawals

`func (o *GetVaultActivity) GetWithdrawals() []VaultBinActivity`

GetWithdrawals returns the Withdrawals field if non-nil, zero value otherwise.

### GetWithdrawalsOk

`func (o *GetVaultActivity) GetWithdrawalsOk() (*[]VaultBinActivity, bool)`

GetWithdrawalsOk returns a tuple with the Withdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawals

`func (o *GetVaultActivity) SetWithdrawals(v []VaultBinActivity)`

SetWithdrawals sets Withdrawals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


