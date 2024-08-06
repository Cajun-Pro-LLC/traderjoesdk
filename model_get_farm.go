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

// checks if the GetFarm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFarm{}

// GetFarm struct for GetFarm
type GetFarm struct {
	FarmId               string       `json:"farmId"`
	VaultId              string       `json:"vaultId"`
	Liquidity            string       `json:"liquidity"`
	LiquidityRaw         int32        `json:"liquidityRaw"`
	LiquidityUsd         float32      `json:"liquidityUsd"`
	AptDecimals          int32        `json:"aptDecimals"`
	Apr1d                float32      `json:"apr1d"`
	RewardsPerSec        float32      `json:"rewardsPerSec"`
	Reward               GetReward    `json:"reward"`
	Rewarder             *GetRewarder `json:"rewarder,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFarm GetFarm

// NewGetFarm instantiates a new GetFarm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFarm(farmId string, vaultId string, liquidity string, liquidityRaw int32, liquidityUsd float32, aptDecimals int32, apr1d float32, rewardsPerSec float32, reward GetReward) *GetFarm {
	this := GetFarm{}
	this.FarmId = farmId
	this.VaultId = vaultId
	this.Liquidity = liquidity
	this.LiquidityRaw = liquidityRaw
	this.LiquidityUsd = liquidityUsd
	this.AptDecimals = aptDecimals
	this.Apr1d = apr1d
	this.RewardsPerSec = rewardsPerSec
	this.Reward = reward
	return &this
}

// NewGetFarmWithDefaults instantiates a new GetFarm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFarmWithDefaults() *GetFarm {
	this := GetFarm{}
	return &this
}

// GetFarmId returns the FarmId field value
func (o *GetFarm) GetFarmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FarmId
}

// GetFarmIdOk returns a tuple with the FarmId field value
// and a boolean to check if the value has been set.
func (o *GetFarm) GetFarmIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FarmId, true
}

// SetFarmId sets field value
func (o *GetFarm) SetFarmId(v string) {
	o.FarmId = v
}

// GetVaultId returns the VaultId field value
func (o *GetFarm) GetVaultId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultId
}

// GetVaultIdOk returns a tuple with the VaultId field value
// and a boolean to check if the value has been set.
func (o *GetFarm) GetVaultIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultId, true
}

// SetVaultId sets field value
func (o *GetFarm) SetVaultId(v string) {
	o.VaultId = v
}

// GetLiquidity returns the Liquidity field value
func (o *GetFarm) GetLiquidity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Liquidity
}

// GetLiquidityOk returns a tuple with the Liquidity field value
// and a boolean to check if the value has been set.
func (o *GetFarm) GetLiquidityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Liquidity, true
}

// SetLiquidity sets field value
func (o *GetFarm) SetLiquidity(v string) {
	o.Liquidity = v
}

// GetLiquidityRaw returns the LiquidityRaw field value
func (o *GetFarm) GetLiquidityRaw() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LiquidityRaw
}

// GetLiquidityRawOk returns a tuple with the LiquidityRaw field value
// and a boolean to check if the value has been set.
func (o *GetFarm) GetLiquidityRawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LiquidityRaw, true
}

// SetLiquidityRaw sets field value
func (o *GetFarm) SetLiquidityRaw(v int32) {
	o.LiquidityRaw = v
}

// GetLiquidityUsd returns the LiquidityUsd field value
func (o *GetFarm) GetLiquidityUsd() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LiquidityUsd
}

// GetLiquidityUsdOk returns a tuple with the LiquidityUsd field value
// and a boolean to check if the value has been set.
func (o *GetFarm) GetLiquidityUsdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LiquidityUsd, true
}

// SetLiquidityUsd sets field value
func (o *GetFarm) SetLiquidityUsd(v float32) {
	o.LiquidityUsd = v
}

// GetAptDecimals returns the AptDecimals field value
func (o *GetFarm) GetAptDecimals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AptDecimals
}

// GetAptDecimalsOk returns a tuple with the AptDecimals field value
// and a boolean to check if the value has been set.
func (o *GetFarm) GetAptDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AptDecimals, true
}

// SetAptDecimals sets field value
func (o *GetFarm) SetAptDecimals(v int32) {
	o.AptDecimals = v
}

// GetApr1d returns the Apr1d field value
func (o *GetFarm) GetApr1d() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Apr1d
}

// GetApr1dOk returns a tuple with the Apr1d field value
// and a boolean to check if the value has been set.
func (o *GetFarm) GetApr1dOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Apr1d, true
}

// SetApr1d sets field value
func (o *GetFarm) SetApr1d(v float32) {
	o.Apr1d = v
}

// GetRewardsPerSec returns the RewardsPerSec field value
func (o *GetFarm) GetRewardsPerSec() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RewardsPerSec
}

// GetRewardsPerSecOk returns a tuple with the RewardsPerSec field value
// and a boolean to check if the value has been set.
func (o *GetFarm) GetRewardsPerSecOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardsPerSec, true
}

// SetRewardsPerSec sets field value
func (o *GetFarm) SetRewardsPerSec(v float32) {
	o.RewardsPerSec = v
}

// GetReward returns the Reward field value
func (o *GetFarm) GetReward() GetReward {
	if o == nil {
		var ret GetReward
		return ret
	}

	return o.Reward
}

// GetRewardOk returns a tuple with the Reward field value
// and a boolean to check if the value has been set.
func (o *GetFarm) GetRewardOk() (*GetReward, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reward, true
}

// SetReward sets field value
func (o *GetFarm) SetReward(v GetReward) {
	o.Reward = v
}

// GetRewarder returns the Rewarder field value if set, zero value otherwise.
func (o *GetFarm) GetRewarder() GetRewarder {
	if o == nil || IsNil(o.Rewarder) {
		var ret GetRewarder
		return ret
	}
	return *o.Rewarder
}

// GetRewarderOk returns a tuple with the Rewarder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFarm) GetRewarderOk() (*GetRewarder, bool) {
	if o == nil || IsNil(o.Rewarder) {
		return nil, false
	}
	return o.Rewarder, true
}

// HasRewarder returns a boolean if a field has been set.
func (o *GetFarm) HasRewarder() bool {
	if o != nil && !IsNil(o.Rewarder) {
		return true
	}

	return false
}

// SetRewarder gets a reference to the given GetRewarder and assigns it to the Rewarder field.
func (o *GetFarm) SetRewarder(v GetRewarder) {
	o.Rewarder = &v
}

func (o GetFarm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFarm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["farmId"] = o.FarmId
	toSerialize["vaultId"] = o.VaultId
	toSerialize["liquidity"] = o.Liquidity
	toSerialize["liquidityRaw"] = o.LiquidityRaw
	toSerialize["liquidityUsd"] = o.LiquidityUsd
	toSerialize["aptDecimals"] = o.AptDecimals
	toSerialize["apr1d"] = o.Apr1d
	toSerialize["rewardsPerSec"] = o.RewardsPerSec
	toSerialize["reward"] = o.Reward
	if !IsNil(o.Rewarder) {
		toSerialize["rewarder"] = o.Rewarder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFarm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"farmId",
		"vaultId",
		"liquidity",
		"liquidityRaw",
		"liquidityUsd",
		"aptDecimals",
		"apr1d",
		"rewardsPerSec",
		"reward",
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

	varGetFarm := _GetFarm{}

	err = json.Unmarshal(data, &varGetFarm)

	if err != nil {
		return err
	}

	*o = GetFarm(varGetFarm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "farmId")
		delete(additionalProperties, "vaultId")
		delete(additionalProperties, "liquidity")
		delete(additionalProperties, "liquidityRaw")
		delete(additionalProperties, "liquidityUsd")
		delete(additionalProperties, "aptDecimals")
		delete(additionalProperties, "apr1d")
		delete(additionalProperties, "rewardsPerSec")
		delete(additionalProperties, "reward")
		delete(additionalProperties, "rewarder")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFarm struct {
	value *GetFarm
	isSet bool
}

func (v NullableGetFarm) Get() *GetFarm {
	return v.value
}

func (v *NullableGetFarm) Set(val *GetFarm) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFarm) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFarm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFarm(val *GetFarm) *NullableGetFarm {
	return &NullableGetFarm{value: val, isSet: true}
}

func (v NullableGetFarm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFarm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
