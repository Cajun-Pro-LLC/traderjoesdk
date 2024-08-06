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

// LeaderboardFeesSortType An enumeration.
type LeaderboardFeesSortType string

// List of LeaderboardFeesSortType
const (
	FEES LeaderboardFeesSortType = "fees"
)

// All allowed values of LeaderboardFeesSortType enum
var AllowedLeaderboardFeesSortTypeEnumValues = []LeaderboardFeesSortType{
	"fees",
}

func (v *LeaderboardFeesSortType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LeaderboardFeesSortType(value)
	for _, existing := range AllowedLeaderboardFeesSortTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LeaderboardFeesSortType", value)
}

// NewLeaderboardFeesSortTypeFromValue returns a pointer to a valid LeaderboardFeesSortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLeaderboardFeesSortTypeFromValue(v string) (*LeaderboardFeesSortType, error) {
	ev := LeaderboardFeesSortType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LeaderboardFeesSortType: valid values are %v", v, AllowedLeaderboardFeesSortTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LeaderboardFeesSortType) IsValid() bool {
	for _, existing := range AllowedLeaderboardFeesSortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LeaderboardFeesSortType value
func (v LeaderboardFeesSortType) Ptr() *LeaderboardFeesSortType {
	return &v
}

type NullableLeaderboardFeesSortType struct {
	value *LeaderboardFeesSortType
	isSet bool
}

func (v NullableLeaderboardFeesSortType) Get() *LeaderboardFeesSortType {
	return v.value
}

func (v *NullableLeaderboardFeesSortType) Set(val *LeaderboardFeesSortType) {
	v.value = val
	v.isSet = true
}

func (v NullableLeaderboardFeesSortType) IsSet() bool {
	return v.isSet
}

func (v *NullableLeaderboardFeesSortType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLeaderboardFeesSortType(val *LeaderboardFeesSortType) *NullableLeaderboardFeesSortType {
	return &NullableLeaderboardFeesSortType{value: val, isSet: true}
}

func (v NullableLeaderboardFeesSortType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLeaderboardFeesSortType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
