# DexAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** |  | 
**Timestamp** | **int32** |  | 
**ReserveUsd** | **float32** |  | 
**ReserveNative** | **float32** |  | 
**VolumeUsd** | **float32** |  | 
**VolumeNative** | **float32** |  | 
**FeesUsd** | **float32** |  | 
**FeesNative** | **float32** |  | 
**ProtocolFeesUsd** | **float32** |  | 
**ProtocolFeesNative** | **float32** |  | 

## Methods

### NewDexAnalytics

`func NewDexAnalytics(date time.Time, timestamp int32, reserveUsd float32, reserveNative float32, volumeUsd float32, volumeNative float32, feesUsd float32, feesNative float32, protocolFeesUsd float32, protocolFeesNative float32, ) *DexAnalytics`

NewDexAnalytics instantiates a new DexAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDexAnalyticsWithDefaults

`func NewDexAnalyticsWithDefaults() *DexAnalytics`

NewDexAnalyticsWithDefaults instantiates a new DexAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *DexAnalytics) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DexAnalytics) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DexAnalytics) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetTimestamp

`func (o *DexAnalytics) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DexAnalytics) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DexAnalytics) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetReserveUsd

`func (o *DexAnalytics) GetReserveUsd() float32`

GetReserveUsd returns the ReserveUsd field if non-nil, zero value otherwise.

### GetReserveUsdOk

`func (o *DexAnalytics) GetReserveUsdOk() (*float32, bool)`

GetReserveUsdOk returns a tuple with the ReserveUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveUsd

`func (o *DexAnalytics) SetReserveUsd(v float32)`

SetReserveUsd sets ReserveUsd field to given value.


### GetReserveNative

`func (o *DexAnalytics) GetReserveNative() float32`

GetReserveNative returns the ReserveNative field if non-nil, zero value otherwise.

### GetReserveNativeOk

`func (o *DexAnalytics) GetReserveNativeOk() (*float32, bool)`

GetReserveNativeOk returns a tuple with the ReserveNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveNative

`func (o *DexAnalytics) SetReserveNative(v float32)`

SetReserveNative sets ReserveNative field to given value.


### GetVolumeUsd

`func (o *DexAnalytics) GetVolumeUsd() float32`

GetVolumeUsd returns the VolumeUsd field if non-nil, zero value otherwise.

### GetVolumeUsdOk

`func (o *DexAnalytics) GetVolumeUsdOk() (*float32, bool)`

GetVolumeUsdOk returns a tuple with the VolumeUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUsd

`func (o *DexAnalytics) SetVolumeUsd(v float32)`

SetVolumeUsd sets VolumeUsd field to given value.


### GetVolumeNative

`func (o *DexAnalytics) GetVolumeNative() float32`

GetVolumeNative returns the VolumeNative field if non-nil, zero value otherwise.

### GetVolumeNativeOk

`func (o *DexAnalytics) GetVolumeNativeOk() (*float32, bool)`

GetVolumeNativeOk returns a tuple with the VolumeNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeNative

`func (o *DexAnalytics) SetVolumeNative(v float32)`

SetVolumeNative sets VolumeNative field to given value.


### GetFeesUsd

`func (o *DexAnalytics) GetFeesUsd() float32`

GetFeesUsd returns the FeesUsd field if non-nil, zero value otherwise.

### GetFeesUsdOk

`func (o *DexAnalytics) GetFeesUsdOk() (*float32, bool)`

GetFeesUsdOk returns a tuple with the FeesUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesUsd

`func (o *DexAnalytics) SetFeesUsd(v float32)`

SetFeesUsd sets FeesUsd field to given value.


### GetFeesNative

`func (o *DexAnalytics) GetFeesNative() float32`

GetFeesNative returns the FeesNative field if non-nil, zero value otherwise.

### GetFeesNativeOk

`func (o *DexAnalytics) GetFeesNativeOk() (*float32, bool)`

GetFeesNativeOk returns a tuple with the FeesNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesNative

`func (o *DexAnalytics) SetFeesNative(v float32)`

SetFeesNative sets FeesNative field to given value.


### GetProtocolFeesUsd

`func (o *DexAnalytics) GetProtocolFeesUsd() float32`

GetProtocolFeesUsd returns the ProtocolFeesUsd field if non-nil, zero value otherwise.

### GetProtocolFeesUsdOk

`func (o *DexAnalytics) GetProtocolFeesUsdOk() (*float32, bool)`

GetProtocolFeesUsdOk returns a tuple with the ProtocolFeesUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFeesUsd

`func (o *DexAnalytics) SetProtocolFeesUsd(v float32)`

SetProtocolFeesUsd sets ProtocolFeesUsd field to given value.


### GetProtocolFeesNative

`func (o *DexAnalytics) GetProtocolFeesNative() float32`

GetProtocolFeesNative returns the ProtocolFeesNative field if non-nil, zero value otherwise.

### GetProtocolFeesNativeOk

`func (o *DexAnalytics) GetProtocolFeesNativeOk() (*float32, bool)`

GetProtocolFeesNativeOk returns a tuple with the ProtocolFeesNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFeesNative

`func (o *DexAnalytics) SetProtocolFeesNative(v float32)`

SetProtocolFeesNative sets ProtocolFeesNative field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


