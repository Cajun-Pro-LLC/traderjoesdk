# ClaimableReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | **string** |  | 
**Epoch** | **int32** |  | 
**ClaimableRewards** | [**[]GetLBPairReward**](GetLBPairReward.md) |  | 

## Methods

### NewClaimableReward

`func NewClaimableReward(market string, epoch int32, claimableRewards []GetLBPairReward, ) *ClaimableReward`

NewClaimableReward instantiates a new ClaimableReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimableRewardWithDefaults

`func NewClaimableRewardWithDefaults() *ClaimableReward`

NewClaimableRewardWithDefaults instantiates a new ClaimableReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *ClaimableReward) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *ClaimableReward) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *ClaimableReward) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetEpoch

`func (o *ClaimableReward) GetEpoch() int32`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *ClaimableReward) GetEpochOk() (*int32, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *ClaimableReward) SetEpoch(v int32)`

SetEpoch sets Epoch field to given value.


### GetClaimableRewards

`func (o *ClaimableReward) GetClaimableRewards() []GetLBPairReward`

GetClaimableRewards returns the ClaimableRewards field if non-nil, zero value otherwise.

### GetClaimableRewardsOk

`func (o *ClaimableReward) GetClaimableRewardsOk() (*[]GetLBPairReward, bool)`

GetClaimableRewardsOk returns a tuple with the ClaimableRewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimableRewards

`func (o *ClaimableReward) SetClaimableRewards(v []GetLBPairReward)`

SetClaimableRewards sets ClaimableRewards field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


