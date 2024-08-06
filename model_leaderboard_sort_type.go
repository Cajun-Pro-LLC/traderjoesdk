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
	"gopkg.in/validator.v2"
)

// LeaderboardSortType - An enumeration.
type LeaderboardSortType struct {
	LeaderboardFeesSortType *LeaderboardFeesSortType
	VolumeParam             *VolumeParam
}

// LeaderboardFeesSortTypeAsLeaderboardSortType is a convenience function that returns LeaderboardFeesSortType wrapped in LeaderboardSortType
func LeaderboardFeesSortTypeAsLeaderboardSortType(v *LeaderboardFeesSortType) LeaderboardSortType {
	return LeaderboardSortType{
		LeaderboardFeesSortType: v,
	}
}

// VolumeParamAsLeaderboardSortType is a convenience function that returns VolumeParam wrapped in LeaderboardSortType
func VolumeParamAsLeaderboardSortType(v *VolumeParam) LeaderboardSortType {
	return LeaderboardSortType{
		VolumeParam: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LeaderboardSortType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LeaderboardFeesSortType
	err = newStrictDecoder(data).Decode(&dst.LeaderboardFeesSortType)
	if err == nil {
		jsonLeaderboardFeesSortType, _ := json.Marshal(dst.LeaderboardFeesSortType)
		if string(jsonLeaderboardFeesSortType) == "{}" { // empty struct
			dst.LeaderboardFeesSortType = nil
		} else {
			if err = validator.Validate(dst.LeaderboardFeesSortType); err != nil {
				dst.LeaderboardFeesSortType = nil
			} else {
				match++
			}
		}
	} else {
		dst.LeaderboardFeesSortType = nil
	}

	// try to unmarshal data into VolumeParam
	err = newStrictDecoder(data).Decode(&dst.VolumeParam)
	if err == nil {
		jsonVolumeParam, _ := json.Marshal(dst.VolumeParam)
		if string(jsonVolumeParam) == "{}" { // empty struct
			dst.VolumeParam = nil
		} else {
			if err = validator.Validate(dst.VolumeParam); err != nil {
				dst.VolumeParam = nil
			} else {
				match++
			}
		}
	} else {
		dst.VolumeParam = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LeaderboardFeesSortType = nil
		dst.VolumeParam = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LeaderboardSortType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LeaderboardSortType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LeaderboardSortType) MarshalJSON() ([]byte, error) {
	if src.LeaderboardFeesSortType != nil {
		return json.Marshal(&src.LeaderboardFeesSortType)
	}

	if src.VolumeParam != nil {
		return json.Marshal(&src.VolumeParam)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LeaderboardSortType) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.LeaderboardFeesSortType != nil {
		return obj.LeaderboardFeesSortType
	}

	if obj.VolumeParam != nil {
		return obj.VolumeParam
	}

	// all schemas are nil
	return nil
}

type NullableLeaderboardSortType struct {
	value *LeaderboardSortType
	isSet bool
}

func (v NullableLeaderboardSortType) Get() *LeaderboardSortType {
	return v.value
}

func (v *NullableLeaderboardSortType) Set(val *LeaderboardSortType) {
	v.value = val
	v.isSet = true
}

func (v NullableLeaderboardSortType) IsSet() bool {
	return v.isSet
}

func (v *NullableLeaderboardSortType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLeaderboardSortType(val *LeaderboardSortType) *NullableLeaderboardSortType {
	return &NullableLeaderboardSortType{value: val, isSet: true}
}

func (v NullableLeaderboardSortType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLeaderboardSortType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
