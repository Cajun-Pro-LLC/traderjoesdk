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

// checks if the GetVaultPair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVaultPair{}

// GetVaultPair struct for GetVaultPair
type GetVaultPair struct {
	Address              string      `json:"address"`
	Chain                Chain       `json:"chain"`
	Version              PairVersion `json:"version"`
	BinStep              int32       `json:"binStep"`
	BaseFeePct           float32     `json:"baseFeePct"`
	AdditionalProperties map[string]interface{}
}

type _GetVaultPair GetVaultPair

// NewGetVaultPair instantiates a new GetVaultPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVaultPair(address string, chain Chain, version PairVersion, binStep int32, baseFeePct float32) *GetVaultPair {
	this := GetVaultPair{}
	this.Address = address
	this.Chain = chain
	this.Version = version
	this.BinStep = binStep
	this.BaseFeePct = baseFeePct
	return &this
}

// NewGetVaultPairWithDefaults instantiates a new GetVaultPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVaultPairWithDefaults() *GetVaultPair {
	this := GetVaultPair{}
	return &this
}

// GetAddress returns the Address field value
func (o *GetVaultPair) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *GetVaultPair) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *GetVaultPair) SetAddress(v string) {
	o.Address = v
}

// GetChain returns the Chain field value
func (o *GetVaultPair) GetChain() Chain {
	if o == nil {
		var ret Chain
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *GetVaultPair) GetChainOk() (*Chain, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *GetVaultPair) SetChain(v Chain) {
	o.Chain = v
}

// GetVersion returns the Version field value
func (o *GetVaultPair) GetVersion() PairVersion {
	if o == nil {
		var ret PairVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetVaultPair) GetVersionOk() (*PairVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetVaultPair) SetVersion(v PairVersion) {
	o.Version = v
}

// GetBinStep returns the BinStep field value
func (o *GetVaultPair) GetBinStep() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BinStep
}

// GetBinStepOk returns a tuple with the BinStep field value
// and a boolean to check if the value has been set.
func (o *GetVaultPair) GetBinStepOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinStep, true
}

// SetBinStep sets field value
func (o *GetVaultPair) SetBinStep(v int32) {
	o.BinStep = v
}

// GetBaseFeePct returns the BaseFeePct field value
func (o *GetVaultPair) GetBaseFeePct() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BaseFeePct
}

// GetBaseFeePctOk returns a tuple with the BaseFeePct field value
// and a boolean to check if the value has been set.
func (o *GetVaultPair) GetBaseFeePctOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseFeePct, true
}

// SetBaseFeePct sets field value
func (o *GetVaultPair) SetBaseFeePct(v float32) {
	o.BaseFeePct = v
}

func (o GetVaultPair) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVaultPair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["chain"] = o.Chain
	toSerialize["version"] = o.Version
	toSerialize["binStep"] = o.BinStep
	toSerialize["baseFeePct"] = o.BaseFeePct

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetVaultPair) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"chain",
		"version",
		"binStep",
		"baseFeePct",
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

	varGetVaultPair := _GetVaultPair{}

	err = json.Unmarshal(data, &varGetVaultPair)

	if err != nil {
		return err
	}

	*o = GetVaultPair(varGetVaultPair)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "chain")
		delete(additionalProperties, "version")
		delete(additionalProperties, "binStep")
		delete(additionalProperties, "baseFeePct")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetVaultPair struct {
	value *GetVaultPair
	isSet bool
}

func (v NullableGetVaultPair) Get() *GetVaultPair {
	return v.value
}

func (v *NullableGetVaultPair) Set(val *GetVaultPair) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVaultPair) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVaultPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVaultPair(val *GetVaultPair) *NullableGetVaultPair {
	return &NullableGetVaultPair{value: val, isSet: true}
}

func (v NullableGetVaultPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVaultPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
