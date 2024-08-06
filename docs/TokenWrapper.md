# TokenWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Name** | **string** |  | 
**Symbol** | **string** |  | 
**Decimals** | **int32** |  | 
**PriceUsd** | **float32** |  | 
**PriceNative** | **string** |  | 

## Methods

### NewTokenWrapper

`func NewTokenWrapper(address string, name string, symbol string, decimals int32, priceUsd float32, priceNative string, ) *TokenWrapper`

NewTokenWrapper instantiates a new TokenWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWrapperWithDefaults

`func NewTokenWrapperWithDefaults() *TokenWrapper`

NewTokenWrapperWithDefaults instantiates a new TokenWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TokenWrapper) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenWrapper) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenWrapper) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetName

`func (o *TokenWrapper) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenWrapper) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenWrapper) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *TokenWrapper) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenWrapper) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenWrapper) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimals

`func (o *TokenWrapper) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenWrapper) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenWrapper) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetPriceUsd

`func (o *TokenWrapper) GetPriceUsd() float32`

GetPriceUsd returns the PriceUsd field if non-nil, zero value otherwise.

### GetPriceUsdOk

`func (o *TokenWrapper) GetPriceUsdOk() (*float32, bool)`

GetPriceUsdOk returns a tuple with the PriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUsd

`func (o *TokenWrapper) SetPriceUsd(v float32)`

SetPriceUsd sets PriceUsd field to given value.


### GetPriceNative

`func (o *TokenWrapper) GetPriceNative() string`

GetPriceNative returns the PriceNative field if non-nil, zero value otherwise.

### GetPriceNativeOk

`func (o *TokenWrapper) GetPriceNativeOk() (*string, bool)`

GetPriceNativeOk returns a tuple with the PriceNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNative

`func (o *TokenWrapper) SetPriceNative(v string)`

SetPriceNative sets PriceNative field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


