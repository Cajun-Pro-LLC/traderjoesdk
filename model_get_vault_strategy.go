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

// checks if the GetVaultStrategy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVaultStrategy{}

// GetVaultStrategy struct for GetVaultStrategy
type GetVaultStrategy struct {
	Address              string  `json:"address"`
	Chain                Chain   `json:"chain"`
	AumAnnualFeePct      float32 `json:"aumAnnualFeePct"`
	AdditionalProperties map[string]interface{}
}

type _GetVaultStrategy GetVaultStrategy

// NewGetVaultStrategy instantiates a new GetVaultStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVaultStrategy(address string, chain Chain, aumAnnualFeePct float32) *GetVaultStrategy {
	this := GetVaultStrategy{}
	this.Address = address
	this.Chain = chain
	this.AumAnnualFeePct = aumAnnualFeePct
	return &this
}

// NewGetVaultStrategyWithDefaults instantiates a new GetVaultStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVaultStrategyWithDefaults() *GetVaultStrategy {
	this := GetVaultStrategy{}
	return &this
}

// GetAddress returns the Address field value
func (o *GetVaultStrategy) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *GetVaultStrategy) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *GetVaultStrategy) SetAddress(v string) {
	o.Address = v
}

// GetChain returns the Chain field value
func (o *GetVaultStrategy) GetChain() Chain {
	if o == nil {
		var ret Chain
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *GetVaultStrategy) GetChainOk() (*Chain, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *GetVaultStrategy) SetChain(v Chain) {
	o.Chain = v
}

// GetAumAnnualFeePct returns the AumAnnualFeePct field value
func (o *GetVaultStrategy) GetAumAnnualFeePct() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AumAnnualFeePct
}

// GetAumAnnualFeePctOk returns a tuple with the AumAnnualFeePct field value
// and a boolean to check if the value has been set.
func (o *GetVaultStrategy) GetAumAnnualFeePctOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AumAnnualFeePct, true
}

// SetAumAnnualFeePct sets field value
func (o *GetVaultStrategy) SetAumAnnualFeePct(v float32) {
	o.AumAnnualFeePct = v
}

func (o GetVaultStrategy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVaultStrategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["chain"] = o.Chain
	toSerialize["aumAnnualFeePct"] = o.AumAnnualFeePct

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetVaultStrategy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"chain",
		"aumAnnualFeePct",
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

	varGetVaultStrategy := _GetVaultStrategy{}

	err = json.Unmarshal(data, &varGetVaultStrategy)

	if err != nil {
		return err
	}

	*o = GetVaultStrategy(varGetVaultStrategy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "chain")
		delete(additionalProperties, "aumAnnualFeePct")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetVaultStrategy struct {
	value *GetVaultStrategy
	isSet bool
}

func (v NullableGetVaultStrategy) Get() *GetVaultStrategy {
	return v.value
}

func (v *NullableGetVaultStrategy) Set(val *GetVaultStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVaultStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVaultStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVaultStrategy(val *GetVaultStrategy) *NullableGetVaultStrategy {
	return &NullableGetVaultStrategy{value: val, isSet: true}
}

func (v NullableGetVaultStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVaultStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
