# PodStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | [**[]ContainerStatusDto**](ContainerStatusDto.md) |  | 
**Name** | **string** |  | 
**RestartCount** | **int32** |  | 
**ServiceVersion** | **string** |  | 
**StartedAt** | Pointer to **NullableInt32** | Unix timestamp with millisecond precision | [optional] 
**State** | [**ServiceStateDto**](ServiceStateDto.md) |  | 
**StateMessage** | **string** |  | 
**StateReason** | **string** |  | 

## Methods

### NewPodStatusDto

`func NewPodStatusDto(containers []ContainerStatusDto, name string, restartCount int32, serviceVersion string, state ServiceStateDto, stateMessage string, stateReason string, ) *PodStatusDto`

NewPodStatusDto instantiates a new PodStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodStatusDtoWithDefaults

`func NewPodStatusDtoWithDefaults() *PodStatusDto`

NewPodStatusDtoWithDefaults instantiates a new PodStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *PodStatusDto) GetContainers() []ContainerStatusDto`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *PodStatusDto) GetContainersOk() (*[]ContainerStatusDto, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *PodStatusDto) SetContainers(v []ContainerStatusDto)`

SetContainers sets Containers field to given value.


### GetName

`func (o *PodStatusDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PodStatusDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PodStatusDto) SetName(v string)`

SetName sets Name field to given value.


### GetRestartCount

`func (o *PodStatusDto) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *PodStatusDto) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *PodStatusDto) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.


### GetServiceVersion

`func (o *PodStatusDto) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *PodStatusDto) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *PodStatusDto) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.


### GetStartedAt

`func (o *PodStatusDto) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *PodStatusDto) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *PodStatusDto) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *PodStatusDto) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *PodStatusDto) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *PodStatusDto) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetState

`func (o *PodStatusDto) GetState() ServiceStateDto`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PodStatusDto) GetStateOk() (*ServiceStateDto, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PodStatusDto) SetState(v ServiceStateDto)`

SetState sets State field to given value.


### GetStateMessage

`func (o *PodStatusDto) GetStateMessage() string`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *PodStatusDto) GetStateMessageOk() (*string, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *PodStatusDto) SetStateMessage(v string)`

SetStateMessage sets StateMessage field to given value.


### GetStateReason

`func (o *PodStatusDto) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *PodStatusDto) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *PodStatusDto) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


