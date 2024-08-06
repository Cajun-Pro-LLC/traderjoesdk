# GetUserRewardHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Epoch** | **int32** |  | 
**EpochStart** | **int32** |  | 
**EpochEnd** | **int32** |  | 
**Progress** | **float32** |  | 
**Rewards** | [**[]GetLBPairReward**](GetLBPairReward.md) |  | 

## Methods

### NewGetUserRewardHistoryResponse

`func NewGetUserRewardHistoryResponse(epoch int32, epochStart int32, epochEnd int32, progress float32, rewards []GetLBPairReward, ) *GetUserRewardHistoryResponse`

NewGetUserRewardHistoryResponse instantiates a new GetUserRewardHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserRewardHistoryResponseWithDefaults

`func NewGetUserRewardHistoryResponseWithDefaults() *GetUserRewardHistoryResponse`

NewGetUserRewardHistoryResponseWithDefaults instantiates a new GetUserRewardHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpoch

`func (o *GetUserRewardHistoryResponse) GetEpoch() int32`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *GetUserRewardHistoryResponse) GetEpochOk() (*int32, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *GetUserRewardHistoryResponse) SetEpoch(v int32)`

SetEpoch sets Epoch field to given value.


### GetEpochStart

`func (o *GetUserRewardHistoryResponse) GetEpochStart() int32`

GetEpochStart returns the EpochStart field if non-nil, zero value otherwise.

### GetEpochStartOk

`func (o *GetUserRewardHistoryResponse) GetEpochStartOk() (*int32, bool)`

GetEpochStartOk returns a tuple with the EpochStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochStart

`func (o *GetUserRewardHistoryResponse) SetEpochStart(v int32)`

SetEpochStart sets EpochStart field to given value.


### GetEpochEnd

`func (o *GetUserRewardHistoryResponse) GetEpochEnd() int32`

GetEpochEnd returns the EpochEnd field if non-nil, zero value otherwise.

### GetEpochEndOk

`func (o *GetUserRewardHistoryResponse) GetEpochEndOk() (*int32, bool)`

GetEpochEndOk returns a tuple with the EpochEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochEnd

`func (o *GetUserRewardHistoryResponse) SetEpochEnd(v int32)`

SetEpochEnd sets EpochEnd field to given value.


### GetProgress

`func (o *GetUserRewardHistoryResponse) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *GetUserRewardHistoryResponse) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *GetUserRewardHistoryResponse) SetProgress(v float32)`

SetProgress sets Progress field to given value.


### GetRewards

`func (o *GetUserRewardHistoryResponse) GetRewards() []GetLBPairReward`

GetRewards returns the Rewards field if non-nil, zero value otherwise.

### GetRewardsOk

`func (o *GetUserRewardHistoryResponse) GetRewardsOk() (*[]GetLBPairReward, bool)`

GetRewardsOk returns a tuple with the Rewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewards

`func (o *GetUserRewardHistoryResponse) SetRewards(v []GetLBPairReward)`

SetRewards sets Rewards field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


