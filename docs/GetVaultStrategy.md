# GetVaultStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Chain** | [**Chain**](Chain.md) |  | 
**AumAnnualFeePct** | **float32** |  | 

## Methods

### NewGetVaultStrategy

`func NewGetVaultStrategy(address string, chain Chain, aumAnnualFeePct float32, ) *GetVaultStrategy`

NewGetVaultStrategy instantiates a new GetVaultStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVaultStrategyWithDefaults

`func NewGetVaultStrategyWithDefaults() *GetVaultStrategy`

NewGetVaultStrategyWithDefaults instantiates a new GetVaultStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GetVaultStrategy) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetVaultStrategy) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetVaultStrategy) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChain

`func (o *GetVaultStrategy) GetChain() Chain`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *GetVaultStrategy) GetChainOk() (*Chain, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *GetVaultStrategy) SetChain(v Chain)`

SetChain sets Chain field to given value.


### GetAumAnnualFeePct

`func (o *GetVaultStrategy) GetAumAnnualFeePct() float32`

GetAumAnnualFeePct returns the AumAnnualFeePct field if non-nil, zero value otherwise.

### GetAumAnnualFeePctOk

`func (o *GetVaultStrategy) GetAumAnnualFeePctOk() (*float32, bool)`

GetAumAnnualFeePctOk returns a tuple with the AumAnnualFeePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAumAnnualFeePct

`func (o *GetVaultStrategy) SetAumAnnualFeePct(v float32)`

SetAumAnnualFeePct sets AumAnnualFeePct field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


