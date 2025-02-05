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
	"time"
)

// checks if the UserFeesEarnedPerBin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFeesEarnedPerBin{}

// UserFeesEarnedPerBin struct for UserFeesEarnedPerBin
type UserFeesEarnedPerBin struct {
	BinId                 int32     `json:"binId"`
	MostRecentDepositTime time.Time `json:"mostRecentDepositTime"`
	Timestamp             int32     `json:"timestamp"`
	AccruedFeesX          float32   `json:"accruedFeesX"`
	AccruedFeesY          float32   `json:"accruedFeesY"`
	AccruedFeesL          float32   `json:"accruedFeesL"`
	PriceXY               float32   `json:"priceXY"`
	PriceYX               float32   `json:"priceYX"`
	AdditionalProperties  map[string]interface{}
}

type _UserFeesEarnedPerBin UserFeesEarnedPerBin

// NewUserFeesEarnedPerBin instantiates a new UserFeesEarnedPerBin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFeesEarnedPerBin(binId int32, mostRecentDepositTime time.Time, timestamp int32, accruedFeesX float32, accruedFeesY float32, accruedFeesL float32, priceXY float32, priceYX float32) *UserFeesEarnedPerBin {
	this := UserFeesEarnedPerBin{}
	this.BinId = binId
	this.MostRecentDepositTime = mostRecentDepositTime
	this.Timestamp = timestamp
	this.AccruedFeesX = accruedFeesX
	this.AccruedFeesY = accruedFeesY
	this.AccruedFeesL = accruedFeesL
	this.PriceXY = priceXY
	this.PriceYX = priceYX
	return &this
}

// NewUserFeesEarnedPerBinWithDefaults instantiates a new UserFeesEarnedPerBin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFeesEarnedPerBinWithDefaults() *UserFeesEarnedPerBin {
	this := UserFeesEarnedPerBin{}
	return &this
}

// GetBinId returns the BinId field value
func (o *UserFeesEarnedPerBin) GetBinId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BinId
}

// GetBinIdOk returns a tuple with the BinId field value
// and a boolean to check if the value has been set.
func (o *UserFeesEarnedPerBin) GetBinIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinId, true
}

// SetBinId sets field value
func (o *UserFeesEarnedPerBin) SetBinId(v int32) {
	o.BinId = v
}

// GetMostRecentDepositTime returns the MostRecentDepositTime field value
func (o *UserFeesEarnedPerBin) GetMostRecentDepositTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.MostRecentDepositTime
}

// GetMostRecentDepositTimeOk returns a tuple with the MostRecentDepositTime field value
// and a boolean to check if the value has been set.
func (o *UserFeesEarnedPerBin) GetMostRecentDepositTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MostRecentDepositTime, true
}

// SetMostRecentDepositTime sets field value
func (o *UserFeesEarnedPerBin) SetMostRecentDepositTime(v time.Time) {
	o.MostRecentDepositTime = v
}

// GetTimestamp returns the Timestamp field value
func (o *UserFeesEarnedPerBin) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *UserFeesEarnedPerBin) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *UserFeesEarnedPerBin) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetAccruedFeesX returns the AccruedFeesX field value
func (o *UserFeesEarnedPerBin) GetAccruedFeesX() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AccruedFeesX
}

// GetAccruedFeesXOk returns a tuple with the AccruedFeesX field value
// and a boolean to check if the value has been set.
func (o *UserFeesEarnedPerBin) GetAccruedFeesXOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccruedFeesX, true
}

// SetAccruedFeesX sets field value
func (o *UserFeesEarnedPerBin) SetAccruedFeesX(v float32) {
	o.AccruedFeesX = v
}

// GetAccruedFeesY returns the AccruedFeesY field value
func (o *UserFeesEarnedPerBin) GetAccruedFeesY() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AccruedFeesY
}

// GetAccruedFeesYOk returns a tuple with the AccruedFeesY field value
// and a boolean to check if the value has been set.
func (o *UserFeesEarnedPerBin) GetAccruedFeesYOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccruedFeesY, true
}

// SetAccruedFeesY sets field value
func (o *UserFeesEarnedPerBin) SetAccruedFeesY(v float32) {
	o.AccruedFeesY = v
}

// GetAccruedFeesL returns the AccruedFeesL field value
func (o *UserFeesEarnedPerBin) GetAccruedFeesL() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AccruedFeesL
}

// GetAccruedFeesLOk returns a tuple with the AccruedFeesL field value
// and a boolean to check if the value has been set.
func (o *UserFeesEarnedPerBin) GetAccruedFeesLOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccruedFeesL, true
}

// SetAccruedFeesL sets field value
func (o *UserFeesEarnedPerBin) SetAccruedFeesL(v float32) {
	o.AccruedFeesL = v
}

// GetPriceXY returns the PriceXY field value
func (o *UserFeesEarnedPerBin) GetPriceXY() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PriceXY
}

// GetPriceXYOk returns a tuple with the PriceXY field value
// and a boolean to check if the value has been set.
func (o *UserFeesEarnedPerBin) GetPriceXYOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceXY, true
}

// SetPriceXY sets field value
func (o *UserFeesEarnedPerBin) SetPriceXY(v float32) {
	o.PriceXY = v
}

// GetPriceYX returns the PriceYX field value
func (o *UserFeesEarnedPerBin) GetPriceYX() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PriceYX
}

// GetPriceYXOk returns a tuple with the PriceYX field value
// and a boolean to check if the value has been set.
func (o *UserFeesEarnedPerBin) GetPriceYXOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceYX, true
}

// SetPriceYX sets field value
func (o *UserFeesEarnedPerBin) SetPriceYX(v float32) {
	o.PriceYX = v
}

func (o UserFeesEarnedPerBin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFeesEarnedPerBin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["binId"] = o.BinId
	toSerialize["mostRecentDepositTime"] = o.MostRecentDepositTime
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["accruedFeesX"] = o.AccruedFeesX
	toSerialize["accruedFeesY"] = o.AccruedFeesY
	toSerialize["accruedFeesL"] = o.AccruedFeesL
	toSerialize["priceXY"] = o.PriceXY
	toSerialize["priceYX"] = o.PriceYX

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFeesEarnedPerBin) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"binId",
		"mostRecentDepositTime",
		"timestamp",
		"accruedFeesX",
		"accruedFeesY",
		"accruedFeesL",
		"priceXY",
		"priceYX",
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

	varUserFeesEarnedPerBin := _UserFeesEarnedPerBin{}

	err = json.Unmarshal(data, &varUserFeesEarnedPerBin)

	if err != nil {
		return err
	}

	*o = UserFeesEarnedPerBin(varUserFeesEarnedPerBin)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "binId")
		delete(additionalProperties, "mostRecentDepositTime")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "accruedFeesX")
		delete(additionalProperties, "accruedFeesY")
		delete(additionalProperties, "accruedFeesL")
		delete(additionalProperties, "priceXY")
		delete(additionalProperties, "priceYX")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserFeesEarnedPerBin struct {
	value *UserFeesEarnedPerBin
	isSet bool
}

func (v NullableUserFeesEarnedPerBin) Get() *UserFeesEarnedPerBin {
	return v.value
}

func (v *NullableUserFeesEarnedPerBin) Set(val *UserFeesEarnedPerBin) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFeesEarnedPerBin) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFeesEarnedPerBin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFeesEarnedPerBin(val *UserFeesEarnedPerBin) *NullableUserFeesEarnedPerBin {
	return &NullableUserFeesEarnedPerBin{value: val, isSet: true}
}

func (v NullableUserFeesEarnedPerBin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFeesEarnedPerBin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
