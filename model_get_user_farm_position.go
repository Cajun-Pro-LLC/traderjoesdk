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

// checks if the GetUserFarmPosition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserFarmPosition{}

// GetUserFarmPosition struct for GetUserFarmPosition
type GetUserFarmPosition struct {
	FarmId               string  `json:"farmId"`
	UserPosition         string  `json:"userPosition"`
	UserPositionRaw      string  `json:"userPositionRaw"`
	UserPositionUsd      string  `json:"userPositionUsd"`
	PendingJoe           float32 `json:"pendingJoe"`
	AdditionalProperties map[string]interface{}
}

type _GetUserFarmPosition GetUserFarmPosition

// NewGetUserFarmPosition instantiates a new GetUserFarmPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserFarmPosition(farmId string, userPosition string, userPositionRaw string, userPositionUsd string, pendingJoe float32) *GetUserFarmPosition {
	this := GetUserFarmPosition{}
	this.FarmId = farmId
	this.UserPosition = userPosition
	this.UserPositionRaw = userPositionRaw
	this.UserPositionUsd = userPositionUsd
	this.PendingJoe = pendingJoe
	return &this
}

// NewGetUserFarmPositionWithDefaults instantiates a new GetUserFarmPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserFarmPositionWithDefaults() *GetUserFarmPosition {
	this := GetUserFarmPosition{}
	return &this
}

// GetFarmId returns the FarmId field value
func (o *GetUserFarmPosition) GetFarmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FarmId
}

// GetFarmIdOk returns a tuple with the FarmId field value
// and a boolean to check if the value has been set.
func (o *GetUserFarmPosition) GetFarmIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FarmId, true
}

// SetFarmId sets field value
func (o *GetUserFarmPosition) SetFarmId(v string) {
	o.FarmId = v
}

// GetUserPosition returns the UserPosition field value
func (o *GetUserFarmPosition) GetUserPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserPosition
}

// GetUserPositionOk returns a tuple with the UserPosition field value
// and a boolean to check if the value has been set.
func (o *GetUserFarmPosition) GetUserPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserPosition, true
}

// SetUserPosition sets field value
func (o *GetUserFarmPosition) SetUserPosition(v string) {
	o.UserPosition = v
}

// GetUserPositionRaw returns the UserPositionRaw field value
func (o *GetUserFarmPosition) GetUserPositionRaw() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserPositionRaw
}

// GetUserPositionRawOk returns a tuple with the UserPositionRaw field value
// and a boolean to check if the value has been set.
func (o *GetUserFarmPosition) GetUserPositionRawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserPositionRaw, true
}

// SetUserPositionRaw sets field value
func (o *GetUserFarmPosition) SetUserPositionRaw(v string) {
	o.UserPositionRaw = v
}

// GetUserPositionUsd returns the UserPositionUsd field value
func (o *GetUserFarmPosition) GetUserPositionUsd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserPositionUsd
}

// GetUserPositionUsdOk returns a tuple with the UserPositionUsd field value
// and a boolean to check if the value has been set.
func (o *GetUserFarmPosition) GetUserPositionUsdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserPositionUsd, true
}

// SetUserPositionUsd sets field value
func (o *GetUserFarmPosition) SetUserPositionUsd(v string) {
	o.UserPositionUsd = v
}

// GetPendingJoe returns the PendingJoe field value
func (o *GetUserFarmPosition) GetPendingJoe() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PendingJoe
}

// GetPendingJoeOk returns a tuple with the PendingJoe field value
// and a boolean to check if the value has been set.
func (o *GetUserFarmPosition) GetPendingJoeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingJoe, true
}

// SetPendingJoe sets field value
func (o *GetUserFarmPosition) SetPendingJoe(v float32) {
	o.PendingJoe = v
}

func (o GetUserFarmPosition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserFarmPosition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["farmId"] = o.FarmId
	toSerialize["userPosition"] = o.UserPosition
	toSerialize["userPositionRaw"] = o.UserPositionRaw
	toSerialize["userPositionUsd"] = o.UserPositionUsd
	toSerialize["pendingJoe"] = o.PendingJoe

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUserFarmPosition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"farmId",
		"userPosition",
		"userPositionRaw",
		"userPositionUsd",
		"pendingJoe",
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

	varGetUserFarmPosition := _GetUserFarmPosition{}

	err = json.Unmarshal(data, &varGetUserFarmPosition)

	if err != nil {
		return err
	}

	*o = GetUserFarmPosition(varGetUserFarmPosition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "farmId")
		delete(additionalProperties, "userPosition")
		delete(additionalProperties, "userPositionRaw")
		delete(additionalProperties, "userPositionUsd")
		delete(additionalProperties, "pendingJoe")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUserFarmPosition struct {
	value *GetUserFarmPosition
	isSet bool
}

func (v NullableGetUserFarmPosition) Get() *GetUserFarmPosition {
	return v.value
}

func (v *NullableGetUserFarmPosition) Set(val *GetUserFarmPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserFarmPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserFarmPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserFarmPosition(val *GetUserFarmPosition) *NullableGetUserFarmPosition {
	return &NullableGetUserFarmPosition{value: val, isSet: true}
}

func (v NullableGetUserFarmPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserFarmPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
