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

// checks if the GetVaultActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVaultActivity{}

// GetVaultActivity struct for GetVaultActivity
type GetVaultActivity struct {
	Date                 time.Time          `json:"date"`
	Timestamp            int32              `json:"timestamp"`
	TransactionHash      string             `json:"transactionHash"`
	Deposits             []VaultBinActivity `json:"deposits"`
	Withdrawals          []VaultBinActivity `json:"withdrawals"`
	AdditionalProperties map[string]interface{}
}

type _GetVaultActivity GetVaultActivity

// NewGetVaultActivity instantiates a new GetVaultActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVaultActivity(date time.Time, timestamp int32, transactionHash string, deposits []VaultBinActivity, withdrawals []VaultBinActivity) *GetVaultActivity {
	this := GetVaultActivity{}
	this.Date = date
	this.Timestamp = timestamp
	this.TransactionHash = transactionHash
	this.Deposits = deposits
	this.Withdrawals = withdrawals
	return &this
}

// NewGetVaultActivityWithDefaults instantiates a new GetVaultActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVaultActivityWithDefaults() *GetVaultActivity {
	this := GetVaultActivity{}
	return &this
}

// GetDate returns the Date field value
func (o *GetVaultActivity) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *GetVaultActivity) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *GetVaultActivity) SetDate(v time.Time) {
	o.Date = v
}

// GetTimestamp returns the Timestamp field value
func (o *GetVaultActivity) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *GetVaultActivity) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *GetVaultActivity) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *GetVaultActivity) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *GetVaultActivity) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *GetVaultActivity) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetDeposits returns the Deposits field value
func (o *GetVaultActivity) GetDeposits() []VaultBinActivity {
	if o == nil {
		var ret []VaultBinActivity
		return ret
	}

	return o.Deposits
}

// GetDepositsOk returns a tuple with the Deposits field value
// and a boolean to check if the value has been set.
func (o *GetVaultActivity) GetDepositsOk() ([]VaultBinActivity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deposits, true
}

// SetDeposits sets field value
func (o *GetVaultActivity) SetDeposits(v []VaultBinActivity) {
	o.Deposits = v
}

// GetWithdrawals returns the Withdrawals field value
func (o *GetVaultActivity) GetWithdrawals() []VaultBinActivity {
	if o == nil {
		var ret []VaultBinActivity
		return ret
	}

	return o.Withdrawals
}

// GetWithdrawalsOk returns a tuple with the Withdrawals field value
// and a boolean to check if the value has been set.
func (o *GetVaultActivity) GetWithdrawalsOk() ([]VaultBinActivity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Withdrawals, true
}

// SetWithdrawals sets field value
func (o *GetVaultActivity) SetWithdrawals(v []VaultBinActivity) {
	o.Withdrawals = v
}

func (o GetVaultActivity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVaultActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["transactionHash"] = o.TransactionHash
	toSerialize["deposits"] = o.Deposits
	toSerialize["withdrawals"] = o.Withdrawals

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetVaultActivity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"timestamp",
		"transactionHash",
		"deposits",
		"withdrawals",
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

	varGetVaultActivity := _GetVaultActivity{}

	err = json.Unmarshal(data, &varGetVaultActivity)

	if err != nil {
		return err
	}

	*o = GetVaultActivity(varGetVaultActivity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "date")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "transactionHash")
		delete(additionalProperties, "deposits")
		delete(additionalProperties, "withdrawals")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetVaultActivity struct {
	value *GetVaultActivity
	isSet bool
}

func (v NullableGetVaultActivity) Get() *GetVaultActivity {
	return v.value
}

func (v *NullableGetVaultActivity) Set(val *GetVaultActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVaultActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVaultActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVaultActivity(val *GetVaultActivity) *NullableGetVaultActivity {
	return &NullableGetVaultActivity{value: val, isSet: true}
}

func (v NullableGetVaultActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVaultActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
