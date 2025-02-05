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

// checks if the UserPoolPosition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPoolPosition{}

// UserPoolPosition struct for UserPoolPosition
type UserPoolPosition struct {
	PoolAddress          string                `json:"poolAddress"`
	PairName             string                `json:"pairName"`
	Status               PairStatus            `json:"status"`
	Version              PairVersion           `json:"version"`
	Chain                string                `json:"chain"`
	LbBinStep            float32               `json:"lbBinStep"`
	LbBaseFeePct         float32               `json:"lbBaseFeePct"`
	LbMaxFeePct          float32               `json:"lbMaxFeePct"`
	BinIds               []int32               `json:"binIds"`
	TokenX               UserPoolPositionToken `json:"tokenX"`
	TokenY               UserPoolPositionToken `json:"tokenY"`
	AdditionalProperties map[string]interface{}
}

type _UserPoolPosition UserPoolPosition

// NewUserPoolPosition instantiates a new UserPoolPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPoolPosition(poolAddress string, pairName string, status PairStatus, version PairVersion, chain string, lbBinStep float32, lbBaseFeePct float32, lbMaxFeePct float32, binIds []int32, tokenX UserPoolPositionToken, tokenY UserPoolPositionToken) *UserPoolPosition {
	this := UserPoolPosition{}
	this.PoolAddress = poolAddress
	this.PairName = pairName
	this.Status = status
	this.Version = version
	this.Chain = chain
	this.LbBinStep = lbBinStep
	this.LbBaseFeePct = lbBaseFeePct
	this.LbMaxFeePct = lbMaxFeePct
	this.BinIds = binIds
	this.TokenX = tokenX
	this.TokenY = tokenY
	return &this
}

// NewUserPoolPositionWithDefaults instantiates a new UserPoolPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPoolPositionWithDefaults() *UserPoolPosition {
	this := UserPoolPosition{}
	return &this
}

// GetPoolAddress returns the PoolAddress field value
func (o *UserPoolPosition) GetPoolAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoolAddress
}

// GetPoolAddressOk returns a tuple with the PoolAddress field value
// and a boolean to check if the value has been set.
func (o *UserPoolPosition) GetPoolAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolAddress, true
}

// SetPoolAddress sets field value
func (o *UserPoolPosition) SetPoolAddress(v string) {
	o.PoolAddress = v
}

// GetPairName returns the PairName field value
func (o *UserPoolPosition) GetPairName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PairName
}

// GetPairNameOk returns a tuple with the PairName field value
// and a boolean to check if the value has been set.
func (o *UserPoolPosition) GetPairNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PairName, true
}

// SetPairName sets field value
func (o *UserPoolPosition) SetPairName(v string) {
	o.PairName = v
}

// GetStatus returns the Status field value
func (o *UserPoolPosition) GetStatus() PairStatus {
	if o == nil {
		var ret PairStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UserPoolPosition) GetStatusOk() (*PairStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UserPoolPosition) SetStatus(v PairStatus) {
	o.Status = v
}

// GetVersion returns the Version field value
func (o *UserPoolPosition) GetVersion() PairVersion {
	if o == nil {
		var ret PairVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *UserPoolPosition) GetVersionOk() (*PairVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *UserPoolPosition) SetVersion(v PairVersion) {
	o.Version = v
}

// GetChain returns the Chain field value
func (o *UserPoolPosition) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *UserPoolPosition) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *UserPoolPosition) SetChain(v string) {
	o.Chain = v
}

// GetLbBinStep returns the LbBinStep field value
func (o *UserPoolPosition) GetLbBinStep() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LbBinStep
}

// GetLbBinStepOk returns a tuple with the LbBinStep field value
// and a boolean to check if the value has been set.
func (o *UserPoolPosition) GetLbBinStepOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LbBinStep, true
}

// SetLbBinStep sets field value
func (o *UserPoolPosition) SetLbBinStep(v float32) {
	o.LbBinStep = v
}

// GetLbBaseFeePct returns the LbBaseFeePct field value
func (o *UserPoolPosition) GetLbBaseFeePct() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LbBaseFeePct
}

// GetLbBaseFeePctOk returns a tuple with the LbBaseFeePct field value
// and a boolean to check if the value has been set.
func (o *UserPoolPosition) GetLbBaseFeePctOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LbBaseFeePct, true
}

