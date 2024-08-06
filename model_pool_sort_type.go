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

// PoolSortType An enumeration.
type PoolSortType string

// List of PoolSortType
const (
	pool_sort_liquidity PoolSortType = "liquidity"
	pool_sort_name      PoolSortType = "name"
	pool_sort_volume    PoolSortType = "volume"
)

// All allowed values of PoolSortType enum
var AllowedPoolSortTypeEnumValues = []PoolSortType{
	"liquidity",
	"name",
	"volume",
}

func (v *PoolSortType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PoolSortType(value)
	for _, existing := range AllowedPoolSortTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PoolSortType", value)
}

// NewPoolSortTypeFromValue returns a pointer to a valid PoolSortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPoolSortTypeFromValue(v string) (*PoolSortType, error) {
	ev := PoolSortType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PoolSortType: valid values are %v", v, AllowedPoolSortTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PoolSortType) IsValid() bool {
	for _, existing := range AllowedPoolSortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PoolSortType value
func (v PoolSortType) Ptr() *PoolSortType {
	return &v
}

type NullablePoolSortType struct {
	value *PoolSortType
	isSet bool
}

func (v NullablePoolSortType) Get() *PoolSortType {
	return v.value
}

func (v *NullablePoolSortType) Set(val *PoolSortType) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolSortType) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolSortType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolSortType(val *PoolSortType) *NullablePoolSortType {
	return &NullablePoolSortType{value: val, isSet: true}
}

func (v NullablePoolSortType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolSortType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
