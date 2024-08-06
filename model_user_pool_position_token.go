/*
Trader Joe Dex API

<p>Discover DeFi with Trader Joe, a leading decentralized exchange. Trade a wide variety of tokens, earn rewards, and engage in secure, peer-to-peer transactions. Trader Joe makes DeFi easy and accessible.

API version: 1.0.0
Contact: public-api@traderjoexyz.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package traderjoesdk

import (
	"encoding/json"
	"fmt"
)

// checks if the UserPoolPositionToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPoolPositionToken{}

// UserPoolPositionToken struct for UserPoolPositionToken
type UserPoolPositionToken struct {
	Address              string  `json:"address"`
	Name                 string  `json:"name"`
	Symbol               string  `json:"symbol"`
	Decimals             int32   `json:"decimals"`
	PriceUsd             float32 `json:"priceUsd"`
	AdditionalProperties map[string]interface{}
}

type _UserPoolPositionToken UserPoolPositionToken

// NewUserPoolPositionToken instantiates a new UserPoolPositionToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPoolPositionToken(address string, name string, symbol string, decimals int32, priceUsd float32) *UserPoolPositionToken {
	this := UserPoolPositionToken{}
	this.Address = address
	this.Name = name
	this.Symbol = symbol
	this.Decimals = decimals
	this.PriceUsd = priceUsd
	return &this
}

// NewUserPoolPositionTokenWithDefaults instantiates a new UserPoolPositionToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPoolPositionTokenWithDefaults() *UserPoolPositionToken {
	this := UserPoolPositionToken{}
	return &this
}

// GetAddress returns the Address field value
func (o *UserPoolPositionToken) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *UserPoolPositionToken) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *UserPoolPositionToken) SetAddress(v string) {
	o.Address = v
}

// GetName returns the Name field value
func (o *UserPoolPositionToken) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserPoolPositionToken) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserPoolPositionToken) SetName(v string) {
	o.Name = v
}

// GetSymbol returns the Symbol field value
func (o *UserPoolPositionToken) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *UserPoolPositionToken) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *UserPoolPositionToken) SetSymbol(v string) {
	o.Symbol = v
}

// GetDecimals returns the Decimals field value
func (o *UserPoolPositionToken) GetDecimals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value
// and a boolean to check if the value has been set.
func (o *UserPoolPositionToken) GetDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decimals, true
}

// SetDecimals sets field value
func (o *UserPoolPositionToken) SetDecimals(v int32) {
	o.Decimals = v
}

// GetPriceUsd returns the PriceUsd field value
func (o *UserPoolPositionToken) GetPriceUsd() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PriceUsd
}

// GetPriceUsdOk returns a tuple with the PriceUsd field value
// and a boolean to check if the value has been set.
func (o *UserPoolPositionToken) GetPriceUsdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceUsd, true
}

// SetPriceUsd sets field value
func (o *UserPoolPositionToken) SetPriceUsd(v float32) {
	o.PriceUsd = v
}

func (o UserPoolPositionToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPoolPositionToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["name"] = o.Name
	toSerialize["symbol"] = o.Symbol
	toSerialize["decimals"] = o.Decimals
	toSerialize["priceUsd"] = o.PriceUsd

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserPoolPositionToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"name",
		"symbol",
		"decimals",
		"priceUsd",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserPoolPositionToken := _UserPoolPositionToken{}

	err = json.Unmarshal(data, &varUserPoolPositionToken)

	if err != nil {
		return err
	}

	*o = UserPoolPositionToken(varUserPoolPositionToken)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "name")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "decimals")
		delete(additionalProperties, "priceUsd")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserPoolPositionToken struct {
	value *UserPoolPositionToken
	isSet bool
}

func (v NullableUserPoolPositionToken) Get() *UserPoolPositionToken {
	return v.value
}

func (v *NullableUserPoolPositionToken) Set(val *UserPoolPositionToken) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPoolPositionToken) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPoolPositionToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPoolPositionToken(val *UserPoolPositionToken) *NullableUserPoolPositionToken {
	return &NullableUserPoolPositionToken{value: val, isSet: true}
}

func (v NullableUserPoolPositionToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPoolPositionToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