// SetLbBaseFeePct sets field value
func (o *UserPoolPosition) SetLbBaseFeePct(v float32) {
	o.LbBaseFeePct = v
}

// GetLbMaxFeePct returns the LbMaxFeePct field value
func (o *UserPoolPosition) GetLbMaxFeePct() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LbMaxFeePct
}

// GetLbMaxFeePctOk returns a tuple with the LbMaxFeePct field value
// and a boolean to check if the value has been set.
func (o *UserPoolPosition) GetLbMaxFeePctOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LbMaxFeePct, true
}

// SetLbMaxFeePct sets field value
func (o *UserPoolPosition) SetLbMaxFeePct(v float32) {
	o.LbMaxFeePct = v
}

// GetBinIds returns the BinIds field value
func (o *UserPoolPosition) GetBinIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.BinIds
}

// GetBinIdsOk returns a tuple with the BinIds field value
// and a boolean to check if the value has been set.
func (o *UserPoolPosition) GetBinIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BinIds, true
}

// SetBinIds sets field value
func (o *UserPoolPosition) SetBinIds(v []int32) {
	o.BinIds = v
}

// GetTokenX returns the TokenX field value
func (o *UserPoolPosition) GetTokenX() UserPoolPositionToken {
	if o == nil {
		var ret UserPoolPositionToken
		return ret
	}

	return o.TokenX
}

// GetTokenXOk returns a tuple with the TokenX field value
// and a boolean to check if the value has been set.
func (o *UserPoolPosition) GetTokenXOk() (*UserPoolPositionToken, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenX, true
}

// SetTokenX sets field value
func (o *UserPoolPosition) SetTokenX(v UserPoolPositionToken) {
	o.TokenX = v
}

// GetTokenY returns the TokenY field value
func (o *UserPoolPosition) GetTokenY() UserPoolPositionToken {
	if o == nil {
		var ret UserPoolPositionToken
		return ret
	}

	return o.TokenY
}

// GetTokenYOk returns a tuple with the TokenY field value
// and a boolean to check if the value has been set.
func (o *UserPoolPosition) GetTokenYOk() (*UserPoolPositionToken, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenY, true
}

// SetTokenY sets field value
func (o *UserPoolPosition) SetTokenY(v UserPoolPositionToken) {
	o.TokenY = v
}

func (o UserPoolPosition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPoolPosition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["poolAddress"] = o.PoolAddress
	toSerialize["pairName"] = o.PairName
	toSerialize["status"] = o.Status
	toSerialize["version"] = o.Version
	toSerialize["chain"] = o.Chain
	toSerialize["lbBinStep"] = o.LbBinStep
	toSerialize["lbBaseFeePct"] = o.LbBaseFeePct
	toSerialize["lbMaxFeePct"] = o.LbMaxFeePct
	toSerialize["binIds"] = o.BinIds
	toSerialize["tokenX"] = o.TokenX
	toSerialize["tokenY"] = o.TokenY

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserPoolPosition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"poolAddress",
		"pairName",
		"status",
		"version",
		"chain",
		"lbBinStep",
		"lbBaseFeePct",
		"lbMaxFeePct",
		"binIds",
		"tokenX",
		"tokenY",
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

	varUserPoolPosition := _UserPoolPosition{}

	err = json.Unmarshal(data, &varUserPoolPosition)

	if err != nil {
		return err
	}

	*o = UserPoolPosition(varUserPoolPosition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "poolAddress")
		delete(additionalProperties, "pairName")
		delete(additionalProperties, "status")
		delete(additionalProperties, "version")
		delete(additionalProperties, "chain")
		delete(additionalProperties, "lbBinStep")
		delete(additionalProperties, "lbBaseFeePct")
		delete(additionalProperties, "lbMaxFeePct")
		delete(additionalProperties, "binIds")
		delete(additionalProperties, "tokenX")
		delete(additionalProperties, "tokenY")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserPoolPosition struct {
	value *UserPoolPosition
	isSet bool
}

func (v NullableUserPoolPosition) Get() *UserPoolPosition {
	return v.value
}

func (v *NullableUserPoolPosition) Set(val *UserPoolPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPoolPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPoolPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPoolPosition(val *UserPoolPosition) *NullableUserPoolPosition {
	return &NullableUserPoolPosition{value: val, isSet: true}
}

func (v NullableUserPoolPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPoolPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
