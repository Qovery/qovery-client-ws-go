# ContainerStateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartedAt** | Pointer to **NullableInt32** | Unix timestamp with millisecond precision | [optional] 
**State** | [**ServiceStateDto**](ServiceStateDto.md) |  | 
**StateMessage** | Pointer to **NullableString** |  | [optional] 
**StateReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContainerStateDto

`func NewContainerStateDto(state ServiceStateDto, ) *ContainerStateDto`

NewContainerStateDto instantiates a new ContainerStateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerStateDtoWithDefaults

`func NewContainerStateDtoWithDefaults() *ContainerStateDto`

NewContainerStateDtoWithDefaults instantiates a new ContainerStateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartedAt

`func (o *ContainerStateDto) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ContainerStateDto) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ContainerStateDto) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ContainerStateDto) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *ContainerStateDto) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *ContainerStateDto) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetState

`func (o *ContainerStateDto) GetState() ServiceStateDto`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContainerStateDto) GetStateOk() (*ServiceStateDto, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContainerStateDto) SetState(v ServiceStateDto)`

SetState sets State field to given value.


### GetStateMessage

`func (o *ContainerStateDto) GetStateMessage() string`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *ContainerStateDto) GetStateMessageOk() (*string, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *ContainerStateDto) SetStateMessage(v string)`

SetStateMessage sets StateMessage field to given value.

### HasStateMessage

`func (o *ContainerStateDto) HasStateMessage() bool`

HasStateMessage returns a boolean if a field has been set.

### SetStateMessageNil

`func (o *ContainerStateDto) SetStateMessageNil(b bool)`

 SetStateMessageNil sets the value for StateMessage to be an explicit nil

### UnsetStateMessage
`func (o *ContainerStateDto) UnsetStateMessage()`

UnsetStateMessage ensures that no value is present for StateMessage, not even an explicit nil
### GetStateReason

`func (o *ContainerStateDto) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *ContainerStateDto) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *ContainerStateDto) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *ContainerStateDto) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### SetStateReasonNil

`func (o *ContainerStateDto) SetStateReasonNil(b bool)`

 SetStateReasonNil sets the value for StateReason to be an explicit nil

### UnsetStateReason
`func (o *ContainerStateDto) UnsetStateReason()`

UnsetStateReason ensures that no value is present for StateReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


