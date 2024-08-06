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

// AllParam An enumeration
type AllParam string

// List of AllParam
const (
	ALL AllParam = "all"
)

// All allowed values of AllParam enum
var AllowedAllParamEnumValues = []AllParam{
	"all",
}

func (v *AllParam) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AllParam(value)
	for _, existing := range AllowedAllParamEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AllParam", value)
}

// NewAllParamFromValue returns a pointer to a valid AllParam
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAllParamFromValue(v string) (*AllParam, error) {
	ev := AllParam(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AllParam: valid values are %v", v, AllowedAllParamEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AllParam) IsValid() bool {
	for _, existing := range AllowedAllParamEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AllParam value
func (v AllParam) Ptr() *AllParam {
	return &v
}

type NullableAllParam struct {
	value *AllParam
	isSet bool
}

func (v NullableAllParam) Get() *AllParam {
	return v.value
}

func (v *NullableAllParam) Set(val *AllParam) {
	v.value = val
	v.isSet = true
}

func (v NullableAllParam) IsSet() bool {
	return v.isSet
}

func (v *NullableAllParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllParam(val *AllParam) *NullableAllParam {
	return &NullableAllParam{value: val, isSet: true}
}

func (v NullableAllParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
